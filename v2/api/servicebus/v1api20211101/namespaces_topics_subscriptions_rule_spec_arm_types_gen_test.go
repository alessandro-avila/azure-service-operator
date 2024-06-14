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

func Test_Action_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Action_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAction_ARM, Action_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAction_ARM runs a test to see if a specific instance of Action_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForAction_ARM(subject Action_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Action_ARM
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

// Generator of Action_ARM instances for property testing - lazily instantiated by Action_ARMGenerator()
var action_ARMGenerator gopter.Gen

// Action_ARMGenerator returns a generator of Action_ARM instances for property testing.
func Action_ARMGenerator() gopter.Gen {
	if action_ARMGenerator != nil {
		return action_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAction_ARM(generators)
	action_ARMGenerator = gen.Struct(reflect.TypeOf(Action_ARM{}), generators)

	return action_ARMGenerator
}

// AddIndependentPropertyGeneratorsForAction_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAction_ARM(gens map[string]gopter.Gen) {
	gens["CompatibilityLevel"] = gen.PtrOf(gen.Int())
	gens["RequiresPreprocessing"] = gen.PtrOf(gen.Bool())
	gens["SqlExpression"] = gen.PtrOf(gen.AlphaString())
}

func Test_CorrelationFilter_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of CorrelationFilter_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCorrelationFilter_ARM, CorrelationFilter_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCorrelationFilter_ARM runs a test to see if a specific instance of CorrelationFilter_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForCorrelationFilter_ARM(subject CorrelationFilter_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual CorrelationFilter_ARM
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

// Generator of CorrelationFilter_ARM instances for property testing - lazily instantiated by
// CorrelationFilter_ARMGenerator()
var correlationFilter_ARMGenerator gopter.Gen

// CorrelationFilter_ARMGenerator returns a generator of CorrelationFilter_ARM instances for property testing.
func CorrelationFilter_ARMGenerator() gopter.Gen {
	if correlationFilter_ARMGenerator != nil {
		return correlationFilter_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCorrelationFilter_ARM(generators)
	correlationFilter_ARMGenerator = gen.Struct(reflect.TypeOf(CorrelationFilter_ARM{}), generators)

	return correlationFilter_ARMGenerator
}

// AddIndependentPropertyGeneratorsForCorrelationFilter_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForCorrelationFilter_ARM(gens map[string]gopter.Gen) {
	gens["ContentType"] = gen.PtrOf(gen.AlphaString())
	gens["CorrelationId"] = gen.PtrOf(gen.AlphaString())
	gens["Label"] = gen.PtrOf(gen.AlphaString())
	gens["MessageId"] = gen.PtrOf(gen.AlphaString())
	gens["Properties"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["ReplyTo"] = gen.PtrOf(gen.AlphaString())
	gens["ReplyToSessionId"] = gen.PtrOf(gen.AlphaString())
	gens["RequiresPreprocessing"] = gen.PtrOf(gen.Bool())
	gens["SessionId"] = gen.PtrOf(gen.AlphaString())
	gens["To"] = gen.PtrOf(gen.AlphaString())
}

func Test_Namespaces_Topics_Subscriptions_Rule_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Namespaces_Topics_Subscriptions_Rule_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespaces_Topics_Subscriptions_Rule_Spec_ARM, Namespaces_Topics_Subscriptions_Rule_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespaces_Topics_Subscriptions_Rule_Spec_ARM runs a test to see if a specific instance of Namespaces_Topics_Subscriptions_Rule_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespaces_Topics_Subscriptions_Rule_Spec_ARM(subject Namespaces_Topics_Subscriptions_Rule_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Namespaces_Topics_Subscriptions_Rule_Spec_ARM
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

// Generator of Namespaces_Topics_Subscriptions_Rule_Spec_ARM instances for property testing - lazily instantiated by
// Namespaces_Topics_Subscriptions_Rule_Spec_ARMGenerator()
var namespaces_Topics_Subscriptions_Rule_Spec_ARMGenerator gopter.Gen

// Namespaces_Topics_Subscriptions_Rule_Spec_ARMGenerator returns a generator of Namespaces_Topics_Subscriptions_Rule_Spec_ARM instances for property testing.
// We first initialize namespaces_Topics_Subscriptions_Rule_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Namespaces_Topics_Subscriptions_Rule_Spec_ARMGenerator() gopter.Gen {
	if namespaces_Topics_Subscriptions_Rule_Spec_ARMGenerator != nil {
		return namespaces_Topics_Subscriptions_Rule_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespaces_Topics_Subscriptions_Rule_Spec_ARM(generators)
	namespaces_Topics_Subscriptions_Rule_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(Namespaces_Topics_Subscriptions_Rule_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespaces_Topics_Subscriptions_Rule_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForNamespaces_Topics_Subscriptions_Rule_Spec_ARM(generators)
	namespaces_Topics_Subscriptions_Rule_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(Namespaces_Topics_Subscriptions_Rule_Spec_ARM{}), generators)

	return namespaces_Topics_Subscriptions_Rule_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForNamespaces_Topics_Subscriptions_Rule_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespaces_Topics_Subscriptions_Rule_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForNamespaces_Topics_Subscriptions_Rule_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespaces_Topics_Subscriptions_Rule_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(Ruleproperties_ARMGenerator())
}

func Test_Ruleproperties_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Ruleproperties_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRuleproperties_ARM, Ruleproperties_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRuleproperties_ARM runs a test to see if a specific instance of Ruleproperties_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRuleproperties_ARM(subject Ruleproperties_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Ruleproperties_ARM
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

// Generator of Ruleproperties_ARM instances for property testing - lazily instantiated by Ruleproperties_ARMGenerator()
var ruleproperties_ARMGenerator gopter.Gen

// Ruleproperties_ARMGenerator returns a generator of Ruleproperties_ARM instances for property testing.
// We first initialize ruleproperties_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Ruleproperties_ARMGenerator() gopter.Gen {
	if ruleproperties_ARMGenerator != nil {
		return ruleproperties_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRuleproperties_ARM(generators)
	ruleproperties_ARMGenerator = gen.Struct(reflect.TypeOf(Ruleproperties_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRuleproperties_ARM(generators)
	AddRelatedPropertyGeneratorsForRuleproperties_ARM(generators)
	ruleproperties_ARMGenerator = gen.Struct(reflect.TypeOf(Ruleproperties_ARM{}), generators)

	return ruleproperties_ARMGenerator
}

// AddIndependentPropertyGeneratorsForRuleproperties_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRuleproperties_ARM(gens map[string]gopter.Gen) {
	gens["FilterType"] = gen.PtrOf(gen.OneConstOf(FilterType_ARM_CorrelationFilter, FilterType_ARM_SqlFilter))
}

// AddRelatedPropertyGeneratorsForRuleproperties_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRuleproperties_ARM(gens map[string]gopter.Gen) {
	gens["Action"] = gen.PtrOf(Action_ARMGenerator())
	gens["CorrelationFilter"] = gen.PtrOf(CorrelationFilter_ARMGenerator())
	gens["SqlFilter"] = gen.PtrOf(SqlFilter_ARMGenerator())
}

func Test_SqlFilter_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlFilter_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlFilter_ARM, SqlFilter_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlFilter_ARM runs a test to see if a specific instance of SqlFilter_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlFilter_ARM(subject SqlFilter_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlFilter_ARM
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

// Generator of SqlFilter_ARM instances for property testing - lazily instantiated by SqlFilter_ARMGenerator()
var sqlFilter_ARMGenerator gopter.Gen

// SqlFilter_ARMGenerator returns a generator of SqlFilter_ARM instances for property testing.
func SqlFilter_ARMGenerator() gopter.Gen {
	if sqlFilter_ARMGenerator != nil {
		return sqlFilter_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlFilter_ARM(generators)
	sqlFilter_ARMGenerator = gen.Struct(reflect.TypeOf(SqlFilter_ARM{}), generators)

	return sqlFilter_ARMGenerator
}

// AddIndependentPropertyGeneratorsForSqlFilter_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlFilter_ARM(gens map[string]gopter.Gen) {
	gens["CompatibilityLevel"] = gen.PtrOf(gen.Int())
	gens["RequiresPreprocessing"] = gen.PtrOf(gen.Bool())
	gens["SqlExpression"] = gen.PtrOf(gen.AlphaString())
}
