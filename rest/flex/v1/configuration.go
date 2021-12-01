/*
 * Twilio - Flex
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.23.2
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"net/url"
)

// Optional parameters for the method 'FetchConfiguration'
type FetchConfigurationParams struct {
	// The Pinned UI version of the Configuration resource to fetch.
	UiVersion *string `json:"UiVersion,omitempty"`
}

func (params *FetchConfigurationParams) SetUiVersion(UiVersion string) *FetchConfigurationParams {
	params.UiVersion = &UiVersion
	return params
}

func (c *ApiService) FetchConfiguration(params *FetchConfigurationParams) (*FlexV1Configuration, error) {
	path := "/v1/Configuration"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.UiVersion != nil {
		data.Set("UiVersion", *params.UiVersion)
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &FlexV1Configuration{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
