# FeeModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Memo** | Pointer to **string** | Additional text that describes the fee. | [optional] 
**Tags** | Pointer to **string** | Descriptive metadata about the fee. | [optional] 
**Token** | **string** | Unique identifier of the fee. | 

## Methods

### NewFeeModel

`func NewFeeModel(token string, ) *FeeModel`

NewFeeModel instantiates a new FeeModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeeModelWithDefaults

`func NewFeeModelWithDefaults() *FeeModel`

NewFeeModelWithDefaults instantiates a new FeeModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMemo

`func (o *FeeModel) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *FeeModel) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *FeeModel) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *FeeModel) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetTags

`func (o *FeeModel) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FeeModel) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FeeModel) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *FeeModel) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetToken

`func (o *FeeModel) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *FeeModel) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *FeeModel) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


