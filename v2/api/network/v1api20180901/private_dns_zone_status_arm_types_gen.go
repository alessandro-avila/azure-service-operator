// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20180901

type PrivateDnsZone_STATUS_ARM struct {
	// Etag: The ETag of the zone.
	Etag *string `json:"etag,omitempty"`

	// Id: Fully qualified resource Id for the resource. Example -
	// '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateDnsZoneName}'.
	Id *string `json:"id,omitempty"`

	// Location: The Azure Region where the resource lives
	Location *string `json:"location,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// Properties: Properties of the Private DNS zone.
	Properties *PrivateZoneProperties_STATUS_ARM `json:"properties,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	// Type: The type of the resource. Example - 'Microsoft.Network/privateDnsZones'.
	Type *string `json:"type,omitempty"`
}

// Represents the properties of the Private DNS zone.
type PrivateZoneProperties_STATUS_ARM struct {
	// MaxNumberOfRecordSets: The maximum number of record sets that can be created in this Private DNS zone. This is a
	// read-only property and any attempt to set this value will be ignored.
	MaxNumberOfRecordSets *int `json:"maxNumberOfRecordSets,omitempty"`

	// MaxNumberOfVirtualNetworkLinks: The maximum number of virtual networks that can be linked to this Private DNS zone. This
	// is a read-only property and any attempt to set this value will be ignored.
	MaxNumberOfVirtualNetworkLinks *int `json:"maxNumberOfVirtualNetworkLinks,omitempty"`

	// MaxNumberOfVirtualNetworkLinksWithRegistration: The maximum number of virtual networks that can be linked to this
	// Private DNS zone with registration enabled. This is a read-only property and any attempt to set this value will be
	// ignored.
	MaxNumberOfVirtualNetworkLinksWithRegistration *int `json:"maxNumberOfVirtualNetworkLinksWithRegistration,omitempty"`

	// NumberOfRecordSets: The current number of record sets in this Private DNS zone. This is a read-only property and any
	// attempt to set this value will be ignored.
	NumberOfRecordSets *int `json:"numberOfRecordSets,omitempty"`

	// NumberOfVirtualNetworkLinks: The current number of virtual networks that are linked to this Private DNS zone. This is a
	// read-only property and any attempt to set this value will be ignored.
	NumberOfVirtualNetworkLinks *int `json:"numberOfVirtualNetworkLinks,omitempty"`

	// NumberOfVirtualNetworkLinksWithRegistration: The current number of virtual networks that are linked to this Private DNS
	// zone with registration enabled. This is a read-only property and any attempt to set this value will be ignored.
	NumberOfVirtualNetworkLinksWithRegistration *int `json:"numberOfVirtualNetworkLinksWithRegistration,omitempty"`

	// ProvisioningState: The provisioning state of the resource. This is a read-only property and any attempt to set this
	// value will be ignored.
	ProvisioningState *PrivateZoneProperties_ProvisioningState_STATUS_ARM `json:"provisioningState,omitempty"`
}

type PrivateZoneProperties_ProvisioningState_STATUS_ARM string

const (
	PrivateZoneProperties_ProvisioningState_STATUS_ARM_Canceled  = PrivateZoneProperties_ProvisioningState_STATUS_ARM("Canceled")
	PrivateZoneProperties_ProvisioningState_STATUS_ARM_Creating  = PrivateZoneProperties_ProvisioningState_STATUS_ARM("Creating")
	PrivateZoneProperties_ProvisioningState_STATUS_ARM_Deleting  = PrivateZoneProperties_ProvisioningState_STATUS_ARM("Deleting")
	PrivateZoneProperties_ProvisioningState_STATUS_ARM_Failed    = PrivateZoneProperties_ProvisioningState_STATUS_ARM("Failed")
	PrivateZoneProperties_ProvisioningState_STATUS_ARM_Succeeded = PrivateZoneProperties_ProvisioningState_STATUS_ARM("Succeeded")
	PrivateZoneProperties_ProvisioningState_STATUS_ARM_Updating  = PrivateZoneProperties_ProvisioningState_STATUS_ARM("Updating")
)

// Mapping from string to PrivateZoneProperties_ProvisioningState_STATUS_ARM
var privateZoneProperties_ProvisioningState_STATUS_ARM_Values = map[string]PrivateZoneProperties_ProvisioningState_STATUS_ARM{
	"canceled":  PrivateZoneProperties_ProvisioningState_STATUS_ARM_Canceled,
	"creating":  PrivateZoneProperties_ProvisioningState_STATUS_ARM_Creating,
	"deleting":  PrivateZoneProperties_ProvisioningState_STATUS_ARM_Deleting,
	"failed":    PrivateZoneProperties_ProvisioningState_STATUS_ARM_Failed,
	"succeeded": PrivateZoneProperties_ProvisioningState_STATUS_ARM_Succeeded,
	"updating":  PrivateZoneProperties_ProvisioningState_STATUS_ARM_Updating,
}
