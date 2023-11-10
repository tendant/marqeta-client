# ProgramFundingSourceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | **string** | The account identifier. | 
**Active** | Pointer to **bool** | Indicates whether the program funding source is active. This field is returned if it exists in the resource. | [optional] 
**CreatedTime** | **time.Time** | Date and time when the resource was created, in UTC. | 
**LastModifiedTime** | **time.Time** | Date and time when the resource was last modified, in UTC. | 
**Name** | **string** | Name of the program funding source. | 
**Token** | **string** | Unique identifier of the funding source. | 

## Methods

### NewProgramFundingSourceResponse

`func NewProgramFundingSourceResponse(account string, createdTime time.Time, lastModifiedTime time.Time, name string, token string, ) *ProgramFundingSourceResponse`

NewProgramFundingSourceResponse instantiates a new ProgramFundingSourceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProgramFundingSourceResponseWithDefaults

`func NewProgramFundingSourceResponseWithDefaults() *ProgramFundingSourceResponse`

NewProgramFundingSourceResponseWithDefaults instantiates a new ProgramFundingSourceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *ProgramFundingSourceResponse) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ProgramFundingSourceResponse) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ProgramFundingSourceResponse) SetAccount(v string)`

SetAccount sets Account field to given value.


### GetActive

`func (o *ProgramFundingSourceResponse) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ProgramFundingSourceResponse) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ProgramFundingSourceResponse) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ProgramFundingSourceResponse) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCreatedTime

`func (o *ProgramFundingSourceResponse) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *ProgramFundingSourceResponse) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *ProgramFundingSourceResponse) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetLastModifiedTime

`func (o *ProgramFundingSourceResponse) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *ProgramFundingSourceResponse) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *ProgramFundingSourceResponse) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.


### GetName

`func (o *ProgramFundingSourceResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProgramFundingSourceResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProgramFundingSourceResponse) SetName(v string)`

SetName sets Name field to given value.


### GetToken

`func (o *ProgramFundingSourceResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ProgramFundingSourceResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ProgramFundingSourceResponse) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


