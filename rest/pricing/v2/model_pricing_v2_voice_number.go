/*
 * Twilio - Pricing
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.23.2
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// PricingV2VoiceNumber struct for PricingV2VoiceNumber
type PricingV2VoiceNumber struct {
	// The name of the country
	Country *string `json:"country,omitempty"`
	// The destination phone number, in E.164 format
	DestinationNumber *string                                    `json:"destination_number,omitempty"`
	InboundCallPrice  *PricingV2VoiceVoiceNumberInboundCallPrice `json:"inbound_call_price,omitempty"`
	// The ISO country code
	IsoCountry *string `json:"iso_country,omitempty"`
	// The origination phone number, in E.164 format
	OriginationNumber *string `json:"origination_number,omitempty"`
	// The list of OutboundCallPriceWithOrigin records
	OutboundCallPrices *[]PricingV2VoiceVoiceNumberOutboundCallPrices `json:"outbound_call_prices,omitempty"`
	// The currency in which prices are measured, in ISO 4127 format (e.g. usd, eur, jpy)
	PriceUnit *string `json:"price_unit,omitempty"`
	// The absolute URL of the resource
	Url *string `json:"url,omitempty"`
}
