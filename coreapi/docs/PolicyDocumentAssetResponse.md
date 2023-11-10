# PolicyDocumentAssetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetCreatedTime** | Pointer to **time.Time** | Date and time when the asset was created. | [optional] 
**AssetToken** | Pointer to **string** | Unique identifier of the asset, which is a type of document that contains finalized values. The values are finalized when the bundle containing the document is created. | [optional] 
**AssetUrls** | Pointer to [**PolicyDocumentAssetURLs**](PolicyDocumentAssetURLs.md) |  | [optional] 

## Methods

### NewPolicyDocumentAssetResponse

`func NewPolicyDocumentAssetResponse() *PolicyDocumentAssetResponse`

NewPolicyDocumentAssetResponse instantiates a new PolicyDocumentAssetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyDocumentAssetResponseWithDefaults

`func NewPolicyDocumentAssetResponseWithDefaults() *PolicyDocumentAssetResponse`

NewPolicyDocumentAssetResponseWithDefaults instantiates a new PolicyDocumentAssetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetCreatedTime

`func (o *PolicyDocumentAssetResponse) GetAssetCreatedTime() time.Time`

GetAssetCreatedTime returns the AssetCreatedTime field if non-nil, zero value otherwise.

### GetAssetCreatedTimeOk

`func (o *PolicyDocumentAssetResponse) GetAssetCreatedTimeOk() (*time.Time, bool)`

GetAssetCreatedTimeOk returns a tuple with the AssetCreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetCreatedTime

`func (o *PolicyDocumentAssetResponse) SetAssetCreatedTime(v time.Time)`

SetAssetCreatedTime sets AssetCreatedTime field to given value.

### HasAssetCreatedTime

`func (o *PolicyDocumentAssetResponse) HasAssetCreatedTime() bool`

HasAssetCreatedTime returns a boolean if a field has been set.

### GetAssetToken

`func (o *PolicyDocumentAssetResponse) GetAssetToken() string`

GetAssetToken returns the AssetToken field if non-nil, zero value otherwise.

### GetAssetTokenOk

`func (o *PolicyDocumentAssetResponse) GetAssetTokenOk() (*string, bool)`

GetAssetTokenOk returns a tuple with the AssetToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetToken

`func (o *PolicyDocumentAssetResponse) SetAssetToken(v string)`

SetAssetToken sets AssetToken field to given value.

### HasAssetToken

`func (o *PolicyDocumentAssetResponse) HasAssetToken() bool`

HasAssetToken returns a boolean if a field has been set.

### GetAssetUrls

`func (o *PolicyDocumentAssetResponse) GetAssetUrls() PolicyDocumentAssetURLs`

GetAssetUrls returns the AssetUrls field if non-nil, zero value otherwise.

### GetAssetUrlsOk

`func (o *PolicyDocumentAssetResponse) GetAssetUrlsOk() (*PolicyDocumentAssetURLs, bool)`

GetAssetUrlsOk returns a tuple with the AssetUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetUrls

`func (o *PolicyDocumentAssetResponse) SetAssetUrls(v PolicyDocumentAssetURLs)`

SetAssetUrls sets AssetUrls field to given value.

### HasAssetUrls

`func (o *PolicyDocumentAssetResponse) HasAssetUrls() bool`

HasAssetUrls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


