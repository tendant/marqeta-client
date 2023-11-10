# WebhookBaseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Indicates whether the webhook is active. | [optional] [default to true]
**Config** | [**WebhookConfigModel**](WebhookConfigModel.md) |  | 
**Events** | **[]string** | Specifies the types of events for which notifications are sent.  The wildcard character &#x60;\\*&#x60; indicates that you receive all webhook notifications, or all notifications of a specified category. For example, &#x60;\\*&#x60; indicates that you receive all webhook notifications; &#x60;transaction.*&#x60; indicates that you receive all &#x60;transaction&#x60; webhook notifications.  *NOTE:* You can only use the wildcard character with the _base_ type events, not subcategories. For example, you cannot subscribe to &#x60;cardtransition.fulfillment.\\*&#x60; events, but you can subscribe to &#x60;cardtransition.*&#x60;. | 
**Name** | **string** | Descriptive name of the webhook. | 

## Methods

### NewWebhookBaseModel

`func NewWebhookBaseModel(config WebhookConfigModel, events []string, name string, ) *WebhookBaseModel`

NewWebhookBaseModel instantiates a new WebhookBaseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookBaseModelWithDefaults

`func NewWebhookBaseModelWithDefaults() *WebhookBaseModel`

NewWebhookBaseModelWithDefaults instantiates a new WebhookBaseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *WebhookBaseModel) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *WebhookBaseModel) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *WebhookBaseModel) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *WebhookBaseModel) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetConfig

`func (o *WebhookBaseModel) GetConfig() WebhookConfigModel`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *WebhookBaseModel) GetConfigOk() (*WebhookConfigModel, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *WebhookBaseModel) SetConfig(v WebhookConfigModel)`

SetConfig sets Config field to given value.


### GetEvents

`func (o *WebhookBaseModel) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *WebhookBaseModel) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *WebhookBaseModel) SetEvents(v []string)`

SetEvents sets Events field to given value.


### GetName

`func (o *WebhookBaseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebhookBaseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebhookBaseModel) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


