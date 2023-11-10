# DepositAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountNumber** | **string** |  | 
**AllowImmediateCredit** | Pointer to **bool** |  | [optional] [default to false]
**BusinessToken** | Pointer to **string** |  | [optional] 
**RoutingNumber** | **string** |  | 
**Token** | **string** |  | 
**UserToken** | Pointer to **string** |  | [optional] 

## Methods

### NewDepositAccount

`func NewDepositAccount(accountNumber string, routingNumber string, token string, ) *DepositAccount`

NewDepositAccount instantiates a new DepositAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDepositAccountWithDefaults

`func NewDepositAccountWithDefaults() *DepositAccount`

NewDepositAccountWithDefaults instantiates a new DepositAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNumber

`func (o *DepositAccount) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *DepositAccount) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *DepositAccount) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.


### GetAllowImmediateCredit

`func (o *DepositAccount) GetAllowImmediateCredit() bool`

GetAllowImmediateCredit returns the AllowImmediateCredit field if non-nil, zero value otherwise.

### GetAllowImmediateCreditOk

`func (o *DepositAccount) GetAllowImmediateCreditOk() (*bool, bool)`

GetAllowImmediateCreditOk returns a tuple with the AllowImmediateCredit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowImmediateCredit

`func (o *DepositAccount) SetAllowImmediateCredit(v bool)`

SetAllowImmediateCredit sets AllowImmediateCredit field to given value.

### HasAllowImmediateCredit

`func (o *DepositAccount) HasAllowImmediateCredit() bool`

HasAllowImmediateCredit returns a boolean if a field has been set.

### GetBusinessToken

`func (o *DepositAccount) GetBusinessToken() string`

GetBusinessToken returns the BusinessToken field if non-nil, zero value otherwise.

### GetBusinessTokenOk

`func (o *DepositAccount) GetBusinessTokenOk() (*string, bool)`

GetBusinessTokenOk returns a tuple with the BusinessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessToken

`func (o *DepositAccount) SetBusinessToken(v string)`

SetBusinessToken sets BusinessToken field to given value.

### HasBusinessToken

`func (o *DepositAccount) HasBusinessToken() bool`

HasBusinessToken returns a boolean if a field has been set.

### GetRoutingNumber

`func (o *DepositAccount) GetRoutingNumber() string`

GetRoutingNumber returns the RoutingNumber field if non-nil, zero value otherwise.

### GetRoutingNumberOk

`func (o *DepositAccount) GetRoutingNumberOk() (*string, bool)`

GetRoutingNumberOk returns a tuple with the RoutingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingNumber

`func (o *DepositAccount) SetRoutingNumber(v string)`

SetRoutingNumber sets RoutingNumber field to given value.


### GetToken

`func (o *DepositAccount) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DepositAccount) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DepositAccount) SetToken(v string)`

SetToken sets Token field to given value.


### GetUserToken

`func (o *DepositAccount) GetUserToken() string`

GetUserToken returns the UserToken field if non-nil, zero value otherwise.

### GetUserTokenOk

`func (o *DepositAccount) GetUserTokenOk() (*string, bool)`

GetUserTokenOk returns a tuple with the UserToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserToken

`func (o *DepositAccount) SetUserToken(v string)`

SetUserToken sets UserToken field to given value.

### HasUserToken

`func (o *DepositAccount) HasUserToken() bool`

HasUserToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


