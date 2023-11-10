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

// checks if the Fee type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Fee{}

// Fee Contains details about the fee.
type Fee struct {
	// Indicates whether the fee is active.
	Active bool `json:"active"`
	// Amount of the fee.
	Amount float32 `json:"amount"`
	// Date and time when the `fees` object was created, in UTC.
	CreatedTime time.Time `json:"created_time"`
	// The three-digit ISO 4217 currency code.
	CurrencyCode string `json:"currency_code"`
	// Date and time when the GPA order was last modified, in UTC.
	LastModifiedTime time.Time `json:"last_modified_time"`
	// Name of the fee.
	Name string `json:"name"`
	RealTimeAssessment *RealTimeFeeAssessment `json:"real_time_assessment,omitempty"`
	// Descriptive metadata about the fee.
	Tags *string `json:"tags,omitempty"`
	// Unique identifier of the `fees` object.
	Token string `json:"token"`
}

// NewFee instantiates a new Fee object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFee(active bool, amount float32, createdTime time.Time, currencyCode string, lastModifiedTime time.Time, name string, token string) *Fee {
	this := Fee{}
	this.Active = active
	this.Amount = amount
	this.CreatedTime = createdTime
	this.CurrencyCode = currencyCode
	this.LastModifiedTime = lastModifiedTime
	this.Name = name
	this.Token = token
	return &this
}

// NewFeeWithDefaults instantiates a new Fee object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeeWithDefaults() *Fee {
	this := Fee{}
	var active bool = false
	this.Active = active
	return &this
}

// GetActive returns the Active field value
func (o *Fee) GetActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Active
}

// GetActiveOk returns a tuple with the Active field value
// and a boolean to check if the value has been set.
func (o *Fee) GetActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Active, true
}

// SetActive sets field value
func (o *Fee) SetActive(v bool) {
	o.Active = v
}

// GetAmount returns the Amount field value
func (o *Fee) GetAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *Fee) GetAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *Fee) SetAmount(v float32) {
	o.Amount = v
}

// GetCreatedTime returns the CreatedTime field value
func (o *Fee) GetCreatedTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value
// and a boolean to check if the value has been set.
func (o *Fee) GetCreatedTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedTime, true
}

// SetCreatedTime sets field value
func (o *Fee) SetCreatedTime(v time.Time) {
	o.CreatedTime = v
}

// GetCurrencyCode returns the CurrencyCode field value
func (o *Fee) GetCurrencyCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value
// and a boolean to check if the value has been set.
func (o *Fee) GetCurrencyCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrencyCode, true
}

// SetCurrencyCode sets field value
func (o *Fee) SetCurrencyCode(v string) {
	o.CurrencyCode = v
}

// GetLastModifiedTime returns the LastModifiedTime field value
func (o *Fee) GetLastModifiedTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastModifiedTime
}

// GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field value
// and a boolean to check if the value has been set.
func (o *Fee) GetLastModifiedTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastModifiedTime, true
}

// SetLastModifiedTime sets field value
func (o *Fee) SetLastModifiedTime(v time.Time) {
	o.LastModifiedTime = v
}

// GetName returns the Name field value
func (o *Fee) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Fee) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Fee) SetName(v string) {
	o.Name = v
}

// GetRealTimeAssessment returns the RealTimeAssessment field value if set, zero value otherwise.
func (o *Fee) GetRealTimeAssessment() RealTimeFeeAssessment {
	if o == nil || IsNil(o.RealTimeAssessment) {
		var ret RealTimeFeeAssessment
		return ret
	}
	return *o.RealTimeAssessment
}

// GetRealTimeAssessmentOk returns a tuple with the RealTimeAssessment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Fee) GetRealTimeAssessmentOk() (*RealTimeFeeAssessment, bool) {
	if o == nil || IsNil(o.RealTimeAssessment) {
		return nil, false
	}
	return o.RealTimeAssessment, true
}

// HasRealTimeAssessment returns a boolean if a field has been set.
func (o *Fee) HasRealTimeAssessment() bool {
	if o != nil && !IsNil(o.RealTimeAssessment) {
		return true
	}

	return false
}

// SetRealTimeAssessment gets a reference to the given RealTimeFeeAssessment and assigns it to the RealTimeAssessment field.
func (o *Fee) SetRealTimeAssessment(v RealTimeFeeAssessment) {
	o.RealTimeAssessment = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Fee) GetTags() string {
	if o == nil || IsNil(o.Tags) {
		var ret string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Fee) GetTagsOk() (*string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Fee) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given string and assigns it to the Tags field.
func (o *Fee) SetTags(v string) {
	o.Tags = &v
}

// GetToken returns the Token field value
func (o *Fee) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *Fee) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *Fee) SetToken(v string) {
	o.Token = v
}

func (o Fee) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Fee) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["active"] = o.Active
	toSerialize["amount"] = o.Amount
	toSerialize["created_time"] = o.CreatedTime
	toSerialize["currency_code"] = o.CurrencyCode
	toSerialize["last_modified_time"] = o.LastModifiedTime
	toSerialize["name"] = o.Name
	if !IsNil(o.RealTimeAssessment) {
		toSerialize["real_time_assessment"] = o.RealTimeAssessment
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	toSerialize["token"] = o.Token
	return toSerialize, nil
}

type NullableFee struct {
	value *Fee
	isSet bool
}

func (v NullableFee) Get() *Fee {
	return v.value
}

func (v *NullableFee) Set(val *Fee) {
	v.value = val
	v.isSet = true
}

func (v NullableFee) IsSet() bool {
	return v.isSet
}

func (v *NullableFee) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFee(val *Fee) *NullableFee {
	return &NullableFee{value: val, isSet: true}
}

func (v NullableFee) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFee) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


