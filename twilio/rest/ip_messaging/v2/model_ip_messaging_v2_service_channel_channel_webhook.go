/*
 * Twilio - Ip_messaging
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

// IpMessagingV2ServiceChannelChannelWebhook struct for IpMessagingV2ServiceChannelChannelWebhook
type IpMessagingV2ServiceChannelChannelWebhook struct {
	AccountSid    *string                 `json:"AccountSid,omitempty"`
	ChannelSid    *string                 `json:"ChannelSid,omitempty"`
	Configuration *map[string]interface{} `json:"Configuration,omitempty"`
	DateCreated   *time.Time              `json:"DateCreated,omitempty"`
	DateUpdated   *time.Time              `json:"DateUpdated,omitempty"`
	ServiceSid    *string                 `json:"ServiceSid,omitempty"`
	Sid           *string                 `json:"Sid,omitempty"`
	Type          *string                 `json:"Type,omitempty"`
	Url           *string                 `json:"Url,omitempty"`
}
