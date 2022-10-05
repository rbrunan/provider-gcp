/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ProjectIAMAuditConfigAuditLogConfigObservation struct {
}

type ProjectIAMAuditConfigAuditLogConfigParameters struct {

	// +kubebuilder:validation:Optional
	ExemptedMembers []*string `json:"exemptedMembers,omitempty" tf:"exempted_members,omitempty"`

	// +kubebuilder:validation:Required
	LogType *string `json:"logType" tf:"log_type,omitempty"`
}

type ProjectIAMAuditConfigObservation struct {
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ProjectIAMAuditConfigParameters struct {

	// +kubebuilder:validation:Required
	AuditLogConfig []ProjectIAMAuditConfigAuditLogConfigParameters `json:"auditLogConfig" tf:"audit_log_config,omitempty"`

	// +crossplane:generate:reference:type=Project
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Reference to a Project to populate project.
	// +kubebuilder:validation:Optional
	ProjectRef *v1.Reference `json:"projectRef,omitempty" tf:"-"`

	// Selector for a Project to populate project.
	// +kubebuilder:validation:Optional
	ProjectSelector *v1.Selector `json:"projectSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	Service *string `json:"service" tf:"service,omitempty"`
}

// ProjectIAMAuditConfigSpec defines the desired state of ProjectIAMAuditConfig
type ProjectIAMAuditConfigSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProjectIAMAuditConfigParameters `json:"forProvider"`
}

// ProjectIAMAuditConfigStatus defines the observed state of ProjectIAMAuditConfig.
type ProjectIAMAuditConfigStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProjectIAMAuditConfigObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ProjectIAMAuditConfig is the Schema for the ProjectIAMAuditConfigs API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type ProjectIAMAuditConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ProjectIAMAuditConfigSpec   `json:"spec"`
	Status            ProjectIAMAuditConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProjectIAMAuditConfigList contains a list of ProjectIAMAuditConfigs
type ProjectIAMAuditConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProjectIAMAuditConfig `json:"items"`
}

// Repository type metadata.
var (
	ProjectIAMAuditConfig_Kind             = "ProjectIAMAuditConfig"
	ProjectIAMAuditConfig_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ProjectIAMAuditConfig_Kind}.String()
	ProjectIAMAuditConfig_KindAPIVersion   = ProjectIAMAuditConfig_Kind + "." + CRDGroupVersion.String()
	ProjectIAMAuditConfig_GroupVersionKind = CRDGroupVersion.WithKind(ProjectIAMAuditConfig_Kind)
)

func init() {
	SchemeBuilder.Register(&ProjectIAMAuditConfig{}, &ProjectIAMAuditConfigList{})
}