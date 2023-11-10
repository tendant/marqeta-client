# DirectDepositAccountResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountNumber** | **string** |  | 
**AllowImmediateCredit** | **bool** |  | [default to false]
**BusinessToken** | **string** |  | 
**CreatedTime** | **time.Time** | yyyy-MM-ddTHH:mm:ssZ | 
**LastModifiedTime** | **time.Time** | yyyy-MM-ddTHH:mm:ssZ | 
**RoutingNumber** | **string** |  | 
**State** | **string** |  | 
**Token** | **string** |  | 
**Type** | Pointer to **string** |  | [optional] 
**UserToken** | **string** |  | 

## Methods

### NewDirectDepositAccountResponse

`func NewDirectDepositAccountResponse(accountNumber string, allowImmediateCredit bool, businessToken string, createdTime time.Time, lastModifiedTime time.Time, routingNumber string, state string, token string, userToken string, ) *DirectDepositAccountResponse`

NewDirectDepositAccountResponse instantiates a new DirectDepositAccountResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectDepositAccountResponseWithDefaults

`func NewDirectDepositAccountResponseWithDefaults() *DirectDepositAccountResponse`

NewDirectDepositAccountResponseWithDefaults instantiates a new DirectDepositAccountResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNumber

`func (o *DirectDepositAccountResponse) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *DirectDepositAccountResponse) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *DirectDepositAccountResponse) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.


### GetAllowImmediateCredit

`func (o *DirectDepositAccountResponse) GetAllowImmediateCredit() bool`

GetAllowImmediateCredit returns the AllowImmediateCredit field if non-nil, zero value otherwise.

### GetAllowImmediateCreditOk

`func (o *DirectDepositAccountResponse) GetAllowImmediateCreditOk() (*bool, bool)`

GetAllowImmediateCreditOk returns a tuple with the AllowImmediateCredit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowImmediateCredit

`func (o *DirectDepositAccountResponse) SetAllowImmediateCredit(v bool)`

SetAllowImmediateCredit sets AllowImmediateCredit field to given value.


### GetBusinessToken

`func (o *DirectDepositAccountResponse) GetBusinessToken() string`

GetBusinessToken returns the BusinessToken field if non-nil, zero value otherwise.

### GetBusinessTokenOk

`func (o *DirectDepositAccountResponse) GetBusinessTokenOk() (*string, bool)`

GetBusinessTokenOk returns a tuple with the BusinessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessToken

`func (o *DirectDepositAccountResponse) SetBusinessToken(v string)`

SetBusinessToken sets BusinessToken field to given value.


### GetCreatedTime

`func (o *DirectDepositAccountResponse) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *DirectDepositAccountResponse) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *DirectDepositAccountResponse) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetLastModifiedTime

`func (o *DirectDepositAccountResponse) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *DirectDepositAccountResponse) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *DirectDepositAccountResponse) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.


### GetRoutingNumber

`func (o *DirectDepositAccountResponse) GetRoutingNumber() string`

GetRoutingNumber returns the RoutingNumber field if non-nil, zero value otherwise.

### GetRoutingNumberOk

`func (o *DirectDepositAccountResponse) GetRoutingNumberOk() (*string, bool)`

GetRoutingNumberOk returns a tuple with the RoutingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingNumber

`func (o *DirectDepositAccountResponse) SetRoutingNumber(v string)`

SetRoutingNumber sets RoutingNumber field to given value.


### GetState

`func (o *DirectDepositAccountResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DirectDepositAccountResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DirectDepositAccountResponse) SetState(v string)`

SetState sets State field to given value.


### GetToken

`func (o *DirectDepositAccountResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DirectDepositAccountResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DirectDepositAccountResponse) SetToken(v string)`

SetToken sets Token field to given value.


### GetType

`func (o *DirectDepositAccountResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DirectDepositAccountResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DirectDepositAccountResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DirectDepositAccountResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUserToken

`func (o *DirectDepositAccountResponse) GetUserToken() string`

GetUserToken returns the UserToken field if non-nil, zero value otherwise.

### GetUserTokenOk

`func (o *DirectDepositAccountResponse) GetUserTokenOk() (*string, bool)`

GetUserTokenOk returns a tuple with the UserToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserToken

`func (o *DirectDepositAccountResponse) SetUserToken(v string)`

SetUserToken sets UserToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


