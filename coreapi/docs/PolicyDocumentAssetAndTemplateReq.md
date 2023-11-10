# PolicyDocumentAssetAndTemplateReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetToken** | **string** | Unique identifier of the asset, which is the version of a document that is based on the template and contains finalized values. The values are finalized when the bundle containing the document is created. | 
**TemplateToken** | **string** | Unique identifier of the template, which is the version of a document that serves as an initial disclosure but does not contain finalized values. The values are finalized in the asset version of the document. | 

## Methods

### NewPolicyDocumentAssetAndTemplateReq

`func NewPolicyDocumentAssetAndTemplateReq(assetToken string, templateToken string, ) *PolicyDocumentAssetAndTemplateReq`

NewPolicyDocumentAssetAndTemplateReq instantiates a new PolicyDocumentAssetAndTemplateReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyDocumentAssetAndTemplateReqWithDefaults

`func NewPolicyDocumentAssetAndTemplateReqWithDefaults() *PolicyDocumentAssetAndTemplateReq`

NewPolicyDocumentAssetAndTemplateReqWithDefaults instantiates a new PolicyDocumentAssetAndTemplateReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetToken

`func (o *PolicyDocumentAssetAndTemplateReq) GetAssetToken() string`

GetAssetToken returns the AssetToken field if non-nil, zero value otherwise.

### GetAssetTokenOk

`func (o *PolicyDocumentAssetAndTemplateReq) GetAssetTokenOk() (*string, bool)`

GetAssetTokenOk returns a tuple with the AssetToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetToken

`func (o *PolicyDocumentAssetAndTemplateReq) SetAssetToken(v string)`

SetAssetToken sets AssetToken field to given value.


### GetTemplateToken

`func (o *PolicyDocumentAssetAndTemplateReq) GetTemplateToken() string`

GetTemplateToken returns the TemplateToken field if non-nil, zero value otherwise.

### GetTemplateTokenOk

`func (o *PolicyDocumentAssetAndTemplateReq) GetTemplateTokenOk() (*string, bool)`

GetTemplateTokenOk returns a tuple with the TemplateToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateToken

`func (o *PolicyDocumentAssetAndTemplateReq) SetTemplateToken(v string)`

SetTemplateToken sets TemplateToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


