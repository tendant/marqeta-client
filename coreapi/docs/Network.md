# Network

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConversionRate** | Pointer to **float32** | Conversion rate between the origination currency and the settlement currency.  Returned when the transaction currency is different from the origination currency. | [optional] 
**DynamicCurrencyConversion** | Pointer to **bool** | Indicates whether currency conversion was performed dynamically at the point of sale. | [optional] [default to false]
**OriginalAmount** | Pointer to **float32** | Amount of the transaction in the currency in which it originated. | [optional] 
**OriginalCurrencyCode** | Pointer to **string** | Currency type of the origination currency. | [optional] 
**SettlementData** | Pointer to [**SettlementData**](SettlementData.md) |  | [optional] 

## Methods

### NewNetwork

`func NewNetwork() *Network`

NewNetwork instantiates a new Network object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkWithDefaults

`func NewNetworkWithDefaults() *Network`

NewNetworkWithDefaults instantiates a new Network object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConversionRate

`func (o *Network) GetConversionRate() float32`

GetConversionRate returns the ConversionRate field if non-nil, zero value otherwise.

### GetConversionRateOk

`func (o *Network) GetConversionRateOk() (*float32, bool)`

GetConversionRateOk returns a tuple with the ConversionRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionRate

`func (o *Network) SetConversionRate(v float32)`

SetConversionRate sets ConversionRate field to given value.

### HasConversionRate

`func (o *Network) HasConversionRate() bool`

HasConversionRate returns a boolean if a field has been set.

### GetDynamicCurrencyConversion

`func (o *Network) GetDynamicCurrencyConversion() bool`

GetDynamicCurrencyConversion returns the DynamicCurrencyConversion field if non-nil, zero value otherwise.

### GetDynamicCurrencyConversionOk

`func (o *Network) GetDynamicCurrencyConversionOk() (*bool, bool)`

GetDynamicCurrencyConversionOk returns a tuple with the DynamicCurrencyConversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicCurrencyConversion

`func (o *Network) SetDynamicCurrencyConversion(v bool)`

SetDynamicCurrencyConversion sets DynamicCurrencyConversion field to given value.

### HasDynamicCurrencyConversion

`func (o *Network) HasDynamicCurrencyConversion() bool`

HasDynamicCurrencyConversion returns a boolean if a field has been set.

### GetOriginalAmount

`func (o *Network) GetOriginalAmount() float32`

GetOriginalAmount returns the OriginalAmount field if non-nil, zero value otherwise.

### GetOriginalAmountOk

`func (o *Network) GetOriginalAmountOk() (*float32, bool)`

GetOriginalAmountOk returns a tuple with the OriginalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalAmount

`func (o *Network) SetOriginalAmount(v float32)`

SetOriginalAmount sets OriginalAmount field to given value.

### HasOriginalAmount

`func (o *Network) HasOriginalAmount() bool`

HasOriginalAmount returns a boolean if a field has been set.

### GetOriginalCurrencyCode

`func (o *Network) GetOriginalCurrencyCode() string`

GetOriginalCurrencyCode returns the OriginalCurrencyCode field if non-nil, zero value otherwise.

### GetOriginalCurrencyCodeOk

`func (o *Network) GetOriginalCurrencyCodeOk() (*string, bool)`

GetOriginalCurrencyCodeOk returns a tuple with the OriginalCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalCurrencyCode

`func (o *Network) SetOriginalCurrencyCode(v string)`

SetOriginalCurrencyCode sets OriginalCurrencyCode field to given value.

### HasOriginalCurrencyCode

`func (o *Network) HasOriginalCurrencyCode() bool`

HasOriginalCurrencyCode returns a boolean if a field has been set.

### GetSettlementData

`func (o *Network) GetSettlementData() SettlementData`

GetSettlementData returns the SettlementData field if non-nil, zero value otherwise.

### GetSettlementDataOk

`func (o *Network) GetSettlementDataOk() (*SettlementData, bool)`

GetSettlementDataOk returns a tuple with the SettlementData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementData

`func (o *Network) SetSettlementData(v SettlementData)`

SetSettlementData sets SettlementData field to given value.

### HasSettlementData

`func (o *Network) HasSettlementData() bool`

HasSettlementData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


