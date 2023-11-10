/*
Core API

Marqeta's Core API endpoints, conveniently annotated to enable code generation (including SDKs), test cases, and documentation. Currently in beta.

API version: 3.0.11
Contact: support@marqeta.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package marqeta_coreapi_client

import (
	"encoding/json"
)

// checks if the ImagesCard type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImagesCard{}

// ImagesCard Specifies personalized images that appear on the card.
type ImagesCard struct {
	// Specifies a PNG image to display on the card.
	Name *string `json:"name,omitempty"`
	// Specifies the color of the image displayed on the card.
	ThermalColor *string `json:"thermal_color,omitempty"`
}

// NewImagesCard instantiates a new ImagesCard object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImagesCard() *ImagesCard {
	this := ImagesCard{}
	return &this
}

// NewImagesCardWithDefaults instantiates a new ImagesCard object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImagesCardWithDefaults() *ImagesCard {
	this := ImagesCard{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ImagesCard) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImagesCard) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ImagesCard) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ImagesCard) SetName(v string) {
	o.Name = &v
}

// GetThermalColor returns the ThermalColor field value if set, zero value otherwise.
func (o *ImagesCard) GetThermalColor() string {
	if o == nil || IsNil(o.ThermalColor) {
		var ret string
		return ret
	}
	return *o.ThermalColor
}

// GetThermalColorOk returns a tuple with the ThermalColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImagesCard) GetThermalColorOk() (*string, bool) {
	if o == nil || IsNil(o.ThermalColor) {
		return nil, false
	}
	return o.ThermalColor, true
}

// HasThermalColor returns a boolean if a field has been set.
func (o *ImagesCard) HasThermalColor() bool {
	if o != nil && !IsNil(o.ThermalColor) {
		return true
	}

	return false
}

// SetThermalColor gets a reference to the given string and assigns it to the ThermalColor field.
func (o *ImagesCard) SetThermalColor(v string) {
	o.ThermalColor = &v
}

func (o ImagesCard) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImagesCard) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ThermalColor) {
		toSerialize["thermal_color"] = o.ThermalColor
	}
	return toSerialize, nil
}

type NullableImagesCard struct {
	value *ImagesCard
	isSet bool
}

func (v NullableImagesCard) Get() *ImagesCard {
	return v.value
}

func (v *NullableImagesCard) Set(val *ImagesCard) {
	v.value = val
	v.isSet = true
}

func (v NullableImagesCard) IsSet() bool {
	return v.isSet
}

func (v *NullableImagesCard) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImagesCard(val *ImagesCard) *NullableImagesCard {
	return &NullableImagesCard{value: val, isSet: true}
}

func (v NullableImagesCard) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImagesCard) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


