# ProgramTransferTypeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Memo** | **string** | Memo or note describing the program transfer type. | 
**ProgramFundingSourceToken** | **string** | Unique identifier of the program funding source to which program transfers will be credited.  Send a &#x60;GET&#x60; request to &#x60;/fundingsources/program&#x60; to retrieve program funding source tokens. | 
**Tags** | Pointer to **string** | Comma-delimited list of tags describing the program transfer type. | [optional] 
**Token** | Pointer to **string** | Unique identifier of the program transfer type.  If you do not include a token, the system will generate one automatically. This token is necessary for use in other API calls, so we recommend that rather than let the system generate one, you use a simple string that is easy to remember. This value cannot be updated. | [optional] 

## Methods

### NewProgramTransferTypeRequest

`func NewProgramTransferTypeRequest(memo string, programFundingSourceToken string, ) *ProgramTransferTypeRequest`

NewProgramTransferTypeRequest instantiates a new ProgramTransferTypeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProgramTransferTypeRequestWithDefaults

`func NewProgramTransferTypeRequestWithDefaults() *ProgramTransferTypeRequest`

NewProgramTransferTypeRequestWithDefaults instantiates a new ProgramTransferTypeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMemo

`func (o *ProgramTransferTypeRequest) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *ProgramTransferTypeRequest) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *ProgramTransferTypeRequest) SetMemo(v string)`

SetMemo sets Memo field to given value.


### GetProgramFundingSourceToken

`func (o *ProgramTransferTypeRequest) GetProgramFundingSourceToken() string`

GetProgramFundingSourceToken returns the ProgramFundingSourceToken field if non-nil, zero value otherwise.

### GetProgramFundingSourceTokenOk

`func (o *ProgramTransferTypeRequest) GetProgramFundingSourceTokenOk() (*string, bool)`

GetProgramFundingSourceTokenOk returns a tuple with the ProgramFundingSourceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramFundingSourceToken

`func (o *ProgramTransferTypeRequest) SetProgramFundingSourceToken(v string)`

SetProgramFundingSourceToken sets ProgramFundingSourceToken field to given value.


### GetTags

`func (o *ProgramTransferTypeRequest) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ProgramTransferTypeRequest) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ProgramTransferTypeRequest) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ProgramTransferTypeRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetToken

`func (o *ProgramTransferTypeRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ProgramTransferTypeRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ProgramTransferTypeRequest) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *ProgramTransferTypeRequest) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


