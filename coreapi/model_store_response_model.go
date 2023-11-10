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

// checks if the StoreResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StoreResponseModel{}

// StoreResponseModel struct for StoreResponseModel
type StoreResponseModel struct {
	Active *bool `json:"active,omitempty"`
	Address1 string `json:"address1"`
	Address2 *string `json:"address2,omitempty"`
	City string `json:"city"`
	Contact *string `json:"contact,omitempty"`
	ContactEmail *string `json:"contact_email,omitempty"`
	Country *string `json:"country,omitempty"`
	// yyyy-MM-ddTHH:mm:ssZ
	CreatedTime time.Time `json:"created_time"`
	KeyedAuthCvvEnforced *bool `json:"keyed_auth_cvv_enforced,omitempty"`
	// yyyy-MM-ddTHH:mm:ssZ
	LastModifiedTime time.Time `json:"last_modified_time"`
	Latitude *float32 `json:"latitude,omitempty"`
	Longitude *float32 `json:"longitude,omitempty"`
	MerchantToken string `json:"merchant_token"`
	Mid string `json:"mid"`
	Name string `json:"name"`
	NetworkMid *string `json:"network_mid,omitempty"`
	PartialApprovalCapable *bool `json:"partial_approval_capable,omitempty"`
	// 1 char max
	PartialAuthFlag *bool `json:"partial_auth_flag,omitempty"`
	Phone *string `json:"phone,omitempty"`
	PostalCode *string `json:"postal_code,omitempty"`
	Province *string `json:"province,omitempty"`
	State string `json:"state"`
	// The unique identifier of the merchant
	Token *string `json:"token,omitempty"`
	Zip *string `json:"zip,omitempty"`
}

// NewStoreResponseModel instantiates a new StoreResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoreResponseModel(address1 string, city string, createdTime time.Time, lastModifiedTime time.Time, merchantToken string, mid string, name string, state string) *StoreResponseModel {
	this := StoreResponseModel{}
	var active bool = true
	this.Active = &active
	this.Address1 = address1
	this.City = city
	this.CreatedTime = createdTime
	var keyedAuthCvvEnforced bool = false
	this.KeyedAuthCvvEnforced = &keyedAuthCvvEnforced
	this.LastModifiedTime = lastModifiedTime
	this.MerchantToken = merchantToken
	this.Mid = mid
	this.Name = name
	var partialApprovalCapable bool = false
	this.PartialApprovalCapable = &partialApprovalCapable
	var partialAuthFlag bool = true
	this.PartialAuthFlag = &partialAuthFlag
	this.State = state
	return &this
}

// NewStoreResponseModelWithDefaults instantiates a new StoreResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoreResponseModelWithDefaults() *StoreResponseModel {
	this := StoreResponseModel{}
	var active bool = true
	this.Active = &active
	var keyedAuthCvvEnforced bool = false
	this.KeyedAuthCvvEnforced = &keyedAuthCvvEnforced
	var partialApprovalCapable bool = false
	this.PartialApprovalCapable = &partialApprovalCapable
	var partialAuthFlag bool = true
	this.PartialAuthFlag = &partialAuthFlag
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *StoreResponseModel) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreResponseModel) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *StoreResponseModel) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *StoreResponseModel) SetActive(v bool) {
	o.Active = &v
}

// GetAddress1 returns the Address1 field value
func (o *StoreResponseModel) GetAddress1() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address1
}

// GetAddress1Ok returns a tuple with the Address1 field value
// and a boolean to check if the value has been set.
func (o *StoreResponseModel) GetAddress1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address1, true
}

// SetAddress1 sets field value
func (o *StoreResponseModel) SetAddress1(v string) {
	o.Address1 = v
}

// GetAddress2 returns the Address2 field value if set, zero value otherwise.
func (o *StoreResponseModel) GetAddress2() string {
	if o == nil || IsNil(o.Address2) {
		var ret string
		return ret
	}
	return *o.Address2
}

// GetAddress2Ok returns a tuple with the Address2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreResponseModel) GetAddress2Ok() (*string, bool) {
	if o == nil || IsNil(o.Address2) {
		return nil, false
	}
	return o.Address2, true
}

// HasAddress2 returns a boolean if a field has been set.
func (o *StoreResponseModel) HasAddress2() bool {
	if o != nil && !IsNil(o.Address2) {
		return true
	}

	return false
}

// SetAddress2 gets a reference to the given string and assigns it to the Address2 field.
func (o *StoreResponseModel) SetAddress2(v string) {
	o.Address2 = &v
}

// GetCity returns the City field value
func (o *StoreResponseModel) GetCity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.City
}

// GetCityOk returns a tuple with the City field value
// and a boolean to check if the value has been set.
func (o *StoreResponseModel) GetCityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.City, true
}

// SetCity sets field value
func (o *StoreResponseModel) SetCity(v string) {
	o.City = v
}

// GetContact returns the Contact field value if set, zero value otherwise.
func (o *StoreResponseModel) GetContact() string {
	if o == nil || IsNil(o.Contact) {
		var ret string
		return ret
	}
	return *o.Contact
}

// GetContactOk returns a tuple with the Contact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreResponseModel) GetContactOk() (*string, bool) {
	if o == nil || IsNil(o.Contact) {
		return nil, false
	}
	return o.Contact, true
}

// HasContact returns a boolean if a field has been set.
func (o *StoreResponseModel) HasContact() bool {
	if o != nil && !IsNil(o.Contact) {
		return true
	}

	return false
}

// SetContact gets a reference to the given string and assigns it to the Contact field.
func (o *StoreResponseModel) SetContact(v string) {
	o.Contact = &v
}

// GetContactEmail returns the ContactEmail field value if set, zero value otherwise.
func (o *StoreResponseModel) GetContactEmail() string {
	if o == nil || IsNil(o.ContactEmail) {
		var ret string
		return ret
	}
	return *o.ContactEmail
}

// GetContactEmailOk returns a tuple with the ContactEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreResponseModel) GetContactEmailOk() (*string, bool) {
	if o == nil || IsNil(o.ContactEmail) {
		return nil, false
	}
	return o.ContactEmail, true
}

// HasContactEmail returns a boolean if a field has been set.
func (o *StoreResponseModel) HasContactEmail() bool {
	if o != nil && !IsNil(o.ContactEmail) {
		return true
	}

	return false
}

// SetContactEmail gets a reference to the given string and assigns it to the ContactEmail field.
func (o *StoreResponseModel) SetContactEmail(v string) {
	o.ContactEmail = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *StoreResponseModel) GetCountry() string {
	if o == nil || IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreResponseModel) GetCountryOk() (*string, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *StoreResponseModel) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *StoreResponseModel) SetCountry(v string) {
	o.Country = &v
}

// GetCreatedTime returns the CreatedTime field value
func (o *StoreResponseModel) GetCreatedTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value
// and a boolean to check if the value has been set.
func (o *StoreResponseModel) GetCreatedTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedTime, true
}

// SetCreatedTime sets field value
func (o *StoreResponseModel) SetCreatedTime(v time.Time) {
	o.CreatedTime = v
}

// GetKeyedAuthCvvEnforced returns the KeyedAuthCvvEnforced field value if set, zero value otherwise.
func (o *StoreResponseModel) GetKeyedAuthCvvEnforced() bool {
	if o == nil || IsNil(o.KeyedAuthCvvEnforced) {
		var ret bool
		return ret
	}
	return *o.KeyedAuthCvvEnforced
}

// GetKeyedAuthCvvEnforcedOk returns a tuple with the KeyedAuthCvvEnforced field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreResponseModel) GetKeyedAuthCvvEnforcedOk() (*bool, bool) {
	if o == nil || IsNil(o.KeyedAuthCvvEnforced) {
		return nil, false
	}
	return o.KeyedAuthCvvEnforced, true
}

// HasKeyedAuthCvvEnforced returns a boolean if a field has been set.
func (o *StoreResponseModel) HasKeyedAuthCvvEnforced() bool {
	if o != nil && !IsNil(o.KeyedAuthCvvEnforced) {
		return true
	}

	return false
}

// SetKeyedAuthCvvEnforced gets a reference to the given bool and assigns it to the KeyedAuthCvvEnforced field.
func (o *StoreResponseModel) SetKeyedAuthCvvEnforced(v bool) {
	o.KeyedAuthCvvEnforced = &v
}

// GetLastModifiedTime returns the LastModifiedTime field value
func (o *StoreResponseModel) GetLastModifiedTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastModifiedTime
}

// GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field value
// and a boolean to check if the value has been set.
func (o *StoreResponseModel) GetLastModifiedTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastModifiedTime, true
}

// SetLastModifiedTime sets field value
func (o *StoreResponseModel) SetLastModifiedTime(v time.Time) {
	o.LastModifiedTime = v
}

// GetLatitude returns the Latitude field value if set, zero value otherwise.
func (o *StoreResponseModel) GetLatitude() float32 {
	if o == nil || IsNil(o.Latitude) {
		var ret float32
		return ret
	}
	return *o.Latitude
}

// GetLatitudeOk returns a tuple with the Latitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreResponseModel) GetLatitudeOk() (*float32, bool) {
	if o == nil || IsNil(o.Latitude) {
		return nil, false
	}
	return o.Latitude, true
}

// HasLatitude returns a boolean if a field has been set.
func (o *StoreResponseModel) HasLatitude() bool {
	if o != nil && !IsNil(o.Latitude) {
		return true
	}

	return false
}

// SetLatitude gets a reference to the given float32 and assigns it to the Latitude field.
func (o *StoreResponseModel) SetLatitude(v float32) {
	o.Latitude = &v
}

// GetLongitude returns the Longitude field value if set, zero value otherwise.
func (o *StoreResponseModel) GetLongitude() float32 {
	if o == nil || IsNil(o.Longitude) {
		var ret float32
		return ret
	}
	return *o.Longitude
}

// GetLongitudeOk returns a tuple with the Longitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreResponseModel) GetLongitudeOk() (*float32, bool) {
	if o == nil || IsNil(o.Longitude) {
		return nil, false
	}
	return o.Longitude, true
}

// HasLongitude returns a boolean if a field has been set.
func (o *StoreResponseModel) HasLongitude() bool {
	if o != nil && !IsNil(o.Longitude) {
		return true
	}

	return false
}

// SetLongitude gets a reference to the given float32 and assigns it to the Longitude field.
func (o *StoreResponseModel) SetLongitude(v float32) {
	o.Longitude = &v
}

// GetMerchantToken returns the MerchantToken field value
func (o *StoreResponseModel) GetMerchantToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantToken
}

// GetMerchantTokenOk returns a tuple with the MerchantToken field value
// and a boolean to check if the value has been set.
func (o *StoreResponseModel) GetMerchantTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantToken, true
}

// SetMerchantToken sets field value
func (o *StoreResponseModel) SetMerchantToken(v string) {
	o.MerchantToken = v
}

// GetMid returns the Mid field value
func (o *StoreResponseModel) GetMid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mid
}

// GetMidOk returns a tuple with the Mid field value
// and a boolean to check if the value has been set.
func (o *StoreResponseModel) GetMidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mid, true
}

// SetMid sets field value
func (o *StoreResponseModel) SetMid(v string) {
	o.Mid = v
}

// GetName returns the Name field value
func (o *StoreResponseModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *StoreResponseModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *StoreResponseModel) SetName(v string) {
	o.Name = v
}

// GetNetworkMid returns the NetworkMid field value if set, zero value otherwise.
func (o *StoreResponseModel) GetNetworkMid() string {
	if o == nil || IsNil(o.NetworkMid) {
		var ret string
		return ret
	}
	return *o.NetworkMid
}

// GetNetworkMidOk returns a tuple with the NetworkMid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreResponseModel) GetNetworkMidOk() (*string, bool) {
	if o == nil || IsNil(o.NetworkMid) {
		return nil, false
	}
	return o.NetworkMid, true
}

// HasNetworkMid returns a boolean if a field has been set.
func (o *StoreResponseModel) HasNetworkMid() bool {
	if o != nil && !IsNil(o.NetworkMid) {
		return true
	}

	return false
}

// SetNetworkMid gets a reference to the given string and assigns it to the NetworkMid field.
func (o *StoreResponseModel) SetNetworkMid(v string) {
	o.NetworkMid = &v
}

// GetPartialApprovalCapable returns the PartialApprovalCapable field value if set, zero value otherwise.
func (o *StoreResponseModel) GetPartialApprovalCapable() bool {
	if o == nil || IsNil(o.PartialApprovalCapable) {
		var ret bool
		return ret
	}
	return *o.PartialApprovalCapable
}

// GetPartialApprovalCapableOk returns a tuple with the PartialApprovalCapable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreResponseModel) GetPartialApprovalCapableOk() (*bool, bool) {
	if o == nil || IsNil(o.PartialApprovalCapable) {
		return nil, false
	}
	return o.PartialApprovalCapable, true
}

// HasPartialApprovalCapable returns a boolean if a field has been set.
func (o *StoreResponseModel) HasPartialApprovalCapable() bool {
	if o != nil && !IsNil(o.PartialApprovalCapable) {
		return true
	}

	return false
}

// SetPartialApprovalCapable gets a reference to the given bool and assigns it to the PartialApprovalCapable field.
func (o *StoreResponseModel) SetPartialApprovalCapable(v bool) {
	o.PartialApprovalCapable = &v
}

// GetPartialAuthFlag returns the PartialAuthFlag field value if set, zero value otherwise.
func (o *StoreResponseModel) GetPartialAuthFlag() bool {
	if o == nil || IsNil(o.PartialAuthFlag) {
		var ret bool
		return ret
	}
	return *o.PartialAuthFlag
}

// GetPartialAuthFlagOk returns a tuple with the PartialAuthFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreResponseModel) GetPartialAuthFlagOk() (*bool, bool) {
	if o == nil || IsNil(o.PartialAuthFlag) {
		return nil, false
	}
	return o.PartialAuthFlag, true
}

// HasPartialAuthFlag returns a boolean if a field has been set.
func (o *StoreResponseModel) HasPartialAuthFlag() bool {
	if o != nil && !IsNil(o.PartialAuthFlag) {
		return true
	}

	return false
}

// SetPartialAuthFlag gets a reference to the given bool and assigns it to the PartialAuthFlag field.
func (o *StoreResponseModel) SetPartialAuthFlag(v bool) {
	o.PartialAuthFlag = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *StoreResponseModel) GetPhone() string {
	if o == nil || IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreResponseModel) GetPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *StoreResponseModel) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *StoreResponseModel) SetPhone(v string) {
	o.Phone = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *StoreResponseModel) GetPostalCode() string {
	if o == nil || IsNil(o.PostalCode) {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreResponseModel) GetPostalCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PostalCode) {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *StoreResponseModel) HasPostalCode() bool {
	if o != nil && !IsNil(o.PostalCode) {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *StoreResponseModel) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetProvince returns the Province field value if set, zero value otherwise.
func (o *StoreResponseModel) GetProvince() string {
	if o == nil || IsNil(o.Province) {
		var ret string
		return ret
	}
	return *o.Province
}

// GetProvinceOk returns a tuple with the Province field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreResponseModel) GetProvinceOk() (*string, bool) {
	if o == nil || IsNil(o.Province) {
		return nil, false
	}
	return o.Province, true
}

// HasProvince returns a boolean if a field has been set.
func (o *StoreResponseModel) HasProvince() bool {
	if o != nil && !IsNil(o.Province) {
		return true
	}

	return false
}

// SetProvince gets a reference to the given string and assigns it to the Province field.
func (o *StoreResponseModel) SetProvince(v string) {
	o.Province = &v
}

// GetState returns the State field value
func (o *StoreResponseModel) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *StoreResponseModel) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *StoreResponseModel) SetState(v string) {
	o.State = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *StoreResponseModel) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreResponseModel) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *StoreResponseModel) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *StoreResponseModel) SetToken(v string) {
	o.Token = &v
}

// GetZip returns the Zip field value if set, zero value otherwise.
func (o *StoreResponseModel) GetZip() string {
	if o == nil || IsNil(o.Zip) {
		var ret string
		return ret
	}
	return *o.Zip
}

// GetZipOk returns a tuple with the Zip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreResponseModel) GetZipOk() (*string, bool) {
	if o == nil || IsNil(o.Zip) {
		return nil, false
	}
	return o.Zip, true
}

// HasZip returns a boolean if a field has been set.
func (o *StoreResponseModel) HasZip() bool {
	if o != nil && !IsNil(o.Zip) {
		return true
	}

	return false
}

// SetZip gets a reference to the given string and assigns it to the Zip field.
func (o *StoreResponseModel) SetZip(v string) {
	o.Zip = &v
}

func (o StoreResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StoreResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	toSerialize["address1"] = o.Address1
	if !IsNil(o.Address2) {
		toSerialize["address2"] = o.Address2
	}
	toSerialize["city"] = o.City
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
	if !IsNil(o.KeyedAuthCvvEnforced) {
		toSerialize["keyed_auth_cvv_enforced"] = o.KeyedAuthCvvEnforced
	}
	toSerialize["last_modified_time"] = o.LastModifiedTime
	if !IsNil(o.Latitude) {
		toSerialize["latitude"] = o.Latitude
	}
	if !IsNil(o.Longitude) {
		toSerialize["longitude"] = o.Longitude
	}
	toSerialize["merchant_token"] = o.MerchantToken
	toSerialize["mid"] = o.Mid
	toSerialize["name"] = o.Name
	if !IsNil(o.NetworkMid) {
		toSerialize["network_mid"] = o.NetworkMid
	}
	if !IsNil(o.PartialApprovalCapable) {
		toSerialize["partial_approval_capable"] = o.PartialApprovalCapable
	}
	if !IsNil(o.PartialAuthFlag) {
		toSerialize["partial_auth_flag"] = o.PartialAuthFlag
	}
	if !IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	if !IsNil(o.PostalCode) {
		toSerialize["postal_code"] = o.PostalCode
	}
	if !IsNil(o.Province) {
		toSerialize["province"] = o.Province
	}
	toSerialize["state"] = o.State
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	if !IsNil(o.Zip) {
		toSerialize["zip"] = o.Zip
	}
	return toSerialize, nil
}

type NullableStoreResponseModel struct {
	value *StoreResponseModel
	isSet bool
}

func (v NullableStoreResponseModel) Get() *StoreResponseModel {
	return v.value
}

func (v *NullableStoreResponseModel) Set(val *StoreResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableStoreResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableStoreResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoreResponseModel(val *StoreResponseModel) *NullableStoreResponseModel {
	return &NullableStoreResponseModel{value: val, isSet: true}
}

func (v NullableStoreResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoreResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


