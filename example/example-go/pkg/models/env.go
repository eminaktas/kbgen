package models

type EnvMap map[string]*Env

type EnvValueFrom struct {
	// +kubebuilder:validation:Optional
	FieldRef *ObjectFieldSelector `json:"fieldRef,omitempty" yaml:"fieldRef,omitempty"`
	// +kubebuilder:validation:Optional
	ConfigMapKeyRef *ObjectKeySelector `json:"configMapKeyRef,omitempty" yaml:"configMapKeyRef,omitempty"`
	// +kubebuilder:validation:Optional
	ResourceFieldRef *ResourceFieldSelector `json:"resourceFieldRef,omitempty" yaml:"resourceFieldRef,omitempty"`
	// +kubebuilder:validation:Optional
	SecretKeyRef *ObjectKeySelector `json:"secretKeyRef,omitempty" yaml:"secretKeyRef,omitempty"`
}

type ResourceFieldSelector struct {
	// +kubebuilder:pruning:PreserveUnknownFields
	// +kubebuilder:validation:Optional
	Divisor *int `json:"divisor,omitempty" yaml:"divisor,omitempty"`
	// +kubebuilder:validation:Required
	Resource string `json:"resource" yaml:"resource"`
	// +kubebuilder:validation:Optional
	ContainerName string `json:"containerName,omitempty" yaml:"containerName,omitempty"`
}

type ObjectKeySelector struct {
	// +kubebuilder:validation:Required
	Name string `json:"name" yaml:"name"`
	// +kubebuilder:validation:Required
	Key string `json:"key" yaml:"key"`
}

type EnvFromSource struct {
	// +kubebuilder:validation:Optional
	ConfigMapRef string `json:"configMapRef,omitempty" yaml:"configMapRef,omitempty"`
	// +kubebuilder:validation:Optional
	SecretRef string `json:"secretRef,omitempty" yaml:"secretRef,omitempty"`
}

type Env struct {
	// +kubebuilder:validation:Optional
	ValueFrom *EnvValueFrom `json:"valueFrom,omitempty" yaml:"valueFrom,omitempty"`
	// +kubebuilder:validation:Required
	Name string `json:"name" yaml:"name"`
	// +kubebuilder:validation:Optional
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}

type ObjectFieldSelector struct {
	// +kubebuilder:validation:Optional
	ApiVersion string `json:"apiVersion,omitempty" yaml:"apiVersion,omitempty"`
	// +kubebuilder:validation:Required
	FieldPath string `json:"fieldPath" yaml:"fieldPath"`
}
