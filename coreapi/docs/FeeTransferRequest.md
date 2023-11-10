# FeeTransferRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessToken** | **string** | Specifies the business account holder to which the fee applies.  Pass either &#x60;business_token&#x60; or &#x60;user_token&#x60;, not both. | 
**Fees** | [**[]FeeModel**](FeeModel.md) | Contains attributes that define characteristics of one or more fees. | 
**Tags** | Pointer to **string** | Metadata about the transfer. | [optional] 
**Token** | Pointer to **string** | Unique identifier of the fee transfer.  If you do not include a token, the system will generate one automatically. This token is necessary for use in other API calls, so we recommend that rather than let the system generate one, you use a simple string that is easy to remember. This value cannot be updated. | [optional] 
**UserToken** | **string** | Specifies the user account holder to which the fee applies.  Pass either &#x60;user_token&#x60; or &#x60;business_token&#x60;, not both. | 

## Methods

### NewFeeTransferRequest

`func NewFeeTransferRequest(businessToken string, fees []FeeModel, userToken string, ) *FeeTransferRequest`

NewFeeTransferRequest instantiates a new FeeTransferRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeeTransferRequestWithDefaults

`func NewFeeTransferRequestWithDefaults() *FeeTransferRequest`

NewFeeTransferRequestWithDefaults instantiates a new FeeTransferRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessToken

`func (o *FeeTransferRequest) GetBusinessToken() string`

GetBusinessToken returns the BusinessToken field if non-nil, zero value otherwise.

### GetBusinessTokenOk

`func (o *FeeTransferRequest) GetBusinessTokenOk() (*string, bool)`

GetBusinessTokenOk returns a tuple with the BusinessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessToken

`func (o *FeeTransferRequest) SetBusinessToken(v string)`

SetBusinessToken sets BusinessToken field to given value.


### GetFees

`func (o *FeeTransferRequest) GetFees() []FeeModel`

GetFees returns the Fees field if non-nil, zero value otherwise.

### GetFeesOk

`func (o *FeeTransferRequest) GetFeesOk() (*[]FeeModel, bool)`

GetFeesOk returns a tuple with the Fees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFees

`func (o *FeeTransferRequest) SetFees(v []FeeModel)`

SetFees sets Fees field to given value.


### GetTags

`func (o *FeeTransferRequest) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FeeTransferRequest) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FeeTransferRequest) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *FeeTransferRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetToken

`func (o *FeeTransferRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *FeeTransferRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *FeeTransferRequest) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *FeeTransferRequest) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUserToken

`func (o *FeeTransferRequest) GetUserToken() string`

GetUserToken returns the UserToken field if non-nil, zero value otherwise.

### GetUserTokenOk

`func (o *FeeTransferRequest) GetUserTokenOk() (*string, bool)`

GetUserTokenOk returns a tuple with the UserToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserToken

`func (o *FeeTransferRequest) SetUserToken(v string)`

SetUserToken sets UserToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


