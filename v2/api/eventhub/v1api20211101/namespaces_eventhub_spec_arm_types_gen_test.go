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

func Test_CaptureDescription_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of CaptureDescription_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCaptureDescription_ARM, CaptureDescription_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCaptureDescription_ARM runs a test to see if a specific instance of CaptureDescription_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForCaptureDescription_ARM(subject CaptureDescription_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual CaptureDescription_ARM
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

// Generator of CaptureDescription_ARM instances for property testing - lazily instantiated by
// CaptureDescription_ARMGenerator()
var captureDescription_ARMGenerator gopter.Gen

// CaptureDescription_ARMGenerator returns a generator of CaptureDescription_ARM instances for property testing.
// We first initialize captureDescription_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func CaptureDescription_ARMGenerator() gopter.Gen {
	if captureDescription_ARMGenerator != nil {
		return captureDescription_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCaptureDescription_ARM(generators)
	captureDescription_ARMGenerator = gen.Struct(reflect.TypeOf(CaptureDescription_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCaptureDescription_ARM(generators)
	AddRelatedPropertyGeneratorsForCaptureDescription_ARM(generators)
	captureDescription_ARMGenerator = gen.Struct(reflect.TypeOf(CaptureDescription_ARM{}), generators)

	return captureDescription_ARMGenerator
}

// AddIndependentPropertyGeneratorsForCaptureDescription_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForCaptureDescription_ARM(gens map[string]gopter.Gen) {
	gens["Enabled"] = gen.PtrOf(gen.Bool())
	gens["Encoding"] = gen.PtrOf(gen.OneConstOf(CaptureDescription_Encoding_ARM_Avro, CaptureDescription_Encoding_ARM_AvroDeflate))
	gens["IntervalInSeconds"] = gen.PtrOf(gen.Int())
	gens["SizeLimitInBytes"] = gen.PtrOf(gen.Int())
	gens["SkipEmptyArchives"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForCaptureDescription_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForCaptureDescription_ARM(gens map[string]gopter.Gen) {
	gens["Destination"] = gen.PtrOf(Destination_ARMGenerator())
}

func Test_Destination_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Destination_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDestination_ARM, Destination_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDestination_ARM runs a test to see if a specific instance of Destination_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDestination_ARM(subject Destination_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Destination_ARM
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

// Generator of Destination_ARM instances for property testing - lazily instantiated by Destination_ARMGenerator()
var destination_ARMGenerator gopter.Gen

// Destination_ARMGenerator returns a generator of Destination_ARM instances for property testing.
// We first initialize destination_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Destination_ARMGenerator() gopter.Gen {
	if destination_ARMGenerator != nil {
		return destination_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDestination_ARM(generators)
	destination_ARMGenerator = gen.Struct(reflect.TypeOf(Destination_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDestination_ARM(generators)
	AddRelatedPropertyGeneratorsForDestination_ARM(generators)
	destination_ARMGenerator = gen.Struct(reflect.TypeOf(Destination_ARM{}), generators)

	return destination_ARMGenerator
}

// AddIndependentPropertyGeneratorsForDestination_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDestination_ARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDestination_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDestination_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(Destination_Properties_ARMGenerator())
}

func Test_Destination_Properties_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Destination_Properties_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDestination_Properties_ARM, Destination_Properties_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDestination_Properties_ARM runs a test to see if a specific instance of Destination_Properties_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDestination_Properties_ARM(subject Destination_Properties_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Destination_Properties_ARM
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

// Generator of Destination_Properties_ARM instances for property testing - lazily instantiated by
// Destination_Properties_ARMGenerator()
var destination_Properties_ARMGenerator gopter.Gen

// Destination_Properties_ARMGenerator returns a generator of Destination_Properties_ARM instances for property testing.
func Destination_Properties_ARMGenerator() gopter.Gen {
	if destination_Properties_ARMGenerator != nil {
		return destination_Properties_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDestination_Properties_ARM(generators)
	destination_Properties_ARMGenerator = gen.Struct(reflect.TypeOf(Destination_Properties_ARM{}), generators)

	return destination_Properties_ARMGenerator
}

// AddIndependentPropertyGeneratorsForDestination_Properties_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDestination_Properties_ARM(gens map[string]gopter.Gen) {
	gens["ArchiveNameFormat"] = gen.PtrOf(gen.AlphaString())
	gens["BlobContainer"] = gen.PtrOf(gen.AlphaString())
	gens["DataLakeAccountName"] = gen.PtrOf(gen.AlphaString())
	gens["DataLakeFolderPath"] = gen.PtrOf(gen.AlphaString())
	gens["DataLakeSubscriptionId"] = gen.PtrOf(gen.AlphaString())
	gens["StorageAccountResourceId"] = gen.PtrOf(gen.AlphaString())
}

func Test_Namespaces_Eventhub_Properties_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Namespaces_Eventhub_Properties_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespaces_Eventhub_Properties_Spec_ARM, Namespaces_Eventhub_Properties_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespaces_Eventhub_Properties_Spec_ARM runs a test to see if a specific instance of Namespaces_Eventhub_Properties_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespaces_Eventhub_Properties_Spec_ARM(subject Namespaces_Eventhub_Properties_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Namespaces_Eventhub_Properties_Spec_ARM
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

// Generator of Namespaces_Eventhub_Properties_Spec_ARM instances for property testing - lazily instantiated by
// Namespaces_Eventhub_Properties_Spec_ARMGenerator()
var namespaces_Eventhub_Properties_Spec_ARMGenerator gopter.Gen

// Namespaces_Eventhub_Properties_Spec_ARMGenerator returns a generator of Namespaces_Eventhub_Properties_Spec_ARM instances for property testing.
// We first initialize namespaces_Eventhub_Properties_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Namespaces_Eventhub_Properties_Spec_ARMGenerator() gopter.Gen {
	if namespaces_Eventhub_Properties_Spec_ARMGenerator != nil {
		return namespaces_Eventhub_Properties_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespaces_Eventhub_Properties_Spec_ARM(generators)
	namespaces_Eventhub_Properties_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(Namespaces_Eventhub_Properties_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespaces_Eventhub_Properties_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForNamespaces_Eventhub_Properties_Spec_ARM(generators)
	namespaces_Eventhub_Properties_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(Namespaces_Eventhub_Properties_Spec_ARM{}), generators)

	return namespaces_Eventhub_Properties_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForNamespaces_Eventhub_Properties_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespaces_Eventhub_Properties_Spec_ARM(gens map[string]gopter.Gen) {
	gens["MessageRetentionInDays"] = gen.PtrOf(gen.Int())
	gens["PartitionCount"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForNamespaces_Eventhub_Properties_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespaces_Eventhub_Properties_Spec_ARM(gens map[string]gopter.Gen) {
	gens["CaptureDescription"] = gen.PtrOf(CaptureDescription_ARMGenerator())
}

func Test_Namespaces_Eventhub_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Namespaces_Eventhub_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespaces_Eventhub_Spec_ARM, Namespaces_Eventhub_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespaces_Eventhub_Spec_ARM runs a test to see if a specific instance of Namespaces_Eventhub_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespaces_Eventhub_Spec_ARM(subject Namespaces_Eventhub_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Namespaces_Eventhub_Spec_ARM
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

// Generator of Namespaces_Eventhub_Spec_ARM instances for property testing - lazily instantiated by
// Namespaces_Eventhub_Spec_ARMGenerator()
var namespaces_Eventhub_Spec_ARMGenerator gopter.Gen

// Namespaces_Eventhub_Spec_ARMGenerator returns a generator of Namespaces_Eventhub_Spec_ARM instances for property testing.
// We first initialize namespaces_Eventhub_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Namespaces_Eventhub_Spec_ARMGenerator() gopter.Gen {
	if namespaces_Eventhub_Spec_ARMGenerator != nil {
		return namespaces_Eventhub_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespaces_Eventhub_Spec_ARM(generators)
	namespaces_Eventhub_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(Namespaces_Eventhub_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespaces_Eventhub_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForNamespaces_Eventhub_Spec_ARM(generators)
	namespaces_Eventhub_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(Namespaces_Eventhub_Spec_ARM{}), generators)

	return namespaces_Eventhub_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForNamespaces_Eventhub_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespaces_Eventhub_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForNamespaces_Eventhub_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespaces_Eventhub_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(Namespaces_Eventhub_Properties_Spec_ARMGenerator())
}
