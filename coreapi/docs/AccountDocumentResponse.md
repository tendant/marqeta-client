# AccountDocumentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptedAt** | Pointer to **time.Time** | Date and time when the user accepted the document on Marqeta&#39;s credit platform, in UTC. | [optional] 
**AssetToken** | Pointer to **string** | Unique identifier of the document on a credit account. | [optional] 
**AssetUrls** | Pointer to [**PolicyDocumentAssetURLs**](PolicyDocumentAssetURLs.md) |  | [optional] 
**EffectiveFrom** | Pointer to **time.Time** | Date and time when the document goes into effect on Marqeta&#39;s credit platform, in UTC. | [optional] 

## Methods

### NewAccountDocumentResponse

`func NewAccountDocumentResponse() *AccountDocumentResponse`

NewAccountDocumentResponse instantiates a new AccountDocumentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountDocumentResponseWithDefaults

`func NewAccountDocumentResponseWithDefaults() *AccountDocumentResponse`

NewAccountDocumentResponseWithDefaults instantiates a new AccountDocumentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptedAt

`func (o *AccountDocumentResponse) GetAcceptedAt() time.Time`

GetAcceptedAt returns the AcceptedAt field if non-nil, zero value otherwise.

### GetAcceptedAtOk

`func (o *AccountDocumentResponse) GetAcceptedAtOk() (*time.Time, bool)`

GetAcceptedAtOk returns a tuple with the AcceptedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedAt

`func (o *AccountDocumentResponse) SetAcceptedAt(v time.Time)`

SetAcceptedAt sets AcceptedAt field to given value.

### HasAcceptedAt

`func (o *AccountDocumentResponse) HasAcceptedAt() bool`

HasAcceptedAt returns a boolean if a field has been set.

### GetAssetToken

`func (o *AccountDocumentResponse) GetAssetToken() string`

GetAssetToken returns the AssetToken field if non-nil, zero value otherwise.

### GetAssetTokenOk

`func (o *AccountDocumentResponse) GetAssetTokenOk() (*string, bool)`

GetAssetTokenOk returns a tuple with the AssetToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetToken

`func (o *AccountDocumentResponse) SetAssetToken(v string)`

SetAssetToken sets AssetToken field to given value.

### HasAssetToken

`func (o *AccountDocumentResponse) HasAssetToken() bool`

HasAssetToken returns a boolean if a field has been set.

### GetAssetUrls

`func (o *AccountDocumentResponse) GetAssetUrls() PolicyDocumentAssetURLs`

GetAssetUrls returns the AssetUrls field if non-nil, zero value otherwise.

### GetAssetUrlsOk

`func (o *AccountDocumentResponse) GetAssetUrlsOk() (*PolicyDocumentAssetURLs, bool)`

GetAssetUrlsOk returns a tuple with the AssetUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetUrls

`func (o *AccountDocumentResponse) SetAssetUrls(v PolicyDocumentAssetURLs)`

SetAssetUrls sets AssetUrls field to given value.

### HasAssetUrls

`func (o *AccountDocumentResponse) HasAssetUrls() bool`

HasAssetUrls returns a boolean if a field has been set.

### GetEffectiveFrom

`func (o *AccountDocumentResponse) GetEffectiveFrom() time.Time`

GetEffectiveFrom returns the EffectiveFrom field if non-nil, zero value otherwise.

### GetEffectiveFromOk

`func (o *AccountDocumentResponse) GetEffectiveFromOk() (*time.Time, bool)`

GetEffectiveFromOk returns a tuple with the EffectiveFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveFrom

`func (o *AccountDocumentResponse) SetEffectiveFrom(v time.Time)`

SetEffectiveFrom sets EffectiveFrom field to given value.

### HasEffectiveFrom

`func (o *AccountDocumentResponse) HasEffectiveFrom() bool`

HasEffectiveFrom returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


