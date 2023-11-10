# RefundDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | Description of the refund. | 
**Method** | [**RefundMethod**](RefundMethod.md) |  | 

## Methods

### NewRefundDetails

`func NewRefundDetails(description string, method RefundMethod, ) *RefundDetails`

NewRefundDetails instantiates a new RefundDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefundDetailsWithDefaults

`func NewRefundDetailsWithDefaults() *RefundDetails`

NewRefundDetailsWithDefaults instantiates a new RefundDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *RefundDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RefundDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RefundDetails) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetMethod

`func (o *RefundDetails) GetMethod() RefundMethod`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *RefundDetails) GetMethodOk() (*RefundMethod, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *RefundDetails) SetMethod(v RefundMethod)`

SetMethod sets Method field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


