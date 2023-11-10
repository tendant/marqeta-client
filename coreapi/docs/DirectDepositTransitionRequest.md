# DirectDepositTransitionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | **string** |  | 
**DirectDepositToken** | **string** |  | 
**IdempotentHash** | Pointer to **string** |  | [optional] 
**Reason** | **string** |  | 
**ReasonCode** | Pointer to **string** |  | [optional] 
**State** | **string** |  | 
**Token** | Pointer to **string** |  | [optional] 

## Methods

### NewDirectDepositTransitionRequest

`func NewDirectDepositTransitionRequest(channel string, directDepositToken string, reason string, state string, ) *DirectDepositTransitionRequest`

NewDirectDepositTransitionRequest instantiates a new DirectDepositTransitionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectDepositTransitionRequestWithDefaults

`func NewDirectDepositTransitionRequestWithDefaults() *DirectDepositTransitionRequest`

NewDirectDepositTransitionRequestWithDefaults instantiates a new DirectDepositTransitionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *DirectDepositTransitionRequest) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *DirectDepositTransitionRequest) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *DirectDepositTransitionRequest) SetChannel(v string)`

SetChannel sets Channel field to given value.


### GetDirectDepositToken

`func (o *DirectDepositTransitionRequest) GetDirectDepositToken() string`

GetDirectDepositToken returns the DirectDepositToken field if non-nil, zero value otherwise.

### GetDirectDepositTokenOk

`func (o *DirectDepositTransitionRequest) GetDirectDepositTokenOk() (*string, bool)`

GetDirectDepositTokenOk returns a tuple with the DirectDepositToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectDepositToken

`func (o *DirectDepositTransitionRequest) SetDirectDepositToken(v string)`

SetDirectDepositToken sets DirectDepositToken field to given value.


### GetIdempotentHash

`func (o *DirectDepositTransitionRequest) GetIdempotentHash() string`

GetIdempotentHash returns the IdempotentHash field if non-nil, zero value otherwise.

### GetIdempotentHashOk

`func (o *DirectDepositTransitionRequest) GetIdempotentHashOk() (*string, bool)`

GetIdempotentHashOk returns a tuple with the IdempotentHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotentHash

`func (o *DirectDepositTransitionRequest) SetIdempotentHash(v string)`

SetIdempotentHash sets IdempotentHash field to given value.

### HasIdempotentHash

`func (o *DirectDepositTransitionRequest) HasIdempotentHash() bool`

HasIdempotentHash returns a boolean if a field has been set.

### GetReason

`func (o *DirectDepositTransitionRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *DirectDepositTransitionRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *DirectDepositTransitionRequest) SetReason(v string)`

SetReason sets Reason field to given value.


### GetReasonCode

`func (o *DirectDepositTransitionRequest) GetReasonCode() string`

GetReasonCode returns the ReasonCode field if non-nil, zero value otherwise.

### GetReasonCodeOk

`func (o *DirectDepositTransitionRequest) GetReasonCodeOk() (*string, bool)`

GetReasonCodeOk returns a tuple with the ReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonCode

`func (o *DirectDepositTransitionRequest) SetReasonCode(v string)`

SetReasonCode sets ReasonCode field to given value.

### HasReasonCode

`func (o *DirectDepositTransitionRequest) HasReasonCode() bool`

HasReasonCode returns a boolean if a field has been set.

### GetState

`func (o *DirectDepositTransitionRequest) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DirectDepositTransitionRequest) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DirectDepositTransitionRequest) SetState(v string)`

SetState sets State field to given value.


### GetToken

`func (o *DirectDepositTransitionRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DirectDepositTransitionRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DirectDepositTransitionRequest) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DirectDepositTransitionRequest) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


