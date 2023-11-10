# ControlTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardToken** | **string** | The unique identifier of the card for which you want to generate a control token. | 
**ControltokenType** | Pointer to **string** | Specifies the type of action completed by this request.  *WARNING:* Sending a request to this endpoint with a &#x60;REVEAL_PIN&#x60; control token requires PCI DSS compliance.  The lifespan of the control token depends on the token type:  * *SET_PIN:* 60 minutes * *REVEAL_PIN:* 5 minutes | [optional] 

## Methods

### NewControlTokenRequest

`func NewControlTokenRequest(cardToken string, ) *ControlTokenRequest`

NewControlTokenRequest instantiates a new ControlTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControlTokenRequestWithDefaults

`func NewControlTokenRequestWithDefaults() *ControlTokenRequest`

NewControlTokenRequestWithDefaults instantiates a new ControlTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardToken

`func (o *ControlTokenRequest) GetCardToken() string`

GetCardToken returns the CardToken field if non-nil, zero value otherwise.

### GetCardTokenOk

`func (o *ControlTokenRequest) GetCardTokenOk() (*string, bool)`

GetCardTokenOk returns a tuple with the CardToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardToken

`func (o *ControlTokenRequest) SetCardToken(v string)`

SetCardToken sets CardToken field to given value.


### GetControltokenType

`func (o *ControlTokenRequest) GetControltokenType() string`

GetControltokenType returns the ControltokenType field if non-nil, zero value otherwise.

### GetControltokenTypeOk

`func (o *ControlTokenRequest) GetControltokenTypeOk() (*string, bool)`

GetControltokenTypeOk returns a tuple with the ControltokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControltokenType

`func (o *ControlTokenRequest) SetControltokenType(v string)`

SetControltokenType sets ControltokenType field to given value.

### HasControltokenType

`func (o *ControlTokenRequest) HasControltokenType() bool`

HasControltokenType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


