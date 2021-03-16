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

// TaskrouterV1WorkspaceWorkerWorkerInstanceStatistics struct for TaskrouterV1WorkspaceWorkerWorkerInstanceStatistics
type TaskrouterV1WorkspaceWorkerWorkerInstanceStatistics struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"AccountSid,omitempty"`
	// An object that contains the cumulative statistics for the Worker
	Cumulative *map[string]interface{} `json:"Cumulative,omitempty"`
	// The absolute URL of the WorkerChannel statistics resource
	Url *string `json:"Url,omitempty"`
	// The SID of the Worker that contains the WorkerChannel
	WorkerSid *string `json:"WorkerSid,omitempty"`
	// The SID of the Workspace that contains the WorkerChannel
	WorkspaceSid *string `json:"WorkspaceSid,omitempty"`
}
