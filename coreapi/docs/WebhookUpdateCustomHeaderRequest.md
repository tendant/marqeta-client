# WebhookUpdateCustomHeaderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomHeader** | Pointer to **map[string]string** | Additional custom information included in the HTTP header. For example, this might contain security information, along with Basic Authentication, when making a JIT Funding request. | [optional] 

## Methods

### NewWebhookUpdateCustomHeaderRequest

`func NewWebhookUpdateCustomHeaderRequest() *WebhookUpdateCustomHeaderRequest`

NewWebhookUpdateCustomHeaderRequest instantiates a new WebhookUpdateCustomHeaderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookUpdateCustomHeaderRequestWithDefaults

`func NewWebhookUpdateCustomHeaderRequestWithDefaults() *WebhookUpdateCustomHeaderRequest`

NewWebhookUpdateCustomHeaderRequestWithDefaults instantiates a new WebhookUpdateCustomHeaderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomHeader

`func (o *WebhookUpdateCustomHeaderRequest) GetCustomHeader() map[string]string`

GetCustomHeader returns the CustomHeader field if non-nil, zero value otherwise.

### GetCustomHeaderOk

`func (o *WebhookUpdateCustomHeaderRequest) GetCustomHeaderOk() (*map[string]string, bool)`

GetCustomHeaderOk returns a tuple with the CustomHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomHeader

`func (o *WebhookUpdateCustomHeaderRequest) SetCustomHeader(v map[string]string)`

SetCustomHeader sets CustomHeader field to given value.

### HasCustomHeader

`func (o *WebhookUpdateCustomHeaderRequest) HasCustomHeader() bool`

HasCustomHeader returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


