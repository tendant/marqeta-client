# CardProductFulfillment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllZeroCardSecurityCode** | Pointer to **bool** | If &#x60;true&#x60;, an all zero code (000) is allowed as a valid value in an authorization request. | [optional] [default to false]
**AllowCardCreation** | Pointer to **bool** | Controls the ability to create cards from this card product; &#x60;true&#x60; allows and &#x60;false&#x60; disallows the creation of cards.  *NOTE:* The card product&#39;s &#x60;active&#x60; field has no effect on card creation or the behavior of this field. | [optional] [default to true]
**BinPrefix** | Pointer to **string** | Prefix of the bank identification number. | [optional] 
**BulkShip** | Pointer to **bool** | Enables bulk ordering of cards of this card product type using the &#x60;/bulkissuances&#x60; endpoint. | [optional] [default to false]
**CardPersonalization** | [**CardPersonalization**](CardPersonalization.md) |  | 
**EnableOfflinePin** | Pointer to **bool** | Enables offline personal identification number (PIN) verification for Europay Mastercard and Visa (EMV, or \&quot;chip-and-PIN\&quot;) card payments. | [optional] [default to false]
**FulfillmentProvider** | Pointer to **string** | Specifies the fulfillment provider.  *NOTE:* Expedited processing is available for cards that are fulfilled by link:http://perfectplastic.com/[Perfect Plastic Printing, window&#x3D;\&quot;_blank\&quot;], link:http://www.idemia.com[IDEMIA, window&#x3D;\&quot;_blank\&quot;], and link:https://www.arroweye.com/[Arroweye Solutions, window&#x3D;\&quot;_blank\&quot;]. You can expedite an order&#39;s processing by using the &#x60;expedite&#x60; field of the &lt;&lt;/core-api/cards, card&gt;&gt; or &lt;&lt;/core-api/bulk-card-orders, bulkissuance&gt;&gt; object. Contact your Marqeta representative for information regarding the cost of expedited service. | [optional] [default to "PERFECTPLASTIC"]
**PackageId** | Pointer to **string** | Card fulfillment provider&#39;s package ID. | [optional] [default to "0"]
**PanLength** | Pointer to **string** | Specifies the length of the primary account number (PAN). | [optional] [default to "16"]
**PaymentInstrument** | Pointer to **string** | Specifies the physical form cards of this card product type will take. | [optional] [default to "PHYSICAL_MSR"]
**Shipping** | Pointer to [**Shipping**](Shipping.md) |  | [optional] 
**UppercaseNameLines** | Pointer to **bool** | A value of &#x60;true&#x60; sets the text in the &#x60;fulfillment.card_personalization.text.name_line_1&#x60; and &#x60;name_line_2&#x60; fields to all uppercase letters. A value of &#x60;false&#x60; leaves the text in its original state. | [optional] [default to true]

## Methods

### NewCardProductFulfillment

`func NewCardProductFulfillment(cardPersonalization CardPersonalization, ) *CardProductFulfillment`

NewCardProductFulfillment instantiates a new CardProductFulfillment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardProductFulfillmentWithDefaults

`func NewCardProductFulfillmentWithDefaults() *CardProductFulfillment`

NewCardProductFulfillmentWithDefaults instantiates a new CardProductFulfillment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllZeroCardSecurityCode

`func (o *CardProductFulfillment) GetAllZeroCardSecurityCode() bool`

GetAllZeroCardSecurityCode returns the AllZeroCardSecurityCode field if non-nil, zero value otherwise.

### GetAllZeroCardSecurityCodeOk

`func (o *CardProductFulfillment) GetAllZeroCardSecurityCodeOk() (*bool, bool)`

GetAllZeroCardSecurityCodeOk returns a tuple with the AllZeroCardSecurityCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllZeroCardSecurityCode

`func (o *CardProductFulfillment) SetAllZeroCardSecurityCode(v bool)`

SetAllZeroCardSecurityCode sets AllZeroCardSecurityCode field to given value.

### HasAllZeroCardSecurityCode

`func (o *CardProductFulfillment) HasAllZeroCardSecurityCode() bool`

HasAllZeroCardSecurityCode returns a boolean if a field has been set.

### GetAllowCardCreation

`func (o *CardProductFulfillment) GetAllowCardCreation() bool`

GetAllowCardCreation returns the AllowCardCreation field if non-nil, zero value otherwise.

### GetAllowCardCreationOk

`func (o *CardProductFulfillment) GetAllowCardCreationOk() (*bool, bool)`

GetAllowCardCreationOk returns a tuple with the AllowCardCreation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCardCreation

`func (o *CardProductFulfillment) SetAllowCardCreation(v bool)`

SetAllowCardCreation sets AllowCardCreation field to given value.

### HasAllowCardCreation

`func (o *CardProductFulfillment) HasAllowCardCreation() bool`

HasAllowCardCreation returns a boolean if a field has been set.

### GetBinPrefix

`func (o *CardProductFulfillment) GetBinPrefix() string`

GetBinPrefix returns the BinPrefix field if non-nil, zero value otherwise.

### GetBinPrefixOk

`func (o *CardProductFulfillment) GetBinPrefixOk() (*string, bool)`

GetBinPrefixOk returns a tuple with the BinPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinPrefix

`func (o *CardProductFulfillment) SetBinPrefix(v string)`

SetBinPrefix sets BinPrefix field to given value.

### HasBinPrefix

`func (o *CardProductFulfillment) HasBinPrefix() bool`

HasBinPrefix returns a boolean if a field has been set.

### GetBulkShip

`func (o *CardProductFulfillment) GetBulkShip() bool`

GetBulkShip returns the BulkShip field if non-nil, zero value otherwise.

### GetBulkShipOk

`func (o *CardProductFulfillment) GetBulkShipOk() (*bool, bool)`

GetBulkShipOk returns a tuple with the BulkShip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulkShip

`func (o *CardProductFulfillment) SetBulkShip(v bool)`

SetBulkShip sets BulkShip field to given value.

### HasBulkShip

`func (o *CardProductFulfillment) HasBulkShip() bool`

HasBulkShip returns a boolean if a field has been set.

### GetCardPersonalization

`func (o *CardProductFulfillment) GetCardPersonalization() CardPersonalization`

GetCardPersonalization returns the CardPersonalization field if non-nil, zero value otherwise.

### GetCardPersonalizationOk

`func (o *CardProductFulfillment) GetCardPersonalizationOk() (*CardPersonalization, bool)`

GetCardPersonalizationOk returns a tuple with the CardPersonalization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardPersonalization

`func (o *CardProductFulfillment) SetCardPersonalization(v CardPersonalization)`

SetCardPersonalization sets CardPersonalization field to given value.


### GetEnableOfflinePin

`func (o *CardProductFulfillment) GetEnableOfflinePin() bool`

GetEnableOfflinePin returns the EnableOfflinePin field if non-nil, zero value otherwise.

### GetEnableOfflinePinOk

`func (o *CardProductFulfillment) GetEnableOfflinePinOk() (*bool, bool)`

GetEnableOfflinePinOk returns a tuple with the EnableOfflinePin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableOfflinePin

`func (o *CardProductFulfillment) SetEnableOfflinePin(v bool)`

SetEnableOfflinePin sets EnableOfflinePin field to given value.

### HasEnableOfflinePin

`func (o *CardProductFulfillment) HasEnableOfflinePin() bool`

HasEnableOfflinePin returns a boolean if a field has been set.

### GetFulfillmentProvider

`func (o *CardProductFulfillment) GetFulfillmentProvider() string`

GetFulfillmentProvider returns the FulfillmentProvider field if non-nil, zero value otherwise.

### GetFulfillmentProviderOk

`func (o *CardProductFulfillment) GetFulfillmentProviderOk() (*string, bool)`

GetFulfillmentProviderOk returns a tuple with the FulfillmentProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentProvider

`func (o *CardProductFulfillment) SetFulfillmentProvider(v string)`

SetFulfillmentProvider sets FulfillmentProvider field to given value.

### HasFulfillmentProvider

`func (o *CardProductFulfillment) HasFulfillmentProvider() bool`

HasFulfillmentProvider returns a boolean if a field has been set.

### GetPackageId

`func (o *CardProductFulfillment) GetPackageId() string`

GetPackageId returns the PackageId field if non-nil, zero value otherwise.

### GetPackageIdOk

`func (o *CardProductFulfillment) GetPackageIdOk() (*string, bool)`

GetPackageIdOk returns a tuple with the PackageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageId

`func (o *CardProductFulfillment) SetPackageId(v string)`

SetPackageId sets PackageId field to given value.

### HasPackageId

`func (o *CardProductFulfillment) HasPackageId() bool`

HasPackageId returns a boolean if a field has been set.

### GetPanLength

`func (o *CardProductFulfillment) GetPanLength() string`

GetPanLength returns the PanLength field if non-nil, zero value otherwise.

### GetPanLengthOk

`func (o *CardProductFulfillment) GetPanLengthOk() (*string, bool)`

GetPanLengthOk returns a tuple with the PanLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPanLength

`func (o *CardProductFulfillment) SetPanLength(v string)`

SetPanLength sets PanLength field to given value.

### HasPanLength

`func (o *CardProductFulfillment) HasPanLength() bool`

HasPanLength returns a boolean if a field has been set.

### GetPaymentInstrument

`func (o *CardProductFulfillment) GetPaymentInstrument() string`

GetPaymentInstrument returns the PaymentInstrument field if non-nil, zero value otherwise.

### GetPaymentInstrumentOk

`func (o *CardProductFulfillment) GetPaymentInstrumentOk() (*string, bool)`

GetPaymentInstrumentOk returns a tuple with the PaymentInstrument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentInstrument

`func (o *CardProductFulfillment) SetPaymentInstrument(v string)`

SetPaymentInstrument sets PaymentInstrument field to given value.

### HasPaymentInstrument

`func (o *CardProductFulfillment) HasPaymentInstrument() bool`

HasPaymentInstrument returns a boolean if a field has been set.

### GetShipping

`func (o *CardProductFulfillment) GetShipping() Shipping`

GetShipping returns the Shipping field if non-nil, zero value otherwise.

### GetShippingOk

`func (o *CardProductFulfillment) GetShippingOk() (*Shipping, bool)`

GetShippingOk returns a tuple with the Shipping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipping

`func (o *CardProductFulfillment) SetShipping(v Shipping)`

SetShipping sets Shipping field to given value.

### HasShipping

`func (o *CardProductFulfillment) HasShipping() bool`

HasShipping returns a boolean if a field has been set.

### GetUppercaseNameLines

`func (o *CardProductFulfillment) GetUppercaseNameLines() bool`

GetUppercaseNameLines returns the UppercaseNameLines field if non-nil, zero value otherwise.

### GetUppercaseNameLinesOk

`func (o *CardProductFulfillment) GetUppercaseNameLinesOk() (*bool, bool)`

GetUppercaseNameLinesOk returns a tuple with the UppercaseNameLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUppercaseNameLines

`func (o *CardProductFulfillment) SetUppercaseNameLines(v bool)`

SetUppercaseNameLines sets UppercaseNameLines field to given value.

### HasUppercaseNameLines

`func (o *CardProductFulfillment) HasUppercaseNameLines() bool`

HasUppercaseNameLines returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


