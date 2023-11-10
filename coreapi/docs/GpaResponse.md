# GpaResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **float32** | Amount funded. | 
**BusinessToken** | Pointer to **string** | Unique identifier of the business.  This field is returned if it exists in the resource. | [optional] 
**CreatedTime** | **time.Time** | Date and time when the GPA order was created, in UTC. | 
**CurrencyCode** | **string** | Three-digit ISO 4217 currency code. | 
**Fees** | Pointer to [**[]FeeDetail**](FeeDetail.md) | List of fees associated with the funding transaction.  This array is returned if it exists in the resource. | [optional] 
**Funding** | [**Funding**](Funding.md) |  | 
**FundingSourceAddressToken** | Pointer to **string** | Unique identifier of the funding source address to use for this order. | [optional] 
**FundingSourceToken** | **string** | Unique identifier of the funding source to use for this order. | 
**GatewayMessage** | Pointer to **string** | Message about the status of the funding request. Useful for determining whether it was approved and completed successfully, declined by the gateway, or timed out.  This field is returned if it exists in the resource. | [optional] 
**GatewayToken** | Pointer to **int64** | Unique identifier of the JIT Funding request and response.  This field is returned if it exists in the resource. | [optional] 
**JitFunding** | Pointer to [**JitFundingApi**](JitFundingApi.md) |  | [optional] 
**LastModifiedTime** | **time.Time** | Date and time when the GPA order was last modified, in UTC. | 
**Memo** | Pointer to **string** | Additional descriptive text.  This field is returned if it exists in the resource. | [optional] 
**Response** | [**Response**](Response.md) |  | 
**State** | **string** | Current status of the funding transaction. | 
**Tags** | Pointer to **string** | Comma-delimited list of tags describing the GPA order.  This field is returned if it exists in the resource. | [optional] 
**Token** | **string** | Unique identifier of the GPA order. | 
**TransactionToken** | **string** | Unique identifier of the transaction being funded. | 
**UserToken** | Pointer to **string** | Unique identifier of the user resource.  This field is returned if it exists in the resource. | [optional] 

## Methods

### NewGpaResponse

`func NewGpaResponse(amount float32, createdTime time.Time, currencyCode string, funding Funding, fundingSourceToken string, lastModifiedTime time.Time, response Response, state string, token string, transactionToken string, ) *GpaResponse`

NewGpaResponse instantiates a new GpaResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGpaResponseWithDefaults

`func NewGpaResponseWithDefaults() *GpaResponse`

NewGpaResponseWithDefaults instantiates a new GpaResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *GpaResponse) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *GpaResponse) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *GpaResponse) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetBusinessToken

`func (o *GpaResponse) GetBusinessToken() string`

GetBusinessToken returns the BusinessToken field if non-nil, zero value otherwise.

### GetBusinessTokenOk

`func (o *GpaResponse) GetBusinessTokenOk() (*string, bool)`

GetBusinessTokenOk returns a tuple with the BusinessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessToken

`func (o *GpaResponse) SetBusinessToken(v string)`

SetBusinessToken sets BusinessToken field to given value.

### HasBusinessToken

`func (o *GpaResponse) HasBusinessToken() bool`

HasBusinessToken returns a boolean if a field has been set.

### GetCreatedTime

`func (o *GpaResponse) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *GpaResponse) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *GpaResponse) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetCurrencyCode

`func (o *GpaResponse) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *GpaResponse) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *GpaResponse) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.


### GetFees

`func (o *GpaResponse) GetFees() []FeeDetail`

GetFees returns the Fees field if non-nil, zero value otherwise.

### GetFeesOk

`func (o *GpaResponse) GetFeesOk() (*[]FeeDetail, bool)`

GetFeesOk returns a tuple with the Fees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFees

`func (o *GpaResponse) SetFees(v []FeeDetail)`

SetFees sets Fees field to given value.

### HasFees

`func (o *GpaResponse) HasFees() bool`

HasFees returns a boolean if a field has been set.

### GetFunding

`func (o *GpaResponse) GetFunding() Funding`

GetFunding returns the Funding field if non-nil, zero value otherwise.

### GetFundingOk

`func (o *GpaResponse) GetFundingOk() (*Funding, bool)`

GetFundingOk returns a tuple with the Funding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunding

`func (o *GpaResponse) SetFunding(v Funding)`

SetFunding sets Funding field to given value.


### GetFundingSourceAddressToken

`func (o *GpaResponse) GetFundingSourceAddressToken() string`

GetFundingSourceAddressToken returns the FundingSourceAddressToken field if non-nil, zero value otherwise.

### GetFundingSourceAddressTokenOk

`func (o *GpaResponse) GetFundingSourceAddressTokenOk() (*string, bool)`

GetFundingSourceAddressTokenOk returns a tuple with the FundingSourceAddressToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingSourceAddressToken

`func (o *GpaResponse) SetFundingSourceAddressToken(v string)`

SetFundingSourceAddressToken sets FundingSourceAddressToken field to given value.

### HasFundingSourceAddressToken

`func (o *GpaResponse) HasFundingSourceAddressToken() bool`

HasFundingSourceAddressToken returns a boolean if a field has been set.

### GetFundingSourceToken

`func (o *GpaResponse) GetFundingSourceToken() string`

GetFundingSourceToken returns the FundingSourceToken field if non-nil, zero value otherwise.

### GetFundingSourceTokenOk

`func (o *GpaResponse) GetFundingSourceTokenOk() (*string, bool)`

GetFundingSourceTokenOk returns a tuple with the FundingSourceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingSourceToken

`func (o *GpaResponse) SetFundingSourceToken(v string)`

SetFundingSourceToken sets FundingSourceToken field to given value.


### GetGatewayMessage

`func (o *GpaResponse) GetGatewayMessage() string`

GetGatewayMessage returns the GatewayMessage field if non-nil, zero value otherwise.

### GetGatewayMessageOk

`func (o *GpaResponse) GetGatewayMessageOk() (*string, bool)`

GetGatewayMessageOk returns a tuple with the GatewayMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayMessage

`func (o *GpaResponse) SetGatewayMessage(v string)`

SetGatewayMessage sets GatewayMessage field to given value.

### HasGatewayMessage

`func (o *GpaResponse) HasGatewayMessage() bool`

HasGatewayMessage returns a boolean if a field has been set.

### GetGatewayToken

`func (o *GpaResponse) GetGatewayToken() int64`

GetGatewayToken returns the GatewayToken field if non-nil, zero value otherwise.

### GetGatewayTokenOk

`func (o *GpaResponse) GetGatewayTokenOk() (*int64, bool)`

GetGatewayTokenOk returns a tuple with the GatewayToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayToken

`func (o *GpaResponse) SetGatewayToken(v int64)`

SetGatewayToken sets GatewayToken field to given value.

### HasGatewayToken

`func (o *GpaResponse) HasGatewayToken() bool`

HasGatewayToken returns a boolean if a field has been set.

### GetJitFunding

`func (o *GpaResponse) GetJitFunding() JitFundingApi`

GetJitFunding returns the JitFunding field if non-nil, zero value otherwise.

### GetJitFundingOk

`func (o *GpaResponse) GetJitFundingOk() (*JitFundingApi, bool)`

GetJitFundingOk returns a tuple with the JitFunding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJitFunding

`func (o *GpaResponse) SetJitFunding(v JitFundingApi)`

SetJitFunding sets JitFunding field to given value.

### HasJitFunding

`func (o *GpaResponse) HasJitFunding() bool`

HasJitFunding returns a boolean if a field has been set.

### GetLastModifiedTime

`func (o *GpaResponse) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *GpaResponse) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *GpaResponse) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.


### GetMemo

`func (o *GpaResponse) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *GpaResponse) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *GpaResponse) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *GpaResponse) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetResponse

`func (o *GpaResponse) GetResponse() Response`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *GpaResponse) GetResponseOk() (*Response, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *GpaResponse) SetResponse(v Response)`

SetResponse sets Response field to given value.


### GetState

`func (o *GpaResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GpaResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GpaResponse) SetState(v string)`

SetState sets State field to given value.


### GetTags

`func (o *GpaResponse) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GpaResponse) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GpaResponse) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GpaResponse) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetToken

`func (o *GpaResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GpaResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GpaResponse) SetToken(v string)`

SetToken sets Token field to given value.


### GetTransactionToken

`func (o *GpaResponse) GetTransactionToken() string`

GetTransactionToken returns the TransactionToken field if non-nil, zero value otherwise.

### GetTransactionTokenOk

`func (o *GpaResponse) GetTransactionTokenOk() (*string, bool)`

GetTransactionTokenOk returns a tuple with the TransactionToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionToken

`func (o *GpaResponse) SetTransactionToken(v string)`

SetTransactionToken sets TransactionToken field to given value.


### GetUserToken

`func (o *GpaResponse) GetUserToken() string`

GetUserToken returns the UserToken field if non-nil, zero value otherwise.

### GetUserTokenOk

`func (o *GpaResponse) GetUserTokenOk() (*string, bool)`

GetUserTokenOk returns a tuple with the UserToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserToken

`func (o *GpaResponse) SetUserToken(v string)`

SetUserToken sets UserToken field to given value.

### HasUserToken

`func (o *GpaResponse) HasUserToken() bool`

HasUserToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


