# Poi

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Atm** | Pointer to **bool** | If set to &#x60;true&#x60;, cards can be used for withdrawing cash at an ATM and for receiving cash back at a point of sale (POS). | [optional] [default to false]
**Ecommerce** | Pointer to **bool** | If set to &#x60;true&#x60;, cards can be used for online purchases. | [optional] [default to true]
**Other** | Pointer to [**OtherPoi**](OtherPoi.md) |  | [optional] 

## Methods

### NewPoi

`func NewPoi() *Poi`

NewPoi instantiates a new Poi object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoiWithDefaults

`func NewPoiWithDefaults() *Poi`

NewPoiWithDefaults instantiates a new Poi object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAtm

`func (o *Poi) GetAtm() bool`

GetAtm returns the Atm field if non-nil, zero value otherwise.

### GetAtmOk

`func (o *Poi) GetAtmOk() (*bool, bool)`

GetAtmOk returns a tuple with the Atm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtm

`func (o *Poi) SetAtm(v bool)`

SetAtm sets Atm field to given value.

### HasAtm

`func (o *Poi) HasAtm() bool`

HasAtm returns a boolean if a field has been set.

### GetEcommerce

`func (o *Poi) GetEcommerce() bool`

GetEcommerce returns the Ecommerce field if non-nil, zero value otherwise.

### GetEcommerceOk

`func (o *Poi) GetEcommerceOk() (*bool, bool)`

GetEcommerceOk returns a tuple with the Ecommerce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcommerce

`func (o *Poi) SetEcommerce(v bool)`

SetEcommerce sets Ecommerce field to given value.

### HasEcommerce

`func (o *Poi) HasEcommerce() bool`

HasEcommerce returns a boolean if a field has been set.

### GetOther

`func (o *Poi) GetOther() OtherPoi`

GetOther returns the Other field if non-nil, zero value otherwise.

### GetOtherOk

`func (o *Poi) GetOtherOk() (*OtherPoi, bool)`

GetOtherOk returns a tuple with the Other field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOther

`func (o *Poi) SetOther(v OtherPoi)`

SetOther sets Other field to given value.

### HasOther

`func (o *Poi) HasOther() bool`

HasOther returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


