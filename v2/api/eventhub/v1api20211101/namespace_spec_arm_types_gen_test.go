// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20211101

import (
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/kr/pretty"
	"github.com/kylelemons/godebug/diff"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"os"
	"reflect"
	"testing"
)

func Test_Encryption_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Encryption_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForEncryption_ARM, Encryption_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForEncryption_ARM runs a test to see if a specific instance of Encryption_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForEncryption_ARM(subject Encryption_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Encryption_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of Encryption_ARM instances for property testing - lazily instantiated by Encryption_ARMGenerator()
var encryption_ARMGenerator gopter.Gen

// Encryption_ARMGenerator returns a generator of Encryption_ARM instances for property testing.
// We first initialize encryption_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Encryption_ARMGenerator() gopter.Gen {
	if encryption_ARMGenerator != nil {
		return encryption_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEncryption_ARM(generators)
	encryption_ARMGenerator = gen.Struct(reflect.TypeOf(Encryption_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEncryption_ARM(generators)
	AddRelatedPropertyGeneratorsForEncryption_ARM(generators)
	encryption_ARMGenerator = gen.Struct(reflect.TypeOf(Encryption_ARM{}), generators)

	return encryption_ARMGenerator
}

// AddIndependentPropertyGeneratorsForEncryption_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForEncryption_ARM(gens map[string]gopter.Gen) {
	gens["KeySource"] = gen.PtrOf(gen.OneConstOf(Encryption_KeySource_ARM_MicrosoftKeyVault))
	gens["RequireInfrastructureEncryption"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForEncryption_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForEncryption_ARM(gens map[string]gopter.Gen) {
	gens["KeyVaultProperties"] = gen.SliceOf(KeyVaultProperties_ARMGenerator())
}

func Test_Identity_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Identity_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForIdentity_ARM, Identity_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForIdentity_ARM runs a test to see if a specific instance of Identity_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForIdentity_ARM(subject Identity_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Identity_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of Identity_ARM instances for property testing - lazily instantiated by Identity_ARMGenerator()
var identity_ARMGenerator gopter.Gen

// Identity_ARMGenerator returns a generator of Identity_ARM instances for property testing.
// We first initialize identity_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Identity_ARMGenerator() gopter.Gen {
	if identity_ARMGenerator != nil {
		return identity_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIdentity_ARM(generators)
	identity_ARMGenerator = gen.Struct(reflect.TypeOf(Identity_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIdentity_ARM(generators)
	AddRelatedPropertyGeneratorsForIdentity_ARM(generators)
	identity_ARMGenerator = gen.Struct(reflect.TypeOf(Identity_ARM{}), generators)

	return identity_ARMGenerator
}

// AddIndependentPropertyGeneratorsForIdentity_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForIdentity_ARM(gens map[string]gopter.Gen) {
	gens["Type"] = gen.PtrOf(gen.OneConstOf(
		Identity_Type_ARM_None,
		Identity_Type_ARM_SystemAssigned,
		Identity_Type_ARM_SystemAssignedUserAssigned,
		Identity_Type_ARM_UserAssigned))
}

// AddRelatedPropertyGeneratorsForIdentity_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForIdentity_ARM(gens map[string]gopter.Gen) {
	gens["UserAssignedIdentities"] = gen.MapOf(
		gen.AlphaString(),
		UserAssignedIdentityDetails_ARMGenerator())
}

func Test_KeyVaultProperties_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of KeyVaultProperties_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForKeyVaultProperties_ARM, KeyVaultProperties_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForKeyVaultProperties_ARM runs a test to see if a specific instance of KeyVaultProperties_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForKeyVaultProperties_ARM(subject KeyVaultProperties_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual KeyVaultProperties_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of KeyVaultProperties_ARM instances for property testing - lazily instantiated by
// KeyVaultProperties_ARMGenerator()
var keyVaultProperties_ARMGenerator gopter.Gen

// KeyVaultProperties_ARMGenerator returns a generator of KeyVaultProperties_ARM instances for property testing.
// We first initialize keyVaultProperties_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func KeyVaultProperties_ARMGenerator() gopter.Gen {
	if keyVaultProperties_ARMGenerator != nil {
		return keyVaultProperties_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForKeyVaultProperties_ARM(generators)
	keyVaultProperties_ARMGenerator = gen.Struct(reflect.TypeOf(KeyVaultProperties_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForKeyVaultProperties_ARM(generators)
	AddRelatedPropertyGeneratorsForKeyVaultProperties_ARM(generators)
	keyVaultProperties_ARMGenerator = gen.Struct(reflect.TypeOf(KeyVaultProperties_ARM{}), generators)

	return keyVaultProperties_ARMGenerator
}

// AddIndependentPropertyGeneratorsForKeyVaultProperties_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForKeyVaultProperties_ARM(gens map[string]gopter.Gen) {
	gens["KeyName"] = gen.PtrOf(gen.AlphaString())
	gens["KeyVaultUri"] = gen.PtrOf(gen.AlphaString())
	gens["KeyVersion"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForKeyVaultProperties_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForKeyVaultProperties_ARM(gens map[string]gopter.Gen) {
	gens["Identity"] = gen.PtrOf(UserAssignedIdentityProperties_ARMGenerator())
}

func Test_Namespace_Properties_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Namespace_Properties_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespace_Properties_Spec_ARM, Namespace_Properties_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespace_Properties_Spec_ARM runs a test to see if a specific instance of Namespace_Properties_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespace_Properties_Spec_ARM(subject Namespace_Properties_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Namespace_Properties_Spec_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of Namespace_Properties_Spec_ARM instances for property testing - lazily instantiated by
// Namespace_Properties_Spec_ARMGenerator()
var namespace_Properties_Spec_ARMGenerator gopter.Gen

// Namespace_Properties_Spec_ARMGenerator returns a generator of Namespace_Properties_Spec_ARM instances for property testing.
// We first initialize namespace_Properties_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Namespace_Properties_Spec_ARMGenerator() gopter.Gen {
	if namespace_Properties_Spec_ARMGenerator != nil {
		return namespace_Properties_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespace_Properties_Spec_ARM(generators)
	namespace_Properties_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(Namespace_Properties_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespace_Properties_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForNamespace_Properties_Spec_ARM(generators)
	namespace_Properties_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(Namespace_Properties_Spec_ARM{}), generators)

	return namespace_Properties_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForNamespace_Properties_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespace_Properties_Spec_ARM(gens map[string]gopter.Gen) {
	gens["AlternateName"] = gen.PtrOf(gen.AlphaString())
	gens["ClusterArmId"] = gen.PtrOf(gen.AlphaString())
	gens["DisableLocalAuth"] = gen.PtrOf(gen.Bool())
	gens["IsAutoInflateEnabled"] = gen.PtrOf(gen.Bool())
	gens["KafkaEnabled"] = gen.PtrOf(gen.Bool())
	gens["MaximumThroughputUnits"] = gen.PtrOf(gen.Int())
	gens["ZoneRedundant"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForNamespace_Properties_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespace_Properties_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Encryption"] = gen.PtrOf(Encryption_ARMGenerator())
}

func Test_Namespace_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Namespace_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespace_Spec_ARM, Namespace_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespace_Spec_ARM runs a test to see if a specific instance of Namespace_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespace_Spec_ARM(subject Namespace_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Namespace_Spec_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of Namespace_Spec_ARM instances for property testing - lazily instantiated by Namespace_Spec_ARMGenerator()
var namespace_Spec_ARMGenerator gopter.Gen

// Namespace_Spec_ARMGenerator returns a generator of Namespace_Spec_ARM instances for property testing.
// We first initialize namespace_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Namespace_Spec_ARMGenerator() gopter.Gen {
	if namespace_Spec_ARMGenerator != nil {
		return namespace_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespace_Spec_ARM(generators)
	namespace_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(Namespace_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespace_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForNamespace_Spec_ARM(generators)
	namespace_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(Namespace_Spec_ARM{}), generators)

	return namespace_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForNamespace_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespace_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForNamespace_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespace_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Identity"] = gen.PtrOf(Identity_ARMGenerator())
	gens["Properties"] = gen.PtrOf(Namespace_Properties_Spec_ARMGenerator())
	gens["Sku"] = gen.PtrOf(Sku_ARMGenerator())
}

func Test_Sku_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Sku_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSku_ARM, Sku_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSku_ARM runs a test to see if a specific instance of Sku_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSku_ARM(subject Sku_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Sku_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of Sku_ARM instances for property testing - lazily instantiated by Sku_ARMGenerator()
var sku_ARMGenerator gopter.Gen

// Sku_ARMGenerator returns a generator of Sku_ARM instances for property testing.
func Sku_ARMGenerator() gopter.Gen {
	if sku_ARMGenerator != nil {
		return sku_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSku_ARM(generators)
	sku_ARMGenerator = gen.Struct(reflect.TypeOf(Sku_ARM{}), generators)

	return sku_ARMGenerator
}

// AddIndependentPropertyGeneratorsForSku_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSku_ARM(gens map[string]gopter.Gen) {
	gens["Capacity"] = gen.PtrOf(gen.Int())
	gens["Name"] = gen.PtrOf(gen.OneConstOf(Sku_Name_ARM_Basic, Sku_Name_ARM_Premium, Sku_Name_ARM_Standard))
	gens["Tier"] = gen.PtrOf(gen.OneConstOf(Sku_Tier_ARM_Basic, Sku_Tier_ARM_Premium, Sku_Tier_ARM_Standard))
}

func Test_UserAssignedIdentityDetails_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of UserAssignedIdentityDetails_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForUserAssignedIdentityDetails_ARM, UserAssignedIdentityDetails_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForUserAssignedIdentityDetails_ARM runs a test to see if a specific instance of UserAssignedIdentityDetails_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForUserAssignedIdentityDetails_ARM(subject UserAssignedIdentityDetails_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual UserAssignedIdentityDetails_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of UserAssignedIdentityDetails_ARM instances for property testing - lazily instantiated by
// UserAssignedIdentityDetails_ARMGenerator()
var userAssignedIdentityDetails_ARMGenerator gopter.Gen

// UserAssignedIdentityDetails_ARMGenerator returns a generator of UserAssignedIdentityDetails_ARM instances for property testing.
func UserAssignedIdentityDetails_ARMGenerator() gopter.Gen {
	if userAssignedIdentityDetails_ARMGenerator != nil {
		return userAssignedIdentityDetails_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	userAssignedIdentityDetails_ARMGenerator = gen.Struct(reflect.TypeOf(UserAssignedIdentityDetails_ARM{}), generators)

	return userAssignedIdentityDetails_ARMGenerator
}

func Test_UserAssignedIdentityProperties_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of UserAssignedIdentityProperties_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForUserAssignedIdentityProperties_ARM, UserAssignedIdentityProperties_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForUserAssignedIdentityProperties_ARM runs a test to see if a specific instance of UserAssignedIdentityProperties_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForUserAssignedIdentityProperties_ARM(subject UserAssignedIdentityProperties_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual UserAssignedIdentityProperties_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of UserAssignedIdentityProperties_ARM instances for property testing - lazily instantiated by
// UserAssignedIdentityProperties_ARMGenerator()
var userAssignedIdentityProperties_ARMGenerator gopter.Gen

// UserAssignedIdentityProperties_ARMGenerator returns a generator of UserAssignedIdentityProperties_ARM instances for property testing.
func UserAssignedIdentityProperties_ARMGenerator() gopter.Gen {
	if userAssignedIdentityProperties_ARMGenerator != nil {
		return userAssignedIdentityProperties_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForUserAssignedIdentityProperties_ARM(generators)
	userAssignedIdentityProperties_ARMGenerator = gen.Struct(reflect.TypeOf(UserAssignedIdentityProperties_ARM{}), generators)

	return userAssignedIdentityProperties_ARMGenerator
}

// AddIndependentPropertyGeneratorsForUserAssignedIdentityProperties_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForUserAssignedIdentityProperties_ARM(gens map[string]gopter.Gen) {
	gens["UserAssignedIdentity"] = gen.PtrOf(gen.AlphaString())
}
