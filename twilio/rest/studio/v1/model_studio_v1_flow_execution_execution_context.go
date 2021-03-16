/*
 * Twilio - Studio
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.11.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// StudioV1FlowExecutionExecutionContext struct for StudioV1FlowExecutionExecutionContext
type StudioV1FlowExecutionExecutionContext struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"AccountSid,omitempty"`
	// The current state of the flow
	Context *map[string]interface{} `json:"Context,omitempty"`
	// The SID of the Execution
	ExecutionSid *string `json:"ExecutionSid,omitempty"`
	// The SID of the Flow
	FlowSid *string `json:"FlowSid,omitempty"`
	// The absolute URL of the resource
	Url *string `json:"Url,omitempty"`
}
