# WebPushProvisioningApplePayJWSModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Header** | [**WebPushProvisioningApplePayJWSHeader**](WebPushProvisioningApplePayJWSHeader.md) |  | 
**Payload** | **string** | JSON Web Signature (JWS) message payload. | 
**Protected** | **string** | Contains header parameters that are integrity-protected by the JSON Web Signature (JWS). | 
**Signature** | **string** | The JSON Web Signature (JWS). | 

## Methods

### NewWebPushProvisioningApplePayJWSModel

`func NewWebPushProvisioningApplePayJWSModel(header WebPushProvisioningApplePayJWSHeader, payload string, protected string, signature string, ) *WebPushProvisioningApplePayJWSModel`

NewWebPushProvisioningApplePayJWSModel instantiates a new WebPushProvisioningApplePayJWSModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebPushProvisioningApplePayJWSModelWithDefaults

`func NewWebPushProvisioningApplePayJWSModelWithDefaults() *WebPushProvisioningApplePayJWSModel`

NewWebPushProvisioningApplePayJWSModelWithDefaults instantiates a new WebPushProvisioningApplePayJWSModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeader

`func (o *WebPushProvisioningApplePayJWSModel) GetHeader() WebPushProvisioningApplePayJWSHeader`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *WebPushProvisioningApplePayJWSModel) GetHeaderOk() (*WebPushProvisioningApplePayJWSHeader, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *WebPushProvisioningApplePayJWSModel) SetHeader(v WebPushProvisioningApplePayJWSHeader)`

SetHeader sets Header field to given value.


### GetPayload

`func (o *WebPushProvisioningApplePayJWSModel) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *WebPushProvisioningApplePayJWSModel) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *WebPushProvisioningApplePayJWSModel) SetPayload(v string)`

SetPayload sets Payload field to given value.


### GetProtected

`func (o *WebPushProvisioningApplePayJWSModel) GetProtected() string`

GetProtected returns the Protected field if non-nil, zero value otherwise.

### GetProtectedOk

`func (o *WebPushProvisioningApplePayJWSModel) GetProtectedOk() (*string, bool)`

GetProtectedOk returns a tuple with the Protected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtected

`func (o *WebPushProvisioningApplePayJWSModel) SetProtected(v string)`

SetProtected sets Protected field to given value.


### GetSignature

`func (o *WebPushProvisioningApplePayJWSModel) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *WebPushProvisioningApplePayJWSModel) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *WebPushProvisioningApplePayJWSModel) SetSignature(v string)`

SetSignature sets Signature field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


