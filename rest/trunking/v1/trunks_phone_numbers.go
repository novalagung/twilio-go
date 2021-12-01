/*
 * Twilio - Trunking
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

// Optional parameters for the method 'CreatePhoneNumber'
type CreatePhoneNumberParams struct {
	// The SID of the [Incoming Phone Number](https://www.twilio.com/docs/phone-numbers/api/incomingphonenumber-resource) that you want to associate with the trunk.
	PhoneNumberSid *string `json:"PhoneNumberSid,omitempty"`
}

func (params *CreatePhoneNumberParams) SetPhoneNumberSid(PhoneNumberSid string) *CreatePhoneNumberParams {
	params.PhoneNumberSid = &PhoneNumberSid
	return params
}

func (c *ApiService) CreatePhoneNumber(TrunkSid string, params *CreatePhoneNumberParams) (*TrunkingV1PhoneNumber, error) {
	path := "/v1/Trunks/{TrunkSid}/PhoneNumbers"
	path = strings.Replace(path, "{"+"TrunkSid"+"}", TrunkSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PhoneNumberSid != nil {
		data.Set("PhoneNumberSid", *params.PhoneNumberSid)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TrunkingV1PhoneNumber{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

func (c *ApiService) DeletePhoneNumber(TrunkSid string, Sid string) error {
	path := "/v1/Trunks/{TrunkSid}/PhoneNumbers/{Sid}"
	path = strings.Replace(path, "{"+"TrunkSid"+"}", TrunkSid, -1)
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

func (c *ApiService) FetchPhoneNumber(TrunkSid string, Sid string) (*TrunkingV1PhoneNumber, error) {
	path := "/v1/Trunks/{TrunkSid}/PhoneNumbers/{Sid}"
	path = strings.Replace(path, "{"+"TrunkSid"+"}", TrunkSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TrunkingV1PhoneNumber{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListPhoneNumber'
type ListPhoneNumberParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListPhoneNumberParams) SetPageSize(PageSize int) *ListPhoneNumberParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListPhoneNumberParams) SetLimit(Limit int) *ListPhoneNumberParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of PhoneNumber records from the API. Request is executed immediately.
func (c *ApiService) PagePhoneNumber(TrunkSid string, params *ListPhoneNumberParams, pageToken, pageNumber string) (*ListPhoneNumberResponse, error) {
	path := "/v1/Trunks/{TrunkSid}/PhoneNumbers"

	path = strings.Replace(path, "{"+"TrunkSid"+"}", TrunkSid, -1)

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

	ps := &ListPhoneNumberResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists PhoneNumber records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListPhoneNumber(TrunkSid string, params *ListPhoneNumberParams) ([]TrunkingV1PhoneNumber, error) {
	if params == nil {
		params = &ListPhoneNumberParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PagePhoneNumber(TrunkSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []TrunkingV1PhoneNumber

	for response != nil {
		records = append(records, response.PhoneNumbers...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListPhoneNumberResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListPhoneNumberResponse)
	}

	return records, err
}

// Streams PhoneNumber records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamPhoneNumber(TrunkSid string, params *ListPhoneNumberParams) (chan TrunkingV1PhoneNumber, error) {
	if params == nil {
		params = &ListPhoneNumberParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PagePhoneNumber(TrunkSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan TrunkingV1PhoneNumber, 1)

	go func() {
		for response != nil {
			for item := range response.PhoneNumbers {
				channel <- response.PhoneNumbers[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListPhoneNumberResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListPhoneNumberResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListPhoneNumberResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListPhoneNumberResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
