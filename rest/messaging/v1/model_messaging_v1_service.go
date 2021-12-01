/*
 * Twilio - Messaging
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

// MessagingV1Service struct for MessagingV1Service
type MessagingV1Service struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// Whether to enable Area Code Geomatch on the Service Instance
	AreaCodeGeomatch *bool `json:"area_code_geomatch,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The HTTP method we use to call fallback_url
	FallbackMethod *string `json:"fallback_method,omitempty"`
	// Whether to enable Fallback to Long Code for messages sent through the Service instance
	FallbackToLongCode *bool `json:"fallback_to_long_code,omitempty"`
	// The URL that we call using fallback_method if an error occurs while retrieving or executing the TwiML from the Inbound Request URL. This field will be overridden if the `use_inbound_webhook_on_number` field is enabled.
	FallbackUrl *string `json:"fallback_url,omitempty"`
	// The string that you assigned to describe the resource
	FriendlyName *string `json:"friendly_name,omitempty"`
	// The HTTP method we use to call inbound_request_url
	InboundMethod *string `json:"inbound_method,omitempty"`
	// The URL we call using inbound_method when a message is received by any phone number or short code in the Service. This field will be overridden if the `use_inbound_webhook_on_number` field is enabled.
	InboundRequestUrl *string `json:"inbound_request_url,omitempty"`
	// The absolute URLs of related resources
	Links *map[string]interface{} `json:"links,omitempty"`
	// Whether to enable the MMS Converter for messages sent through the Service instance
	MmsConverter *bool `json:"mms_converter,omitempty"`
	// Reserved
	ScanMessageContent *string `json:"scan_message_content,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// Whether to enable Encoding for messages sent through the Service instance
	SmartEncoding *bool `json:"smart_encoding,omitempty"`
	// The URL we call to pass status updates about message delivery
	StatusCallback *string `json:"status_callback,omitempty"`
	// Whether to enable Sticky Sender on the Service instance
	StickySender *bool `json:"sticky_sender,omitempty"`
	// Reserved
	SynchronousValidation *bool `json:"synchronous_validation,omitempty"`
	// The absolute URL of the Service resource
	Url *string `json:"url,omitempty"`
	// Whether US A2P campaign is registered for this Service.
	UsAppToPersonRegistered *bool `json:"us_app_to_person_registered,omitempty"`
	// If enabled, the webhook url configured on the phone number will be used and will override the `inbound_request_url`/`fallback_url` url called when an inbound message is received.
	UseInboundWebhookOnNumber *bool `json:"use_inbound_webhook_on_number,omitempty"`
	// A string describing the scenario in which the Messaging Service will be used
	Usecase *string `json:"usecase,omitempty"`
	// How long, in seconds, messages sent from the Service are valid
	ValidityPeriod *int `json:"validity_period,omitempty"`
}
