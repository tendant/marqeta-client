# FeeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Indicates whether the fee is active. | [optional] [default to true]
**Amount** | **float32** | Amount of the fee. | 
**Category** | Pointer to **string** | Specifies if the fee is a standalone fee or a real-time fee. | [optional] 
**CurrencyCode** | **string** | Three-digit ISO 4217 currency code. | 
**FeeAttributes** | Pointer to [**FeeAttributes**](FeeAttributes.md) |  | [optional] 
**Name** | **string** | Name of the fee request. | 
**Tags** | Pointer to **string** | Descriptive metadata about the fee. | [optional] 
**Token** | Pointer to **string** | Unique identifier of the fee.  If you do not include a token, the system will generate one automatically. This token is necessary for use in other API calls, so Marqeta recommends that rather than let the system generate one, you use a simple string that is easy to remember. This value cannot be updated. | [optional] 
**Type** | Pointer to **string** | Specifies if the fee is a flat fee or a percentage of the transaction amount. | [optional] 

## Methods

### NewFeeRequest

`func NewFeeRequest(amount float32, currencyCode string, name string, ) *FeeRequest`

NewFeeRequest instantiates a new FeeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeeRequestWithDefaults

`func NewFeeRequestWithDefaults() *FeeRequest`

NewFeeRequestWithDefaults instantiates a new FeeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *FeeRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *FeeRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *FeeRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *FeeRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAmount

`func (o *FeeRequest) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *FeeRequest) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *FeeRequest) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetCategory

`func (o *FeeRequest) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *FeeRequest) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *FeeRequest) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *FeeRequest) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *FeeRequest) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *FeeRequest) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *FeeRequest) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.


### GetFeeAttributes

`func (o *FeeRequest) GetFeeAttributes() FeeAttributes`

GetFeeAttributes returns the FeeAttributes field if non-nil, zero value otherwise.

### GetFeeAttributesOk

`func (o *FeeRequest) GetFeeAttributesOk() (*FeeAttributes, bool)`

GetFeeAttributesOk returns a tuple with the FeeAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeAttributes

`func (o *FeeRequest) SetFeeAttributes(v FeeAttributes)`

SetFeeAttributes sets FeeAttributes field to given value.

### HasFeeAttributes

`func (o *FeeRequest) HasFeeAttributes() bool`

HasFeeAttributes returns a boolean if a field has been set.

### GetName

`func (o *FeeRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FeeRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FeeRequest) SetName(v string)`

SetName sets Name field to given value.


### GetTags

`func (o *FeeRequest) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FeeRequest) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FeeRequest) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *FeeRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetToken

`func (o *FeeRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *FeeRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *FeeRequest) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *FeeRequest) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetType

`func (o *FeeRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FeeRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FeeRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FeeRequest) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


