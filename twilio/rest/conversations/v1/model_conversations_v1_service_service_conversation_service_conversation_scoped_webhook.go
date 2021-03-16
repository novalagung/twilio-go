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

import (
	"time"
)

// ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook struct for ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook
type ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook struct {
	// The unique ID of the Account responsible for this conversation.
	AccountSid *string `json:"AccountSid,omitempty"`
	// The SID of the Conversation Service that the resource is associated with.
	ChatServiceSid *string `json:"ChatServiceSid,omitempty"`
	// The configuration of this webhook.
	Configuration *map[string]interface{} `json:"Configuration,omitempty"`
	// The unique ID of the Conversation for this webhook.
	ConversationSid *string `json:"ConversationSid,omitempty"`
	// The date that this resource was created.
	DateCreated *time.Time `json:"DateCreated,omitempty"`
	// The date that this resource was last updated.
	DateUpdated *time.Time `json:"DateUpdated,omitempty"`
	// A 34 character string that uniquely identifies this resource.
	Sid *string `json:"Sid,omitempty"`
	// The target of this webhook.
	Target *string `json:"Target,omitempty"`
	// An absolute URL for this webhook.
	Url *string `json:"Url,omitempty"`
}
