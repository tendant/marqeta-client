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

// checks if the FeeListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FeeListResponse{}

// FeeListResponse struct for FeeListResponse
type FeeListResponse struct {
	// Number of resources to retrieve.  This field is returned if there are resources in your returned array.
	Count *int32 `json:"count,omitempty"`
	// An array of fee objects.  Objects are returned as appropriate to your query.
	Data []FeeResponse `json:"data,omitempty"`
	// Sort order index of the last resource in the returned array.  This field is returned if there are resources in your returned array.
	EndIndex *int32 `json:"end_index,omitempty"`
	// A value of `true` indicates that more unreturned resources exist. A value of `false` indicates that no more unreturned resources exist.  This field is returned if there are resources in your returned array.
	IsMore *bool `json:"is_more,omitempty"`
	// Sort order index of the first resource in the returned array.  This field is returned if there are resources in your returned array.
	StartIndex *int32 `json:"start_index,omitempty"`
}

// NewFeeListResponse instantiates a new FeeListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeeListResponse() *FeeListResponse {
	this := FeeListResponse{}
	var isMore bool = false
	this.IsMore = &isMore
	return &this
}

// NewFeeListResponseWithDefaults instantiates a new FeeListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeeListResponseWithDefaults() *FeeListResponse {
	this := FeeListResponse{}
	var isMore bool = false
	this.IsMore = &isMore
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *FeeListResponse) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeeListResponse) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *FeeListResponse) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *FeeListResponse) SetCount(v int32) {
	o.Count = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *FeeListResponse) GetData() []FeeResponse {
	if o == nil || IsNil(o.Data) {
		var ret []FeeResponse
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeeListResponse) GetDataOk() ([]FeeResponse, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *FeeListResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []FeeResponse and assigns it to the Data field.
func (o *FeeListResponse) SetData(v []FeeResponse) {
	o.Data = v
}

// GetEndIndex returns the EndIndex field value if set, zero value otherwise.
func (o *FeeListResponse) GetEndIndex() int32 {
	if o == nil || IsNil(o.EndIndex) {
		var ret int32
		return ret
	}
	return *o.EndIndex
}

// GetEndIndexOk returns a tuple with the EndIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeeListResponse) GetEndIndexOk() (*int32, bool) {
	if o == nil || IsNil(o.EndIndex) {
		return nil, false
	}
	return o.EndIndex, true
}

// HasEndIndex returns a boolean if a field has been set.
func (o *FeeListResponse) HasEndIndex() bool {
	if o != nil && !IsNil(o.EndIndex) {
		return true
	}

	return false
}

// SetEndIndex gets a reference to the given int32 and assigns it to the EndIndex field.
func (o *FeeListResponse) SetEndIndex(v int32) {
	o.EndIndex = &v
}

// GetIsMore returns the IsMore field value if set, zero value otherwise.
func (o *FeeListResponse) GetIsMore() bool {
	if o == nil || IsNil(o.IsMore) {
		var ret bool
		return ret
	}
	return *o.IsMore
}

// GetIsMoreOk returns a tuple with the IsMore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeeListResponse) GetIsMoreOk() (*bool, bool) {
	if o == nil || IsNil(o.IsMore) {
		return nil, false
	}
	return o.IsMore, true
}

// HasIsMore returns a boolean if a field has been set.
func (o *FeeListResponse) HasIsMore() bool {
	if o != nil && !IsNil(o.IsMore) {
		return true
	}

	return false
}

// SetIsMore gets a reference to the given bool and assigns it to the IsMore field.
func (o *FeeListResponse) SetIsMore(v bool) {
	o.IsMore = &v
}

// GetStartIndex returns the StartIndex field value if set, zero value otherwise.
func (o *FeeListResponse) GetStartIndex() int32 {
	if o == nil || IsNil(o.StartIndex) {
		var ret int32
		return ret
	}
	return *o.StartIndex
}

// GetStartIndexOk returns a tuple with the StartIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeeListResponse) GetStartIndexOk() (*int32, bool) {
	if o == nil || IsNil(o.StartIndex) {
		return nil, false
	}
	return o.StartIndex, true
}

// HasStartIndex returns a boolean if a field has been set.
func (o *FeeListResponse) HasStartIndex() bool {
	if o != nil && !IsNil(o.StartIndex) {
		return true
	}

	return false
}

// SetStartIndex gets a reference to the given int32 and assigns it to the StartIndex field.
func (o *FeeListResponse) SetStartIndex(v int32) {
	o.StartIndex = &v
}

func (o FeeListResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FeeListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.EndIndex) {
		toSerialize["end_index"] = o.EndIndex
	}
	if !IsNil(o.IsMore) {
		toSerialize["is_more"] = o.IsMore
	}
	if !IsNil(o.StartIndex) {
		toSerialize["start_index"] = o.StartIndex
	}
	return toSerialize, nil
}

type NullableFeeListResponse struct {
	value *FeeListResponse
	isSet bool
}

func (v NullableFeeListResponse) Get() *FeeListResponse {
	return v.value
}

func (v *NullableFeeListResponse) Set(val *FeeListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFeeListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFeeListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeeListResponse(val *FeeListResponse) *NullableFeeListResponse {
	return &NullableFeeListResponse{value: val, isSet: true}
}

func (v NullableFeeListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeeListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


