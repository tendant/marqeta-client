# WebhookResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Indicates whether the webhook is active. This field is returned if you included it in your webhook. | [optional] [default to true]
**Config** | [**WebhookConfigModel**](WebhookConfigModel.md) |  | 
**CreatedTime** | Pointer to **time.Time** | Date and time when the webhook event was created, in UTC. | [optional] 
**Events** | **[]string** | Specifies the types of events for which notifications are sent.  The wildcard character &#x60;\\*&#x60; indicates that you receive all webhook notifications, or all notifications of a specified category. For example, &#x60;\\*&#x60; indicates that you receive all webhook notifications; &#x60;transaction.*&#x60; indicates that you receive all &#x60;transaction&#x60; webhook notifications.  *NOTE:* You can only use the wildcard character with the _base_ type events, not subcategories. For example, you cannot subscribe to &#x60;cardtransition.fulfillment.\\*&#x60; events, but you can subscribe to &#x60;cardtransition.*&#x60;. | 
**LastModifiedTime** | Pointer to **time.Time** | Date and time when the associated object was last modified, in UTC. | [optional] 
**Name** | **string** | Descriptive name of the webhook. | 
**Token** | Pointer to **string** | Unique identifier of the webhook. | [optional] 

## Methods

### NewWebhookResponseModel

`func NewWebhookResponseModel(config WebhookConfigModel, events []string, name string, ) *WebhookResponseModel`

NewWebhookResponseModel instantiates a new WebhookResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookResponseModelWithDefaults

`func NewWebhookResponseModelWithDefaults() *WebhookResponseModel`

NewWebhookResponseModelWithDefaults instantiates a new WebhookResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *WebhookResponseModel) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *WebhookResponseModel) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *WebhookResponseModel) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *WebhookResponseModel) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetConfig

`func (o *WebhookResponseModel) GetConfig() WebhookConfigModel`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *WebhookResponseModel) GetConfigOk() (*WebhookConfigModel, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *WebhookResponseModel) SetConfig(v WebhookConfigModel)`

SetConfig sets Config field to given value.


### GetCreatedTime

`func (o *WebhookResponseModel) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *WebhookResponseModel) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *WebhookResponseModel) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *WebhookResponseModel) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetEvents

`func (o *WebhookResponseModel) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *WebhookResponseModel) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *WebhookResponseModel) SetEvents(v []string)`

SetEvents sets Events field to given value.


### GetLastModifiedTime

`func (o *WebhookResponseModel) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *WebhookResponseModel) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *WebhookResponseModel) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.

### HasLastModifiedTime

`func (o *WebhookResponseModel) HasLastModifiedTime() bool`

HasLastModifiedTime returns a boolean if a field has been set.

### GetName

`func (o *WebhookResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebhookResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebhookResponseModel) SetName(v string)`

SetName sets Name field to given value.


### GetToken

`func (o *WebhookResponseModel) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *WebhookResponseModel) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *WebhookResponseModel) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *WebhookResponseModel) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


