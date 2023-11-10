# PolicyRewardReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description of the reward policy. | [optional] 
**Name** | **string** | Name of the reward policy. | 
**Rules** | [**[]PolicyRewardRule**](PolicyRewardRule.md) | One or more reward rules. | 
**Token** | Pointer to **string** | Unique identifier of the reward policy. | [optional] 

## Methods

### NewPolicyRewardReq

`func NewPolicyRewardReq(name string, rules []PolicyRewardRule, ) *PolicyRewardReq`

NewPolicyRewardReq instantiates a new PolicyRewardReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyRewardReqWithDefaults

`func NewPolicyRewardReqWithDefaults() *PolicyRewardReq`

NewPolicyRewardReqWithDefaults instantiates a new PolicyRewardReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *PolicyRewardReq) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PolicyRewardReq) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PolicyRewardReq) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PolicyRewardReq) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *PolicyRewardReq) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PolicyRewardReq) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PolicyRewardReq) SetName(v string)`

SetName sets Name field to given value.


### GetRules

`func (o *PolicyRewardReq) GetRules() []PolicyRewardRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *PolicyRewardReq) GetRulesOk() (*[]PolicyRewardRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *PolicyRewardReq) SetRules(v []PolicyRewardRule)`

SetRules sets Rules field to given value.


### GetToken

`func (o *PolicyRewardReq) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PolicyRewardReq) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PolicyRewardReq) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *PolicyRewardReq) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


