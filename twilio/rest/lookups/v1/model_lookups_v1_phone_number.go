/*
 * Twilio - Lookups
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.11.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// LookupsV1PhoneNumber struct for LookupsV1PhoneNumber
type LookupsV1PhoneNumber struct {
	// A JSON string with the results of the Add-ons you specified
	AddOns *map[string]interface{} `json:"AddOns,omitempty"`
	// The name of the phone number's owner
	CallerName *map[string]interface{} `json:"CallerName,omitempty"`
	// The telecom company that provides the phone number
	Carrier *map[string]interface{} `json:"Carrier,omitempty"`
	// The ISO country code for the phone number
	CountryCode *string `json:"CountryCode,omitempty"`
	// The phone number, in national format
	NationalFormat *string `json:"NationalFormat,omitempty"`
	// The phone number in E.164 format
	PhoneNumber *string `json:"PhoneNumber,omitempty"`
	// The absolute URL of the resource
	Url *string `json:"Url,omitempty"`
}
