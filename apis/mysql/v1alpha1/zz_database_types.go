// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type DatabaseInitParameters struct {

	// The default character set to use when
	// a table is created without specifying an explicit character set. Defaults
	// to utf8mb4.
	DefaultCharacterSet *string `json:"defaultCharacterSet,omitempty" tf:"default_character_set,omitempty"`

	// The default collation to use when a table
	// is created without specifying an explicit collation. Defaults to
	// utf8mb4_general_ci. Each character set has its own set of collations, so
	// changing the character set requires also changing the collation.
	DefaultCollation *string `json:"defaultCollation,omitempty" tf:"default_collation,omitempty"`
}

type DatabaseObservation struct {

	// The default character set to use when
	// a table is created without specifying an explicit character set. Defaults
	// to utf8mb4.
	DefaultCharacterSet *string `json:"defaultCharacterSet,omitempty" tf:"default_character_set,omitempty"`

	// The default collation to use when a table
	// is created without specifying an explicit collation. Defaults to
	// utf8mb4_general_ci. Each character set has its own set of collations, so
	// changing the character set requires also changing the collation.
	DefaultCollation *string `json:"defaultCollation,omitempty" tf:"default_collation,omitempty"`

	// The id of the database.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type DatabaseParameters struct {

	// The default character set to use when
	// a table is created without specifying an explicit character set. Defaults
	// to utf8mb4.
	// +kubebuilder:validation:Optional
	DefaultCharacterSet *string `json:"defaultCharacterSet,omitempty" tf:"default_character_set,omitempty"`

	// The default collation to use when a table
	// is created without specifying an explicit collation. Defaults to
	// utf8mb4_general_ci. Each character set has its own set of collations, so
	// changing the character set requires also changing the collation.
	// +kubebuilder:validation:Optional
	DefaultCollation *string `json:"defaultCollation,omitempty" tf:"default_collation,omitempty"`
}

// DatabaseSpec defines the desired state of Database
type DatabaseSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DatabaseParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider DatabaseInitParameters `json:"initProvider,omitempty"`
}

// DatabaseStatus defines the observed state of Database.
type DatabaseStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DatabaseObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Database is the Schema for the Databases API. Creates and manages a database on a MySQL server.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,mysql}
type Database struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatabaseSpec   `json:"spec"`
	Status            DatabaseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DatabaseList contains a list of Databases
type DatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Database `json:"items"`
}

// Repository type metadata.
var (
	Database_Kind             = "Database"
	Database_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Database_Kind}.String()
	Database_KindAPIVersion   = Database_Kind + "." + CRDGroupVersion.String()
	Database_GroupVersionKind = CRDGroupVersion.WithKind(Database_Kind)
)

func init() {
	SchemeBuilder.Register(&Database{}, &DatabaseList{})
}