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

// checks if the PrimaryContactInfoModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrimaryContactInfoModel{}

// PrimaryContactInfoModel Describes the business' primary contact person.
type PrimaryContactInfoModel struct {
	// Business department of the primary contact.
	Department *string `json:"department,omitempty"`
	// Email address of the primary contact.
	Email *string `json:"email,omitempty"`
	// Phone extension of the primary contact.
	Extension *string `json:"extension,omitempty"`
	// Fax number of the primary contact.
	Fax *string `json:"fax,omitempty"`
	// Full name of the primary contact.
	FullName *string `json:"full_name,omitempty"`
	// Mobile phone number of the primary contact.
	Mobile *string `json:"mobile,omitempty"`
	// Phone number of the primary contact.
	Phone *string `json:"phone,omitempty"`
	// Title of the primary contact.
	Title *string `json:"title,omitempty"`
}

// NewPrimaryContactInfoModel instantiates a new PrimaryContactInfoModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrimaryContactInfoModel() *PrimaryContactInfoModel {
	this := PrimaryContactInfoModel{}
	return &this
}

// NewPrimaryContactInfoModelWithDefaults instantiates a new PrimaryContactInfoModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrimaryContactInfoModelWithDefaults() *PrimaryContactInfoModel {
	this := PrimaryContactInfoModel{}
	return &this
}

// GetDepartment returns the Department field value if set, zero value otherwise.
func (o *PrimaryContactInfoModel) GetDepartment() string {
	if o == nil || IsNil(o.Department) {
		var ret string
		return ret
	}
	return *o.Department
}

// GetDepartmentOk returns a tuple with the Department field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrimaryContactInfoModel) GetDepartmentOk() (*string, bool) {
	if o == nil || IsNil(o.Department) {
		return nil, false
	}
	return o.Department, true
}

// HasDepartment returns a boolean if a field has been set.
func (o *PrimaryContactInfoModel) HasDepartment() bool {
	if o != nil && !IsNil(o.Department) {
		return true
	}

	return false
}

// SetDepartment gets a reference to the given string and assigns it to the Department field.
func (o *PrimaryContactInfoModel) SetDepartment(v string) {
	o.Department = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *PrimaryContactInfoModel) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrimaryContactInfoModel) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *PrimaryContactInfoModel) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *PrimaryContactInfoModel) SetEmail(v string) {
	o.Email = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *PrimaryContactInfoModel) GetExtension() string {
	if o == nil || IsNil(o.Extension) {
		var ret string
		return ret
	}
	return *o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrimaryContactInfoModel) GetExtensionOk() (*string, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *PrimaryContactInfoModel) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given string and assigns it to the Extension field.
func (o *PrimaryContactInfoModel) SetExtension(v string) {
	o.Extension = &v
}

// GetFax returns the Fax field value if set, zero value otherwise.
func (o *PrimaryContactInfoModel) GetFax() string {
	if o == nil || IsNil(o.Fax) {
		var ret string
		return ret
	}
	return *o.Fax
}

// GetFaxOk returns a tuple with the Fax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrimaryContactInfoModel) GetFaxOk() (*string, bool) {
	if o == nil || IsNil(o.Fax) {
		return nil, false
	}
	return o.Fax, true
}

// HasFax returns a boolean if a field has been set.
func (o *PrimaryContactInfoModel) HasFax() bool {
	if o != nil && !IsNil(o.Fax) {
		return true
	}

	return false
}

// SetFax gets a reference to the given string and assigns it to the Fax field.
func (o *PrimaryContactInfoModel) SetFax(v string) {
	o.Fax = &v
}

// GetFullName returns the FullName field value if set, zero value otherwise.
func (o *PrimaryContactInfoModel) GetFullName() string {
	if o == nil || IsNil(o.FullName) {
		var ret string
		return ret
	}
	return *o.FullName
}

// GetFullNameOk returns a tuple with the FullName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrimaryContactInfoModel) GetFullNameOk() (*string, bool) {
	if o == nil || IsNil(o.FullName) {
		return nil, false
	}
	return o.FullName, true
}

// HasFullName returns a boolean if a field has been set.
func (o *PrimaryContactInfoModel) HasFullName() bool {
	if o != nil && !IsNil(o.FullName) {
		return true
	}

	return false
}

// SetFullName gets a reference to the given string and assigns it to the FullName field.
func (o *PrimaryContactInfoModel) SetFullName(v string) {
	o.FullName = &v
}

// GetMobile returns the Mobile field value if set, zero value otherwise.
func (o *PrimaryContactInfoModel) GetMobile() string {
	if o == nil || IsNil(o.Mobile) {
		var ret string
		return ret
	}
	return *o.Mobile
}

// GetMobileOk returns a tuple with the Mobile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrimaryContactInfoModel) GetMobileOk() (*string, bool) {
	if o == nil || IsNil(o.Mobile) {
		return nil, false
	}
	return o.Mobile, true
}

// HasMobile returns a boolean if a field has been set.
func (o *PrimaryContactInfoModel) HasMobile() bool {
	if o != nil && !IsNil(o.Mobile) {
		return true
	}

	return false
}

// SetMobile gets a reference to the given string and assigns it to the Mobile field.
func (o *PrimaryContactInfoModel) SetMobile(v string) {
	o.Mobile = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *PrimaryContactInfoModel) GetPhone() string {
	if o == nil || IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrimaryContactInfoModel) GetPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *PrimaryContactInfoModel) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *PrimaryContactInfoModel) SetPhone(v string) {
	o.Phone = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *PrimaryContactInfoModel) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrimaryContactInfoModel) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *PrimaryContactInfoModel) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *PrimaryContactInfoModel) SetTitle(v string) {
	o.Title = &v
}

func (o PrimaryContactInfoModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrimaryContactInfoModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Department) {
		toSerialize["department"] = o.Department
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Extension) {
		toSerialize["extension"] = o.Extension
	}
	if !IsNil(o.Fax) {
		toSerialize["fax"] = o.Fax
	}
	if !IsNil(o.FullName) {
		toSerialize["full_name"] = o.FullName
	}
	if !IsNil(o.Mobile) {
		toSerialize["mobile"] = o.Mobile
	}
	if !IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	return toSerialize, nil
}

type NullablePrimaryContactInfoModel struct {
	value *PrimaryContactInfoModel
	isSet bool
}

func (v NullablePrimaryContactInfoModel) Get() *PrimaryContactInfoModel {
	return v.value
}

func (v *NullablePrimaryContactInfoModel) Set(val *PrimaryContactInfoModel) {
	v.value = val
	v.isSet = true
}

func (v NullablePrimaryContactInfoModel) IsSet() bool {
	return v.isSet
}

func (v *NullablePrimaryContactInfoModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrimaryContactInfoModel(val *PrimaryContactInfoModel) *NullablePrimaryContactInfoModel {
	return &NullablePrimaryContactInfoModel{value: val, isSet: true}
}

func (v NullablePrimaryContactInfoModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrimaryContactInfoModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

