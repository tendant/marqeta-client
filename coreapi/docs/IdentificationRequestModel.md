# IdentificationRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpirationDate** | Pointer to **string** | Expiration date of the identification, if applicable. | [optional] 
**Type** | **string** | Type of identification.  *NOTE:* Full Social Security Number (SSN) is required for US-based cardholder KYC verification. Nine digits only, no delimiters. &#x60;123456789&#x60;, for example. | [readonly] 
**Value** | Pointer to **string** | Number associated with the identification. | [optional] [readonly] 

## Methods

### NewIdentificationRequestModel

`func NewIdentificationRequestModel(type_ string, ) *IdentificationRequestModel`

NewIdentificationRequestModel instantiates a new IdentificationRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentificationRequestModelWithDefaults

`func NewIdentificationRequestModelWithDefaults() *IdentificationRequestModel`

NewIdentificationRequestModelWithDefaults instantiates a new IdentificationRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpirationDate

`func (o *IdentificationRequestModel) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *IdentificationRequestModel) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *IdentificationRequestModel) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *IdentificationRequestModel) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetType

`func (o *IdentificationRequestModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IdentificationRequestModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IdentificationRequestModel) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *IdentificationRequestModel) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *IdentificationRequestModel) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *IdentificationRequestModel) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *IdentificationRequestModel) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


