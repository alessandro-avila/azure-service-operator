// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20230301

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

func Test_PrometheusRuleGroupAction_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrometheusRuleGroupAction_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrometheusRuleGroupAction_STATUS_ARM, PrometheusRuleGroupAction_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrometheusRuleGroupAction_STATUS_ARM runs a test to see if a specific instance of PrometheusRuleGroupAction_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPrometheusRuleGroupAction_STATUS_ARM(subject PrometheusRuleGroupAction_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrometheusRuleGroupAction_STATUS_ARM
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

// Generator of PrometheusRuleGroupAction_STATUS_ARM instances for property testing - lazily instantiated by
// PrometheusRuleGroupAction_STATUS_ARMGenerator()
var prometheusRuleGroupAction_STATUS_ARMGenerator gopter.Gen

// PrometheusRuleGroupAction_STATUS_ARMGenerator returns a generator of PrometheusRuleGroupAction_STATUS_ARM instances for property testing.
func PrometheusRuleGroupAction_STATUS_ARMGenerator() gopter.Gen {
	if prometheusRuleGroupAction_STATUS_ARMGenerator != nil {
		return prometheusRuleGroupAction_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrometheusRuleGroupAction_STATUS_ARM(generators)
	prometheusRuleGroupAction_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(PrometheusRuleGroupAction_STATUS_ARM{}), generators)

	return prometheusRuleGroupAction_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForPrometheusRuleGroupAction_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrometheusRuleGroupAction_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["ActionGroupId"] = gen.PtrOf(gen.AlphaString())
	gens["ActionProperties"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
}

func Test_PrometheusRuleGroupProperties_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrometheusRuleGroupProperties_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrometheusRuleGroupProperties_STATUS_ARM, PrometheusRuleGroupProperties_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrometheusRuleGroupProperties_STATUS_ARM runs a test to see if a specific instance of PrometheusRuleGroupProperties_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPrometheusRuleGroupProperties_STATUS_ARM(subject PrometheusRuleGroupProperties_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrometheusRuleGroupProperties_STATUS_ARM
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

// Generator of PrometheusRuleGroupProperties_STATUS_ARM instances for property testing - lazily instantiated by
// PrometheusRuleGroupProperties_STATUS_ARMGenerator()
var prometheusRuleGroupProperties_STATUS_ARMGenerator gopter.Gen

// PrometheusRuleGroupProperties_STATUS_ARMGenerator returns a generator of PrometheusRuleGroupProperties_STATUS_ARM instances for property testing.
// We first initialize prometheusRuleGroupProperties_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PrometheusRuleGroupProperties_STATUS_ARMGenerator() gopter.Gen {
	if prometheusRuleGroupProperties_STATUS_ARMGenerator != nil {
		return prometheusRuleGroupProperties_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrometheusRuleGroupProperties_STATUS_ARM(generators)
	prometheusRuleGroupProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(PrometheusRuleGroupProperties_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrometheusRuleGroupProperties_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForPrometheusRuleGroupProperties_STATUS_ARM(generators)
	prometheusRuleGroupProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(PrometheusRuleGroupProperties_STATUS_ARM{}), generators)

	return prometheusRuleGroupProperties_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForPrometheusRuleGroupProperties_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrometheusRuleGroupProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["ClusterName"] = gen.PtrOf(gen.AlphaString())
	gens["Description"] = gen.PtrOf(gen.AlphaString())
	gens["Enabled"] = gen.PtrOf(gen.Bool())
	gens["Interval"] = gen.PtrOf(gen.AlphaString())
	gens["Scopes"] = gen.SliceOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForPrometheusRuleGroupProperties_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrometheusRuleGroupProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Rules"] = gen.SliceOf(PrometheusRule_STATUS_ARMGenerator())
}

func Test_PrometheusRuleGroup_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrometheusRuleGroup_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrometheusRuleGroup_STATUS_ARM, PrometheusRuleGroup_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrometheusRuleGroup_STATUS_ARM runs a test to see if a specific instance of PrometheusRuleGroup_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPrometheusRuleGroup_STATUS_ARM(subject PrometheusRuleGroup_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrometheusRuleGroup_STATUS_ARM
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

// Generator of PrometheusRuleGroup_STATUS_ARM instances for property testing - lazily instantiated by
// PrometheusRuleGroup_STATUS_ARMGenerator()
var prometheusRuleGroup_STATUS_ARMGenerator gopter.Gen

// PrometheusRuleGroup_STATUS_ARMGenerator returns a generator of PrometheusRuleGroup_STATUS_ARM instances for property testing.
// We first initialize prometheusRuleGroup_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PrometheusRuleGroup_STATUS_ARMGenerator() gopter.Gen {
	if prometheusRuleGroup_STATUS_ARMGenerator != nil {
		return prometheusRuleGroup_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrometheusRuleGroup_STATUS_ARM(generators)
	prometheusRuleGroup_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(PrometheusRuleGroup_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrometheusRuleGroup_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForPrometheusRuleGroup_STATUS_ARM(generators)
	prometheusRuleGroup_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(PrometheusRuleGroup_STATUS_ARM{}), generators)

	return prometheusRuleGroup_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForPrometheusRuleGroup_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrometheusRuleGroup_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForPrometheusRuleGroup_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrometheusRuleGroup_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(PrometheusRuleGroupProperties_STATUS_ARMGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_STATUS_ARMGenerator())
}

func Test_PrometheusRuleResolveConfiguration_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrometheusRuleResolveConfiguration_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrometheusRuleResolveConfiguration_STATUS_ARM, PrometheusRuleResolveConfiguration_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrometheusRuleResolveConfiguration_STATUS_ARM runs a test to see if a specific instance of PrometheusRuleResolveConfiguration_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPrometheusRuleResolveConfiguration_STATUS_ARM(subject PrometheusRuleResolveConfiguration_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrometheusRuleResolveConfiguration_STATUS_ARM
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

// Generator of PrometheusRuleResolveConfiguration_STATUS_ARM instances for property testing - lazily instantiated by
// PrometheusRuleResolveConfiguration_STATUS_ARMGenerator()
var prometheusRuleResolveConfiguration_STATUS_ARMGenerator gopter.Gen

// PrometheusRuleResolveConfiguration_STATUS_ARMGenerator returns a generator of PrometheusRuleResolveConfiguration_STATUS_ARM instances for property testing.
func PrometheusRuleResolveConfiguration_STATUS_ARMGenerator() gopter.Gen {
	if prometheusRuleResolveConfiguration_STATUS_ARMGenerator != nil {
		return prometheusRuleResolveConfiguration_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrometheusRuleResolveConfiguration_STATUS_ARM(generators)
	prometheusRuleResolveConfiguration_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(PrometheusRuleResolveConfiguration_STATUS_ARM{}), generators)

	return prometheusRuleResolveConfiguration_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForPrometheusRuleResolveConfiguration_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrometheusRuleResolveConfiguration_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["AutoResolved"] = gen.PtrOf(gen.Bool())
	gens["TimeToResolve"] = gen.PtrOf(gen.AlphaString())
}

func Test_PrometheusRule_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrometheusRule_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrometheusRule_STATUS_ARM, PrometheusRule_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrometheusRule_STATUS_ARM runs a test to see if a specific instance of PrometheusRule_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPrometheusRule_STATUS_ARM(subject PrometheusRule_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrometheusRule_STATUS_ARM
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

// Generator of PrometheusRule_STATUS_ARM instances for property testing - lazily instantiated by
// PrometheusRule_STATUS_ARMGenerator()
var prometheusRule_STATUS_ARMGenerator gopter.Gen

// PrometheusRule_STATUS_ARMGenerator returns a generator of PrometheusRule_STATUS_ARM instances for property testing.
// We first initialize prometheusRule_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PrometheusRule_STATUS_ARMGenerator() gopter.Gen {
	if prometheusRule_STATUS_ARMGenerator != nil {
		return prometheusRule_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrometheusRule_STATUS_ARM(generators)
	prometheusRule_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(PrometheusRule_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrometheusRule_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForPrometheusRule_STATUS_ARM(generators)
	prometheusRule_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(PrometheusRule_STATUS_ARM{}), generators)

	return prometheusRule_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForPrometheusRule_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrometheusRule_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Alert"] = gen.PtrOf(gen.AlphaString())
	gens["Annotations"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Enabled"] = gen.PtrOf(gen.Bool())
	gens["Expression"] = gen.PtrOf(gen.AlphaString())
	gens["For"] = gen.PtrOf(gen.AlphaString())
	gens["Labels"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Record"] = gen.PtrOf(gen.AlphaString())
	gens["Severity"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForPrometheusRule_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrometheusRule_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Actions"] = gen.SliceOf(PrometheusRuleGroupAction_STATUS_ARMGenerator())
	gens["ResolveConfiguration"] = gen.PtrOf(PrometheusRuleResolveConfiguration_STATUS_ARMGenerator())
}

func Test_SystemData_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SystemData_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSystemData_STATUS_ARM, SystemData_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSystemData_STATUS_ARM runs a test to see if a specific instance of SystemData_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSystemData_STATUS_ARM(subject SystemData_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SystemData_STATUS_ARM
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

// Generator of SystemData_STATUS_ARM instances for property testing - lazily instantiated by
// SystemData_STATUS_ARMGenerator()
var systemData_STATUS_ARMGenerator gopter.Gen

// SystemData_STATUS_ARMGenerator returns a generator of SystemData_STATUS_ARM instances for property testing.
func SystemData_STATUS_ARMGenerator() gopter.Gen {
	if systemData_STATUS_ARMGenerator != nil {
		return systemData_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSystemData_STATUS_ARM(generators)
	systemData_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(SystemData_STATUS_ARM{}), generators)

	return systemData_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForSystemData_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSystemData_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["CreatedAt"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedBy"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedByType"] = gen.PtrOf(gen.OneConstOf(
		SystemData_CreatedByType_STATUS_ARM_Application,
		SystemData_CreatedByType_STATUS_ARM_Key,
		SystemData_CreatedByType_STATUS_ARM_ManagedIdentity,
		SystemData_CreatedByType_STATUS_ARM_User))
	gens["LastModifiedAt"] = gen.PtrOf(gen.AlphaString())
	gens["LastModifiedBy"] = gen.PtrOf(gen.AlphaString())
	gens["LastModifiedByType"] = gen.PtrOf(gen.OneConstOf(
		SystemData_LastModifiedByType_STATUS_ARM_Application,
		SystemData_LastModifiedByType_STATUS_ARM_Key,
		SystemData_LastModifiedByType_STATUS_ARM_ManagedIdentity,
		SystemData_LastModifiedByType_STATUS_ARM_User))
}
