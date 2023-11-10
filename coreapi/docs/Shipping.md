# Shipping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CareOfLine** | Pointer to **string** | Adds the specified value as a care of (C/O) line to the mailing carrier.  *NOTE:* This field can be specified on cards, card products, and bulk card orders. If you specify this field at multiple levels, the order of precedence is: card, bulk card order, card product. | [optional] 
**Method** | Pointer to **string** | Specifies the shipping service. | [optional] 
**RecipientAddress** | Pointer to [**FulfillmentAddressRequest**](FulfillmentAddressRequest.md) |  | [optional] 
**ReturnAddress** | Pointer to [**FulfillmentAddressRequest**](FulfillmentAddressRequest.md) |  | [optional] 

## Methods

### NewShipping

`func NewShipping() *Shipping`

NewShipping instantiates a new Shipping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShippingWithDefaults

`func NewShippingWithDefaults() *Shipping`

NewShippingWithDefaults instantiates a new Shipping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCareOfLine

`func (o *Shipping) GetCareOfLine() string`

GetCareOfLine returns the CareOfLine field if non-nil, zero value otherwise.

### GetCareOfLineOk

`func (o *Shipping) GetCareOfLineOk() (*string, bool)`

GetCareOfLineOk returns a tuple with the CareOfLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCareOfLine

`func (o *Shipping) SetCareOfLine(v string)`

SetCareOfLine sets CareOfLine field to given value.

### HasCareOfLine

`func (o *Shipping) HasCareOfLine() bool`

HasCareOfLine returns a boolean if a field has been set.

### GetMethod

`func (o *Shipping) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *Shipping) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *Shipping) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *Shipping) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetRecipientAddress

`func (o *Shipping) GetRecipientAddress() FulfillmentAddressRequest`

GetRecipientAddress returns the RecipientAddress field if non-nil, zero value otherwise.

### GetRecipientAddressOk

`func (o *Shipping) GetRecipientAddressOk() (*FulfillmentAddressRequest, bool)`

GetRecipientAddressOk returns a tuple with the RecipientAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientAddress

`func (o *Shipping) SetRecipientAddress(v FulfillmentAddressRequest)`

SetRecipientAddress sets RecipientAddress field to given value.

### HasRecipientAddress

`func (o *Shipping) HasRecipientAddress() bool`

HasRecipientAddress returns a boolean if a field has been set.

### GetReturnAddress

`func (o *Shipping) GetReturnAddress() FulfillmentAddressRequest`

GetReturnAddress returns the ReturnAddress field if non-nil, zero value otherwise.

### GetReturnAddressOk

`func (o *Shipping) GetReturnAddressOk() (*FulfillmentAddressRequest, bool)`

GetReturnAddressOk returns a tuple with the ReturnAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnAddress

`func (o *Shipping) SetReturnAddress(v FulfillmentAddressRequest)`

SetReturnAddress sets ReturnAddress field to given value.

### HasReturnAddress

`func (o *Shipping) HasReturnAddress() bool`

HasReturnAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


