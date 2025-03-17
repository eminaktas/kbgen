package models

func (in *Container) DeepCopyInto(out *Container) {
	*out = *in
}

func (in *Container) DeepCopy() *Container {
	if in == nil {
		return nil
	}
	out := new(Container)
	in.DeepCopyInto(out)
	return out
}

func (in *Server) DeepCopyInto(out *Server) {
	*out = *in
}

func (in *Server) DeepCopy() *Server {
	if in == nil {
		return nil
	}
	out := new(Server)
	in.DeepCopyInto(out)
	return out
}

type Container struct {
	// +optional
	Image string `json:"image,omitempty" yaml:"image,omitempty"`
	// +kubebuilder:validation:Enum=Always;IfNotPresent;Never
	// +optional
	ImagePullPolicy string `json:"imagePullPolicy,omitempty" yaml:"imagePullPolicy,omitempty"`

	Main `json:",inline" yaml:",inline"`
}

type Server struct {
	// +optional
	Annotations map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	// +required
	Image string `json:"image" yaml:"image"`
	// +kubebuilder:validation:Enum=Always;IfNotPresent;Never
	// +optional
	ImagePullPolicy string `json:"imagePullPolicy,omitempty" yaml:"imagePullPolicy,omitempty"`
	// +optional
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	// +required
	MainContainer *Container `json:"mainContainer" yaml:"mainContainer"`
	// +optional
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	// +optional
	Namespace string `json:"namespace,omitempty" yaml:"namespace,omitempty"`
	// +required
	Replicas *int `json:"replicas" yaml:"replicas"`
	// +optional
	Services []*Service `json:"services,omitempty" yaml:"services,omitempty"`
	// +optional
	Volumes []*Volume `json:"volumes,omitempty" yaml:"volumes,omitempty"`
	// +kubebuilder:validation:Enum=Deployment;StatefulSet
	// +required
	WorkloadType string `json:"workloadType" yaml:"workloadType"`
}
