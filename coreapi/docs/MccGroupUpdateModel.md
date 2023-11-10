# MccGroupUpdateModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Indicates whether the MCC group is active or inactive. | [optional] [default to false]
**Config** | Pointer to [**Config**](Config.md) |  | [optional] 
**Mccs** | Pointer to **[]string** | The set of merchant category codes that you want to include in this group. For each element, valid characters are 0-9, and the length must be 4 digits. You can also specify a range like \&quot;9876-9880\&quot;. An MCC can belong to more than one group.  Updating the merchant category codes for the group completely replaces the group&#39;s existing codes. For example, if the current MCC group is &#x60;[\&quot;1234\&quot;]&#x60; and you want to add the 2345 code (while retaining the existing code), you must specify &#x60;[\&quot;1234\&quot;, \&quot;2345\&quot;]&#x60; in this field. | [optional] 
**Name** | Pointer to **string** | The name of the MCC group. | [optional] 

## Methods

### NewMccGroupUpdateModel

`func NewMccGroupUpdateModel() *MccGroupUpdateModel`

NewMccGroupUpdateModel instantiates a new MccGroupUpdateModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMccGroupUpdateModelWithDefaults

`func NewMccGroupUpdateModelWithDefaults() *MccGroupUpdateModel`

NewMccGroupUpdateModelWithDefaults instantiates a new MccGroupUpdateModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *MccGroupUpdateModel) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *MccGroupUpdateModel) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *MccGroupUpdateModel) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *MccGroupUpdateModel) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetConfig

`func (o *MccGroupUpdateModel) GetConfig() Config`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *MccGroupUpdateModel) GetConfigOk() (*Config, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *MccGroupUpdateModel) SetConfig(v Config)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *MccGroupUpdateModel) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetMccs

`func (o *MccGroupUpdateModel) GetMccs() []string`

GetMccs returns the Mccs field if non-nil, zero value otherwise.

### GetMccsOk

`func (o *MccGroupUpdateModel) GetMccsOk() (*[]string, bool)`

GetMccsOk returns a tuple with the Mccs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMccs

`func (o *MccGroupUpdateModel) SetMccs(v []string)`

SetMccs sets Mccs field to given value.

### HasMccs

`func (o *MccGroupUpdateModel) HasMccs() bool`

HasMccs returns a boolean if a field has been set.

### GetName

`func (o *MccGroupUpdateModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MccGroupUpdateModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MccGroupUpdateModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MccGroupUpdateModel) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


