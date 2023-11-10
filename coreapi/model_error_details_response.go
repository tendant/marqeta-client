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

// checks if the ErrorDetailsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorDetailsResponse{}

// ErrorDetailsResponse Contains details on an application error.
type ErrorDetailsResponse struct {
	// Unique identifier of the application that resulted in an error.
	ApplicationToken string `json:"application_token"`
	// Message describing the error.
	Message string `json:"message"`
	// Unique identifier of the error details.
	Token string `json:"token"`
	// Type of error.
	Type string `json:"type"`
}

// NewErrorDetailsResponse instantiates a new ErrorDetailsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorDetailsResponse(applicationToken string, message string, token string, type_ string) *ErrorDetailsResponse {
	this := ErrorDetailsResponse{}
	this.ApplicationToken = applicationToken
	this.Message = message
	this.Token = token
	this.Type = type_
	return &this
}

// NewErrorDetailsResponseWithDefaults instantiates a new ErrorDetailsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorDetailsResponseWithDefaults() *ErrorDetailsResponse {
	this := ErrorDetailsResponse{}
	return &this
}

// GetApplicationToken returns the ApplicationToken field value
func (o *ErrorDetailsResponse) GetApplicationToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApplicationToken
}

// GetApplicationTokenOk returns a tuple with the ApplicationToken field value
// and a boolean to check if the value has been set.
func (o *ErrorDetailsResponse) GetApplicationTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApplicationToken, true
}

// SetApplicationToken sets field value
func (o *ErrorDetailsResponse) SetApplicationToken(v string) {
	o.ApplicationToken = v
}

// GetMessage returns the Message field value
func (o *ErrorDetailsResponse) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ErrorDetailsResponse) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *ErrorDetailsResponse) SetMessage(v string) {
	o.Message = v
}

// GetToken returns the Token field value
func (o *ErrorDetailsResponse) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *ErrorDetailsResponse) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *ErrorDetailsResponse) SetToken(v string) {
	o.Token = v
}

// GetType returns the Type field value
func (o *ErrorDetailsResponse) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ErrorDetailsResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ErrorDetailsResponse) SetType(v string) {
	o.Type = v
}

func (o ErrorDetailsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorDetailsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["application_token"] = o.ApplicationToken
	toSerialize["message"] = o.Message
	toSerialize["token"] = o.Token
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableErrorDetailsResponse struct {
	value *ErrorDetailsResponse
	isSet bool
}

func (v NullableErrorDetailsResponse) Get() *ErrorDetailsResponse {
	return v.value
}

func (v *NullableErrorDetailsResponse) Set(val *ErrorDetailsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorDetailsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorDetailsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorDetailsResponse(val *ErrorDetailsResponse) *NullableErrorDetailsResponse {
	return &NullableErrorDetailsResponse{value: val, isSet: true}
}

func (v NullableErrorDetailsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorDetailsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


