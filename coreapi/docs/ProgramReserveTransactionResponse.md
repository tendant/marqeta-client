# ProgramReserveTransactionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **float32** | Amount of the program reserve account credit or debit. Sometimes referred to as a _program funding account_. | [optional] 
**CreatedTime** | **time.Time** | Date and time when the resource was created, in UTC. | 
**CurrencyCode** | Pointer to **string** | Three-digit ISO 4217 currency code. | [optional] 
**IsCollateral** | Pointer to **bool** |  | [optional] 
**LastModifiedTime** | **time.Time** | The date and time when the resource was last modified, in UTC. | 
**Memo** | Pointer to **string** | Memo or note describing the transaction. | [optional] 
**State** | Pointer to **string** | Transaction state. | [optional] 
**Tags** | Pointer to **string** | Comma-delimited list of tags describing the transaction. | [optional] 
**Token** | Pointer to **string** | The unique identifier of the transaction response. | [optional] 
**TransactionToken** | Pointer to **string** | Unique identifier of the transaction. | [optional] 
**Type** | Pointer to **string** | Transaction type. | [optional] 

## Methods

### NewProgramReserveTransactionResponse

`func NewProgramReserveTransactionResponse(createdTime time.Time, lastModifiedTime time.Time, ) *ProgramReserveTransactionResponse`

NewProgramReserveTransactionResponse instantiates a new ProgramReserveTransactionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProgramReserveTransactionResponseWithDefaults

`func NewProgramReserveTransactionResponseWithDefaults() *ProgramReserveTransactionResponse`

NewProgramReserveTransactionResponseWithDefaults instantiates a new ProgramReserveTransactionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *ProgramReserveTransactionResponse) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ProgramReserveTransactionResponse) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ProgramReserveTransactionResponse) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *ProgramReserveTransactionResponse) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCreatedTime

`func (o *ProgramReserveTransactionResponse) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *ProgramReserveTransactionResponse) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *ProgramReserveTransactionResponse) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetCurrencyCode

`func (o *ProgramReserveTransactionResponse) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *ProgramReserveTransactionResponse) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *ProgramReserveTransactionResponse) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *ProgramReserveTransactionResponse) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetIsCollateral

`func (o *ProgramReserveTransactionResponse) GetIsCollateral() bool`

GetIsCollateral returns the IsCollateral field if non-nil, zero value otherwise.

### GetIsCollateralOk

`func (o *ProgramReserveTransactionResponse) GetIsCollateralOk() (*bool, bool)`

GetIsCollateralOk returns a tuple with the IsCollateral field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCollateral

`func (o *ProgramReserveTransactionResponse) SetIsCollateral(v bool)`

SetIsCollateral sets IsCollateral field to given value.

### HasIsCollateral

`func (o *ProgramReserveTransactionResponse) HasIsCollateral() bool`

HasIsCollateral returns a boolean if a field has been set.

### GetLastModifiedTime

`func (o *ProgramReserveTransactionResponse) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *ProgramReserveTransactionResponse) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *ProgramReserveTransactionResponse) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.


### GetMemo

`func (o *ProgramReserveTransactionResponse) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *ProgramReserveTransactionResponse) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *ProgramReserveTransactionResponse) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *ProgramReserveTransactionResponse) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetState

`func (o *ProgramReserveTransactionResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ProgramReserveTransactionResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ProgramReserveTransactionResponse) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ProgramReserveTransactionResponse) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTags

`func (o *ProgramReserveTransactionResponse) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ProgramReserveTransactionResponse) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ProgramReserveTransactionResponse) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ProgramReserveTransactionResponse) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetToken

`func (o *ProgramReserveTransactionResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ProgramReserveTransactionResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ProgramReserveTransactionResponse) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *ProgramReserveTransactionResponse) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTransactionToken

`func (o *ProgramReserveTransactionResponse) GetTransactionToken() string`

GetTransactionToken returns the TransactionToken field if non-nil, zero value otherwise.

### GetTransactionTokenOk

`func (o *ProgramReserveTransactionResponse) GetTransactionTokenOk() (*string, bool)`

GetTransactionTokenOk returns a tuple with the TransactionToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionToken

`func (o *ProgramReserveTransactionResponse) SetTransactionToken(v string)`

SetTransactionToken sets TransactionToken field to given value.

### HasTransactionToken

`func (o *ProgramReserveTransactionResponse) HasTransactionToken() bool`

HasTransactionToken returns a boolean if a field has been set.

### GetType

`func (o *ProgramReserveTransactionResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProgramReserveTransactionResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProgramReserveTransactionResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ProgramReserveTransactionResponse) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


