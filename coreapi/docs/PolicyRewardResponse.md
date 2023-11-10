# PolicyRewardResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedTime** | Pointer to **time.Time** | Date and time when the reward policy was created on Marqeta&#39;s credit platform, in UTC. | [optional] 
**Description** | Pointer to **string** | Description of the reward policy. | [optional] 
**Name** | Pointer to **string** | Name of the reward policy. | [optional] 
**Rules** | Pointer to [**[]PolicyRewardRule**](PolicyRewardRule.md) | One or more reward rules | [optional] 
**Token** | Pointer to **string** | Unique identifier of the reward policy. | [optional] 
**UpdatedTime** | Pointer to **time.Time** | Date and time when the reward policy was last updated on Marqeta&#39;s credit platform, in UTC. | [optional] 

## Methods

### NewPolicyRewardResponse

`func NewPolicyRewardResponse() *PolicyRewardResponse`

NewPolicyRewardResponse instantiates a new PolicyRewardResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyRewardResponseWithDefaults

`func NewPolicyRewardResponseWithDefaults() *PolicyRewardResponse`

NewPolicyRewardResponseWithDefaults instantiates a new PolicyRewardResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedTime

`func (o *PolicyRewardResponse) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *PolicyRewardResponse) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *PolicyRewardResponse) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *PolicyRewardResponse) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDescription

`func (o *PolicyRewardResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PolicyRewardResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PolicyRewardResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PolicyRewardResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *PolicyRewardResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PolicyRewardResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PolicyRewardResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PolicyRewardResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRules

`func (o *PolicyRewardResponse) GetRules() []PolicyRewardRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *PolicyRewardResponse) GetRulesOk() (*[]PolicyRewardRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *PolicyRewardResponse) SetRules(v []PolicyRewardRule)`

SetRules sets Rules field to given value.

### HasRules

`func (o *PolicyRewardResponse) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetToken

`func (o *PolicyRewardResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PolicyRewardResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PolicyRewardResponse) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *PolicyRewardResponse) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUpdatedTime

`func (o *PolicyRewardResponse) GetUpdatedTime() time.Time`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *PolicyRewardResponse) GetUpdatedTimeOk() (*time.Time, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *PolicyRewardResponse) SetUpdatedTime(v time.Time)`

SetUpdatedTime sets UpdatedTime field to given value.

### HasUpdatedTime

`func (o *PolicyRewardResponse) HasUpdatedTime() bool`

HasUpdatedTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


