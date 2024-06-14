// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20211101

type Namespaces_Topics_Subscriptions_Rule_STATUS_ARM struct {
	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Location: The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// Properties: Properties of Rule resource
	Properties *Ruleproperties_STATUS_ARM `json:"properties,omitempty"`

	// SystemData: The system meta data relating to this resource.
	SystemData *SystemData_STATUS_ARM `json:"systemData,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.EventHub/Namespaces" or "Microsoft.EventHub/Namespaces/EventHubs"
	Type *string `json:"type,omitempty"`
}

// Description of Rule Resource.
type Ruleproperties_STATUS_ARM struct {
	// Action: Represents the filter actions which are allowed for the transformation of a message that have been matched by a
	// filter expression.
	Action *Action_STATUS_ARM `json:"action,omitempty"`

	// CorrelationFilter: Properties of correlationFilter
	CorrelationFilter *CorrelationFilter_STATUS_ARM `json:"correlationFilter,omitempty"`

	// FilterType: Filter type that is evaluated against a BrokeredMessage.
	FilterType *FilterType_STATUS_ARM `json:"filterType,omitempty"`

	// SqlFilter: Properties of sqlFilter
	SqlFilter *SqlFilter_STATUS_ARM `json:"sqlFilter,omitempty"`
}

// Represents the filter actions which are allowed for the transformation of a message that have been matched by a filter
// expression.
type Action_STATUS_ARM struct {
	// CompatibilityLevel: This property is reserved for future use. An integer value showing the compatibility level,
	// currently hard-coded to 20.
	CompatibilityLevel *int `json:"compatibilityLevel,omitempty"`

	// RequiresPreprocessing: Value that indicates whether the rule action requires preprocessing.
	RequiresPreprocessing *bool `json:"requiresPreprocessing,omitempty"`

	// SqlExpression: SQL expression. e.g. MyProperty='ABC'
	SqlExpression *string `json:"sqlExpression,omitempty"`
}

// Represents the correlation filter expression.
type CorrelationFilter_STATUS_ARM struct {
	// ContentType: Content type of the message.
	ContentType *string `json:"contentType,omitempty"`

	// CorrelationId: Identifier of the correlation.
	CorrelationId *string `json:"correlationId,omitempty"`

	// Label: Application specific label.
	Label *string `json:"label,omitempty"`

	// MessageId: Identifier of the message.
	MessageId *string `json:"messageId,omitempty"`

	// Properties: dictionary object for custom filters
	Properties map[string]string `json:"properties,omitempty"`

	// ReplyTo: Address of the queue to reply to.
	ReplyTo *string `json:"replyTo,omitempty"`

	// ReplyToSessionId: Session identifier to reply to.
	ReplyToSessionId *string `json:"replyToSessionId,omitempty"`

	// RequiresPreprocessing: Value that indicates whether the rule action requires preprocessing.
	RequiresPreprocessing *bool `json:"requiresPreprocessing,omitempty"`

	// SessionId: Session identifier.
	SessionId *string `json:"sessionId,omitempty"`

	// To: Address to send to.
	To *string `json:"to,omitempty"`
}

// Rule filter types
type FilterType_STATUS_ARM string

const (
	FilterType_STATUS_ARM_CorrelationFilter = FilterType_STATUS_ARM("CorrelationFilter")
	FilterType_STATUS_ARM_SqlFilter         = FilterType_STATUS_ARM("SqlFilter")
)

// Mapping from string to FilterType_STATUS_ARM
var filterType_STATUS_ARM_Values = map[string]FilterType_STATUS_ARM{
	"correlationfilter": FilterType_STATUS_ARM_CorrelationFilter,
	"sqlfilter":         FilterType_STATUS_ARM_SqlFilter,
}

// Represents a filter which is a composition of an expression and an action that is executed in the pub/sub pipeline.
type SqlFilter_STATUS_ARM struct {
	// CompatibilityLevel: This property is reserved for future use. An integer value showing the compatibility level,
	// currently hard-coded to 20.
	CompatibilityLevel *int `json:"compatibilityLevel,omitempty"`

	// RequiresPreprocessing: Value that indicates whether the rule action requires preprocessing.
	RequiresPreprocessing *bool `json:"requiresPreprocessing,omitempty"`

	// SqlExpression: The SQL expression. e.g. MyProperty='ABC'
	SqlExpression *string `json:"sqlExpression,omitempty"`
}
