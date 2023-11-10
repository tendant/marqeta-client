# JitFundingApi

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActingUserToken** | Pointer to **string** | User who conducted the transaction.  Can be a child user configured to share its parent&#39;s account balance. | [optional] 
**AddressVerification** | Pointer to [**JitAddressVerification**](JitAddressVerification.md) |  | [optional] 
**Amount** | **float32** | Requested amount of funding. | 
**Balances** | Pointer to [**map[string]CardholderBalance**](CardholderBalance.md) | Contains the GPA&#39;s balance details. | [optional] 
**BusinessToken** | Pointer to **string** | Holder of the business account that was funded. | [optional] 
**DeclineReason** | Pointer to **string** | Reason why the transaction was declined. | [optional] 
**IncrementalAuthorizationJitFundingTokens** | Pointer to **[]string** | Array of tokens referencing the JIT Funding tokens of all previous associated incremental authorization JIT Funding requests. Useful for ascertaining the final transaction amount when the original amount was incremented. | [optional] 
**Memo** | Pointer to **string** | Additional information that describes the JIT Funding transaction. | [optional] 
**Method** | **string** | JIT Funding response type. See &lt;&lt;/core-api/gateway-jit-funding-messages#_the_jit_funding_object, The jit_funding object&gt;&gt; for the purpose, funding event type, and description of each method. | 
**OriginalJitFundingToken** | Pointer to **string** | Unique identifier of the first associated JIT Funding message. Useful for correlating related JIT Funding messages (that is, those associated with the same GPA order). Not included in the first of any set of related messages. | [optional] 
**Tags** | Pointer to **string** | Customer-defined tags related to the JIT Funding transaction. | [optional] 
**Token** | **string** | Existing JIT Funding token matching the &#x60;funding.gateway_log.transaction_id&#x60; field of the associated GPA order.  *NOTE:* The &#x60;transaction_id&#x60; field updates if a subsequent JIT Funding message associated with that GPA order is sent. If multiple JIT Funding messages are associated with the same GPA order, the &#x60;transaction_id&#x60; field matches the token of the most recent message. | 
**UserToken** | **string** | Holder of the user account that was funded. | 

## Methods

### NewJitFundingApi

`func NewJitFundingApi(amount float32, method string, token string, userToken string, ) *JitFundingApi`

NewJitFundingApi instantiates a new JitFundingApi object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJitFundingApiWithDefaults

`func NewJitFundingApiWithDefaults() *JitFundingApi`

NewJitFundingApiWithDefaults instantiates a new JitFundingApi object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActingUserToken

`func (o *JitFundingApi) GetActingUserToken() string`

GetActingUserToken returns the ActingUserToken field if non-nil, zero value otherwise.

### GetActingUserTokenOk

`func (o *JitFundingApi) GetActingUserTokenOk() (*string, bool)`

GetActingUserTokenOk returns a tuple with the ActingUserToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActingUserToken

`func (o *JitFundingApi) SetActingUserToken(v string)`

SetActingUserToken sets ActingUserToken field to given value.

### HasActingUserToken

`func (o *JitFundingApi) HasActingUserToken() bool`

HasActingUserToken returns a boolean if a field has been set.

### GetAddressVerification

`func (o *JitFundingApi) GetAddressVerification() JitAddressVerification`

GetAddressVerification returns the AddressVerification field if non-nil, zero value otherwise.

### GetAddressVerificationOk

`func (o *JitFundingApi) GetAddressVerificationOk() (*JitAddressVerification, bool)`

GetAddressVerificationOk returns a tuple with the AddressVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressVerification

`func (o *JitFundingApi) SetAddressVerification(v JitAddressVerification)`

SetAddressVerification sets AddressVerification field to given value.

### HasAddressVerification

`func (o *JitFundingApi) HasAddressVerification() bool`

HasAddressVerification returns a boolean if a field has been set.

### GetAmount

`func (o *JitFundingApi) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *JitFundingApi) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *JitFundingApi) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetBalances

`func (o *JitFundingApi) GetBalances() map[string]CardholderBalance`

GetBalances returns the Balances field if non-nil, zero value otherwise.

### GetBalancesOk

`func (o *JitFundingApi) GetBalancesOk() (*map[string]CardholderBalance, bool)`

GetBalancesOk returns a tuple with the Balances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalances

`func (o *JitFundingApi) SetBalances(v map[string]CardholderBalance)`

SetBalances sets Balances field to given value.

### HasBalances

`func (o *JitFundingApi) HasBalances() bool`

HasBalances returns a boolean if a field has been set.

### GetBusinessToken

`func (o *JitFundingApi) GetBusinessToken() string`

GetBusinessToken returns the BusinessToken field if non-nil, zero value otherwise.

### GetBusinessTokenOk

`func (o *JitFundingApi) GetBusinessTokenOk() (*string, bool)`

GetBusinessTokenOk returns a tuple with the BusinessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessToken

`func (o *JitFundingApi) SetBusinessToken(v string)`

SetBusinessToken sets BusinessToken field to given value.

### HasBusinessToken

`func (o *JitFundingApi) HasBusinessToken() bool`

HasBusinessToken returns a boolean if a field has been set.

### GetDeclineReason

`func (o *JitFundingApi) GetDeclineReason() string`

GetDeclineReason returns the DeclineReason field if non-nil, zero value otherwise.

### GetDeclineReasonOk

`func (o *JitFundingApi) GetDeclineReasonOk() (*string, bool)`

GetDeclineReasonOk returns a tuple with the DeclineReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclineReason

`func (o *JitFundingApi) SetDeclineReason(v string)`

SetDeclineReason sets DeclineReason field to given value.

### HasDeclineReason

`func (o *JitFundingApi) HasDeclineReason() bool`

HasDeclineReason returns a boolean if a field has been set.

### GetIncrementalAuthorizationJitFundingTokens

`func (o *JitFundingApi) GetIncrementalAuthorizationJitFundingTokens() []string`

GetIncrementalAuthorizationJitFundingTokens returns the IncrementalAuthorizationJitFundingTokens field if non-nil, zero value otherwise.

### GetIncrementalAuthorizationJitFundingTokensOk

`func (o *JitFundingApi) GetIncrementalAuthorizationJitFundingTokensOk() (*[]string, bool)`

GetIncrementalAuthorizationJitFundingTokensOk returns a tuple with the IncrementalAuthorizationJitFundingTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncrementalAuthorizationJitFundingTokens

`func (o *JitFundingApi) SetIncrementalAuthorizationJitFundingTokens(v []string)`

SetIncrementalAuthorizationJitFundingTokens sets IncrementalAuthorizationJitFundingTokens field to given value.

### HasIncrementalAuthorizationJitFundingTokens

`func (o *JitFundingApi) HasIncrementalAuthorizationJitFundingTokens() bool`

HasIncrementalAuthorizationJitFundingTokens returns a boolean if a field has been set.

### GetMemo

`func (o *JitFundingApi) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *JitFundingApi) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *JitFundingApi) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *JitFundingApi) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetMethod

`func (o *JitFundingApi) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *JitFundingApi) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *JitFundingApi) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetOriginalJitFundingToken

`func (o *JitFundingApi) GetOriginalJitFundingToken() string`

GetOriginalJitFundingToken returns the OriginalJitFundingToken field if non-nil, zero value otherwise.

### GetOriginalJitFundingTokenOk

`func (o *JitFundingApi) GetOriginalJitFundingTokenOk() (*string, bool)`

GetOriginalJitFundingTokenOk returns a tuple with the OriginalJitFundingToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalJitFundingToken

`func (o *JitFundingApi) SetOriginalJitFundingToken(v string)`

SetOriginalJitFundingToken sets OriginalJitFundingToken field to given value.

### HasOriginalJitFundingToken

`func (o *JitFundingApi) HasOriginalJitFundingToken() bool`

HasOriginalJitFundingToken returns a boolean if a field has been set.

### GetTags

`func (o *JitFundingApi) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *JitFundingApi) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *JitFundingApi) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *JitFundingApi) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetToken

`func (o *JitFundingApi) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *JitFundingApi) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *JitFundingApi) SetToken(v string)`

SetToken sets Token field to given value.


### GetUserToken

`func (o *JitFundingApi) GetUserToken() string`

GetUserToken returns the UserToken field if non-nil, zero value otherwise.

### GetUserTokenOk

`func (o *JitFundingApi) GetUserTokenOk() (*string, bool)`

GetUserTokenOk returns a tuple with the UserToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserToken

`func (o *JitFundingApi) SetUserToken(v string)`

SetUserToken sets UserToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


