// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20201201

import (
	"fmt"
	storage "github.com/Azure/azure-service-operator/v2/api/cache/v1api20201201/storage"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Generator information:
// - Generated from: /redis/resource-manager/Microsoft.Cache/stable/2020-12-01/redis.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redis/{name}/linkedServers/{linkedServerName}
type RedisLinkedServer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Redis_LinkedServer_Spec   `json:"spec,omitempty"`
	Status            Redis_LinkedServer_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &RedisLinkedServer{}

// GetConditions returns the conditions of the resource
func (server *RedisLinkedServer) GetConditions() conditions.Conditions {
	return server.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (server *RedisLinkedServer) SetConditions(conditions conditions.Conditions) {
	server.Status.Conditions = conditions
}

var _ conversion.Convertible = &RedisLinkedServer{}

// ConvertFrom populates our RedisLinkedServer from the provided hub RedisLinkedServer
func (server *RedisLinkedServer) ConvertFrom(hub conversion.Hub) error {
	// intermediate variable for conversion
	var source storage.RedisLinkedServer

	err := source.ConvertFrom(hub)
	if err != nil {
		return errors.Wrap(err, "converting from hub to source")
	}

	err = server.AssignProperties_From_RedisLinkedServer(&source)
	if err != nil {
		return errors.Wrap(err, "converting from source to server")
	}

	return nil
}

// ConvertTo populates the provided hub RedisLinkedServer from our RedisLinkedServer
func (server *RedisLinkedServer) ConvertTo(hub conversion.Hub) error {
	// intermediate variable for conversion
	var destination storage.RedisLinkedServer
	err := server.AssignProperties_To_RedisLinkedServer(&destination)
	if err != nil {
		return errors.Wrap(err, "converting to destination from server")
	}
	err = destination.ConvertTo(hub)
	if err != nil {
		return errors.Wrap(err, "converting from destination to hub")
	}

	return nil
}

// +kubebuilder:webhook:path=/mutate-cache-azure-com-v1api20201201-redislinkedserver,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=cache.azure.com,resources=redislinkedservers,verbs=create;update,versions=v1api20201201,name=default.v1api20201201.redislinkedservers.cache.azure.com,admissionReviewVersions=v1

var _ admission.Defaulter = &RedisLinkedServer{}

// Default applies defaults to the RedisLinkedServer resource
func (server *RedisLinkedServer) Default() {
	server.defaultImpl()
	var temp any = server
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (server *RedisLinkedServer) defaultAzureName() {
	if server.Spec.AzureName == "" {
		server.Spec.AzureName = server.Name
	}
}

// defaultImpl applies the code generated defaults to the RedisLinkedServer resource
func (server *RedisLinkedServer) defaultImpl() { server.defaultAzureName() }

var _ genruntime.KubernetesResource = &RedisLinkedServer{}

// AzureName returns the Azure name of the resource
func (server *RedisLinkedServer) AzureName() string {
	return server.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-12-01"
func (server RedisLinkedServer) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (server *RedisLinkedServer) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (server *RedisLinkedServer) GetSpec() genruntime.ConvertibleSpec {
	return &server.Spec
}

// GetStatus returns the status of this resource
func (server *RedisLinkedServer) GetStatus() genruntime.ConvertibleStatus {
	return &server.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (server *RedisLinkedServer) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Cache/redis/linkedServers"
func (server *RedisLinkedServer) GetType() string {
	return "Microsoft.Cache/redis/linkedServers"
}

// NewEmptyStatus returns a new empty (blank) status
func (server *RedisLinkedServer) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Redis_LinkedServer_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (server *RedisLinkedServer) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(server.Spec)
	return server.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (server *RedisLinkedServer) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Redis_LinkedServer_STATUS); ok {
		server.Status = *st
		return nil
	}

	// Convert status to required version
	var st Redis_LinkedServer_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	server.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-cache-azure-com-v1api20201201-redislinkedserver,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=cache.azure.com,resources=redislinkedservers,verbs=create;update,versions=v1api20201201,name=validate.v1api20201201.redislinkedservers.cache.azure.com,admissionReviewVersions=v1

var _ admission.Validator = &RedisLinkedServer{}

// ValidateCreate validates the creation of the resource
func (server *RedisLinkedServer) ValidateCreate() (admission.Warnings, error) {
	validations := server.createValidations()
	var temp any = server
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(validations)
}

// ValidateDelete validates the deletion of the resource
func (server *RedisLinkedServer) ValidateDelete() (admission.Warnings, error) {
	validations := server.deleteValidations()
	var temp any = server
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(validations)
}

// ValidateUpdate validates an update of the resource
func (server *RedisLinkedServer) ValidateUpdate(old runtime.Object) (admission.Warnings, error) {
	validations := server.updateValidations()
	var temp any = server
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(old, validations)
}

// createValidations validates the creation of the resource
func (server *RedisLinkedServer) createValidations() []func() (admission.Warnings, error) {
	return []func() (admission.Warnings, error){server.validateResourceReferences, server.validateOwnerReference}
}

// deleteValidations validates the deletion of the resource
func (server *RedisLinkedServer) deleteValidations() []func() (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (server *RedisLinkedServer) updateValidations() []func(old runtime.Object) (admission.Warnings, error) {
	return []func(old runtime.Object) (admission.Warnings, error){
		func(old runtime.Object) (admission.Warnings, error) {
			return server.validateResourceReferences()
		},
		server.validateWriteOnceProperties,
		func(old runtime.Object) (admission.Warnings, error) {
			return server.validateOwnerReference()
		},
	}
}

// validateOwnerReference validates the owner field
func (server *RedisLinkedServer) validateOwnerReference() (admission.Warnings, error) {
	return genruntime.ValidateOwner(server)
}

// validateResourceReferences validates all resource references
func (server *RedisLinkedServer) validateResourceReferences() (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&server.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (server *RedisLinkedServer) validateWriteOnceProperties(old runtime.Object) (admission.Warnings, error) {
	oldObj, ok := old.(*RedisLinkedServer)
	if !ok {
		return nil, nil
	}

	return genruntime.ValidateWriteOnceProperties(oldObj, server)
}

// AssignProperties_From_RedisLinkedServer populates our RedisLinkedServer from the provided source RedisLinkedServer
func (server *RedisLinkedServer) AssignProperties_From_RedisLinkedServer(source *storage.RedisLinkedServer) error {

	// ObjectMeta
	server.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec Redis_LinkedServer_Spec
	err := spec.AssignProperties_From_Redis_LinkedServer_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_Redis_LinkedServer_Spec() to populate field Spec")
	}
	server.Spec = spec

	// Status
	var status Redis_LinkedServer_STATUS
	err = status.AssignProperties_From_Redis_LinkedServer_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_Redis_LinkedServer_STATUS() to populate field Status")
	}
	server.Status = status

	// No error
	return nil
}

// AssignProperties_To_RedisLinkedServer populates the provided destination RedisLinkedServer from our RedisLinkedServer
func (server *RedisLinkedServer) AssignProperties_To_RedisLinkedServer(destination *storage.RedisLinkedServer) error {

	// ObjectMeta
	destination.ObjectMeta = *server.ObjectMeta.DeepCopy()

	// Spec
	var spec storage.Redis_LinkedServer_Spec
	err := server.Spec.AssignProperties_To_Redis_LinkedServer_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_Redis_LinkedServer_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status storage.Redis_LinkedServer_STATUS
	err = server.Status.AssignProperties_To_Redis_LinkedServer_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_Redis_LinkedServer_STATUS() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (server *RedisLinkedServer) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: server.Spec.OriginalVersion(),
		Kind:    "RedisLinkedServer",
	}
}

// +kubebuilder:object:root=true
// Generator information:
// - Generated from: /redis/resource-manager/Microsoft.Cache/stable/2020-12-01/redis.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redis/{name}/linkedServers/{linkedServerName}
type RedisLinkedServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RedisLinkedServer `json:"items"`
}

type Redis_LinkedServer_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName string `json:"azureName,omitempty"`

	// +kubebuilder:validation:Required
	// LinkedRedisCacheLocation: Location of the linked redis cache.
	LinkedRedisCacheLocation *string `json:"linkedRedisCacheLocation,omitempty"`

	// +kubebuilder:validation:Required
	// LinkedRedisCacheReference: Fully qualified resourceId of the linked redis cache.
	LinkedRedisCacheReference *genruntime.ResourceReference `armReference:"LinkedRedisCacheId" json:"linkedRedisCacheReference,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a cache.azure.com/Redis resource
	Owner *genruntime.KnownResourceReference `group:"cache.azure.com" json:"owner,omitempty" kind:"Redis"`

	// +kubebuilder:validation:Required
	// ServerRole: Role of the linked server.
	ServerRole *RedisLinkedServerCreateProperties_ServerRole `json:"serverRole,omitempty"`
}

var _ genruntime.ARMTransformer = &Redis_LinkedServer_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (server *Redis_LinkedServer_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if server == nil {
		return nil, nil
	}
	result := &Redis_LinkedServer_Spec_ARM{}

	// Set property "Name":
	result.Name = resolved.Name

	// Set property "Properties":
	if server.LinkedRedisCacheLocation != nil ||
		server.LinkedRedisCacheReference != nil ||
		server.ServerRole != nil {
		result.Properties = &RedisLinkedServerCreateProperties_ARM{}
	}
	if server.LinkedRedisCacheLocation != nil {
		linkedRedisCacheLocation := *server.LinkedRedisCacheLocation
		result.Properties.LinkedRedisCacheLocation = &linkedRedisCacheLocation
	}
	if server.LinkedRedisCacheReference != nil {
		linkedRedisCacheIdARMID, err := resolved.ResolvedReferences.Lookup(*server.LinkedRedisCacheReference)
		if err != nil {
			return nil, err
		}
		linkedRedisCacheId := linkedRedisCacheIdARMID
		result.Properties.LinkedRedisCacheId = &linkedRedisCacheId
	}
	if server.ServerRole != nil {
		var temp string
		temp = string(*server.ServerRole)
		serverRole := RedisLinkedServerCreateProperties_ServerRole_ARM(temp)
		result.Properties.ServerRole = &serverRole
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (server *Redis_LinkedServer_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &Redis_LinkedServer_Spec_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (server *Redis_LinkedServer_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(Redis_LinkedServer_Spec_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected Redis_LinkedServer_Spec_ARM, got %T", armInput)
	}

	// Set property "AzureName":
	server.SetAzureName(genruntime.ExtractKubernetesResourceNameFromARMName(typedInput.Name))

	// Set property "LinkedRedisCacheLocation":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.LinkedRedisCacheLocation != nil {
			linkedRedisCacheLocation := *typedInput.Properties.LinkedRedisCacheLocation
			server.LinkedRedisCacheLocation = &linkedRedisCacheLocation
		}
	}

	// no assignment for property "LinkedRedisCacheReference"

	// Set property "Owner":
	server.Owner = &genruntime.KnownResourceReference{
		Name:  owner.Name,
		ARMID: owner.ARMID,
	}

	// Set property "ServerRole":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.ServerRole != nil {
			var temp string
			temp = string(*typedInput.Properties.ServerRole)
			serverRole := RedisLinkedServerCreateProperties_ServerRole(temp)
			server.ServerRole = &serverRole
		}
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &Redis_LinkedServer_Spec{}

// ConvertSpecFrom populates our Redis_LinkedServer_Spec from the provided source
func (server *Redis_LinkedServer_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*storage.Redis_LinkedServer_Spec)
	if ok {
		// Populate our instance from source
		return server.AssignProperties_From_Redis_LinkedServer_Spec(src)
	}

	// Convert to an intermediate form
	src = &storage.Redis_LinkedServer_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = server.AssignProperties_From_Redis_LinkedServer_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our Redis_LinkedServer_Spec
func (server *Redis_LinkedServer_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*storage.Redis_LinkedServer_Spec)
	if ok {
		// Populate destination from our instance
		return server.AssignProperties_To_Redis_LinkedServer_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &storage.Redis_LinkedServer_Spec{}
	err := server.AssignProperties_To_Redis_LinkedServer_Spec(dst)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecTo()")
	}

	// Update dst from our instance
	err = dst.ConvertSpecTo(destination)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecTo()")
	}

	return nil
}

// AssignProperties_From_Redis_LinkedServer_Spec populates our Redis_LinkedServer_Spec from the provided source Redis_LinkedServer_Spec
func (server *Redis_LinkedServer_Spec) AssignProperties_From_Redis_LinkedServer_Spec(source *storage.Redis_LinkedServer_Spec) error {

	// AzureName
	server.AzureName = source.AzureName

	// LinkedRedisCacheLocation
	server.LinkedRedisCacheLocation = genruntime.ClonePointerToString(source.LinkedRedisCacheLocation)

	// LinkedRedisCacheReference
	if source.LinkedRedisCacheReference != nil {
		linkedRedisCacheReference := source.LinkedRedisCacheReference.Copy()
		server.LinkedRedisCacheReference = &linkedRedisCacheReference
	} else {
		server.LinkedRedisCacheReference = nil
	}

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		server.Owner = &owner
	} else {
		server.Owner = nil
	}

	// ServerRole
	if source.ServerRole != nil {
		serverRole := *source.ServerRole
		serverRoleTemp := genruntime.ToEnum(serverRole, redisLinkedServerCreateProperties_ServerRole_Values)
		server.ServerRole = &serverRoleTemp
	} else {
		server.ServerRole = nil
	}

	// No error
	return nil
}

// AssignProperties_To_Redis_LinkedServer_Spec populates the provided destination Redis_LinkedServer_Spec from our Redis_LinkedServer_Spec
func (server *Redis_LinkedServer_Spec) AssignProperties_To_Redis_LinkedServer_Spec(destination *storage.Redis_LinkedServer_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AzureName
	destination.AzureName = server.AzureName

	// LinkedRedisCacheLocation
	destination.LinkedRedisCacheLocation = genruntime.ClonePointerToString(server.LinkedRedisCacheLocation)

	// LinkedRedisCacheReference
	if server.LinkedRedisCacheReference != nil {
		linkedRedisCacheReference := server.LinkedRedisCacheReference.Copy()
		destination.LinkedRedisCacheReference = &linkedRedisCacheReference
	} else {
		destination.LinkedRedisCacheReference = nil
	}

	// OriginalVersion
	destination.OriginalVersion = server.OriginalVersion()

	// Owner
	if server.Owner != nil {
		owner := server.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// ServerRole
	if server.ServerRole != nil {
		serverRole := string(*server.ServerRole)
		destination.ServerRole = &serverRole
	} else {
		destination.ServerRole = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// OriginalVersion returns the original API version used to create the resource.
func (server *Redis_LinkedServer_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

// SetAzureName sets the Azure name of the resource
func (server *Redis_LinkedServer_Spec) SetAzureName(azureName string) { server.AzureName = azureName }

type Redis_LinkedServer_STATUS struct {
	// Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`

	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// LinkedRedisCacheId: Fully qualified resourceId of the linked redis cache.
	LinkedRedisCacheId *string `json:"linkedRedisCacheId,omitempty"`

	// LinkedRedisCacheLocation: Location of the linked redis cache.
	LinkedRedisCacheLocation *string `json:"linkedRedisCacheLocation,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// ProvisioningState: Terminal state of the link between primary and secondary redis cache.
	ProvisioningState *string `json:"provisioningState,omitempty"`

	// ServerRole: Role of the linked server.
	ServerRole *RedisLinkedServerProperties_ServerRole_STATUS `json:"serverRole,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Redis_LinkedServer_STATUS{}

// ConvertStatusFrom populates our Redis_LinkedServer_STATUS from the provided source
func (server *Redis_LinkedServer_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*storage.Redis_LinkedServer_STATUS)
	if ok {
		// Populate our instance from source
		return server.AssignProperties_From_Redis_LinkedServer_STATUS(src)
	}

	// Convert to an intermediate form
	src = &storage.Redis_LinkedServer_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = server.AssignProperties_From_Redis_LinkedServer_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our Redis_LinkedServer_STATUS
func (server *Redis_LinkedServer_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*storage.Redis_LinkedServer_STATUS)
	if ok {
		// Populate destination from our instance
		return server.AssignProperties_To_Redis_LinkedServer_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &storage.Redis_LinkedServer_STATUS{}
	err := server.AssignProperties_To_Redis_LinkedServer_STATUS(dst)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusTo()")
	}

	// Update dst from our instance
	err = dst.ConvertStatusTo(destination)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusTo()")
	}

	return nil
}

var _ genruntime.FromARMConverter = &Redis_LinkedServer_STATUS{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (server *Redis_LinkedServer_STATUS) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &Redis_LinkedServer_STATUS_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (server *Redis_LinkedServer_STATUS) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(Redis_LinkedServer_STATUS_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected Redis_LinkedServer_STATUS_ARM, got %T", armInput)
	}

	// no assignment for property "Conditions"

	// Set property "Id":
	if typedInput.Id != nil {
		id := *typedInput.Id
		server.Id = &id
	}

	// Set property "LinkedRedisCacheId":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.LinkedRedisCacheId != nil {
			linkedRedisCacheId := *typedInput.Properties.LinkedRedisCacheId
			server.LinkedRedisCacheId = &linkedRedisCacheId
		}
	}

	// Set property "LinkedRedisCacheLocation":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.LinkedRedisCacheLocation != nil {
			linkedRedisCacheLocation := *typedInput.Properties.LinkedRedisCacheLocation
			server.LinkedRedisCacheLocation = &linkedRedisCacheLocation
		}
	}

	// Set property "Name":
	if typedInput.Name != nil {
		name := *typedInput.Name
		server.Name = &name
	}

	// Set property "ProvisioningState":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.ProvisioningState != nil {
			provisioningState := *typedInput.Properties.ProvisioningState
			server.ProvisioningState = &provisioningState
		}
	}

	// Set property "ServerRole":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.ServerRole != nil {
			var temp string
			temp = string(*typedInput.Properties.ServerRole)
			serverRole := RedisLinkedServerProperties_ServerRole_STATUS(temp)
			server.ServerRole = &serverRole
		}
	}

	// Set property "Type":
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		server.Type = &typeVar
	}

	// No error
	return nil
}

// AssignProperties_From_Redis_LinkedServer_STATUS populates our Redis_LinkedServer_STATUS from the provided source Redis_LinkedServer_STATUS
func (server *Redis_LinkedServer_STATUS) AssignProperties_From_Redis_LinkedServer_STATUS(source *storage.Redis_LinkedServer_STATUS) error {

	// Conditions
	server.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Id
	server.Id = genruntime.ClonePointerToString(source.Id)

	// LinkedRedisCacheId
	server.LinkedRedisCacheId = genruntime.ClonePointerToString(source.LinkedRedisCacheId)

	// LinkedRedisCacheLocation
	server.LinkedRedisCacheLocation = genruntime.ClonePointerToString(source.LinkedRedisCacheLocation)

	// Name
	server.Name = genruntime.ClonePointerToString(source.Name)

	// ProvisioningState
	server.ProvisioningState = genruntime.ClonePointerToString(source.ProvisioningState)

	// ServerRole
	if source.ServerRole != nil {
		serverRole := *source.ServerRole
		serverRoleTemp := genruntime.ToEnum(serverRole, redisLinkedServerProperties_ServerRole_STATUS_Values)
		server.ServerRole = &serverRoleTemp
	} else {
		server.ServerRole = nil
	}

	// Type
	server.Type = genruntime.ClonePointerToString(source.Type)

	// No error
	return nil
}

// AssignProperties_To_Redis_LinkedServer_STATUS populates the provided destination Redis_LinkedServer_STATUS from our Redis_LinkedServer_STATUS
func (server *Redis_LinkedServer_STATUS) AssignProperties_To_Redis_LinkedServer_STATUS(destination *storage.Redis_LinkedServer_STATUS) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(server.Conditions)

	// Id
	destination.Id = genruntime.ClonePointerToString(server.Id)

	// LinkedRedisCacheId
	destination.LinkedRedisCacheId = genruntime.ClonePointerToString(server.LinkedRedisCacheId)

	// LinkedRedisCacheLocation
	destination.LinkedRedisCacheLocation = genruntime.ClonePointerToString(server.LinkedRedisCacheLocation)

	// Name
	destination.Name = genruntime.ClonePointerToString(server.Name)

	// ProvisioningState
	destination.ProvisioningState = genruntime.ClonePointerToString(server.ProvisioningState)

	// ServerRole
	if server.ServerRole != nil {
		serverRole := string(*server.ServerRole)
		destination.ServerRole = &serverRole
	} else {
		destination.ServerRole = nil
	}

	// Type
	destination.Type = genruntime.ClonePointerToString(server.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// +kubebuilder:validation:Enum={"Primary","Secondary"}
type RedisLinkedServerCreateProperties_ServerRole string

const (
	RedisLinkedServerCreateProperties_ServerRole_Primary   = RedisLinkedServerCreateProperties_ServerRole("Primary")
	RedisLinkedServerCreateProperties_ServerRole_Secondary = RedisLinkedServerCreateProperties_ServerRole("Secondary")
)

// Mapping from string to RedisLinkedServerCreateProperties_ServerRole
var redisLinkedServerCreateProperties_ServerRole_Values = map[string]RedisLinkedServerCreateProperties_ServerRole{
	"primary":   RedisLinkedServerCreateProperties_ServerRole_Primary,
	"secondary": RedisLinkedServerCreateProperties_ServerRole_Secondary,
}

type RedisLinkedServerProperties_ServerRole_STATUS string

const (
	RedisLinkedServerProperties_ServerRole_STATUS_Primary   = RedisLinkedServerProperties_ServerRole_STATUS("Primary")
	RedisLinkedServerProperties_ServerRole_STATUS_Secondary = RedisLinkedServerProperties_ServerRole_STATUS("Secondary")
)

// Mapping from string to RedisLinkedServerProperties_ServerRole_STATUS
var redisLinkedServerProperties_ServerRole_STATUS_Values = map[string]RedisLinkedServerProperties_ServerRole_STATUS{
	"primary":   RedisLinkedServerProperties_ServerRole_STATUS_Primary,
	"secondary": RedisLinkedServerProperties_ServerRole_STATUS_Secondary,
}

func init() {
	SchemeBuilder.Register(&RedisLinkedServer{}, &RedisLinkedServerList{})
}
