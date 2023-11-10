# AccountTransitionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountToken** | **string** | Unique identifier of the credit account for which to transition a status. | 
**CreatedTime** | **time.Time** | Date and time when the transition record was created on Marqeta&#39;s credit platform, in UTC. | 
**OriginalStatus** | [**AccountStatusEnum**](AccountStatusEnum.md) |  | 
**Status** | [**AccountStatusEnum**](AccountStatusEnum.md) |  | 
**Token** | **string** | Unique identifier of the credit account transition. | 

## Methods

### NewAccountTransitionResponse

`func NewAccountTransitionResponse(accountToken string, createdTime time.Time, originalStatus AccountStatusEnum, status AccountStatusEnum, token string, ) *AccountTransitionResponse`

NewAccountTransitionResponse instantiates a new AccountTransitionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountTransitionResponseWithDefaults

`func NewAccountTransitionResponseWithDefaults() *AccountTransitionResponse`

NewAccountTransitionResponseWithDefaults instantiates a new AccountTransitionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountToken

`func (o *AccountTransitionResponse) GetAccountToken() string`

GetAccountToken returns the AccountToken field if non-nil, zero value otherwise.

### GetAccountTokenOk

`func (o *AccountTransitionResponse) GetAccountTokenOk() (*string, bool)`

GetAccountTokenOk returns a tuple with the AccountToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountToken

`func (o *AccountTransitionResponse) SetAccountToken(v string)`

SetAccountToken sets AccountToken field to given value.


### GetCreatedTime

`func (o *AccountTransitionResponse) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *AccountTransitionResponse) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *AccountTransitionResponse) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetOriginalStatus

`func (o *AccountTransitionResponse) GetOriginalStatus() AccountStatusEnum`

GetOriginalStatus returns the OriginalStatus field if non-nil, zero value otherwise.

### GetOriginalStatusOk

`func (o *AccountTransitionResponse) GetOriginalStatusOk() (*AccountStatusEnum, bool)`

GetOriginalStatusOk returns a tuple with the OriginalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalStatus

`func (o *AccountTransitionResponse) SetOriginalStatus(v AccountStatusEnum)`

SetOriginalStatus sets OriginalStatus field to given value.


### GetStatus

`func (o *AccountTransitionResponse) GetStatus() AccountStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AccountTransitionResponse) GetStatusOk() (*AccountStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AccountTransitionResponse) SetStatus(v AccountStatusEnum)`

SetStatus sets Status field to given value.


### GetToken

`func (o *AccountTransitionResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AccountTransitionResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AccountTransitionResponse) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


