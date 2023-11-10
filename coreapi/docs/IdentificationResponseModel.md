# IdentificationResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpirationDate** | Pointer to **string** | Expiration date for the identification, if applicable. | [optional] [readonly] 
**Type** | Pointer to **string** | Type of identification. | [optional] [readonly] 
**Value** | Pointer to **string** | Number associated with the identification. | [optional] [readonly] 

## Methods

### NewIdentificationResponseModel

`func NewIdentificationResponseModel() *IdentificationResponseModel`

NewIdentificationResponseModel instantiates a new IdentificationResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentificationResponseModelWithDefaults

`func NewIdentificationResponseModelWithDefaults() *IdentificationResponseModel`

NewIdentificationResponseModelWithDefaults instantiates a new IdentificationResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpirationDate

`func (o *IdentificationResponseModel) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *IdentificationResponseModel) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *IdentificationResponseModel) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *IdentificationResponseModel) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetType

`func (o *IdentificationResponseModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IdentificationResponseModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IdentificationResponseModel) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IdentificationResponseModel) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *IdentificationResponseModel) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *IdentificationResponseModel) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *IdentificationResponseModel) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *IdentificationResponseModel) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


