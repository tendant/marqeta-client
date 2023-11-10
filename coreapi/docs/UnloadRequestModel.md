# UnloadRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **float32** | Amount of funds to return to the funding source. | 
**FundingSourceAddressToken** | Pointer to **string** | Unique identifier of the funding source to use for this GPA unload order.  Send a &#x60;GET&#x60; request to &#x60;/fundingsources/addresses/user/{token}&#x60; to retrieve addresses for a specific user. | [optional] 
**Memo** | Pointer to **string** | Additional descriptive text about the GPA unload. | [optional] 
**OriginalOrderToken** | **string** | Unique identifier of the original GPA order. | 
**Tags** | Pointer to **string** | Comma-delimited list of tags describing the GPA unload order. | [optional] 
**Token** | Pointer to **string** | Unique identifier of the GPA unload order.  If you do not include a token, the system will generate one automatically. This token is necessary for use in other calls, so we recommend that rather than let the system generate one, you use a simple string that is easy to remember. This value cannot be updated. | [optional] 

## Methods

### NewUnloadRequestModel

`func NewUnloadRequestModel(amount float32, originalOrderToken string, ) *UnloadRequestModel`

NewUnloadRequestModel instantiates a new UnloadRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnloadRequestModelWithDefaults

`func NewUnloadRequestModelWithDefaults() *UnloadRequestModel`

NewUnloadRequestModelWithDefaults instantiates a new UnloadRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *UnloadRequestModel) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *UnloadRequestModel) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *UnloadRequestModel) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetFundingSourceAddressToken

`func (o *UnloadRequestModel) GetFundingSourceAddressToken() string`

GetFundingSourceAddressToken returns the FundingSourceAddressToken field if non-nil, zero value otherwise.

### GetFundingSourceAddressTokenOk

`func (o *UnloadRequestModel) GetFundingSourceAddressTokenOk() (*string, bool)`

GetFundingSourceAddressTokenOk returns a tuple with the FundingSourceAddressToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingSourceAddressToken

`func (o *UnloadRequestModel) SetFundingSourceAddressToken(v string)`

SetFundingSourceAddressToken sets FundingSourceAddressToken field to given value.

### HasFundingSourceAddressToken

`func (o *UnloadRequestModel) HasFundingSourceAddressToken() bool`

HasFundingSourceAddressToken returns a boolean if a field has been set.

### GetMemo

`func (o *UnloadRequestModel) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *UnloadRequestModel) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *UnloadRequestModel) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *UnloadRequestModel) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetOriginalOrderToken

`func (o *UnloadRequestModel) GetOriginalOrderToken() string`

GetOriginalOrderToken returns the OriginalOrderToken field if non-nil, zero value otherwise.

### GetOriginalOrderTokenOk

`func (o *UnloadRequestModel) GetOriginalOrderTokenOk() (*string, bool)`

GetOriginalOrderTokenOk returns a tuple with the OriginalOrderToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalOrderToken

`func (o *UnloadRequestModel) SetOriginalOrderToken(v string)`

SetOriginalOrderToken sets OriginalOrderToken field to given value.


### GetTags

`func (o *UnloadRequestModel) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UnloadRequestModel) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UnloadRequestModel) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UnloadRequestModel) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetToken

`func (o *UnloadRequestModel) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UnloadRequestModel) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UnloadRequestModel) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *UnloadRequestModel) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


