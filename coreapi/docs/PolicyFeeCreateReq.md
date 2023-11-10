# PolicyFeeCreateReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | [**PolicyFeeAccount**](PolicyFeeAccount.md) |  | 
**Description** | Pointer to **string** | Description of the fee policy. | [optional] 
**Name** | **string** | Name of the fee policy. | 
**Token** | Pointer to **string** | Unique identifier of the fee policy. | [optional] 

## Methods

### NewPolicyFeeCreateReq

`func NewPolicyFeeCreateReq(account PolicyFeeAccount, name string, ) *PolicyFeeCreateReq`

NewPolicyFeeCreateReq instantiates a new PolicyFeeCreateReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyFeeCreateReqWithDefaults

`func NewPolicyFeeCreateReqWithDefaults() *PolicyFeeCreateReq`

NewPolicyFeeCreateReqWithDefaults instantiates a new PolicyFeeCreateReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *PolicyFeeCreateReq) GetAccount() PolicyFeeAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *PolicyFeeCreateReq) GetAccountOk() (*PolicyFeeAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *PolicyFeeCreateReq) SetAccount(v PolicyFeeAccount)`

SetAccount sets Account field to given value.


### GetDescription

`func (o *PolicyFeeCreateReq) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PolicyFeeCreateReq) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PolicyFeeCreateReq) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PolicyFeeCreateReq) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *PolicyFeeCreateReq) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PolicyFeeCreateReq) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PolicyFeeCreateReq) SetName(v string)`

SetName sets Name field to given value.


### GetToken

`func (o *PolicyFeeCreateReq) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PolicyFeeCreateReq) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PolicyFeeCreateReq) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *PolicyFeeCreateReq) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


