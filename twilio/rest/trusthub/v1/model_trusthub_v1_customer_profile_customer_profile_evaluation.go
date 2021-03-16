/*
 * Twilio - Trusthub
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.11.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// TrusthubV1CustomerProfileCustomerProfileEvaluation struct for TrusthubV1CustomerProfileCustomerProfileEvaluation
type TrusthubV1CustomerProfileCustomerProfileEvaluation struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"AccountSid,omitempty"`
	// The unique string that identifies the resource
	CustomerProfileSid *string    `json:"CustomerProfileSid,omitempty"`
	DateCreated        *time.Time `json:"DateCreated,omitempty"`
	// The unique string of a policy
	PolicySid *string `json:"PolicySid,omitempty"`
	// The results of the Evaluation resource
	Results *[]map[string]interface{} `json:"Results,omitempty"`
	// The unique string that identifies the Evaluation resource
	Sid *string `json:"Sid,omitempty"`
	// The compliance status of the Evaluation resource
	Status *string `json:"Status,omitempty"`
	Url    *string `json:"Url,omitempty"`
}
