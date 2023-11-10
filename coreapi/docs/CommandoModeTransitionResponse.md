# CommandoModeTransitionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommandoModeToken** | Pointer to **string** | Unique identifier of the Commando Mode control set. | [optional] 
**CreatedTime** | **time.Time** | Date and time when the resource was created, in UTC. | 
**Name** | Pointer to **string** | Identifies the user who changed the Commando Mode control set&#39;s state. | [optional] 
**Token** | Pointer to **string** | Unique identifier of the Command Mode control set transition object. | [optional] 
**Transition** | Pointer to [**CommandoModeNestedTransition**](CommandoModeNestedTransition.md) |  | [optional] 
**Type** | Pointer to **string** | Specifies the type of event that triggered the Commando Mode transition, such as a &#x60;connection_error&#x60; or &#x60;response_timeout&#x60;. | [optional] 

## Methods

### NewCommandoModeTransitionResponse

`func NewCommandoModeTransitionResponse(createdTime time.Time, ) *CommandoModeTransitionResponse`

NewCommandoModeTransitionResponse instantiates a new CommandoModeTransitionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommandoModeTransitionResponseWithDefaults

`func NewCommandoModeTransitionResponseWithDefaults() *CommandoModeTransitionResponse`

NewCommandoModeTransitionResponseWithDefaults instantiates a new CommandoModeTransitionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommandoModeToken

`func (o *CommandoModeTransitionResponse) GetCommandoModeToken() string`

GetCommandoModeToken returns the CommandoModeToken field if non-nil, zero value otherwise.

### GetCommandoModeTokenOk

`func (o *CommandoModeTransitionResponse) GetCommandoModeTokenOk() (*string, bool)`

GetCommandoModeTokenOk returns a tuple with the CommandoModeToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandoModeToken

`func (o *CommandoModeTransitionResponse) SetCommandoModeToken(v string)`

SetCommandoModeToken sets CommandoModeToken field to given value.

### HasCommandoModeToken

`func (o *CommandoModeTransitionResponse) HasCommandoModeToken() bool`

HasCommandoModeToken returns a boolean if a field has been set.

### GetCreatedTime

`func (o *CommandoModeTransitionResponse) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *CommandoModeTransitionResponse) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *CommandoModeTransitionResponse) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetName

`func (o *CommandoModeTransitionResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CommandoModeTransitionResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CommandoModeTransitionResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CommandoModeTransitionResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetToken

`func (o *CommandoModeTransitionResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CommandoModeTransitionResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CommandoModeTransitionResponse) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CommandoModeTransitionResponse) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTransition

`func (o *CommandoModeTransitionResponse) GetTransition() CommandoModeNestedTransition`

GetTransition returns the Transition field if non-nil, zero value otherwise.

### GetTransitionOk

`func (o *CommandoModeTransitionResponse) GetTransitionOk() (*CommandoModeNestedTransition, bool)`

GetTransitionOk returns a tuple with the Transition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransition

`func (o *CommandoModeTransitionResponse) SetTransition(v CommandoModeNestedTransition)`

SetTransition sets Transition field to given value.

### HasTransition

`func (o *CommandoModeTransitionResponse) HasTransition() bool`

HasTransition returns a boolean if a field has been set.

### GetType

`func (o *CommandoModeTransitionResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CommandoModeTransitionResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CommandoModeTransitionResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CommandoModeTransitionResponse) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


