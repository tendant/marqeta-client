# UserValidationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BirthDate** | **bool** | Indicates whether the &#x60;birth_date&#x60; field in the request was validated. | [default to false]
**Phone** | **bool** | Indicates whether the &#x60;phone&#x60; field in the request was validated. | [default to false]
**RandomNameLine1Postfix** | **bool** | Indicates whether the &#x60;random_name_line1_postfix&#x60; field in the request was validated. | [default to false]
**Ssn** | **bool** | Indicates whether the &#x60;ssn&#x60; field in the request was validated. | [default to false]

## Methods

### NewUserValidationResponse

`func NewUserValidationResponse(birthDate bool, phone bool, randomNameLine1Postfix bool, ssn bool, ) *UserValidationResponse`

NewUserValidationResponse instantiates a new UserValidationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserValidationResponseWithDefaults

`func NewUserValidationResponseWithDefaults() *UserValidationResponse`

NewUserValidationResponseWithDefaults instantiates a new UserValidationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBirthDate

`func (o *UserValidationResponse) GetBirthDate() bool`

GetBirthDate returns the BirthDate field if non-nil, zero value otherwise.

### GetBirthDateOk

`func (o *UserValidationResponse) GetBirthDateOk() (*bool, bool)`

GetBirthDateOk returns a tuple with the BirthDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthDate

`func (o *UserValidationResponse) SetBirthDate(v bool)`

SetBirthDate sets BirthDate field to given value.


### GetPhone

`func (o *UserValidationResponse) GetPhone() bool`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *UserValidationResponse) GetPhoneOk() (*bool, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *UserValidationResponse) SetPhone(v bool)`

SetPhone sets Phone field to given value.


### GetRandomNameLine1Postfix

`func (o *UserValidationResponse) GetRandomNameLine1Postfix() bool`

GetRandomNameLine1Postfix returns the RandomNameLine1Postfix field if non-nil, zero value otherwise.

### GetRandomNameLine1PostfixOk

`func (o *UserValidationResponse) GetRandomNameLine1PostfixOk() (*bool, bool)`

GetRandomNameLine1PostfixOk returns a tuple with the RandomNameLine1Postfix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRandomNameLine1Postfix

`func (o *UserValidationResponse) SetRandomNameLine1Postfix(v bool)`

SetRandomNameLine1Postfix sets RandomNameLine1Postfix field to given value.


### GetSsn

`func (o *UserValidationResponse) GetSsn() bool`

GetSsn returns the Ssn field if non-nil, zero value otherwise.

### GetSsnOk

`func (o *UserValidationResponse) GetSsnOk() (*bool, bool)`

GetSsnOk returns a tuple with the Ssn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsn

`func (o *UserValidationResponse) SetSsn(v bool)`

SetSsn sets Ssn field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


