package models

type Server struct {
	// +kubebuilder:validation:Required
	Replicas *int `json:"replicas" yaml:"replicas"`
	// +kubebuilder:validation:Optional
	Services []*Service `json:"services,omitempty" yaml:"services,omitempty"`
	// +kubebuilder:validation:Enum=Deployment;StatefulSet
	// +kubebuilder:validation:Required
	WorkloadType string `json:"workloadType" yaml:"workloadType"`
	// +kubebuilder:validation:Optional
	Volumes []*Volume `json:"volumes,omitempty" yaml:"volumes,omitempty"`
	// +kubebuilder:validation:Required
	MainContainer *Container `json:"mainContainer" yaml:"mainContainer"`
	// +kubebuilder:validation:Required
	Image string `json:"image" yaml:"image"`
	// +kubebuilder:validation:Optional
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	// +kubebuilder:validation:Optional
	Annotations map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	// +kubebuilder:validation:Enum=Always;IfNotPresent;Never
	// +kubebuilder:validation:Optional
	ImagePullPolicy string `json:"imagePullPolicy,omitempty" yaml:"imagePullPolicy,omitempty"`
	// +kubebuilder:validation:Optional
	Namespace string `json:"namespace,omitempty" yaml:"namespace,omitempty"`
	// +kubebuilder:validation:Optional
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
}

func (in *Server) DeepCopyInto(out *Server) {
	*out = *in
}

type Container struct {
	// +kubebuilder:validation:Optional
	Image string `json:"image,omitempty" yaml:"image,omitempty"`
	// +kubebuilder:validation:Enum=Always;IfNotPresent;Never
	// +kubebuilder:validation:Optional
	ImagePullPolicy string `json:"imagePullPolicy,omitempty" yaml:"imagePullPolicy,omitempty"`

	Main `json:",inline" yaml:",inline"`
}
