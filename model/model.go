package model
import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/api/core/v1"
)

type CheService struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +optional
	Spec CheServiceSpec `json:"spec,omitempty"`

}


type CheServiceSpec struct {
	Services []v1.Service `json:"services,omitempty"`

	Pods []v1.Pod `json:"pods,omitempty"`
}