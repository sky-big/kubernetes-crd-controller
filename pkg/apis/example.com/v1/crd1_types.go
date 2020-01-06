package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type CRD1 struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +optional
	Spec CRD1Spec `json:"spec,omitempty"`
	// +optional
	Status CRD1Status `json:"status,omitempty"`
}

type CRD1Spec struct {
	// +optional
	Generation int64 `json:"generation,omitempty"`

	// +optional
	V1 int64 `json:"v1,omitempty"`
}

type CRD1Status struct {
	// +optional
	Phase int64 `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type CRD1List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []CRD1 `json:"items"`
}
