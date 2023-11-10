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

// checks if the MerchantResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantResponseModel{}

// MerchantResponseModel struct for MerchantResponseModel
type MerchantResponseModel struct {
	Active *bool `json:"active,omitempty"`
	Address1 *string `json:"address1,omitempty"`
	Address2 *string `json:"address2,omitempty"`
	City *string `json:"city,omitempty"`
	Contact *string `json:"contact,omitempty"`
	ContactEmail *string `json:"contact_email,omitempty"`
	Country *string `json:"country,omitempty"`
	// yyyy-MM-ddTHH:mm:ssZ
	CreatedTime time.Time `json:"created_time"`
	// yyyy-MM-ddTHH:mm:ssZ
	LastModifiedTime time.Time `json:"last_modified_time"`
	Latitude *float32 `json:"latitude,omitempty"`
	Longitude *float32 `json:"longitude,omitempty"`
	Name string `json:"name"`
	PartialAuthFlag *bool `json:"partial_auth_flag,omitempty"`
	Phone *string `json:"phone,omitempty"`
	Province *string `json:"province,omitempty"`
	State *string `json:"state,omitempty"`
	// The unique identifier of the merchant
	Token *string `json:"token,omitempty"`
	Zip *string `json:"zip,omitempty"`
}

// NewMerchantResponseModel instantiates a new MerchantResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantResponseModel(createdTime time.Time, lastModifiedTime time.Time, name string) *MerchantResponseModel {
	this := MerchantResponseModel{}
	var active bool = true
	this.Active = &active
	this.CreatedTime = createdTime
	this.LastModifiedTime = lastModifiedTime
	this.Name = name
	var partialAuthFlag bool = true
	this.PartialAuthFlag = &partialAuthFlag
	return &this
}

// NewMerchantResponseModelWithDefaults instantiates a new MerchantResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantResponseModelWithDefaults() *MerchantResponseModel {
	this := MerchantResponseModel{}
	var active bool = true
	this.Active = &active
	var partialAuthFlag bool = true
	this.PartialAuthFlag = &partialAuthFlag
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *MerchantResponseModel) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantResponseModel) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *MerchantResponseModel) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *MerchantResponseModel) SetActive(v bool) {
	o.Active = &v
}

// GetAddress1 returns the Address1 field value if set, zero value otherwise.
func (o *MerchantResponseModel) GetAddress1() string {
	if o == nil || IsNil(o.Address1) {
		var ret string
		return ret
	}
	return *o.Address1
}

// GetAddress1Ok returns a tuple with the Address1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantResponseModel) GetAddress1Ok() (*string, bool) {
	if o == nil || IsNil(o.Address1) {
		return nil, false
	}
	return o.Address1, true
}

// HasAddress1 returns a boolean if a field has been set.
func (o *MerchantResponseModel) HasAddress1() bool {
	if o != nil && !IsNil(o.Address1) {
		return true
	}

	return false
}

// SetAddress1 gets a reference to the given string and assigns it to the Address1 field.
func (o *MerchantResponseModel) SetAddress1(v string) {
	o.Address1 = &v
}

// GetAddress2 returns the Address2 field value if set, zero value otherwise.
func (o *MerchantResponseModel) GetAddress2() string {
	if o == nil || IsNil(o.Address2) {
		var ret string
		return ret
	}
	return *o.Address2
}

// GetAddress2Ok returns a tuple with the Address2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantResponseModel) GetAddress2Ok() (*string, bool) {
	if o == nil || IsNil(o.Address2) {
		return nil, false
	}
	return o.Address2, true
}

// HasAddress2 returns a boolean if a field has been set.
func (o *MerchantResponseModel) HasAddress2() bool {
	if o != nil && !IsNil(o.Address2) {
		return true
	}

	return false
}

// SetAddress2 gets a reference to the given string and assigns it to the Address2 field.
func (o *MerchantResponseModel) SetAddress2(v string) {
	o.Address2 = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *MerchantResponseModel) GetCity() string {
	if o == nil || IsNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantResponseModel) GetCityOk() (*string, bool) {
	if o == nil || IsNil(o.City) {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *MerchantResponseModel) HasCity() bool {
	if o != nil && !IsNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *MerchantResponseModel) SetCity(v string) {
	o.City = &v
}

// GetContact returns the Contact field value if set, zero value otherwise.
func (o *MerchantResponseModel) GetContact() string {
	if o == nil || IsNil(o.Contact) {
		var ret string
		return ret
	}
	return *o.Contact
}

// GetContactOk returns a tuple with the Contact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantResponseModel) GetContactOk() (*string, bool) {
	if o == nil || IsNil(o.Contact) {
		return nil, false
	}
	return o.Contact, true
}

// HasContact returns a boolean if a field has been set.
func (o *MerchantResponseModel) HasContact() bool {
	if o != nil && !IsNil(o.Contact) {
		return true
	}

	return false
}

// SetContact gets a reference to the given string and assigns it to the Contact field.
func (o *MerchantResponseModel) SetContact(v string) {
	o.Contact = &v
}

// GetContactEmail returns the ContactEmail field value if set, zero value otherwise.
func (o *MerchantResponseModel) GetContactEmail() string {
	if o == nil || IsNil(o.ContactEmail) {
		var ret string
		return ret
	}
	return *o.ContactEmail
}

// GetContactEmailOk returns a tuple with the ContactEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantResponseModel) GetContactEmailOk() (*string, bool) {
	if o == nil || IsNil(o.ContactEmail) {
		return nil, false
	}
	return o.ContactEmail, true
}

// HasContactEmail returns a boolean if a field has been set.
func (o *MerchantResponseModel) HasContactEmail() bool {
	if o != nil && !IsNil(o.ContactEmail) {
		return true
	}

	return false
}

// SetContactEmail gets a reference to the given string and assigns it to the ContactEmail field.
func (o *MerchantResponseModel) SetContactEmail(v string) {
	o.ContactEmail = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *MerchantResponseModel) GetCountry() string {
	if o == nil || IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantResponseModel) GetCountryOk() (*string, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *MerchantResponseModel) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *MerchantResponseModel) SetCountry(v string) {
	o.Country = &v
}

// GetCreatedTime returns the CreatedTime field value
func (o *MerchantResponseModel) GetCreatedTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value
// and a boolean to check if the value has been set.
func (o *MerchantResponseModel) GetCreatedTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedTime, true
}

// SetCreatedTime sets field value
func (o *MerchantResponseModel) SetCreatedTime(v time.Time) {
	o.CreatedTime = v
}

// GetLastModifiedTime returns the LastModifiedTime field value
func (o *MerchantResponseModel) GetLastModifiedTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastModifiedTime
}

// GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field value
// and a boolean to check if the value has been set.
func (o *MerchantResponseModel) GetLastModifiedTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastModifiedTime, true
}

// SetLastModifiedTime sets field value
func (o *MerchantResponseModel) SetLastModifiedTime(v time.Time) {
	o.LastModifiedTime = v
}

// GetLatitude returns the Latitude field value if set, zero value otherwise.
func (o *MerchantResponseModel) GetLatitude() float32 {
	if o == nil || IsNil(o.Latitude) {
		var ret float32
		return ret
	}
	return *o.Latitude
}

// GetLatitudeOk returns a tuple with the Latitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantResponseModel) GetLatitudeOk() (*float32, bool) {
	if o == nil || IsNil(o.Latitude) {
		return nil, false
	}
	return o.Latitude, true
}

// HasLatitude returns a boolean if a field has been set.
func (o *MerchantResponseModel) HasLatitude() bool {
	if o != nil && !IsNil(o.Latitude) {
		return true
	}

	return false
}

// SetLatitude gets a reference to the given float32 and assigns it to the Latitude field.
func (o *MerchantResponseModel) SetLatitude(v float32) {
	o.Latitude = &v
}

// GetLongitude returns the Longitude field value if set, zero value otherwise.
func (o *MerchantResponseModel) GetLongitude() float32 {
	if o == nil || IsNil(o.Longitude) {
		var ret float32
		return ret
	}
	return *o.Longitude
}

// GetLongitudeOk returns a tuple with the Longitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantResponseModel) GetLongitudeOk() (*float32, bool) {
	if o == nil || IsNil(o.Longitude) {
		return nil, false
	}
	return o.Longitude, true
}

// HasLongitude returns a boolean if a field has been set.
func (o *MerchantResponseModel) HasLongitude() bool {
	if o != nil && !IsNil(o.Longitude) {
		return true
	}

	return false
}

// SetLongitude gets a reference to the given float32 and assigns it to the Longitude field.
func (o *MerchantResponseModel) SetLongitude(v float32) {
	o.Longitude = &v
}

// GetName returns the Name field value
func (o *MerchantResponseModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MerchantResponseModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *MerchantResponseModel) SetName(v string) {
	o.Name = v
}

// GetPartialAuthFlag returns the PartialAuthFlag field value if set, zero value otherwise.
func (o *MerchantResponseModel) GetPartialAuthFlag() bool {
	if o == nil || IsNil(o.PartialAuthFlag) {
		var ret bool
		return ret
	}
	return *o.PartialAuthFlag
}

// GetPartialAuthFlagOk returns a tuple with the PartialAuthFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantResponseModel) GetPartialAuthFlagOk() (*bool, bool) {
	if o == nil || IsNil(o.PartialAuthFlag) {
		return nil, false
	}
	return o.PartialAuthFlag, true
}

// HasPartialAuthFlag returns a boolean if a field has been set.
func (o *MerchantResponseModel) HasPartialAuthFlag() bool {
	if o != nil && !IsNil(o.PartialAuthFlag) {
		return true
	}

	return false
}

// SetPartialAuthFlag gets a reference to the given bool and assigns it to the PartialAuthFlag field.
func (o *MerchantResponseModel) SetPartialAuthFlag(v bool) {
	o.PartialAuthFlag = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *MerchantResponseModel) GetPhone() string {
	if o == nil || IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantResponseModel) GetPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *MerchantResponseModel) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *MerchantResponseModel) SetPhone(v string) {
	o.Phone = &v
}

// GetProvince returns the Province field value if set, zero value otherwise.
func (o *MerchantResponseModel) GetProvince() string {
	if o == nil || IsNil(o.Province) {
		var ret string
		return ret
	}
	return *o.Province
}

// GetProvinceOk returns a tuple with the Province field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantResponseModel) GetProvinceOk() (*string, bool) {
	if o == nil || IsNil(o.Province) {
		return nil, false
	}
	return o.Province, true
}

// HasProvince returns a boolean if a field has been set.
func (o *MerchantResponseModel) HasProvince() bool {
	if o != nil && !IsNil(o.Province) {
		return true
	}

	return false
}

// SetProvince gets a reference to the given string and assigns it to the Province field.
func (o *MerchantResponseModel) SetProvince(v string) {
	o.Province = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *MerchantResponseModel) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantResponseModel) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *MerchantResponseModel) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *MerchantResponseModel) SetState(v string) {
	o.State = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *MerchantResponseModel) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantResponseModel) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *MerchantResponseModel) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *MerchantResponseModel) SetToken(v string) {
	o.Token = &v
}

// GetZip returns the Zip field value if set, zero value otherwise.
func (o *MerchantResponseModel) GetZip() string {
	if o == nil || IsNil(o.Zip) {
		var ret string
		return ret
	}
	return *o.Zip
}

// GetZipOk returns a tuple with the Zip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantResponseModel) GetZipOk() (*string, bool) {
	if o == nil || IsNil(o.Zip) {
		return nil, false
	}
	return o.Zip, true
}

// HasZip returns a boolean if a field has been set.
func (o *MerchantResponseModel) HasZip() bool {
	if o != nil && !IsNil(o.Zip) {
		return true
	}

	return false
}

// SetZip gets a reference to the given string and assigns it to the Zip field.
func (o *MerchantResponseModel) SetZip(v string) {
	o.Zip = &v
}

func (o MerchantResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.Address1) {
		toSerialize["address1"] = o.Address1
	}
	if !IsNil(o.Address2) {
		toSerialize["address2"] = o.Address2
	}
	if !IsNil(o.City) {
		toSerialize["city"] = o.City
	}
	if !IsNil(o.Contact) {
		toSerialize["contact"] = o.Contact
	}
	if !IsNil(o.ContactEmail) {
		toSerialize["contact_email"] = o.ContactEmail
	}
	if !IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	toSerialize["created_time"] = o.CreatedTime
	toSerialize["last_modified_time"] = o.LastModifiedTime
	if !IsNil(o.Latitude) {
		toSerialize["latitude"] = o.Latitude
	}
	if !IsNil(o.Longitude) {
		toSerialize["longitude"] = o.Longitude
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.PartialAuthFlag) {
		toSerialize["partial_auth_flag"] = o.PartialAuthFlag
	}
	if !IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	if !IsNil(o.Province) {
		toSerialize["province"] = o.Province
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	if !IsNil(o.Zip) {
		toSerialize["zip"] = o.Zip
	}
	return toSerialize, nil
}

type NullableMerchantResponseModel struct {
	value *MerchantResponseModel
	isSet bool
}

func (v NullableMerchantResponseModel) Get() *MerchantResponseModel {
	return v.value
}

func (v *NullableMerchantResponseModel) Set(val *MerchantResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantResponseModel(val *MerchantResponseModel) *NullableMerchantResponseModel {
	return &NullableMerchantResponseModel{value: val, isSet: true}
}

func (v NullableMerchantResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

