/*
 * Twilio - Ip_messaging
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

func (c *ApiService) DeleteUserBinding(ServiceSid string, UserSid string, Sid string) error {
	path := "/v2/Services/{ServiceSid}/Users/{UserSid}/Bindings/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"UserSid"+"}", UserSid, -1)
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

func (c *ApiService) FetchUserBinding(ServiceSid string, UserSid string, Sid string) (*IpMessagingV2ServiceUserUserBinding, error) {
	path := "/v2/Services/{ServiceSid}/Users/{UserSid}/Bindings/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"UserSid"+"}", UserSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &IpMessagingV2ServiceUserUserBinding{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListUserBinding'
type ListUserBindingParams struct {
	//
	BindingType *[]string `json:"BindingType,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListUserBindingParams) SetBindingType(BindingType []string) *ListUserBindingParams {
	params.BindingType = &BindingType
	return params
}
func (params *ListUserBindingParams) SetPageSize(PageSize int) *ListUserBindingParams {
	params.PageSize = &PageSize
	return params
}

// Retrieve a single page of UserBinding records from the API. Request is executed immediately.
func (c *ApiService) PageUserBinding(ServiceSid string, UserSid string, params *ListUserBindingParams, pageToken string, pageNumber string) (*ListUserBindingResponse, error) {
	path := "/v2/Services/{ServiceSid}/Users/{UserSid}/Bindings"

	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"UserSid"+"}", UserSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.BindingType != nil {
		for _, item := range *params.BindingType {
			data.Add("BindingType", item)
		}
	}
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

	ps := &ListUserBindingResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists UserBinding records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListUserBinding(ServiceSid string, UserSid string, params *ListUserBindingParams, limit int) ([]IpMessagingV2ServiceUserUserBinding, error) {
	if params == nil {
		params = &ListUserBindingParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, limit))

	response, err := c.PageUserBinding(ServiceSid, UserSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []IpMessagingV2ServiceUserUserBinding

	for response != nil {
		records = append(records, response.Bindings...)

		var record interface{}
		if record, err = client.GetNext(response, &curRecord, limit, c.getNextListUserBindingResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListUserBindingResponse)
	}

	return records, err
}

// Streams UserBinding records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamUserBinding(ServiceSid string, UserSid string, params *ListUserBindingParams, limit int) (chan IpMessagingV2ServiceUserUserBinding, error) {
	if params == nil {
		params = &ListUserBindingParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, limit))

	response, err := c.PageUserBinding(ServiceSid, UserSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan IpMessagingV2ServiceUserUserBinding, 1)

	go func() {
		for response != nil {
			for item := range response.Bindings {
				channel <- response.Bindings[item]
			}

			var record interface{}
			if record, err = client.GetNext(response, &curRecord, limit, c.getNextListUserBindingResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListUserBindingResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListUserBindingResponse(nextPageUri string) (interface{}, error) {
	if nextPageUri == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(c.baseURL+nextPageUri, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListUserBindingResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
