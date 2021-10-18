/*
 * Twilio - Conversations
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.22.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ConversationsV1ServiceNotification struct for ConversationsV1ServiceNotification
type ConversationsV1ServiceNotification struct {
	// The unique ID of the Account responsible for this configuration.
	AccountSid *string `json:"account_sid,omitempty"`
	// The Push Notification configuration for being added to a Conversation.
	AddedToConversation *map[string]interface{} `json:"added_to_conversation,omitempty"`
	// The SID of the Conversation Service that the Configuration applies to.
	ChatServiceSid *string `json:"chat_service_sid,omitempty"`
	// Weather the notification logging is enabled.
	LogEnabled *bool `json:"log_enabled,omitempty"`
	// The Push Notification configuration for New Messages.
	NewMessage *map[string]interface{} `json:"new_message,omitempty"`
	// The Push Notification configuration for being removed from a Conversation.
	RemovedFromConversation *map[string]interface{} `json:"removed_from_conversation,omitempty"`
	// An absolute URL for this configuration.
	Url *string `json:"url,omitempty"`
}
