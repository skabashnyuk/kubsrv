package v1

type CheService struct {
	TypeMeta   `json:",inline"`
	ObjectMeta `json:"metadata,omitempty"`

	Spec CheServiceSpec `json:"spec,omitempty"`
}

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

type CheServiceSpec struct {
	Services []Service `json:"services,omitempty"`
	//
	Pods []Pod `json:"pods,omitempty"`

	Commands []CheCommand `json:"commands,omitempty"`

	Version
}

type Service struct {
	TypeMeta   `json:",inline"`
	ObjectMeta `json:"metadata,omitempty"`
	// Spec defines the behavior of a service.
	Spec ServiceSpec `json:"spec,omitempty"`
}

// ServiceSpec describes the attributes that a user creates on a service.
type ServiceSpec struct {
	// The list of ports that are exposed by this service.
	Ports []ServicePort `json:"ports,omitempty"`

	// Route service traffic to pods with label keys and values matching this
	// selector. If empty or not present, the service is assumed to have an
	// external process managing its endpoints, which Kubernetes will not
	// modify. Only applies to types ClusterIP, NodePort, and LoadBalancer.
	// Ignored if type is ExternalName.
	// More info: https://kubernetes.io/docs/concepts/services-networking/service/
	// +optional
	Selector map[string]string `json:"selector,omitempty"`
}

// ServicePort contains information on service's port.
type ServicePort struct {
	// The name of this port within the service. This must be a DNS_LABEL.
	// All ports within a ServiceSpec must have unique names. This maps to
	// the 'Name' field in EndpointPort objects.
	// Optional if only one ServicePort is defined on this service.
	// +optional
	Name string `json:"name,omitempty"`

	// The IP protocol for this port. Supports "TCP" and "UDP".
	// Default is TCP.
	// +optional
	Protocol Protocol `json:"protocol,omitempty"`

	// The port that will be exposed by this service.
	Port int32 `json:"port"`

	// Number or name of the port to access on the pods targeted by the service.
	// Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.
	// If this is a string, it will be looked up as a named port in the
	// target Pod's container ports. If this is not specified, the value
	// of the 'port' field is used (an identity map).
	// This field is ignored for services with clusterIP=None, and should be
	// omitted or set equal to the 'port' field.
	// More info: https://kubernetes.io/docs/concepts/services-networking/service/#defining-a-service
	// +optional
	TargetPort int32 `json:"targetPort,omitempty"`


}

// Protocol defines network protocols supported for things like container ports.
type Protocol string

const (
	// ProtocolTCP is the TCP protocol.
	ProtocolTCP Protocol = "TCP"
	// ProtocolUDP is the UDP protocol.
	ProtocolUDP Protocol = "UDP"
)

// Pod is a collection of containers that can run on a host. This resource is created
// by clients and scheduled onto hosts.
type Pod struct {
	TypeMeta   `json:",inline"`
	ObjectMeta `json:"metadata,omitempty"`
	// Specification of the desired behavior of the pod.
	Spec PodSpec `json:"spec,omitempty"`
}

type PodSpec struct {
	// List of volumes that can be mounted by containers belonging to the pod.
	Volumes []Volume `json:"volumes,omitempty"`

	// List of containers belonging to the pod.
	Containers []Container `json:"containers"`
}

type Volume struct {
	// Volume's name.
	Name string `json:"name" `
	// VolumeSource represents the location and type of the mounted volume.
	// If not specified, the Volume is implied to be an EmptyDir.
	// This implied behavior is deprecated and will be removed in a future version.
	VolumeSource `json:",inline"`
}

type VolumeSource struct {
	// PersistentVolumeClaimVolumeSource represents a reference to a
	// PersistentVolumeClaim in the same namespace.
	PersistentVolumeClaim *PersistentVolumeClaimVolumeSource `json:"persistentVolumeClaim,omitempty"`
}

// PersistentVolumeClaimVolumeSource references the user's PVC in the same namespace.
// This volume finds the bound PV and mounts that volume for the pod. A
// PersistentVolumeClaimVolumeSource is, essentially, a wrapper around another
// type of volume that is owned by someone else (the system).
type PersistentVolumeClaimVolumeSource struct {
	// ClaimName is the name of a PersistentVolumeClaim in the same namespace as the pod using this volume.
	// More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims
	ClaimName string `json:"claimName"`
	// Will force the ReadOnly setting in VolumeMounts.
	// Default false.
	// +optional
	ReadOnly bool `json:"readOnly,omitempty"`
}



// A single application container that you want to run within a pod.
type Container struct {

	// Name of the container specified as a DNS_LABEL.
	// Each container in a pod must have a unique name (DNS_LABEL).
	// Cannot be updated.
	Name string `json:"name"`
	// Docker image name.
	// More info: https://kubernetes.io/docs/concepts/containers/images
	// This field is optional to allow higher level config management to default or override
	// container images in workload controllers like Deployments and StatefulSets.
	// +optional
	Image string `json:"image,omitempty"`

	// List of environment variables to set in the container.

	Env []EnvVar `json:"env,omitempty"`
	// Pod volumes to mount into the container's filesystem.

	VolumeMounts []VolumeMount `json:"volumeMounts,omitempty"`
}

// EnvVar represents an environment variable present in a Container.
type EnvVar struct {
	// Name of the environment variable. Must be a C_IDENTIFIER.
	Name string `json:"name" protobuf:"bytes,1,opt,name=name"`

	// Optional: no more than one of the following may be specified.

	// Variable references $(VAR_NAME) are expanded
	// using the previous defined environment variables in the container and
	// any service environment variables. If a variable cannot be resolved,
	// the reference in the input string will be unchanged. The $(VAR_NAME)
	// syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped
	// references will never be expanded, regardless of whether the variable
	// exists or not.
	// Defaults to "".
	// +optional
	Value string `json:"value,omitempty" protobuf:"bytes,2,opt,name=value"`
}

// VolumeMount describes a mounting of a Volume within a container.
type VolumeMount struct {
	// This must match the Name of a Volume.
	Name string `json:"name" protobuf:"bytes,1,opt,name=name"`
	// Mounted read-only if true, read-write otherwise (false or unspecified).
	// Defaults to false.
	// +optional
	ReadOnly bool `json:"readOnly,omitempty" protobuf:"varint,2,opt,name=readOnly"`
	// Path within the container at which the volume should be mounted.  Must
	// not contain ':'.
	MountPath string `json:"mountPath" protobuf:"bytes,3,opt,name=mountPath"`
	// Path within the volume from which the container's volume should be mounted.
	// Defaults to "" (volume's root).
	// +optional
	SubPath string `json:"subPath,omitempty" protobuf:"bytes,4,opt,name=subPath"`
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
