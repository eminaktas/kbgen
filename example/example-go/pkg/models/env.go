package models

type EnvMap map[string]*Env

type EnvValueFrom struct {
	// +optional
	SecretKeyRef *ObjectKeySelector `json:"secretKeyRef,omitempty" yaml:"secretKeyRef,omitempty"`
	// +optional
	ConfigMapKeyRef *ObjectKeySelector `json:"configMapKeyRef,omitempty" yaml:"configMapKeyRef,omitempty"`
	// +optional
	ResourceFieldRef *ResourceFieldSelector `json:"resourceFieldRef,omitempty" yaml:"resourceFieldRef,omitempty"`
	// +optional
	FieldRef *ObjectFieldSelector `json:"fieldRef,omitempty" yaml:"fieldRef,omitempty"`
}

func (in *EnvValueFrom) DeepCopyInto(out *EnvValueFrom) {
	*out = *in
}

func (in *EnvValueFrom) DeepCopy() *EnvValueFrom {
	if in == nil {
		return nil
	}
	out := new(EnvValueFrom)
	in.DeepCopyInto(out)
	return out
}

type ResourceFieldSelector struct {
	// +required
	Resource string `json:"resource" yaml:"resource"`
	// +optional
	ContainerName string `json:"containerName,omitempty" yaml:"containerName,omitempty"`
	// +kubebuilder:pruning:PreserveUnknownFields
	// +optional
	Divisor *int `json:"divisor,omitempty" yaml:"divisor,omitempty"`
}

func (in *ResourceFieldSelector) DeepCopyInto(out *ResourceFieldSelector) {
	*out = *in
}

func (in *ResourceFieldSelector) DeepCopy() *ResourceFieldSelector {
	if in == nil {
		return nil
	}
	out := new(ResourceFieldSelector)
	in.DeepCopyInto(out)
	return out
}

type ObjectKeySelector struct {
	// +required
	Key string `json:"key" yaml:"key"`
	// +required
	Name string `json:"name" yaml:"name"`
}

func (in *ObjectKeySelector) DeepCopyInto(out *ObjectKeySelector) {
	*out = *in
}

func (in *ObjectKeySelector) DeepCopy() *ObjectKeySelector {
	if in == nil {
		return nil
	}
	out := new(ObjectKeySelector)
	in.DeepCopyInto(out)
	return out
}

type ObjectFieldSelector struct {
	// +optional
	ApiVersion string `json:"apiVersion,omitempty" yaml:"apiVersion,omitempty"`
	// +required
	FieldPath string `json:"fieldPath" yaml:"fieldPath"`
}

func (in *ObjectFieldSelector) DeepCopyInto(out *ObjectFieldSelector) {
	*out = *in
}

func (in *ObjectFieldSelector) DeepCopy() *ObjectFieldSelector {
	if in == nil {
		return nil
	}
	out := new(ObjectFieldSelector)
	in.DeepCopyInto(out)
	return out
}

type EnvFromSource struct {
	// +optional
	SecretRef string `json:"secretRef,omitempty" yaml:"secretRef,omitempty"`
	// +optional
	ConfigMapRef string `json:"configMapRef,omitempty" yaml:"configMapRef,omitempty"`
}

func (in *EnvFromSource) DeepCopyInto(out *EnvFromSource) {
	*out = *in
}

func (in *EnvFromSource) DeepCopy() *EnvFromSource {
	if in == nil {
		return nil
	}
	out := new(EnvFromSource)
	in.DeepCopyInto(out)
	return out
}

type Env struct {
	// +optional
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
	// +required
	Name string `json:"name" yaml:"name"`
	// +optional
	ValueFrom *EnvValueFrom `json:"valueFrom,omitempty" yaml:"valueFrom,omitempty"`
}

func (in *Env) DeepCopyInto(out *Env) {
	*out = *in
}

func (in *Env) DeepCopy() *Env {
	if in == nil {
		return nil
	}
	out := new(Env)
	in.DeepCopyInto(out)
	return out
}
