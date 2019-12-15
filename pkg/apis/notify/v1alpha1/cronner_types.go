package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// CronnerSpec defines the desired state of Cronner
type CronnerSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	Cluster       string   `json:"cluster"`
	ContactMethod string   `json:"contactMethod"`
	CronjobName   string   `json:"cronjobName"`
	CurrentPodID  string   `json:"currentPodID"`
	Name          string   `json:"name"`
	Namespace     string   `json:"namespace"`
	SMTPDl        []string `json:"smtpDl"`
	SMTPHost      string   `json:"smtpHost"`
	SMTPPort      int64    `json:"smtpPort"`
	SMTPTls       bool     `json:"smtpTls"`
}

// CronnerStatus defines the observed state of Cronner
type CronnerStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	Cluster       string   `json:"cluster"`
	ContactMethod string   `json:"contactMethod"`
	CronjobName   string   `json:"cronjobName"`
	CurrentPodID  string   `json:"currentPodID"`
	Name          string   `json:"name"`
	Namespace     string   `json:"namespace"`
	SMTPDl        []string `json:"smtpDl"`
	SMTPHost      string   `json:"smtpHost"`
	SMTPPort      int64    `json:"smtpPort"`
	SMTPTls       bool     `json:"smtpTls"`

	Items []Cronner `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Cronner is the Schema for the cronners API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=cronners,scope=Namespaced
type Cronner struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CronnerSpec   `json:"spec,omitempty"`
	Status CronnerStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CronnerList contains a list of Cronner
type CronnerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Cronner `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Cronner{}, &CronnerList{})
}
