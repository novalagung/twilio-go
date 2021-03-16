/*
 * Twilio - Pricing
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.11.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// PricingV2VoiceVoiceNumberOutboundCallPrices struct for PricingV2VoiceVoiceNumberOutboundCallPrices
type PricingV2VoiceVoiceNumberOutboundCallPrices struct {
	BasePrice           float32  `json:"BasePrice,omitempty"`
	CurrentPrice        float32  `json:"CurrentPrice,omitempty"`
	OriginationPrefixes []string `json:"OriginationPrefixes,omitempty"`
}
