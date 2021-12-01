/*
 * Twilio - Wireless
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.23.2
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListRatePlanResponse struct for ListRatePlanResponse
type ListRatePlanResponse struct {
	Meta      ListCommandResponseMeta `json:"meta,omitempty"`
	RatePlans []WirelessV1RatePlan    `json:"rate_plans,omitempty"`
}
