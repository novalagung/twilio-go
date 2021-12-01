/*
 * Twilio - Monitor
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.23.2
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListAlertResponse struct for ListAlertResponse
type ListAlertResponse struct {
	Alerts []MonitorV1Alert      `json:"alerts,omitempty"`
	Meta   ListAlertResponseMeta `json:"meta,omitempty"`
}
