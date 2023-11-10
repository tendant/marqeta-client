# WebhookRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Indicates whether the webhook is active. | [optional] [default to true]
**Config** | [**WebhookConfigModel**](WebhookConfigModel.md) |  | 
**Events** | **[]string** | Specifies the types of events for which notifications are sent.  The wildcard character &#x60;\\*&#x60; indicates that you receive all webhook notifications, or all notifications of a specified category. For example, &#x60;\\*&#x60; indicates that you receive all webhook notifications; &#x60;transaction.*&#x60; indicates that you receive all &#x60;transaction&#x60; webhook notifications.  *NOTE:* You can only use the wildcard character with the _base_ type events, not subcategories. For example, you cannot subscribe to &#x60;cardtransition.fulfillment.\\*&#x60; events, but you can subscribe to &#x60;cardtransition.*&#x60;. | 
**Name** | **string** | Descriptive name of the webhook. | 
**Token** | Pointer to **string** | Unique identifier of the webhook. | [optional] 

## Methods

### NewWebhookRequestModel

`func NewWebhookRequestModel(config WebhookConfigModel, events []string, name string, ) *WebhookRequestModel`

NewWebhookRequestModel instantiates a new WebhookRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookRequestModelWithDefaults

`func NewWebhookRequestModelWithDefaults() *WebhookRequestModel`

NewWebhookRequestModelWithDefaults instantiates a new WebhookRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *WebhookRequestModel) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *WebhookRequestModel) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *WebhookRequestModel) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *WebhookRequestModel) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetConfig

`func (o *WebhookRequestModel) GetConfig() WebhookConfigModel`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *WebhookRequestModel) GetConfigOk() (*WebhookConfigModel, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *WebhookRequestModel) SetConfig(v WebhookConfigModel)`

SetConfig sets Config field to given value.


### GetEvents

`func (o *WebhookRequestModel) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *WebhookRequestModel) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *WebhookRequestModel) SetEvents(v []string)`

SetEvents sets Events field to given value.


### GetName

`func (o *WebhookRequestModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebhookRequestModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebhookRequestModel) SetName(v string)`

SetName sets Name field to given value.


### GetToken

`func (o *WebhookRequestModel) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *WebhookRequestModel) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *WebhookRequestModel) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *WebhookRequestModel) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


