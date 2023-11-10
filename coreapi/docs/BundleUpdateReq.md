# BundleUpdateReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AprPolicyToken** | **string** | Unique identifier of the bundle&#39;s APR policy. | 
**CreditProductPolicyToken** | **string** | Unique identifier of the bundle&#39;s credit product policy. | 
**Description** | **string** | Description of the bundle. | 
**DocumentPolicyToken** | **string** | Unique identifier of the bundle&#39;s document policy. | 
**FeePolicyToken** | **string** | Unique identifier of the bundle&#39;s fee policy. | 
**Name** | **string** | Name of the bundle. | 
**OfferPolicyToken** | Pointer to **string** | Unique identifier of the bundle&#39;s offer policy. | [optional] 
**RewardPolicyToken** | Pointer to **string** | Unique identifier of the bundle&#39;s reward policy. | [optional] 

## Methods

### NewBundleUpdateReq

`func NewBundleUpdateReq(aprPolicyToken string, creditProductPolicyToken string, description string, documentPolicyToken string, feePolicyToken string, name string, ) *BundleUpdateReq`

NewBundleUpdateReq instantiates a new BundleUpdateReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBundleUpdateReqWithDefaults

`func NewBundleUpdateReqWithDefaults() *BundleUpdateReq`

NewBundleUpdateReqWithDefaults instantiates a new BundleUpdateReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAprPolicyToken

`func (o *BundleUpdateReq) GetAprPolicyToken() string`

GetAprPolicyToken returns the AprPolicyToken field if non-nil, zero value otherwise.

### GetAprPolicyTokenOk

`func (o *BundleUpdateReq) GetAprPolicyTokenOk() (*string, bool)`

GetAprPolicyTokenOk returns a tuple with the AprPolicyToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAprPolicyToken

`func (o *BundleUpdateReq) SetAprPolicyToken(v string)`

SetAprPolicyToken sets AprPolicyToken field to given value.


### GetCreditProductPolicyToken

`func (o *BundleUpdateReq) GetCreditProductPolicyToken() string`

GetCreditProductPolicyToken returns the CreditProductPolicyToken field if non-nil, zero value otherwise.

### GetCreditProductPolicyTokenOk

`func (o *BundleUpdateReq) GetCreditProductPolicyTokenOk() (*string, bool)`

GetCreditProductPolicyTokenOk returns a tuple with the CreditProductPolicyToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditProductPolicyToken

`func (o *BundleUpdateReq) SetCreditProductPolicyToken(v string)`

SetCreditProductPolicyToken sets CreditProductPolicyToken field to given value.


### GetDescription

`func (o *BundleUpdateReq) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BundleUpdateReq) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BundleUpdateReq) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDocumentPolicyToken

`func (o *BundleUpdateReq) GetDocumentPolicyToken() string`

GetDocumentPolicyToken returns the DocumentPolicyToken field if non-nil, zero value otherwise.

### GetDocumentPolicyTokenOk

`func (o *BundleUpdateReq) GetDocumentPolicyTokenOk() (*string, bool)`

GetDocumentPolicyTokenOk returns a tuple with the DocumentPolicyToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentPolicyToken

`func (o *BundleUpdateReq) SetDocumentPolicyToken(v string)`

SetDocumentPolicyToken sets DocumentPolicyToken field to given value.


### GetFeePolicyToken

`func (o *BundleUpdateReq) GetFeePolicyToken() string`

GetFeePolicyToken returns the FeePolicyToken field if non-nil, zero value otherwise.

### GetFeePolicyTokenOk

`func (o *BundleUpdateReq) GetFeePolicyTokenOk() (*string, bool)`

GetFeePolicyTokenOk returns a tuple with the FeePolicyToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeePolicyToken

`func (o *BundleUpdateReq) SetFeePolicyToken(v string)`

SetFeePolicyToken sets FeePolicyToken field to given value.


### GetName

`func (o *BundleUpdateReq) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BundleUpdateReq) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BundleUpdateReq) SetName(v string)`

SetName sets Name field to given value.


### GetOfferPolicyToken

`func (o *BundleUpdateReq) GetOfferPolicyToken() string`

GetOfferPolicyToken returns the OfferPolicyToken field if non-nil, zero value otherwise.

### GetOfferPolicyTokenOk

`func (o *BundleUpdateReq) GetOfferPolicyTokenOk() (*string, bool)`

GetOfferPolicyTokenOk returns a tuple with the OfferPolicyToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferPolicyToken

`func (o *BundleUpdateReq) SetOfferPolicyToken(v string)`

SetOfferPolicyToken sets OfferPolicyToken field to given value.

### HasOfferPolicyToken

`func (o *BundleUpdateReq) HasOfferPolicyToken() bool`

HasOfferPolicyToken returns a boolean if a field has been set.

### GetRewardPolicyToken

`func (o *BundleUpdateReq) GetRewardPolicyToken() string`

GetRewardPolicyToken returns the RewardPolicyToken field if non-nil, zero value otherwise.

### GetRewardPolicyTokenOk

`func (o *BundleUpdateReq) GetRewardPolicyTokenOk() (*string, bool)`

GetRewardPolicyTokenOk returns a tuple with the RewardPolicyToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardPolicyToken

`func (o *BundleUpdateReq) SetRewardPolicyToken(v string)`

SetRewardPolicyToken sets RewardPolicyToken field to given value.

### HasRewardPolicyToken

`func (o *BundleUpdateReq) HasRewardPolicyToken() bool`

HasRewardPolicyToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


