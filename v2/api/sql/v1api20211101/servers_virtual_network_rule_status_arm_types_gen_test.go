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

func Test_Servers_VirtualNetworkRule_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Servers_VirtualNetworkRule_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServers_VirtualNetworkRule_STATUS_ARM, Servers_VirtualNetworkRule_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServers_VirtualNetworkRule_STATUS_ARM runs a test to see if a specific instance of Servers_VirtualNetworkRule_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForServers_VirtualNetworkRule_STATUS_ARM(subject Servers_VirtualNetworkRule_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Servers_VirtualNetworkRule_STATUS_ARM
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

// Generator of Servers_VirtualNetworkRule_STATUS_ARM instances for property testing - lazily instantiated by
// Servers_VirtualNetworkRule_STATUS_ARMGenerator()
var servers_VirtualNetworkRule_STATUS_ARMGenerator gopter.Gen

// Servers_VirtualNetworkRule_STATUS_ARMGenerator returns a generator of Servers_VirtualNetworkRule_STATUS_ARM instances for property testing.
// We first initialize servers_VirtualNetworkRule_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Servers_VirtualNetworkRule_STATUS_ARMGenerator() gopter.Gen {
	if servers_VirtualNetworkRule_STATUS_ARMGenerator != nil {
		return servers_VirtualNetworkRule_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServers_VirtualNetworkRule_STATUS_ARM(generators)
	servers_VirtualNetworkRule_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Servers_VirtualNetworkRule_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServers_VirtualNetworkRule_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForServers_VirtualNetworkRule_STATUS_ARM(generators)
	servers_VirtualNetworkRule_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Servers_VirtualNetworkRule_STATUS_ARM{}), generators)

	return servers_VirtualNetworkRule_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForServers_VirtualNetworkRule_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServers_VirtualNetworkRule_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForServers_VirtualNetworkRule_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServers_VirtualNetworkRule_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(VirtualNetworkRuleProperties_STATUS_ARMGenerator())
}

func Test_VirtualNetworkRuleProperties_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualNetworkRuleProperties_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualNetworkRuleProperties_STATUS_ARM, VirtualNetworkRuleProperties_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualNetworkRuleProperties_STATUS_ARM runs a test to see if a specific instance of VirtualNetworkRuleProperties_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualNetworkRuleProperties_STATUS_ARM(subject VirtualNetworkRuleProperties_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualNetworkRuleProperties_STATUS_ARM
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

// Generator of VirtualNetworkRuleProperties_STATUS_ARM instances for property testing - lazily instantiated by
// VirtualNetworkRuleProperties_STATUS_ARMGenerator()
var virtualNetworkRuleProperties_STATUS_ARMGenerator gopter.Gen

// VirtualNetworkRuleProperties_STATUS_ARMGenerator returns a generator of VirtualNetworkRuleProperties_STATUS_ARM instances for property testing.
func VirtualNetworkRuleProperties_STATUS_ARMGenerator() gopter.Gen {
	if virtualNetworkRuleProperties_STATUS_ARMGenerator != nil {
		return virtualNetworkRuleProperties_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworkRuleProperties_STATUS_ARM(generators)
	virtualNetworkRuleProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetworkRuleProperties_STATUS_ARM{}), generators)

	return virtualNetworkRuleProperties_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForVirtualNetworkRuleProperties_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualNetworkRuleProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["IgnoreMissingVnetServiceEndpoint"] = gen.PtrOf(gen.Bool())
	gens["State"] = gen.PtrOf(gen.OneConstOf(
		VirtualNetworkRuleProperties_State_STATUS_ARM_Deleting,
		VirtualNetworkRuleProperties_State_STATUS_ARM_Failed,
		VirtualNetworkRuleProperties_State_STATUS_ARM_InProgress,
		VirtualNetworkRuleProperties_State_STATUS_ARM_Initializing,
		VirtualNetworkRuleProperties_State_STATUS_ARM_Ready,
		VirtualNetworkRuleProperties_State_STATUS_ARM_Unknown))
	gens["VirtualNetworkSubnetId"] = gen.PtrOf(gen.AlphaString())
}
