# OtherPoi

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Allow** | Pointer to **bool** | If set to &#x60;true&#x60;, card transactions at points of interaction other than e-commerce or ATMs are allowed. This group includes points of sale (POS). | [optional] [default to true]
**CardPresenceRequired** | Pointer to **bool** | If set to &#x60;true&#x60;, cards of this card product type are required to be present during the transaction, such as in IVR scenarios. | [optional] [default to false]
**CardholderPresenceRequired** | Pointer to **bool** | If set to &#x60;true&#x60;, the cardholder is required to be present during the transaction, such as in a restaurant where the card is present but the cardholder might not be present when the card is swiped. | [optional] [default to false]
**Track1DiscretionaryData** | Pointer to **string** |  | [optional] 
**Track2DiscretionaryData** | Pointer to **string** |  | [optional] 
**UseStaticPin** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewOtherPoi

`func NewOtherPoi() *OtherPoi`

NewOtherPoi instantiates a new OtherPoi object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOtherPoiWithDefaults

`func NewOtherPoiWithDefaults() *OtherPoi`

NewOtherPoiWithDefaults instantiates a new OtherPoi object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllow

`func (o *OtherPoi) GetAllow() bool`

GetAllow returns the Allow field if non-nil, zero value otherwise.

### GetAllowOk

`func (o *OtherPoi) GetAllowOk() (*bool, bool)`

GetAllowOk returns a tuple with the Allow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllow

`func (o *OtherPoi) SetAllow(v bool)`

SetAllow sets Allow field to given value.

### HasAllow

`func (o *OtherPoi) HasAllow() bool`

HasAllow returns a boolean if a field has been set.

### GetCardPresenceRequired

`func (o *OtherPoi) GetCardPresenceRequired() bool`

GetCardPresenceRequired returns the CardPresenceRequired field if non-nil, zero value otherwise.

### GetCardPresenceRequiredOk

`func (o *OtherPoi) GetCardPresenceRequiredOk() (*bool, bool)`

GetCardPresenceRequiredOk returns a tuple with the CardPresenceRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardPresenceRequired

`func (o *OtherPoi) SetCardPresenceRequired(v bool)`

SetCardPresenceRequired sets CardPresenceRequired field to given value.

### HasCardPresenceRequired

`func (o *OtherPoi) HasCardPresenceRequired() bool`

HasCardPresenceRequired returns a boolean if a field has been set.

### GetCardholderPresenceRequired

`func (o *OtherPoi) GetCardholderPresenceRequired() bool`

GetCardholderPresenceRequired returns the CardholderPresenceRequired field if non-nil, zero value otherwise.

### GetCardholderPresenceRequiredOk

`func (o *OtherPoi) GetCardholderPresenceRequiredOk() (*bool, bool)`

GetCardholderPresenceRequiredOk returns a tuple with the CardholderPresenceRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardholderPresenceRequired

`func (o *OtherPoi) SetCardholderPresenceRequired(v bool)`

SetCardholderPresenceRequired sets CardholderPresenceRequired field to given value.

### HasCardholderPresenceRequired

`func (o *OtherPoi) HasCardholderPresenceRequired() bool`

HasCardholderPresenceRequired returns a boolean if a field has been set.

### GetTrack1DiscretionaryData

`func (o *OtherPoi) GetTrack1DiscretionaryData() string`

GetTrack1DiscretionaryData returns the Track1DiscretionaryData field if non-nil, zero value otherwise.

### GetTrack1DiscretionaryDataOk

`func (o *OtherPoi) GetTrack1DiscretionaryDataOk() (*string, bool)`

GetTrack1DiscretionaryDataOk returns a tuple with the Track1DiscretionaryData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrack1DiscretionaryData

`func (o *OtherPoi) SetTrack1DiscretionaryData(v string)`

SetTrack1DiscretionaryData sets Track1DiscretionaryData field to given value.

### HasTrack1DiscretionaryData

`func (o *OtherPoi) HasTrack1DiscretionaryData() bool`

HasTrack1DiscretionaryData returns a boolean if a field has been set.

### GetTrack2DiscretionaryData

`func (o *OtherPoi) GetTrack2DiscretionaryData() string`

GetTrack2DiscretionaryData returns the Track2DiscretionaryData field if non-nil, zero value otherwise.

### GetTrack2DiscretionaryDataOk

`func (o *OtherPoi) GetTrack2DiscretionaryDataOk() (*string, bool)`

GetTrack2DiscretionaryDataOk returns a tuple with the Track2DiscretionaryData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrack2DiscretionaryData

`func (o *OtherPoi) SetTrack2DiscretionaryData(v string)`

SetTrack2DiscretionaryData sets Track2DiscretionaryData field to given value.

### HasTrack2DiscretionaryData

`func (o *OtherPoi) HasTrack2DiscretionaryData() bool`

HasTrack2DiscretionaryData returns a boolean if a field has been set.

### GetUseStaticPin

`func (o *OtherPoi) GetUseStaticPin() bool`

GetUseStaticPin returns the UseStaticPin field if non-nil, zero value otherwise.

### GetUseStaticPinOk

`func (o *OtherPoi) GetUseStaticPinOk() (*bool, bool)`

GetUseStaticPinOk returns a tuple with the UseStaticPin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseStaticPin

`func (o *OtherPoi) SetUseStaticPin(v bool)`

SetUseStaticPin sets UseStaticPin field to given value.

### HasUseStaticPin

`func (o *OtherPoi) HasUseStaticPin() bool`

HasUseStaticPin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


