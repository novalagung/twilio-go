/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.23.2
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ApiV2010CallFeedback struct for ApiV2010CallFeedback
type ApiV2010CallFeedback struct {
	// The unique sid that identifies this account
	AccountSid *string `json:"account_sid,omitempty"`
	// The date this resource was created
	DateCreated *string `json:"date_created,omitempty"`
	// The date this resource was last updated
	DateUpdated *string `json:"date_updated,omitempty"`
	// Issues experienced during the call
	Issues *[]string `json:"issues,omitempty"`
	// 1 to 5 quality score
	QualityScore *int `json:"quality_score,omitempty"`
	// A string that uniquely identifies this feedback resource
	Sid *string `json:"sid,omitempty"`
}
