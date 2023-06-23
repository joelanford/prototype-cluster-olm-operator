package manifests

import (
	"embed"
	"io/fs"

	configv1 "github.com/openshift/api/config/v1"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/yaml"
)

//go:embed rukpak
var rukpakManifests embed.FS

//go:embed catalogd
var catalogdManifests embed.FS

//go:embed operator-controller
var operatorControllerManifests embed.FS

var RukpakAssets = Assets{manifests: rukpakManifests}
var CatalogdAssets = Assets{manifests: catalogdManifests}
var OperatorControllerAssets = Assets{manifests: operatorControllerManifests}

type Assets struct {
	manifests fs.FS
}

func (a Assets) Files(filter func(data []byte) bool) []string {
	var files []string
	if err := fs.WalkDir(a.manifests, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}

		data, err := a.Read(path)
		if err != nil {
			return err
		}
		if filter != nil && filter(data) {
			files = append(files, path)
		}
		return nil
	}); err != nil {
		panic(err)
	}
	return files
}

func (a Assets) Read(name string) ([]byte, error) {
	return fs.ReadFile(a.manifests, name)
}

func (a Assets) RelatedObjects(rm meta.RESTMapper) ([]configv1.ObjectReference, error) {
	files := a.Files(nil)
	relatedObjects := make([]configv1.ObjectReference, 0, len(files))
	for _, f := range files {
		data, err := a.Read(f)
		if err != nil {
			return nil, err
		}
		var u unstructured.Unstructured
		if err := yaml.Unmarshal(data, &u); err != nil {
			return nil, err
		}

		m, err := rm.RESTMapping(u.GroupVersionKind().GroupKind(), u.GroupVersionKind().Version)
		if err != nil {
			return nil, err
		}
		relatedObjects = append(relatedObjects, configv1.ObjectReference{
			Group:     m.GroupVersionKind.Group,
			Resource:  m.Resource.Resource,
			Namespace: u.GetNamespace(),
			Name:      u.GetName(),
		})
	}
	return relatedObjects, nil
}
