# PolicyRewardRuleFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | [**AmountFilter**](AmountFilter.md) |  | 
**MccDynamic** | Pointer to [**MccDynamicFilter**](MccDynamicFilter.md) |  | [optional] 

## Methods

### NewPolicyRewardRuleFilters

`func NewPolicyRewardRuleFilters(amount AmountFilter, ) *PolicyRewardRuleFilters`

NewPolicyRewardRuleFilters instantiates a new PolicyRewardRuleFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyRewardRuleFiltersWithDefaults

`func NewPolicyRewardRuleFiltersWithDefaults() *PolicyRewardRuleFilters`

NewPolicyRewardRuleFiltersWithDefaults instantiates a new PolicyRewardRuleFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *PolicyRewardRuleFilters) GetAmount() AmountFilter`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PolicyRewardRuleFilters) GetAmountOk() (*AmountFilter, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PolicyRewardRuleFilters) SetAmount(v AmountFilter)`

SetAmount sets Amount field to given value.


### GetMccDynamic

`func (o *PolicyRewardRuleFilters) GetMccDynamic() MccDynamicFilter`

GetMccDynamic returns the MccDynamic field if non-nil, zero value otherwise.

### GetMccDynamicOk

`func (o *PolicyRewardRuleFilters) GetMccDynamicOk() (*MccDynamicFilter, bool)`

GetMccDynamicOk returns a tuple with the MccDynamic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMccDynamic

`func (o *PolicyRewardRuleFilters) SetMccDynamic(v MccDynamicFilter)`

SetMccDynamic sets MccDynamic field to given value.

### HasMccDynamic

`func (o *PolicyRewardRuleFilters) HasMccDynamic() bool`

HasMccDynamic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


