/*
 * Twilio - Trusthub
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

// Optional parameters for the method 'CreateTrustProductEntityAssignment'
type CreateTrustProductEntityAssignmentParams struct {
	// The SID of an object bag that holds information of the different items.
	ObjectSid *string `json:"ObjectSid,omitempty"`
}

func (params *CreateTrustProductEntityAssignmentParams) SetObjectSid(ObjectSid string) *CreateTrustProductEntityAssignmentParams {
	params.ObjectSid = &ObjectSid
	return params
}

// Create a new Assigned Item.
func (c *ApiService) CreateTrustProductEntityAssignment(TrustProductSid string, params *CreateTrustProductEntityAssignmentParams) (*TrusthubV1TrustProductTrustProductEntityAssignment, error) {
	path := "/v1/TrustProducts/{TrustProductSid}/EntityAssignments"
	path = strings.Replace(path, "{"+"TrustProductSid"+"}", TrustProductSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.ObjectSid != nil {
		data.Set("ObjectSid", *params.ObjectSid)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TrusthubV1TrustProductTrustProductEntityAssignment{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Remove an Assignment Item Instance.
func (c *ApiService) DeleteTrustProductEntityAssignment(TrustProductSid string, Sid string) error {
	path := "/v1/TrustProducts/{TrustProductSid}/EntityAssignments/{Sid}"
	path = strings.Replace(path, "{"+"TrustProductSid"+"}", TrustProductSid, -1)
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

// Fetch specific Assigned Item Instance.
func (c *ApiService) FetchTrustProductEntityAssignment(TrustProductSid string, Sid string) (*TrusthubV1TrustProductTrustProductEntityAssignment, error) {
	path := "/v1/TrustProducts/{TrustProductSid}/EntityAssignments/{Sid}"
	path = strings.Replace(path, "{"+"TrustProductSid"+"}", TrustProductSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TrusthubV1TrustProductTrustProductEntityAssignment{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListTrustProductEntityAssignment'
type ListTrustProductEntityAssignmentParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListTrustProductEntityAssignmentParams) SetPageSize(PageSize int) *ListTrustProductEntityAssignmentParams {
	params.PageSize = &PageSize
	return params
}

// Retrieve a single page of TrustProductEntityAssignment records from the API. Request is executed immediately.
func (c *ApiService) PageTrustProductEntityAssignment(TrustProductSid string, params *ListTrustProductEntityAssignmentParams, pageToken string, pageNumber string) (*ListTrustProductEntityAssignmentResponse, error) {
	path := "/v1/TrustProducts/{TrustProductSid}/EntityAssignments"

	path = strings.Replace(path, "{"+"TrustProductSid"+"}", TrustProductSid, -1)

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

	ps := &ListTrustProductEntityAssignmentResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists TrustProductEntityAssignment records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListTrustProductEntityAssignment(TrustProductSid string, params *ListTrustProductEntityAssignmentParams, limit int) ([]TrusthubV1TrustProductTrustProductEntityAssignment, error) {
	if params == nil {
		params = &ListTrustProductEntityAssignmentParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, limit))

	response, err := c.PageTrustProductEntityAssignment(TrustProductSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []TrusthubV1TrustProductTrustProductEntityAssignment

	for response != nil {
		records = append(records, response.Results...)

		var record interface{}
		if record, err = client.GetNext(response, &curRecord, limit, c.getNextListTrustProductEntityAssignmentResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListTrustProductEntityAssignmentResponse)
	}

	return records, err
}

// Streams TrustProductEntityAssignment records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamTrustProductEntityAssignment(TrustProductSid string, params *ListTrustProductEntityAssignmentParams, limit int) (chan TrusthubV1TrustProductTrustProductEntityAssignment, error) {
	if params == nil {
		params = &ListTrustProductEntityAssignmentParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, limit))

	response, err := c.PageTrustProductEntityAssignment(TrustProductSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan TrusthubV1TrustProductTrustProductEntityAssignment, 1)

	go func() {
		for response != nil {
			for item := range response.Results {
				channel <- response.Results[item]
			}

			var record interface{}
			if record, err = client.GetNext(response, &curRecord, limit, c.getNextListTrustProductEntityAssignmentResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListTrustProductEntityAssignmentResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListTrustProductEntityAssignmentResponse(nextPageUri string) (interface{}, error) {
	if nextPageUri == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(c.baseURL+nextPageUri, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListTrustProductEntityAssignmentResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
