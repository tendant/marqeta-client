# CardFulfillmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardFulfillmentReason** | Pointer to **string** | Reason for card fulfillment. | [optional] 
**CardPersonalization** | [**CardPersonalization**](CardPersonalization.md) |  | 
**Shipping** | Pointer to [**Shipping**](Shipping.md) |  | [optional] 

## Methods

### NewCardFulfillmentRequest

`func NewCardFulfillmentRequest(cardPersonalization CardPersonalization, ) *CardFulfillmentRequest`

NewCardFulfillmentRequest instantiates a new CardFulfillmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardFulfillmentRequestWithDefaults

`func NewCardFulfillmentRequestWithDefaults() *CardFulfillmentRequest`

NewCardFulfillmentRequestWithDefaults instantiates a new CardFulfillmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardFulfillmentReason

`func (o *CardFulfillmentRequest) GetCardFulfillmentReason() string`

GetCardFulfillmentReason returns the CardFulfillmentReason field if non-nil, zero value otherwise.

### GetCardFulfillmentReasonOk

`func (o *CardFulfillmentRequest) GetCardFulfillmentReasonOk() (*string, bool)`

GetCardFulfillmentReasonOk returns a tuple with the CardFulfillmentReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardFulfillmentReason

`func (o *CardFulfillmentRequest) SetCardFulfillmentReason(v string)`

SetCardFulfillmentReason sets CardFulfillmentReason field to given value.

### HasCardFulfillmentReason

`func (o *CardFulfillmentRequest) HasCardFulfillmentReason() bool`

HasCardFulfillmentReason returns a boolean if a field has been set.

### GetCardPersonalization

`func (o *CardFulfillmentRequest) GetCardPersonalization() CardPersonalization`

GetCardPersonalization returns the CardPersonalization field if non-nil, zero value otherwise.

### GetCardPersonalizationOk

`func (o *CardFulfillmentRequest) GetCardPersonalizationOk() (*CardPersonalization, bool)`

GetCardPersonalizationOk returns a tuple with the CardPersonalization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardPersonalization

`func (o *CardFulfillmentRequest) SetCardPersonalization(v CardPersonalization)`

SetCardPersonalization sets CardPersonalization field to given value.


### GetShipping

`func (o *CardFulfillmentRequest) GetShipping() Shipping`

GetShipping returns the Shipping field if non-nil, zero value otherwise.

### GetShippingOk

`func (o *CardFulfillmentRequest) GetShippingOk() (*Shipping, bool)`

GetShippingOk returns a tuple with the Shipping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipping

`func (o *CardFulfillmentRequest) SetShipping(v Shipping)`

SetShipping sets Shipping field to given value.

### HasShipping

`func (o *CardFulfillmentRequest) HasShipping() bool`

HasShipping returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


