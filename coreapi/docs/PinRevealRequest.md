# PinRevealRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardholderVerificationMethod** | **string** | The supplemental method used to verify the cardholder&#39;s identity before revealing the card&#39;s personal identification number (PIN).  The possible cardholder verification methods are:  * *BIOMETRIC_FACE:* In-app authentication via facial recognition * *BIOMETRIC_FINGERPRINT:* In-app authentication via biometric fingerprint * *EXP_CVV:* In-app authentication by entering the card&#39;s expiration date and card verification value (CVV) * *LOGIN:* In-app authentication by re-entering the app password * *OTP:* Two-factor authentication involving a one-time password (OTP) * *OTP_CVV:* Two-factor authentication involving the card&#39;s CVV and an OTP * *OTHER:* Authentication that relies on other secure methods | 
**ControlToken** | **string** | Unique value generated as a result of issuing a &#x60;POST&#x60; request to the &#x60;/pins/controltoken&#x60; endpoint. This value cannot be updated. | 

## Methods

### NewPinRevealRequest

`func NewPinRevealRequest(cardholderVerificationMethod string, controlToken string, ) *PinRevealRequest`

NewPinRevealRequest instantiates a new PinRevealRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPinRevealRequestWithDefaults

`func NewPinRevealRequestWithDefaults() *PinRevealRequest`

NewPinRevealRequestWithDefaults instantiates a new PinRevealRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardholderVerificationMethod

`func (o *PinRevealRequest) GetCardholderVerificationMethod() string`

GetCardholderVerificationMethod returns the CardholderVerificationMethod field if non-nil, zero value otherwise.

### GetCardholderVerificationMethodOk

`func (o *PinRevealRequest) GetCardholderVerificationMethodOk() (*string, bool)`

GetCardholderVerificationMethodOk returns a tuple with the CardholderVerificationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardholderVerificationMethod

`func (o *PinRevealRequest) SetCardholderVerificationMethod(v string)`

SetCardholderVerificationMethod sets CardholderVerificationMethod field to given value.


### GetControlToken

`func (o *PinRevealRequest) GetControlToken() string`

GetControlToken returns the ControlToken field if non-nil, zero value otherwise.

### GetControlTokenOk

`func (o *PinRevealRequest) GetControlTokenOk() (*string, bool)`

GetControlTokenOk returns a tuple with the ControlToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlToken

`func (o *PinRevealRequest) SetControlToken(v string)`

SetControlToken sets ControlToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


