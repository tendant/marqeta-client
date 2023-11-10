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

// checks if the CommandoModeResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommandoModeResponse{}

// CommandoModeResponse struct for CommandoModeResponse
type CommandoModeResponse struct {
	CommandoModeEnables *CommandoModeEnables `json:"commando_mode_enables,omitempty"`
	// Date and time when the resource was created, in UTC.
	CreatedTime time.Time `json:"created_time"`
	CurrentState *CommandoModeNestedTransition `json:"current_state,omitempty"`
	// Date and time when the resource was last updated, in UTC.
	LastModifiedTime time.Time `json:"last_modified_time"`
	// Unique identifier of the associated program gateway funding source.
	ProgramGatewayFundingSourceToken *string `json:"program_gateway_funding_source_token,omitempty"`
	RealTimeStandinCriteria *RealTimeStandinCriteria `json:"real_time_standin_criteria,omitempty"`
	// Unique identifier of the Commando Mode control set.
	Token *string `json:"token,omitempty"`
}

// NewCommandoModeResponse instantiates a new CommandoModeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommandoModeResponse(createdTime time.Time, lastModifiedTime time.Time) *CommandoModeResponse {
	this := CommandoModeResponse{}
	this.CreatedTime = createdTime
	this.LastModifiedTime = lastModifiedTime
	return &this
}

// NewCommandoModeResponseWithDefaults instantiates a new CommandoModeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommandoModeResponseWithDefaults() *CommandoModeResponse {
	this := CommandoModeResponse{}
	return &this
}

// GetCommandoModeEnables returns the CommandoModeEnables field value if set, zero value otherwise.
func (o *CommandoModeResponse) GetCommandoModeEnables() CommandoModeEnables {
	if o == nil || IsNil(o.CommandoModeEnables) {
		var ret CommandoModeEnables
		return ret
	}
	return *o.CommandoModeEnables
}

// GetCommandoModeEnablesOk returns a tuple with the CommandoModeEnables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommandoModeResponse) GetCommandoModeEnablesOk() (*CommandoModeEnables, bool) {
	if o == nil || IsNil(o.CommandoModeEnables) {
		return nil, false
	}
	return o.CommandoModeEnables, true
}

// HasCommandoModeEnables returns a boolean if a field has been set.
func (o *CommandoModeResponse) HasCommandoModeEnables() bool {
	if o != nil && !IsNil(o.CommandoModeEnables) {
		return true
	}

	return false
}

// SetCommandoModeEnables gets a reference to the given CommandoModeEnables and assigns it to the CommandoModeEnables field.
func (o *CommandoModeResponse) SetCommandoModeEnables(v CommandoModeEnables) {
	o.CommandoModeEnables = &v
}

// GetCreatedTime returns the CreatedTime field value
func (o *CommandoModeResponse) GetCreatedTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value
// and a boolean to check if the value has been set.
func (o *CommandoModeResponse) GetCreatedTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedTime, true
}

// SetCreatedTime sets field value
func (o *CommandoModeResponse) SetCreatedTime(v time.Time) {
	o.CreatedTime = v
}

// GetCurrentState returns the CurrentState field value if set, zero value otherwise.
func (o *CommandoModeResponse) GetCurrentState() CommandoModeNestedTransition {
	if o == nil || IsNil(o.CurrentState) {
		var ret CommandoModeNestedTransition
		return ret
	}
	return *o.CurrentState
}

// GetCurrentStateOk returns a tuple with the CurrentState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommandoModeResponse) GetCurrentStateOk() (*CommandoModeNestedTransition, bool) {
	if o == nil || IsNil(o.CurrentState) {
		return nil, false
	}
	return o.CurrentState, true
}

// HasCurrentState returns a boolean if a field has been set.
func (o *CommandoModeResponse) HasCurrentState() bool {
	if o != nil && !IsNil(o.CurrentState) {
		return true
	}

	return false
}

// SetCurrentState gets a reference to the given CommandoModeNestedTransition and assigns it to the CurrentState field.
func (o *CommandoModeResponse) SetCurrentState(v CommandoModeNestedTransition) {
	o.CurrentState = &v
}

// GetLastModifiedTime returns the LastModifiedTime field value
func (o *CommandoModeResponse) GetLastModifiedTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastModifiedTime
}

// GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field value
// and a boolean to check if the value has been set.
func (o *CommandoModeResponse) GetLastModifiedTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastModifiedTime, true
}

// SetLastModifiedTime sets field value
func (o *CommandoModeResponse) SetLastModifiedTime(v time.Time) {
	o.LastModifiedTime = v
}

// GetProgramGatewayFundingSourceToken returns the ProgramGatewayFundingSourceToken field value if set, zero value otherwise.
func (o *CommandoModeResponse) GetProgramGatewayFundingSourceToken() string {
	if o == nil || IsNil(o.ProgramGatewayFundingSourceToken) {
		var ret string
		return ret
	}
	return *o.ProgramGatewayFundingSourceToken
}

// GetProgramGatewayFundingSourceTokenOk returns a tuple with the ProgramGatewayFundingSourceToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommandoModeResponse) GetProgramGatewayFundingSourceTokenOk() (*string, bool) {
	if o == nil || IsNil(o.ProgramGatewayFundingSourceToken) {
		return nil, false
	}
	return o.ProgramGatewayFundingSourceToken, true
}

// HasProgramGatewayFundingSourceToken returns a boolean if a field has been set.
func (o *CommandoModeResponse) HasProgramGatewayFundingSourceToken() bool {
	if o != nil && !IsNil(o.ProgramGatewayFundingSourceToken) {
		return true
	}

	return false
}

// SetProgramGatewayFundingSourceToken gets a reference to the given string and assigns it to the ProgramGatewayFundingSourceToken field.
func (o *CommandoModeResponse) SetProgramGatewayFundingSourceToken(v string) {
	o.ProgramGatewayFundingSourceToken = &v
}

// GetRealTimeStandinCriteria returns the RealTimeStandinCriteria field value if set, zero value otherwise.
func (o *CommandoModeResponse) GetRealTimeStandinCriteria() RealTimeStandinCriteria {
	if o == nil || IsNil(o.RealTimeStandinCriteria) {
		var ret RealTimeStandinCriteria
		return ret
	}
	return *o.RealTimeStandinCriteria
}

// GetRealTimeStandinCriteriaOk returns a tuple with the RealTimeStandinCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommandoModeResponse) GetRealTimeStandinCriteriaOk() (*RealTimeStandinCriteria, bool) {
	if o == nil || IsNil(o.RealTimeStandinCriteria) {
		return nil, false
	}
	return o.RealTimeStandinCriteria, true
}

// HasRealTimeStandinCriteria returns a boolean if a field has been set.
func (o *CommandoModeResponse) HasRealTimeStandinCriteria() bool {
	if o != nil && !IsNil(o.RealTimeStandinCriteria) {
		return true
	}

	return false
}

// SetRealTimeStandinCriteria gets a reference to the given RealTimeStandinCriteria and assigns it to the RealTimeStandinCriteria field.
func (o *CommandoModeResponse) SetRealTimeStandinCriteria(v RealTimeStandinCriteria) {
	o.RealTimeStandinCriteria = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *CommandoModeResponse) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommandoModeResponse) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *CommandoModeResponse) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *CommandoModeResponse) SetToken(v string) {
	o.Token = &v
}

func (o CommandoModeResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommandoModeResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CommandoModeEnables) {
		toSerialize["commando_mode_enables"] = o.CommandoModeEnables
	}
	toSerialize["created_time"] = o.CreatedTime
	if !IsNil(o.CurrentState) {
		toSerialize["current_state"] = o.CurrentState
	}
	toSerialize["last_modified_time"] = o.LastModifiedTime
	if !IsNil(o.ProgramGatewayFundingSourceToken) {
		toSerialize["program_gateway_funding_source_token"] = o.ProgramGatewayFundingSourceToken
	}
	if !IsNil(o.RealTimeStandinCriteria) {
		toSerialize["real_time_standin_criteria"] = o.RealTimeStandinCriteria
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	return toSerialize, nil
}

type NullableCommandoModeResponse struct {
	value *CommandoModeResponse
	isSet bool
}

func (v NullableCommandoModeResponse) Get() *CommandoModeResponse {
	return v.value
}

func (v *NullableCommandoModeResponse) Set(val *CommandoModeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCommandoModeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCommandoModeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommandoModeResponse(val *CommandoModeResponse) *NullableCommandoModeResponse {
	return &NullableCommandoModeResponse{value: val, isSet: true}
}

func (v NullableCommandoModeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommandoModeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


