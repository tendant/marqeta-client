# BundleCreateReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AprPolicyToken** | **string** | Unique identifier of the APR policy. | 
**CreditProductPolicyToken** | **string** | Unique identifier of the credit product policy. | 
**Description** | **string** | Description of the bundle. | 
**DocumentPolicyToken** | **string** | Unique identifier of the document policy. | 
**FeePolicyToken** | **string** | Unique identifier of the fee policy. | 
**Name** | **string** | Name of the bundle. | 
**OfferPolicyToken** | Pointer to **string** | Unique identifier of the offer policy. | [optional] 
**RewardPolicyToken** | Pointer to **string** | Unique identifier of the reward policy. | [optional] 
**Status** | [**BundleResourceStatus**](BundleResourceStatus.md) |  | 
**Token** | Pointer to **string** | Unique identifier of the bundle. | [optional] 

## Methods

### NewBundleCreateReq

`func NewBundleCreateReq(aprPolicyToken string, creditProductPolicyToken string, description string, documentPolicyToken string, feePolicyToken string, name string, status BundleResourceStatus, ) *BundleCreateReq`

NewBundleCreateReq instantiates a new BundleCreateReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBundleCreateReqWithDefaults

`func NewBundleCreateReqWithDefaults() *BundleCreateReq`

NewBundleCreateReqWithDefaults instantiates a new BundleCreateReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAprPolicyToken

`func (o *BundleCreateReq) GetAprPolicyToken() string`

GetAprPolicyToken returns the AprPolicyToken field if non-nil, zero value otherwise.

### GetAprPolicyTokenOk

`func (o *BundleCreateReq) GetAprPolicyTokenOk() (*string, bool)`

GetAprPolicyTokenOk returns a tuple with the AprPolicyToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAprPolicyToken

`func (o *BundleCreateReq) SetAprPolicyToken(v string)`

SetAprPolicyToken sets AprPolicyToken field to given value.


### GetCreditProductPolicyToken

`func (o *BundleCreateReq) GetCreditProductPolicyToken() string`

GetCreditProductPolicyToken returns the CreditProductPolicyToken field if non-nil, zero value otherwise.

### GetCreditProductPolicyTokenOk

`func (o *BundleCreateReq) GetCreditProductPolicyTokenOk() (*string, bool)`

GetCreditProductPolicyTokenOk returns a tuple with the CreditProductPolicyToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditProductPolicyToken

`func (o *BundleCreateReq) SetCreditProductPolicyToken(v string)`

SetCreditProductPolicyToken sets CreditProductPolicyToken field to given value.


### GetDescription

`func (o *BundleCreateReq) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BundleCreateReq) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BundleCreateReq) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDocumentPolicyToken

`func (o *BundleCreateReq) GetDocumentPolicyToken() string`

GetDocumentPolicyToken returns the DocumentPolicyToken field if non-nil, zero value otherwise.

### GetDocumentPolicyTokenOk

`func (o *BundleCreateReq) GetDocumentPolicyTokenOk() (*string, bool)`

GetDocumentPolicyTokenOk returns a tuple with the DocumentPolicyToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentPolicyToken

`func (o *BundleCreateReq) SetDocumentPolicyToken(v string)`

SetDocumentPolicyToken sets DocumentPolicyToken field to given value.


### GetFeePolicyToken

`func (o *BundleCreateReq) GetFeePolicyToken() string`

GetFeePolicyToken returns the FeePolicyToken field if non-nil, zero value otherwise.

### GetFeePolicyTokenOk

`func (o *BundleCreateReq) GetFeePolicyTokenOk() (*string, bool)`

GetFeePolicyTokenOk returns a tuple with the FeePolicyToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeePolicyToken

`func (o *BundleCreateReq) SetFeePolicyToken(v string)`

SetFeePolicyToken sets FeePolicyToken field to given value.


### GetName

`func (o *BundleCreateReq) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BundleCreateReq) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BundleCreateReq) SetName(v string)`

SetName sets Name field to given value.


### GetOfferPolicyToken

`func (o *BundleCreateReq) GetOfferPolicyToken() string`

GetOfferPolicyToken returns the OfferPolicyToken field if non-nil, zero value otherwise.

### GetOfferPolicyTokenOk

`func (o *BundleCreateReq) GetOfferPolicyTokenOk() (*string, bool)`

GetOfferPolicyTokenOk returns a tuple with the OfferPolicyToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferPolicyToken

`func (o *BundleCreateReq) SetOfferPolicyToken(v string)`

SetOfferPolicyToken sets OfferPolicyToken field to given value.

### HasOfferPolicyToken

`func (o *BundleCreateReq) HasOfferPolicyToken() bool`

HasOfferPolicyToken returns a boolean if a field has been set.

### GetRewardPolicyToken

`func (o *BundleCreateReq) GetRewardPolicyToken() string`

GetRewardPolicyToken returns the RewardPolicyToken field if non-nil, zero value otherwise.

### GetRewardPolicyTokenOk

`func (o *BundleCreateReq) GetRewardPolicyTokenOk() (*string, bool)`

GetRewardPolicyTokenOk returns a tuple with the RewardPolicyToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardPolicyToken

`func (o *BundleCreateReq) SetRewardPolicyToken(v string)`

SetRewardPolicyToken sets RewardPolicyToken field to given value.

### HasRewardPolicyToken

`func (o *BundleCreateReq) HasRewardPolicyToken() bool`

HasRewardPolicyToken returns a boolean if a field has been set.

### GetStatus

`func (o *BundleCreateReq) GetStatus() BundleResourceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BundleCreateReq) GetStatusOk() (*BundleResourceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BundleCreateReq) SetStatus(v BundleResourceStatus)`

SetStatus sets Status field to given value.


### GetToken

`func (o *BundleCreateReq) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *BundleCreateReq) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *BundleCreateReq) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *BundleCreateReq) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


