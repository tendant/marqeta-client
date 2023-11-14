/*
Core API

Marqeta's Core API endpoints, conveniently annotated to enable code generation (including SDKs), test cases, and documentation. Currently in beta.

API version: 3.0.11
Contact: support@marqeta.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package coreapi

import (
	"encoding/json"
)

// checks if the Link type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Link{}

// Link Link to balances of related user GPAs.
type Link struct {
	// URL of the requested `/balances` resource.
	Href string `json:"href"`
	// The HTTP method of the link.
	Method string `json:"method"`
	// Specifies the relationship between the current resource and the linked resource.
	Rel string `json:"rel"`
}

// NewLink instantiates a new Link object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLink(href string, method string, rel string) *Link {
	this := Link{}
	this.Href = href
	this.Method = method
	this.Rel = rel
	return &this
}

// NewLinkWithDefaults instantiates a new Link object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkWithDefaults() *Link {
	this := Link{}
	return &this
}

// GetHref returns the Href field value
func (o *Link) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *Link) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *Link) SetHref(v string) {
	o.Href = v
}

// GetMethod returns the Method field value
func (o *Link) GetMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *Link) GetMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *Link) SetMethod(v string) {
	o.Method = v
}

// GetRel returns the Rel field value
func (o *Link) GetRel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Rel
}

// GetRelOk returns a tuple with the Rel field value
// and a boolean to check if the value has been set.
func (o *Link) GetRelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rel, true
}

// SetRel sets field value
func (o *Link) SetRel(v string) {
	o.Rel = v
}

func (o Link) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Link) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["href"] = o.Href
	toSerialize["method"] = o.Method
	toSerialize["rel"] = o.Rel
	return toSerialize, nil
}

type NullableLink struct {
	value *Link
	isSet bool
}

func (v NullableLink) Get() *Link {
	return v.value
}

func (v *NullableLink) Set(val *Link) {
	v.value = val
	v.isSet = true
}

func (v NullableLink) IsSet() bool {
	return v.isSet
}

func (v *NullableLink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLink(val *Link) *NullableLink {
	return &NullableLink{value: val, isSet: true}
}

func (v NullableLink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


