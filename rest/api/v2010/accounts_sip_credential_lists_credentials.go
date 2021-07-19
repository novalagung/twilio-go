/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.19.0
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

// Optional parameters for the method 'CreateSipCredential'
type CreateSipCredentialParams struct {
	// The unique id of the Account that is responsible for this resource.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// The password that the username will use when authenticating SIP requests. The password must be a minimum of 12 characters, contain at least 1 digit, and have mixed case. (eg `IWasAtSignal2018`)
	Password *string `json:"Password,omitempty"`
	// The username that will be passed when authenticating SIP requests. The username should be sent in response to Twilio's challenge of the initial INVITE. It can be up to 32 characters long.
	Username *string `json:"Username,omitempty"`
}

func (params *CreateSipCredentialParams) SetPathAccountSid(PathAccountSid string) *CreateSipCredentialParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *CreateSipCredentialParams) SetPassword(Password string) *CreateSipCredentialParams {
	params.Password = &Password
	return params
}
func (params *CreateSipCredentialParams) SetUsername(Username string) *CreateSipCredentialParams {
	params.Username = &Username
	return params
}

// Create a new credential resource.
func (c *ApiService) CreateSipCredential(CredentialListSid string, params *CreateSipCredentialParams) (*ApiV2010AccountSipSipCredentialListSipCredential, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/SIP/CredentialLists/{CredentialListSid}/Credentials.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"CredentialListSid"+"}", CredentialListSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Password != nil {
		data.Set("Password", *params.Password)
	}
	if params != nil && params.Username != nil {
		data.Set("Username", *params.Username)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010AccountSipSipCredentialListSipCredential{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'DeleteSipCredential'
type DeleteSipCredentialParams struct {
	// The unique id of the Account that is responsible for this resource.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *DeleteSipCredentialParams) SetPathAccountSid(PathAccountSid string) *DeleteSipCredentialParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

// Delete a credential resource.
func (c *ApiService) DeleteSipCredential(CredentialListSid string, Sid string, params *DeleteSipCredentialParams) error {
	path := "/2010-04-01/Accounts/{AccountSid}/SIP/CredentialLists/{CredentialListSid}/Credentials/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"CredentialListSid"+"}", CredentialListSid, -1)
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

// Optional parameters for the method 'FetchSipCredential'
type FetchSipCredentialParams struct {
	// The unique id of the Account that is responsible for this resource.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *FetchSipCredentialParams) SetPathAccountSid(PathAccountSid string) *FetchSipCredentialParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

// Fetch a single credential.
func (c *ApiService) FetchSipCredential(CredentialListSid string, Sid string, params *FetchSipCredentialParams) (*ApiV2010AccountSipSipCredentialListSipCredential, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/SIP/CredentialLists/{CredentialListSid}/Credentials/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"CredentialListSid"+"}", CredentialListSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010AccountSipSipCredentialListSipCredential{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListSipCredential'
type ListSipCredentialParams struct {
	// The unique id of the Account that is responsible for this resource.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListSipCredentialParams) SetPathAccountSid(PathAccountSid string) *ListSipCredentialParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *ListSipCredentialParams) SetPageSize(PageSize int) *ListSipCredentialParams {
	params.PageSize = &PageSize
	return params
}

// Retrieve a single page of SipCredential records from the API. Request is executed immediately.
func (c *ApiService) PageSipCredential(CredentialListSid string, params *ListSipCredentialParams, pageToken string, pageNumber string) (*ListSipCredentialResponse, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/SIP/CredentialLists/{CredentialListSid}/Credentials.json"

	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"CredentialListSid"+"}", CredentialListSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	if pageToken != "" {
		data.Set("PageToken", pageToken)
	}
	if pageToken != "" {
		data.Set("Page", pageNumber)
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListSipCredentialResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists SipCredential records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListSipCredential(CredentialListSid string, params *ListSipCredentialParams, limit int) ([]ApiV2010AccountSipSipCredentialListSipCredential, error) {
	if params == nil {
		params = &ListSipCredentialParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, limit))

	response, err := c.PageSipCredential(CredentialListSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []ApiV2010AccountSipSipCredentialListSipCredential

	for response != nil {
		records = append(records, response.Credentials...)

		var record interface{}
		if record, err = client.GetNext(response, &curRecord, limit, c.getNextListSipCredentialResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListSipCredentialResponse)
	}

	return records, err
}

// Streams SipCredential records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamSipCredential(CredentialListSid string, params *ListSipCredentialParams, limit int) (chan ApiV2010AccountSipSipCredentialListSipCredential, error) {
	if params == nil {
		params = &ListSipCredentialParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, limit))

	response, err := c.PageSipCredential(CredentialListSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan ApiV2010AccountSipSipCredentialListSipCredential, 1)

	go func() {
		for response != nil {
			for item := range response.Credentials {
				channel <- response.Credentials[item]
			}

			var record interface{}
			if record, err = client.GetNext(response, &curRecord, limit, c.getNextListSipCredentialResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListSipCredentialResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListSipCredentialResponse(nextPageUri string) (interface{}, error) {
	if nextPageUri == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(c.baseURL+nextPageUri, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListSipCredentialResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateSipCredential'
type UpdateSipCredentialParams struct {
	// The unique id of the Account that is responsible for this resource.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// The password that the username will use when authenticating SIP requests. The password must be a minimum of 12 characters, contain at least 1 digit, and have mixed case. (eg `IWasAtSignal2018`)
	Password *string `json:"Password,omitempty"`
}

func (params *UpdateSipCredentialParams) SetPathAccountSid(PathAccountSid string) *UpdateSipCredentialParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *UpdateSipCredentialParams) SetPassword(Password string) *UpdateSipCredentialParams {
	params.Password = &Password
	return params
}

// Update a credential resource.
func (c *ApiService) UpdateSipCredential(CredentialListSid string, Sid string, params *UpdateSipCredentialParams) (*ApiV2010AccountSipSipCredentialListSipCredential, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/SIP/CredentialLists/{CredentialListSid}/Credentials/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"CredentialListSid"+"}", CredentialListSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Password != nil {
		data.Set("Password", *params.Password)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010AccountSipSipCredentialListSipCredential{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
