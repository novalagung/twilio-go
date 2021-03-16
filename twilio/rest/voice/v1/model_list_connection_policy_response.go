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

// ListConnectionPolicyResponse struct for ListConnectionPolicyResponse
type ListConnectionPolicyResponse struct {
	ConnectionPolicies []VoiceV1ConnectionPolicy `json:"ConnectionPolicies,omitempty"`
	Meta               ListByocTrunkResponseMeta `json:"Meta,omitempty"`
}
