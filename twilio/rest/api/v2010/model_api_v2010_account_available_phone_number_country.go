/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.11.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ApiV2010AccountAvailablePhoneNumberCountry struct for ApiV2010AccountAvailablePhoneNumberCountry
type ApiV2010AccountAvailablePhoneNumberCountry struct {
	// Whether all phone numbers available in the country are new to the Twilio platform.
	Beta *bool `json:"Beta,omitempty"`
	// The name of the country
	Country *string `json:"Country,omitempty"`
	// The ISO-3166-1 country code of the country.
	CountryCode *string `json:"CountryCode,omitempty"`
	// A list of related resources identified by their relative URIs
	SubresourceUris *map[string]interface{} `json:"SubresourceUris,omitempty"`
	// The URI of the Country resource, relative to `https://api.twilio.com`
	Uri *string `json:"Uri,omitempty"`
}
