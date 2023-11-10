# Result

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Codes** | Pointer to [**[]ResultCode**](ResultCode.md) | An array of KYC verification result code objects. | [optional] 
**Status** | Pointer to **string** | KYC verification status. | [optional] 

## Methods

### NewResult

`func NewResult() *Result`

NewResult instantiates a new Result object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultWithDefaults

`func NewResultWithDefaults() *Result`

NewResultWithDefaults instantiates a new Result object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCodes

`func (o *Result) GetCodes() []ResultCode`

GetCodes returns the Codes field if non-nil, zero value otherwise.

### GetCodesOk

`func (o *Result) GetCodesOk() (*[]ResultCode, bool)`

GetCodesOk returns a tuple with the Codes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodes

`func (o *Result) SetCodes(v []ResultCode)`

SetCodes sets Codes field to given value.

### HasCodes

`func (o *Result) HasCodes() bool`

HasCodes returns a boolean if a field has been set.

### GetStatus

`func (o *Result) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Result) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Result) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Result) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


