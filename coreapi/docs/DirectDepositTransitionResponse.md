# DirectDepositTransitionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to **string** |  | [optional] 
**CreatedTime** | Pointer to **time.Time** |  | [optional] 
**DirectDepositAccountToken** | Pointer to **string** |  | [optional] 
**DirectDepositToken** | Pointer to **string** |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**ReasonCode** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Token** | Pointer to **string** |  | [optional] 
**TransactionToken** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewDirectDepositTransitionResponse

`func NewDirectDepositTransitionResponse() *DirectDepositTransitionResponse`

NewDirectDepositTransitionResponse instantiates a new DirectDepositTransitionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectDepositTransitionResponseWithDefaults

`func NewDirectDepositTransitionResponseWithDefaults() *DirectDepositTransitionResponse`

NewDirectDepositTransitionResponseWithDefaults instantiates a new DirectDepositTransitionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *DirectDepositTransitionResponse) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *DirectDepositTransitionResponse) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *DirectDepositTransitionResponse) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *DirectDepositTransitionResponse) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetCreatedTime

`func (o *DirectDepositTransitionResponse) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *DirectDepositTransitionResponse) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *DirectDepositTransitionResponse) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *DirectDepositTransitionResponse) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDirectDepositAccountToken

`func (o *DirectDepositTransitionResponse) GetDirectDepositAccountToken() string`

GetDirectDepositAccountToken returns the DirectDepositAccountToken field if non-nil, zero value otherwise.

### GetDirectDepositAccountTokenOk

`func (o *DirectDepositTransitionResponse) GetDirectDepositAccountTokenOk() (*string, bool)`

GetDirectDepositAccountTokenOk returns a tuple with the DirectDepositAccountToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectDepositAccountToken

`func (o *DirectDepositTransitionResponse) SetDirectDepositAccountToken(v string)`

SetDirectDepositAccountToken sets DirectDepositAccountToken field to given value.

### HasDirectDepositAccountToken

`func (o *DirectDepositTransitionResponse) HasDirectDepositAccountToken() bool`

HasDirectDepositAccountToken returns a boolean if a field has been set.

### GetDirectDepositToken

`func (o *DirectDepositTransitionResponse) GetDirectDepositToken() string`

GetDirectDepositToken returns the DirectDepositToken field if non-nil, zero value otherwise.

### GetDirectDepositTokenOk

`func (o *DirectDepositTransitionResponse) GetDirectDepositTokenOk() (*string, bool)`

GetDirectDepositTokenOk returns a tuple with the DirectDepositToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectDepositToken

`func (o *DirectDepositTransitionResponse) SetDirectDepositToken(v string)`

SetDirectDepositToken sets DirectDepositToken field to given value.

### HasDirectDepositToken

`func (o *DirectDepositTransitionResponse) HasDirectDepositToken() bool`

HasDirectDepositToken returns a boolean if a field has been set.

### GetReason

`func (o *DirectDepositTransitionResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *DirectDepositTransitionResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *DirectDepositTransitionResponse) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *DirectDepositTransitionResponse) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetReasonCode

`func (o *DirectDepositTransitionResponse) GetReasonCode() string`

GetReasonCode returns the ReasonCode field if non-nil, zero value otherwise.

### GetReasonCodeOk

`func (o *DirectDepositTransitionResponse) GetReasonCodeOk() (*string, bool)`

GetReasonCodeOk returns a tuple with the ReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonCode

`func (o *DirectDepositTransitionResponse) SetReasonCode(v string)`

SetReasonCode sets ReasonCode field to given value.

### HasReasonCode

`func (o *DirectDepositTransitionResponse) HasReasonCode() bool`

HasReasonCode returns a boolean if a field has been set.

### GetState

`func (o *DirectDepositTransitionResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DirectDepositTransitionResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DirectDepositTransitionResponse) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *DirectDepositTransitionResponse) HasState() bool`

HasState returns a boolean if a field has been set.

### GetToken

`func (o *DirectDepositTransitionResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DirectDepositTransitionResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DirectDepositTransitionResponse) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DirectDepositTransitionResponse) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTransactionToken

`func (o *DirectDepositTransitionResponse) GetTransactionToken() string`

GetTransactionToken returns the TransactionToken field if non-nil, zero value otherwise.

### GetTransactionTokenOk

`func (o *DirectDepositTransitionResponse) GetTransactionTokenOk() (*string, bool)`

GetTransactionTokenOk returns a tuple with the TransactionToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionToken

`func (o *DirectDepositTransitionResponse) SetTransactionToken(v string)`

SetTransactionToken sets TransactionToken field to given value.

### HasTransactionToken

`func (o *DirectDepositTransitionResponse) HasTransactionToken() bool`

HasTransactionToken returns a boolean if a field has been set.

### GetType

`func (o *DirectDepositTransitionResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DirectDepositTransitionResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DirectDepositTransitionResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DirectDepositTransitionResponse) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


