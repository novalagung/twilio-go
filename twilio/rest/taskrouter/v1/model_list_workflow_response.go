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

// ListWorkflowResponse struct for ListWorkflowResponse
type ListWorkflowResponse struct {
	Meta      ListWorkspaceResponseMeta       `json:"Meta,omitempty"`
	Workflows []TaskrouterV1WorkspaceWorkflow `json:"Workflows,omitempty"`
}
