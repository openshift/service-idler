// TODO: add apache license boilerplate here

// Api versions allow the api contract for a resource to be changed while keeping
// backward compatibility by support multiple concurrent versions
// of the same resource

// +k8s:openapi-gen=true
// +k8s:deepcopy-gen=package,register
// +k8s:conversion-gen=github.com/openshift/origin-idler/pkg/apis/idling
// +k8s:defaulter-gen=TypeMeta
// +groupName=idling.openshift.io
package v1alpha2 // import "github.com/openshift/origin-idler/pkg/apis/idling/v1alpha2"
