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

// checks if the DirectDepositAccountListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DirectDepositAccountListResponse{}

// DirectDepositAccountListResponse struct for DirectDepositAccountListResponse
type DirectDepositAccountListResponse struct {
	Count *int32 `json:"count,omitempty"`
	Data []DirectDepositAccountResponse `json:"data,omitempty"`
	EndIndex *int32 `json:"end_index,omitempty"`
	IsMore *bool `json:"is_more,omitempty"`
	StartIndex *int32 `json:"start_index,omitempty"`
}

// NewDirectDepositAccountListResponse instantiates a new DirectDepositAccountListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDirectDepositAccountListResponse() *DirectDepositAccountListResponse {
	this := DirectDepositAccountListResponse{}
	var isMore bool = false
	this.IsMore = &isMore
	return &this
}

// NewDirectDepositAccountListResponseWithDefaults instantiates a new DirectDepositAccountListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDirectDepositAccountListResponseWithDefaults() *DirectDepositAccountListResponse {
	this := DirectDepositAccountListResponse{}
	var isMore bool = false
	this.IsMore = &isMore
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *DirectDepositAccountListResponse) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectDepositAccountListResponse) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *DirectDepositAccountListResponse) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *DirectDepositAccountListResponse) SetCount(v int32) {
	o.Count = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *DirectDepositAccountListResponse) GetData() []DirectDepositAccountResponse {
	if o == nil || IsNil(o.Data) {
		var ret []DirectDepositAccountResponse
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectDepositAccountListResponse) GetDataOk() ([]DirectDepositAccountResponse, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *DirectDepositAccountListResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []DirectDepositAccountResponse and assigns it to the Data field.
func (o *DirectDepositAccountListResponse) SetData(v []DirectDepositAccountResponse) {
	o.Data = v
}

// GetEndIndex returns the EndIndex field value if set, zero value otherwise.
func (o *DirectDepositAccountListResponse) GetEndIndex() int32 {
	if o == nil || IsNil(o.EndIndex) {
		var ret int32
		return ret
	}
	return *o.EndIndex
}

// GetEndIndexOk returns a tuple with the EndIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectDepositAccountListResponse) GetEndIndexOk() (*int32, bool) {
	if o == nil || IsNil(o.EndIndex) {
		return nil, false
	}
	return o.EndIndex, true
}

// HasEndIndex returns a boolean if a field has been set.
func (o *DirectDepositAccountListResponse) HasEndIndex() bool {
	if o != nil && !IsNil(o.EndIndex) {
		return true
	}

	return false
}

// SetEndIndex gets a reference to the given int32 and assigns it to the EndIndex field.
func (o *DirectDepositAccountListResponse) SetEndIndex(v int32) {
	o.EndIndex = &v
}

// GetIsMore returns the IsMore field value if set, zero value otherwise.
func (o *DirectDepositAccountListResponse) GetIsMore() bool {
	if o == nil || IsNil(o.IsMore) {
		var ret bool
		return ret
	}
	return *o.IsMore
}

// GetIsMoreOk returns a tuple with the IsMore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectDepositAccountListResponse) GetIsMoreOk() (*bool, bool) {
	if o == nil || IsNil(o.IsMore) {
		return nil, false
	}
	return o.IsMore, true
}

// HasIsMore returns a boolean if a field has been set.
func (o *DirectDepositAccountListResponse) HasIsMore() bool {
	if o != nil && !IsNil(o.IsMore) {
		return true
	}

	return false
}

// SetIsMore gets a reference to the given bool and assigns it to the IsMore field.
func (o *DirectDepositAccountListResponse) SetIsMore(v bool) {
	o.IsMore = &v
}

// GetStartIndex returns the StartIndex field value if set, zero value otherwise.
func (o *DirectDepositAccountListResponse) GetStartIndex() int32 {
	if o == nil || IsNil(o.StartIndex) {
		var ret int32
		return ret
	}
	return *o.StartIndex
}

// GetStartIndexOk returns a tuple with the StartIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectDepositAccountListResponse) GetStartIndexOk() (*int32, bool) {
	if o == nil || IsNil(o.StartIndex) {
		return nil, false
	}
	return o.StartIndex, true
}

// HasStartIndex returns a boolean if a field has been set.
func (o *DirectDepositAccountListResponse) HasStartIndex() bool {
	if o != nil && !IsNil(o.StartIndex) {
		return true
	}

	return false
}

// SetStartIndex gets a reference to the given int32 and assigns it to the StartIndex field.
func (o *DirectDepositAccountListResponse) SetStartIndex(v int32) {
	o.StartIndex = &v
}

func (o DirectDepositAccountListResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DirectDepositAccountListResponse) ToMap() (map[string]interface{}, error) {
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

type NullableDirectDepositAccountListResponse struct {
	value *DirectDepositAccountListResponse
	isSet bool
}

func (v NullableDirectDepositAccountListResponse) Get() *DirectDepositAccountListResponse {
	return v.value
}

func (v *NullableDirectDepositAccountListResponse) Set(val *DirectDepositAccountListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDirectDepositAccountListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDirectDepositAccountListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDirectDepositAccountListResponse(val *DirectDepositAccountListResponse) *NullableDirectDepositAccountListResponse {
	return &NullableDirectDepositAccountListResponse{value: val, isSet: true}
}

func (v NullableDirectDepositAccountListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDirectDepositAccountListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


