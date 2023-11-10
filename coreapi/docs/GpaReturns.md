# GpaReturns

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **float32** | Amount of funds returned to the funding source. | 
**CreatedTime** | **time.Time** | Date and time when the GPA unload order was created, in UTC. | 
**Funding** | [**Funding**](Funding.md) |  | 
**FundingSourceAddressToken** | Pointer to **string** | Identifies the funding source used for this order. | [optional] 
**FundingSourceToken** | **string** | Identifies the funding source used for this order. | 
**JitFunding** | Pointer to [**JitFundingApi**](JitFundingApi.md) |  | [optional] 
**LastModifiedTime** | **time.Time** | Date and time when the GPA unload order was last modified, in UTC. | 
**Memo** | Pointer to **string** | Additional descriptive text. | [optional] 
**OriginalOrderToken** | Pointer to **string** | Identifies the original GPA order. | [optional] 
**Response** | [**Response**](Response.md) |  | 
**State** | **string** | Current status of the GPA unload order. | 
**Tags** | Pointer to **string** | Comma-delimited list of tags describing the GPA order. | [optional] 
**Token** | **string** | Unique identifier of the GPA unload order. | 
**TransactionToken** | **string** | Unique identifier of the transaction. | 

## Methods

### NewGpaReturns

`func NewGpaReturns(amount float32, createdTime time.Time, funding Funding, fundingSourceToken string, lastModifiedTime time.Time, response Response, state string, token string, transactionToken string, ) *GpaReturns`

NewGpaReturns instantiates a new GpaReturns object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGpaReturnsWithDefaults

`func NewGpaReturnsWithDefaults() *GpaReturns`

NewGpaReturnsWithDefaults instantiates a new GpaReturns object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *GpaReturns) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *GpaReturns) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *GpaReturns) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetCreatedTime

`func (o *GpaReturns) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *GpaReturns) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *GpaReturns) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetFunding

`func (o *GpaReturns) GetFunding() Funding`

GetFunding returns the Funding field if non-nil, zero value otherwise.

### GetFundingOk

`func (o *GpaReturns) GetFundingOk() (*Funding, bool)`

GetFundingOk returns a tuple with the Funding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunding

`func (o *GpaReturns) SetFunding(v Funding)`

SetFunding sets Funding field to given value.


### GetFundingSourceAddressToken

`func (o *GpaReturns) GetFundingSourceAddressToken() string`

GetFundingSourceAddressToken returns the FundingSourceAddressToken field if non-nil, zero value otherwise.

### GetFundingSourceAddressTokenOk

`func (o *GpaReturns) GetFundingSourceAddressTokenOk() (*string, bool)`

GetFundingSourceAddressTokenOk returns a tuple with the FundingSourceAddressToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingSourceAddressToken

`func (o *GpaReturns) SetFundingSourceAddressToken(v string)`

SetFundingSourceAddressToken sets FundingSourceAddressToken field to given value.

### HasFundingSourceAddressToken

`func (o *GpaReturns) HasFundingSourceAddressToken() bool`

HasFundingSourceAddressToken returns a boolean if a field has been set.

### GetFundingSourceToken

`func (o *GpaReturns) GetFundingSourceToken() string`

GetFundingSourceToken returns the FundingSourceToken field if non-nil, zero value otherwise.

### GetFundingSourceTokenOk

`func (o *GpaReturns) GetFundingSourceTokenOk() (*string, bool)`

GetFundingSourceTokenOk returns a tuple with the FundingSourceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingSourceToken

`func (o *GpaReturns) SetFundingSourceToken(v string)`

SetFundingSourceToken sets FundingSourceToken field to given value.


### GetJitFunding

`func (o *GpaReturns) GetJitFunding() JitFundingApi`

GetJitFunding returns the JitFunding field if non-nil, zero value otherwise.

### GetJitFundingOk

`func (o *GpaReturns) GetJitFundingOk() (*JitFundingApi, bool)`

GetJitFundingOk returns a tuple with the JitFunding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJitFunding

`func (o *GpaReturns) SetJitFunding(v JitFundingApi)`

SetJitFunding sets JitFunding field to given value.

### HasJitFunding

`func (o *GpaReturns) HasJitFunding() bool`

HasJitFunding returns a boolean if a field has been set.

### GetLastModifiedTime

`func (o *GpaReturns) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *GpaReturns) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *GpaReturns) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.


### GetMemo

`func (o *GpaReturns) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *GpaReturns) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *GpaReturns) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *GpaReturns) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetOriginalOrderToken

`func (o *GpaReturns) GetOriginalOrderToken() string`

GetOriginalOrderToken returns the OriginalOrderToken field if non-nil, zero value otherwise.

### GetOriginalOrderTokenOk

`func (o *GpaReturns) GetOriginalOrderTokenOk() (*string, bool)`

GetOriginalOrderTokenOk returns a tuple with the OriginalOrderToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalOrderToken

`func (o *GpaReturns) SetOriginalOrderToken(v string)`

SetOriginalOrderToken sets OriginalOrderToken field to given value.

### HasOriginalOrderToken

`func (o *GpaReturns) HasOriginalOrderToken() bool`

HasOriginalOrderToken returns a boolean if a field has been set.

### GetResponse

`func (o *GpaReturns) GetResponse() Response`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *GpaReturns) GetResponseOk() (*Response, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *GpaReturns) SetResponse(v Response)`

SetResponse sets Response field to given value.


### GetState

`func (o *GpaReturns) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GpaReturns) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GpaReturns) SetState(v string)`

SetState sets State field to given value.


### GetTags

`func (o *GpaReturns) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GpaReturns) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GpaReturns) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GpaReturns) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetToken

`func (o *GpaReturns) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GpaReturns) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GpaReturns) SetToken(v string)`

SetToken sets Token field to given value.


### GetTransactionToken

`func (o *GpaReturns) GetTransactionToken() string`

GetTransactionToken returns the TransactionToken field if non-nil, zero value otherwise.

### GetTransactionTokenOk

`func (o *GpaReturns) GetTransactionTokenOk() (*string, bool)`

GetTransactionTokenOk returns a tuple with the TransactionToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionToken

`func (o *GpaReturns) SetTransactionToken(v string)`

SetTransactionToken sets TransactionToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


