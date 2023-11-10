# FeeDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fee** | [**Fee**](Fee.md) |  | 
**Memo** | Pointer to **string** | Additional text that describes the fee transfer. | [optional] 
**Tags** | Pointer to **string** | Descriptive metadata about the fee. | [optional] 
**Token** | **string** | Unique identifier of the fee. | 
**TransactionToken** | **string** | Unique identifier of the fee transaction. | 

## Methods

### NewFeeDetail

`func NewFeeDetail(fee Fee, token string, transactionToken string, ) *FeeDetail`

NewFeeDetail instantiates a new FeeDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeeDetailWithDefaults

`func NewFeeDetailWithDefaults() *FeeDetail`

NewFeeDetailWithDefaults instantiates a new FeeDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFee

`func (o *FeeDetail) GetFee() Fee`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *FeeDetail) GetFeeOk() (*Fee, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *FeeDetail) SetFee(v Fee)`

SetFee sets Fee field to given value.


### GetMemo

`func (o *FeeDetail) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *FeeDetail) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *FeeDetail) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *FeeDetail) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetTags

`func (o *FeeDetail) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FeeDetail) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FeeDetail) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *FeeDetail) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetToken

`func (o *FeeDetail) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *FeeDetail) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *FeeDetail) SetToken(v string)`

SetToken sets Token field to given value.


### GetTransactionToken

`func (o *FeeDetail) GetTransactionToken() string`

GetTransactionToken returns the TransactionToken field if non-nil, zero value otherwise.

### GetTransactionTokenOk

`func (o *FeeDetail) GetTransactionTokenOk() (*string, bool)`

GetTransactionTokenOk returns a tuple with the TransactionToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionToken

`func (o *FeeDetail) SetTransactionToken(v string)`

SetTransactionToken sets TransactionToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


