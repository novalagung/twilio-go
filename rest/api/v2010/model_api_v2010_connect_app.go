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

// ApiV2010ConnectApp struct for ApiV2010ConnectApp
type ApiV2010ConnectApp struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The URL to redirect the user to after authorization
	AuthorizeRedirectUrl *string `json:"authorize_redirect_url,omitempty"`
	// The company name set for the Connect App
	CompanyName *string `json:"company_name,omitempty"`
	// The HTTP method we use to call deauthorize_callback_url
	DeauthorizeCallbackMethod *string `json:"deauthorize_callback_method,omitempty"`
	// The URL we call to de-authorize the Connect App
	DeauthorizeCallbackUrl *string `json:"deauthorize_callback_url,omitempty"`
	// The description of the Connect App
	Description *string `json:"description,omitempty"`
	// The string that you assigned to describe the resource
	FriendlyName *string `json:"friendly_name,omitempty"`
	// The URL users can obtain more information
	HomepageUrl *string `json:"homepage_url,omitempty"`
	// The set of permissions that your ConnectApp requests
	Permissions *[]string `json:"permissions,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// The URI of the resource, relative to `https://api.twilio.com`
	Uri *string `json:"uri,omitempty"`
}
