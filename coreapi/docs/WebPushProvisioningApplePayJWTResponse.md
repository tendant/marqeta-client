# WebPushProvisioningApplePayJWTResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Jws** | [**WebPushProvisioningApplePayJWSModel**](WebPushProvisioningApplePayJWSModel.md) |  | 
**State** | **string** | Unique state associated with the digital wallet token. The Marqeta platform returns a universally unique identifier (UUID) in this field. | 

## Methods

### NewWebPushProvisioningApplePayJWTResponse

`func NewWebPushProvisioningApplePayJWTResponse(jws WebPushProvisioningApplePayJWSModel, state string, ) *WebPushProvisioningApplePayJWTResponse`

NewWebPushProvisioningApplePayJWTResponse instantiates a new WebPushProvisioningApplePayJWTResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebPushProvisioningApplePayJWTResponseWithDefaults

`func NewWebPushProvisioningApplePayJWTResponseWithDefaults() *WebPushProvisioningApplePayJWTResponse`

NewWebPushProvisioningApplePayJWTResponseWithDefaults instantiates a new WebPushProvisioningApplePayJWTResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJws

`func (o *WebPushProvisioningApplePayJWTResponse) GetJws() WebPushProvisioningApplePayJWSModel`

GetJws returns the Jws field if non-nil, zero value otherwise.

### GetJwsOk

`func (o *WebPushProvisioningApplePayJWTResponse) GetJwsOk() (*WebPushProvisioningApplePayJWSModel, bool)`

GetJwsOk returns a tuple with the Jws field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJws

`func (o *WebPushProvisioningApplePayJWTResponse) SetJws(v WebPushProvisioningApplePayJWSModel)`

SetJws sets Jws field to given value.


### GetState

`func (o *WebPushProvisioningApplePayJWTResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *WebPushProvisioningApplePayJWTResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *WebPushProvisioningApplePayJWTResponse) SetState(v string)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


