# PolicyDocumentAssetAndTemplateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetCreatedTime** | Pointer to **time.Time** | Date and time when the asset was created. | [optional] 
**AssetToken** | Pointer to **string** | Unique identifier of the asset, which is the version of a document that is based on the template and contains finalized values. The values are finalized when the bundle containing the document is created. | [optional] 
**AssetUrls** | Pointer to [**PolicyDocumentAssetURLs**](PolicyDocumentAssetURLs.md) |  | [optional] 
**TemplateCreatedTime** | Pointer to **time.Time** | Date and time when the template was created. | [optional] 
**TemplateToken** | Pointer to **string** | Unique identifier of the template, which is the version of a document that serves as an initial disclosure but does not contain finalized values. Values are finalized in the asset version of the document. | [optional] 
**TemplateUrls** | Pointer to [**PolicyDocumentTemplateURLs**](PolicyDocumentTemplateURLs.md) |  | [optional] 

## Methods

### NewPolicyDocumentAssetAndTemplateResponse

`func NewPolicyDocumentAssetAndTemplateResponse() *PolicyDocumentAssetAndTemplateResponse`

NewPolicyDocumentAssetAndTemplateResponse instantiates a new PolicyDocumentAssetAndTemplateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyDocumentAssetAndTemplateResponseWithDefaults

`func NewPolicyDocumentAssetAndTemplateResponseWithDefaults() *PolicyDocumentAssetAndTemplateResponse`

NewPolicyDocumentAssetAndTemplateResponseWithDefaults instantiates a new PolicyDocumentAssetAndTemplateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetCreatedTime

`func (o *PolicyDocumentAssetAndTemplateResponse) GetAssetCreatedTime() time.Time`

GetAssetCreatedTime returns the AssetCreatedTime field if non-nil, zero value otherwise.

### GetAssetCreatedTimeOk

`func (o *PolicyDocumentAssetAndTemplateResponse) GetAssetCreatedTimeOk() (*time.Time, bool)`

GetAssetCreatedTimeOk returns a tuple with the AssetCreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetCreatedTime

`func (o *PolicyDocumentAssetAndTemplateResponse) SetAssetCreatedTime(v time.Time)`

SetAssetCreatedTime sets AssetCreatedTime field to given value.

### HasAssetCreatedTime

`func (o *PolicyDocumentAssetAndTemplateResponse) HasAssetCreatedTime() bool`

HasAssetCreatedTime returns a boolean if a field has been set.

### GetAssetToken

`func (o *PolicyDocumentAssetAndTemplateResponse) GetAssetToken() string`

GetAssetToken returns the AssetToken field if non-nil, zero value otherwise.

### GetAssetTokenOk

`func (o *PolicyDocumentAssetAndTemplateResponse) GetAssetTokenOk() (*string, bool)`

GetAssetTokenOk returns a tuple with the AssetToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetToken

`func (o *PolicyDocumentAssetAndTemplateResponse) SetAssetToken(v string)`

SetAssetToken sets AssetToken field to given value.

### HasAssetToken

`func (o *PolicyDocumentAssetAndTemplateResponse) HasAssetToken() bool`

HasAssetToken returns a boolean if a field has been set.

### GetAssetUrls

`func (o *PolicyDocumentAssetAndTemplateResponse) GetAssetUrls() PolicyDocumentAssetURLs`

GetAssetUrls returns the AssetUrls field if non-nil, zero value otherwise.

### GetAssetUrlsOk

`func (o *PolicyDocumentAssetAndTemplateResponse) GetAssetUrlsOk() (*PolicyDocumentAssetURLs, bool)`

GetAssetUrlsOk returns a tuple with the AssetUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetUrls

`func (o *PolicyDocumentAssetAndTemplateResponse) SetAssetUrls(v PolicyDocumentAssetURLs)`

SetAssetUrls sets AssetUrls field to given value.

### HasAssetUrls

`func (o *PolicyDocumentAssetAndTemplateResponse) HasAssetUrls() bool`

HasAssetUrls returns a boolean if a field has been set.

### GetTemplateCreatedTime

`func (o *PolicyDocumentAssetAndTemplateResponse) GetTemplateCreatedTime() time.Time`

GetTemplateCreatedTime returns the TemplateCreatedTime field if non-nil, zero value otherwise.

### GetTemplateCreatedTimeOk

`func (o *PolicyDocumentAssetAndTemplateResponse) GetTemplateCreatedTimeOk() (*time.Time, bool)`

GetTemplateCreatedTimeOk returns a tuple with the TemplateCreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateCreatedTime

`func (o *PolicyDocumentAssetAndTemplateResponse) SetTemplateCreatedTime(v time.Time)`

SetTemplateCreatedTime sets TemplateCreatedTime field to given value.

### HasTemplateCreatedTime

`func (o *PolicyDocumentAssetAndTemplateResponse) HasTemplateCreatedTime() bool`

HasTemplateCreatedTime returns a boolean if a field has been set.

### GetTemplateToken

`func (o *PolicyDocumentAssetAndTemplateResponse) GetTemplateToken() string`

GetTemplateToken returns the TemplateToken field if non-nil, zero value otherwise.

### GetTemplateTokenOk

`func (o *PolicyDocumentAssetAndTemplateResponse) GetTemplateTokenOk() (*string, bool)`

GetTemplateTokenOk returns a tuple with the TemplateToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateToken

`func (o *PolicyDocumentAssetAndTemplateResponse) SetTemplateToken(v string)`

SetTemplateToken sets TemplateToken field to given value.

### HasTemplateToken

`func (o *PolicyDocumentAssetAndTemplateResponse) HasTemplateToken() bool`

HasTemplateToken returns a boolean if a field has been set.

### GetTemplateUrls

`func (o *PolicyDocumentAssetAndTemplateResponse) GetTemplateUrls() PolicyDocumentTemplateURLs`

GetTemplateUrls returns the TemplateUrls field if non-nil, zero value otherwise.

### GetTemplateUrlsOk

`func (o *PolicyDocumentAssetAndTemplateResponse) GetTemplateUrlsOk() (*PolicyDocumentTemplateURLs, bool)`

GetTemplateUrlsOk returns a tuple with the TemplateUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateUrls

`func (o *PolicyDocumentAssetAndTemplateResponse) SetTemplateUrls(v PolicyDocumentTemplateURLs)`

SetTemplateUrls sets TemplateUrls field to given value.

### HasTemplateUrls

`func (o *PolicyDocumentAssetAndTemplateResponse) HasTemplateUrls() bool`

HasTemplateUrls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


