package models

type ContainerPort struct {
	// +kubebuilder:validation:Optional
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	// +kubebuilder:validation:Enum=TCP;UDP;SCTP
	// +kubebuilder:validation:Optional
	Protocol string `json:"protocol,omitempty" yaml:"protocol,omitempty"`
	// +kubebuilder:validation:Optional
	ContainerPort *int `json:"containerPort,omitempty" yaml:"containerPort,omitempty"`
}
