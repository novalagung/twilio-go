/*
 * Twilio - Numbers
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.11.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListSupportingDocumentTypeResponse struct for ListSupportingDocumentTypeResponse
type ListSupportingDocumentTypeResponse struct {
	Meta                    ListBundleResponseMeta                                `json:"Meta,omitempty"`
	SupportingDocumentTypes []NumbersV2RegulatoryComplianceSupportingDocumentType `json:"SupportingDocumentTypes,omitempty"`
}
