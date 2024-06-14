// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20201201

type Redis_LinkedServer_STATUS_ARM struct {
	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// Properties: Properties of the linked server.
	Properties *RedisLinkedServerProperties_STATUS_ARM `json:"properties,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

// Properties of a linked server to be returned in get/put response
type RedisLinkedServerProperties_STATUS_ARM struct {
	// LinkedRedisCacheId: Fully qualified resourceId of the linked redis cache.
	LinkedRedisCacheId *string `json:"linkedRedisCacheId,omitempty"`

	// LinkedRedisCacheLocation: Location of the linked redis cache.
	LinkedRedisCacheLocation *string `json:"linkedRedisCacheLocation,omitempty"`

	// ProvisioningState: Terminal state of the link between primary and secondary redis cache.
	ProvisioningState *string `json:"provisioningState,omitempty"`

	// ServerRole: Role of the linked server.
	ServerRole *RedisLinkedServerProperties_ServerRole_STATUS_ARM `json:"serverRole,omitempty"`
}

type RedisLinkedServerProperties_ServerRole_STATUS_ARM string

const (
	RedisLinkedServerProperties_ServerRole_STATUS_ARM_Primary   = RedisLinkedServerProperties_ServerRole_STATUS_ARM("Primary")
	RedisLinkedServerProperties_ServerRole_STATUS_ARM_Secondary = RedisLinkedServerProperties_ServerRole_STATUS_ARM("Secondary")
)

// Mapping from string to RedisLinkedServerProperties_ServerRole_STATUS_ARM
var redisLinkedServerProperties_ServerRole_STATUS_ARM_Values = map[string]RedisLinkedServerProperties_ServerRole_STATUS_ARM{
	"primary":   RedisLinkedServerProperties_ServerRole_STATUS_ARM_Primary,
	"secondary": RedisLinkedServerProperties_ServerRole_STATUS_ARM_Secondary,
}
