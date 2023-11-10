# AvsControlOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeclineOnAddressNumberMismatch** | Pointer to **bool** | Set to &#x60;true&#x60; to decline account verification or authorization messages whose address number does not match the address number on file.  Set to &#x60;false&#x60; to not decline account verification or authorization messages whose address number does not match the address number on file.  This field is functional only if &#x60;validate &#x3D; true&#x60;. | [optional] [default to false]
**DeclineOnPostalCodeMismatch** | Pointer to **bool** | Set to &#x60;true&#x60; to decline account verification or authorization messages whose postal code does not match the postal code on file.  Set to &#x60;false&#x60; to not decline account verification or authorization messages whose postal code does not match the postal code on file.  This field is functional only if &#x60;validate &#x3D; true&#x60;. | [optional] [default to true]
**Validate** | Pointer to **bool** | Set to &#x60;true&#x60; to require validation of account verification or authorization messages.  Set to &#x60;false&#x60; to forgo validation and approve all account verification and authorization messages. | [optional] [default to true]

## Methods

### NewAvsControlOptions

`func NewAvsControlOptions() *AvsControlOptions`

NewAvsControlOptions instantiates a new AvsControlOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAvsControlOptionsWithDefaults

`func NewAvsControlOptionsWithDefaults() *AvsControlOptions`

NewAvsControlOptionsWithDefaults instantiates a new AvsControlOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeclineOnAddressNumberMismatch

`func (o *AvsControlOptions) GetDeclineOnAddressNumberMismatch() bool`

GetDeclineOnAddressNumberMismatch returns the DeclineOnAddressNumberMismatch field if non-nil, zero value otherwise.

### GetDeclineOnAddressNumberMismatchOk

`func (o *AvsControlOptions) GetDeclineOnAddressNumberMismatchOk() (*bool, bool)`

GetDeclineOnAddressNumberMismatchOk returns a tuple with the DeclineOnAddressNumberMismatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclineOnAddressNumberMismatch

`func (o *AvsControlOptions) SetDeclineOnAddressNumberMismatch(v bool)`

SetDeclineOnAddressNumberMismatch sets DeclineOnAddressNumberMismatch field to given value.

### HasDeclineOnAddressNumberMismatch

`func (o *AvsControlOptions) HasDeclineOnAddressNumberMismatch() bool`

HasDeclineOnAddressNumberMismatch returns a boolean if a field has been set.

### GetDeclineOnPostalCodeMismatch

`func (o *AvsControlOptions) GetDeclineOnPostalCodeMismatch() bool`

GetDeclineOnPostalCodeMismatch returns the DeclineOnPostalCodeMismatch field if non-nil, zero value otherwise.

### GetDeclineOnPostalCodeMismatchOk

`func (o *AvsControlOptions) GetDeclineOnPostalCodeMismatchOk() (*bool, bool)`

GetDeclineOnPostalCodeMismatchOk returns a tuple with the DeclineOnPostalCodeMismatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclineOnPostalCodeMismatch

`func (o *AvsControlOptions) SetDeclineOnPostalCodeMismatch(v bool)`

SetDeclineOnPostalCodeMismatch sets DeclineOnPostalCodeMismatch field to given value.

### HasDeclineOnPostalCodeMismatch

`func (o *AvsControlOptions) HasDeclineOnPostalCodeMismatch() bool`

HasDeclineOnPostalCodeMismatch returns a boolean if a field has been set.

### GetValidate

`func (o *AvsControlOptions) GetValidate() bool`

GetValidate returns the Validate field if non-nil, zero value otherwise.

### GetValidateOk

`func (o *AvsControlOptions) GetValidateOk() (*bool, bool)`

GetValidateOk returns a tuple with the Validate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidate

`func (o *AvsControlOptions) SetValidate(v bool)`

SetValidate sets Validate field to given value.

### HasValidate

`func (o *AvsControlOptions) HasValidate() bool`

HasValidate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


