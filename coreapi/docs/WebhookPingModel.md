# WebhookPingModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pings** | [**[]EchoPingRequest**](EchoPingRequest.md) | Array of ping requests to your webhook endpoint. | 

## Methods

### NewWebhookPingModel

`func NewWebhookPingModel(pings []EchoPingRequest, ) *WebhookPingModel`

NewWebhookPingModel instantiates a new WebhookPingModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookPingModelWithDefaults

`func NewWebhookPingModelWithDefaults() *WebhookPingModel`

NewWebhookPingModelWithDefaults instantiates a new WebhookPingModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPings

`func (o *WebhookPingModel) GetPings() []EchoPingRequest`

GetPings returns the Pings field if non-nil, zero value otherwise.

### GetPingsOk

`func (o *WebhookPingModel) GetPingsOk() (*[]EchoPingRequest, bool)`

GetPingsOk returns a tuple with the Pings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPings

`func (o *WebhookPingModel) SetPings(v []EchoPingRequest)`

SetPings sets Pings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


