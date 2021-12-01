/*
 * Twilio - Conversations
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
	"time"

	"github.com/twilio/twilio-go/client"
)

// Optional parameters for the method 'CreateServiceConversationMessage'
type CreateServiceConversationMessageParams struct {
	// The X-Twilio-Webhook-Enabled HTTP request header
	XTwilioWebhookEnabled *string `json:"X-Twilio-Webhook-Enabled,omitempty"`
	// A string metadata field you can use to store any data you wish. The string value must contain structurally valid JSON if specified.  **Note** that if the attributes are not set \\\"{}\\\" will be returned.
	Attributes *string `json:"Attributes,omitempty"`
	// The channel specific identifier of the message's author. Defaults to `system`.
	Author *string `json:"Author,omitempty"`
	// The content of the message, can be up to 1,600 characters long.
	Body *string `json:"Body,omitempty"`
	// The date that this resource was created.
	DateCreated *time.Time `json:"DateCreated,omitempty"`
	// The date that this resource was last updated. `null` if the message has not been edited.
	DateUpdated *time.Time `json:"DateUpdated,omitempty"`
	// The Media SID to be attached to the new Message.
	MediaSid *string `json:"MediaSid,omitempty"`
}

func (params *CreateServiceConversationMessageParams) SetXTwilioWebhookEnabled(XTwilioWebhookEnabled string) *CreateServiceConversationMessageParams {
	params.XTwilioWebhookEnabled = &XTwilioWebhookEnabled
	return params
}
func (params *CreateServiceConversationMessageParams) SetAttributes(Attributes string) *CreateServiceConversationMessageParams {
	params.Attributes = &Attributes
	return params
}
func (params *CreateServiceConversationMessageParams) SetAuthor(Author string) *CreateServiceConversationMessageParams {
	params.Author = &Author
	return params
}
func (params *CreateServiceConversationMessageParams) SetBody(Body string) *CreateServiceConversationMessageParams {
	params.Body = &Body
	return params
}
func (params *CreateServiceConversationMessageParams) SetDateCreated(DateCreated time.Time) *CreateServiceConversationMessageParams {
	params.DateCreated = &DateCreated
	return params
}
func (params *CreateServiceConversationMessageParams) SetDateUpdated(DateUpdated time.Time) *CreateServiceConversationMessageParams {
	params.DateUpdated = &DateUpdated
	return params
}
func (params *CreateServiceConversationMessageParams) SetMediaSid(MediaSid string) *CreateServiceConversationMessageParams {
	params.MediaSid = &MediaSid
	return params
}

// Add a new message to the conversation in a specific service
func (c *ApiService) CreateServiceConversationMessage(ChatServiceSid string, ConversationSid string, params *CreateServiceConversationMessageParams) (*ConversationsV1ServiceConversationMessage, error) {
	path := "/v1/Services/{ChatServiceSid}/Conversations/{ConversationSid}/Messages"
	path = strings.Replace(path, "{"+"ChatServiceSid"+"}", ChatServiceSid, -1)
	path = strings.Replace(path, "{"+"ConversationSid"+"}", ConversationSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Attributes != nil {
		data.Set("Attributes", *params.Attributes)
	}
	if params != nil && params.Author != nil {
		data.Set("Author", *params.Author)
	}
	if params != nil && params.Body != nil {
		data.Set("Body", *params.Body)
	}
	if params != nil && params.DateCreated != nil {
		data.Set("DateCreated", fmt.Sprint((*params.DateCreated).Format(time.RFC3339)))
	}
	if params != nil && params.DateUpdated != nil {
		data.Set("DateUpdated", fmt.Sprint((*params.DateUpdated).Format(time.RFC3339)))
	}
	if params != nil && params.MediaSid != nil {
		data.Set("MediaSid", *params.MediaSid)
	}

	if params != nil && params.XTwilioWebhookEnabled != nil {
		headers["X-Twilio-Webhook-Enabled"] = *params.XTwilioWebhookEnabled
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ConversationsV1ServiceConversationMessage{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'DeleteServiceConversationMessage'
type DeleteServiceConversationMessageParams struct {
	// The X-Twilio-Webhook-Enabled HTTP request header
	XTwilioWebhookEnabled *string `json:"X-Twilio-Webhook-Enabled,omitempty"`
}

func (params *DeleteServiceConversationMessageParams) SetXTwilioWebhookEnabled(XTwilioWebhookEnabled string) *DeleteServiceConversationMessageParams {
	params.XTwilioWebhookEnabled = &XTwilioWebhookEnabled
	return params
}

// Remove a message from the conversation
func (c *ApiService) DeleteServiceConversationMessage(ChatServiceSid string, ConversationSid string, Sid string, params *DeleteServiceConversationMessageParams) error {
	path := "/v1/Services/{ChatServiceSid}/Conversations/{ConversationSid}/Messages/{Sid}"
	path = strings.Replace(path, "{"+"ChatServiceSid"+"}", ChatServiceSid, -1)
	path = strings.Replace(path, "{"+"ConversationSid"+"}", ConversationSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.XTwilioWebhookEnabled != nil {
		headers["X-Twilio-Webhook-Enabled"] = *params.XTwilioWebhookEnabled
	}

	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

// Fetch a message from the conversation
func (c *ApiService) FetchServiceConversationMessage(ChatServiceSid string, ConversationSid string, Sid string) (*ConversationsV1ServiceConversationMessage, error) {
	path := "/v1/Services/{ChatServiceSid}/Conversations/{ConversationSid}/Messages/{Sid}"
	path = strings.Replace(path, "{"+"ChatServiceSid"+"}", ChatServiceSid, -1)
	path = strings.Replace(path, "{"+"ConversationSid"+"}", ConversationSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ConversationsV1ServiceConversationMessage{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListServiceConversationMessage'
type ListServiceConversationMessageParams struct {
	// The sort order of the returned messages. Can be: `asc` (ascending) or `desc` (descending), with `asc` as the default.
	Order *string `json:"Order,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListServiceConversationMessageParams) SetOrder(Order string) *ListServiceConversationMessageParams {
	params.Order = &Order
	return params
}
func (params *ListServiceConversationMessageParams) SetPageSize(PageSize int) *ListServiceConversationMessageParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListServiceConversationMessageParams) SetLimit(Limit int) *ListServiceConversationMessageParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of ServiceConversationMessage records from the API. Request is executed immediately.
func (c *ApiService) PageServiceConversationMessage(ChatServiceSid string, ConversationSid string, params *ListServiceConversationMessageParams, pageToken, pageNumber string) (*ListServiceConversationMessageResponse, error) {
	path := "/v1/Services/{ChatServiceSid}/Conversations/{ConversationSid}/Messages"

	path = strings.Replace(path, "{"+"ChatServiceSid"+"}", ChatServiceSid, -1)
	path = strings.Replace(path, "{"+"ConversationSid"+"}", ConversationSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Order != nil {
		data.Set("Order", *params.Order)
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

	ps := &ListServiceConversationMessageResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists ServiceConversationMessage records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListServiceConversationMessage(ChatServiceSid string, ConversationSid string, params *ListServiceConversationMessageParams) ([]ConversationsV1ServiceConversationMessage, error) {
	if params == nil {
		params = &ListServiceConversationMessageParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageServiceConversationMessage(ChatServiceSid, ConversationSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []ConversationsV1ServiceConversationMessage

	for response != nil {
		records = append(records, response.Messages...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListServiceConversationMessageResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListServiceConversationMessageResponse)
	}

	return records, err
}

// Streams ServiceConversationMessage records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamServiceConversationMessage(ChatServiceSid string, ConversationSid string, params *ListServiceConversationMessageParams) (chan ConversationsV1ServiceConversationMessage, error) {
	if params == nil {
		params = &ListServiceConversationMessageParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageServiceConversationMessage(ChatServiceSid, ConversationSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan ConversationsV1ServiceConversationMessage, 1)

	go func() {
		for response != nil {
			for item := range response.Messages {
				channel <- response.Messages[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListServiceConversationMessageResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListServiceConversationMessageResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListServiceConversationMessageResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListServiceConversationMessageResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateServiceConversationMessage'
type UpdateServiceConversationMessageParams struct {
	// The X-Twilio-Webhook-Enabled HTTP request header
	XTwilioWebhookEnabled *string `json:"X-Twilio-Webhook-Enabled,omitempty"`
	// A string metadata field you can use to store any data you wish. The string value must contain structurally valid JSON if specified.  **Note** that if the attributes are not set \\\"{}\\\" will be returned.
	Attributes *string `json:"Attributes,omitempty"`
	// The channel specific identifier of the message's author. Defaults to `system`.
	Author *string `json:"Author,omitempty"`
	// The content of the message, can be up to 1,600 characters long.
	Body *string `json:"Body,omitempty"`
	// The date that this resource was created.
	DateCreated *time.Time `json:"DateCreated,omitempty"`
	// The date that this resource was last updated. `null` if the message has not been edited.
	DateUpdated *time.Time `json:"DateUpdated,omitempty"`
}

func (params *UpdateServiceConversationMessageParams) SetXTwilioWebhookEnabled(XTwilioWebhookEnabled string) *UpdateServiceConversationMessageParams {
	params.XTwilioWebhookEnabled = &XTwilioWebhookEnabled
	return params
}
func (params *UpdateServiceConversationMessageParams) SetAttributes(Attributes string) *UpdateServiceConversationMessageParams {
	params.Attributes = &Attributes
	return params
}
func (params *UpdateServiceConversationMessageParams) SetAuthor(Author string) *UpdateServiceConversationMessageParams {
	params.Author = &Author
	return params
}
func (params *UpdateServiceConversationMessageParams) SetBody(Body string) *UpdateServiceConversationMessageParams {
	params.Body = &Body
	return params
}
func (params *UpdateServiceConversationMessageParams) SetDateCreated(DateCreated time.Time) *UpdateServiceConversationMessageParams {
	params.DateCreated = &DateCreated
	return params
}
func (params *UpdateServiceConversationMessageParams) SetDateUpdated(DateUpdated time.Time) *UpdateServiceConversationMessageParams {
	params.DateUpdated = &DateUpdated
	return params
}

// Update an existing message in the conversation
func (c *ApiService) UpdateServiceConversationMessage(ChatServiceSid string, ConversationSid string, Sid string, params *UpdateServiceConversationMessageParams) (*ConversationsV1ServiceConversationMessage, error) {
	path := "/v1/Services/{ChatServiceSid}/Conversations/{ConversationSid}/Messages/{Sid}"
	path = strings.Replace(path, "{"+"ChatServiceSid"+"}", ChatServiceSid, -1)
	path = strings.Replace(path, "{"+"ConversationSid"+"}", ConversationSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Attributes != nil {
		data.Set("Attributes", *params.Attributes)
	}
	if params != nil && params.Author != nil {
		data.Set("Author", *params.Author)
	}
	if params != nil && params.Body != nil {
		data.Set("Body", *params.Body)
	}
	if params != nil && params.DateCreated != nil {
		data.Set("DateCreated", fmt.Sprint((*params.DateCreated).Format(time.RFC3339)))
	}
	if params != nil && params.DateUpdated != nil {
		data.Set("DateUpdated", fmt.Sprint((*params.DateUpdated).Format(time.RFC3339)))
	}

	if params != nil && params.XTwilioWebhookEnabled != nil {
		headers["X-Twilio-Webhook-Enabled"] = *params.XTwilioWebhookEnabled
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ConversationsV1ServiceConversationMessage{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
