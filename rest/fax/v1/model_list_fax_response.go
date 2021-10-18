/*
 * Twilio - Fax
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.22.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListFaxResponse struct for ListFaxResponse
type ListFaxResponse struct {
	Faxes []FaxV1Fax          `json:"faxes,omitempty"`
	Meta  ListFaxResponseMeta `json:"meta,omitempty"`
}
