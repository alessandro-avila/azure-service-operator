// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220701

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type DnsResolvers_InboundEndpoint_Spec_ARM struct {
	// Location: The geo-location where the resource lives
	Location *string `json:"location,omitempty"`
	Name     string  `json:"name,omitempty"`

	// Properties: Properties of the inbound endpoint.
	Properties *InboundEndpointProperties_ARM `json:"properties,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &DnsResolvers_InboundEndpoint_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2022-07-01"
func (endpoint DnsResolvers_InboundEndpoint_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (endpoint *DnsResolvers_InboundEndpoint_Spec_ARM) GetName() string {
	return endpoint.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/dnsResolvers/inboundEndpoints"
func (endpoint *DnsResolvers_InboundEndpoint_Spec_ARM) GetType() string {
	return "Microsoft.Network/dnsResolvers/inboundEndpoints"
}

// Represents the properties of an inbound endpoint for a DNS resolver.
type InboundEndpointProperties_ARM struct {
	// IpConfigurations: IP configurations for the inbound endpoint.
	IpConfigurations []IpConfiguration_ARM `json:"ipConfigurations,omitempty"`
}

// IP configuration.
type IpConfiguration_ARM struct {
	// PrivateIpAddress: Private IP address of the IP configuration.
	PrivateIpAddress *string `json:"privateIpAddress,omitempty"`

	// PrivateIpAllocationMethod: Private IP address allocation method.
	PrivateIpAllocationMethod *IpConfiguration_PrivateIpAllocationMethod_ARM `json:"privateIpAllocationMethod,omitempty"`

	// Subnet: The reference to the subnet bound to the IP configuration.
	Subnet *DnsresolverSubResource_ARM `json:"subnet,omitempty"`
}

// +kubebuilder:validation:Enum={"Dynamic","Static"}
type IpConfiguration_PrivateIpAllocationMethod_ARM string

const (
	IpConfiguration_PrivateIpAllocationMethod_ARM_Dynamic = IpConfiguration_PrivateIpAllocationMethod_ARM("Dynamic")
	IpConfiguration_PrivateIpAllocationMethod_ARM_Static  = IpConfiguration_PrivateIpAllocationMethod_ARM("Static")
)

// Mapping from string to IpConfiguration_PrivateIpAllocationMethod_ARM
var ipConfiguration_PrivateIpAllocationMethod_ARM_Values = map[string]IpConfiguration_PrivateIpAllocationMethod_ARM{
	"dynamic": IpConfiguration_PrivateIpAllocationMethod_ARM_Dynamic,
	"static":  IpConfiguration_PrivateIpAllocationMethod_ARM_Static,
}
