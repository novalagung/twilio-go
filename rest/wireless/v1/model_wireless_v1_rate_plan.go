/*
 * Twilio - Wireless
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.22.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// WirelessV1RatePlan struct for WirelessV1RatePlan
type WirelessV1RatePlan struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// Whether SIMs can use GPRS/3G/4G/LTE data connectivity
	DataEnabled *bool `json:"data_enabled,omitempty"`
	// The total data usage in Megabytes that the Network allows during one month on the home network
	DataLimit *int `json:"data_limit,omitempty"`
	// The model used to meter data usage
	DataMetering *string `json:"data_metering,omitempty"`
	// The date when the resource was created, given as GMT in ISO 8601 format
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The date when the resource was last updated, given as GMT in ISO 8601 format
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The string that you assigned to describe the resource
	FriendlyName *string `json:"friendly_name,omitempty"`
	// The services that SIMs capable of using GPRS/3G/4G/LTE data connectivity can use outside of the United States
	InternationalRoaming *[]string `json:"international_roaming,omitempty"`
	// The total data usage (download and upload combined) in Megabytes that the Network allows during one month when roaming outside the United States
	InternationalRoamingDataLimit *int `json:"international_roaming_data_limit,omitempty"`
	// Whether SIMs can make, send, and receive SMS using Commands
	MessagingEnabled *bool `json:"messaging_enabled,omitempty"`
	// The total data usage in Megabytes that the Network allows during one month on non-home networks in the United States
	NationalRoamingDataLimit *int `json:"national_roaming_data_limit,omitempty"`
	// Whether SIMs can roam on networks other than the home network in the United States
	NationalRoamingEnabled *bool `json:"national_roaming_enabled,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// An application-defined string that uniquely identifies the resource
	UniqueName *string `json:"unique_name,omitempty"`
	// The absolute URL of the resource
	Url *string `json:"url,omitempty"`
	// Whether SIMs can make and receive voice calls
	VoiceEnabled *bool `json:"voice_enabled,omitempty"`
}
