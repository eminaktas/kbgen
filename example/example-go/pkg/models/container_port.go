package models

type ContainerPort struct {
	// +optional
	ContainerPort *int `json:"containerPort,omitempty" yaml:"containerPort,omitempty"`
	// +optional
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	// +kubebuilder:validation:Enum=TCP;UDP;SCTP
	// +optional
	Protocol string `json:"protocol,omitempty" yaml:"protocol,omitempty"`
}

func (in *ContainerPort) DeepCopyInto(out *ContainerPort) {
	*out = *in
}

func (in *ContainerPort) DeepCopy() *ContainerPort {
	if in == nil {
		return nil
	}
	out := new(ContainerPort)
	in.DeepCopyInto(out)
	return out
}
