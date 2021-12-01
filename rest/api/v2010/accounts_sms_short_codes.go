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

// Optional parameters for the method 'FetchShortCode'
type FetchShortCodeParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the ShortCode resource(s) to fetch.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *FetchShortCodeParams) SetPathAccountSid(PathAccountSid string) *FetchShortCodeParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

// Fetch an instance of a short code
func (c *ApiService) FetchShortCode(Sid string, params *FetchShortCodeParams) (*ApiV2010ShortCode, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/SMS/ShortCodes/{Sid}.json"
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

	ps := &ApiV2010ShortCode{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListShortCode'
type ListShortCodeParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the ShortCode resource(s) to read.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// The string that identifies the ShortCode resources to read.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// Only show the ShortCode resources that match this pattern. You can specify partial numbers and use '*' as a wildcard for any digit.
	ShortCode *string `json:"ShortCode,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListShortCodeParams) SetPathAccountSid(PathAccountSid string) *ListShortCodeParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *ListShortCodeParams) SetFriendlyName(FriendlyName string) *ListShortCodeParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *ListShortCodeParams) SetShortCode(ShortCode string) *ListShortCodeParams {
	params.ShortCode = &ShortCode
	return params
}
func (params *ListShortCodeParams) SetPageSize(PageSize int) *ListShortCodeParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListShortCodeParams) SetLimit(Limit int) *ListShortCodeParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of ShortCode records from the API. Request is executed immediately.
func (c *ApiService) PageShortCode(params *ListShortCodeParams, pageToken, pageNumber string) (*ListShortCodeResponse, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/SMS/ShortCodes.json"

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
	if params != nil && params.ShortCode != nil {
		data.Set("ShortCode", *params.ShortCode)
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

	ps := &ListShortCodeResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists ShortCode records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListShortCode(params *ListShortCodeParams) ([]ApiV2010ShortCode, error) {
	if params == nil {
		params = &ListShortCodeParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageShortCode(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []ApiV2010ShortCode

	for response != nil {
		records = append(records, response.ShortCodes...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListShortCodeResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListShortCodeResponse)
	}

	return records, err
}

// Streams ShortCode records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamShortCode(params *ListShortCodeParams) (chan ApiV2010ShortCode, error) {
	if params == nil {
		params = &ListShortCodeParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageShortCode(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan ApiV2010ShortCode, 1)

	go func() {
		for response != nil {
			for item := range response.ShortCodes {
				channel <- response.ShortCodes[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListShortCodeResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListShortCodeResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListShortCodeResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListShortCodeResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateShortCode'
type UpdateShortCodeParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the ShortCode resource(s) to update.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// The API version to use to start a new TwiML session. Can be: `2010-04-01` or `2008-08-01`.
	ApiVersion *string `json:"ApiVersion,omitempty"`
	// A descriptive string that you created to describe this resource. It can be up to 64 characters long. By default, the `FriendlyName` is the short code.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// The HTTP method that we should use to call the `sms_fallback_url`. Can be: `GET` or `POST`.
	SmsFallbackMethod *string `json:"SmsFallbackMethod,omitempty"`
	// The URL that we should call if an error occurs while retrieving or executing the TwiML from `sms_url`.
	SmsFallbackUrl *string `json:"SmsFallbackUrl,omitempty"`
	// The HTTP method we should use when calling the `sms_url`. Can be: `GET` or `POST`.
	SmsMethod *string `json:"SmsMethod,omitempty"`
	// The URL we should call when receiving an incoming SMS message to this short code.
	SmsUrl *string `json:"SmsUrl,omitempty"`
}

func (params *UpdateShortCodeParams) SetPathAccountSid(PathAccountSid string) *UpdateShortCodeParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *UpdateShortCodeParams) SetApiVersion(ApiVersion string) *UpdateShortCodeParams {
	params.ApiVersion = &ApiVersion
	return params
}
func (params *UpdateShortCodeParams) SetFriendlyName(FriendlyName string) *UpdateShortCodeParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *UpdateShortCodeParams) SetSmsFallbackMethod(SmsFallbackMethod string) *UpdateShortCodeParams {
	params.SmsFallbackMethod = &SmsFallbackMethod
	return params
}
func (params *UpdateShortCodeParams) SetSmsFallbackUrl(SmsFallbackUrl string) *UpdateShortCodeParams {
	params.SmsFallbackUrl = &SmsFallbackUrl
	return params
}
func (params *UpdateShortCodeParams) SetSmsMethod(SmsMethod string) *UpdateShortCodeParams {
	params.SmsMethod = &SmsMethod
	return params
}
func (params *UpdateShortCodeParams) SetSmsUrl(SmsUrl string) *UpdateShortCodeParams {
	params.SmsUrl = &SmsUrl
	return params
}

// Update a short code with the following parameters
func (c *ApiService) UpdateShortCode(Sid string, params *UpdateShortCodeParams) (*ApiV2010ShortCode, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/SMS/ShortCodes/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.ApiVersion != nil {
		data.Set("ApiVersion", *params.ApiVersion)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.SmsFallbackMethod != nil {
		data.Set("SmsFallbackMethod", *params.SmsFallbackMethod)
	}
	if params != nil && params.SmsFallbackUrl != nil {
		data.Set("SmsFallbackUrl", *params.SmsFallbackUrl)
	}
	if params != nil && params.SmsMethod != nil {
		data.Set("SmsMethod", *params.SmsMethod)
	}
	if params != nil && params.SmsUrl != nil {
		data.Set("SmsUrl", *params.SmsUrl)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010ShortCode{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
