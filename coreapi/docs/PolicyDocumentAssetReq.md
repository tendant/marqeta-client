# PolicyDocumentAssetReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetToken** | **string** | Unique identifier of the asset, which is a type of document that contains finalized values. The values are finalized when the bundle containing the document is created. | 

## Methods

### NewPolicyDocumentAssetReq

`func NewPolicyDocumentAssetReq(assetToken string, ) *PolicyDocumentAssetReq`

NewPolicyDocumentAssetReq instantiates a new PolicyDocumentAssetReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyDocumentAssetReqWithDefaults

`func NewPolicyDocumentAssetReqWithDefaults() *PolicyDocumentAssetReq`

NewPolicyDocumentAssetReqWithDefaults instantiates a new PolicyDocumentAssetReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetToken

`func (o *PolicyDocumentAssetReq) GetAssetToken() string`

GetAssetToken returns the AssetToken field if non-nil, zero value otherwise.

### GetAssetTokenOk

`func (o *PolicyDocumentAssetReq) GetAssetTokenOk() (*string, bool)`

GetAssetTokenOk returns a tuple with the AssetToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetToken

`func (o *PolicyDocumentAssetReq) SetAssetToken(v string)`

SetAssetToken sets AssetToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


