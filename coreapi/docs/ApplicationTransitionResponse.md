# ApplicationTransitionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationToken** | Pointer to **string** | Unique identifier of the application whose state you transitioned. | [optional] 
**CreatedTime** | Pointer to **time.Time** | Date and time when the application changed states on the Marqeta platform, in UTC. | [optional] 
**OriginalStatus** | Pointer to [**ApplicationResourceState**](ApplicationResourceState.md) |  | [optional] 
**Status** | Pointer to [**ApplicationResourceState**](ApplicationResourceState.md) |  | [optional] 
**Token** | Pointer to **string** | Unique identifier of the transition of an application&#39;s state. | [optional] 

## Methods

### NewApplicationTransitionResponse

`func NewApplicationTransitionResponse() *ApplicationTransitionResponse`

NewApplicationTransitionResponse instantiates a new ApplicationTransitionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationTransitionResponseWithDefaults

`func NewApplicationTransitionResponseWithDefaults() *ApplicationTransitionResponse`

NewApplicationTransitionResponseWithDefaults instantiates a new ApplicationTransitionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationToken

`func (o *ApplicationTransitionResponse) GetApplicationToken() string`

GetApplicationToken returns the ApplicationToken field if non-nil, zero value otherwise.

### GetApplicationTokenOk

`func (o *ApplicationTransitionResponse) GetApplicationTokenOk() (*string, bool)`

GetApplicationTokenOk returns a tuple with the ApplicationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationToken

`func (o *ApplicationTransitionResponse) SetApplicationToken(v string)`

SetApplicationToken sets ApplicationToken field to given value.

### HasApplicationToken

`func (o *ApplicationTransitionResponse) HasApplicationToken() bool`

HasApplicationToken returns a boolean if a field has been set.

### GetCreatedTime

`func (o *ApplicationTransitionResponse) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *ApplicationTransitionResponse) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *ApplicationTransitionResponse) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *ApplicationTransitionResponse) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetOriginalStatus

`func (o *ApplicationTransitionResponse) GetOriginalStatus() ApplicationResourceState`

GetOriginalStatus returns the OriginalStatus field if non-nil, zero value otherwise.

### GetOriginalStatusOk

`func (o *ApplicationTransitionResponse) GetOriginalStatusOk() (*ApplicationResourceState, bool)`

GetOriginalStatusOk returns a tuple with the OriginalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalStatus

`func (o *ApplicationTransitionResponse) SetOriginalStatus(v ApplicationResourceState)`

SetOriginalStatus sets OriginalStatus field to given value.

### HasOriginalStatus

`func (o *ApplicationTransitionResponse) HasOriginalStatus() bool`

HasOriginalStatus returns a boolean if a field has been set.

### GetStatus

`func (o *ApplicationTransitionResponse) GetStatus() ApplicationResourceState`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApplicationTransitionResponse) GetStatusOk() (*ApplicationResourceState, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApplicationTransitionResponse) SetStatus(v ApplicationResourceState)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApplicationTransitionResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetToken

`func (o *ApplicationTransitionResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ApplicationTransitionResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ApplicationTransitionResponse) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *ApplicationTransitionResponse) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


