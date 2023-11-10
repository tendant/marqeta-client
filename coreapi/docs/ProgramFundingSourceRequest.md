# ProgramFundingSourceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Indicates whether the program funding source is active. | [optional] 
**Name** | **string** | Name of the program funding source. | 
**Token** | Pointer to **string** | Unique identifier of the funding source. If you do not include a token, the system will generate one automatically. As this token is necessary for use in other calls, we recommend that you define a simple and easy to remember string rather than letting the system generate a token for you. This value cannot be updated. | [optional] 

## Methods

### NewProgramFundingSourceRequest

`func NewProgramFundingSourceRequest(name string, ) *ProgramFundingSourceRequest`

NewProgramFundingSourceRequest instantiates a new ProgramFundingSourceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProgramFundingSourceRequestWithDefaults

`func NewProgramFundingSourceRequestWithDefaults() *ProgramFundingSourceRequest`

NewProgramFundingSourceRequestWithDefaults instantiates a new ProgramFundingSourceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *ProgramFundingSourceRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ProgramFundingSourceRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ProgramFundingSourceRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ProgramFundingSourceRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetName

`func (o *ProgramFundingSourceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProgramFundingSourceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProgramFundingSourceRequest) SetName(v string)`

SetName sets Name field to given value.


### GetToken

`func (o *ProgramFundingSourceRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ProgramFundingSourceRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ProgramFundingSourceRequest) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *ProgramFundingSourceRequest) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


