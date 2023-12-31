package v1alpha1

// This file contains a collection of methods that can be used from go-restful to
// generate Swagger API documentation for its models. Please read this PR for more
// information on the implementation: https://github.com/emicklei/go-restful/pull/215
//
// TODOs are ignored from the parser (e.g. TODO(andronat):... || TODO:...) if and only if
// they are on one line! For multiple line or blocks that you want to ignore use ---.
// Any context after a --- is ignored.
//
// Those methods can be generated by using hack/update-swagger-docs.sh

// AUTO-GENERATED FUNCTIONS START HERE
var map_DelegatedAuthentication = map[string]string{
	"":         "DelegatedAuthentication allows authentication to be disabled.",
	"disabled": "disabled indicates that authentication should be disabled.  By default it will use delegated authentication.",
}

func (DelegatedAuthentication) SwaggerDoc() map[string]string {
	return map_DelegatedAuthentication
}

var map_DelegatedAuthorization = map[string]string{
	"":         "DelegatedAuthorization allows authorization to be disabled.",
	"disabled": "disabled indicates that authorization should be disabled.  By default it will use delegated authorization.",
}

func (DelegatedAuthorization) SwaggerDoc() map[string]string {
	return map_DelegatedAuthorization
}

var map_GenerationHistory = map[string]string{
	"":               "GenerationHistory keeps track of the generation for a given resource so that decisions about forced updated can be made.",
	"group":          "group is the group of the thing you're tracking",
	"resource":       "resource is the resource type of the thing you're tracking",
	"namespace":      "namespace is where the thing you're tracking is",
	"name":           "name is the name of the thing you're tracking",
	"lastGeneration": "lastGeneration is the last generation of the workload controller involved",
}

func (GenerationHistory) SwaggerDoc() map[string]string {
	return map_GenerationHistory
}

var map_GenericOperatorConfig = map[string]string{
	"":               "GenericOperatorConfig provides information to configure an operator\n\nCompatibility level 4: No compatibility is provided, the API can change at any point for any reason. These capabilities should not be used by applications needing long term support.",
	"servingInfo":    "ServingInfo is the HTTP serving information for the controller's endpoints",
	"leaderElection": "leaderElection provides information to elect a leader. Only override this if you have a specific need",
	"authentication": "authentication allows configuration of authentication for the endpoints",
	"authorization":  "authorization allows configuration of authentication for the endpoints",
}

func (GenericOperatorConfig) SwaggerDoc() map[string]string {
	return map_GenericOperatorConfig
}

var map_LoggingConfig = map[string]string{
	"":        "LoggingConfig holds information about configuring logging",
	"level":   "level is passed to glog.",
	"vmodule": "vmodule is passed to glog.",
}

func (LoggingConfig) SwaggerDoc() map[string]string {
	return map_LoggingConfig
}

var map_NodeStatus = map[string]string{
	"":                               "NodeStatus provides information about the current state of a particular node managed by this operator.",
	"nodeName":                       "nodeName is the name of the node",
	"currentDeploymentGeneration":    "currentDeploymentGeneration is the generation of the most recently successful deployment",
	"targetDeploymentGeneration":     "targetDeploymentGeneration is the generation of the deployment we're trying to apply",
	"lastFailedDeploymentGeneration": "lastFailedDeploymentGeneration is the generation of the deployment we tried and failed to deploy.",
	"lastFailedDeploymentErrors":     "lastFailedDeploymentGenerationErrors is a list of the errors during the failed deployment referenced in lastFailedDeploymentGeneration",
}

func (NodeStatus) SwaggerDoc() map[string]string {
	return map_NodeStatus
}

var map_OperatorCondition = map[string]string{
	"": "OperatorCondition is just the standard condition fields.",
}

func (OperatorCondition) SwaggerDoc() map[string]string {
	return map_OperatorCondition
}

var map_OperatorSpec = map[string]string{
	"":                "OperatorSpec contains common fields for an operator to need.  It is intended to be anonymous included inside of the Spec struct for you particular operator.",
	"managementState": "managementState indicates whether and how the operator should manage the component",
	"imagePullSpec":   "imagePullSpec is the image to use for the component.",
	"imagePullPolicy": "imagePullPolicy specifies the image pull policy. One of Always, Never, IfNotPresent. Defaults to Always if :latest tag is specified, or IfNotPresent otherwise.",
	"version":         "version is the desired state in major.minor.micro-patch.  Usually patch is ignored.",
	"logging":         "logging contains glog parameters for the component pods.  It's always a command line arg for the moment",
}

func (OperatorSpec) SwaggerDoc() map[string]string {
	return map_OperatorSpec
}

var map_OperatorStatus = map[string]string{
	"":                           "OperatorStatus contains common fields for an operator to need.  It is intended to be anonymous included inside of the Status struct for you particular operator.",
	"observedGeneration":         "observedGeneration is the last generation change you've dealt with",
	"conditions":                 "conditions is a list of conditions and their status",
	"state":                      "state indicates what the operator has observed to be its current operational status.",
	"taskSummary":                "taskSummary is a high level summary of what the controller is currently attempting to do.  It is high-level, human-readable and not guaranteed in any way. (I needed this for debugging and realized it made a great summary).",
	"currentVersionAvailability": "currentVersionAvailability is availability information for the current version.  If it is unmanged or removed, this doesn't exist.",
	"targetVersionAvailability":  "targetVersionAvailability is availability information for the target version if we are migrating",
}

func (OperatorStatus) SwaggerDoc() map[string]string {
	return map_OperatorStatus
}

var map_StaticPodOperatorStatus = map[string]string{
	"":                                    "StaticPodOperatorStatus is status for controllers that manage static pods.  There are different needs because individual node status must be tracked.",
	"latestAvailableDeploymentGeneration": "latestAvailableDeploymentGeneration is the deploymentID of the most recent deployment",
	"nodeStatuses":                        "nodeStatuses track the deployment values and errors across individual nodes",
}

func (StaticPodOperatorStatus) SwaggerDoc() map[string]string {
	return map_StaticPodOperatorStatus
}

var map_VersionAvailability = map[string]string{
	"":                "VersionAvailability gives information about the synchronization and operational status of a particular version of the component",
	"version":         "version is the level this availability applies to",
	"updatedReplicas": "updatedReplicas indicates how many replicas are at the desired state",
	"readyReplicas":   "readyReplicas indicates how many replicas are ready and at the desired state",
	"errors":          "errors indicates what failures are associated with the operator trying to manage this version",
	"generations":     "generations allows an operator to track what the generation of \"important\" resources was the last time we updated them",
}

func (VersionAvailability) SwaggerDoc() map[string]string {
	return map_VersionAvailability
}

var map_ImageContentSourcePolicy = map[string]string{
	"":         "ImageContentSourcePolicy holds cluster-wide information about how to handle registry mirror rules. When multiple policies are defined, the outcome of the behavior is defined on each field.\n\nCompatibility level 4: No compatibility is provided, the API can change at any point for any reason. These capabilities should not be used by applications needing long term support.",
	"metadata": "metadata is the standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
	"spec":     "spec holds user settable values for configuration",
}

func (ImageContentSourcePolicy) SwaggerDoc() map[string]string {
	return map_ImageContentSourcePolicy
}

var map_ImageContentSourcePolicyList = map[string]string{
	"":         "ImageContentSourcePolicyList lists the items in the ImageContentSourcePolicy CRD.\n\nCompatibility level 4: No compatibility is provided, the API can change at any point for any reason. These capabilities should not be used by applications needing long term support.",
	"metadata": "metadata is the standard list's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
}

func (ImageContentSourcePolicyList) SwaggerDoc() map[string]string {
	return map_ImageContentSourcePolicyList
}

var map_ImageContentSourcePolicySpec = map[string]string{
	"":                        "ImageContentSourcePolicySpec is the specification of the ImageContentSourcePolicy CRD.",
	"repositoryDigestMirrors": "repositoryDigestMirrors allows images referenced by image digests in pods to be pulled from alternative mirrored repository locations. The image pull specification provided to the pod will be compared to the source locations described in RepositoryDigestMirrors and the image may be pulled down from any of the mirrors in the list instead of the specified repository allowing administrators to choose a potentially faster mirror. Only image pull specifications that have an image digest will have this behavior applied to them - tags will continue to be pulled from the specified repository in the pull spec.\n\nEach “source” repository is treated independently; configurations for different “source” repositories don’t interact.\n\nWhen multiple policies are defined for the same “source” repository, the sets of defined mirrors will be merged together, preserving the relative order of the mirrors, if possible. For example, if policy A has mirrors `a, b, c` and policy B has mirrors `c, d, e`, the mirrors will be used in the order `a, b, c, d, e`.  If the orders of mirror entries conflict (e.g. `a, b` vs. `b, a`) the configuration is not rejected but the resulting order is unspecified.",
}

func (ImageContentSourcePolicySpec) SwaggerDoc() map[string]string {
	return map_ImageContentSourcePolicySpec
}

var map_RepositoryDigestMirrors = map[string]string{
	"":        "RepositoryDigestMirrors holds cluster-wide information about how to handle mirros in the registries config. Note: the mirrors only work when pulling the images that are referenced by their digests.",
	"source":  "source is the repository that users refer to, e.g. in image pull specifications.",
	"mirrors": "mirrors is one or more repositories that may also contain the same images. The order of mirrors in this list is treated as the user's desired priority, while source is by default considered lower priority than all mirrors. Other cluster configuration, including (but not limited to) other repositoryDigestMirrors objects, may impact the exact order mirrors are contacted in, or some mirrors may be contacted in parallel, so this should be considered a preference rather than a guarantee of ordering.",
}

func (RepositoryDigestMirrors) SwaggerDoc() map[string]string {
	return map_RepositoryDigestMirrors
}

var map_OLM = map[string]string{
	"":         "OLM provides information to configure an operator to manage the OLM controllers\n\nCompatibility level 4: No compatibility is provided, the API can change at any point for any reason. These capabilities should not be used by applications needing long term support.",
	"metadata": "metadata is the standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
	"spec":     "spec holds user settable values for configuration",
	"status":   "status holds observed values from the cluster. They may not be overridden.",
}

func (OLM) SwaggerDoc() map[string]string {
	return map_OLM
}

var map_OLMList = map[string]string{
	"":         "OLMList is a collection of items\n\nCompatibility level 4: No compatibility is provided, the API can change at any point for any reason. These capabilities should not be used by applications needing long term support.",
	"metadata": "metadata is the standard list's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
	"items":    "Items contains the items",
}

func (OLMList) SwaggerDoc() map[string]string {
	return map_OLMList
}

// AUTO-GENERATED FUNCTIONS END HERE
