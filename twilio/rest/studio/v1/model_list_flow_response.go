/*
 * Twilio - Studio
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.11.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListFlowResponse struct for ListFlowResponse
type ListFlowResponse struct {
	Flows []StudioV1Flow       `json:"Flows,omitempty"`
	Meta  ListFlowResponseMeta `json:"Meta,omitempty"`
}
