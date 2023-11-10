# AccountHolderGroupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to [**AccountHolderGroupConfig**](AccountHolderGroupConfig.md) |  | [optional] 
**Name** | Pointer to **string** | Descriptive name for the account holder group.  This field is returned if it exists in the resource. | [optional] 
**Token** | Pointer to **string** | Unique identifier of the account holder group.  This field is always returned. | [optional] 

## Methods

### NewAccountHolderGroupResponse

`func NewAccountHolderGroupResponse() *AccountHolderGroupResponse`

NewAccountHolderGroupResponse instantiates a new AccountHolderGroupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountHolderGroupResponseWithDefaults

`func NewAccountHolderGroupResponseWithDefaults() *AccountHolderGroupResponse`

NewAccountHolderGroupResponseWithDefaults instantiates a new AccountHolderGroupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *AccountHolderGroupResponse) GetConfig() AccountHolderGroupConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AccountHolderGroupResponse) GetConfigOk() (*AccountHolderGroupConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AccountHolderGroupResponse) SetConfig(v AccountHolderGroupConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *AccountHolderGroupResponse) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetName

`func (o *AccountHolderGroupResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountHolderGroupResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountHolderGroupResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AccountHolderGroupResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetToken

`func (o *AccountHolderGroupResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AccountHolderGroupResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AccountHolderGroupResponse) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *AccountHolderGroupResponse) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


