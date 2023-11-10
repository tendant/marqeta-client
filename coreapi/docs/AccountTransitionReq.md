# AccountTransitionReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**AccountStatusEnum**](AccountStatusEnum.md) |  | 
**Token** | Pointer to **string** | Unique identifier of the credit account transition. | [optional] 

## Methods

### NewAccountTransitionReq

`func NewAccountTransitionReq(status AccountStatusEnum, ) *AccountTransitionReq`

NewAccountTransitionReq instantiates a new AccountTransitionReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountTransitionReqWithDefaults

`func NewAccountTransitionReqWithDefaults() *AccountTransitionReq`

NewAccountTransitionReqWithDefaults instantiates a new AccountTransitionReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AccountTransitionReq) GetStatus() AccountStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AccountTransitionReq) GetStatusOk() (*AccountStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AccountTransitionReq) SetStatus(v AccountStatusEnum)`

SetStatus sets Status field to given value.


### GetToken

`func (o *AccountTransitionReq) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AccountTransitionReq) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AccountTransitionReq) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *AccountTransitionReq) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


