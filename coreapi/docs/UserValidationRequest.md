# UserValidationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BirthDate** | Pointer to **time.Time** | Date of birth of the user associated with this card. | [optional] 
**Phone** | Pointer to **string** | Telephone number of the user associated with this card. | [optional] 
**RandomNameLine1Postfix** | Pointer to **string** | Random six-digit numeric postfix generated for some bulk card orders.  See &lt;&lt;/core-api/bulk-card-orders, Bulk Card Orders&gt;&gt; for more information about numeric postfixes. | [optional] 
**Ssn** | Pointer to **string** | Social Security Number (SSN) of the user associated with this card. | [optional] 

## Methods

### NewUserValidationRequest

`func NewUserValidationRequest() *UserValidationRequest`

NewUserValidationRequest instantiates a new UserValidationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserValidationRequestWithDefaults

`func NewUserValidationRequestWithDefaults() *UserValidationRequest`

NewUserValidationRequestWithDefaults instantiates a new UserValidationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBirthDate

`func (o *UserValidationRequest) GetBirthDate() time.Time`

GetBirthDate returns the BirthDate field if non-nil, zero value otherwise.

### GetBirthDateOk

`func (o *UserValidationRequest) GetBirthDateOk() (*time.Time, bool)`

GetBirthDateOk returns a tuple with the BirthDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthDate

`func (o *UserValidationRequest) SetBirthDate(v time.Time)`

SetBirthDate sets BirthDate field to given value.

### HasBirthDate

`func (o *UserValidationRequest) HasBirthDate() bool`

HasBirthDate returns a boolean if a field has been set.

### GetPhone

`func (o *UserValidationRequest) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *UserValidationRequest) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *UserValidationRequest) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *UserValidationRequest) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetRandomNameLine1Postfix

`func (o *UserValidationRequest) GetRandomNameLine1Postfix() string`

GetRandomNameLine1Postfix returns the RandomNameLine1Postfix field if non-nil, zero value otherwise.

### GetRandomNameLine1PostfixOk

`func (o *UserValidationRequest) GetRandomNameLine1PostfixOk() (*string, bool)`

GetRandomNameLine1PostfixOk returns a tuple with the RandomNameLine1Postfix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRandomNameLine1Postfix

`func (o *UserValidationRequest) SetRandomNameLine1Postfix(v string)`

SetRandomNameLine1Postfix sets RandomNameLine1Postfix field to given value.

### HasRandomNameLine1Postfix

`func (o *UserValidationRequest) HasRandomNameLine1Postfix() bool`

HasRandomNameLine1Postfix returns a boolean if a field has been set.

### GetSsn

`func (o *UserValidationRequest) GetSsn() string`

GetSsn returns the Ssn field if non-nil, zero value otherwise.

### GetSsnOk

`func (o *UserValidationRequest) GetSsnOk() (*string, bool)`

GetSsnOk returns a tuple with the Ssn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsn

`func (o *UserValidationRequest) SetSsn(v string)`

SetSsn sets Ssn field to given value.

### HasSsn

`func (o *UserValidationRequest) HasSsn() bool`

HasSsn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


