/*
 * Twilio - Autopilot
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
	"net/url"

	"strings"
)

// Returns JSON actions for the Task.
func (c *ApiService) FetchTaskActions(AssistantSid string, TaskSid string) (*AutopilotV1TaskActions, error) {
	path := "/v1/Assistants/{AssistantSid}/Tasks/{TaskSid}/Actions"
	path = strings.Replace(path, "{"+"AssistantSid"+"}", AssistantSid, -1)
	path = strings.Replace(path, "{"+"TaskSid"+"}", TaskSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &AutopilotV1TaskActions{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'UpdateTaskActions'
type UpdateTaskActionsParams struct {
	// The JSON string that specifies the [actions](https://www.twilio.com/docs/autopilot/actions) that instruct the Assistant on how to perform the task.
	Actions *map[string]interface{} `json:"Actions,omitempty"`
}

func (params *UpdateTaskActionsParams) SetActions(Actions map[string]interface{}) *UpdateTaskActionsParams {
	params.Actions = &Actions
	return params
}

// Updates the actions of an Task identified by {TaskSid} or {TaskUniqueName}.
func (c *ApiService) UpdateTaskActions(AssistantSid string, TaskSid string, params *UpdateTaskActionsParams) (*AutopilotV1TaskActions, error) {
	path := "/v1/Assistants/{AssistantSid}/Tasks/{TaskSid}/Actions"
	path = strings.Replace(path, "{"+"AssistantSid"+"}", AssistantSid, -1)
	path = strings.Replace(path, "{"+"TaskSid"+"}", TaskSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Actions != nil {
		v, err := json.Marshal(params.Actions)

		if err != nil {
			return nil, err
		}

		data.Set("Actions", string(v))
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &AutopilotV1TaskActions{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
