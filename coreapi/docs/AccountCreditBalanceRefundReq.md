# AccountCreditBalanceRefundReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **float32** | Amount of the credit balance refund.  The maximum refund amount is the amount that brings the account balance to $0. For example, $4000 is the maximum refund amount for a -$4000 account balance. | 
**CurrencyCode** | [**CurrencyCode**](CurrencyCode.md) |  | [default to CURRENCYCODE_USD]
**Description** | **string** | Description for the credit balance refund. | 
**Method** | [**RefundMethod**](RefundMethod.md) |  | 
**Token** | Pointer to **string** | Unique identifier of the credit balance refund. | [optional] 

## Methods

### NewAccountCreditBalanceRefundReq

`func NewAccountCreditBalanceRefundReq(amount float32, currencyCode CurrencyCode, description string, method RefundMethod, ) *AccountCreditBalanceRefundReq`

NewAccountCreditBalanceRefundReq instantiates a new AccountCreditBalanceRefundReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountCreditBalanceRefundReqWithDefaults

`func NewAccountCreditBalanceRefundReqWithDefaults() *AccountCreditBalanceRefundReq`

NewAccountCreditBalanceRefundReqWithDefaults instantiates a new AccountCreditBalanceRefundReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *AccountCreditBalanceRefundReq) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *AccountCreditBalanceRefundReq) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *AccountCreditBalanceRefundReq) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetCurrencyCode

`func (o *AccountCreditBalanceRefundReq) GetCurrencyCode() CurrencyCode`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *AccountCreditBalanceRefundReq) GetCurrencyCodeOk() (*CurrencyCode, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *AccountCreditBalanceRefundReq) SetCurrencyCode(v CurrencyCode)`

SetCurrencyCode sets CurrencyCode field to given value.


### GetDescription

`func (o *AccountCreditBalanceRefundReq) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AccountCreditBalanceRefundReq) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AccountCreditBalanceRefundReq) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetMethod

`func (o *AccountCreditBalanceRefundReq) GetMethod() RefundMethod`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *AccountCreditBalanceRefundReq) GetMethodOk() (*RefundMethod, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *AccountCreditBalanceRefundReq) SetMethod(v RefundMethod)`

SetMethod sets Method field to given value.


### GetToken

`func (o *AccountCreditBalanceRefundReq) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AccountCreditBalanceRefundReq) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AccountCreditBalanceRefundReq) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *AccountCreditBalanceRefundReq) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


