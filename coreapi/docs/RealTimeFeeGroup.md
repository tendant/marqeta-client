# RealTimeFeeGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | **bool** | Indicates whether the real-time fee group is active. | [default to false]
**CreatedTime** | Pointer to **time.Time** | Date and time when the real-time fee group was created, in UTC. | [optional] 
**FeeTokens** | Pointer to **[]string** | Specifies the fees in this real-time fee group. | [optional] 
**LastModifiedTime** | Pointer to **time.Time** | Date and time when the real-time fee group was last modified, in UTC. | [optional] 
**Name** | **string** | Descriptive name for the real-time fee group. | 
**Token** | **string** | Unique identifier of the real-time fee group. | 

## Methods

### NewRealTimeFeeGroup

`func NewRealTimeFeeGroup(active bool, name string, token string, ) *RealTimeFeeGroup`

NewRealTimeFeeGroup instantiates a new RealTimeFeeGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRealTimeFeeGroupWithDefaults

`func NewRealTimeFeeGroupWithDefaults() *RealTimeFeeGroup`

NewRealTimeFeeGroupWithDefaults instantiates a new RealTimeFeeGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *RealTimeFeeGroup) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *RealTimeFeeGroup) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *RealTimeFeeGroup) SetActive(v bool)`

SetActive sets Active field to given value.


### GetCreatedTime

`func (o *RealTimeFeeGroup) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *RealTimeFeeGroup) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *RealTimeFeeGroup) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *RealTimeFeeGroup) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetFeeTokens

`func (o *RealTimeFeeGroup) GetFeeTokens() []string`

GetFeeTokens returns the FeeTokens field if non-nil, zero value otherwise.

### GetFeeTokensOk

`func (o *RealTimeFeeGroup) GetFeeTokensOk() (*[]string, bool)`

GetFeeTokensOk returns a tuple with the FeeTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeTokens

`func (o *RealTimeFeeGroup) SetFeeTokens(v []string)`

SetFeeTokens sets FeeTokens field to given value.

### HasFeeTokens

`func (o *RealTimeFeeGroup) HasFeeTokens() bool`

HasFeeTokens returns a boolean if a field has been set.

### GetLastModifiedTime

`func (o *RealTimeFeeGroup) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *RealTimeFeeGroup) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *RealTimeFeeGroup) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.

### HasLastModifiedTime

`func (o *RealTimeFeeGroup) HasLastModifiedTime() bool`

HasLastModifiedTime returns a boolean if a field has been set.

### GetName

`func (o *RealTimeFeeGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RealTimeFeeGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RealTimeFeeGroup) SetName(v string)`

SetName sets Name field to given value.


### GetToken

`func (o *RealTimeFeeGroup) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *RealTimeFeeGroup) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *RealTimeFeeGroup) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


