package types

//TypeMeta describes an individual object in an API response or request with strings representing the type of the
// object and its API schema version. Structures that are versioned or persisted should inline TypeMeta
type TypeMeta struct {
	APIVersion string `json:"apiVersion,omitempty" yaml:"apiVersion,omitempty" `
	Kind       string `json:"kind,omitempty" yaml:"kind,omitempty"`
}

type ObjectMeta struct {
	// Object name. Name must be unique.
	Name string `json:"name,omitempty"`
	// Map of string keys and values that can be used to organize and categorize
	// (scope and select) objects. May match selectors of replication controllers
	// and services.
	Labels map[string]string `json:"labels,omitempty"`
}

type Version struct {
	Version string `json:"version"`
}

type CheService struct {
	TypeMeta   `json:",inline"`
	ObjectMeta `json:"metadata,omitempty"`

	Spec CheServiceSpec `json:"spec,omitempty"`
}

type CheServiceSpec struct {
	Version

	// List of containers belonging to the service.
	Containers []Container `json:"containers"`
}

// A single application container that you want to run within a pod.
type Container struct {
	// Docker image name.
	Image string `json:"image"`

	// List of environment variables to set in the container. Cannot be updated.
	Env []EnvVar `json:"env,omitempty"`

	Resources ResourceRequirements `json:"resources,omitempty"`

	// List of container commands
	Commands []Command `json:"commands,omitempty"`

	// List of container servers
	Servers []Server `json:"servers"`

	// List of container volumes
	Volumes []Volume `json:"volumes,omitempty"`
}

type Command struct {
	Name string `json:"name"`

	WorkingDir string `json:"working-dir"`

	Command []string `json:"command"`
}

type Volume struct {
	Name string `json:"name"`

	//path of the volume in running container
	MountPath string `json:"mountPath"`
}

type Server struct {
	Name string `json:"name"`

	Port int32 `json:"port"`

	Protocol string `json:"protocol"`

	Attributes map[string]string `json:"attributes,omitempty"`
}

// ResourceRequirements describes the compute resource requirements.
type ResourceRequirements struct {
	// Requests describes the minimum amount of compute resources required.
	Requests ResourceList `json:"requests,omitempty"`
}

// ResourceList is a set of (resource name, quantity) pairs.
type ResourceList map[string]string

// EnvVar represents an environment variable present in a Container.
type EnvVar struct {
	// Name of the environment variable. Must be a C_IDENTIFIER.
	Name string `json:"name"`

	// Optional: no more than one of the following may be specified.

	// Variable references $(VAR_NAME) are expanded
	// using the previous defined environment variables in the container and
	// any service environment variables. If a variable cannot be resolved,
	// the reference in the input string will be unchanged. The $(VAR_NAME)
	// syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped
	// references will never be expanded, regardless of whether the variable
	// exists or not.
	// Defaults to "".
	Value string `json:"value,omitempty"`
}

type CheFeature struct {
	TypeMeta   `json:",inline"`
	ObjectMeta `json:"metadata,omitempty"`

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
