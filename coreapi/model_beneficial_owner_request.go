/*
Core API

Marqeta's Core API endpoints, conveniently annotated to enable code generation (including SDKs), test cases, and documentation. Currently in beta.

API version: 3.0.11
Contact: support@marqeta.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package coreapi

import (
	"encoding/json"
	"time"
)

// checks if the BeneficialOwnerRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BeneficialOwnerRequest{}

// BeneficialOwnerRequest Contains information about the beneficial owner of the business, if applicable.  This object is required for KYC verification in the United States if the business is private and has a beneficial owner. This object is not required for publicly held businesses.  Do not include information about the proprietor or business officer in a `beneficial_owner` object. If the proprietor or officer of the business is also a beneficial owner, you must indicate that in the `proprietor_is_beneficial_owner` field in the body field details of the business.
type BeneficialOwnerRequest struct {
	// Date of birth of the beneficial owner.
	Dob *time.Time `json:"dob,omitempty"`
	// First name of the beneficial owner.
	FirstName *string `json:"first_name,omitempty"`
	Home *AddressRequestModel `json:"home,omitempty"`
	// Last name of the beneficial owner.
	LastName *string `json:"last_name,omitempty"`
	// Middle name of the beneficial owner.
	MiddleName *string `json:"middle_name,omitempty"`
	// Ten-digit phone number of the beneficial owner.
	Phone *string `json:"phone,omitempty"`
	// Nine-digit Social Security Number (SSN) of the beneficial owner.
	Ssn *string `json:"ssn,omitempty"`
	// Title of the beneficial owner.
	Title *string `json:"title,omitempty"`
}

// NewBeneficialOwnerRequest instantiates a new BeneficialOwnerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBeneficialOwnerRequest() *BeneficialOwnerRequest {
	this := BeneficialOwnerRequest{}
	return &this
}

// NewBeneficialOwnerRequestWithDefaults instantiates a new BeneficialOwnerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBeneficialOwnerRequestWithDefaults() *BeneficialOwnerRequest {
	this := BeneficialOwnerRequest{}
	return &this
}

// GetDob returns the Dob field value if set, zero value otherwise.
func (o *BeneficialOwnerRequest) GetDob() time.Time {
	if o == nil || IsNil(o.Dob) {
		var ret time.Time
		return ret
	}
	return *o.Dob
}

// GetDobOk returns a tuple with the Dob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BeneficialOwnerRequest) GetDobOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Dob) {
		return nil, false
	}
	return o.Dob, true
}

// HasDob returns a boolean if a field has been set.
func (o *BeneficialOwnerRequest) HasDob() bool {
	if o != nil && !IsNil(o.Dob) {
		return true
	}

	return false
}

// SetDob gets a reference to the given time.Time and assigns it to the Dob field.
func (o *BeneficialOwnerRequest) SetDob(v time.Time) {
	o.Dob = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *BeneficialOwnerRequest) GetFirstName() string {
	if o == nil || IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BeneficialOwnerRequest) GetFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *BeneficialOwnerRequest) HasFirstName() bool {
	if o != nil && !IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *BeneficialOwnerRequest) SetFirstName(v string) {
	o.FirstName = &v
}

// GetHome returns the Home field value if set, zero value otherwise.
func (o *BeneficialOwnerRequest) GetHome() AddressRequestModel {
	if o == nil || IsNil(o.Home) {
		var ret AddressRequestModel
		return ret
	}
	return *o.Home
}

// GetHomeOk returns a tuple with the Home field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BeneficialOwnerRequest) GetHomeOk() (*AddressRequestModel, bool) {
	if o == nil || IsNil(o.Home) {
		return nil, false
	}
	return o.Home, true
}

// HasHome returns a boolean if a field has been set.
func (o *BeneficialOwnerRequest) HasHome() bool {
	if o != nil && !IsNil(o.Home) {
		return true
	}

	return false
}

// SetHome gets a reference to the given AddressRequestModel and assigns it to the Home field.
func (o *BeneficialOwnerRequest) SetHome(v AddressRequestModel) {
	o.Home = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *BeneficialOwnerRequest) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BeneficialOwnerRequest) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *BeneficialOwnerRequest) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *BeneficialOwnerRequest) SetLastName(v string) {
	o.LastName = &v
}

// GetMiddleName returns the MiddleName field value if set, zero value otherwise.
func (o *BeneficialOwnerRequest) GetMiddleName() string {
	if o == nil || IsNil(o.MiddleName) {
		var ret string
		return ret
	}
	return *o.MiddleName
}

// GetMiddleNameOk returns a tuple with the MiddleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BeneficialOwnerRequest) GetMiddleNameOk() (*string, bool) {
	if o == nil || IsNil(o.MiddleName) {
		return nil, false
	}
	return o.MiddleName, true
}

// HasMiddleName returns a boolean if a field has been set.
func (o *BeneficialOwnerRequest) HasMiddleName() bool {
	if o != nil && !IsNil(o.MiddleName) {
		return true
	}

	return false
}

// SetMiddleName gets a reference to the given string and assigns it to the MiddleName field.
func (o *BeneficialOwnerRequest) SetMiddleName(v string) {
	o.MiddleName = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *BeneficialOwnerRequest) GetPhone() string {
	if o == nil || IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BeneficialOwnerRequest) GetPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *BeneficialOwnerRequest) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *BeneficialOwnerRequest) SetPhone(v string) {
	o.Phone = &v
}

// GetSsn returns the Ssn field value if set, zero value otherwise.
func (o *BeneficialOwnerRequest) GetSsn() string {
	if o == nil || IsNil(o.Ssn) {
		var ret string
		return ret
	}
	return *o.Ssn
}

// GetSsnOk returns a tuple with the Ssn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BeneficialOwnerRequest) GetSsnOk() (*string, bool) {
	if o == nil || IsNil(o.Ssn) {
		return nil, false
	}
	return o.Ssn, true
}

// HasSsn returns a boolean if a field has been set.
func (o *BeneficialOwnerRequest) HasSsn() bool {
	if o != nil && !IsNil(o.Ssn) {
		return true
	}

	return false
}

// SetSsn gets a reference to the given string and assigns it to the Ssn field.
func (o *BeneficialOwnerRequest) SetSsn(v string) {
	o.Ssn = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *BeneficialOwnerRequest) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BeneficialOwnerRequest) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *BeneficialOwnerRequest) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *BeneficialOwnerRequest) SetTitle(v string) {
	o.Title = &v
}

func (o BeneficialOwnerRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BeneficialOwnerRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Dob) {
		toSerialize["dob"] = o.Dob
	}
	if !IsNil(o.FirstName) {
		toSerialize["first_name"] = o.FirstName
	}
	if !IsNil(o.Home) {
		toSerialize["home"] = o.Home
	}
	if !IsNil(o.LastName) {
		toSerialize["last_name"] = o.LastName
	}
	if !IsNil(o.MiddleName) {
		toSerialize["middle_name"] = o.MiddleName
	}
	if !IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	if !IsNil(o.Ssn) {
		toSerialize["ssn"] = o.Ssn
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	return toSerialize, nil
}

type NullableBeneficialOwnerRequest struct {
	value *BeneficialOwnerRequest
	isSet bool
}

func (v NullableBeneficialOwnerRequest) Get() *BeneficialOwnerRequest {
	return v.value
}

func (v *NullableBeneficialOwnerRequest) Set(val *BeneficialOwnerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBeneficialOwnerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBeneficialOwnerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBeneficialOwnerRequest(val *BeneficialOwnerRequest) *NullableBeneficialOwnerRequest {
	return &NullableBeneficialOwnerRequest{value: val, isSet: true}
}

func (v NullableBeneficialOwnerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBeneficialOwnerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


