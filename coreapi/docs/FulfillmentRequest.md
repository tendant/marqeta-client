# FulfillmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardPersonalization** | [**CardPersonalization**](CardPersonalization.md) |  | 
**Shipping** | Pointer to [**Shipping**](Shipping.md) |  | [optional] 

## Methods

### NewFulfillmentRequest

`func NewFulfillmentRequest(cardPersonalization CardPersonalization, ) *FulfillmentRequest`

NewFulfillmentRequest instantiates a new FulfillmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFulfillmentRequestWithDefaults

`func NewFulfillmentRequestWithDefaults() *FulfillmentRequest`

NewFulfillmentRequestWithDefaults instantiates a new FulfillmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardPersonalization

`func (o *FulfillmentRequest) GetCardPersonalization() CardPersonalization`

GetCardPersonalization returns the CardPersonalization field if non-nil, zero value otherwise.

### GetCardPersonalizationOk

`func (o *FulfillmentRequest) GetCardPersonalizationOk() (*CardPersonalization, bool)`

GetCardPersonalizationOk returns a tuple with the CardPersonalization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardPersonalization

`func (o *FulfillmentRequest) SetCardPersonalization(v CardPersonalization)`

SetCardPersonalization sets CardPersonalization field to given value.


### GetShipping

`func (o *FulfillmentRequest) GetShipping() Shipping`

GetShipping returns the Shipping field if non-nil, zero value otherwise.

### GetShippingOk

`func (o *FulfillmentRequest) GetShippingOk() (*Shipping, bool)`

GetShippingOk returns a tuple with the Shipping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipping

`func (o *FulfillmentRequest) SetShipping(v Shipping)`

SetShipping sets Shipping field to given value.

### HasShipping

`func (o *FulfillmentRequest) HasShipping() bool`

HasShipping returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


