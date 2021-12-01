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

import (
	"encoding/json"
	"fmt"
	"net/url"

	"strings"

	"github.com/twilio/twilio-go/client"
)

// Optional parameters for the method 'CreateNewSigningKey'
type CreateNewSigningKeyParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will be responsible for the new Key resource.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// A descriptive string that you create to describe the resource. It can be up to 64 characters long.
	FriendlyName *string `json:"FriendlyName,omitempty"`
}

func (params *CreateNewSigningKeyParams) SetPathAccountSid(PathAccountSid string) *CreateNewSigningKeyParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *CreateNewSigningKeyParams) SetFriendlyName(FriendlyName string) *CreateNewSigningKeyParams {
	params.FriendlyName = &FriendlyName
	return params
}

// Create a new Signing Key for the account making the request.
func (c *ApiService) CreateNewSigningKey(params *CreateNewSigningKeyParams) (*ApiV2010NewSigningKey, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/SigningKeys.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010NewSigningKey{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'DeleteSigningKey'
type DeleteSigningKeyParams struct {
	//
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *DeleteSigningKeyParams) SetPathAccountSid(PathAccountSid string) *DeleteSigningKeyParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

func (c *ApiService) DeleteSigningKey(Sid string, params *DeleteSigningKeyParams) error {
	path := "/2010-04-01/Accounts/{AccountSid}/SigningKeys/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

// Optional parameters for the method 'FetchSigningKey'
type FetchSigningKeyParams struct {
	//
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *FetchSigningKeyParams) SetPathAccountSid(PathAccountSid string) *FetchSigningKeyParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

func (c *ApiService) FetchSigningKey(Sid string, params *FetchSigningKeyParams) (*ApiV2010SigningKey, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/SigningKeys/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010SigningKey{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListSigningKey'
type ListSigningKeyParams struct {
	//
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListSigningKeyParams) SetPathAccountSid(PathAccountSid string) *ListSigningKeyParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *ListSigningKeyParams) SetPageSize(PageSize int) *ListSigningKeyParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListSigningKeyParams) SetLimit(Limit int) *ListSigningKeyParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of SigningKey records from the API. Request is executed immediately.
func (c *ApiService) PageSigningKey(params *ListSigningKeyParams, pageToken, pageNumber string) (*ListSigningKeyResponse, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/SigningKeys.json"

	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	if pageToken != "" {
		data.Set("PageToken", pageToken)
	}
	if pageNumber != "" {
		data.Set("Page", pageNumber)
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListSigningKeyResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists SigningKey records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListSigningKey(params *ListSigningKeyParams) ([]ApiV2010SigningKey, error) {
	if params == nil {
		params = &ListSigningKeyParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageSigningKey(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []ApiV2010SigningKey

	for response != nil {
		records = append(records, response.SigningKeys...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListSigningKeyResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListSigningKeyResponse)
	}

	return records, err
}

// Streams SigningKey records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamSigningKey(params *ListSigningKeyParams) (chan ApiV2010SigningKey, error) {
	if params == nil {
		params = &ListSigningKeyParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageSigningKey(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan ApiV2010SigningKey, 1)

	go func() {
		for response != nil {
			for item := range response.SigningKeys {
				channel <- response.SigningKeys[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListSigningKeyResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListSigningKeyResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListSigningKeyResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListSigningKeyResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateSigningKey'
type UpdateSigningKeyParams struct {
	//
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	//
	FriendlyName *string `json:"FriendlyName,omitempty"`
}

func (params *UpdateSigningKeyParams) SetPathAccountSid(PathAccountSid string) *UpdateSigningKeyParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *UpdateSigningKeyParams) SetFriendlyName(FriendlyName string) *UpdateSigningKeyParams {
	params.FriendlyName = &FriendlyName
	return params
}

func (c *ApiService) UpdateSigningKey(Sid string, params *UpdateSigningKeyParams) (*ApiV2010SigningKey, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/SigningKeys/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010SigningKey{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
