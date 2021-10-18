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

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

// Optional parameters for the method 'CreateToken'
type CreateTokenParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// The duration in seconds for which the generated credentials are valid. The default value is 86400 (24 hours).
	Ttl *int `json:"Ttl,omitempty"`
}

func (params *CreateTokenParams) SetPathAccountSid(PathAccountSid string) *CreateTokenParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *CreateTokenParams) SetTtl(Ttl int) *CreateTokenParams {
	params.Ttl = &Ttl
	return params
}

// Create a new token for ICE servers
func (c *ApiService) CreateToken(params *CreateTokenParams) (*ApiV2010Token, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Tokens.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Ttl != nil {
		data.Set("Ttl", fmt.Sprint(*params.Ttl))
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010Token{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
