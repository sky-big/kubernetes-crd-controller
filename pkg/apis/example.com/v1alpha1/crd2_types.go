package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type CRD2 struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +optional
	Spec CRD2Spec `json:"spec,omitempty"`
	// +optional
	Status CRD2Status `json:"status,omitempty"`
}

type CRD2Spec struct {
	// +optional
	Generation int64 `json:"generation,omitempty"`

	// +optional
	V1alpha1 int64 `json:"v1alpha1,omitempty"`
}

type CRD2Status struct {
	// +optional
	Phase int64 `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type CRD2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []CRD2 `json:"items"`
}
