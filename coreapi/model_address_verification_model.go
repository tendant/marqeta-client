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

// checks if the AddressVerificationModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddressVerificationModel{}

// AddressVerificationModel Contains address verification data consisting of address data entered by the cardholder, address data held by the Marqeta platform, and an assertion by the Marqeta platform as to whether the two sets of data match.
type AddressVerificationModel struct {
	OnFile *AvsInformation `json:"on_file,omitempty"`
	Request *AvsInformation `json:"request,omitempty"`
	Response *Response `json:"response,omitempty"`
}

// NewAddressVerificationModel instantiates a new AddressVerificationModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressVerificationModel() *AddressVerificationModel {
	this := AddressVerificationModel{}
	return &this
}

// NewAddressVerificationModelWithDefaults instantiates a new AddressVerificationModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressVerificationModelWithDefaults() *AddressVerificationModel {
	this := AddressVerificationModel{}
	return &this
}

// GetOnFile returns the OnFile field value if set, zero value otherwise.
func (o *AddressVerificationModel) GetOnFile() AvsInformation {
	if o == nil || IsNil(o.OnFile) {
		var ret AvsInformation
		return ret
	}
	return *o.OnFile
}

// GetOnFileOk returns a tuple with the OnFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressVerificationModel) GetOnFileOk() (*AvsInformation, bool) {
	if o == nil || IsNil(o.OnFile) {
		return nil, false
	}
	return o.OnFile, true
}

// HasOnFile returns a boolean if a field has been set.
func (o *AddressVerificationModel) HasOnFile() bool {
	if o != nil && !IsNil(o.OnFile) {
		return true
	}

	return false
}

// SetOnFile gets a reference to the given AvsInformation and assigns it to the OnFile field.
func (o *AddressVerificationModel) SetOnFile(v AvsInformation) {
	o.OnFile = &v
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *AddressVerificationModel) GetRequest() AvsInformation {
	if o == nil || IsNil(o.Request) {
		var ret AvsInformation
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressVerificationModel) GetRequestOk() (*AvsInformation, bool) {
	if o == nil || IsNil(o.Request) {
		return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *AddressVerificationModel) HasRequest() bool {
	if o != nil && !IsNil(o.Request) {
		return true
	}

	return false
}

// SetRequest gets a reference to the given AvsInformation and assigns it to the Request field.
func (o *AddressVerificationModel) SetRequest(v AvsInformation) {
	o.Request = &v
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *AddressVerificationModel) GetResponse() Response {
	if o == nil || IsNil(o.Response) {
		var ret Response
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressVerificationModel) GetResponseOk() (*Response, bool) {
	if o == nil || IsNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *AddressVerificationModel) HasResponse() bool {
	if o != nil && !IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given Response and assigns it to the Response field.
func (o *AddressVerificationModel) SetResponse(v Response) {
	o.Response = &v
}

func (o AddressVerificationModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddressVerificationModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OnFile) {
		toSerialize["on_file"] = o.OnFile
	}
	if !IsNil(o.Request) {
		toSerialize["request"] = o.Request
	}
	if !IsNil(o.Response) {
		toSerialize["response"] = o.Response
	}
	return toSerialize, nil
}

type NullableAddressVerificationModel struct {
	value *AddressVerificationModel
	isSet bool
}

func (v NullableAddressVerificationModel) Get() *AddressVerificationModel {
	return v.value
}

func (v *NullableAddressVerificationModel) Set(val *AddressVerificationModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressVerificationModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressVerificationModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressVerificationModel(val *AddressVerificationModel) *NullableAddressVerificationModel {
	return &NullableAddressVerificationModel{value: val, isSet: true}
}

func (v NullableAddressVerificationModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressVerificationModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

