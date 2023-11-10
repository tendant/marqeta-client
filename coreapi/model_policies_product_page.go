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

// checks if the PoliciesProductPage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PoliciesProductPage{}

// PoliciesProductPage List response details for product policies.
type PoliciesProductPage struct {
	// Number of resources returned.
	Count int32 `json:"count"`
	// One or more credit product policies.
	Data []PolicyProductResponse `json:"data"`
	// Sort order index of the last resource in the returned array.
	EndIndex int32 `json:"end_index"`
	// A value of `true` indicates that more unreturned resources exist.
	IsMore bool `json:"is_more"`
	// Sort order index of the first resource in the returned array.
	StartIndex int32 `json:"start_index"`
}

// NewPoliciesProductPage instantiates a new PoliciesProductPage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPoliciesProductPage(count int32, data []PolicyProductResponse, endIndex int32, isMore bool, startIndex int32) *PoliciesProductPage {
	this := PoliciesProductPage{}
	this.Count = count
	this.Data = data
	this.EndIndex = endIndex
	this.IsMore = isMore
	this.StartIndex = startIndex
	return &this
}

// NewPoliciesProductPageWithDefaults instantiates a new PoliciesProductPage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoliciesProductPageWithDefaults() *PoliciesProductPage {
	this := PoliciesProductPage{}
	return &this
}

// GetCount returns the Count field value
func (o *PoliciesProductPage) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *PoliciesProductPage) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *PoliciesProductPage) SetCount(v int32) {
	o.Count = v
}

// GetData returns the Data field value
func (o *PoliciesProductPage) GetData() []PolicyProductResponse {
	if o == nil {
		var ret []PolicyProductResponse
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PoliciesProductPage) GetDataOk() ([]PolicyProductResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *PoliciesProductPage) SetData(v []PolicyProductResponse) {
	o.Data = v
}

// GetEndIndex returns the EndIndex field value
func (o *PoliciesProductPage) GetEndIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EndIndex
}

// GetEndIndexOk returns a tuple with the EndIndex field value
// and a boolean to check if the value has been set.
func (o *PoliciesProductPage) GetEndIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndIndex, true
}

// SetEndIndex sets field value
func (o *PoliciesProductPage) SetEndIndex(v int32) {
	o.EndIndex = v
}

// GetIsMore returns the IsMore field value
func (o *PoliciesProductPage) GetIsMore() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsMore
}

// GetIsMoreOk returns a tuple with the IsMore field value
// and a boolean to check if the value has been set.
func (o *PoliciesProductPage) GetIsMoreOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsMore, true
}

// SetIsMore sets field value
func (o *PoliciesProductPage) SetIsMore(v bool) {
	o.IsMore = v
}

// GetStartIndex returns the StartIndex field value
func (o *PoliciesProductPage) GetStartIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.StartIndex
}

// GetStartIndexOk returns a tuple with the StartIndex field value
// and a boolean to check if the value has been set.
func (o *PoliciesProductPage) GetStartIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartIndex, true
}

// SetStartIndex sets field value
func (o *PoliciesProductPage) SetStartIndex(v int32) {
	o.StartIndex = v
}

func (o PoliciesProductPage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PoliciesProductPage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["count"] = o.Count
	toSerialize["data"] = o.Data
	toSerialize["end_index"] = o.EndIndex
	toSerialize["is_more"] = o.IsMore
	toSerialize["start_index"] = o.StartIndex
	return toSerialize, nil
}

type NullablePoliciesProductPage struct {
	value *PoliciesProductPage
	isSet bool
}

func (v NullablePoliciesProductPage) Get() *PoliciesProductPage {
	return v.value
}

func (v *NullablePoliciesProductPage) Set(val *PoliciesProductPage) {
	v.value = val
	v.isSet = true
}

func (v NullablePoliciesProductPage) IsSet() bool {
	return v.isSet
}

func (v *NullablePoliciesProductPage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePoliciesProductPage(val *PoliciesProductPage) *NullablePoliciesProductPage {
	return &NullablePoliciesProductPage{value: val, isSet: true}
}

func (v NullablePoliciesProductPage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePoliciesProductPage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


