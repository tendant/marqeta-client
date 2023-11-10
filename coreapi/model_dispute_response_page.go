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

// checks if the DisputeResponsePage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DisputeResponsePage{}

// DisputeResponsePage Return paginated disputes.
type DisputeResponsePage struct {
	// Number of resources returned.
	Count int32 `json:"count"`
	// Contains one or more disputes on a credit account.
	Data []DisputeResponse `json:"data"`
	// Sort order index of the last resource in the returned array.
	EndIndex int32 `json:"end_index"`
	// A value of `true` indicates that more unreturned resources exist.
	IsMore bool `json:"is_more"`
	// Sort order index of the first resource in the returned array.
	StartIndex int32 `json:"start_index"`
}

// NewDisputeResponsePage instantiates a new DisputeResponsePage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDisputeResponsePage(count int32, data []DisputeResponse, endIndex int32, isMore bool, startIndex int32) *DisputeResponsePage {
	this := DisputeResponsePage{}
	this.Count = count
	this.Data = data
	this.EndIndex = endIndex
	this.IsMore = isMore
	this.StartIndex = startIndex
	return &this
}

// NewDisputeResponsePageWithDefaults instantiates a new DisputeResponsePage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDisputeResponsePageWithDefaults() *DisputeResponsePage {
	this := DisputeResponsePage{}
	return &this
}

// GetCount returns the Count field value
func (o *DisputeResponsePage) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *DisputeResponsePage) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *DisputeResponsePage) SetCount(v int32) {
	o.Count = v
}

// GetData returns the Data field value
func (o *DisputeResponsePage) GetData() []DisputeResponse {
	if o == nil {
		var ret []DisputeResponse
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *DisputeResponsePage) GetDataOk() ([]DisputeResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *DisputeResponsePage) SetData(v []DisputeResponse) {
	o.Data = v
}

// GetEndIndex returns the EndIndex field value
func (o *DisputeResponsePage) GetEndIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EndIndex
}

// GetEndIndexOk returns a tuple with the EndIndex field value
// and a boolean to check if the value has been set.
func (o *DisputeResponsePage) GetEndIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndIndex, true
}

// SetEndIndex sets field value
func (o *DisputeResponsePage) SetEndIndex(v int32) {
	o.EndIndex = v
}

// GetIsMore returns the IsMore field value
func (o *DisputeResponsePage) GetIsMore() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsMore
}

// GetIsMoreOk returns a tuple with the IsMore field value
// and a boolean to check if the value has been set.
func (o *DisputeResponsePage) GetIsMoreOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsMore, true
}

// SetIsMore sets field value
func (o *DisputeResponsePage) SetIsMore(v bool) {
	o.IsMore = v
}

// GetStartIndex returns the StartIndex field value
func (o *DisputeResponsePage) GetStartIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.StartIndex
}

// GetStartIndexOk returns a tuple with the StartIndex field value
// and a boolean to check if the value has been set.
func (o *DisputeResponsePage) GetStartIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartIndex, true
}

// SetStartIndex sets field value
func (o *DisputeResponsePage) SetStartIndex(v int32) {
	o.StartIndex = v
}

func (o DisputeResponsePage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DisputeResponsePage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["count"] = o.Count
	toSerialize["data"] = o.Data
	toSerialize["end_index"] = o.EndIndex
	toSerialize["is_more"] = o.IsMore
	toSerialize["start_index"] = o.StartIndex
	return toSerialize, nil
}

type NullableDisputeResponsePage struct {
	value *DisputeResponsePage
	isSet bool
}

func (v NullableDisputeResponsePage) Get() *DisputeResponsePage {
	return v.value
}

func (v *NullableDisputeResponsePage) Set(val *DisputeResponsePage) {
	v.value = val
	v.isSet = true
}

func (v NullableDisputeResponsePage) IsSet() bool {
	return v.isSet
}

func (v *NullableDisputeResponsePage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDisputeResponsePage(val *DisputeResponsePage) *NullableDisputeResponsePage {
	return &NullableDisputeResponsePage{value: val, isSet: true}
}

func (v NullableDisputeResponsePage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDisputeResponsePage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

