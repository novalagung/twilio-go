/*
 * Twilio - Serverless
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

// ServerlessV1ServiceEnvironmentLog struct for ServerlessV1ServiceEnvironmentLog
type ServerlessV1ServiceEnvironmentLog struct {
	// The SID of the Account that created the Log resource
	AccountSid *string `json:"AccountSid,omitempty"`
	// The SID of the build that corresponds to the log
	BuildSid *string `json:"BuildSid,omitempty"`
	// The ISO 8601 date and time in GMT when the Log resource was created
	DateCreated *time.Time `json:"DateCreated,omitempty"`
	// The SID of the deployment that corresponds to the log
	DeploymentSid *string `json:"DeploymentSid,omitempty"`
	// The SID of the environment in which the log occurred
	EnvironmentSid *string `json:"EnvironmentSid,omitempty"`
	// The SID of the function whose invocation produced the log
	FunctionSid *string `json:"FunctionSid,omitempty"`
	// The log level
	Level *string `json:"Level,omitempty"`
	// The log message
	Message *string `json:"Message,omitempty"`
	// The SID of the request associated with the log
	RequestSid *string `json:"RequestSid,omitempty"`
	// The SID of the Service that the Log resource is associated with
	ServiceSid *string `json:"ServiceSid,omitempty"`
	// The unique string that identifies the Log resource
	Sid *string `json:"Sid,omitempty"`
	// The absolute URL of the Log resource
	Url *string `json:"Url,omitempty"`
}
