/*
 * Twilio - Events
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.11.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListSubscribedEventResponse struct for ListSubscribedEventResponse
type ListSubscribedEventResponse struct {
	Meta  ListVersionResponseMeta               `json:"Meta,omitempty"`
	Types []EventsV1SubscriptionSubscribedEvent `json:"Types,omitempty"`
}
