# AccountCreditBalanceRefundResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountToken** | **string** | Unique identifier of the credit account that needs the credit balance refund. | 
**Amount** | **float32** | Amount of the credit balance refund.  The maximum refund amount is the amount that brings the account balance to $0. For example, $4000 is the maximum refund amount for a -$4000 account balance. | 
**CreatedTime** | **time.Time** | Date and time when the credit balance refund was created. | 
**Description** | **string** | Description for the credit credit balance refund. | 
**Method** | [**RefundMethod**](RefundMethod.md) |  | 
**Status** | [**PaymentStatus**](PaymentStatus.md) |  | 
**Token** | **string** | Unique identifier of the credit balance refund.  If in the &#x60;detail_object&#x60;, unique identifier of the detail object. | 
**UpdatedTime** | **time.Time** | Date and time when the credit balance refund was last updated. | 

## Methods

### NewAccountCreditBalanceRefundResponse

`func NewAccountCreditBalanceRefundResponse(accountToken string, amount float32, createdTime time.Time, description string, method RefundMethod, status PaymentStatus, token string, updatedTime time.Time, ) *AccountCreditBalanceRefundResponse`

NewAccountCreditBalanceRefundResponse instantiates a new AccountCreditBalanceRefundResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountCreditBalanceRefundResponseWithDefaults

`func NewAccountCreditBalanceRefundResponseWithDefaults() *AccountCreditBalanceRefundResponse`

NewAccountCreditBalanceRefundResponseWithDefaults instantiates a new AccountCreditBalanceRefundResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountToken

`func (o *AccountCreditBalanceRefundResponse) GetAccountToken() string`

GetAccountToken returns the AccountToken field if non-nil, zero value otherwise.

### GetAccountTokenOk

`func (o *AccountCreditBalanceRefundResponse) GetAccountTokenOk() (*string, bool)`

GetAccountTokenOk returns a tuple with the AccountToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountToken

`func (o *AccountCreditBalanceRefundResponse) SetAccountToken(v string)`

SetAccountToken sets AccountToken field to given value.


### GetAmount

`func (o *AccountCreditBalanceRefundResponse) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *AccountCreditBalanceRefundResponse) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *AccountCreditBalanceRefundResponse) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetCreatedTime

`func (o *AccountCreditBalanceRefundResponse) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *AccountCreditBalanceRefundResponse) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *AccountCreditBalanceRefundResponse) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetDescription

`func (o *AccountCreditBalanceRefundResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AccountCreditBalanceRefundResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AccountCreditBalanceRefundResponse) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetMethod

`func (o *AccountCreditBalanceRefundResponse) GetMethod() RefundMethod`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *AccountCreditBalanceRefundResponse) GetMethodOk() (*RefundMethod, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *AccountCreditBalanceRefundResponse) SetMethod(v RefundMethod)`

SetMethod sets Method field to given value.


### GetStatus

`func (o *AccountCreditBalanceRefundResponse) GetStatus() PaymentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AccountCreditBalanceRefundResponse) GetStatusOk() (*PaymentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AccountCreditBalanceRefundResponse) SetStatus(v PaymentStatus)`

SetStatus sets Status field to given value.


### GetToken

`func (o *AccountCreditBalanceRefundResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AccountCreditBalanceRefundResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AccountCreditBalanceRefundResponse) SetToken(v string)`

SetToken sets Token field to given value.


### GetUpdatedTime

`func (o *AccountCreditBalanceRefundResponse) GetUpdatedTime() time.Time`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *AccountCreditBalanceRefundResponse) GetUpdatedTimeOk() (*time.Time, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *AccountCreditBalanceRefundResponse) SetUpdatedTime(v time.Time)`

SetUpdatedTime sets UpdatedTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


