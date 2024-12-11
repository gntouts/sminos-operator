package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// InstanceSpec defines the desired state of Instance
type InstanceSpec struct {
	// BrokerProperties defines environment variables for broker Pods
	BrokerProperties map[string]string `json:"brokerProperties,omitempty"`

	// Capacity is the number of slots for the Instance
	Capacity uint `json:"capacity"`

	// CdiName is the CDI fully qualified name of the device linked to the Instance
	CdiName string `json:"cdiName"`

	// ConfigurationName is the name of the corresponding Configuration
	ConfigurationName string `json:"configurationName"`

	// DeviceUsage is a map of capability slots to node names
	DeviceUsage map[string]string `json:"deviceUsage,omitempty"`

	// Nodes is a list of nodes that can access this capability instance
	Nodes []string `json:"nodes,omitempty"`

	// Shared defines whether the capability is shared by multiple nodes
	Shared bool `json:"shared,omitempty"`
}

// InstanceStatus defines the observed state of Sminos
type InstanceStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// Instance is the Schema for the Instance API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type Instance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   InstanceSpec   `json:"spec"`
	Status InstanceStatus `json:"status,omitempty"`
}

// InstanceList contains a list of Instance
// +kubebuilder:object:root=true
type InstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []Instance `json:"items"`
}

// DiscoveryProperty represents an individual property in discoveryProperties
type DiscoveryProperty struct {
	Name      string  `json:"name"`
	Value     *string `json:"value,omitempty"`
	ValueFrom *struct {
		SecretKeyRef *struct {
			Key       string `json:"key"`
			Name      string `json:"name"`
			Namespace string `json:"namespace"`
			Optional  *bool  `json:"optional,omitempty"`
		} `json:"secretKeyRef,omitempty"`
		ConfigMapKeyRef *struct {
			Key       string `json:"key"`
			Name      string `json:"name"`
			Namespace string `json:"namespace"`
			Optional  *bool  `json:"optional,omitempty"`
		} `json:"configMapKeyRef,omitempty"`
	} `json:"valueFrom,omitempty"`
}

// DiscoveryHandlerInfo contains details about a discovery handler
type DiscoveryHandlerInfo struct {
	Name                string              `json:"name"`
	DiscoveryDetails    string              `json:"discoveryDetails"`
	DiscoveryProperties []DiscoveryProperty `json:"discoveryProperties,omitempty"`
}

// ConfigurationSpec defines the desired state of Configuration
type ConfigurationSpec struct {
	DiscoveryHandler DiscoveryHandlerInfo `json:"discoveryHandler"`
	Capacity         *int                 `json:"capacity,omitempty"`
	BrokerSpec       *struct {
		BrokerJobSpec *map[string]interface{} `json:"brokerJobSpec,omitempty"`
		BrokerPodSpec *map[string]interface{} `json:"brokerPodSpec,omitempty"`
	} `json:"brokerSpec,omitempty"`
	InstanceServiceSpec      *map[string]interface{} `json:"instanceServiceSpec,omitempty"`
	ConfigurationServiceSpec *map[string]interface{} `json:"configurationServiceSpec,omitempty"`
	BrokerProperties         map[string]string       `json:"brokerProperties,omitempty"`
}

// Configuration is the Schema for the Configuration API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type Configuration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ConfigurationSpec `json:"spec"`
}

// ConfigurationList contains a list of Configuration
// +kubebuilder:object:root=true
type ConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []Configuration `json:"items"`
}~

func init() {
	SchemeBuilder.Register(&Instance{}, &InstanceList{}, &Configuration{}, &ConfigurationList{})
}
