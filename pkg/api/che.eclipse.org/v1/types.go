package v1

type CheService struct {
	TypeMeta   `json:",inline"`
	ObjectMeta `json:"metadata,omitempty"`

	Spec CheServiceSpec `json:"spec,omitempty"`
}

type TypeMeta struct {
}

type ObjectMeta struct {
}

type Version struct {
	Version string `json:"version"`
}

type CheServiceSpec struct {
	Services []Service `json:"services,omitempty"`
	//
	Pods []Pod `json:"pods,omitempty"`

	Commands []CheCommand `json:"commands,omitempty"`

	Version
}

type Service struct {
}

type Pod struct {
}

type CheCommand struct {
	TypeMeta   `json:",inline"`
	ObjectMeta `json:"metadata,omitempty"`

	// +optional
	Spec CheCommandSpec `json:"spec,omitempty"`
}

type CheCommandSpec struct {
	TargetLabelSelector string `json:"target-label-selector"`

	WorkingDirectory string `json:"working-dir,omitempty"`

	Commands []string `json:"command"`
}

type CheFeature struct {
	TypeMeta   `json:",inline"`
	ObjectMeta `json:"metadata,omitempty"`

	// +optional
	Spec CheFeatureSpec `json:"spec,omitempty"`
}

type CheFeatureSpec struct {
	Version
	Services []CheServiceReference `json:"services"`
}

type CheServiceReference struct {
	Name       string                `json:"name"`
	Version    string                `json:"version"`
	Parameters []CheServiceParameter `json:"parameters,omitempty"`
}

type CheServiceParameter struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
