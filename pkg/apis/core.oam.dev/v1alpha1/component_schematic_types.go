package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PortProtocol string

const (
	TCP PortProtocol = "TCP"
	UDP              = "UDP"
)

type WorkloadSetting struct {
	Name          string        `json:"name"`
	Description   string        `json:"description"`
	ParameterType ParameterType `json:"type"`
	Required      bool          `json:"required"`
	Default       string        `json:"default"`
	FromParam     string        `json:"fromParam"`
}

/// CPU describes a CPU resource allocation for a container.
///
/// The minimum number of logical cpus required for running this container.
type CPU struct {
	Required float64 `json:"required"`
}

/// Memory describes the memory allocation for a container.
///
/// The minimum amount of memory in MB required for running this container. The value should be a positive integer, greater than zero.
type Memory struct {
	Required string `json:"required"`
}

/// GPU describes a Container's need for a GPU.
///
/// The minimum number of gpus required for running this container.
type GPU struct {
	Required float64 `json:"required"`
}

/// Volume describes a path that is attached to a Container.
///
/// It specifies not only the location, but also the requirements.
type Volume struct {
	Name          string        `json:"name"`
	MountPath     string        `json:"mountPath"`
	AccessMode    AccessMode    `json:"accessMode"`
	SharingPolicy SharingPolicy `json:"sharingPolicy"`
	Disk          Disk          `json:"disk"`
}

/// AccessMode defines the access modes for file systems.
///
/// Currently, only read/write and read-only are supported.
type AccessMode string

const (
	RW AccessMode = "RW"
	RO AccessMode = "RO"
)

/// SharingPolicy defines whether a filesystem can be shared across containers.
///
/// An Exclusive filesystem can only be attached to one container.
type SharingPolicy string

const (
	Shared    SharingPolicy = "Shared"
	Exclusive SharingPolicy = "Exclusive"
)

// Disk describes the disk requirements for backing a Volume.
type Disk struct {
	Required  string `json:"required"`
	Ephemeral bool   `json:"ephemeral"`
}

type ExtendedResource struct {
	Name     string `json:"name"`
	Required string `json:"required"`
}

type Resources struct {
	Cpu      CPU                `json:"cpu"`
	Memory   Memory             `json:"memory"`
	Gpu      GPU                `json:"gpu"`
	Volumes  []Volume           `json:"volumes"`
	Extended []ExtendedResource `json:"extended"`
}

type Env struct {
	Name      string `json:"name"`
	Value     string `json:"value"`
	FromParam string `json:"fromParam"`
}

type Port struct {
	Name          string       `json:"name"`
	ContainerPort int32        `json:"port"`
	Protocol      PortProtocol `json:"protocol"`
}

type Exec struct {
	Command []string `json:"command"`
}

type HttpHeader struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type HttpGet struct {
	Path        string       `json:"path"`
	Port        int32        `json:"port"`
	HttpHeaders []HttpHeader `json:"httpHeaders"`
}

type TcpSocket struct {
	Port int32 `json:"port"`
}

type HealthProbe struct {
	Exec                Exec      `json:"exec"`
	HttpGet             HttpGet   `json:"httpGet"`
	TcpSocket           TcpSocket `json:"tcpSocket"`
	InitialDelaySeconds int32     `json:"initialDelaySeconds"`
	PeriodSeconds       int32     `json:"periodSeconds"`
	TimeoutSeconds      int32     `json:"timeoutSeconds"`
	SuccessThreshold    int32     `json:"successThreshold"`
	FailureThreshold    int32     `json:"failureThreshold"`
}

type ConfigFile struct {
	Path      string `json:"path"`
	Value     string `json:"value"`
	FromParam string `json:"fromParam"`
}

type Container struct {
	Name            string       `json:"name"`
	Image           string       `json:"image"`
	Resources       Resources    `json:"resources"`
	Cmd             []string     `json:"cmd"`
	Args            []string     `json:"args"`
	Env             []Env        `json:"env"`
	Config          []ConfigFile `json:"config"`
	Ports           []Port       `json:"ports"`
	LivenessProbe   HealthProbe  `json:"livenessProbe"`
	ReadinessProbe  HealthProbe  `json:"readinessProbe"`
	ImagePullSecret string       `json:"imagePullSecret"`
}

type ParameterType string

const (
	Boolean ParameterType = "boolean"
	String  ParameterType = "string"
	Number  ParameterType = "number"
	Null    ParameterType = "null"
)

type Parameter struct {
	Name          string        `json:"name"`
	Description   string        `json:"description"`
	ParameterType ParameterType `json:"type"`
	Required      bool          `json:"required"`
	Default       string        `json:"default"`
}

// ComponentSpec defines the desired state of ComponentSchematic
type ComponentSpec struct {
	Parameters       []Parameter       `json:"parameters"`
	WorkloadType     string            `json:"workloadType"`
	OsType           string            `json:"osType"`
	Arch             string            `json:"arch"`
	Containers       []Container       `json:"containers"`
	WorkloadSettings []WorkloadSetting `json:"workloadSettings"`
}

type ComponentStatus struct {
}

// +genclient

// ComponentSchematic is the Schema for the components API
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ComponentSchematic struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComponentSpec   `json:"spec,omitempty"`
	Status ComponentStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ComponentSchematicList contains a list of ComponentSchematic
type ComponentSchematicList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComponentSchematic `json:"items"`
}
