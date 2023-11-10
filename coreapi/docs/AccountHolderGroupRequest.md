# AccountHolderGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to [**AccountHolderGroupConfig**](AccountHolderGroupConfig.md) |  | [optional] 
**Name** | Pointer to **string** | Descriptive name for the account holder group. | [optional] 
**Token** | Pointer to **string** | Unique identifier of the account holder group. | [optional] 

## Methods

### NewAccountHolderGroupRequest

`func NewAccountHolderGroupRequest() *AccountHolderGroupRequest`

NewAccountHolderGroupRequest instantiates a new AccountHolderGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountHolderGroupRequestWithDefaults

`func NewAccountHolderGroupRequestWithDefaults() *AccountHolderGroupRequest`

NewAccountHolderGroupRequestWithDefaults instantiates a new AccountHolderGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *AccountHolderGroupRequest) GetConfig() AccountHolderGroupConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AccountHolderGroupRequest) GetConfigOk() (*AccountHolderGroupConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AccountHolderGroupRequest) SetConfig(v AccountHolderGroupConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *AccountHolderGroupRequest) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetName

`func (o *AccountHolderGroupRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountHolderGroupRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountHolderGroupRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AccountHolderGroupRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetToken

`func (o *AccountHolderGroupRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AccountHolderGroupRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AccountHolderGroupRequest) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *AccountHolderGroupRequest) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


