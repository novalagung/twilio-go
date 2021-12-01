/*
 * Twilio - Trusthub
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

// Optional parameters for the method 'CreateCustomerProfile'
type CreateCustomerProfileParams struct {
	// The email address that will receive updates when the Customer-Profile resource changes status.
	Email *string `json:"Email,omitempty"`
	// The string that you assigned to describe the resource.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// The unique string of a policy that is associated to the Customer-Profile resource.
	PolicySid *string `json:"PolicySid,omitempty"`
	// The URL we call to inform your application of status changes.
	StatusCallback *string `json:"StatusCallback,omitempty"`
}

func (params *CreateCustomerProfileParams) SetEmail(Email string) *CreateCustomerProfileParams {
	params.Email = &Email
	return params
}
func (params *CreateCustomerProfileParams) SetFriendlyName(FriendlyName string) *CreateCustomerProfileParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *CreateCustomerProfileParams) SetPolicySid(PolicySid string) *CreateCustomerProfileParams {
	params.PolicySid = &PolicySid
	return params
}
func (params *CreateCustomerProfileParams) SetStatusCallback(StatusCallback string) *CreateCustomerProfileParams {
	params.StatusCallback = &StatusCallback
	return params
}

// Create a new Customer-Profile.
func (c *ApiService) CreateCustomerProfile(params *CreateCustomerProfileParams) (*TrusthubV1CustomerProfile, error) {
	path := "/v1/CustomerProfiles"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Email != nil {
		data.Set("Email", *params.Email)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.PolicySid != nil {
		data.Set("PolicySid", *params.PolicySid)
	}
	if params != nil && params.StatusCallback != nil {
		data.Set("StatusCallback", *params.StatusCallback)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TrusthubV1CustomerProfile{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Delete a specific Customer-Profile.
func (c *ApiService) DeleteCustomerProfile(Sid string) error {
	path := "/v1/CustomerProfiles/{Sid}"
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

// Fetch a specific Customer-Profile instance.
func (c *ApiService) FetchCustomerProfile(Sid string) (*TrusthubV1CustomerProfile, error) {
	path := "/v1/CustomerProfiles/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TrusthubV1CustomerProfile{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListCustomerProfile'
type ListCustomerProfileParams struct {
	// The verification status of the Customer-Profile resource.
	Status *string `json:"Status,omitempty"`
	// The string that you assigned to describe the resource.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// The unique string of a policy that is associated to the Customer-Profile resource.
	PolicySid *string `json:"PolicySid,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListCustomerProfileParams) SetStatus(Status string) *ListCustomerProfileParams {
	params.Status = &Status
	return params
}
func (params *ListCustomerProfileParams) SetFriendlyName(FriendlyName string) *ListCustomerProfileParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *ListCustomerProfileParams) SetPolicySid(PolicySid string) *ListCustomerProfileParams {
	params.PolicySid = &PolicySid
	return params
}
func (params *ListCustomerProfileParams) SetPageSize(PageSize int) *ListCustomerProfileParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListCustomerProfileParams) SetLimit(Limit int) *ListCustomerProfileParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of CustomerProfile records from the API. Request is executed immediately.
func (c *ApiService) PageCustomerProfile(params *ListCustomerProfileParams, pageToken, pageNumber string) (*ListCustomerProfileResponse, error) {
	path := "/v1/CustomerProfiles"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Status != nil {
		data.Set("Status", *params.Status)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.PolicySid != nil {
		data.Set("PolicySid", *params.PolicySid)
	}
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

	ps := &ListCustomerProfileResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists CustomerProfile records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListCustomerProfile(params *ListCustomerProfileParams) ([]TrusthubV1CustomerProfile, error) {
	if params == nil {
		params = &ListCustomerProfileParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageCustomerProfile(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []TrusthubV1CustomerProfile

	for response != nil {
		records = append(records, response.Results...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListCustomerProfileResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListCustomerProfileResponse)
	}

	return records, err
}

// Streams CustomerProfile records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamCustomerProfile(params *ListCustomerProfileParams) (chan TrusthubV1CustomerProfile, error) {
	if params == nil {
		params = &ListCustomerProfileParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageCustomerProfile(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan TrusthubV1CustomerProfile, 1)

	go func() {
		for response != nil {
			for item := range response.Results {
				channel <- response.Results[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListCustomerProfileResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListCustomerProfileResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListCustomerProfileResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListCustomerProfileResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateCustomerProfile'
type UpdateCustomerProfileParams struct {
	// The email address that will receive updates when the Customer-Profile resource changes status.
	Email *string `json:"Email,omitempty"`
	// The string that you assigned to describe the resource.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// The verification status of the Customer-Profile resource.
	Status *string `json:"Status,omitempty"`
	// The URL we call to inform your application of status changes.
	StatusCallback *string `json:"StatusCallback,omitempty"`
}

func (params *UpdateCustomerProfileParams) SetEmail(Email string) *UpdateCustomerProfileParams {
	params.Email = &Email
	return params
}
func (params *UpdateCustomerProfileParams) SetFriendlyName(FriendlyName string) *UpdateCustomerProfileParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *UpdateCustomerProfileParams) SetStatus(Status string) *UpdateCustomerProfileParams {
	params.Status = &Status
	return params
}
func (params *UpdateCustomerProfileParams) SetStatusCallback(StatusCallback string) *UpdateCustomerProfileParams {
	params.StatusCallback = &StatusCallback
	return params
}

// Updates a Customer-Profile in an account.
func (c *ApiService) UpdateCustomerProfile(Sid string, params *UpdateCustomerProfileParams) (*TrusthubV1CustomerProfile, error) {
	path := "/v1/CustomerProfiles/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Email != nil {
		data.Set("Email", *params.Email)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.Status != nil {
		data.Set("Status", *params.Status)
	}
	if params != nil && params.StatusCallback != nil {
		data.Set("StatusCallback", *params.StatusCallback)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TrusthubV1CustomerProfile{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
