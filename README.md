# Run Locally

Before you begin, setup an OpenShift cluster and download a kube config file for it.

```sh
# Setup the environment
export KUBECONFIG=<pathToOpenShiftKubeConfig>
export POD_NAME=cluster-olm-operator-pod
export OPERATOR_NAME=cluster-olm-operator

# Create the OLM operator API and instance
kubectl apply -f ./vendor/github.com/openshift/api/operator/v1alpha1/0000_10_config-operator_01_olm.crd.yaml
kubectl apply -f ./olm-operator.yaml

# Create the cluster-olm-operator namesace (this is just where the lease will be created when running locally)
kubectl create namespace openshift-cluster-olm-operator

# Install cert-manager
kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.12.0/cert-manager.yaml

# Run the operator
go run ./ operator --namespace=openshift-cluster-olm-operator --kubeconfig "${KUBECONFIG}" --listen 127.0.0.1:8443
```

# TODO

1. Remove cert-manager dependency (use service-ca-operator instead)

   - The StaticResourceController can't handle the issuers and certificates present in the rukpak manifests, which means no cert secrets are created.
   - The deployments can't start because they are waiting on the certificate secrets to be created to fulfill their volume mounts.

