/*
 * Twilio - Voice
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.11.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// VoiceV1DialingPermissionsDialingPermissionsCountry struct for VoiceV1DialingPermissionsDialingPermissionsCountry
type VoiceV1DialingPermissionsDialingPermissionsCountry struct {
	// The name of the continent in which the country is located
	Continent *string `json:"Continent,omitempty"`
	// The E.164 assigned country codes(s)
	CountryCodes *[]string `json:"CountryCodes,omitempty"`
	// Whether dialing to high-risk special services numbers is enabled
	HighRiskSpecialNumbersEnabled *bool `json:"HighRiskSpecialNumbersEnabled,omitempty"`
	// Whether dialing to high-risk toll fraud numbers is enabled, else `false`
	HighRiskTollfraudNumbersEnabled *bool `json:"HighRiskTollfraudNumbersEnabled,omitempty"`
	// The ISO country code
	IsoCode *string `json:"IsoCode,omitempty"`
	// A list of URLs related to this resource
	Links *map[string]interface{} `json:"Links,omitempty"`
	// Whether dialing to low-risk numbers is enabled
	LowRiskNumbersEnabled *bool `json:"LowRiskNumbersEnabled,omitempty"`
	// The name of the country
	Name *string `json:"Name,omitempty"`
	// The absolute URL of this resource
	Url *string `json:"Url,omitempty"`
}
