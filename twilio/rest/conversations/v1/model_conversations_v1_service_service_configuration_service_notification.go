/*
 * Twilio - Conversations
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.11.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ConversationsV1ServiceServiceConfigurationServiceNotification struct for ConversationsV1ServiceServiceConfigurationServiceNotification
type ConversationsV1ServiceServiceConfigurationServiceNotification struct {
	// The unique ID of the Account responsible for this configuration.
	AccountSid *string `json:"AccountSid,omitempty"`
	// The Push Notification configuration for being added to a Conversation.
	AddedToConversation *map[string]interface{} `json:"AddedToConversation,omitempty"`
	// The SID of the Conversation Service that the Configuration applies to.
	ChatServiceSid *string `json:"ChatServiceSid,omitempty"`
	// Weather the notification logging is enabled.
	LogEnabled *bool `json:"LogEnabled,omitempty"`
	// The Push Notification configuration for New Messages.
	NewMessage *map[string]interface{} `json:"NewMessage,omitempty"`
	// The Push Notification configuration for being removed from a Conversation.
	RemovedFromConversation *map[string]interface{} `json:"RemovedFromConversation,omitempty"`
	// An absolute URL for this configuration.
	Url *string `json:"Url,omitempty"`
}
