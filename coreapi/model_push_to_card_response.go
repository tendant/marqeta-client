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

// checks if the PushToCardResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PushToCardResponse{}

// PushToCardResponse struct for PushToCardResponse
type PushToCardResponse struct {
	Address1 *string `json:"address_1,omitempty"`
	Address2 *string `json:"address_2,omitempty"`
	City *string `json:"city,omitempty"`
	Country *string `json:"country,omitempty"`
	// yyyy-MM-ddTHH:mm:ssZ
	CreatedTime time.Time `json:"created_time"`
	ExpDate *string `json:"exp_date,omitempty"`
	FastFundTransferEligible *bool `json:"fast_fund_transfer_eligible,omitempty"`
	GamblingFundTransferEligible *bool `json:"gambling_fund_transfer_eligible,omitempty"`
	LastFour *string `json:"last_four,omitempty"`
	// yyyy-MM-ddTHH:mm:ssZ
	LastModifiedTime time.Time `json:"last_modified_time"`
	LastName *string `json:"last_name,omitempty"`
	NameOnCard *string `json:"name_on_card,omitempty"`
	PostalCode *string `json:"postal_code,omitempty"`
	State *string `json:"state,omitempty"`
	Token *string `json:"token,omitempty"`
}

// NewPushToCardResponse instantiates a new PushToCardResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPushToCardResponse(createdTime time.Time, lastModifiedTime time.Time) *PushToCardResponse {
	this := PushToCardResponse{}
	this.CreatedTime = createdTime
	var fastFundTransferEligible bool = false
	this.FastFundTransferEligible = &fastFundTransferEligible
	var gamblingFundTransferEligible bool = false
	this.GamblingFundTransferEligible = &gamblingFundTransferEligible
	this.LastModifiedTime = lastModifiedTime
	return &this
}

// NewPushToCardResponseWithDefaults instantiates a new PushToCardResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPushToCardResponseWithDefaults() *PushToCardResponse {
	this := PushToCardResponse{}
	var fastFundTransferEligible bool = false
	this.FastFundTransferEligible = &fastFundTransferEligible
	var gamblingFundTransferEligible bool = false
	this.GamblingFundTransferEligible = &gamblingFundTransferEligible
	return &this
}

// GetAddress1 returns the Address1 field value if set, zero value otherwise.
func (o *PushToCardResponse) GetAddress1() string {
	if o == nil || IsNil(o.Address1) {
		var ret string
		return ret
	}
	return *o.Address1
}

// GetAddress1Ok returns a tuple with the Address1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PushToCardResponse) GetAddress1Ok() (*string, bool) {
	if o == nil || IsNil(o.Address1) {
		return nil, false
	}
	return o.Address1, true
}

// HasAddress1 returns a boolean if a field has been set.
func (o *PushToCardResponse) HasAddress1() bool {
	if o != nil && !IsNil(o.Address1) {
		return true
	}

	return false
}

// SetAddress1 gets a reference to the given string and assigns it to the Address1 field.
func (o *PushToCardResponse) SetAddress1(v string) {
	o.Address1 = &v
}

// GetAddress2 returns the Address2 field value if set, zero value otherwise.
func (o *PushToCardResponse) GetAddress2() string {
	if o == nil || IsNil(o.Address2) {
		var ret string
		return ret
	}
	return *o.Address2
}

// GetAddress2Ok returns a tuple with the Address2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PushToCardResponse) GetAddress2Ok() (*string, bool) {
	if o == nil || IsNil(o.Address2) {
		return nil, false
	}
	return o.Address2, true
}

// HasAddress2 returns a boolean if a field has been set.
func (o *PushToCardResponse) HasAddress2() bool {
	if o != nil && !IsNil(o.Address2) {
		return true
	}

	return false
}

// SetAddress2 gets a reference to the given string and assigns it to the Address2 field.
func (o *PushToCardResponse) SetAddress2(v string) {
	o.Address2 = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *PushToCardResponse) GetCity() string {
	if o == nil || IsNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PushToCardResponse) GetCityOk() (*string, bool) {
	if o == nil || IsNil(o.City) {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *PushToCardResponse) HasCity() bool {
	if o != nil && !IsNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *PushToCardResponse) SetCity(v string) {
	o.City = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *PushToCardResponse) GetCountry() string {
	if o == nil || IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PushToCardResponse) GetCountryOk() (*string, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *PushToCardResponse) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *PushToCardResponse) SetCountry(v string) {
	o.Country = &v
}

// GetCreatedTime returns the CreatedTime field value
func (o *PushToCardResponse) GetCreatedTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value
// and a boolean to check if the value has been set.
func (o *PushToCardResponse) GetCreatedTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedTime, true
}

// SetCreatedTime sets field value
func (o *PushToCardResponse) SetCreatedTime(v time.Time) {
	o.CreatedTime = v
}

// GetExpDate returns the ExpDate field value if set, zero value otherwise.
func (o *PushToCardResponse) GetExpDate() string {
	if o == nil || IsNil(o.ExpDate) {
		var ret string
		return ret
	}
	return *o.ExpDate
}

// GetExpDateOk returns a tuple with the ExpDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PushToCardResponse) GetExpDateOk() (*string, bool) {
	if o == nil || IsNil(o.ExpDate) {
		return nil, false
	}
	return o.ExpDate, true
}

// HasExpDate returns a boolean if a field has been set.
func (o *PushToCardResponse) HasExpDate() bool {
	if o != nil && !IsNil(o.ExpDate) {
		return true
	}

	return false
}

// SetExpDate gets a reference to the given string and assigns it to the ExpDate field.
func (o *PushToCardResponse) SetExpDate(v string) {
	o.ExpDate = &v
}

// GetFastFundTransferEligible returns the FastFundTransferEligible field value if set, zero value otherwise.
func (o *PushToCardResponse) GetFastFundTransferEligible() bool {
	if o == nil || IsNil(o.FastFundTransferEligible) {
		var ret bool
		return ret
	}
	return *o.FastFundTransferEligible
}

// GetFastFundTransferEligibleOk returns a tuple with the FastFundTransferEligible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PushToCardResponse) GetFastFundTransferEligibleOk() (*bool, bool) {
	if o == nil || IsNil(o.FastFundTransferEligible) {
		return nil, false
	}
	return o.FastFundTransferEligible, true
}

// HasFastFundTransferEligible returns a boolean if a field has been set.
func (o *PushToCardResponse) HasFastFundTransferEligible() bool {
	if o != nil && !IsNil(o.FastFundTransferEligible) {
		return true
	}

	return false
}

// SetFastFundTransferEligible gets a reference to the given bool and assigns it to the FastFundTransferEligible field.
func (o *PushToCardResponse) SetFastFundTransferEligible(v bool) {
	o.FastFundTransferEligible = &v
}

// GetGamblingFundTransferEligible returns the GamblingFundTransferEligible field value if set, zero value otherwise.
func (o *PushToCardResponse) GetGamblingFundTransferEligible() bool {
	if o == nil || IsNil(o.GamblingFundTransferEligible) {
		var ret bool
		return ret
	}
	return *o.GamblingFundTransferEligible
}

// GetGamblingFundTransferEligibleOk returns a tuple with the GamblingFundTransferEligible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PushToCardResponse) GetGamblingFundTransferEligibleOk() (*bool, bool) {
	if o == nil || IsNil(o.GamblingFundTransferEligible) {
		return nil, false
	}
	return o.GamblingFundTransferEligible, true
}

// HasGamblingFundTransferEligible returns a boolean if a field has been set.
func (o *PushToCardResponse) HasGamblingFundTransferEligible() bool {
	if o != nil && !IsNil(o.GamblingFundTransferEligible) {
		return true
	}

	return false
}

// SetGamblingFundTransferEligible gets a reference to the given bool and assigns it to the GamblingFundTransferEligible field.
func (o *PushToCardResponse) SetGamblingFundTransferEligible(v bool) {
	o.GamblingFundTransferEligible = &v
}

// GetLastFour returns the LastFour field value if set, zero value otherwise.
func (o *PushToCardResponse) GetLastFour() string {
	if o == nil || IsNil(o.LastFour) {
		var ret string
		return ret
	}
	return *o.LastFour
}

// GetLastFourOk returns a tuple with the LastFour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PushToCardResponse) GetLastFourOk() (*string, bool) {
	if o == nil || IsNil(o.LastFour) {
		return nil, false
	}
	return o.LastFour, true
}

// HasLastFour returns a boolean if a field has been set.
func (o *PushToCardResponse) HasLastFour() bool {
	if o != nil && !IsNil(o.LastFour) {
		return true
	}

	return false
}

// SetLastFour gets a reference to the given string and assigns it to the LastFour field.
func (o *PushToCardResponse) SetLastFour(v string) {
	o.LastFour = &v
}

// GetLastModifiedTime returns the LastModifiedTime field value
func (o *PushToCardResponse) GetLastModifiedTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastModifiedTime
}

// GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field value
// and a boolean to check if the value has been set.
func (o *PushToCardResponse) GetLastModifiedTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastModifiedTime, true
}

// SetLastModifiedTime sets field value
func (o *PushToCardResponse) SetLastModifiedTime(v time.Time) {
	o.LastModifiedTime = v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *PushToCardResponse) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PushToCardResponse) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *PushToCardResponse) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *PushToCardResponse) SetLastName(v string) {
	o.LastName = &v
}

// GetNameOnCard returns the NameOnCard field value if set, zero value otherwise.
func (o *PushToCardResponse) GetNameOnCard() string {
	if o == nil || IsNil(o.NameOnCard) {
		var ret string
		return ret
	}
	return *o.NameOnCard
}

// GetNameOnCardOk returns a tuple with the NameOnCard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PushToCardResponse) GetNameOnCardOk() (*string, bool) {
	if o == nil || IsNil(o.NameOnCard) {
		return nil, false
	}
	return o.NameOnCard, true
}

// HasNameOnCard returns a boolean if a field has been set.
func (o *PushToCardResponse) HasNameOnCard() bool {
	if o != nil && !IsNil(o.NameOnCard) {
		return true
	}

	return false
}

// SetNameOnCard gets a reference to the given string and assigns it to the NameOnCard field.
func (o *PushToCardResponse) SetNameOnCard(v string) {
	o.NameOnCard = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *PushToCardResponse) GetPostalCode() string {
	if o == nil || IsNil(o.PostalCode) {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PushToCardResponse) GetPostalCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PostalCode) {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *PushToCardResponse) HasPostalCode() bool {
	if o != nil && !IsNil(o.PostalCode) {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *PushToCardResponse) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *PushToCardResponse) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PushToCardResponse) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *PushToCardResponse) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *PushToCardResponse) SetState(v string) {
	o.State = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *PushToCardResponse) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PushToCardResponse) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *PushToCardResponse) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *PushToCardResponse) SetToken(v string) {
	o.Token = &v
}

func (o PushToCardResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PushToCardResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address1) {
		toSerialize["address_1"] = o.Address1
	}
	if !IsNil(o.Address2) {
		toSerialize["address_2"] = o.Address2
	}
	if !IsNil(o.City) {
		toSerialize["city"] = o.City
	}
	if !IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	toSerialize["created_time"] = o.CreatedTime
	if !IsNil(o.ExpDate) {
		toSerialize["exp_date"] = o.ExpDate
	}
	if !IsNil(o.FastFundTransferEligible) {
		toSerialize["fast_fund_transfer_eligible"] = o.FastFundTransferEligible
	}
	if !IsNil(o.GamblingFundTransferEligible) {
		toSerialize["gambling_fund_transfer_eligible"] = o.GamblingFundTransferEligible
	}
	if !IsNil(o.LastFour) {
		toSerialize["last_four"] = o.LastFour
	}
	toSerialize["last_modified_time"] = o.LastModifiedTime
	if !IsNil(o.LastName) {
		toSerialize["last_name"] = o.LastName
	}
	if !IsNil(o.NameOnCard) {
		toSerialize["name_on_card"] = o.NameOnCard
	}
	if !IsNil(o.PostalCode) {
		toSerialize["postal_code"] = o.PostalCode
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	return toSerialize, nil
}

type NullablePushToCardResponse struct {
	value *PushToCardResponse
	isSet bool
}

func (v NullablePushToCardResponse) Get() *PushToCardResponse {
	return v.value
}

func (v *NullablePushToCardResponse) Set(val *PushToCardResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePushToCardResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePushToCardResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePushToCardResponse(val *PushToCardResponse) *NullablePushToCardResponse {
	return &NullablePushToCardResponse{value: val, isSet: true}
}

func (v NullablePushToCardResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePushToCardResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

