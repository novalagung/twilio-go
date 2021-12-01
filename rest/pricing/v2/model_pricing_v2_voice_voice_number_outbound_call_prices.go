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

// PricingV2VoiceVoiceNumberOutboundCallPrices struct for PricingV2VoiceVoiceNumberOutboundCallPrices
type PricingV2VoiceVoiceNumberOutboundCallPrices struct {
	BasePrice           float32  `json:"base_price,omitempty"`
	CurrentPrice        float32  `json:"current_price,omitempty"`
	OriginationPrefixes []string `json:"origination_prefixes,omitempty"`
}
