# PinRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ControlToken** | **string** | Unique value generated as a result of issuing a &#x60;POST&#x60; request to the &#x60;/pins/controltoken&#x60; endpoint. This value cannot be updated. | 
**Pin** | **string** | Four-digit number to associate with the card. | 

## Methods

### NewPinRequest

`func NewPinRequest(controlToken string, pin string, ) *PinRequest`

NewPinRequest instantiates a new PinRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPinRequestWithDefaults

`func NewPinRequestWithDefaults() *PinRequest`

NewPinRequestWithDefaults instantiates a new PinRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetControlToken

`func (o *PinRequest) GetControlToken() string`

GetControlToken returns the ControlToken field if non-nil, zero value otherwise.

### GetControlTokenOk

`func (o *PinRequest) GetControlTokenOk() (*string, bool)`

GetControlTokenOk returns a tuple with the ControlToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlToken

`func (o *PinRequest) SetControlToken(v string)`

SetControlToken sets ControlToken field to given value.


### GetPin

`func (o *PinRequest) GetPin() string`

GetPin returns the Pin field if non-nil, zero value otherwise.

### GetPinOk

`func (o *PinRequest) GetPinOk() (*string, bool)`

GetPinOk returns a tuple with the Pin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPin

`func (o *PinRequest) SetPin(v string)`

SetPin sets Pin field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


