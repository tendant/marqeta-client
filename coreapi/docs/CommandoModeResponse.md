# CommandoModeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommandoModeEnables** | Pointer to [**CommandoModeEnables**](CommandoModeEnables.md) |  | [optional] 
**CreatedTime** | **time.Time** | Date and time when the resource was created, in UTC. | 
**CurrentState** | Pointer to [**CommandoModeNestedTransition**](CommandoModeNestedTransition.md) |  | [optional] 
**LastModifiedTime** | **time.Time** | Date and time when the resource was last updated, in UTC. | 
**ProgramGatewayFundingSourceToken** | Pointer to **string** | Unique identifier of the associated program gateway funding source. | [optional] 
**RealTimeStandinCriteria** | Pointer to [**RealTimeStandinCriteria**](RealTimeStandinCriteria.md) |  | [optional] 
**Token** | Pointer to **string** | Unique identifier of the Commando Mode control set. | [optional] 

## Methods

### NewCommandoModeResponse

`func NewCommandoModeResponse(createdTime time.Time, lastModifiedTime time.Time, ) *CommandoModeResponse`

NewCommandoModeResponse instantiates a new CommandoModeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommandoModeResponseWithDefaults

`func NewCommandoModeResponseWithDefaults() *CommandoModeResponse`

NewCommandoModeResponseWithDefaults instantiates a new CommandoModeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommandoModeEnables

`func (o *CommandoModeResponse) GetCommandoModeEnables() CommandoModeEnables`

GetCommandoModeEnables returns the CommandoModeEnables field if non-nil, zero value otherwise.

### GetCommandoModeEnablesOk

`func (o *CommandoModeResponse) GetCommandoModeEnablesOk() (*CommandoModeEnables, bool)`

GetCommandoModeEnablesOk returns a tuple with the CommandoModeEnables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandoModeEnables

`func (o *CommandoModeResponse) SetCommandoModeEnables(v CommandoModeEnables)`

SetCommandoModeEnables sets CommandoModeEnables field to given value.

### HasCommandoModeEnables

`func (o *CommandoModeResponse) HasCommandoModeEnables() bool`

HasCommandoModeEnables returns a boolean if a field has been set.

### GetCreatedTime

`func (o *CommandoModeResponse) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *CommandoModeResponse) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *CommandoModeResponse) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetCurrentState

`func (o *CommandoModeResponse) GetCurrentState() CommandoModeNestedTransition`

GetCurrentState returns the CurrentState field if non-nil, zero value otherwise.

### GetCurrentStateOk

`func (o *CommandoModeResponse) GetCurrentStateOk() (*CommandoModeNestedTransition, bool)`

GetCurrentStateOk returns a tuple with the CurrentState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentState

`func (o *CommandoModeResponse) SetCurrentState(v CommandoModeNestedTransition)`

SetCurrentState sets CurrentState field to given value.

### HasCurrentState

`func (o *CommandoModeResponse) HasCurrentState() bool`

HasCurrentState returns a boolean if a field has been set.

### GetLastModifiedTime

`func (o *CommandoModeResponse) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *CommandoModeResponse) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *CommandoModeResponse) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.


### GetProgramGatewayFundingSourceToken

`func (o *CommandoModeResponse) GetProgramGatewayFundingSourceToken() string`

GetProgramGatewayFundingSourceToken returns the ProgramGatewayFundingSourceToken field if non-nil, zero value otherwise.

### GetProgramGatewayFundingSourceTokenOk

`func (o *CommandoModeResponse) GetProgramGatewayFundingSourceTokenOk() (*string, bool)`

GetProgramGatewayFundingSourceTokenOk returns a tuple with the ProgramGatewayFundingSourceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramGatewayFundingSourceToken

`func (o *CommandoModeResponse) SetProgramGatewayFundingSourceToken(v string)`

SetProgramGatewayFundingSourceToken sets ProgramGatewayFundingSourceToken field to given value.

### HasProgramGatewayFundingSourceToken

`func (o *CommandoModeResponse) HasProgramGatewayFundingSourceToken() bool`

HasProgramGatewayFundingSourceToken returns a boolean if a field has been set.

### GetRealTimeStandinCriteria

`func (o *CommandoModeResponse) GetRealTimeStandinCriteria() RealTimeStandinCriteria`

GetRealTimeStandinCriteria returns the RealTimeStandinCriteria field if non-nil, zero value otherwise.

### GetRealTimeStandinCriteriaOk

`func (o *CommandoModeResponse) GetRealTimeStandinCriteriaOk() (*RealTimeStandinCriteria, bool)`

GetRealTimeStandinCriteriaOk returns a tuple with the RealTimeStandinCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealTimeStandinCriteria

`func (o *CommandoModeResponse) SetRealTimeStandinCriteria(v RealTimeStandinCriteria)`

SetRealTimeStandinCriteria sets RealTimeStandinCriteria field to given value.

### HasRealTimeStandinCriteria

`func (o *CommandoModeResponse) HasRealTimeStandinCriteria() bool`

HasRealTimeStandinCriteria returns a boolean if a field has been set.

### GetToken

`func (o *CommandoModeResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CommandoModeResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CommandoModeResponse) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CommandoModeResponse) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


