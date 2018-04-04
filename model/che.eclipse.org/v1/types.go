package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/api/core/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type CheService struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +optional
	Spec CheServiceSpec `json:"spec,omitempty"`
}

type Version struct {
	Version string `json:"version"`
}

type CheServiceSpec struct {
	Services []v1.Service `json:"services,omitempty"`
	//
	Pods []v1.Pod `json:"pods,omitempty"`

	Commands []CheCommand `json:"commands,omitempty"`

	Version
}

type CheServiceSpecTesst struct {
	Commands []CheCommand `json:"commands,omitempty"`
}

type CheCommand struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +optional
	Spec CheCommandSpec `json:"spec,omitempty"`
}

type CheCommandSpec struct {
	TargetLabelSelector string `json:"target-label-selector"`

	WorkingDirectory string `json:"working-dir,omitempty"`

	Commands []string `json:"command"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type CheFeature struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +optional
	Spec CheFeatureSpec `json:"spec,omitempty"`
}

type CheFeatureSpec struct {
	Version
	Services []CheServiceReference `json:"services"`
}

type CheServiceReference struct {
	Name       string                `json:"name"`
	Version    string                `json:"version"`
	Parameters []CheServiceParameter `json:"parameters,omitempty"`
}

type CheServiceParameter struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
