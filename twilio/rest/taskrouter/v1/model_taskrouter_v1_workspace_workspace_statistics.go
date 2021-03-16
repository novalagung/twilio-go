/*
 * Twilio - Taskrouter
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.11.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// TaskrouterV1WorkspaceWorkspaceStatistics struct for TaskrouterV1WorkspaceWorkspaceStatistics
type TaskrouterV1WorkspaceWorkspaceStatistics struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"AccountSid,omitempty"`
	// An object that contains the cumulative statistics for the Workspace
	Cumulative *map[string]interface{} `json:"Cumulative,omitempty"`
	// n object that contains the real-time statistics for the Workspace
	Realtime *map[string]interface{} `json:"Realtime,omitempty"`
	// The absolute URL of the Workspace statistics resource
	Url *string `json:"Url,omitempty"`
	// The SID of the Workspace
	WorkspaceSid *string `json:"WorkspaceSid,omitempty"`
}
