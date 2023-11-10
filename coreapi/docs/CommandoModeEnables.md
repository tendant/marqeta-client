# CommandoModeEnables

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthControls** | Pointer to **[]string** | Unique identifiers of the authorization controls enabled while in Commando Mode. | [optional] 
**IgnoreCardSuspendedState** | Pointer to **bool** | If set to &#x60;true&#x60;, transactions conducted while Commando Mode is enabled proceed even when the card is suspended. If set to &#x60;false&#x60;, transactions conducted while Commando Mode is enabled are declined if the card is suspended. | [optional] [default to false]
**ProgramFundingSource** | **string** | Unique identifier of the program funding source that substitutes for the program gateway funding source upon Commando Mode enablement. | 
**UseCacheBalance** | Pointer to **bool** | This field is not currently in use. | [optional] [default to false]
**VelocityControls** | Pointer to **[]string** | Unique identifiers of the velocity controls enabled while in Commando Mode.  Velocity controls that are enabled in Commando Mode are inactive until a Commando Mode event occurs. When Commando Mode velocity controls are activated, they conform to the &#x60;velocity_window&#x60; specified in that velocity control. For example, a &#x60;velocity_window&#x60; of &#x60;DAY&#x60; is one calendar day starting at 00:00:00 UTC. If a Commando Mode event occurs at 11:59:59 UTC, the &#x60;DAY&#x60; window includes all transactions that occurred between 00:00:00 and 11:59:59 on that calendar day. | [optional] 

## Methods

### NewCommandoModeEnables

`func NewCommandoModeEnables(programFundingSource string, ) *CommandoModeEnables`

NewCommandoModeEnables instantiates a new CommandoModeEnables object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommandoModeEnablesWithDefaults

`func NewCommandoModeEnablesWithDefaults() *CommandoModeEnables`

NewCommandoModeEnablesWithDefaults instantiates a new CommandoModeEnables object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthControls

`func (o *CommandoModeEnables) GetAuthControls() []string`

GetAuthControls returns the AuthControls field if non-nil, zero value otherwise.

### GetAuthControlsOk

`func (o *CommandoModeEnables) GetAuthControlsOk() (*[]string, bool)`

GetAuthControlsOk returns a tuple with the AuthControls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthControls

`func (o *CommandoModeEnables) SetAuthControls(v []string)`

SetAuthControls sets AuthControls field to given value.

### HasAuthControls

`func (o *CommandoModeEnables) HasAuthControls() bool`

HasAuthControls returns a boolean if a field has been set.

### GetIgnoreCardSuspendedState

`func (o *CommandoModeEnables) GetIgnoreCardSuspendedState() bool`

GetIgnoreCardSuspendedState returns the IgnoreCardSuspendedState field if non-nil, zero value otherwise.

### GetIgnoreCardSuspendedStateOk

`func (o *CommandoModeEnables) GetIgnoreCardSuspendedStateOk() (*bool, bool)`

GetIgnoreCardSuspendedStateOk returns a tuple with the IgnoreCardSuspendedState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreCardSuspendedState

`func (o *CommandoModeEnables) SetIgnoreCardSuspendedState(v bool)`

SetIgnoreCardSuspendedState sets IgnoreCardSuspendedState field to given value.

### HasIgnoreCardSuspendedState

`func (o *CommandoModeEnables) HasIgnoreCardSuspendedState() bool`

HasIgnoreCardSuspendedState returns a boolean if a field has been set.

### GetProgramFundingSource

`func (o *CommandoModeEnables) GetProgramFundingSource() string`

GetProgramFundingSource returns the ProgramFundingSource field if non-nil, zero value otherwise.

### GetProgramFundingSourceOk

`func (o *CommandoModeEnables) GetProgramFundingSourceOk() (*string, bool)`

GetProgramFundingSourceOk returns a tuple with the ProgramFundingSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramFundingSource

`func (o *CommandoModeEnables) SetProgramFundingSource(v string)`

SetProgramFundingSource sets ProgramFundingSource field to given value.


### GetUseCacheBalance

`func (o *CommandoModeEnables) GetUseCacheBalance() bool`

GetUseCacheBalance returns the UseCacheBalance field if non-nil, zero value otherwise.

### GetUseCacheBalanceOk

`func (o *CommandoModeEnables) GetUseCacheBalanceOk() (*bool, bool)`

GetUseCacheBalanceOk returns a tuple with the UseCacheBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCacheBalance

`func (o *CommandoModeEnables) SetUseCacheBalance(v bool)`

SetUseCacheBalance sets UseCacheBalance field to given value.

### HasUseCacheBalance

`func (o *CommandoModeEnables) HasUseCacheBalance() bool`

HasUseCacheBalance returns a boolean if a field has been set.

### GetVelocityControls

`func (o *CommandoModeEnables) GetVelocityControls() []string`

GetVelocityControls returns the VelocityControls field if non-nil, zero value otherwise.

### GetVelocityControlsOk

`func (o *CommandoModeEnables) GetVelocityControlsOk() (*[]string, bool)`

GetVelocityControlsOk returns a tuple with the VelocityControls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVelocityControls

`func (o *CommandoModeEnables) SetVelocityControls(v []string)`

SetVelocityControls sets VelocityControls field to given value.

### HasVelocityControls

`func (o *CommandoModeEnables) HasVelocityControls() bool`

HasVelocityControls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


