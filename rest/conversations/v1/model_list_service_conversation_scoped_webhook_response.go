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

// ListServiceConversationScopedWebhookResponse struct for ListServiceConversationScopedWebhookResponse
type ListServiceConversationScopedWebhookResponse struct {
	Meta     ListConversationResponseMeta                      `json:"meta,omitempty"`
	Webhooks []ConversationsV1ServiceConversationScopedWebhook `json:"webhooks,omitempty"`
}
