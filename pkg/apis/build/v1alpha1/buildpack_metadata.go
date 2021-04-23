package v1alpha1

type BuildpackMetadataList []BuildpackMetadata

// +k8s:openapi-gen=true
type BuildpackMetadata struct {
	Id       string `json:"id"`
	Version  string `json:"version"`
	Homepage string `json:"homepage,omitempty"`
}
