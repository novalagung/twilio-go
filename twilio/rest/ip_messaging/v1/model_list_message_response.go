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

// ListMessageResponse struct for ListMessageResponse
type ListMessageResponse struct {
	Messages []IpMessagingV1ServiceChannelMessage `json:"Messages,omitempty"`
	Meta     ListCredentialResponseMeta           `json:"Meta,omitempty"`
}
