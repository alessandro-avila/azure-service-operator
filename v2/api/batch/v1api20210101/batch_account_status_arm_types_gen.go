// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20210101

// Contains information about an Azure Batch account.
type BatchAccount_STATUS_ARM struct {
	// Id: The ID of the resource.
	Id *string `json:"id,omitempty"`

	// Identity: The identity of the Batch account.
	Identity *BatchAccountIdentity_STATUS_ARM `json:"identity,omitempty"`

	// Location: The location of the resource.
	Location *string `json:"location,omitempty"`

	// Name: The name of the resource.
	Name *string `json:"name,omitempty"`

	// Properties: The properties associated with the account.
	Properties *BatchAccountProperties_STATUS_ARM `json:"properties,omitempty"`

	// Tags: The tags of the resource.
	Tags map[string]string `json:"tags,omitempty"`

	// Type: The type of the resource.
	Type *string `json:"type,omitempty"`
}

// The identity of the Batch account, if configured. This is only used when the user specifies 'Microsoft.KeyVault' as
// their Batch account encryption configuration.
type BatchAccountIdentity_STATUS_ARM struct {
	// PrincipalId: The principal id of the Batch account. This property will only be provided for a system assigned identity.
	PrincipalId *string `json:"principalId,omitempty"`

	// TenantId: The tenant id associated with the Batch account. This property will only be provided for a system assigned
	// identity.
	TenantId *string `json:"tenantId,omitempty"`

	// Type: The type of identity used for the Batch account.
	Type *BatchAccountIdentity_Type_STATUS_ARM `json:"type,omitempty"`

	// UserAssignedIdentities: The list of user identities associated with the Batch account. The user identity dictionary key
	// references will be ARM resource ids in the form:
	// '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}'.
	UserAssignedIdentities map[string]BatchAccountIdentity_UserAssignedIdentities_STATUS_ARM `json:"userAssignedIdentities,omitempty"`
}

// Account specific properties.
type BatchAccountProperties_STATUS_ARM struct {
	// AccountEndpoint: The account endpoint used to interact with the Batch service.
	AccountEndpoint              *string `json:"accountEndpoint,omitempty"`
	ActiveJobAndJobScheduleQuota *int    `json:"activeJobAndJobScheduleQuota,omitempty"`

	// AutoStorage: Contains information about the auto-storage account associated with a Batch account.
	AutoStorage *AutoStorageProperties_STATUS_ARM `json:"autoStorage,omitempty"`

	// DedicatedCoreQuota: For accounts with PoolAllocationMode set to UserSubscription, quota is managed on the subscription
	// so this value is not returned.
	DedicatedCoreQuota *int `json:"dedicatedCoreQuota,omitempty"`

	// DedicatedCoreQuotaPerVMFamily: A list of the dedicated core quota per Virtual Machine family for the Batch account. For
	// accounts with PoolAllocationMode set to UserSubscription, quota is managed on the subscription so this value is not
	// returned.
	DedicatedCoreQuotaPerVMFamily []VirtualMachineFamilyCoreQuota_STATUS_ARM `json:"dedicatedCoreQuotaPerVMFamily,omitempty"`

	// DedicatedCoreQuotaPerVMFamilyEnforced: Batch is transitioning its core quota system for dedicated cores to be enforced
	// per Virtual Machine family. During this transitional phase, the dedicated core quota per Virtual Machine family may not
	// yet be enforced. If this flag is false, dedicated core quota is enforced via the old dedicatedCoreQuota property on the
	// account and does not consider Virtual Machine family. If this flag is true, dedicated core quota is enforced via the
	// dedicatedCoreQuotaPerVMFamily property on the account, and the old dedicatedCoreQuota does not apply.
	DedicatedCoreQuotaPerVMFamilyEnforced *bool `json:"dedicatedCoreQuotaPerVMFamilyEnforced,omitempty"`

	// Encryption: Configures how customer data is encrypted inside the Batch account. By default, accounts are encrypted using
	// a Microsoft managed key. For additional control, a customer-managed key can be used instead.
	Encryption *EncryptionProperties_STATUS_ARM `json:"encryption,omitempty"`

	// KeyVaultReference: Identifies the Azure key vault associated with a Batch account.
	KeyVaultReference *KeyVaultReference_STATUS_ARM `json:"keyVaultReference,omitempty"`

	// LowPriorityCoreQuota: For accounts with PoolAllocationMode set to UserSubscription, quota is managed on the subscription
	// so this value is not returned.
	LowPriorityCoreQuota *int `json:"lowPriorityCoreQuota,omitempty"`

	// PoolAllocationMode: The allocation mode for creating pools in the Batch account.
	PoolAllocationMode *PoolAllocationMode_STATUS_ARM `json:"poolAllocationMode,omitempty"`
	PoolQuota          *int                           `json:"poolQuota,omitempty"`

	// PrivateEndpointConnections: List of private endpoint connections associated with the Batch account
	PrivateEndpointConnections []PrivateEndpointConnection_STATUS_ARM `json:"privateEndpointConnections,omitempty"`

	// ProvisioningState: The provisioned state of the resource
	ProvisioningState *BatchAccountProperties_ProvisioningState_STATUS_ARM `json:"provisioningState,omitempty"`

	// PublicNetworkAccess: If not specified, the default value is 'enabled'.
	PublicNetworkAccess *PublicNetworkAccessType_STATUS_ARM `json:"publicNetworkAccess,omitempty"`
}

// Contains information about the auto-storage account associated with a Batch account.
type AutoStorageProperties_STATUS_ARM struct {
	// LastKeySync: The UTC time at which storage keys were last synchronized with the Batch account.
	LastKeySync *string `json:"lastKeySync,omitempty"`

	// StorageAccountId: The resource ID of the storage account to be used for auto-storage account.
	StorageAccountId *string `json:"storageAccountId,omitempty"`
}

type BatchAccountIdentity_Type_STATUS_ARM string

const (
	BatchAccountIdentity_Type_STATUS_ARM_None           = BatchAccountIdentity_Type_STATUS_ARM("None")
	BatchAccountIdentity_Type_STATUS_ARM_SystemAssigned = BatchAccountIdentity_Type_STATUS_ARM("SystemAssigned")
	BatchAccountIdentity_Type_STATUS_ARM_UserAssigned   = BatchAccountIdentity_Type_STATUS_ARM("UserAssigned")
)

// Mapping from string to BatchAccountIdentity_Type_STATUS_ARM
var batchAccountIdentity_Type_STATUS_ARM_Values = map[string]BatchAccountIdentity_Type_STATUS_ARM{
	"none":           BatchAccountIdentity_Type_STATUS_ARM_None,
	"systemassigned": BatchAccountIdentity_Type_STATUS_ARM_SystemAssigned,
	"userassigned":   BatchAccountIdentity_Type_STATUS_ARM_UserAssigned,
}

type BatchAccountIdentity_UserAssignedIdentities_STATUS_ARM struct {
	// ClientId: The client id of user assigned identity.
	ClientId *string `json:"clientId,omitempty"`

	// PrincipalId: The principal id of user assigned identity.
	PrincipalId *string `json:"principalId,omitempty"`
}

type BatchAccountProperties_ProvisioningState_STATUS_ARM string

const (
	BatchAccountProperties_ProvisioningState_STATUS_ARM_Cancelled = BatchAccountProperties_ProvisioningState_STATUS_ARM("Cancelled")
	BatchAccountProperties_ProvisioningState_STATUS_ARM_Creating  = BatchAccountProperties_ProvisioningState_STATUS_ARM("Creating")
	BatchAccountProperties_ProvisioningState_STATUS_ARM_Deleting  = BatchAccountProperties_ProvisioningState_STATUS_ARM("Deleting")
	BatchAccountProperties_ProvisioningState_STATUS_ARM_Failed    = BatchAccountProperties_ProvisioningState_STATUS_ARM("Failed")
	BatchAccountProperties_ProvisioningState_STATUS_ARM_Invalid   = BatchAccountProperties_ProvisioningState_STATUS_ARM("Invalid")
	BatchAccountProperties_ProvisioningState_STATUS_ARM_Succeeded = BatchAccountProperties_ProvisioningState_STATUS_ARM("Succeeded")
)

// Mapping from string to BatchAccountProperties_ProvisioningState_STATUS_ARM
var batchAccountProperties_ProvisioningState_STATUS_ARM_Values = map[string]BatchAccountProperties_ProvisioningState_STATUS_ARM{
	"cancelled": BatchAccountProperties_ProvisioningState_STATUS_ARM_Cancelled,
	"creating":  BatchAccountProperties_ProvisioningState_STATUS_ARM_Creating,
	"deleting":  BatchAccountProperties_ProvisioningState_STATUS_ARM_Deleting,
	"failed":    BatchAccountProperties_ProvisioningState_STATUS_ARM_Failed,
	"invalid":   BatchAccountProperties_ProvisioningState_STATUS_ARM_Invalid,
	"succeeded": BatchAccountProperties_ProvisioningState_STATUS_ARM_Succeeded,
}

// Configures how customer data is encrypted inside the Batch account. By default, accounts are encrypted using a Microsoft
// managed key. For additional control, a customer-managed key can be used instead.
type EncryptionProperties_STATUS_ARM struct {
	// KeySource: Type of the key source.
	KeySource *EncryptionProperties_KeySource_STATUS_ARM `json:"keySource,omitempty"`

	// KeyVaultProperties: Additional details when using Microsoft.KeyVault
	KeyVaultProperties *KeyVaultProperties_STATUS_ARM `json:"keyVaultProperties,omitempty"`
}

// Identifies the Azure key vault associated with a Batch account.
type KeyVaultReference_STATUS_ARM struct {
	// Id: The resource ID of the Azure key vault associated with the Batch account.
	Id *string `json:"id,omitempty"`

	// Url: The URL of the Azure key vault associated with the Batch account.
	Url *string `json:"url,omitempty"`
}

// The allocation mode for creating pools in the Batch account.
type PoolAllocationMode_STATUS_ARM string

const (
	PoolAllocationMode_STATUS_ARM_BatchService     = PoolAllocationMode_STATUS_ARM("BatchService")
	PoolAllocationMode_STATUS_ARM_UserSubscription = PoolAllocationMode_STATUS_ARM("UserSubscription")
)

// Mapping from string to PoolAllocationMode_STATUS_ARM
var poolAllocationMode_STATUS_ARM_Values = map[string]PoolAllocationMode_STATUS_ARM{
	"batchservice":     PoolAllocationMode_STATUS_ARM_BatchService,
	"usersubscription": PoolAllocationMode_STATUS_ARM_UserSubscription,
}

// Contains information about a private link resource.
type PrivateEndpointConnection_STATUS_ARM struct {
	// Id: The ID of the resource.
	Id *string `json:"id,omitempty"`
}

// The network access type for operating on the resources in the Batch account.
type PublicNetworkAccessType_STATUS_ARM string

const (
	PublicNetworkAccessType_STATUS_ARM_Disabled = PublicNetworkAccessType_STATUS_ARM("Disabled")
	PublicNetworkAccessType_STATUS_ARM_Enabled  = PublicNetworkAccessType_STATUS_ARM("Enabled")
)

// Mapping from string to PublicNetworkAccessType_STATUS_ARM
var publicNetworkAccessType_STATUS_ARM_Values = map[string]PublicNetworkAccessType_STATUS_ARM{
	"disabled": PublicNetworkAccessType_STATUS_ARM_Disabled,
	"enabled":  PublicNetworkAccessType_STATUS_ARM_Enabled,
}

// A VM Family and its associated core quota for the Batch account.
type VirtualMachineFamilyCoreQuota_STATUS_ARM struct {
	// CoreQuota: The core quota for the VM family for the Batch account.
	CoreQuota *int `json:"coreQuota,omitempty"`

	// Name: The Virtual Machine family name.
	Name *string `json:"name,omitempty"`
}

type EncryptionProperties_KeySource_STATUS_ARM string

const (
	EncryptionProperties_KeySource_STATUS_ARM_MicrosoftBatch    = EncryptionProperties_KeySource_STATUS_ARM("Microsoft.Batch")
	EncryptionProperties_KeySource_STATUS_ARM_MicrosoftKeyVault = EncryptionProperties_KeySource_STATUS_ARM("Microsoft.KeyVault")
)

// Mapping from string to EncryptionProperties_KeySource_STATUS_ARM
var encryptionProperties_KeySource_STATUS_ARM_Values = map[string]EncryptionProperties_KeySource_STATUS_ARM{
	"microsoft.batch":    EncryptionProperties_KeySource_STATUS_ARM_MicrosoftBatch,
	"microsoft.keyvault": EncryptionProperties_KeySource_STATUS_ARM_MicrosoftKeyVault,
}

// KeyVault configuration when using an encryption KeySource of Microsoft.KeyVault.
type KeyVaultProperties_STATUS_ARM struct {
	// KeyIdentifier: Full path to the versioned secret. Example
	// https://mykeyvault.vault.azure.net/keys/testkey/6e34a81fef704045975661e297a4c053. To be usable the following
	// prerequisites must be met:
	// The Batch Account has a System Assigned identity
	// The account identity has been granted Key/Get, Key/Unwrap and Key/Wrap permissions
	// The KeyVault has soft-delete and purge protection enabled
	KeyIdentifier *string `json:"keyIdentifier,omitempty"`
}
