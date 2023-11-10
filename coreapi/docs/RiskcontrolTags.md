# RiskcontrolTags

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RuleName** | Pointer to **string** | Name of the rule, as defined in the Real-Time Decisioning dashboard, that was triggered. | [optional] 
**Tag** | Pointer to **string** | Tag name defined in the rule definition in the Real-Time Decisioning dashboard. | [optional] 
**Values** | Pointer to **[]string** | Value associated with the tag. | [optional] 

## Methods

### NewRiskcontrolTags

`func NewRiskcontrolTags() *RiskcontrolTags`

NewRiskcontrolTags instantiates a new RiskcontrolTags object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskcontrolTagsWithDefaults

`func NewRiskcontrolTagsWithDefaults() *RiskcontrolTags`

NewRiskcontrolTagsWithDefaults instantiates a new RiskcontrolTags object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRuleName

`func (o *RiskcontrolTags) GetRuleName() string`

GetRuleName returns the RuleName field if non-nil, zero value otherwise.

### GetRuleNameOk

`func (o *RiskcontrolTags) GetRuleNameOk() (*string, bool)`

GetRuleNameOk returns a tuple with the RuleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleName

`func (o *RiskcontrolTags) SetRuleName(v string)`

SetRuleName sets RuleName field to given value.

### HasRuleName

`func (o *RiskcontrolTags) HasRuleName() bool`

HasRuleName returns a boolean if a field has been set.

### GetTag

`func (o *RiskcontrolTags) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *RiskcontrolTags) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *RiskcontrolTags) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *RiskcontrolTags) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetValues

`func (o *RiskcontrolTags) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *RiskcontrolTags) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *RiskcontrolTags) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *RiskcontrolTags) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


