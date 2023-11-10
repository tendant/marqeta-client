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
	"time"
)

// checks if the ApplicationTransitionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationTransitionResponse{}

// ApplicationTransitionResponse Contains information on the transition of an application's state.
type ApplicationTransitionResponse struct {
	// Unique identifier of the application whose state you transitioned.
	ApplicationToken *string `json:"application_token,omitempty"`
	// Date and time when the application changed states on the Marqeta platform, in UTC.
	CreatedTime *time.Time `json:"created_time,omitempty"`
	OriginalStatus *ApplicationResourceState `json:"original_status,omitempty"`
	Status *ApplicationResourceState `json:"status,omitempty"`
	// Unique identifier of the transition of an application's state.
	Token *string `json:"token,omitempty"`
}

// NewApplicationTransitionResponse instantiates a new ApplicationTransitionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationTransitionResponse() *ApplicationTransitionResponse {
	this := ApplicationTransitionResponse{}
	return &this
}

// NewApplicationTransitionResponseWithDefaults instantiates a new ApplicationTransitionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationTransitionResponseWithDefaults() *ApplicationTransitionResponse {
	this := ApplicationTransitionResponse{}
	return &this
}

// GetApplicationToken returns the ApplicationToken field value if set, zero value otherwise.
func (o *ApplicationTransitionResponse) GetApplicationToken() string {
	if o == nil || IsNil(o.ApplicationToken) {
		var ret string
		return ret
	}
	return *o.ApplicationToken
}

// GetApplicationTokenOk returns a tuple with the ApplicationToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationTransitionResponse) GetApplicationTokenOk() (*string, bool) {
	if o == nil || IsNil(o.ApplicationToken) {
		return nil, false
	}
	return o.ApplicationToken, true
}

// HasApplicationToken returns a boolean if a field has been set.
func (o *ApplicationTransitionResponse) HasApplicationToken() bool {
	if o != nil && !IsNil(o.ApplicationToken) {
		return true
	}

	return false
}

// SetApplicationToken gets a reference to the given string and assigns it to the ApplicationToken field.
func (o *ApplicationTransitionResponse) SetApplicationToken(v string) {
	o.ApplicationToken = &v
}

// GetCreatedTime returns the CreatedTime field value if set, zero value otherwise.
func (o *ApplicationTransitionResponse) GetCreatedTime() time.Time {
	if o == nil || IsNil(o.CreatedTime) {
		var ret time.Time
		return ret
	}
	return *o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationTransitionResponse) GetCreatedTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedTime) {
		return nil, false
	}
	return o.CreatedTime, true
}

// HasCreatedTime returns a boolean if a field has been set.
func (o *ApplicationTransitionResponse) HasCreatedTime() bool {
	if o != nil && !IsNil(o.CreatedTime) {
		return true
	}

	return false
}

// SetCreatedTime gets a reference to the given time.Time and assigns it to the CreatedTime field.
func (o *ApplicationTransitionResponse) SetCreatedTime(v time.Time) {
	o.CreatedTime = &v
}

// GetOriginalStatus returns the OriginalStatus field value if set, zero value otherwise.
func (o *ApplicationTransitionResponse) GetOriginalStatus() ApplicationResourceState {
	if o == nil || IsNil(o.OriginalStatus) {
		var ret ApplicationResourceState
		return ret
	}
	return *o.OriginalStatus
}

// GetOriginalStatusOk returns a tuple with the OriginalStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationTransitionResponse) GetOriginalStatusOk() (*ApplicationResourceState, bool) {
	if o == nil || IsNil(o.OriginalStatus) {
		return nil, false
	}
	return o.OriginalStatus, true
}

// HasOriginalStatus returns a boolean if a field has been set.
func (o *ApplicationTransitionResponse) HasOriginalStatus() bool {
	if o != nil && !IsNil(o.OriginalStatus) {
		return true
	}

	return false
}

// SetOriginalStatus gets a reference to the given ApplicationResourceState and assigns it to the OriginalStatus field.
func (o *ApplicationTransitionResponse) SetOriginalStatus(v ApplicationResourceState) {
	o.OriginalStatus = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ApplicationTransitionResponse) GetStatus() ApplicationResourceState {
	if o == nil || IsNil(o.Status) {
		var ret ApplicationResourceState
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationTransitionResponse) GetStatusOk() (*ApplicationResourceState, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ApplicationTransitionResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ApplicationResourceState and assigns it to the Status field.
func (o *ApplicationTransitionResponse) SetStatus(v ApplicationResourceState) {
	o.Status = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *ApplicationTransitionResponse) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationTransitionResponse) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *ApplicationTransitionResponse) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *ApplicationTransitionResponse) SetToken(v string) {
	o.Token = &v
}

func (o ApplicationTransitionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationTransitionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApplicationToken) {
		toSerialize["application_token"] = o.ApplicationToken
	}
	if !IsNil(o.CreatedTime) {
		toSerialize["created_time"] = o.CreatedTime
	}
	if !IsNil(o.OriginalStatus) {
		toSerialize["original_status"] = o.OriginalStatus
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	return toSerialize, nil
}

type NullableApplicationTransitionResponse struct {
	value *ApplicationTransitionResponse
	isSet bool
}

func (v NullableApplicationTransitionResponse) Get() *ApplicationTransitionResponse {
	return v.value
}

func (v *NullableApplicationTransitionResponse) Set(val *ApplicationTransitionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationTransitionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationTransitionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationTransitionResponse(val *ApplicationTransitionResponse) *NullableApplicationTransitionResponse {
	return &NullableApplicationTransitionResponse{value: val, isSet: true}
}

func (v NullableApplicationTransitionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationTransitionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

