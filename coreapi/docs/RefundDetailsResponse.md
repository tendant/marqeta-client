# RefundDetailsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | Description of the refund. | 
**Method** | [**RefundMethod**](RefundMethod.md) |  | 
**Status** | [**RefundStatus**](RefundStatus.md) |  | 

## Methods

### NewRefundDetailsResponse

`func NewRefundDetailsResponse(description string, method RefundMethod, status RefundStatus, ) *RefundDetailsResponse`

NewRefundDetailsResponse instantiates a new RefundDetailsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefundDetailsResponseWithDefaults

`func NewRefundDetailsResponseWithDefaults() *RefundDetailsResponse`

NewRefundDetailsResponseWithDefaults instantiates a new RefundDetailsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *RefundDetailsResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RefundDetailsResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RefundDetailsResponse) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetMethod

`func (o *RefundDetailsResponse) GetMethod() RefundMethod`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *RefundDetailsResponse) GetMethodOk() (*RefundMethod, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *RefundDetailsResponse) SetMethod(v RefundMethod)`

SetMethod sets Method field to given value.


### GetStatus

`func (o *RefundDetailsResponse) GetStatus() RefundStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RefundDetailsResponse) GetStatusOk() (*RefundStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RefundDetailsResponse) SetStatus(v RefundStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


