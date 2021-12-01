/*
 * Twilio - Conversations
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.23.2
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListServiceRoleResponse struct for ListServiceRoleResponse
type ListServiceRoleResponse struct {
	Meta  ListConversationResponseMeta `json:"meta,omitempty"`
	Roles []ConversationsV1ServiceRole `json:"roles,omitempty"`
}
