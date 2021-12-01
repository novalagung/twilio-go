/*
 * Twilio - Supersim
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.23.2
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListIpCommandResponse struct for ListIpCommandResponse
type ListIpCommandResponse struct {
	IpCommands []SupersimV1IpCommand   `json:"ip_commands,omitempty"`
	Meta       ListCommandResponseMeta `json:"meta,omitempty"`
}
