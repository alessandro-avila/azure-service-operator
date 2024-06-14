// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220701

// Private link service resource.
type PrivateLinkService_STATUS_PrivateLinkService_SubResourceEmbedded_ARM struct {
	// Etag: A unique read-only string that changes whenever the resource is updated.
	Etag *string `json:"etag,omitempty"`

	// ExtendedLocation: The extended location of the load balancer.
	ExtendedLocation *ExtendedLocation_STATUS_ARM `json:"extendedLocation,omitempty"`

	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Location: Resource location.
	Location *string `json:"location,omitempty"`

	// Name: Resource name.
	Name *string `json:"name,omitempty"`

	// Properties: Properties of the private link service.
	Properties *PrivateLinkServiceProperties_STATUS_ARM `json:"properties,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	// Type: Resource type.
	Type *string `json:"type,omitempty"`
}

// Properties of the private link service.
type PrivateLinkServiceProperties_STATUS_ARM struct {
	// Alias: The alias of the private link service.
	Alias *string `json:"alias,omitempty"`

	// AutoApproval: The auto-approval list of the private link service.
	AutoApproval *ResourceSet_STATUS_ARM `json:"autoApproval,omitempty"`

	// EnableProxyProtocol: Whether the private link service is enabled for proxy protocol or not.
	EnableProxyProtocol *bool `json:"enableProxyProtocol,omitempty"`

	// Fqdns: The list of Fqdn.
	Fqdns []string `json:"fqdns,omitempty"`

	// IpConfigurations: An array of private link service IP configurations.
	IpConfigurations []PrivateLinkServiceIpConfiguration_STATUS_ARM `json:"ipConfigurations,omitempty"`

	// LoadBalancerFrontendIpConfigurations: An array of references to the load balancer IP configurations.
	LoadBalancerFrontendIpConfigurations []FrontendIPConfiguration_STATUS_PrivateLinkService_SubResourceEmbedded_ARM `json:"loadBalancerFrontendIpConfigurations,omitempty"`

	// NetworkInterfaces: An array of references to the network interfaces created for this private link service.
	NetworkInterfaces []NetworkInterface_STATUS_PrivateLinkService_SubResourceEmbedded_ARM `json:"networkInterfaces,omitempty"`

	// PrivateEndpointConnections: An array of list about connections to the private endpoint.
	PrivateEndpointConnections []PrivateEndpointConnection_STATUS_ARM `json:"privateEndpointConnections,omitempty"`

	// ProvisioningState: The provisioning state of the private link service resource.
	ProvisioningState *ApplicationGatewayProvisioningState_STATUS_ARM `json:"provisioningState,omitempty"`

	// Visibility: The visibility list of the private link service.
	Visibility *ResourceSet_STATUS_ARM `json:"visibility,omitempty"`
}

// Frontend IP address of the load balancer.
type FrontendIPConfiguration_STATUS_PrivateLinkService_SubResourceEmbedded_ARM struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`
}

// A network interface in a resource group.
type NetworkInterface_STATUS_PrivateLinkService_SubResourceEmbedded_ARM struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`
}

// PrivateEndpointConnection resource.
type PrivateEndpointConnection_STATUS_ARM struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`
}

// The private link service ip configuration.
type PrivateLinkServiceIpConfiguration_STATUS_ARM struct {
	// Etag: A unique read-only string that changes whenever the resource is updated.
	Etag *string `json:"etag,omitempty"`

	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Name: The name of private link service ip configuration.
	Name *string `json:"name,omitempty"`

	// Properties: Properties of the private link service ip configuration.
	Properties *PrivateLinkServiceIpConfigurationProperties_STATUS_ARM `json:"properties,omitempty"`

	// Type: The resource type.
	Type *string `json:"type,omitempty"`
}

// The base resource set for visibility and auto-approval.
type ResourceSet_STATUS_ARM struct {
	// Subscriptions: The list of subscriptions.
	Subscriptions []string `json:"subscriptions,omitempty"`
}

// Properties of private link service IP configuration.
type PrivateLinkServiceIpConfigurationProperties_STATUS_ARM struct {
	// Primary: Whether the ip configuration is primary or not.
	Primary *bool `json:"primary,omitempty"`

	// PrivateIPAddress: The private IP address of the IP configuration.
	PrivateIPAddress *string `json:"privateIPAddress,omitempty"`

	// PrivateIPAddressVersion: Whether the specific IP configuration is IPv4 or IPv6. Default is IPv4.
	PrivateIPAddressVersion *IPVersion_STATUS_ARM `json:"privateIPAddressVersion,omitempty"`

	// PrivateIPAllocationMethod: The private IP address allocation method.
	PrivateIPAllocationMethod *IPAllocationMethod_STATUS_ARM `json:"privateIPAllocationMethod,omitempty"`

	// ProvisioningState: The provisioning state of the private link service IP configuration resource.
	ProvisioningState *ApplicationGatewayProvisioningState_STATUS_ARM `json:"provisioningState,omitempty"`

	// Subnet: The reference to the subnet resource.
	Subnet *Subnet_STATUS_PrivateLinkService_SubResourceEmbedded_ARM `json:"subnet,omitempty"`
}

// IP address allocation method.
type IPAllocationMethod_STATUS_ARM string

const (
	IPAllocationMethod_STATUS_ARM_Dynamic = IPAllocationMethod_STATUS_ARM("Dynamic")
	IPAllocationMethod_STATUS_ARM_Static  = IPAllocationMethod_STATUS_ARM("Static")
)

// Mapping from string to IPAllocationMethod_STATUS_ARM
var iPAllocationMethod_STATUS_ARM_Values = map[string]IPAllocationMethod_STATUS_ARM{
	"dynamic": IPAllocationMethod_STATUS_ARM_Dynamic,
	"static":  IPAllocationMethod_STATUS_ARM_Static,
}

// Subnet in a virtual network resource.
type Subnet_STATUS_PrivateLinkService_SubResourceEmbedded_ARM struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`
}
