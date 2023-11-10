# PolicyFeeUpdateReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to [**PolicyFeeAccount**](PolicyFeeAccount.md) |  | [optional] 
**Description** | Pointer to **string** | Description of the fee policy. | [optional] 
**Name** | **string** | Name of the fee policy. | 
**Token** | Pointer to **string** | Unique identifier of the fee policy. | [optional] 

## Methods

### NewPolicyFeeUpdateReq

`func NewPolicyFeeUpdateReq(name string, ) *PolicyFeeUpdateReq`

NewPolicyFeeUpdateReq instantiates a new PolicyFeeUpdateReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyFeeUpdateReqWithDefaults

`func NewPolicyFeeUpdateReqWithDefaults() *PolicyFeeUpdateReq`

NewPolicyFeeUpdateReqWithDefaults instantiates a new PolicyFeeUpdateReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *PolicyFeeUpdateReq) GetAccount() PolicyFeeAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *PolicyFeeUpdateReq) GetAccountOk() (*PolicyFeeAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *PolicyFeeUpdateReq) SetAccount(v PolicyFeeAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *PolicyFeeUpdateReq) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetDescription

`func (o *PolicyFeeUpdateReq) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PolicyFeeUpdateReq) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PolicyFeeUpdateReq) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PolicyFeeUpdateReq) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *PolicyFeeUpdateReq) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PolicyFeeUpdateReq) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PolicyFeeUpdateReq) SetName(v string)`

SetName sets Name field to given value.


### GetToken

`func (o *PolicyFeeUpdateReq) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PolicyFeeUpdateReq) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PolicyFeeUpdateReq) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *PolicyFeeUpdateReq) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


