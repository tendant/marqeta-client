# Webhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Endpoint** | **string** | Valid URL | 
**Password** | **string** | Authentication password | 
**Secret** | Pointer to **string** | Authentication secret | [optional] 
**Username** | **string** | Authentication username | 

## Methods

### NewWebhook

`func NewWebhook(endpoint string, password string, username string, ) *Webhook`

NewWebhook instantiates a new Webhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookWithDefaults

`func NewWebhookWithDefaults() *Webhook`

NewWebhookWithDefaults instantiates a new Webhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpoint

`func (o *Webhook) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *Webhook) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *Webhook) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetPassword

`func (o *Webhook) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *Webhook) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *Webhook) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetSecret

`func (o *Webhook) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *Webhook) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *Webhook) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *Webhook) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetUsername

`func (o *Webhook) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *Webhook) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *Webhook) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


