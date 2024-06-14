// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20230501

import (
	"fmt"
	storage "github.com/Azure/azure-service-operator/v2/api/cdn/v1api20230501/storage"
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
// - Generated from: /cdn/resource-manager/Microsoft.Cdn/stable/2023-05-01/afdx.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/ruleSets/{ruleSetName}
type RuleSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Profiles_RuleSet_Spec   `json:"spec,omitempty"`
	Status            Profiles_RuleSet_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &RuleSet{}

// GetConditions returns the conditions of the resource
func (ruleSet *RuleSet) GetConditions() conditions.Conditions {
	return ruleSet.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (ruleSet *RuleSet) SetConditions(conditions conditions.Conditions) {
	ruleSet.Status.Conditions = conditions
}

var _ conversion.Convertible = &RuleSet{}

// ConvertFrom populates our RuleSet from the provided hub RuleSet
func (ruleSet *RuleSet) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*storage.RuleSet)
	if !ok {
		return fmt.Errorf("expected cdn/v1api20230501/storage/RuleSet but received %T instead", hub)
	}

	return ruleSet.AssignProperties_From_RuleSet(source)
}

// ConvertTo populates the provided hub RuleSet from our RuleSet
func (ruleSet *RuleSet) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*storage.RuleSet)
	if !ok {
		return fmt.Errorf("expected cdn/v1api20230501/storage/RuleSet but received %T instead", hub)
	}

	return ruleSet.AssignProperties_To_RuleSet(destination)
}

// +kubebuilder:webhook:path=/mutate-cdn-azure-com-v1api20230501-ruleset,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=cdn.azure.com,resources=rulesets,verbs=create;update,versions=v1api20230501,name=default.v1api20230501.rulesets.cdn.azure.com,admissionReviewVersions=v1

var _ admission.Defaulter = &RuleSet{}

// Default applies defaults to the RuleSet resource
func (ruleSet *RuleSet) Default() {
	ruleSet.defaultImpl()
	var temp any = ruleSet
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (ruleSet *RuleSet) defaultAzureName() {
	if ruleSet.Spec.AzureName == "" {
		ruleSet.Spec.AzureName = ruleSet.Name
	}
}

// defaultImpl applies the code generated defaults to the RuleSet resource
func (ruleSet *RuleSet) defaultImpl() { ruleSet.defaultAzureName() }

var _ genruntime.ImportableResource = &RuleSet{}

// InitializeSpec initializes the spec for this resource from the given status
func (ruleSet *RuleSet) InitializeSpec(status genruntime.ConvertibleStatus) error {
	if s, ok := status.(*Profiles_RuleSet_STATUS); ok {
		return ruleSet.Spec.Initialize_From_Profiles_RuleSet_STATUS(s)
	}

	return fmt.Errorf("expected Status of type Profiles_RuleSet_STATUS but received %T instead", status)
}

var _ genruntime.KubernetesResource = &RuleSet{}

// AzureName returns the Azure name of the resource
func (ruleSet *RuleSet) AzureName() string {
	return ruleSet.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-05-01"
func (ruleSet RuleSet) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (ruleSet *RuleSet) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (ruleSet *RuleSet) GetSpec() genruntime.ConvertibleSpec {
	return &ruleSet.Spec
}

// GetStatus returns the status of this resource
func (ruleSet *RuleSet) GetStatus() genruntime.ConvertibleStatus {
	return &ruleSet.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (ruleSet *RuleSet) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Cdn/profiles/ruleSets"
func (ruleSet *RuleSet) GetType() string {
	return "Microsoft.Cdn/profiles/ruleSets"
}

// NewEmptyStatus returns a new empty (blank) status
func (ruleSet *RuleSet) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Profiles_RuleSet_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (ruleSet *RuleSet) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(ruleSet.Spec)
	return ruleSet.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (ruleSet *RuleSet) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Profiles_RuleSet_STATUS); ok {
		ruleSet.Status = *st
		return nil
	}

	// Convert status to required version
	var st Profiles_RuleSet_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	ruleSet.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-cdn-azure-com-v1api20230501-ruleset,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=cdn.azure.com,resources=rulesets,verbs=create;update,versions=v1api20230501,name=validate.v1api20230501.rulesets.cdn.azure.com,admissionReviewVersions=v1

var _ admission.Validator = &RuleSet{}

// ValidateCreate validates the creation of the resource
func (ruleSet *RuleSet) ValidateCreate() (admission.Warnings, error) {
	validations := ruleSet.createValidations()
	var temp any = ruleSet
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(validations)
}

// ValidateDelete validates the deletion of the resource
func (ruleSet *RuleSet) ValidateDelete() (admission.Warnings, error) {
	validations := ruleSet.deleteValidations()
	var temp any = ruleSet
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(validations)
}

// ValidateUpdate validates an update of the resource
func (ruleSet *RuleSet) ValidateUpdate(old runtime.Object) (admission.Warnings, error) {
	validations := ruleSet.updateValidations()
	var temp any = ruleSet
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(old, validations)
}

// createValidations validates the creation of the resource
func (ruleSet *RuleSet) createValidations() []func() (admission.Warnings, error) {
	return []func() (admission.Warnings, error){ruleSet.validateResourceReferences, ruleSet.validateOwnerReference}
}

// deleteValidations validates the deletion of the resource
func (ruleSet *RuleSet) deleteValidations() []func() (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (ruleSet *RuleSet) updateValidations() []func(old runtime.Object) (admission.Warnings, error) {
	return []func(old runtime.Object) (admission.Warnings, error){
		func(old runtime.Object) (admission.Warnings, error) {
			return ruleSet.validateResourceReferences()
		},
		ruleSet.validateWriteOnceProperties,
		func(old runtime.Object) (admission.Warnings, error) {
			return ruleSet.validateOwnerReference()
		},
	}
}

// validateOwnerReference validates the owner field
func (ruleSet *RuleSet) validateOwnerReference() (admission.Warnings, error) {
	return genruntime.ValidateOwner(ruleSet)
}

// validateResourceReferences validates all resource references
func (ruleSet *RuleSet) validateResourceReferences() (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&ruleSet.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (ruleSet *RuleSet) validateWriteOnceProperties(old runtime.Object) (admission.Warnings, error) {
	oldObj, ok := old.(*RuleSet)
	if !ok {
		return nil, nil
	}

	return genruntime.ValidateWriteOnceProperties(oldObj, ruleSet)
}

// AssignProperties_From_RuleSet populates our RuleSet from the provided source RuleSet
func (ruleSet *RuleSet) AssignProperties_From_RuleSet(source *storage.RuleSet) error {

	// ObjectMeta
	ruleSet.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec Profiles_RuleSet_Spec
	err := spec.AssignProperties_From_Profiles_RuleSet_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_Profiles_RuleSet_Spec() to populate field Spec")
	}
	ruleSet.Spec = spec

	// Status
	var status Profiles_RuleSet_STATUS
	err = status.AssignProperties_From_Profiles_RuleSet_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_Profiles_RuleSet_STATUS() to populate field Status")
	}
	ruleSet.Status = status

	// No error
	return nil
}

// AssignProperties_To_RuleSet populates the provided destination RuleSet from our RuleSet
func (ruleSet *RuleSet) AssignProperties_To_RuleSet(destination *storage.RuleSet) error {

	// ObjectMeta
	destination.ObjectMeta = *ruleSet.ObjectMeta.DeepCopy()

	// Spec
	var spec storage.Profiles_RuleSet_Spec
	err := ruleSet.Spec.AssignProperties_To_Profiles_RuleSet_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_Profiles_RuleSet_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status storage.Profiles_RuleSet_STATUS
	err = ruleSet.Status.AssignProperties_To_Profiles_RuleSet_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_Profiles_RuleSet_STATUS() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (ruleSet *RuleSet) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: ruleSet.Spec.OriginalVersion(),
		Kind:    "RuleSet",
	}
}

// +kubebuilder:object:root=true
// Generator information:
// - Generated from: /cdn/resource-manager/Microsoft.Cdn/stable/2023-05-01/afdx.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/ruleSets/{ruleSetName}
type RuleSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RuleSet `json:"items"`
}

type Profiles_RuleSet_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName string `json:"azureName,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a cdn.azure.com/Profile resource
	Owner *genruntime.KnownResourceReference `group:"cdn.azure.com" json:"owner,omitempty" kind:"Profile"`
}

var _ genruntime.ARMTransformer = &Profiles_RuleSet_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (ruleSet *Profiles_RuleSet_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if ruleSet == nil {
		return nil, nil
	}
	result := &Profiles_RuleSet_Spec_ARM{}

	// Set property "Name":
	result.Name = resolved.Name
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (ruleSet *Profiles_RuleSet_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &Profiles_RuleSet_Spec_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (ruleSet *Profiles_RuleSet_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(Profiles_RuleSet_Spec_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected Profiles_RuleSet_Spec_ARM, got %T", armInput)
	}

	// Set property "AzureName":
	ruleSet.SetAzureName(genruntime.ExtractKubernetesResourceNameFromARMName(typedInput.Name))

	// Set property "Owner":
	ruleSet.Owner = &genruntime.KnownResourceReference{
		Name:  owner.Name,
		ARMID: owner.ARMID,
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &Profiles_RuleSet_Spec{}

// ConvertSpecFrom populates our Profiles_RuleSet_Spec from the provided source
func (ruleSet *Profiles_RuleSet_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*storage.Profiles_RuleSet_Spec)
	if ok {
		// Populate our instance from source
		return ruleSet.AssignProperties_From_Profiles_RuleSet_Spec(src)
	}

	// Convert to an intermediate form
	src = &storage.Profiles_RuleSet_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = ruleSet.AssignProperties_From_Profiles_RuleSet_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our Profiles_RuleSet_Spec
func (ruleSet *Profiles_RuleSet_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*storage.Profiles_RuleSet_Spec)
	if ok {
		// Populate destination from our instance
		return ruleSet.AssignProperties_To_Profiles_RuleSet_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &storage.Profiles_RuleSet_Spec{}
	err := ruleSet.AssignProperties_To_Profiles_RuleSet_Spec(dst)
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

// AssignProperties_From_Profiles_RuleSet_Spec populates our Profiles_RuleSet_Spec from the provided source Profiles_RuleSet_Spec
func (ruleSet *Profiles_RuleSet_Spec) AssignProperties_From_Profiles_RuleSet_Spec(source *storage.Profiles_RuleSet_Spec) error {

	// AzureName
	ruleSet.AzureName = source.AzureName

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		ruleSet.Owner = &owner
	} else {
		ruleSet.Owner = nil
	}

	// No error
	return nil
}

// AssignProperties_To_Profiles_RuleSet_Spec populates the provided destination Profiles_RuleSet_Spec from our Profiles_RuleSet_Spec
func (ruleSet *Profiles_RuleSet_Spec) AssignProperties_To_Profiles_RuleSet_Spec(destination *storage.Profiles_RuleSet_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AzureName
	destination.AzureName = ruleSet.AzureName

	// OriginalVersion
	destination.OriginalVersion = ruleSet.OriginalVersion()

	// Owner
	if ruleSet.Owner != nil {
		owner := ruleSet.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
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

// Initialize_From_Profiles_RuleSet_STATUS populates our Profiles_RuleSet_Spec from the provided source Profiles_RuleSet_STATUS
func (ruleSet *Profiles_RuleSet_Spec) Initialize_From_Profiles_RuleSet_STATUS(source *Profiles_RuleSet_STATUS) error {

	// No error
	return nil
}

// OriginalVersion returns the original API version used to create the resource.
func (ruleSet *Profiles_RuleSet_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

// SetAzureName sets the Azure name of the resource
func (ruleSet *Profiles_RuleSet_Spec) SetAzureName(azureName string) { ruleSet.AzureName = azureName }

type Profiles_RuleSet_STATUS struct {
	// Conditions: The observed state of the resource
	Conditions       []conditions.Condition                     `json:"conditions,omitempty"`
	DeploymentStatus *RuleSetProperties_DeploymentStatus_STATUS `json:"deploymentStatus,omitempty"`

	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Name: Resource name.
	Name *string `json:"name,omitempty"`

	// ProfileName: The name of the profile which holds the rule set.
	ProfileName *string `json:"profileName,omitempty"`

	// ProvisioningState: Provisioning status
	ProvisioningState *RuleSetProperties_ProvisioningState_STATUS `json:"provisioningState,omitempty"`

	// SystemData: Read only system data
	SystemData *SystemData_STATUS `json:"systemData,omitempty"`

	// Type: Resource type.
	Type *string `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Profiles_RuleSet_STATUS{}

// ConvertStatusFrom populates our Profiles_RuleSet_STATUS from the provided source
func (ruleSet *Profiles_RuleSet_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*storage.Profiles_RuleSet_STATUS)
	if ok {
		// Populate our instance from source
		return ruleSet.AssignProperties_From_Profiles_RuleSet_STATUS(src)
	}

	// Convert to an intermediate form
	src = &storage.Profiles_RuleSet_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = ruleSet.AssignProperties_From_Profiles_RuleSet_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our Profiles_RuleSet_STATUS
func (ruleSet *Profiles_RuleSet_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*storage.Profiles_RuleSet_STATUS)
	if ok {
		// Populate destination from our instance
		return ruleSet.AssignProperties_To_Profiles_RuleSet_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &storage.Profiles_RuleSet_STATUS{}
	err := ruleSet.AssignProperties_To_Profiles_RuleSet_STATUS(dst)
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

var _ genruntime.FromARMConverter = &Profiles_RuleSet_STATUS{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (ruleSet *Profiles_RuleSet_STATUS) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &Profiles_RuleSet_STATUS_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (ruleSet *Profiles_RuleSet_STATUS) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(Profiles_RuleSet_STATUS_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected Profiles_RuleSet_STATUS_ARM, got %T", armInput)
	}

	// no assignment for property "Conditions"

	// Set property "DeploymentStatus":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.DeploymentStatus != nil {
			var temp string
			temp = string(*typedInput.Properties.DeploymentStatus)
			deploymentStatus := RuleSetProperties_DeploymentStatus_STATUS(temp)
			ruleSet.DeploymentStatus = &deploymentStatus
		}
	}

	// Set property "Id":
	if typedInput.Id != nil {
		id := *typedInput.Id
		ruleSet.Id = &id
	}

	// Set property "Name":
	if typedInput.Name != nil {
		name := *typedInput.Name
		ruleSet.Name = &name
	}

	// Set property "ProfileName":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.ProfileName != nil {
			profileName := *typedInput.Properties.ProfileName
			ruleSet.ProfileName = &profileName
		}
	}

	// Set property "ProvisioningState":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.ProvisioningState != nil {
			var temp string
			temp = string(*typedInput.Properties.ProvisioningState)
			provisioningState := RuleSetProperties_ProvisioningState_STATUS(temp)
			ruleSet.ProvisioningState = &provisioningState
		}
	}

	// Set property "SystemData":
	if typedInput.SystemData != nil {
		var systemData1 SystemData_STATUS
		err := systemData1.PopulateFromARM(owner, *typedInput.SystemData)
		if err != nil {
			return err
		}
		systemData := systemData1
		ruleSet.SystemData = &systemData
	}

	// Set property "Type":
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		ruleSet.Type = &typeVar
	}

	// No error
	return nil
}

// AssignProperties_From_Profiles_RuleSet_STATUS populates our Profiles_RuleSet_STATUS from the provided source Profiles_RuleSet_STATUS
func (ruleSet *Profiles_RuleSet_STATUS) AssignProperties_From_Profiles_RuleSet_STATUS(source *storage.Profiles_RuleSet_STATUS) error {

	// Conditions
	ruleSet.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// DeploymentStatus
	if source.DeploymentStatus != nil {
		deploymentStatus := *source.DeploymentStatus
		deploymentStatusTemp := genruntime.ToEnum(deploymentStatus, ruleSetProperties_DeploymentStatus_STATUS_Values)
		ruleSet.DeploymentStatus = &deploymentStatusTemp
	} else {
		ruleSet.DeploymentStatus = nil
	}

	// Id
	ruleSet.Id = genruntime.ClonePointerToString(source.Id)

	// Name
	ruleSet.Name = genruntime.ClonePointerToString(source.Name)

	// ProfileName
	ruleSet.ProfileName = genruntime.ClonePointerToString(source.ProfileName)

	// ProvisioningState
	if source.ProvisioningState != nil {
		provisioningState := *source.ProvisioningState
		provisioningStateTemp := genruntime.ToEnum(provisioningState, ruleSetProperties_ProvisioningState_STATUS_Values)
		ruleSet.ProvisioningState = &provisioningStateTemp
	} else {
		ruleSet.ProvisioningState = nil
	}

	// SystemData
	if source.SystemData != nil {
		var systemDatum SystemData_STATUS
		err := systemDatum.AssignProperties_From_SystemData_STATUS(source.SystemData)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_SystemData_STATUS() to populate field SystemData")
		}
		ruleSet.SystemData = &systemDatum
	} else {
		ruleSet.SystemData = nil
	}

	// Type
	ruleSet.Type = genruntime.ClonePointerToString(source.Type)

	// No error
	return nil
}

// AssignProperties_To_Profiles_RuleSet_STATUS populates the provided destination Profiles_RuleSet_STATUS from our Profiles_RuleSet_STATUS
func (ruleSet *Profiles_RuleSet_STATUS) AssignProperties_To_Profiles_RuleSet_STATUS(destination *storage.Profiles_RuleSet_STATUS) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(ruleSet.Conditions)

	// DeploymentStatus
	if ruleSet.DeploymentStatus != nil {
		deploymentStatus := string(*ruleSet.DeploymentStatus)
		destination.DeploymentStatus = &deploymentStatus
	} else {
		destination.DeploymentStatus = nil
	}

	// Id
	destination.Id = genruntime.ClonePointerToString(ruleSet.Id)

	// Name
	destination.Name = genruntime.ClonePointerToString(ruleSet.Name)

	// ProfileName
	destination.ProfileName = genruntime.ClonePointerToString(ruleSet.ProfileName)

	// ProvisioningState
	if ruleSet.ProvisioningState != nil {
		provisioningState := string(*ruleSet.ProvisioningState)
		destination.ProvisioningState = &provisioningState
	} else {
		destination.ProvisioningState = nil
	}

	// SystemData
	if ruleSet.SystemData != nil {
		var systemDatum storage.SystemData_STATUS
		err := ruleSet.SystemData.AssignProperties_To_SystemData_STATUS(&systemDatum)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_SystemData_STATUS() to populate field SystemData")
		}
		destination.SystemData = &systemDatum
	} else {
		destination.SystemData = nil
	}

	// Type
	destination.Type = genruntime.ClonePointerToString(ruleSet.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

type RuleSetProperties_DeploymentStatus_STATUS string

const (
	RuleSetProperties_DeploymentStatus_STATUS_Failed     = RuleSetProperties_DeploymentStatus_STATUS("Failed")
	RuleSetProperties_DeploymentStatus_STATUS_InProgress = RuleSetProperties_DeploymentStatus_STATUS("InProgress")
	RuleSetProperties_DeploymentStatus_STATUS_NotStarted = RuleSetProperties_DeploymentStatus_STATUS("NotStarted")
	RuleSetProperties_DeploymentStatus_STATUS_Succeeded  = RuleSetProperties_DeploymentStatus_STATUS("Succeeded")
)

// Mapping from string to RuleSetProperties_DeploymentStatus_STATUS
var ruleSetProperties_DeploymentStatus_STATUS_Values = map[string]RuleSetProperties_DeploymentStatus_STATUS{
	"failed":     RuleSetProperties_DeploymentStatus_STATUS_Failed,
	"inprogress": RuleSetProperties_DeploymentStatus_STATUS_InProgress,
	"notstarted": RuleSetProperties_DeploymentStatus_STATUS_NotStarted,
	"succeeded":  RuleSetProperties_DeploymentStatus_STATUS_Succeeded,
}

type RuleSetProperties_ProvisioningState_STATUS string

const (
	RuleSetProperties_ProvisioningState_STATUS_Creating  = RuleSetProperties_ProvisioningState_STATUS("Creating")
	RuleSetProperties_ProvisioningState_STATUS_Deleting  = RuleSetProperties_ProvisioningState_STATUS("Deleting")
	RuleSetProperties_ProvisioningState_STATUS_Failed    = RuleSetProperties_ProvisioningState_STATUS("Failed")
	RuleSetProperties_ProvisioningState_STATUS_Succeeded = RuleSetProperties_ProvisioningState_STATUS("Succeeded")
	RuleSetProperties_ProvisioningState_STATUS_Updating  = RuleSetProperties_ProvisioningState_STATUS("Updating")
)

// Mapping from string to RuleSetProperties_ProvisioningState_STATUS
var ruleSetProperties_ProvisioningState_STATUS_Values = map[string]RuleSetProperties_ProvisioningState_STATUS{
	"creating":  RuleSetProperties_ProvisioningState_STATUS_Creating,
	"deleting":  RuleSetProperties_ProvisioningState_STATUS_Deleting,
	"failed":    RuleSetProperties_ProvisioningState_STATUS_Failed,
	"succeeded": RuleSetProperties_ProvisioningState_STATUS_Succeeded,
	"updating":  RuleSetProperties_ProvisioningState_STATUS_Updating,
}

func init() {
	SchemeBuilder.Register(&RuleSet{}, &RuleSetList{})
}
