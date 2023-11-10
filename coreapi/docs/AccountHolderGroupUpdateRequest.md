# AccountHolderGroupUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to [**AccountHolderGroupConfig**](AccountHolderGroupConfig.md) |  | [optional] 
**Name** | Pointer to **string** | Descriptive name for the account holder group. | [optional] 

## Methods

### NewAccountHolderGroupUpdateRequest

`func NewAccountHolderGroupUpdateRequest() *AccountHolderGroupUpdateRequest`

NewAccountHolderGroupUpdateRequest instantiates a new AccountHolderGroupUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountHolderGroupUpdateRequestWithDefaults

`func NewAccountHolderGroupUpdateRequestWithDefaults() *AccountHolderGroupUpdateRequest`

NewAccountHolderGroupUpdateRequestWithDefaults instantiates a new AccountHolderGroupUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *AccountHolderGroupUpdateRequest) GetConfig() AccountHolderGroupConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AccountHolderGroupUpdateRequest) GetConfigOk() (*AccountHolderGroupConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AccountHolderGroupUpdateRequest) SetConfig(v AccountHolderGroupConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *AccountHolderGroupUpdateRequest) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetName

`func (o *AccountHolderGroupUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountHolderGroupUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountHolderGroupUpdateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AccountHolderGroupUpdateRequest) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


