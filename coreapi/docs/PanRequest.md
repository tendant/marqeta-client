# PanRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CvvNumber** | Pointer to **string** | Three-digit card verification value (CVV2) included on the back of the card.  This value cannot be updated. | [optional] 
**Expiration** | Pointer to **string** | Card expiration date. | [optional] 
**Pan** | **string** | Primary account number (PAN) of the card whose information you want to retrieve.  Send a &#x60;GET&#x60; request to &#x60;/cards/{token}/showpan&#x60; to retrieve the PAN for a specific card. | 

## Methods

### NewPanRequest

`func NewPanRequest(pan string, ) *PanRequest`

NewPanRequest instantiates a new PanRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPanRequestWithDefaults

`func NewPanRequestWithDefaults() *PanRequest`

NewPanRequestWithDefaults instantiates a new PanRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCvvNumber

`func (o *PanRequest) GetCvvNumber() string`

GetCvvNumber returns the CvvNumber field if non-nil, zero value otherwise.

### GetCvvNumberOk

`func (o *PanRequest) GetCvvNumberOk() (*string, bool)`

GetCvvNumberOk returns a tuple with the CvvNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvvNumber

`func (o *PanRequest) SetCvvNumber(v string)`

SetCvvNumber sets CvvNumber field to given value.

### HasCvvNumber

`func (o *PanRequest) HasCvvNumber() bool`

HasCvvNumber returns a boolean if a field has been set.

### GetExpiration

`func (o *PanRequest) GetExpiration() string`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *PanRequest) GetExpirationOk() (*string, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *PanRequest) SetExpiration(v string)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *PanRequest) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### GetPan

`func (o *PanRequest) GetPan() string`

GetPan returns the Pan field if non-nil, zero value otherwise.

### GetPanOk

`func (o *PanRequest) GetPanOk() (*string, bool)`

GetPanOk returns a tuple with the Pan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPan

`func (o *PanRequest) SetPan(v string)`

SetPan sets Pan field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


