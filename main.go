package main

import (
	"context"
	goflag "flag"
	"fmt"
	"os"
	"regexp"
	"strings"

	_ "cloud.google.com/go/compute/metadata"
	configv1 "github.com/openshift/api/config/v1"
	"github.com/openshift/library-go/pkg/controller/controllercmd"
	"github.com/openshift/library-go/pkg/operator/deploymentcontroller"
	"github.com/openshift/library-go/pkg/operator/staticresourcecontroller"
	"github.com/openshift/library-go/pkg/operator/status"
	"github.com/openshift/library-go/pkg/operator/v1helpers"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apimachinery/pkg/version"
	"k8s.io/component-base/cli"
	utilflag "k8s.io/component-base/cli/flag"
	"sigs.k8s.io/yaml"

	"github.com/openshift/cluster-olm-operator/manifests"
	"github.com/openshift/cluster-olm-operator/pkg/clients"
)

func main() {
	pflag.CommandLine.SetNormalizeFunc(utilflag.WordSepNormalizeFunc)
	pflag.CommandLine.AddGoFlagSet(goflag.CommandLine)

	command := newRootCommand()
	code := cli.Run(command)
	os.Exit(code)
}

func newRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cluster-olm-operator",
		Short: "OpenShift cluster olm operator",
	}
	cmd.AddCommand(newOperatorCommand())
	return cmd
}

func newOperatorCommand() *cobra.Command {
	cmd := controllercmd.NewControllerCommandConfig(
		"cluster-olm-operator",
		version.Info{Major: "0", Minor: "0", GitVersion: "0.0.1"},
		runOperator,
	).NewCommandWithContext(context.Background())
	cmd.Use = "operator"
	cmd.Short = "Start the cluster olm operator"
	return cmd
}

func runOperator(ctx context.Context, cc *controllercmd.ControllerContext) error {
	cl, err := clients.New(cc)
	if err != nil {
		return err
	}

	allAssets := map[string]manifests.Assets{
		"Rukpak":             manifests.RukpakAssets,
		"Catalogd":           manifests.CatalogdAssets,
		"OperatorController": manifests.OperatorControllerAssets,
	}
	var (
		namespaces     = sets.New[string]()
		relatedObjects []configv1.ObjectReference
	)
	for _, componentAssets := range allAssets {
		componentObjects, err := componentAssets.RelatedObjects(cl.RESTMapper)
		if err != nil {
			return err
		}
		for _, obj := range componentObjects {
			namespaces.Insert(obj.Namespace)
		}
		relatedObjects = append(relatedObjects, componentObjects...)
	}

	kubeInformers := v1helpers.NewKubeInformersForNamespaces(cl.KubeClient, namespaces.UnsortedList()...)

	toStart := []func(<-chan struct{}){
		kubeInformers.Start,
		cl.ConfigInformerFactory.Start,
		cl.OperatorClient.Informer().Run,
	}

	var toRun []func(context.Context, int)

	noDeploymentsFilter := func(data []byte) bool {
		var u unstructured.Unstructured
		if err := yaml.Unmarshal(data, &u); err != nil {
			panic(err)
		}
		switch u.GetKind() {
		case "Deployment":
			return false
		}
		return true
	}

	deploymentFilter := func(data []byte) bool {
		var u unstructured.Unstructured
		if err := yaml.Unmarshal(data, &u); err != nil {
			panic(err)
		}
		switch u.GetKind() {
		case "Deployment":
			return true
		}
		return false
	}

	clientHolder := cl.ClientHolder().WithKubernetesInformers(kubeInformers)
	for name, componentAssets := range allAssets {
		c := staticresourcecontroller.NewStaticResourceController(
			fmt.Sprintf("StaticResources%s", name),
			componentAssets.Read,
			componentAssets.Files(noDeploymentsFilter),
			clientHolder,
			cl.OperatorClient,
			cc.EventRecorder.ForComponent(name),
		).AddKubeInformers(kubeInformers)
		toRun = append(toRun, c.Run)

		for _, file := range componentAssets.Files(deploymentFilter) {
			m, err := componentAssets.Read(file)
			if err != nil {
				return err
			}

			var u unstructured.Unstructured
			if err := yaml.Unmarshal(m, &u); err != nil {
				return err
			}

			baseName := fmt.Sprintf("Deployment%s", capName(u.GetName()))
			d := deploymentcontroller.NewDeploymentController(
				baseName,
				m,
				cc.EventRecorder.ForComponent(baseName),
				cl.OperatorClient,
				cl.KubeClient,
				cl.KubeInformerFactory.Apps().V1().Deployments(),
				nil,
				nil,
			)
			toRun = append(toRun, d.Run)
		}
	}
	toStart = append(toStart, cl.KubeInformerFactory.Start)

	versionGetter := status.NewVersionGetter()
	clusterOperatorCtrl := status.NewClusterOperatorStatusController(
		"cluster-olm-operator",
		relatedObjects,
		cl.ConfigClient.ConfigV1(),
		cl.ConfigInformerFactory.Config().V1().ClusterOperators(),
		cl.OperatorClient,
		versionGetter,
		cc.EventRecorder.ForComponent("cluster-olm-operator"),
	)
	toRun = append(toRun, clusterOperatorCtrl.Run)

	for _, start := range toStart {
		go start(ctx.Done())
	}
	for _, run := range toRun {
		go run(ctx, 1)
	}

	<-ctx.Done()
	return nil
}

func capName(in string) string {
	return regexp.MustCompile(`(^[a-z])|(\-[a-z])`).ReplaceAllStringFunc(in, func(s string) string { return strings.TrimPrefix(strings.ToUpper(s), "-") })
}
