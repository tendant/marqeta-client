# ChargebackResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **float32** | Amount of the chargeback. | 
**Channel** | **string** | Channel the chargeback came through. | 
**CreatedTime** | **time.Time** | Date and time when the chargeback was created. Not returned for transactions when the associated chargeback is in the &#x60;INITIATED&#x60; state. | 
**CreditUser** | **bool** | Whether to credit the user for the chargeback amount. | [default to false]
**LastModifiedTime** | **time.Time** | Date and time when the chargeback was last modified. Not returned for transactions when the associated chargeback is in the &#x60;INITIATED&#x60; state. | 
**Memo** | Pointer to **string** | Additional comments about the chargeback. | [optional] 
**Network** | **string** | Network handling the chargeback. | 
**NetworkCaseId** | Pointer to **string** | Network-assigned identifier of the chargeback. | [optional] 
**ReasonCode** | Pointer to **string** | Identifies the standardized reason for the chargeback. | [optional] 
**State** | **string** | State of the case. | 
**Token** | **string** | Unique identifier of the chargeback. | 
**TransactionToken** | **string** | Unique identifier of the transaction being charged back. | 

## Methods

### NewChargebackResponse

`func NewChargebackResponse(amount float32, channel string, createdTime time.Time, creditUser bool, lastModifiedTime time.Time, network string, state string, token string, transactionToken string, ) *ChargebackResponse`

NewChargebackResponse instantiates a new ChargebackResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChargebackResponseWithDefaults

`func NewChargebackResponseWithDefaults() *ChargebackResponse`

NewChargebackResponseWithDefaults instantiates a new ChargebackResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *ChargebackResponse) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ChargebackResponse) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ChargebackResponse) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetChannel

`func (o *ChargebackResponse) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *ChargebackResponse) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *ChargebackResponse) SetChannel(v string)`

SetChannel sets Channel field to given value.


### GetCreatedTime

`func (o *ChargebackResponse) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *ChargebackResponse) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *ChargebackResponse) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetCreditUser

`func (o *ChargebackResponse) GetCreditUser() bool`

GetCreditUser returns the CreditUser field if non-nil, zero value otherwise.

### GetCreditUserOk

`func (o *ChargebackResponse) GetCreditUserOk() (*bool, bool)`

GetCreditUserOk returns a tuple with the CreditUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditUser

`func (o *ChargebackResponse) SetCreditUser(v bool)`

SetCreditUser sets CreditUser field to given value.


### GetLastModifiedTime

`func (o *ChargebackResponse) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *ChargebackResponse) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *ChargebackResponse) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.


### GetMemo

`func (o *ChargebackResponse) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *ChargebackResponse) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *ChargebackResponse) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *ChargebackResponse) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetNetwork

`func (o *ChargebackResponse) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *ChargebackResponse) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *ChargebackResponse) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetNetworkCaseId

`func (o *ChargebackResponse) GetNetworkCaseId() string`

GetNetworkCaseId returns the NetworkCaseId field if non-nil, zero value otherwise.

### GetNetworkCaseIdOk

`func (o *ChargebackResponse) GetNetworkCaseIdOk() (*string, bool)`

GetNetworkCaseIdOk returns a tuple with the NetworkCaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkCaseId

`func (o *ChargebackResponse) SetNetworkCaseId(v string)`

SetNetworkCaseId sets NetworkCaseId field to given value.

### HasNetworkCaseId

`func (o *ChargebackResponse) HasNetworkCaseId() bool`

HasNetworkCaseId returns a boolean if a field has been set.

### GetReasonCode

`func (o *ChargebackResponse) GetReasonCode() string`

GetReasonCode returns the ReasonCode field if non-nil, zero value otherwise.

### GetReasonCodeOk

`func (o *ChargebackResponse) GetReasonCodeOk() (*string, bool)`

GetReasonCodeOk returns a tuple with the ReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonCode

`func (o *ChargebackResponse) SetReasonCode(v string)`

SetReasonCode sets ReasonCode field to given value.

### HasReasonCode

`func (o *ChargebackResponse) HasReasonCode() bool`

HasReasonCode returns a boolean if a field has been set.

### GetState

`func (o *ChargebackResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ChargebackResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ChargebackResponse) SetState(v string)`

SetState sets State field to given value.


### GetToken

`func (o *ChargebackResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ChargebackResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ChargebackResponse) SetToken(v string)`

SetToken sets Token field to given value.


### GetTransactionToken

`func (o *ChargebackResponse) GetTransactionToken() string`

GetTransactionToken returns the TransactionToken field if non-nil, zero value otherwise.

### GetTransactionTokenOk

`func (o *ChargebackResponse) GetTransactionTokenOk() (*string, bool)`

GetTransactionTokenOk returns a tuple with the TransactionToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionToken

`func (o *ChargebackResponse) SetTransactionToken(v string)`

SetTransactionToken sets TransactionToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


