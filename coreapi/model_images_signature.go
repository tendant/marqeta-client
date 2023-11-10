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

// checks if the ImagesSignature type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImagesSignature{}

// ImagesSignature Specifies an image of the cardholder's signature.
type ImagesSignature struct {
	// Specifies a PNG image of the cardholder's signature.
	Name *string `json:"name,omitempty"`
}

// NewImagesSignature instantiates a new ImagesSignature object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImagesSignature() *ImagesSignature {
	this := ImagesSignature{}
	return &this
}

// NewImagesSignatureWithDefaults instantiates a new ImagesSignature object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImagesSignatureWithDefaults() *ImagesSignature {
	this := ImagesSignature{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ImagesSignature) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImagesSignature) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ImagesSignature) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ImagesSignature) SetName(v string) {
	o.Name = &v
}

func (o ImagesSignature) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImagesSignature) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableImagesSignature struct {
	value *ImagesSignature
	isSet bool
}

func (v NullableImagesSignature) Get() *ImagesSignature {
	return v.value
}

func (v *NullableImagesSignature) Set(val *ImagesSignature) {
	v.value = val
	v.isSet = true
}

func (v NullableImagesSignature) IsSet() bool {
	return v.isSet
}

func (v *NullableImagesSignature) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImagesSignature(val *ImagesSignature) *NullableImagesSignature {
	return &NullableImagesSignature{value: val, isSet: true}
}

func (v NullableImagesSignature) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImagesSignature) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

