# FulfillmentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardPersonalization** | [**CardPersonalization**](CardPersonalization.md) |  | 
**Shipping** | Pointer to [**ShippingInformationResponse**](ShippingInformationResponse.md) |  | [optional] 

## Methods

### NewFulfillmentResponse

`func NewFulfillmentResponse(cardPersonalization CardPersonalization, ) *FulfillmentResponse`

NewFulfillmentResponse instantiates a new FulfillmentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFulfillmentResponseWithDefaults

`func NewFulfillmentResponseWithDefaults() *FulfillmentResponse`

NewFulfillmentResponseWithDefaults instantiates a new FulfillmentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardPersonalization

`func (o *FulfillmentResponse) GetCardPersonalization() CardPersonalization`

GetCardPersonalization returns the CardPersonalization field if non-nil, zero value otherwise.

### GetCardPersonalizationOk

`func (o *FulfillmentResponse) GetCardPersonalizationOk() (*CardPersonalization, bool)`

GetCardPersonalizationOk returns a tuple with the CardPersonalization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardPersonalization

`func (o *FulfillmentResponse) SetCardPersonalization(v CardPersonalization)`

SetCardPersonalization sets CardPersonalization field to given value.


### GetShipping

`func (o *FulfillmentResponse) GetShipping() ShippingInformationResponse`

GetShipping returns the Shipping field if non-nil, zero value otherwise.

### GetShippingOk

`func (o *FulfillmentResponse) GetShippingOk() (*ShippingInformationResponse, bool)`

GetShippingOk returns a tuple with the Shipping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipping

`func (o *FulfillmentResponse) SetShipping(v ShippingInformationResponse)`

SetShipping sets Shipping field to given value.

### HasShipping

`func (o *FulfillmentResponse) HasShipping() bool`

HasShipping returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


