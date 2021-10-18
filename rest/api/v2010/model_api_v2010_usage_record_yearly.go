/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.22.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ApiV2010UsageRecordYearly struct for ApiV2010UsageRecordYearly
type ApiV2010UsageRecordYearly struct {
	// The SID of the Account accrued the usage
	AccountSid *string `json:"account_sid,omitempty"`
	// The API version used to create the resource
	ApiVersion *string `json:"api_version,omitempty"`
	// Usage records up to date as of this timestamp
	AsOf *string `json:"as_of,omitempty"`
	// The category of usage
	Category *string `json:"category,omitempty"`
	// The number of usage events
	Count *string `json:"count,omitempty"`
	// The units in which count is measured
	CountUnit *string `json:"count_unit,omitempty"`
	// A plain-language description of the usage category
	Description *string `json:"description,omitempty"`
	// The last date for which usage is included in the UsageRecord
	EndDate *string `json:"end_date,omitempty"`
	// The total price of the usage
	Price *float32 `json:"price,omitempty"`
	// The currency in which `price` is measured
	PriceUnit *string `json:"price_unit,omitempty"`
	// The first date for which usage is included in this UsageRecord
	StartDate *string `json:"start_date,omitempty"`
	// A list of related resources identified by their relative URIs
	SubresourceUris *map[string]interface{} `json:"subresource_uris,omitempty"`
	// The URI of the resource, relative to `https://api.twilio.com`
	Uri *string `json:"uri,omitempty"`
	// The amount of usage
	Usage *string `json:"usage,omitempty"`
	// The units in which usage is measured
	UsageUnit *string `json:"usage_unit,omitempty"`
}
