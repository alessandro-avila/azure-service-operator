// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20230501preview

type Service_Backend_STATUS_ARM struct {
	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// Properties: Backend entity contract properties.
	Properties *BackendContractProperties_STATUS_ARM `json:"properties,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

// Parameters supplied to the Create Backend operation.
type BackendContractProperties_STATUS_ARM struct {
	// CircuitBreaker: Backend Circuit Breaker Configuration
	CircuitBreaker *BackendCircuitBreaker_STATUS_ARM `json:"circuitBreaker,omitempty"`

	// Credentials: Backend Credentials Contract Properties
	Credentials *BackendCredentialsContract_STATUS_ARM `json:"credentials,omitempty"`

	// Description: Backend Description.
	Description *string `json:"description,omitempty"`

	// Pool: Backend pool information
	Pool *BackendPool_STATUS_ARM `json:"pool,omitempty"`

	// Properties: Backend Properties contract
	Properties *BackendProperties_STATUS_ARM `json:"properties,omitempty"`

	// Protocol: Backend communication protocol.
	Protocol *BackendContractProperties_Protocol_STATUS_ARM `json:"protocol,omitempty"`

	// Proxy: Backend gateway Contract Properties
	Proxy *BackendProxyContract_STATUS_ARM `json:"proxy,omitempty"`

	// ResourceId: Management Uri of the Resource in External System. This URL can be the Arm Resource Id of Logic Apps,
	// Function Apps or API Apps.
	ResourceId *string `json:"resourceId,omitempty"`

	// Title: Backend Title.
	Title *string `json:"title,omitempty"`

	// Tls: Backend TLS Properties
	Tls *BackendTlsProperties_STATUS_ARM `json:"tls,omitempty"`

	// Type: Type of the backend. A backend can be either Single or Pool.
	Type *BackendContractProperties_Type_STATUS_ARM `json:"type,omitempty"`

	// Url: Runtime Url of the Backend.
	Url *string `json:"url,omitempty"`
}

// The configuration of the backend circuit breaker
type BackendCircuitBreaker_STATUS_ARM struct {
	// Rules: The rules for tripping the backend.
	Rules []CircuitBreakerRule_STATUS_ARM `json:"rules,omitempty"`
}

type BackendContractProperties_Protocol_STATUS_ARM string

const (
	BackendContractProperties_Protocol_STATUS_ARM_Http = BackendContractProperties_Protocol_STATUS_ARM("http")
	BackendContractProperties_Protocol_STATUS_ARM_Soap = BackendContractProperties_Protocol_STATUS_ARM("soap")
)

// Mapping from string to BackendContractProperties_Protocol_STATUS_ARM
var backendContractProperties_Protocol_STATUS_ARM_Values = map[string]BackendContractProperties_Protocol_STATUS_ARM{
	"http": BackendContractProperties_Protocol_STATUS_ARM_Http,
	"soap": BackendContractProperties_Protocol_STATUS_ARM_Soap,
}

type BackendContractProperties_Type_STATUS_ARM string

const (
	BackendContractProperties_Type_STATUS_ARM_Pool   = BackendContractProperties_Type_STATUS_ARM("Pool")
	BackendContractProperties_Type_STATUS_ARM_Single = BackendContractProperties_Type_STATUS_ARM("Single")
)

// Mapping from string to BackendContractProperties_Type_STATUS_ARM
var backendContractProperties_Type_STATUS_ARM_Values = map[string]BackendContractProperties_Type_STATUS_ARM{
	"pool":   BackendContractProperties_Type_STATUS_ARM_Pool,
	"single": BackendContractProperties_Type_STATUS_ARM_Single,
}

// Details of the Credentials used to connect to Backend.
type BackendCredentialsContract_STATUS_ARM struct {
	// Authorization: Authorization header authentication
	Authorization *BackendAuthorizationHeaderCredentials_STATUS_ARM `json:"authorization,omitempty"`

	// Certificate: List of Client Certificate Thumbprints. Will be ignored if certificatesIds are provided.
	Certificate []string `json:"certificate,omitempty"`

	// CertificateIds: List of Client Certificate Ids.
	CertificateIds []string `json:"certificateIds,omitempty"`

	// Header: Header Parameter description.
	Header map[string][]string `json:"header,omitempty"`

	// Query: Query Parameter description.
	Query map[string][]string `json:"query,omitempty"`
}

// Backend pool information
type BackendPool_STATUS_ARM struct {
	// Services: The list of backend entities belonging to a pool.
	Services []BackendPoolItem_STATUS_ARM `json:"services,omitempty"`
}

// Properties specific to the Backend Type.
type BackendProperties_STATUS_ARM struct {
	// ServiceFabricCluster: Backend Service Fabric Cluster Properties
	ServiceFabricCluster *BackendServiceFabricClusterProperties_STATUS_ARM `json:"serviceFabricCluster,omitempty"`
}

// Details of the Backend WebProxy Server to use in the Request to Backend.
type BackendProxyContract_STATUS_ARM struct {
	// Url: WebProxy Server AbsoluteUri property which includes the entire URI stored in the Uri instance, including all
	// fragments and query strings.
	Url *string `json:"url,omitempty"`

	// Username: Username to connect to the WebProxy server
	Username *string `json:"username,omitempty"`
}

// Properties controlling TLS Certificate Validation.
type BackendTlsProperties_STATUS_ARM struct {
	// ValidateCertificateChain: Flag indicating whether SSL certificate chain validation should be done when using self-signed
	// certificates for this backend host.
	ValidateCertificateChain *bool `json:"validateCertificateChain,omitempty"`

	// ValidateCertificateName: Flag indicating whether SSL certificate name validation should be done when using self-signed
	// certificates for this backend host.
	ValidateCertificateName *bool `json:"validateCertificateName,omitempty"`
}

// Authorization header information.
type BackendAuthorizationHeaderCredentials_STATUS_ARM struct {
	// Parameter: Authentication Parameter value.
	Parameter *string `json:"parameter,omitempty"`

	// Scheme: Authentication Scheme name.
	Scheme *string `json:"scheme,omitempty"`
}

// Backend pool service information
type BackendPoolItem_STATUS_ARM struct {
	// Id: The unique ARM id of the backend entity. The ARM id should refer to an already existing backend entity.
	Id *string `json:"id,omitempty"`
}

// Properties of the Service Fabric Type Backend.
type BackendServiceFabricClusterProperties_STATUS_ARM struct {
	// ClientCertificateId: The client certificate id for the management endpoint.
	ClientCertificateId *string `json:"clientCertificateId,omitempty"`

	// ClientCertificatethumbprint: The client certificate thumbprint for the management endpoint. Will be ignored if
	// certificatesIds are provided
	ClientCertificatethumbprint *string `json:"clientCertificatethumbprint,omitempty"`

	// ManagementEndpoints: The cluster management endpoint.
	ManagementEndpoints []string `json:"managementEndpoints,omitempty"`

	// MaxPartitionResolutionRetries: Maximum number of retries while attempting resolve the partition.
	MaxPartitionResolutionRetries *int `json:"maxPartitionResolutionRetries,omitempty"`

	// ServerCertificateThumbprints: Thumbprints of certificates cluster management service uses for tls communication
	ServerCertificateThumbprints []string `json:"serverCertificateThumbprints,omitempty"`

	// ServerX509Names: Server X509 Certificate Names Collection
	ServerX509Names []X509CertificateName_STATUS_ARM `json:"serverX509Names,omitempty"`
}

// Rule configuration to trip the backend.
type CircuitBreakerRule_STATUS_ARM struct {
	// FailureCondition: The conditions for tripping the circuit breaker.
	FailureCondition *CircuitBreakerFailureCondition_STATUS_ARM `json:"failureCondition,omitempty"`

	// Name: The rule name.
	Name *string `json:"name,omitempty"`

	// TripDuration: The duration for which the circuit will be tripped.
	TripDuration *string `json:"tripDuration,omitempty"`
}

// The trip conditions of the circuit breaker
type CircuitBreakerFailureCondition_STATUS_ARM struct {
	// Count: The threshold for opening the circuit.
	Count *int `json:"count,omitempty"`

	// ErrorReasons: The error reasons which are considered as failure.
	ErrorReasons []CircuitBreakerFailureCondition_ErrorReasons_STATUS_ARM `json:"errorReasons,omitempty"`

	// Interval: The interval during which the failures are counted.
	Interval *string `json:"interval,omitempty"`

	// Percentage: The threshold for opening the circuit.
	Percentage *int `json:"percentage,omitempty"`

	// StatusCodeRanges: The status code ranges which are considered as failure.
	StatusCodeRanges []FailureStatusCodeRange_STATUS_ARM `json:"statusCodeRanges,omitempty"`
}

// Properties of server X509Names.
type X509CertificateName_STATUS_ARM struct {
	// IssuerCertificateThumbprint: Thumbprint for the Issuer of the Certificate.
	IssuerCertificateThumbprint *string `json:"issuerCertificateThumbprint,omitempty"`

	// Name: Common Name of the Certificate.
	Name *string `json:"name,omitempty"`
}

// The failure http status code range
type FailureStatusCodeRange_STATUS_ARM struct {
	// Max: The maximum http status code.
	Max *int `json:"max,omitempty"`

	// Min: The minimum http status code.
	Min *int `json:"min,omitempty"`
}
