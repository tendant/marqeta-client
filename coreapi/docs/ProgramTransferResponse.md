# ProgramTransferResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **float32** | Amount of program transfer. | 
**BusinessToken** | Pointer to **string** | Unique identifier of the business account holder. Returned if &#x60;user_token&#x60; is not specified. | [optional] 
**CreatedTime** | Pointer to **time.Time** | Date and time when the program transfer object was created, in UTC. | [optional] 
**CurrencyCode** | **string** | Three-digit ISO 4217 currency code. | 
**Fees** | Pointer to [**[]FeeDetail**](FeeDetail.md) | Contains attributes that define characteristics of one or more fees. | [optional] 
**JitFunding** | Pointer to [**JitFundingApi**](JitFundingApi.md) |  | [optional] 
**Memo** | Pointer to **string** | Additional description of the program transfer. | [optional] 
**Tags** | Pointer to **string** | Comma-delimited list of tags describing the program transfer. | [optional] 
**Token** | Pointer to **string** | Unique identifier of the program transfer. | [optional] 
**TransactionToken** | **string** | Unique identifier of the transaction. | 
**TypeToken** | **string** | Unique identifier of the program transfer type. | 
**UserToken** | Pointer to **string** | Unique identifier of the user account holder. Returned if &#x60;business_token&#x60; is not specified. | [optional] 

## Methods

### NewProgramTransferResponse

`func NewProgramTransferResponse(amount float32, currencyCode string, transactionToken string, typeToken string, ) *ProgramTransferResponse`

NewProgramTransferResponse instantiates a new ProgramTransferResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProgramTransferResponseWithDefaults

`func NewProgramTransferResponseWithDefaults() *ProgramTransferResponse`

NewProgramTransferResponseWithDefaults instantiates a new ProgramTransferResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *ProgramTransferResponse) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ProgramTransferResponse) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ProgramTransferResponse) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetBusinessToken

`func (o *ProgramTransferResponse) GetBusinessToken() string`

GetBusinessToken returns the BusinessToken field if non-nil, zero value otherwise.

### GetBusinessTokenOk

`func (o *ProgramTransferResponse) GetBusinessTokenOk() (*string, bool)`

GetBusinessTokenOk returns a tuple with the BusinessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessToken

`func (o *ProgramTransferResponse) SetBusinessToken(v string)`

SetBusinessToken sets BusinessToken field to given value.

### HasBusinessToken

`func (o *ProgramTransferResponse) HasBusinessToken() bool`

HasBusinessToken returns a boolean if a field has been set.

### GetCreatedTime

`func (o *ProgramTransferResponse) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *ProgramTransferResponse) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *ProgramTransferResponse) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *ProgramTransferResponse) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *ProgramTransferResponse) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *ProgramTransferResponse) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *ProgramTransferResponse) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.


### GetFees

`func (o *ProgramTransferResponse) GetFees() []FeeDetail`

GetFees returns the Fees field if non-nil, zero value otherwise.

### GetFeesOk

`func (o *ProgramTransferResponse) GetFeesOk() (*[]FeeDetail, bool)`

GetFeesOk returns a tuple with the Fees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFees

`func (o *ProgramTransferResponse) SetFees(v []FeeDetail)`

SetFees sets Fees field to given value.

### HasFees

`func (o *ProgramTransferResponse) HasFees() bool`

HasFees returns a boolean if a field has been set.

### GetJitFunding

`func (o *ProgramTransferResponse) GetJitFunding() JitFundingApi`

GetJitFunding returns the JitFunding field if non-nil, zero value otherwise.

### GetJitFundingOk

`func (o *ProgramTransferResponse) GetJitFundingOk() (*JitFundingApi, bool)`

GetJitFundingOk returns a tuple with the JitFunding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJitFunding

`func (o *ProgramTransferResponse) SetJitFunding(v JitFundingApi)`

SetJitFunding sets JitFunding field to given value.

### HasJitFunding

`func (o *ProgramTransferResponse) HasJitFunding() bool`

HasJitFunding returns a boolean if a field has been set.

### GetMemo

`func (o *ProgramTransferResponse) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *ProgramTransferResponse) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *ProgramTransferResponse) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *ProgramTransferResponse) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetTags

`func (o *ProgramTransferResponse) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ProgramTransferResponse) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ProgramTransferResponse) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ProgramTransferResponse) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetToken

`func (o *ProgramTransferResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ProgramTransferResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ProgramTransferResponse) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *ProgramTransferResponse) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTransactionToken

`func (o *ProgramTransferResponse) GetTransactionToken() string`

GetTransactionToken returns the TransactionToken field if non-nil, zero value otherwise.

### GetTransactionTokenOk

`func (o *ProgramTransferResponse) GetTransactionTokenOk() (*string, bool)`

GetTransactionTokenOk returns a tuple with the TransactionToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionToken

`func (o *ProgramTransferResponse) SetTransactionToken(v string)`

SetTransactionToken sets TransactionToken field to given value.


### GetTypeToken

`func (o *ProgramTransferResponse) GetTypeToken() string`

GetTypeToken returns the TypeToken field if non-nil, zero value otherwise.

### GetTypeTokenOk

`func (o *ProgramTransferResponse) GetTypeTokenOk() (*string, bool)`

GetTypeTokenOk returns a tuple with the TypeToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeToken

`func (o *ProgramTransferResponse) SetTypeToken(v string)`

SetTypeToken sets TypeToken field to given value.


### GetUserToken

`func (o *ProgramTransferResponse) GetUserToken() string`

GetUserToken returns the UserToken field if non-nil, zero value otherwise.

### GetUserTokenOk

`func (o *ProgramTransferResponse) GetUserTokenOk() (*string, bool)`

GetUserTokenOk returns a tuple with the UserToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserToken

`func (o *ProgramTransferResponse) SetUserToken(v string)`

SetUserToken sets UserToken field to given value.

### HasUserToken

`func (o *ProgramTransferResponse) HasUserToken() bool`

HasUserToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


