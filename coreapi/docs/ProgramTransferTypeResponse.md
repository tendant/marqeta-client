# ProgramTransferTypeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedTime** | Pointer to **time.Time** | The date and time when the program transfer type object was created, in UTC. | [optional] 
**LastModifiedTime** | Pointer to **time.Time** | Date and time when the program transfer type object was last modified, in UTC. | [optional] 
**Memo** | Pointer to **string** | Memo or note describing the program transfer type. | [optional] 
**ProgramFundingSourceToken** | **string** | Unique identifier of the program funding source to which program transfers will be credited. | 
**Tags** | Pointer to **string** | Comma-delimited list of tags describing the program transfer type. | [optional] 
**Token** | **string** | Unique identifier of the program transfer type request object. | 

## Methods

### NewProgramTransferTypeResponse

`func NewProgramTransferTypeResponse(programFundingSourceToken string, token string, ) *ProgramTransferTypeResponse`

NewProgramTransferTypeResponse instantiates a new ProgramTransferTypeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProgramTransferTypeResponseWithDefaults

`func NewProgramTransferTypeResponseWithDefaults() *ProgramTransferTypeResponse`

NewProgramTransferTypeResponseWithDefaults instantiates a new ProgramTransferTypeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedTime

`func (o *ProgramTransferTypeResponse) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *ProgramTransferTypeResponse) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *ProgramTransferTypeResponse) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *ProgramTransferTypeResponse) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetLastModifiedTime

`func (o *ProgramTransferTypeResponse) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *ProgramTransferTypeResponse) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *ProgramTransferTypeResponse) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.

### HasLastModifiedTime

`func (o *ProgramTransferTypeResponse) HasLastModifiedTime() bool`

HasLastModifiedTime returns a boolean if a field has been set.

### GetMemo

`func (o *ProgramTransferTypeResponse) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *ProgramTransferTypeResponse) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *ProgramTransferTypeResponse) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *ProgramTransferTypeResponse) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetProgramFundingSourceToken

`func (o *ProgramTransferTypeResponse) GetProgramFundingSourceToken() string`

GetProgramFundingSourceToken returns the ProgramFundingSourceToken field if non-nil, zero value otherwise.

### GetProgramFundingSourceTokenOk

`func (o *ProgramTransferTypeResponse) GetProgramFundingSourceTokenOk() (*string, bool)`

GetProgramFundingSourceTokenOk returns a tuple with the ProgramFundingSourceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramFundingSourceToken

`func (o *ProgramTransferTypeResponse) SetProgramFundingSourceToken(v string)`

SetProgramFundingSourceToken sets ProgramFundingSourceToken field to given value.


### GetTags

`func (o *ProgramTransferTypeResponse) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ProgramTransferTypeResponse) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ProgramTransferTypeResponse) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ProgramTransferTypeResponse) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetToken

`func (o *ProgramTransferTypeResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ProgramTransferTypeResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ProgramTransferTypeResponse) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


