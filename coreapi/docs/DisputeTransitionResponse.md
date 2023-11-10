# DisputeTransitionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountToken** | **string** | Unique identifier of the credit account on which the dispute was updated. | 
**Amount** | **float32** | Amount of the updated dispute, based on the resolution. | 
**CreatedTime** | **time.Time** | Date and time when the dispute update was created on Marqeta&#39;s credit platform, in UTC. | 
**Notes** | Pointer to **string** | Additional information on the dispute update (for example, a reason for the dispute update). | [optional] 
**Status** | [**DisputeStatus**](DisputeStatus.md) |  | 
**Token** | **string** | Unique identifier of the dispute update. | 

## Methods

### NewDisputeTransitionResponse

`func NewDisputeTransitionResponse(accountToken string, amount float32, createdTime time.Time, status DisputeStatus, token string, ) *DisputeTransitionResponse`

NewDisputeTransitionResponse instantiates a new DisputeTransitionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDisputeTransitionResponseWithDefaults

`func NewDisputeTransitionResponseWithDefaults() *DisputeTransitionResponse`

NewDisputeTransitionResponseWithDefaults instantiates a new DisputeTransitionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountToken

`func (o *DisputeTransitionResponse) GetAccountToken() string`

GetAccountToken returns the AccountToken field if non-nil, zero value otherwise.

### GetAccountTokenOk

`func (o *DisputeTransitionResponse) GetAccountTokenOk() (*string, bool)`

GetAccountTokenOk returns a tuple with the AccountToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountToken

`func (o *DisputeTransitionResponse) SetAccountToken(v string)`

SetAccountToken sets AccountToken field to given value.


### GetAmount

`func (o *DisputeTransitionResponse) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *DisputeTransitionResponse) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *DisputeTransitionResponse) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetCreatedTime

`func (o *DisputeTransitionResponse) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *DisputeTransitionResponse) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *DisputeTransitionResponse) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetNotes

`func (o *DisputeTransitionResponse) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *DisputeTransitionResponse) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *DisputeTransitionResponse) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *DisputeTransitionResponse) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetStatus

`func (o *DisputeTransitionResponse) GetStatus() DisputeStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DisputeTransitionResponse) GetStatusOk() (*DisputeStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DisputeTransitionResponse) SetStatus(v DisputeStatus)`

SetStatus sets Status field to given value.


### GetToken

`func (o *DisputeTransitionResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DisputeTransitionResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DisputeTransitionResponse) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


