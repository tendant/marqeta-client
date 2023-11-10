# CommandoModeNestedTransition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | **string** | Mechanism that changed the Commando Mode control set&#39;s state. | 
**CommandoEnabled** | **bool** | Indicates whether Commando Mode is enabled.  * If &#x60;commando_enabled&#x60; is &#x60;true&#x60; and &#x60;COMMANDO_MANUAL&#x60; is configured, all transactions are processed via Commando Mode. * If &#x60;commando_enabled&#x60; is &#x60;true&#x60; and &#x60;COMMANDO_AUTO&#x60; is configured, Commando Mode is ready to intervene only when a transaction times out or encounters an error. | [default to false]
**Reason** | Pointer to **string** | Describes the reason why the current state of the Commando Mode control set was last changed. | [optional] 
**Username** | Pointer to **string** | Identifies the user who last changed the Commando Mode control set&#39;s state. | [optional] 

## Methods

### NewCommandoModeNestedTransition

`func NewCommandoModeNestedTransition(channel string, commandoEnabled bool, ) *CommandoModeNestedTransition`

NewCommandoModeNestedTransition instantiates a new CommandoModeNestedTransition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommandoModeNestedTransitionWithDefaults

`func NewCommandoModeNestedTransitionWithDefaults() *CommandoModeNestedTransition`

NewCommandoModeNestedTransitionWithDefaults instantiates a new CommandoModeNestedTransition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *CommandoModeNestedTransition) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *CommandoModeNestedTransition) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *CommandoModeNestedTransition) SetChannel(v string)`

SetChannel sets Channel field to given value.


### GetCommandoEnabled

`func (o *CommandoModeNestedTransition) GetCommandoEnabled() bool`

GetCommandoEnabled returns the CommandoEnabled field if non-nil, zero value otherwise.

### GetCommandoEnabledOk

`func (o *CommandoModeNestedTransition) GetCommandoEnabledOk() (*bool, bool)`

GetCommandoEnabledOk returns a tuple with the CommandoEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandoEnabled

`func (o *CommandoModeNestedTransition) SetCommandoEnabled(v bool)`

SetCommandoEnabled sets CommandoEnabled field to given value.


### GetReason

`func (o *CommandoModeNestedTransition) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CommandoModeNestedTransition) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CommandoModeNestedTransition) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *CommandoModeNestedTransition) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetUsername

`func (o *CommandoModeNestedTransition) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CommandoModeNestedTransition) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CommandoModeNestedTransition) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *CommandoModeNestedTransition) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


