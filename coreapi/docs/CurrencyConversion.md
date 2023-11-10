# CurrencyConversion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OriginalAmount** | Pointer to **float32** | Amount of the transaction in the currency in which it originated. Returned if the transaction involves currency conversion. | [optional] 
**ConversionRate** | Pointer to **float32** | Conversion rate between the origination currency and the settlement currency. Returned when the transaction currency is different from the origination currency. | [optional] 
**OriginalCurrencyCode** | Pointer to **string** | The three-digit ISO 4217 currency code. | [optional] 

## Methods

### NewCurrencyConversion

`func NewCurrencyConversion() *CurrencyConversion`

NewCurrencyConversion instantiates a new CurrencyConversion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurrencyConversionWithDefaults

`func NewCurrencyConversionWithDefaults() *CurrencyConversion`

NewCurrencyConversionWithDefaults instantiates a new CurrencyConversion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOriginalAmount

`func (o *CurrencyConversion) GetOriginalAmount() float32`

GetOriginalAmount returns the OriginalAmount field if non-nil, zero value otherwise.

### GetOriginalAmountOk

`func (o *CurrencyConversion) GetOriginalAmountOk() (*float32, bool)`

GetOriginalAmountOk returns a tuple with the OriginalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalAmount

`func (o *CurrencyConversion) SetOriginalAmount(v float32)`

SetOriginalAmount sets OriginalAmount field to given value.

### HasOriginalAmount

`func (o *CurrencyConversion) HasOriginalAmount() bool`

HasOriginalAmount returns a boolean if a field has been set.

### GetConversionRate

`func (o *CurrencyConversion) GetConversionRate() float32`

GetConversionRate returns the ConversionRate field if non-nil, zero value otherwise.

### GetConversionRateOk

`func (o *CurrencyConversion) GetConversionRateOk() (*float32, bool)`

GetConversionRateOk returns a tuple with the ConversionRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionRate

`func (o *CurrencyConversion) SetConversionRate(v float32)`

SetConversionRate sets ConversionRate field to given value.

### HasConversionRate

`func (o *CurrencyConversion) HasConversionRate() bool`

HasConversionRate returns a boolean if a field has been set.

### GetOriginalCurrencyCode

`func (o *CurrencyConversion) GetOriginalCurrencyCode() string`

GetOriginalCurrencyCode returns the OriginalCurrencyCode field if non-nil, zero value otherwise.

### GetOriginalCurrencyCodeOk

`func (o *CurrencyConversion) GetOriginalCurrencyCodeOk() (*string, bool)`

GetOriginalCurrencyCodeOk returns a tuple with the OriginalCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalCurrencyCode

`func (o *CurrencyConversion) SetOriginalCurrencyCode(v string)`

SetOriginalCurrencyCode sets OriginalCurrencyCode field to given value.

### HasOriginalCurrencyCode

`func (o *CurrencyConversion) HasOriginalCurrencyCode() bool`

HasOriginalCurrencyCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


