// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20210601

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

func Test_Backup_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Backup_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForBackup_ARM, Backup_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForBackup_ARM runs a test to see if a specific instance of Backup_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForBackup_ARM(subject Backup_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Backup_ARM
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

// Generator of Backup_ARM instances for property testing - lazily instantiated by Backup_ARMGenerator()
var backup_ARMGenerator gopter.Gen

// Backup_ARMGenerator returns a generator of Backup_ARM instances for property testing.
func Backup_ARMGenerator() gopter.Gen {
	if backup_ARMGenerator != nil {
		return backup_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBackup_ARM(generators)
	backup_ARMGenerator = gen.Struct(reflect.TypeOf(Backup_ARM{}), generators)

	return backup_ARMGenerator
}

// AddIndependentPropertyGeneratorsForBackup_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForBackup_ARM(gens map[string]gopter.Gen) {
	gens["BackupRetentionDays"] = gen.PtrOf(gen.Int())
	gens["GeoRedundantBackup"] = gen.PtrOf(gen.OneConstOf(Backup_GeoRedundantBackup_ARM_Disabled, Backup_GeoRedundantBackup_ARM_Enabled))
}

func Test_FlexibleServer_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FlexibleServer_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFlexibleServer_Spec_ARM, FlexibleServer_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFlexibleServer_Spec_ARM runs a test to see if a specific instance of FlexibleServer_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForFlexibleServer_Spec_ARM(subject FlexibleServer_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FlexibleServer_Spec_ARM
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

// Generator of FlexibleServer_Spec_ARM instances for property testing - lazily instantiated by
// FlexibleServer_Spec_ARMGenerator()
var flexibleServer_Spec_ARMGenerator gopter.Gen

// FlexibleServer_Spec_ARMGenerator returns a generator of FlexibleServer_Spec_ARM instances for property testing.
// We first initialize flexibleServer_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func FlexibleServer_Spec_ARMGenerator() gopter.Gen {
	if flexibleServer_Spec_ARMGenerator != nil {
		return flexibleServer_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlexibleServer_Spec_ARM(generators)
	flexibleServer_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(FlexibleServer_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlexibleServer_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForFlexibleServer_Spec_ARM(generators)
	flexibleServer_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(FlexibleServer_Spec_ARM{}), generators)

	return flexibleServer_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForFlexibleServer_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFlexibleServer_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForFlexibleServer_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFlexibleServer_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ServerProperties_ARMGenerator())
	gens["Sku"] = gen.PtrOf(Sku_ARMGenerator())
}

func Test_HighAvailability_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of HighAvailability_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForHighAvailability_ARM, HighAvailability_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForHighAvailability_ARM runs a test to see if a specific instance of HighAvailability_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForHighAvailability_ARM(subject HighAvailability_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual HighAvailability_ARM
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

// Generator of HighAvailability_ARM instances for property testing - lazily instantiated by
// HighAvailability_ARMGenerator()
var highAvailability_ARMGenerator gopter.Gen

// HighAvailability_ARMGenerator returns a generator of HighAvailability_ARM instances for property testing.
func HighAvailability_ARMGenerator() gopter.Gen {
	if highAvailability_ARMGenerator != nil {
		return highAvailability_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForHighAvailability_ARM(generators)
	highAvailability_ARMGenerator = gen.Struct(reflect.TypeOf(HighAvailability_ARM{}), generators)

	return highAvailability_ARMGenerator
}

// AddIndependentPropertyGeneratorsForHighAvailability_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForHighAvailability_ARM(gens map[string]gopter.Gen) {
	gens["Mode"] = gen.PtrOf(gen.OneConstOf(HighAvailability_Mode_ARM_Disabled, HighAvailability_Mode_ARM_ZoneRedundant))
	gens["StandbyAvailabilityZone"] = gen.PtrOf(gen.AlphaString())
}

func Test_MaintenanceWindow_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MaintenanceWindow_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMaintenanceWindow_ARM, MaintenanceWindow_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMaintenanceWindow_ARM runs a test to see if a specific instance of MaintenanceWindow_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForMaintenanceWindow_ARM(subject MaintenanceWindow_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MaintenanceWindow_ARM
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

// Generator of MaintenanceWindow_ARM instances for property testing - lazily instantiated by
// MaintenanceWindow_ARMGenerator()
var maintenanceWindow_ARMGenerator gopter.Gen

// MaintenanceWindow_ARMGenerator returns a generator of MaintenanceWindow_ARM instances for property testing.
func MaintenanceWindow_ARMGenerator() gopter.Gen {
	if maintenanceWindow_ARMGenerator != nil {
		return maintenanceWindow_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMaintenanceWindow_ARM(generators)
	maintenanceWindow_ARMGenerator = gen.Struct(reflect.TypeOf(MaintenanceWindow_ARM{}), generators)

	return maintenanceWindow_ARMGenerator
}

// AddIndependentPropertyGeneratorsForMaintenanceWindow_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMaintenanceWindow_ARM(gens map[string]gopter.Gen) {
	gens["CustomWindow"] = gen.PtrOf(gen.AlphaString())
	gens["DayOfWeek"] = gen.PtrOf(gen.Int())
	gens["StartHour"] = gen.PtrOf(gen.Int())
	gens["StartMinute"] = gen.PtrOf(gen.Int())
}

func Test_Network_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Network_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNetwork_ARM, Network_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNetwork_ARM runs a test to see if a specific instance of Network_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNetwork_ARM(subject Network_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Network_ARM
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

// Generator of Network_ARM instances for property testing - lazily instantiated by Network_ARMGenerator()
var network_ARMGenerator gopter.Gen

// Network_ARMGenerator returns a generator of Network_ARM instances for property testing.
func Network_ARMGenerator() gopter.Gen {
	if network_ARMGenerator != nil {
		return network_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetwork_ARM(generators)
	network_ARMGenerator = gen.Struct(reflect.TypeOf(Network_ARM{}), generators)

	return network_ARMGenerator
}

// AddIndependentPropertyGeneratorsForNetwork_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNetwork_ARM(gens map[string]gopter.Gen) {
	gens["DelegatedSubnetResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["PrivateDnsZoneArmResourceId"] = gen.PtrOf(gen.AlphaString())
}

func Test_ServerProperties_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServerProperties_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServerProperties_ARM, ServerProperties_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServerProperties_ARM runs a test to see if a specific instance of ServerProperties_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForServerProperties_ARM(subject ServerProperties_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServerProperties_ARM
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

// Generator of ServerProperties_ARM instances for property testing - lazily instantiated by
// ServerProperties_ARMGenerator()
var serverProperties_ARMGenerator gopter.Gen

// ServerProperties_ARMGenerator returns a generator of ServerProperties_ARM instances for property testing.
// We first initialize serverProperties_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ServerProperties_ARMGenerator() gopter.Gen {
	if serverProperties_ARMGenerator != nil {
		return serverProperties_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServerProperties_ARM(generators)
	serverProperties_ARMGenerator = gen.Struct(reflect.TypeOf(ServerProperties_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServerProperties_ARM(generators)
	AddRelatedPropertyGeneratorsForServerProperties_ARM(generators)
	serverProperties_ARMGenerator = gen.Struct(reflect.TypeOf(ServerProperties_ARM{}), generators)

	return serverProperties_ARMGenerator
}

// AddIndependentPropertyGeneratorsForServerProperties_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServerProperties_ARM(gens map[string]gopter.Gen) {
	gens["AdministratorLogin"] = gen.PtrOf(gen.AlphaString())
	gens["AdministratorLoginPassword"] = gen.PtrOf(gen.AlphaString())
	gens["AvailabilityZone"] = gen.PtrOf(gen.AlphaString())
	gens["CreateMode"] = gen.PtrOf(gen.OneConstOf(
		ServerProperties_CreateMode_ARM_Create,
		ServerProperties_CreateMode_ARM_Default,
		ServerProperties_CreateMode_ARM_PointInTimeRestore,
		ServerProperties_CreateMode_ARM_Update))
	gens["PointInTimeUTC"] = gen.PtrOf(gen.AlphaString())
	gens["SourceServerResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["Version"] = gen.PtrOf(gen.OneConstOf(
		ServerVersion_ARM_11,
		ServerVersion_ARM_12,
		ServerVersion_ARM_13,
		ServerVersion_ARM_14))
}

// AddRelatedPropertyGeneratorsForServerProperties_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServerProperties_ARM(gens map[string]gopter.Gen) {
	gens["Backup"] = gen.PtrOf(Backup_ARMGenerator())
	gens["HighAvailability"] = gen.PtrOf(HighAvailability_ARMGenerator())
	gens["MaintenanceWindow"] = gen.PtrOf(MaintenanceWindow_ARMGenerator())
	gens["Network"] = gen.PtrOf(Network_ARMGenerator())
	gens["Storage"] = gen.PtrOf(Storage_ARMGenerator())
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
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tier"] = gen.PtrOf(gen.OneConstOf(Sku_Tier_ARM_Burstable, Sku_Tier_ARM_GeneralPurpose, Sku_Tier_ARM_MemoryOptimized))
}

func Test_Storage_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Storage_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorage_ARM, Storage_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorage_ARM runs a test to see if a specific instance of Storage_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForStorage_ARM(subject Storage_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Storage_ARM
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

// Generator of Storage_ARM instances for property testing - lazily instantiated by Storage_ARMGenerator()
var storage_ARMGenerator gopter.Gen

// Storage_ARMGenerator returns a generator of Storage_ARM instances for property testing.
func Storage_ARMGenerator() gopter.Gen {
	if storage_ARMGenerator != nil {
		return storage_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorage_ARM(generators)
	storage_ARMGenerator = gen.Struct(reflect.TypeOf(Storage_ARM{}), generators)

	return storage_ARMGenerator
}

// AddIndependentPropertyGeneratorsForStorage_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForStorage_ARM(gens map[string]gopter.Gen) {
	gens["StorageSizeGB"] = gen.PtrOf(gen.Int())
}
