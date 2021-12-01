/*
 * Twilio - Flex
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.23.2
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// FlexV1FlexFlow struct for FlexV1FlexFlow
type FlexV1FlexFlow struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The channel type
	ChannelType *string `json:"channel_type,omitempty"`
	// The SID of the chat service
	ChatServiceSid *string `json:"chat_service_sid,omitempty"`
	// The channel contact's Identity
	ContactIdentity *string `json:"contact_identity,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// Whether the Flex Flow is enabled
	Enabled *bool `json:"enabled,omitempty"`
	// The string that you assigned to describe the resource
	FriendlyName *string `json:"friendly_name,omitempty"`
	// An object that contains specific parameters for the integration
	Integration *map[string]interface{} `json:"integration,omitempty"`
	// The software that will handle inbound messages.
	IntegrationType *string `json:"integration_type,omitempty"`
	// Remove active Proxy sessions if the corresponding Task is deleted.
	JanitorEnabled *bool `json:"janitor_enabled,omitempty"`
	// Re-use this chat channel for future interactions with a contact
	LongLived *bool `json:"long_lived,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// The absolute URL of the Flex Flow resource
	Url *string `json:"url,omitempty"`
}
