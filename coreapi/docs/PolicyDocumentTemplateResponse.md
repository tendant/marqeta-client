# PolicyDocumentTemplateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TemplateCreatedTime** | Pointer to **time.Time** | Date and time when the template was created. | [optional] 
**TemplateToken** | Pointer to **string** | Unique identifier of a template, which is a document that serves as an initial disclosure but does not contain finalized values. | [optional] 
**TemplateUrls** | Pointer to [**PolicyDocumentTemplateURLs**](PolicyDocumentTemplateURLs.md) |  | [optional] 

## Methods

### NewPolicyDocumentTemplateResponse

`func NewPolicyDocumentTemplateResponse() *PolicyDocumentTemplateResponse`

NewPolicyDocumentTemplateResponse instantiates a new PolicyDocumentTemplateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyDocumentTemplateResponseWithDefaults

`func NewPolicyDocumentTemplateResponseWithDefaults() *PolicyDocumentTemplateResponse`

NewPolicyDocumentTemplateResponseWithDefaults instantiates a new PolicyDocumentTemplateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplateCreatedTime

`func (o *PolicyDocumentTemplateResponse) GetTemplateCreatedTime() time.Time`

GetTemplateCreatedTime returns the TemplateCreatedTime field if non-nil, zero value otherwise.

### GetTemplateCreatedTimeOk

`func (o *PolicyDocumentTemplateResponse) GetTemplateCreatedTimeOk() (*time.Time, bool)`

GetTemplateCreatedTimeOk returns a tuple with the TemplateCreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateCreatedTime

`func (o *PolicyDocumentTemplateResponse) SetTemplateCreatedTime(v time.Time)`

SetTemplateCreatedTime sets TemplateCreatedTime field to given value.

### HasTemplateCreatedTime

`func (o *PolicyDocumentTemplateResponse) HasTemplateCreatedTime() bool`

HasTemplateCreatedTime returns a boolean if a field has been set.

### GetTemplateToken

`func (o *PolicyDocumentTemplateResponse) GetTemplateToken() string`

GetTemplateToken returns the TemplateToken field if non-nil, zero value otherwise.

### GetTemplateTokenOk

`func (o *PolicyDocumentTemplateResponse) GetTemplateTokenOk() (*string, bool)`

GetTemplateTokenOk returns a tuple with the TemplateToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateToken

`func (o *PolicyDocumentTemplateResponse) SetTemplateToken(v string)`

SetTemplateToken sets TemplateToken field to given value.

### HasTemplateToken

`func (o *PolicyDocumentTemplateResponse) HasTemplateToken() bool`

HasTemplateToken returns a boolean if a field has been set.

### GetTemplateUrls

`func (o *PolicyDocumentTemplateResponse) GetTemplateUrls() PolicyDocumentTemplateURLs`

GetTemplateUrls returns the TemplateUrls field if non-nil, zero value otherwise.

### GetTemplateUrlsOk

`func (o *PolicyDocumentTemplateResponse) GetTemplateUrlsOk() (*PolicyDocumentTemplateURLs, bool)`

GetTemplateUrlsOk returns a tuple with the TemplateUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateUrls

`func (o *PolicyDocumentTemplateResponse) SetTemplateUrls(v PolicyDocumentTemplateURLs)`

SetTemplateUrls sets TemplateUrls field to given value.

### HasTemplateUrls

`func (o *PolicyDocumentTemplateResponse) HasTemplateUrls() bool`

HasTemplateUrls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


