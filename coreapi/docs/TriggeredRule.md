# TriggeredRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alert** | Pointer to **bool** | Indicates if the rule triggered an alert. | [optional] 
**EntityType** | Pointer to **string** | The entity type where the triggered rule was defined. | [optional] 
**RuleName** | Pointer to **string** | Name of the rule, as defined in the Real-Time Decisioning dashboard that was triggered. | [optional] 
**SuppressAlert** | Pointer to **bool** | Indicates if the triggered rule has &#x60;suppress_alert&#x60; in the definition. | [optional] 
**Tags** | Pointer to [**[]Tag**](Tag.md) | All the tags defined in the triggered rule. | [optional] 

## Methods

### NewTriggeredRule

`func NewTriggeredRule() *TriggeredRule`

NewTriggeredRule instantiates a new TriggeredRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriggeredRuleWithDefaults

`func NewTriggeredRuleWithDefaults() *TriggeredRule`

NewTriggeredRuleWithDefaults instantiates a new TriggeredRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlert

`func (o *TriggeredRule) GetAlert() bool`

GetAlert returns the Alert field if non-nil, zero value otherwise.

### GetAlertOk

`func (o *TriggeredRule) GetAlertOk() (*bool, bool)`

GetAlertOk returns a tuple with the Alert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlert

`func (o *TriggeredRule) SetAlert(v bool)`

SetAlert sets Alert field to given value.

### HasAlert

`func (o *TriggeredRule) HasAlert() bool`

HasAlert returns a boolean if a field has been set.

### GetEntityType

`func (o *TriggeredRule) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *TriggeredRule) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *TriggeredRule) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *TriggeredRule) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetRuleName

`func (o *TriggeredRule) GetRuleName() string`

GetRuleName returns the RuleName field if non-nil, zero value otherwise.

### GetRuleNameOk

`func (o *TriggeredRule) GetRuleNameOk() (*string, bool)`

GetRuleNameOk returns a tuple with the RuleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleName

`func (o *TriggeredRule) SetRuleName(v string)`

SetRuleName sets RuleName field to given value.

### HasRuleName

`func (o *TriggeredRule) HasRuleName() bool`

HasRuleName returns a boolean if a field has been set.

### GetSuppressAlert

`func (o *TriggeredRule) GetSuppressAlert() bool`

GetSuppressAlert returns the SuppressAlert field if non-nil, zero value otherwise.

### GetSuppressAlertOk

`func (o *TriggeredRule) GetSuppressAlertOk() (*bool, bool)`

GetSuppressAlertOk returns a tuple with the SuppressAlert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressAlert

`func (o *TriggeredRule) SetSuppressAlert(v bool)`

SetSuppressAlert sets SuppressAlert field to given value.

### HasSuppressAlert

`func (o *TriggeredRule) HasSuppressAlert() bool`

HasSuppressAlert returns a boolean if a field has been set.

### GetTags

`func (o *TriggeredRule) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TriggeredRule) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TriggeredRule) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *TriggeredRule) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


