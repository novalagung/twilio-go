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

// ListTaskChannelResponse struct for ListTaskChannelResponse
type ListTaskChannelResponse struct {
	Channels []TaskrouterV1WorkspaceTaskChannel `json:"Channels,omitempty"`
	Meta     ListWorkspaceResponseMeta          `json:"Meta,omitempty"`
}
