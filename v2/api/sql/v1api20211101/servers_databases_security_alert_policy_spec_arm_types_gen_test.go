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

func Test_DatabaseSecurityAlertPoliciesSecurityAlertsPolicyProperties_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseSecurityAlertPoliciesSecurityAlertsPolicyProperties_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseSecurityAlertPoliciesSecurityAlertsPolicyProperties_ARM, DatabaseSecurityAlertPoliciesSecurityAlertsPolicyProperties_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseSecurityAlertPoliciesSecurityAlertsPolicyProperties_ARM runs a test to see if a specific instance of DatabaseSecurityAlertPoliciesSecurityAlertsPolicyProperties_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseSecurityAlertPoliciesSecurityAlertsPolicyProperties_ARM(subject DatabaseSecurityAlertPoliciesSecurityAlertsPolicyProperties_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseSecurityAlertPoliciesSecurityAlertsPolicyProperties_ARM
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

// Generator of DatabaseSecurityAlertPoliciesSecurityAlertsPolicyProperties_ARM instances for property testing - lazily
// instantiated by DatabaseSecurityAlertPoliciesSecurityAlertsPolicyProperties_ARMGenerator()
var databaseSecurityAlertPoliciesSecurityAlertsPolicyProperties_ARMGenerator gopter.Gen

// DatabaseSecurityAlertPoliciesSecurityAlertsPolicyProperties_ARMGenerator returns a generator of DatabaseSecurityAlertPoliciesSecurityAlertsPolicyProperties_ARM instances for property testing.
func DatabaseSecurityAlertPoliciesSecurityAlertsPolicyProperties_ARMGenerator() gopter.Gen {
	if databaseSecurityAlertPoliciesSecurityAlertsPolicyProperties_ARMGenerator != nil {
		return databaseSecurityAlertPoliciesSecurityAlertsPolicyProperties_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseSecurityAlertPoliciesSecurityAlertsPolicyProperties_ARM(generators)
	databaseSecurityAlertPoliciesSecurityAlertsPolicyProperties_ARMGenerator = gen.Struct(reflect.TypeOf(DatabaseSecurityAlertPoliciesSecurityAlertsPolicyProperties_ARM{}), generators)

	return databaseSecurityAlertPoliciesSecurityAlertsPolicyProperties_ARMGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseSecurityAlertPoliciesSecurityAlertsPolicyProperties_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseSecurityAlertPoliciesSecurityAlertsPolicyProperties_ARM(gens map[string]gopter.Gen) {
	gens["DisabledAlerts"] = gen.SliceOf(gen.AlphaString())
	gens["EmailAccountAdmins"] = gen.PtrOf(gen.Bool())
	gens["EmailAddresses"] = gen.SliceOf(gen.AlphaString())
	gens["RetentionDays"] = gen.PtrOf(gen.Int())
	gens["State"] = gen.PtrOf(gen.OneConstOf(DatabaseSecurityAlertPoliciesSecurityAlertsPolicyProperties_State_ARM_Disabled, DatabaseSecurityAlertPoliciesSecurityAlertsPolicyProperties_State_ARM_Enabled))
	gens["StorageAccountAccessKey"] = gen.PtrOf(gen.AlphaString())
	gens["StorageEndpoint"] = gen.PtrOf(gen.AlphaString())
}

func Test_Servers_Databases_SecurityAlertPolicy_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Servers_Databases_SecurityAlertPolicy_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServers_Databases_SecurityAlertPolicy_Spec_ARM, Servers_Databases_SecurityAlertPolicy_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServers_Databases_SecurityAlertPolicy_Spec_ARM runs a test to see if a specific instance of Servers_Databases_SecurityAlertPolicy_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForServers_Databases_SecurityAlertPolicy_Spec_ARM(subject Servers_Databases_SecurityAlertPolicy_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Servers_Databases_SecurityAlertPolicy_Spec_ARM
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

// Generator of Servers_Databases_SecurityAlertPolicy_Spec_ARM instances for property testing - lazily instantiated by
// Servers_Databases_SecurityAlertPolicy_Spec_ARMGenerator()
var servers_Databases_SecurityAlertPolicy_Spec_ARMGenerator gopter.Gen

// Servers_Databases_SecurityAlertPolicy_Spec_ARMGenerator returns a generator of Servers_Databases_SecurityAlertPolicy_Spec_ARM instances for property testing.
// We first initialize servers_Databases_SecurityAlertPolicy_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Servers_Databases_SecurityAlertPolicy_Spec_ARMGenerator() gopter.Gen {
	if servers_Databases_SecurityAlertPolicy_Spec_ARMGenerator != nil {
		return servers_Databases_SecurityAlertPolicy_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServers_Databases_SecurityAlertPolicy_Spec_ARM(generators)
	servers_Databases_SecurityAlertPolicy_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(Servers_Databases_SecurityAlertPolicy_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServers_Databases_SecurityAlertPolicy_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForServers_Databases_SecurityAlertPolicy_Spec_ARM(generators)
	servers_Databases_SecurityAlertPolicy_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(Servers_Databases_SecurityAlertPolicy_Spec_ARM{}), generators)

	return servers_Databases_SecurityAlertPolicy_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForServers_Databases_SecurityAlertPolicy_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServers_Databases_SecurityAlertPolicy_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForServers_Databases_SecurityAlertPolicy_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServers_Databases_SecurityAlertPolicy_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(DatabaseSecurityAlertPoliciesSecurityAlertsPolicyProperties_ARMGenerator())
}
