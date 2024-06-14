// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20230101

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type StorageAccounts_ManagementPolicy_Spec_ARM struct {
	Name string `json:"name,omitempty"`

	// Properties: Returns the Storage Account Data Policies Rules.
	Properties *ManagementPolicyProperties_ARM `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &StorageAccounts_ManagementPolicy_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-01-01"
func (policy StorageAccounts_ManagementPolicy_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (policy *StorageAccounts_ManagementPolicy_Spec_ARM) GetName() string {
	return policy.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Storage/storageAccounts/managementPolicies"
func (policy *StorageAccounts_ManagementPolicy_Spec_ARM) GetType() string {
	return "Microsoft.Storage/storageAccounts/managementPolicies"
}

// The Storage Account ManagementPolicy properties.
type ManagementPolicyProperties_ARM struct {
	// Policy: The Storage Account ManagementPolicy, in JSON format. See more details in:
	// https://docs.microsoft.com/en-us/azure/storage/common/storage-lifecycle-managment-concepts.
	Policy *ManagementPolicySchema_ARM `json:"policy,omitempty"`
}

// The Storage Account ManagementPolicies Rules. See more details in:
// https://docs.microsoft.com/en-us/azure/storage/common/storage-lifecycle-managment-concepts.
type ManagementPolicySchema_ARM struct {
	// Rules: The Storage Account ManagementPolicies Rules. See more details in:
	// https://docs.microsoft.com/en-us/azure/storage/common/storage-lifecycle-managment-concepts.
	Rules []ManagementPolicyRule_ARM `json:"rules"`
}

// An object that wraps the Lifecycle rule. Each rule is uniquely defined by name.
type ManagementPolicyRule_ARM struct {
	// Definition: An object that defines the Lifecycle rule.
	Definition *ManagementPolicyDefinition_ARM `json:"definition,omitempty"`

	// Enabled: Rule is enabled if set to true.
	Enabled *bool `json:"enabled,omitempty"`

	// Name: A rule name can contain any combination of alpha numeric characters. Rule name is case-sensitive. It must be
	// unique within a policy.
	Name *string `json:"name,omitempty"`

	// Type: The valid value is Lifecycle
	Type *ManagementPolicyRule_Type_ARM `json:"type,omitempty"`
}

// An object that defines the Lifecycle rule. Each definition is made up with a filters set and an actions set.
type ManagementPolicyDefinition_ARM struct {
	// Actions: An object that defines the action set.
	Actions *ManagementPolicyAction_ARM `json:"actions,omitempty"`

	// Filters: An object that defines the filter set.
	Filters *ManagementPolicyFilter_ARM `json:"filters,omitempty"`
}

// +kubebuilder:validation:Enum={"Lifecycle"}
type ManagementPolicyRule_Type_ARM string

const ManagementPolicyRule_Type_ARM_Lifecycle = ManagementPolicyRule_Type_ARM("Lifecycle")

// Mapping from string to ManagementPolicyRule_Type_ARM
var managementPolicyRule_Type_ARM_Values = map[string]ManagementPolicyRule_Type_ARM{
	"lifecycle": ManagementPolicyRule_Type_ARM_Lifecycle,
}

// Actions are applied to the filtered blobs when the execution condition is met.
type ManagementPolicyAction_ARM struct {
	// BaseBlob: The management policy action for base blob
	BaseBlob *ManagementPolicyBaseBlob_ARM `json:"baseBlob,omitempty"`

	// Snapshot: The management policy action for snapshot
	Snapshot *ManagementPolicySnapShot_ARM `json:"snapshot,omitempty"`

	// Version: The management policy action for version
	Version *ManagementPolicyVersion_ARM `json:"version,omitempty"`
}

// Filters limit rule actions to a subset of blobs within the storage account. If multiple filters are defined, a logical
// AND is performed on all filters.
type ManagementPolicyFilter_ARM struct {
	// BlobIndexMatch: An array of blob index tag based filters, there can be at most 10 tag filters
	BlobIndexMatch []TagFilter_ARM `json:"blobIndexMatch"`

	// BlobTypes: An array of predefined enum values. Currently blockBlob supports all tiering and delete actions. Only delete
	// actions are supported for appendBlob.
	BlobTypes []string `json:"blobTypes"`

	// PrefixMatch: An array of strings for prefixes to be match.
	PrefixMatch []string `json:"prefixMatch"`
}

// Management policy action for base blob.
type ManagementPolicyBaseBlob_ARM struct {
	// Delete: The function to delete the blob
	Delete *DateAfterModification_ARM `json:"delete,omitempty"`

	// EnableAutoTierToHotFromCool: This property enables auto tiering of a blob from cool to hot on a blob access. This
	// property requires tierToCool.daysAfterLastAccessTimeGreaterThan.
	EnableAutoTierToHotFromCool *bool `json:"enableAutoTierToHotFromCool,omitempty"`

	// TierToArchive: The function to tier blobs to archive storage.
	TierToArchive *DateAfterModification_ARM `json:"tierToArchive,omitempty"`

	// TierToCold: The function to tier blobs to cold storage.
	TierToCold *DateAfterModification_ARM `json:"tierToCold,omitempty"`

	// TierToCool: The function to tier blobs to cool storage.
	TierToCool *DateAfterModification_ARM `json:"tierToCool,omitempty"`

	// TierToHot: The function to tier blobs to hot storage. This action can only be used with Premium Block Blob Storage
	// Accounts
	TierToHot *DateAfterModification_ARM `json:"tierToHot,omitempty"`
}

// Management policy action for snapshot.
type ManagementPolicySnapShot_ARM struct {
	// Delete: The function to delete the blob snapshot
	Delete *DateAfterCreation_ARM `json:"delete,omitempty"`

	// TierToArchive: The function to tier blob snapshot to archive storage.
	TierToArchive *DateAfterCreation_ARM `json:"tierToArchive,omitempty"`

	// TierToCold: The function to tier blobs to cold storage.
	TierToCold *DateAfterCreation_ARM `json:"tierToCold,omitempty"`

	// TierToCool: The function to tier blob snapshot to cool storage.
	TierToCool *DateAfterCreation_ARM `json:"tierToCool,omitempty"`

	// TierToHot: The function to tier blobs to hot storage. This action can only be used with Premium Block Blob Storage
	// Accounts
	TierToHot *DateAfterCreation_ARM `json:"tierToHot,omitempty"`
}

// Management policy action for blob version.
type ManagementPolicyVersion_ARM struct {
	// Delete: The function to delete the blob version
	Delete *DateAfterCreation_ARM `json:"delete,omitempty"`

	// TierToArchive: The function to tier blob version to archive storage.
	TierToArchive *DateAfterCreation_ARM `json:"tierToArchive,omitempty"`

	// TierToCold: The function to tier blobs to cold storage.
	TierToCold *DateAfterCreation_ARM `json:"tierToCold,omitempty"`

	// TierToCool: The function to tier blob version to cool storage.
	TierToCool *DateAfterCreation_ARM `json:"tierToCool,omitempty"`

	// TierToHot: The function to tier blobs to hot storage. This action can only be used with Premium Block Blob Storage
	// Accounts
	TierToHot *DateAfterCreation_ARM `json:"tierToHot,omitempty"`
}

// Blob index tag based filtering for blob objects
type TagFilter_ARM struct {
	// Name: This is the filter tag name, it can have 1 - 128 characters
	Name *string `json:"name,omitempty"`

	// Op: This is the comparison operator which is used for object comparison and filtering. Only == (equality operator) is
	// currently supported
	Op *string `json:"op,omitempty"`

	// Value: This is the filter tag value field used for tag based filtering, it can have 0 - 256 characters
	Value *string `json:"value,omitempty"`
}

// Object to define snapshot and version action conditions.
type DateAfterCreation_ARM struct {
	// DaysAfterCreationGreaterThan: Value indicating the age in days after creation
	DaysAfterCreationGreaterThan *int `json:"daysAfterCreationGreaterThan,omitempty"`

	// DaysAfterLastTierChangeGreaterThan: Value indicating the age in days after last blob tier change time. This property is
	// only applicable for tierToArchive actions and requires daysAfterCreationGreaterThan to be set for snapshots and blob
	// version based actions. The blob will be archived if both the conditions are satisfied.
	DaysAfterLastTierChangeGreaterThan *int `json:"daysAfterLastTierChangeGreaterThan,omitempty"`
}

// Object to define the base blob action conditions. Properties daysAfterModificationGreaterThan,
// daysAfterLastAccessTimeGreaterThan and daysAfterCreationGreaterThan are mutually exclusive. The
// daysAfterLastTierChangeGreaterThan property is only applicable for tierToArchive actions which requires
// daysAfterModificationGreaterThan to be set, also it cannot be used in conjunction with
// daysAfterLastAccessTimeGreaterThan or daysAfterCreationGreaterThan.
type DateAfterModification_ARM struct {
	// DaysAfterCreationGreaterThan: Value indicating the age in days after blob creation.
	DaysAfterCreationGreaterThan *int `json:"daysAfterCreationGreaterThan,omitempty"`

	// DaysAfterLastAccessTimeGreaterThan: Value indicating the age in days after last blob access. This property can only be
	// used in conjunction with last access time tracking policy
	DaysAfterLastAccessTimeGreaterThan *int `json:"daysAfterLastAccessTimeGreaterThan,omitempty"`

	// DaysAfterLastTierChangeGreaterThan: Value indicating the age in days after last blob tier change time. This property is
	// only applicable for tierToArchive actions and requires daysAfterModificationGreaterThan to be set for baseBlobs based
	// actions. The blob will be archived if both the conditions are satisfied.
	DaysAfterLastTierChangeGreaterThan *int `json:"daysAfterLastTierChangeGreaterThan,omitempty"`

	// DaysAfterModificationGreaterThan: Value indicating the age in days after last modification
	DaysAfterModificationGreaterThan *int `json:"daysAfterModificationGreaterThan,omitempty"`
}
