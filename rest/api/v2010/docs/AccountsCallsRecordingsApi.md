# AccountsCallsRecordingsApi

All URIs are relative to *https://api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCallRecording**](AccountsCallsRecordingsApi.md#CreateCallRecording) | **Post** /2010-04-01/Accounts/{AccountSid}/Calls/{CallSid}/Recordings.json | 
[**DeleteCallRecording**](AccountsCallsRecordingsApi.md#DeleteCallRecording) | **Delete** /2010-04-01/Accounts/{AccountSid}/Calls/{CallSid}/Recordings/{Sid}.json | 
[**FetchCallRecording**](AccountsCallsRecordingsApi.md#FetchCallRecording) | **Get** /2010-04-01/Accounts/{AccountSid}/Calls/{CallSid}/Recordings/{Sid}.json | 
[**ListCallRecording**](AccountsCallsRecordingsApi.md#ListCallRecording) | **Get** /2010-04-01/Accounts/{AccountSid}/Calls/{CallSid}/Recordings.json | 
[**UpdateCallRecording**](AccountsCallsRecordingsApi.md#UpdateCallRecording) | **Post** /2010-04-01/Accounts/{AccountSid}/Calls/{CallSid}/Recordings/{Sid}.json | 



## CreateCallRecording

> ApiV2010CallRecording CreateCallRecording(ctx, CallSidoptional)



Create a recording for the call

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**CallSid** | **string** | The SID of the [Call](https://www.twilio.com/docs/voice/api/call-resource) to associate the resource with.

### Other Parameters

Other parameters are passed through a pointer to a CreateCallRecordingParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
**RecordingChannels** | **string** | The number of channels used in the recording. Can be: &#x60;mono&#x60; or &#x60;dual&#x60; and the default is &#x60;mono&#x60;. &#x60;mono&#x60; records all parties of the call into one channel. &#x60;dual&#x60; records each party of a 2-party call into separate channels.
**RecordingStatusCallback** | **string** | The URL we should call using the &#x60;recording_status_callback_method&#x60; on each recording event specified in  &#x60;recording_status_callback_event&#x60;. For more information, see [RecordingStatusCallback parameters](https://www.twilio.com/docs/voice/api/recording#recordingstatuscallback).
**RecordingStatusCallbackEvent** | **[]string** | The recording status events on which we should call the &#x60;recording_status_callback&#x60; URL. Can be: &#x60;in-progress&#x60;, &#x60;completed&#x60; and &#x60;absent&#x60; and the default is &#x60;completed&#x60;. Separate multiple event values with a space.
**RecordingStatusCallbackMethod** | **string** | The HTTP method we should use to call &#x60;recording_status_callback&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and the default is &#x60;POST&#x60;.
**RecordingTrack** | **string** | The audio track to record for the call. Can be: &#x60;inbound&#x60;, &#x60;outbound&#x60; or &#x60;both&#x60;. The default is &#x60;both&#x60;. &#x60;inbound&#x60; records the audio that is received by Twilio. &#x60;outbound&#x60; records the audio that is generated from Twilio. &#x60;both&#x60; records the audio that is received and generated by Twilio.
**Trim** | **string** | Whether to trim any leading and trailing silence in the recording. Can be: &#x60;trim-silence&#x60; or &#x60;do-not-trim&#x60; and the default is &#x60;do-not-trim&#x60;. &#x60;trim-silence&#x60; trims the silence from the beginning and end of the recording and &#x60;do-not-trim&#x60; does not.

### Return type

[**ApiV2010CallRecording**](ApiV2010CallRecording.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCallRecording

> DeleteCallRecording(ctx, CallSidSidoptional)



Delete a recording from your account

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**CallSid** | **string** | The [Call](https://www.twilio.com/docs/voice/api/call-resource) SID of the resources to delete.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Recording resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteCallRecordingParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording resources to delete.

### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchCallRecording

> ApiV2010CallRecording FetchCallRecording(ctx, CallSidSidoptional)



Fetch an instance of a recording for a call

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**CallSid** | **string** | The [Call](https://www.twilio.com/docs/voice/api/call-resource) SID of the resource to fetch.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Recording resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchCallRecordingParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording resource to fetch.

### Return type

[**ApiV2010CallRecording**](ApiV2010CallRecording.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCallRecording

> ListCallRecordingResponse ListCallRecording(ctx, CallSidoptional)



Retrieve a list of recordings belonging to the call used to make the request

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**CallSid** | **string** | The [Call](https://www.twilio.com/docs/voice/api/call-resource) SID of the resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListCallRecordingParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording resources to read.
**DateCreated** | **string** | The &#x60;date_created&#x60; value, specified as &#x60;YYYY-MM-DD&#x60;, of the resources to read. You can also specify inequality: &#x60;DateCreated&lt;&#x3D;YYYY-MM-DD&#x60; will return recordings generated at or before midnight on a given date, and &#x60;DateCreated&gt;&#x3D;YYYY-MM-DD&#x60; returns recordings generated at or after midnight on a date.
**DateCreatedBefore** | **string** | The &#x60;date_created&#x60; value, specified as &#x60;YYYY-MM-DD&#x60;, of the resources to read. You can also specify inequality: &#x60;DateCreated&lt;&#x3D;YYYY-MM-DD&#x60; will return recordings generated at or before midnight on a given date, and &#x60;DateCreated&gt;&#x3D;YYYY-MM-DD&#x60; returns recordings generated at or after midnight on a date.
**DateCreatedAfter** | **string** | The &#x60;date_created&#x60; value, specified as &#x60;YYYY-MM-DD&#x60;, of the resources to read. You can also specify inequality: &#x60;DateCreated&lt;&#x3D;YYYY-MM-DD&#x60; will return recordings generated at or before midnight on a given date, and &#x60;DateCreated&gt;&#x3D;YYYY-MM-DD&#x60; returns recordings generated at or after midnight on a date.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**ListCallRecordingResponse**](ListCallRecordingResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCallRecording

> ApiV2010CallRecording UpdateCallRecording(ctx, CallSidSidoptional)



Changes the status of the recording to paused, stopped, or in-progress. Note: Pass `Twilio.CURRENT` instead of recording sid to reference current active recording.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**CallSid** | **string** | The [Call](https://www.twilio.com/docs/voice/api/call-resource) SID of the resource to update.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Recording resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateCallRecordingParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording resource to update.
**PauseBehavior** | **string** | Whether to record during a pause. Can be: &#x60;skip&#x60; or &#x60;silence&#x60; and the default is &#x60;silence&#x60;. &#x60;skip&#x60; does not record during the pause period, while &#x60;silence&#x60; will replace the actual audio of the call with silence during the pause period. This parameter only applies when setting &#x60;status&#x60; is set to &#x60;paused&#x60;.
**Status** | **string** | The new status of the recording. Can be: &#x60;stopped&#x60;, &#x60;paused&#x60;, &#x60;in-progress&#x60;.

### Return type

[**ApiV2010CallRecording**](ApiV2010CallRecording.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

