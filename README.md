# Kubernetes CRD Controller

## Overview

Kubernetes CRD Controller is to manage custom resources on the Kubernetes cluster.

## Quick Start

### Deploy Kubernetes CRD Controller

1. Clone the project on your Kubernetes cluster master node:
```
$ git clone https://github.com/sky-big/kubernetes-crd-controller
$ cd kubernetes-crd-controller
```

2. To deploy the Kubernetes CRD Controller on your Kubernetes cluster, please run the following script:
```
$ make install
```

3. Use command ```kubectl get pods``` to check Pulsar Operator deploy status like:
```
$ kubectl get pods
NAME                                         READY   STATUS    RESTARTS   AGE
kubernetes-crd-controller-86d5f5cbd7-6kq2x   1/1     Running   0          40m
```

Now you can use the CRDs provide by Kubernetes CRD Controller to deploy your Pulsar Cluster.

### Define Custom Resource

1. example [define crd1 custom resource](./pkg/apis/example/v1alpha1/crd1_types.go)

```
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
	V1alpha1 int64 `json:"v1alpha1,omitempty"`
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
```

1. example [define crd2 custom resource](./pkg/apis/example/v1alpha1/crd2_types.go)

```
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
```

### Create Custom Resource

1. example [create custom resource](./deploy/crds/crs.yaml):

```
apiVersion: example.dev/v1alpha1
kind: CRD1
metadata:
  name: test-crd1
spec:
  v1alpha1: 1

---

apiVersion: example.dev/v1alpha1
kind: CRD2
metadata:
  name: test-crd2
spec:
  v1alpha1: 2
```

## Generate Custom Resource SDK

```
$ make generate
```

## Generate Image

```
$ make image
```

## Push Image

```
$ make push
```

## Generate Project Vendor

```
$ make vendor
```
