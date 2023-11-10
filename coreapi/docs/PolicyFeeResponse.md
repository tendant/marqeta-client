# PolicyFeeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to [**PolicyFeeAccount**](PolicyFeeAccount.md) |  | [optional] 
**CreatedTime** | Pointer to **time.Time** | Date and time when the fee policy was created on Marqeta&#39;s credit platform, in UTC. | [optional] 
**Description** | Pointer to **string** | Description of the fee policy. | [optional] 
**Name** | Pointer to **string** | Name of the fee policy. | [optional] 
**Token** | Pointer to **string** | Unique identifier of the fee policy. | [optional] 
**UpdatedTime** | Pointer to **time.Time** | Date and time when the fee policy was last updated on Marqeta&#39;s credit platform, in UTC. | [optional] 

## Methods

### NewPolicyFeeResponse

`func NewPolicyFeeResponse() *PolicyFeeResponse`

NewPolicyFeeResponse instantiates a new PolicyFeeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyFeeResponseWithDefaults

`func NewPolicyFeeResponseWithDefaults() *PolicyFeeResponse`

NewPolicyFeeResponseWithDefaults instantiates a new PolicyFeeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *PolicyFeeResponse) GetAccount() PolicyFeeAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *PolicyFeeResponse) GetAccountOk() (*PolicyFeeAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *PolicyFeeResponse) SetAccount(v PolicyFeeAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *PolicyFeeResponse) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetCreatedTime

`func (o *PolicyFeeResponse) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *PolicyFeeResponse) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *PolicyFeeResponse) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *PolicyFeeResponse) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDescription

`func (o *PolicyFeeResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PolicyFeeResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PolicyFeeResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PolicyFeeResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *PolicyFeeResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PolicyFeeResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PolicyFeeResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PolicyFeeResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetToken

`func (o *PolicyFeeResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PolicyFeeResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PolicyFeeResponse) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *PolicyFeeResponse) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUpdatedTime

`func (o *PolicyFeeResponse) GetUpdatedTime() time.Time`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *PolicyFeeResponse) GetUpdatedTimeOk() (*time.Time, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *PolicyFeeResponse) SetUpdatedTime(v time.Time)`

SetUpdatedTime sets UpdatedTime field to given value.

### HasUpdatedTime

`func (o *PolicyFeeResponse) HasUpdatedTime() bool`

HasUpdatedTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


