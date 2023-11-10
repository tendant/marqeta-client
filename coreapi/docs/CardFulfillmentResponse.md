# CardFulfillmentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardFulfillmentReason** | Pointer to **string** | Descriptive reason for the card fulfillment. | [optional] 
**CardPersonalization** | [**CardPersonalization**](CardPersonalization.md) |  | 
**Shipping** | Pointer to [**ShippingInformationResponse**](ShippingInformationResponse.md) |  | [optional] 

## Methods

### NewCardFulfillmentResponse

`func NewCardFulfillmentResponse(cardPersonalization CardPersonalization, ) *CardFulfillmentResponse`

NewCardFulfillmentResponse instantiates a new CardFulfillmentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardFulfillmentResponseWithDefaults

`func NewCardFulfillmentResponseWithDefaults() *CardFulfillmentResponse`

NewCardFulfillmentResponseWithDefaults instantiates a new CardFulfillmentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardFulfillmentReason

`func (o *CardFulfillmentResponse) GetCardFulfillmentReason() string`

GetCardFulfillmentReason returns the CardFulfillmentReason field if non-nil, zero value otherwise.

### GetCardFulfillmentReasonOk

`func (o *CardFulfillmentResponse) GetCardFulfillmentReasonOk() (*string, bool)`

GetCardFulfillmentReasonOk returns a tuple with the CardFulfillmentReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardFulfillmentReason

`func (o *CardFulfillmentResponse) SetCardFulfillmentReason(v string)`

SetCardFulfillmentReason sets CardFulfillmentReason field to given value.

### HasCardFulfillmentReason

`func (o *CardFulfillmentResponse) HasCardFulfillmentReason() bool`

HasCardFulfillmentReason returns a boolean if a field has been set.

### GetCardPersonalization

`func (o *CardFulfillmentResponse) GetCardPersonalization() CardPersonalization`

GetCardPersonalization returns the CardPersonalization field if non-nil, zero value otherwise.

### GetCardPersonalizationOk

`func (o *CardFulfillmentResponse) GetCardPersonalizationOk() (*CardPersonalization, bool)`

GetCardPersonalizationOk returns a tuple with the CardPersonalization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardPersonalization

`func (o *CardFulfillmentResponse) SetCardPersonalization(v CardPersonalization)`

SetCardPersonalization sets CardPersonalization field to given value.


### GetShipping

`func (o *CardFulfillmentResponse) GetShipping() ShippingInformationResponse`

GetShipping returns the Shipping field if non-nil, zero value otherwise.

### GetShippingOk

`func (o *CardFulfillmentResponse) GetShippingOk() (*ShippingInformationResponse, bool)`

GetShippingOk returns a tuple with the Shipping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipping

`func (o *CardFulfillmentResponse) SetShipping(v ShippingInformationResponse)`

SetShipping sets Shipping field to given value.

### HasShipping

`func (o *CardFulfillmentResponse) HasShipping() bool`

HasShipping returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


