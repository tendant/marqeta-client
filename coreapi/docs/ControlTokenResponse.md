# ControlTokenResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ControlToken** | **string** | Unique value generated as a result of issuing a &#x60;POST&#x60; request to the &#x60;/pins/controltoken&#x60; endpoint. This value cannot be updated. | 

## Methods

### NewControlTokenResponse

`func NewControlTokenResponse(controlToken string, ) *ControlTokenResponse`

NewControlTokenResponse instantiates a new ControlTokenResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControlTokenResponseWithDefaults

`func NewControlTokenResponseWithDefaults() *ControlTokenResponse`

NewControlTokenResponseWithDefaults instantiates a new ControlTokenResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetControlToken

`func (o *ControlTokenResponse) GetControlToken() string`

GetControlToken returns the ControlToken field if non-nil, zero value otherwise.

### GetControlTokenOk

`func (o *ControlTokenResponse) GetControlTokenOk() (*string, bool)`

GetControlTokenOk returns a tuple with the ControlToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlToken

`func (o *ControlTokenResponse) SetControlToken(v string)`

SetControlToken sets ControlToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


