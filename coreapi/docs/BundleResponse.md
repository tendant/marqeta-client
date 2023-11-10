# BundleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AprPolicyDetail** | Pointer to [**PolicyAprResponse**](PolicyAprResponse.md) |  | [optional] 
**AprPolicyToken** | Pointer to **string** | Unique identifier of the bundle&#39;s APR policy. | [optional] 
**CreatedTime** | Pointer to **time.Time** | Date and time when the bundle was created on Marqeta&#39;s credit platform, in UTC. | [optional] 
**CreditProductPolicyDetail** | Pointer to [**PolicyProductResponse**](PolicyProductResponse.md) |  | [optional] 
**CreditProductPolicyToken** | Pointer to **string** | Unique identifier of the bundle&#39;s credit product policy. | [optional] 
**Description** | Pointer to **string** | Description of the bundle. | [optional] 
**DocumentPolicyDetail** | Pointer to [**PolicyDocumentResponse**](PolicyDocumentResponse.md) |  | [optional] 
**DocumentPolicyToken** | Pointer to **string** | Unique identifier of the bundle&#39;s document policy. | [optional] 
**FeePolicyDetail** | Pointer to [**PolicyFeeResponse**](PolicyFeeResponse.md) |  | [optional] 
**FeePolicyToken** | Pointer to **string** | Unique identifier of the bundle&#39;s fee policy. | [optional] 
**Name** | Pointer to **string** | Name of the bundle. | [optional] 
**OfferPolicyDetail** | Pointer to [**PolicyOfferResponse**](PolicyOfferResponse.md) |  | [optional] 
**OfferPolicyToken** | Pointer to **string** | Unique identifier of the bundle&#39;s offer policy. | [optional] 
**RewardPolicyDetail** | Pointer to [**PolicyRewardResponse**](PolicyRewardResponse.md) |  | [optional] 
**RewardPolicyToken** | Pointer to **string** | Unique identifier of the bundle&#39;s reward policy. | [optional] 
**Status** | Pointer to [**BundleResourceStatus**](BundleResourceStatus.md) |  | [optional] 
**Token** | Pointer to **string** | Unique identifier of the bundle. | [optional] 
**UpdatedTime** | Pointer to **time.Time** | Date and time when the bundle was last updated on Marqeta&#39;s credit platform, in UTC. | [optional] 

## Methods

### NewBundleResponse

`func NewBundleResponse() *BundleResponse`

NewBundleResponse instantiates a new BundleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBundleResponseWithDefaults

`func NewBundleResponseWithDefaults() *BundleResponse`

NewBundleResponseWithDefaults instantiates a new BundleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAprPolicyDetail

`func (o *BundleResponse) GetAprPolicyDetail() PolicyAprResponse`

GetAprPolicyDetail returns the AprPolicyDetail field if non-nil, zero value otherwise.

### GetAprPolicyDetailOk

`func (o *BundleResponse) GetAprPolicyDetailOk() (*PolicyAprResponse, bool)`

GetAprPolicyDetailOk returns a tuple with the AprPolicyDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAprPolicyDetail

`func (o *BundleResponse) SetAprPolicyDetail(v PolicyAprResponse)`

SetAprPolicyDetail sets AprPolicyDetail field to given value.

### HasAprPolicyDetail

`func (o *BundleResponse) HasAprPolicyDetail() bool`

HasAprPolicyDetail returns a boolean if a field has been set.

### GetAprPolicyToken

`func (o *BundleResponse) GetAprPolicyToken() string`

GetAprPolicyToken returns the AprPolicyToken field if non-nil, zero value otherwise.

### GetAprPolicyTokenOk

`func (o *BundleResponse) GetAprPolicyTokenOk() (*string, bool)`

GetAprPolicyTokenOk returns a tuple with the AprPolicyToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAprPolicyToken

`func (o *BundleResponse) SetAprPolicyToken(v string)`

SetAprPolicyToken sets AprPolicyToken field to given value.

### HasAprPolicyToken

`func (o *BundleResponse) HasAprPolicyToken() bool`

HasAprPolicyToken returns a boolean if a field has been set.

### GetCreatedTime

`func (o *BundleResponse) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *BundleResponse) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *BundleResponse) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *BundleResponse) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetCreditProductPolicyDetail

`func (o *BundleResponse) GetCreditProductPolicyDetail() PolicyProductResponse`

GetCreditProductPolicyDetail returns the CreditProductPolicyDetail field if non-nil, zero value otherwise.

### GetCreditProductPolicyDetailOk

`func (o *BundleResponse) GetCreditProductPolicyDetailOk() (*PolicyProductResponse, bool)`

GetCreditProductPolicyDetailOk returns a tuple with the CreditProductPolicyDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditProductPolicyDetail

`func (o *BundleResponse) SetCreditProductPolicyDetail(v PolicyProductResponse)`

SetCreditProductPolicyDetail sets CreditProductPolicyDetail field to given value.

### HasCreditProductPolicyDetail

`func (o *BundleResponse) HasCreditProductPolicyDetail() bool`

HasCreditProductPolicyDetail returns a boolean if a field has been set.

### GetCreditProductPolicyToken

`func (o *BundleResponse) GetCreditProductPolicyToken() string`

GetCreditProductPolicyToken returns the CreditProductPolicyToken field if non-nil, zero value otherwise.

### GetCreditProductPolicyTokenOk

`func (o *BundleResponse) GetCreditProductPolicyTokenOk() (*string, bool)`

GetCreditProductPolicyTokenOk returns a tuple with the CreditProductPolicyToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditProductPolicyToken

`func (o *BundleResponse) SetCreditProductPolicyToken(v string)`

SetCreditProductPolicyToken sets CreditProductPolicyToken field to given value.

### HasCreditProductPolicyToken

`func (o *BundleResponse) HasCreditProductPolicyToken() bool`

HasCreditProductPolicyToken returns a boolean if a field has been set.

### GetDescription

`func (o *BundleResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BundleResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BundleResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BundleResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDocumentPolicyDetail

`func (o *BundleResponse) GetDocumentPolicyDetail() PolicyDocumentResponse`

GetDocumentPolicyDetail returns the DocumentPolicyDetail field if non-nil, zero value otherwise.

### GetDocumentPolicyDetailOk

`func (o *BundleResponse) GetDocumentPolicyDetailOk() (*PolicyDocumentResponse, bool)`

GetDocumentPolicyDetailOk returns a tuple with the DocumentPolicyDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentPolicyDetail

`func (o *BundleResponse) SetDocumentPolicyDetail(v PolicyDocumentResponse)`

SetDocumentPolicyDetail sets DocumentPolicyDetail field to given value.

### HasDocumentPolicyDetail

`func (o *BundleResponse) HasDocumentPolicyDetail() bool`

HasDocumentPolicyDetail returns a boolean if a field has been set.

### GetDocumentPolicyToken

`func (o *BundleResponse) GetDocumentPolicyToken() string`

GetDocumentPolicyToken returns the DocumentPolicyToken field if non-nil, zero value otherwise.

### GetDocumentPolicyTokenOk

`func (o *BundleResponse) GetDocumentPolicyTokenOk() (*string, bool)`

GetDocumentPolicyTokenOk returns a tuple with the DocumentPolicyToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentPolicyToken

`func (o *BundleResponse) SetDocumentPolicyToken(v string)`

SetDocumentPolicyToken sets DocumentPolicyToken field to given value.

### HasDocumentPolicyToken

`func (o *BundleResponse) HasDocumentPolicyToken() bool`

HasDocumentPolicyToken returns a boolean if a field has been set.

### GetFeePolicyDetail

`func (o *BundleResponse) GetFeePolicyDetail() PolicyFeeResponse`

GetFeePolicyDetail returns the FeePolicyDetail field if non-nil, zero value otherwise.

### GetFeePolicyDetailOk

`func (o *BundleResponse) GetFeePolicyDetailOk() (*PolicyFeeResponse, bool)`

GetFeePolicyDetailOk returns a tuple with the FeePolicyDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeePolicyDetail

`func (o *BundleResponse) SetFeePolicyDetail(v PolicyFeeResponse)`

SetFeePolicyDetail sets FeePolicyDetail field to given value.

### HasFeePolicyDetail

`func (o *BundleResponse) HasFeePolicyDetail() bool`

HasFeePolicyDetail returns a boolean if a field has been set.

### GetFeePolicyToken

`func (o *BundleResponse) GetFeePolicyToken() string`

GetFeePolicyToken returns the FeePolicyToken field if non-nil, zero value otherwise.

### GetFeePolicyTokenOk

`func (o *BundleResponse) GetFeePolicyTokenOk() (*string, bool)`

GetFeePolicyTokenOk returns a tuple with the FeePolicyToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeePolicyToken

`func (o *BundleResponse) SetFeePolicyToken(v string)`

SetFeePolicyToken sets FeePolicyToken field to given value.

### HasFeePolicyToken

`func (o *BundleResponse) HasFeePolicyToken() bool`

HasFeePolicyToken returns a boolean if a field has been set.

### GetName

`func (o *BundleResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BundleResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BundleResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BundleResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOfferPolicyDetail

`func (o *BundleResponse) GetOfferPolicyDetail() PolicyOfferResponse`

GetOfferPolicyDetail returns the OfferPolicyDetail field if non-nil, zero value otherwise.

### GetOfferPolicyDetailOk

`func (o *BundleResponse) GetOfferPolicyDetailOk() (*PolicyOfferResponse, bool)`

GetOfferPolicyDetailOk returns a tuple with the OfferPolicyDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferPolicyDetail

`func (o *BundleResponse) SetOfferPolicyDetail(v PolicyOfferResponse)`

SetOfferPolicyDetail sets OfferPolicyDetail field to given value.

### HasOfferPolicyDetail

`func (o *BundleResponse) HasOfferPolicyDetail() bool`

HasOfferPolicyDetail returns a boolean if a field has been set.

### GetOfferPolicyToken

`func (o *BundleResponse) GetOfferPolicyToken() string`

GetOfferPolicyToken returns the OfferPolicyToken field if non-nil, zero value otherwise.

### GetOfferPolicyTokenOk

`func (o *BundleResponse) GetOfferPolicyTokenOk() (*string, bool)`

GetOfferPolicyTokenOk returns a tuple with the OfferPolicyToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferPolicyToken

`func (o *BundleResponse) SetOfferPolicyToken(v string)`

SetOfferPolicyToken sets OfferPolicyToken field to given value.

### HasOfferPolicyToken

`func (o *BundleResponse) HasOfferPolicyToken() bool`

HasOfferPolicyToken returns a boolean if a field has been set.

### GetRewardPolicyDetail

`func (o *BundleResponse) GetRewardPolicyDetail() PolicyRewardResponse`

GetRewardPolicyDetail returns the RewardPolicyDetail field if non-nil, zero value otherwise.

### GetRewardPolicyDetailOk

`func (o *BundleResponse) GetRewardPolicyDetailOk() (*PolicyRewardResponse, bool)`

GetRewardPolicyDetailOk returns a tuple with the RewardPolicyDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardPolicyDetail

`func (o *BundleResponse) SetRewardPolicyDetail(v PolicyRewardResponse)`

SetRewardPolicyDetail sets RewardPolicyDetail field to given value.

### HasRewardPolicyDetail

`func (o *BundleResponse) HasRewardPolicyDetail() bool`

HasRewardPolicyDetail returns a boolean if a field has been set.

### GetRewardPolicyToken

`func (o *BundleResponse) GetRewardPolicyToken() string`

GetRewardPolicyToken returns the RewardPolicyToken field if non-nil, zero value otherwise.

### GetRewardPolicyTokenOk

`func (o *BundleResponse) GetRewardPolicyTokenOk() (*string, bool)`

GetRewardPolicyTokenOk returns a tuple with the RewardPolicyToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardPolicyToken

`func (o *BundleResponse) SetRewardPolicyToken(v string)`

SetRewardPolicyToken sets RewardPolicyToken field to given value.

### HasRewardPolicyToken

`func (o *BundleResponse) HasRewardPolicyToken() bool`

HasRewardPolicyToken returns a boolean if a field has been set.

### GetStatus

`func (o *BundleResponse) GetStatus() BundleResourceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BundleResponse) GetStatusOk() (*BundleResourceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BundleResponse) SetStatus(v BundleResourceStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BundleResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetToken

`func (o *BundleResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *BundleResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *BundleResponse) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *BundleResponse) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUpdatedTime

`func (o *BundleResponse) GetUpdatedTime() time.Time`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *BundleResponse) GetUpdatedTimeOk() (*time.Time, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *BundleResponse) SetUpdatedTime(v time.Time)`

SetUpdatedTime sets UpdatedTime field to given value.

### HasUpdatedTime

`func (o *BundleResponse) HasUpdatedTime() bool`

HasUpdatedTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


