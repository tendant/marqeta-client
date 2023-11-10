# WebhookConfigModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BasicAuthPassword** | **string** | Password for accessing your webhook endpoint. | 
**BasicAuthUsername** | **string** | Username for accessing your webhook endpoint. | 
**CustomHeader** | Pointer to **map[string]string** | Custom headers to be passed along with the request. | [optional] 
**Secret** | Pointer to **string** | Randomly chosen string used for implementing HMAC-SHA1.  HMAC-SHA1 provides an added layer of security by authenticating the message and validating message integrity. Using this functionality requires that your webhook endpoint verify the message signature. For information about implementing this functionality, see &lt;&lt;/developer-guides/signature-verification, Signature Verification&gt;&gt;. | [optional] 
**Url** | **string** | URL of your webhook endpoint. | 
**UseMtls** | Pointer to **bool** | Set to &#x60;true&#x60; to use MTLS for the webhook. | [optional] [default to false]

## Methods

### NewWebhookConfigModel

`func NewWebhookConfigModel(basicAuthPassword string, basicAuthUsername string, url string, ) *WebhookConfigModel`

NewWebhookConfigModel instantiates a new WebhookConfigModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookConfigModelWithDefaults

`func NewWebhookConfigModelWithDefaults() *WebhookConfigModel`

NewWebhookConfigModelWithDefaults instantiates a new WebhookConfigModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBasicAuthPassword

`func (o *WebhookConfigModel) GetBasicAuthPassword() string`

GetBasicAuthPassword returns the BasicAuthPassword field if non-nil, zero value otherwise.

### GetBasicAuthPasswordOk

`func (o *WebhookConfigModel) GetBasicAuthPasswordOk() (*string, bool)`

GetBasicAuthPasswordOk returns a tuple with the BasicAuthPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuthPassword

`func (o *WebhookConfigModel) SetBasicAuthPassword(v string)`

SetBasicAuthPassword sets BasicAuthPassword field to given value.


### GetBasicAuthUsername

`func (o *WebhookConfigModel) GetBasicAuthUsername() string`

GetBasicAuthUsername returns the BasicAuthUsername field if non-nil, zero value otherwise.

### GetBasicAuthUsernameOk

`func (o *WebhookConfigModel) GetBasicAuthUsernameOk() (*string, bool)`

GetBasicAuthUsernameOk returns a tuple with the BasicAuthUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuthUsername

`func (o *WebhookConfigModel) SetBasicAuthUsername(v string)`

SetBasicAuthUsername sets BasicAuthUsername field to given value.


### GetCustomHeader

`func (o *WebhookConfigModel) GetCustomHeader() map[string]string`

GetCustomHeader returns the CustomHeader field if non-nil, zero value otherwise.

### GetCustomHeaderOk

`func (o *WebhookConfigModel) GetCustomHeaderOk() (*map[string]string, bool)`

GetCustomHeaderOk returns a tuple with the CustomHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomHeader

`func (o *WebhookConfigModel) SetCustomHeader(v map[string]string)`

SetCustomHeader sets CustomHeader field to given value.

### HasCustomHeader

`func (o *WebhookConfigModel) HasCustomHeader() bool`

HasCustomHeader returns a boolean if a field has been set.

### GetSecret

`func (o *WebhookConfigModel) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *WebhookConfigModel) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *WebhookConfigModel) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *WebhookConfigModel) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetUrl

`func (o *WebhookConfigModel) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebhookConfigModel) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebhookConfigModel) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetUseMtls

`func (o *WebhookConfigModel) GetUseMtls() bool`

GetUseMtls returns the UseMtls field if non-nil, zero value otherwise.

### GetUseMtlsOk

`func (o *WebhookConfigModel) GetUseMtlsOk() (*bool, bool)`

GetUseMtlsOk returns a tuple with the UseMtls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseMtls

`func (o *WebhookConfigModel) SetUseMtls(v bool)`

SetUseMtls sets UseMtls field to given value.

### HasUseMtls

`func (o *WebhookConfigModel) HasUseMtls() bool`

HasUseMtls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


