# ShippingInformationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CareOfLine** | Pointer to **string** | Specifies the value of the care of (C/O) line on the mailing carrier.  This field is returned if it exists in the resource. | [optional] 
**Method** | Pointer to **string** | Specifies the shipping service.  This field is returned if it exists in the resource. | [optional] 
**RecipientAddress** | Pointer to [**FulfillmentAddressResponse**](FulfillmentAddressResponse.md) |  | [optional] 
**ReturnAddress** | Pointer to [**FulfillmentAddressResponse**](FulfillmentAddressResponse.md) |  | [optional] 

## Methods

### NewShippingInformationResponse

`func NewShippingInformationResponse() *ShippingInformationResponse`

NewShippingInformationResponse instantiates a new ShippingInformationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShippingInformationResponseWithDefaults

`func NewShippingInformationResponseWithDefaults() *ShippingInformationResponse`

NewShippingInformationResponseWithDefaults instantiates a new ShippingInformationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCareOfLine

`func (o *ShippingInformationResponse) GetCareOfLine() string`

GetCareOfLine returns the CareOfLine field if non-nil, zero value otherwise.

### GetCareOfLineOk

`func (o *ShippingInformationResponse) GetCareOfLineOk() (*string, bool)`

GetCareOfLineOk returns a tuple with the CareOfLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCareOfLine

`func (o *ShippingInformationResponse) SetCareOfLine(v string)`

SetCareOfLine sets CareOfLine field to given value.

### HasCareOfLine

`func (o *ShippingInformationResponse) HasCareOfLine() bool`

HasCareOfLine returns a boolean if a field has been set.

### GetMethod

`func (o *ShippingInformationResponse) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *ShippingInformationResponse) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *ShippingInformationResponse) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *ShippingInformationResponse) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetRecipientAddress

`func (o *ShippingInformationResponse) GetRecipientAddress() FulfillmentAddressResponse`

GetRecipientAddress returns the RecipientAddress field if non-nil, zero value otherwise.

### GetRecipientAddressOk

`func (o *ShippingInformationResponse) GetRecipientAddressOk() (*FulfillmentAddressResponse, bool)`

GetRecipientAddressOk returns a tuple with the RecipientAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientAddress

`func (o *ShippingInformationResponse) SetRecipientAddress(v FulfillmentAddressResponse)`

SetRecipientAddress sets RecipientAddress field to given value.

### HasRecipientAddress

`func (o *ShippingInformationResponse) HasRecipientAddress() bool`

HasRecipientAddress returns a boolean if a field has been set.

### GetReturnAddress

`func (o *ShippingInformationResponse) GetReturnAddress() FulfillmentAddressResponse`

GetReturnAddress returns the ReturnAddress field if non-nil, zero value otherwise.

### GetReturnAddressOk

`func (o *ShippingInformationResponse) GetReturnAddressOk() (*FulfillmentAddressResponse, bool)`

GetReturnAddressOk returns a tuple with the ReturnAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnAddress

`func (o *ShippingInformationResponse) SetReturnAddress(v FulfillmentAddressResponse)`

SetReturnAddress sets ReturnAddress field to given value.

### HasReturnAddress

`func (o *ShippingInformationResponse) HasReturnAddress() bool`

HasReturnAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


