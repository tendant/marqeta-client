# FeeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | **bool** | Specifies whether the fee is active. | [readonly] 
**Amount** | **float32** | Amount of the fee. | 
**Category** | Pointer to **string** | Specifies if the fee is a standalone fee or a real-time fee. | [optional] [readonly] 
**CreatedTime** | **time.Time** | Date and time when the &#x60;fees&#x60; object was created, in UTC. | [readonly] 
**CurrencyCode** | **string** | Three-digit ISO 4217 currency code. | [readonly] 
**FeeAttributes** | Pointer to [**FeeAttributes**](FeeAttributes.md) |  | [optional] 
**LastModifiedTime** | **time.Time** | Date and time when the &#x60;fees&#x60; object was last modified, in UTC. | [readonly] 
**Name** | **string** | Name of the fee. | 
**Tags** | Pointer to **string** | Descriptive metadata about the fee. | [optional] 
**Token** | **string** | Unique identifier of the &#x60;fees&#x60; object. | 
**Type** | Pointer to **string** | Specifies if the fee is a flat fee or a percentage of the transaction amount. | [optional] [readonly] 

## Methods

### NewFeeResponse

`func NewFeeResponse(active bool, amount float32, createdTime time.Time, currencyCode string, lastModifiedTime time.Time, name string, token string, ) *FeeResponse`

NewFeeResponse instantiates a new FeeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeeResponseWithDefaults

`func NewFeeResponseWithDefaults() *FeeResponse`

NewFeeResponseWithDefaults instantiates a new FeeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *FeeResponse) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *FeeResponse) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *FeeResponse) SetActive(v bool)`

SetActive sets Active field to given value.


### GetAmount

`func (o *FeeResponse) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *FeeResponse) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *FeeResponse) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetCategory

`func (o *FeeResponse) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *FeeResponse) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *FeeResponse) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *FeeResponse) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCreatedTime

`func (o *FeeResponse) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *FeeResponse) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *FeeResponse) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetCurrencyCode

`func (o *FeeResponse) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *FeeResponse) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *FeeResponse) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.


### GetFeeAttributes

`func (o *FeeResponse) GetFeeAttributes() FeeAttributes`

GetFeeAttributes returns the FeeAttributes field if non-nil, zero value otherwise.

### GetFeeAttributesOk

`func (o *FeeResponse) GetFeeAttributesOk() (*FeeAttributes, bool)`

GetFeeAttributesOk returns a tuple with the FeeAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeAttributes

`func (o *FeeResponse) SetFeeAttributes(v FeeAttributes)`

SetFeeAttributes sets FeeAttributes field to given value.

### HasFeeAttributes

`func (o *FeeResponse) HasFeeAttributes() bool`

HasFeeAttributes returns a boolean if a field has been set.

### GetLastModifiedTime

`func (o *FeeResponse) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *FeeResponse) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *FeeResponse) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.


### GetName

`func (o *FeeResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FeeResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FeeResponse) SetName(v string)`

SetName sets Name field to given value.


### GetTags

`func (o *FeeResponse) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FeeResponse) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FeeResponse) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *FeeResponse) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetToken

`func (o *FeeResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *FeeResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *FeeResponse) SetToken(v string)`

SetToken sets Token field to given value.


### GetType

`func (o *FeeResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FeeResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FeeResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FeeResponse) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


