/*
 * Twilio - Lookups
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.11.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	twilio "github.com/twilio/twilio-go/client"
)

type DefaultApiService struct {
	baseURL string
	client  *twilio.Client
}

func NewDefaultApiService(client *twilio.Client) *DefaultApiService {
	return &DefaultApiService{
		client:  client,
		baseURL: "https://lookups.twilio.com",
	}
}

// FetchPhoneNumberParams Optional parameters for the method 'FetchPhoneNumber'
type FetchPhoneNumberParams struct {
	CountryCode *string                 `json:"CountryCode,omitempty"`
	Type        *[]string               `json:"Type,omitempty"`
	AddOns      *[]string               `json:"AddOns,omitempty"`
	AddOnsData  *map[string]interface{} `json:"AddOnsData,omitempty"`
}

/*
* FetchPhoneNumber Method for FetchPhoneNumber
* @param PhoneNumber The phone number to lookup in [E.164](https://www.twilio.com/docs/glossary/what-e164) format, which consists of a + followed by the country code and subscriber number.
* @param optional nil or *FetchPhoneNumberParams - Optional Parameters:
* @param "CountryCode" (string) - The [ISO country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) of the phone number to fetch. This is used to specify the country when the phone number is provided in a national format.
* @param "Type" ([]string) - The type of information to return. Can be: `carrier` or `caller-name`. The default is null.  Carrier information costs $0.005 per phone number looked up.  Caller Name information is currently available only in the US and costs $0.01 per phone number looked up.  To retrieve both types on information, specify this parameter twice; once with `carrier` and once with `caller-name` as the value.
* @param "AddOns" ([]string) - The `unique_name` of an Add-on you would like to invoke. Can be the `unique_name` of an Add-on that is installed on your account. You can specify multiple instances of this parameter to invoke multiple Add-ons. For more information about  Add-ons, see the [Add-ons documentation](https://www.twilio.com/docs/add-ons).
* @param "AddOnsData" (map[string]interface{}) - Data specific to the add-on you would like to invoke. The content and format of this value depends on the add-on.
* @return LookupsV1PhoneNumber
 */
func (c *DefaultApiService) FetchPhoneNumber(PhoneNumber string, params *FetchPhoneNumberParams) (*LookupsV1PhoneNumber, error) {
	path := "/v1/PhoneNumbers/{PhoneNumber}"
	path = strings.Replace(path, "{"+"PhoneNumber"+"}", PhoneNumber, -1)

	data := url.Values{}
	headers := 0

	if params != nil && params.CountryCode != nil {
		data.Set("CountryCode", *params.CountryCode)
	}
	if params != nil && params.Type != nil {
		data.Set("Type", strings.Join(*params.Type, ","))
	}
	if params != nil && params.AddOns != nil {
		data.Set("AddOns", strings.Join(*params.AddOns, ","))
	}
	if params != nil && params.AddOnsData != nil {
		v, err := json.Marshal(params.AddOnsData)

		if err != nil {
			return nil, err
		}

		data.Set("AddOnsData", fmt.Sprint(v))
	}

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &LookupsV1PhoneNumber{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
