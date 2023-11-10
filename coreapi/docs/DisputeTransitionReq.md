# DisputeTransitionReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **float32** | Updated amount of the dispute, based on the resolution. | 
**Notes** | Pointer to **string** | Additional information on the dispute update (for example, a reason for the dispute update). | [optional] 
**Status** | [**DisputeStatus**](DisputeStatus.md) |  | 
**Token** | Pointer to **string** | Unique identifier of the dispute update. | [optional] 

## Methods

### NewDisputeTransitionReq

`func NewDisputeTransitionReq(amount float32, status DisputeStatus, ) *DisputeTransitionReq`

NewDisputeTransitionReq instantiates a new DisputeTransitionReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDisputeTransitionReqWithDefaults

`func NewDisputeTransitionReqWithDefaults() *DisputeTransitionReq`

NewDisputeTransitionReqWithDefaults instantiates a new DisputeTransitionReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *DisputeTransitionReq) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *DisputeTransitionReq) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *DisputeTransitionReq) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetNotes

`func (o *DisputeTransitionReq) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *DisputeTransitionReq) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *DisputeTransitionReq) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *DisputeTransitionReq) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetStatus

`func (o *DisputeTransitionReq) GetStatus() DisputeStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DisputeTransitionReq) GetStatusOk() (*DisputeStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DisputeTransitionReq) SetStatus(v DisputeStatus)`

SetStatus sets Status field to given value.


### GetToken

`func (o *DisputeTransitionReq) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DisputeTransitionReq) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DisputeTransitionReq) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DisputeTransitionReq) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


