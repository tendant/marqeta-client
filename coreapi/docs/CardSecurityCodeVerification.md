# CardSecurityCodeVerification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Response** | [**Response**](Response.md) |  | 
**Type** | **string** | Indicates the type of security code. Can have these possible values:  * *CVV1*  the security code stored in the magnetic stripe on the card. * *CVV2*  the security code printed on the card. * *ICVV*  the security code stored on the chip of the card. * *DCVV*  a dynamic security code used in some contactless payments when a card or device is tapped against the card reader. | 

## Methods

### NewCardSecurityCodeVerification

`func NewCardSecurityCodeVerification(response Response, type_ string, ) *CardSecurityCodeVerification`

NewCardSecurityCodeVerification instantiates a new CardSecurityCodeVerification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardSecurityCodeVerificationWithDefaults

`func NewCardSecurityCodeVerificationWithDefaults() *CardSecurityCodeVerification`

NewCardSecurityCodeVerificationWithDefaults instantiates a new CardSecurityCodeVerification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponse

`func (o *CardSecurityCodeVerification) GetResponse() Response`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *CardSecurityCodeVerification) GetResponseOk() (*Response, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *CardSecurityCodeVerification) SetResponse(v Response)`

SetResponse sets Response field to given value.


### GetType

`func (o *CardSecurityCodeVerification) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CardSecurityCodeVerification) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CardSecurityCodeVerification) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


