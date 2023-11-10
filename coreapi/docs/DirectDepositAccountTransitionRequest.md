# DirectDepositAccountTransitionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountToken** | **string** |  | 
**Channel** | **string** |  | 
**Reason** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Token** | Pointer to **string** |  | [optional] 

## Methods

### NewDirectDepositAccountTransitionRequest

`func NewDirectDepositAccountTransitionRequest(accountToken string, channel string, ) *DirectDepositAccountTransitionRequest`

NewDirectDepositAccountTransitionRequest instantiates a new DirectDepositAccountTransitionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectDepositAccountTransitionRequestWithDefaults

`func NewDirectDepositAccountTransitionRequestWithDefaults() *DirectDepositAccountTransitionRequest`

NewDirectDepositAccountTransitionRequestWithDefaults instantiates a new DirectDepositAccountTransitionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountToken

`func (o *DirectDepositAccountTransitionRequest) GetAccountToken() string`

GetAccountToken returns the AccountToken field if non-nil, zero value otherwise.

### GetAccountTokenOk

`func (o *DirectDepositAccountTransitionRequest) GetAccountTokenOk() (*string, bool)`

GetAccountTokenOk returns a tuple with the AccountToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountToken

`func (o *DirectDepositAccountTransitionRequest) SetAccountToken(v string)`

SetAccountToken sets AccountToken field to given value.


### GetChannel

`func (o *DirectDepositAccountTransitionRequest) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *DirectDepositAccountTransitionRequest) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *DirectDepositAccountTransitionRequest) SetChannel(v string)`

SetChannel sets Channel field to given value.


### GetReason

`func (o *DirectDepositAccountTransitionRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *DirectDepositAccountTransitionRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *DirectDepositAccountTransitionRequest) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *DirectDepositAccountTransitionRequest) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetState

`func (o *DirectDepositAccountTransitionRequest) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DirectDepositAccountTransitionRequest) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DirectDepositAccountTransitionRequest) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *DirectDepositAccountTransitionRequest) HasState() bool`

HasState returns a boolean if a field has been set.

### GetToken

`func (o *DirectDepositAccountTransitionRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DirectDepositAccountTransitionRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DirectDepositAccountTransitionRequest) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DirectDepositAccountTransitionRequest) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


