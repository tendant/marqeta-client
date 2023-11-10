# ProgramTransfer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **float32** | Amount of program transfer. | 
**BusinessToken** | Pointer to **string** | Unique identifier of the business. Pass either a &#x60;business_token&#x60; or a &#x60;user_token&#x60;, not both.  Send a &#x60;GET&#x60; request to &#x60;/businesses&#x60; to retrieve business tokens. | [optional] 
**CurrencyCode** | **string** | Three-digit ISO 4217 currency code. | 
**Fees** | Pointer to [**[]FeeModel**](FeeModel.md) | Contains attributes that define characteristics of one or more fees. This array is returned in the response when it is included in the request. | [optional] 
**Memo** | Pointer to **string** | Memo or note describing the program transfer. | [optional] 
**Tags** | Pointer to **string** | Comma-delimited list of tags describing the program transfer. | [optional] 
**Token** | Pointer to **string** | Unique identifier of the program transfer.  If you do not include a token, the system will generate one automatically. This token is necessary for use in other API calls, so we recommend that rather than let the system generate one, you use a simple string that is easy to remember. This value cannot be updated. | [optional] 
**TypeToken** | **string** | Unique identifier of the program transfer type.  Send a &#x60;GET&#x60; request to &#x60;/programtransfers/types&#x60; to retrieve program transfer type tokens. | 
**UserToken** | Pointer to **string** | Unique identifier of the user. Pass either a &#x60;user_token&#x60; or a &#x60;business_token&#x60;, not both.  Send a &#x60;GET&#x60; request to &#x60;/users&#x60; to retrieve business tokens. | [optional] 

## Methods

### NewProgramTransfer

`func NewProgramTransfer(amount float32, currencyCode string, typeToken string, ) *ProgramTransfer`

NewProgramTransfer instantiates a new ProgramTransfer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProgramTransferWithDefaults

`func NewProgramTransferWithDefaults() *ProgramTransfer`

NewProgramTransferWithDefaults instantiates a new ProgramTransfer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *ProgramTransfer) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ProgramTransfer) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ProgramTransfer) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetBusinessToken

`func (o *ProgramTransfer) GetBusinessToken() string`

GetBusinessToken returns the BusinessToken field if non-nil, zero value otherwise.

### GetBusinessTokenOk

`func (o *ProgramTransfer) GetBusinessTokenOk() (*string, bool)`

GetBusinessTokenOk returns a tuple with the BusinessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessToken

`func (o *ProgramTransfer) SetBusinessToken(v string)`

SetBusinessToken sets BusinessToken field to given value.

### HasBusinessToken

`func (o *ProgramTransfer) HasBusinessToken() bool`

HasBusinessToken returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *ProgramTransfer) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *ProgramTransfer) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *ProgramTransfer) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.


### GetFees

`func (o *ProgramTransfer) GetFees() []FeeModel`

GetFees returns the Fees field if non-nil, zero value otherwise.

### GetFeesOk

`func (o *ProgramTransfer) GetFeesOk() (*[]FeeModel, bool)`

GetFeesOk returns a tuple with the Fees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFees

`func (o *ProgramTransfer) SetFees(v []FeeModel)`

SetFees sets Fees field to given value.

### HasFees

`func (o *ProgramTransfer) HasFees() bool`

HasFees returns a boolean if a field has been set.

### GetMemo

`func (o *ProgramTransfer) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *ProgramTransfer) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *ProgramTransfer) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *ProgramTransfer) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetTags

`func (o *ProgramTransfer) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ProgramTransfer) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ProgramTransfer) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ProgramTransfer) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetToken

`func (o *ProgramTransfer) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ProgramTransfer) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ProgramTransfer) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *ProgramTransfer) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTypeToken

`func (o *ProgramTransfer) GetTypeToken() string`

GetTypeToken returns the TypeToken field if non-nil, zero value otherwise.

### GetTypeTokenOk

`func (o *ProgramTransfer) GetTypeTokenOk() (*string, bool)`

GetTypeTokenOk returns a tuple with the TypeToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeToken

`func (o *ProgramTransfer) SetTypeToken(v string)`

SetTypeToken sets TypeToken field to given value.


### GetUserToken

`func (o *ProgramTransfer) GetUserToken() string`

GetUserToken returns the UserToken field if non-nil, zero value otherwise.

### GetUserTokenOk

`func (o *ProgramTransfer) GetUserTokenOk() (*string, bool)`

GetUserTokenOk returns a tuple with the UserToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserToken

`func (o *ProgramTransfer) SetUserToken(v string)`

SetUserToken sets UserToken field to given value.

### HasUserToken

`func (o *ProgramTransfer) HasUserToken() bool`

HasUserToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


