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

// checks if the Application type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Application{}

// Application Contains client application information.
type Application struct {
	// Access code of the client application.
	AccessCode *string `json:"access_code,omitempty"`
	// URL of the client application assets.
	AssetsUrl *string `json:"assets_url,omitempty"`
	// Base URL of the client API.
	ClientApiBaseUrl *string `json:"client_api_base_url,omitempty"`
	// Client application's environment.
	Environment *string `json:"environment,omitempty"`
	// Name of the program on the Marqeta platform.
	Program *string `json:"program,omitempty"`
	// Short code of the program on the Marqeta platform.
	ProgramShortCode *string `json:"program_short_code,omitempty"`
	// Unique identifier of the `application` object.
	Token *string `json:"token,omitempty"`
}

// NewApplication instantiates a new Application object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplication() *Application {
	this := Application{}
	return &this
}

// NewApplicationWithDefaults instantiates a new Application object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationWithDefaults() *Application {
	this := Application{}
	return &this
}

// GetAccessCode returns the AccessCode field value if set, zero value otherwise.
func (o *Application) GetAccessCode() string {
	if o == nil || IsNil(o.AccessCode) {
		var ret string
		return ret
	}
	return *o.AccessCode
}

// GetAccessCodeOk returns a tuple with the AccessCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetAccessCodeOk() (*string, bool) {
	if o == nil || IsNil(o.AccessCode) {
		return nil, false
	}
	return o.AccessCode, true
}

// HasAccessCode returns a boolean if a field has been set.
func (o *Application) HasAccessCode() bool {
	if o != nil && !IsNil(o.AccessCode) {
		return true
	}

	return false
}

// SetAccessCode gets a reference to the given string and assigns it to the AccessCode field.
func (o *Application) SetAccessCode(v string) {
	o.AccessCode = &v
}

// GetAssetsUrl returns the AssetsUrl field value if set, zero value otherwise.
func (o *Application) GetAssetsUrl() string {
	if o == nil || IsNil(o.AssetsUrl) {
		var ret string
		return ret
	}
	return *o.AssetsUrl
}

// GetAssetsUrlOk returns a tuple with the AssetsUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetAssetsUrlOk() (*string, bool) {
	if o == nil || IsNil(o.AssetsUrl) {
		return nil, false
	}
	return o.AssetsUrl, true
}

// HasAssetsUrl returns a boolean if a field has been set.
func (o *Application) HasAssetsUrl() bool {
	if o != nil && !IsNil(o.AssetsUrl) {
		return true
	}

	return false
}

// SetAssetsUrl gets a reference to the given string and assigns it to the AssetsUrl field.
func (o *Application) SetAssetsUrl(v string) {
	o.AssetsUrl = &v
}

// GetClientApiBaseUrl returns the ClientApiBaseUrl field value if set, zero value otherwise.
func (o *Application) GetClientApiBaseUrl() string {
	if o == nil || IsNil(o.ClientApiBaseUrl) {
		var ret string
		return ret
	}
	return *o.ClientApiBaseUrl
}

// GetClientApiBaseUrlOk returns a tuple with the ClientApiBaseUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetClientApiBaseUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ClientApiBaseUrl) {
		return nil, false
	}
	return o.ClientApiBaseUrl, true
}

// HasClientApiBaseUrl returns a boolean if a field has been set.
func (o *Application) HasClientApiBaseUrl() bool {
	if o != nil && !IsNil(o.ClientApiBaseUrl) {
		return true
	}

	return false
}

// SetClientApiBaseUrl gets a reference to the given string and assigns it to the ClientApiBaseUrl field.
func (o *Application) SetClientApiBaseUrl(v string) {
	o.ClientApiBaseUrl = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *Application) GetEnvironment() string {
	if o == nil || IsNil(o.Environment) {
		var ret string
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetEnvironmentOk() (*string, bool) {
	if o == nil || IsNil(o.Environment) {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *Application) HasEnvironment() bool {
	if o != nil && !IsNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given string and assigns it to the Environment field.
func (o *Application) SetEnvironment(v string) {
	o.Environment = &v
}

// GetProgram returns the Program field value if set, zero value otherwise.
func (o *Application) GetProgram() string {
	if o == nil || IsNil(o.Program) {
		var ret string
		return ret
	}
	return *o.Program
}

// GetProgramOk returns a tuple with the Program field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetProgramOk() (*string, bool) {
	if o == nil || IsNil(o.Program) {
		return nil, false
	}
	return o.Program, true
}

// HasProgram returns a boolean if a field has been set.
func (o *Application) HasProgram() bool {
	if o != nil && !IsNil(o.Program) {
		return true
	}

	return false
}

// SetProgram gets a reference to the given string and assigns it to the Program field.
func (o *Application) SetProgram(v string) {
	o.Program = &v
}

// GetProgramShortCode returns the ProgramShortCode field value if set, zero value otherwise.
func (o *Application) GetProgramShortCode() string {
	if o == nil || IsNil(o.ProgramShortCode) {
		var ret string
		return ret
	}
	return *o.ProgramShortCode
}

// GetProgramShortCodeOk returns a tuple with the ProgramShortCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetProgramShortCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ProgramShortCode) {
		return nil, false
	}
	return o.ProgramShortCode, true
}

// HasProgramShortCode returns a boolean if a field has been set.
func (o *Application) HasProgramShortCode() bool {
	if o != nil && !IsNil(o.ProgramShortCode) {
		return true
	}

	return false
}

// SetProgramShortCode gets a reference to the given string and assigns it to the ProgramShortCode field.
func (o *Application) SetProgramShortCode(v string) {
	o.ProgramShortCode = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *Application) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *Application) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *Application) SetToken(v string) {
	o.Token = &v
}

func (o Application) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Application) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccessCode) {
		toSerialize["access_code"] = o.AccessCode
	}
	if !IsNil(o.AssetsUrl) {
		toSerialize["assets_url"] = o.AssetsUrl
	}
	if !IsNil(o.ClientApiBaseUrl) {
		toSerialize["client_api_base_url"] = o.ClientApiBaseUrl
	}
	if !IsNil(o.Environment) {
		toSerialize["environment"] = o.Environment
	}
	if !IsNil(o.Program) {
		toSerialize["program"] = o.Program
	}
	if !IsNil(o.ProgramShortCode) {
		toSerialize["program_short_code"] = o.ProgramShortCode
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	return toSerialize, nil
}

type NullableApplication struct {
	value *Application
	isSet bool
}

func (v NullableApplication) Get() *Application {
	return v.value
}

func (v *NullableApplication) Set(val *Application) {
	v.value = val
	v.isSet = true
}

func (v NullableApplication) IsSet() bool {
	return v.isSet
}

func (v *NullableApplication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplication(val *Application) *NullableApplication {
	return &NullableApplication{value: val, isSet: true}
}

func (v NullableApplication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

