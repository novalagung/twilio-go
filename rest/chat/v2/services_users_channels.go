/*
 * Twilio - Chat
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
	"time"

	"github.com/twilio/twilio-go/client"
)

// Removes User from selected Channel.
func (c *ApiService) DeleteUserChannel(ServiceSid string, UserSid string, ChannelSid string) error {
	path := "/v2/Services/{ServiceSid}/Users/{UserSid}/Channels/{ChannelSid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"UserSid"+"}", UserSid, -1)
	path = strings.Replace(path, "{"+"ChannelSid"+"}", ChannelSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

func (c *ApiService) FetchUserChannel(ServiceSid string, UserSid string, ChannelSid string) (*ChatV2ServiceUserUserChannel, error) {
	path := "/v2/Services/{ServiceSid}/Users/{UserSid}/Channels/{ChannelSid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"UserSid"+"}", UserSid, -1)
	path = strings.Replace(path, "{"+"ChannelSid"+"}", ChannelSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ChatV2ServiceUserUserChannel{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListUserChannel'
type ListUserChannelParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListUserChannelParams) SetPageSize(PageSize int) *ListUserChannelParams {
	params.PageSize = &PageSize
	return params
}

// Retrieve a single page of UserChannel records from the API. Request is executed immediately.
func (c *ApiService) PageUserChannel(ServiceSid string, UserSid string, params *ListUserChannelParams, pageToken string, pageNumber string) (*ListUserChannelResponse, error) {
	path := "/v2/Services/{ServiceSid}/Users/{UserSid}/Channels"

	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"UserSid"+"}", UserSid, -1)

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

	ps := &ListUserChannelResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists UserChannel records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListUserChannel(ServiceSid string, UserSid string, params *ListUserChannelParams, limit int) ([]ChatV2ServiceUserUserChannel, error) {
	if params == nil {
		params = &ListUserChannelParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, limit))

	response, err := c.PageUserChannel(ServiceSid, UserSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []ChatV2ServiceUserUserChannel

	for response != nil {
		records = append(records, response.Channels...)

		var record interface{}
		if record, err = client.GetNext(response, &curRecord, limit, c.getNextListUserChannelResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListUserChannelResponse)
	}

	return records, err
}

// Streams UserChannel records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamUserChannel(ServiceSid string, UserSid string, params *ListUserChannelParams, limit int) (chan ChatV2ServiceUserUserChannel, error) {
	if params == nil {
		params = &ListUserChannelParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, limit))

	response, err := c.PageUserChannel(ServiceSid, UserSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan ChatV2ServiceUserUserChannel, 1)

	go func() {
		for response != nil {
			for item := range response.Channels {
				channel <- response.Channels[item]
			}

			var record interface{}
			if record, err = client.GetNext(response, &curRecord, limit, c.getNextListUserChannelResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListUserChannelResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListUserChannelResponse(nextPageUri string) (interface{}, error) {
	if nextPageUri == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(c.baseURL+nextPageUri, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListUserChannelResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateUserChannel'
type UpdateUserChannelParams struct {
	// The index of the last [Message](https://www.twilio.com/docs/chat/rest/message-resource) in the [Channel](https://www.twilio.com/docs/chat/channels) that the Member has read.
	LastConsumedMessageIndex *int `json:"LastConsumedMessageIndex,omitempty"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp of the last [Message](https://www.twilio.com/docs/chat/rest/message-resource) read event for the Member within the [Channel](https://www.twilio.com/docs/chat/channels).
	LastConsumptionTimestamp *time.Time `json:"LastConsumptionTimestamp,omitempty"`
	// The push notification level to assign to the User Channel. Can be: `default` or `muted`.
	NotificationLevel *string `json:"NotificationLevel,omitempty"`
}

func (params *UpdateUserChannelParams) SetLastConsumedMessageIndex(LastConsumedMessageIndex int) *UpdateUserChannelParams {
	params.LastConsumedMessageIndex = &LastConsumedMessageIndex
	return params
}
func (params *UpdateUserChannelParams) SetLastConsumptionTimestamp(LastConsumptionTimestamp time.Time) *UpdateUserChannelParams {
	params.LastConsumptionTimestamp = &LastConsumptionTimestamp
	return params
}
func (params *UpdateUserChannelParams) SetNotificationLevel(NotificationLevel string) *UpdateUserChannelParams {
	params.NotificationLevel = &NotificationLevel
	return params
}

func (c *ApiService) UpdateUserChannel(ServiceSid string, UserSid string, ChannelSid string, params *UpdateUserChannelParams) (*ChatV2ServiceUserUserChannel, error) {
	path := "/v2/Services/{ServiceSid}/Users/{UserSid}/Channels/{ChannelSid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"UserSid"+"}", UserSid, -1)
	path = strings.Replace(path, "{"+"ChannelSid"+"}", ChannelSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.LastConsumedMessageIndex != nil {
		data.Set("LastConsumedMessageIndex", fmt.Sprint(*params.LastConsumedMessageIndex))
	}
	if params != nil && params.LastConsumptionTimestamp != nil {
		data.Set("LastConsumptionTimestamp", fmt.Sprint((*params.LastConsumptionTimestamp).Format(time.RFC3339)))
	}
	if params != nil && params.NotificationLevel != nil {
		data.Set("NotificationLevel", *params.NotificationLevel)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ChatV2ServiceUserUserChannel{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
