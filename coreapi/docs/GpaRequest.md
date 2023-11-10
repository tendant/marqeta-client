# GpaRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **float32** | Amount to fund. | 
**BusinessToken** | Pointer to **string** | Unique identifier of the business.  Pass either a &#x60;business_token&#x60; or a &#x60;user_token&#x60;, not both.  Send a &#x60;GET&#x60; request to &#x60;/businesses&#x60; to retrieve business tokens. | [optional] 
**CurrencyCode** | **string** | Three-digit ISO 4217 currency code. | 
**Fees** | Pointer to [**[]FeeModel**](FeeModel.md) | List of fees associated with the funding transaction. | [optional] 
**FundingSourceAddressToken** | Pointer to **string** | Unique identifier of the funding source address to use for this order. If your funding source is an ACH account, then a funding source address is not required. If your funding source is a payment card, you must have at least one funding source address in order to create a GPA order. Send a &#x60;GET&#x60; request to &#x60;/fundingsources/addresses/user/{token}&#x60; to retrieve addresses for a specific user. | [optional] 
**FundingSourceToken** | **string** | Unique identifier of the funding source to use for this order.  You do not have to supply a funding source token value in this call if you have a default funding source set up (verify the funding source&#39;s &#x60;is_default_account&#x60; field). If you have only one funding source, then this source is used as the default. If you have multiple funding sources and none are configured as the default, then an error is returned.  Send a &#x60;GET&#x60; request to &#x60;/fundingsources/user/{user_token}&#x60; to retrieve funding source tokens for a user or to &#x60;/fundingsources/business/{business_token}&#x60; to retrieve funding source tokens for a business. | 
**Memo** | Pointer to **string** | Additional descriptive text. | [optional] 
**Tags** | Pointer to **string** | Comma-delimited list of tags describing the GPA order. | [optional] 
**Token** | Pointer to **string** | Unique identifier of the GPA order.  If you do not include a token, the system will generate one automatically. This token is necessary for use in other calls, so we recommend that rather than let the system generate one, you use a simple string that is easy to remember. This value cannot be updated. | [optional] 
**UserToken** | Pointer to **string** | Unique identifier of the user.  Pass either a &#x60;user_token&#x60; or a &#x60;business_token&#x60;, not both.  Send a &#x60;GET&#x60; request to &#x60;/users&#x60; to retrieve business tokens. | [optional] 

## Methods

### NewGpaRequest

`func NewGpaRequest(amount float32, currencyCode string, fundingSourceToken string, ) *GpaRequest`

NewGpaRequest instantiates a new GpaRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGpaRequestWithDefaults

`func NewGpaRequestWithDefaults() *GpaRequest`

NewGpaRequestWithDefaults instantiates a new GpaRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *GpaRequest) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *GpaRequest) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *GpaRequest) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetBusinessToken

`func (o *GpaRequest) GetBusinessToken() string`

GetBusinessToken returns the BusinessToken field if non-nil, zero value otherwise.

### GetBusinessTokenOk

`func (o *GpaRequest) GetBusinessTokenOk() (*string, bool)`

GetBusinessTokenOk returns a tuple with the BusinessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessToken

`func (o *GpaRequest) SetBusinessToken(v string)`

SetBusinessToken sets BusinessToken field to given value.

### HasBusinessToken

`func (o *GpaRequest) HasBusinessToken() bool`

HasBusinessToken returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *GpaRequest) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *GpaRequest) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *GpaRequest) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.


### GetFees

`func (o *GpaRequest) GetFees() []FeeModel`

GetFees returns the Fees field if non-nil, zero value otherwise.

### GetFeesOk

`func (o *GpaRequest) GetFeesOk() (*[]FeeModel, bool)`

GetFeesOk returns a tuple with the Fees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFees

`func (o *GpaRequest) SetFees(v []FeeModel)`

SetFees sets Fees field to given value.

### HasFees

`func (o *GpaRequest) HasFees() bool`

HasFees returns a boolean if a field has been set.

### GetFundingSourceAddressToken

`func (o *GpaRequest) GetFundingSourceAddressToken() string`

GetFundingSourceAddressToken returns the FundingSourceAddressToken field if non-nil, zero value otherwise.

### GetFundingSourceAddressTokenOk

`func (o *GpaRequest) GetFundingSourceAddressTokenOk() (*string, bool)`

GetFundingSourceAddressTokenOk returns a tuple with the FundingSourceAddressToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingSourceAddressToken

`func (o *GpaRequest) SetFundingSourceAddressToken(v string)`

SetFundingSourceAddressToken sets FundingSourceAddressToken field to given value.

### HasFundingSourceAddressToken

`func (o *GpaRequest) HasFundingSourceAddressToken() bool`

HasFundingSourceAddressToken returns a boolean if a field has been set.

### GetFundingSourceToken

`func (o *GpaRequest) GetFundingSourceToken() string`

GetFundingSourceToken returns the FundingSourceToken field if non-nil, zero value otherwise.

### GetFundingSourceTokenOk

`func (o *GpaRequest) GetFundingSourceTokenOk() (*string, bool)`

GetFundingSourceTokenOk returns a tuple with the FundingSourceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingSourceToken

`func (o *GpaRequest) SetFundingSourceToken(v string)`

SetFundingSourceToken sets FundingSourceToken field to given value.


### GetMemo

`func (o *GpaRequest) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *GpaRequest) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *GpaRequest) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *GpaRequest) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetTags

`func (o *GpaRequest) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GpaRequest) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GpaRequest) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GpaRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetToken

`func (o *GpaRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GpaRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GpaRequest) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GpaRequest) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUserToken

`func (o *GpaRequest) GetUserToken() string`

GetUserToken returns the UserToken field if non-nil, zero value otherwise.

### GetUserTokenOk

`func (o *GpaRequest) GetUserTokenOk() (*string, bool)`

GetUserTokenOk returns a tuple with the UserToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserToken

`func (o *GpaRequest) SetUserToken(v string)`

SetUserToken sets UserToken field to given value.

### HasUserToken

`func (o *GpaRequest) HasUserToken() bool`

HasUserToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


