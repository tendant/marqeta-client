# MccGroupModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Indicates if the group is active or inactive. | [optional] [default to false]
**Config** | Pointer to [**Config**](Config.md) |  | [optional] 
**Mccs** | **[]map[string]interface{}** | The set of merchant category codes that you want to include in this group. For each element, valid characters are 0-9, and the length must be 4 digits. You can also specify a range like \&quot;9876-9880\&quot;. An MCC can belong to more than one group. | 
**Name** | **string** | The name of the group. | 
**Token** | Pointer to **string** | The unique identifier of the group.  If you do not include a token, the system will generate one automatically. This token is necessary for use in other API calls, so we recommend that rather than let the system generate one, you use a simple string that is easy to remember. This value cannot be updated. | [optional] 

## Methods

### NewMccGroupModel

`func NewMccGroupModel(mccs []map[string]interface{}, name string, ) *MccGroupModel`

NewMccGroupModel instantiates a new MccGroupModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMccGroupModelWithDefaults

`func NewMccGroupModelWithDefaults() *MccGroupModel`

NewMccGroupModelWithDefaults instantiates a new MccGroupModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *MccGroupModel) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *MccGroupModel) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *MccGroupModel) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *MccGroupModel) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetConfig

`func (o *MccGroupModel) GetConfig() Config`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *MccGroupModel) GetConfigOk() (*Config, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *MccGroupModel) SetConfig(v Config)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *MccGroupModel) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetMccs

`func (o *MccGroupModel) GetMccs() []map[string]interface{}`

GetMccs returns the Mccs field if non-nil, zero value otherwise.

### GetMccsOk

`func (o *MccGroupModel) GetMccsOk() (*[]map[string]interface{}, bool)`

GetMccsOk returns a tuple with the Mccs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMccs

`func (o *MccGroupModel) SetMccs(v []map[string]interface{})`

SetMccs sets Mccs field to given value.


### GetName

`func (o *MccGroupModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MccGroupModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MccGroupModel) SetName(v string)`

SetName sets Name field to given value.


### GetToken

`func (o *MccGroupModel) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *MccGroupModel) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *MccGroupModel) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *MccGroupModel) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


