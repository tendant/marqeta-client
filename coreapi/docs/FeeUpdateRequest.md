# FeeUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Indicates whether the fee is active. | [optional] [default to true]
**Amount** | Pointer to **float32** | Amount of the fee. | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**CurrencyCode** | Pointer to **string** | Three-digit ISO 4217 currency code. | [optional] 
**FeeAttributes** | Pointer to [**FeeAttributes**](FeeAttributes.md) |  | [optional] 
**Name** | Pointer to **string** | Name of the fee request. | [optional] 
**Tags** | Pointer to **string** | Descriptive metadata about the fee. | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewFeeUpdateRequest

`func NewFeeUpdateRequest() *FeeUpdateRequest`

NewFeeUpdateRequest instantiates a new FeeUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeeUpdateRequestWithDefaults

`func NewFeeUpdateRequestWithDefaults() *FeeUpdateRequest`

NewFeeUpdateRequestWithDefaults instantiates a new FeeUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *FeeUpdateRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *FeeUpdateRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *FeeUpdateRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *FeeUpdateRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAmount

`func (o *FeeUpdateRequest) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *FeeUpdateRequest) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *FeeUpdateRequest) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *FeeUpdateRequest) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCategory

`func (o *FeeUpdateRequest) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *FeeUpdateRequest) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *FeeUpdateRequest) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *FeeUpdateRequest) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *FeeUpdateRequest) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *FeeUpdateRequest) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *FeeUpdateRequest) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *FeeUpdateRequest) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetFeeAttributes

`func (o *FeeUpdateRequest) GetFeeAttributes() FeeAttributes`

GetFeeAttributes returns the FeeAttributes field if non-nil, zero value otherwise.

### GetFeeAttributesOk

`func (o *FeeUpdateRequest) GetFeeAttributesOk() (*FeeAttributes, bool)`

GetFeeAttributesOk returns a tuple with the FeeAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeAttributes

`func (o *FeeUpdateRequest) SetFeeAttributes(v FeeAttributes)`

SetFeeAttributes sets FeeAttributes field to given value.

### HasFeeAttributes

`func (o *FeeUpdateRequest) HasFeeAttributes() bool`

HasFeeAttributes returns a boolean if a field has been set.

### GetName

`func (o *FeeUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FeeUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FeeUpdateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FeeUpdateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTags

`func (o *FeeUpdateRequest) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FeeUpdateRequest) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FeeUpdateRequest) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *FeeUpdateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *FeeUpdateRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FeeUpdateRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FeeUpdateRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FeeUpdateRequest) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


