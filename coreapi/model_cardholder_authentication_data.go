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
)

// checks if the CardholderAuthenticationData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CardholderAuthenticationData{}

// CardholderAuthenticationData Contains 3D Secure authentication data:  * `electronic_commerce_indicator`  The level of verification performed. * `verification_result`  The result of the verification. * `verification_value_created_by`  The transaction participant who determined the verification result. * `three_ds_message_version`  The 3D Secure message version used for authentication. * `authentication_method`  The 3D Secure authentication method. * `authentication_status`  The 3D Secure authentication status. * `acquirer_exemption`  Indicates a 3D Secure authentication exemption from the acquirer. * `issuer_exemption`  Indicates a 3D Secure authentication exemption from the issuer.
type CardholderAuthenticationData struct {
	// Indicates 3D Secure authentication exemptions from the acquirer. This array is returned if it is included in the transaction data from the card network.
	AcquirerExemption []string `json:"acquirer_exemption,omitempty"`
	// Specifies the 3D Secure authentication method.
	AuthenticationMethod *string `json:"authentication_method,omitempty"`
	// Specifies the status of the 3D Secure authentication:  * `ATTEMPTED`: Indicates that 3D Secure authentication was requested and processed by the card network. * `DATA_SHARE_EXEMPTION`: Indicates that 3D Secure authentication was for information only and exempted. * `EXEMPTED`: Indicates that 3D Secure authentication was requested but the challenge was exempted. * `EXEMPTION_CLAIMED`: Indicates that 3D Secure authentication was exempted because acquirer transaction risk analysis (TRA) was already performed. * `SCA_EXEMPTION`: Indicates that 3D Secure authentication was exempted because strong customer authentication (SCA) was already performed. * `SUCCESSFUL`: Indicates that 3D Secure authentication was successful. * `SUCCESSFUL_NON_PAYMENT`: Indicates that 3D Secure non-payment authentication was successful. * `THREEDS_REQUESTER_DATA_SHARE_EXEMPTION`: Status is deprecated and will be removed from a future release of the Marqeta platform. After June 2023, use `DATA_SHARE_EXEMPTION` instead. * `THREEDS_REQUESTER_SCA_EXEMPTION`: Status is deprecated and will be removed in a June 2023 release of the Marqeta platform. After June 2023, use `SCA_EXEMPTION` instead. * `THREEDS_REQUESTER_TRA_EXEMPTION`: Status is deprecated and will be removed in a June 2023 release of the Marqeta platform. After June 2023, use `EXEMPTION_CLAIMED` instead. * `UNAVAILABLE`: ** For Visa transactions, this status indicates that 3D Secure authentication was requested, but Marqeta's access control server (ACS) was not available. ** For Mastercard transactions, this status indicates that 3D Secure authentication was not available.
	AuthenticationStatus *string `json:"authentication_status,omitempty"`
	// Authentication amount from the cardholder authentication verification value (CAVV) used to validate an authorization. This field is returned if it is included in the transaction data from the card network.
	CavvAuthenticationAmount *string `json:"cavv_authentication_amount,omitempty"`
	// Status of the 3D Secure authentication attempt, as provided by a transaction participant.  * `authentication_attempted`: Merchant attempted to authenticate, but either the issuer or the cardholder does not participate in 3D Secure. * `authentication_successful`: Cardholder authentication successful. * `no_authentication`: Non-authenticated e-commerce transaction.
	ElectronicCommerceIndicator *string `json:"electronic_commerce_indicator,omitempty"`
	// Indicates a 3D Secure authentication exemption from the issuer This field is returned if it is included in the transaction data from the card network.
	IssuerExemption *string `json:"issuer_exemption,omitempty"`
	// Raw cardholder authentication verification value provided by the card network. This field is returned if it is included in the transaction data from the card network.
	RawCavvData *string `json:"raw_cavv_data,omitempty"`
	// Specifies the 3D Secure message version used for authentication.
	ThreeDsMessageVersion *string `json:"three_ds_message_version,omitempty"`
	// Result of cardholder authentication verification value (CAVV) or accountholder authentication value (AAV) verification performed by the card network.  * `failed` * `not_present` * `not_provided` * `not_verified` * `not_verified_authentication_outage` * `verified` * `verified_amount_checked` * `verified_amount_greater_than_20_percent`: For Mastercard AAV verification, indicates that the original authentication amount and final authorization amount are mismatched, and that the final authorization amount exceeds the original authentication amount by more than 20%. This 20% margin falls outside Mastercard's suggested tolerance for what a European cardholder might reasonably expect when the total transaction amount is not known in advance. * `verified_amount_less_than_20_percent`: For Mastercard AAV verification, indicates that the original authentication amount and final authorization amount are mismatched, and that the final authorization amount exceeds the original authentication amount by 20% or less. This 20% margin falls within Mastercard's suggested tolerance for what a European cardholder might reasonably expect when the total transaction amount is not known in advance.
	VerificationResult *string `json:"verification_result,omitempty"`
	// Transaction participant who determined the verification result.
	VerificationValueCreatedBy *string `json:"verification_value_created_by,omitempty"`
}

// NewCardholderAuthenticationData instantiates a new CardholderAuthenticationData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardholderAuthenticationData() *CardholderAuthenticationData {
	this := CardholderAuthenticationData{}
	return &this
}

// NewCardholderAuthenticationDataWithDefaults instantiates a new CardholderAuthenticationData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardholderAuthenticationDataWithDefaults() *CardholderAuthenticationData {
	this := CardholderAuthenticationData{}
	return &this
}

// GetAcquirerExemption returns the AcquirerExemption field value if set, zero value otherwise.
func (o *CardholderAuthenticationData) GetAcquirerExemption() []string {
	if o == nil || IsNil(o.AcquirerExemption) {
		var ret []string
		return ret
	}
	return o.AcquirerExemption
}

// GetAcquirerExemptionOk returns a tuple with the AcquirerExemption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardholderAuthenticationData) GetAcquirerExemptionOk() ([]string, bool) {
	if o == nil || IsNil(o.AcquirerExemption) {
		return nil, false
	}
	return o.AcquirerExemption, true
}

// HasAcquirerExemption returns a boolean if a field has been set.
func (o *CardholderAuthenticationData) HasAcquirerExemption() bool {
	if o != nil && !IsNil(o.AcquirerExemption) {
		return true
	}

	return false
}

// SetAcquirerExemption gets a reference to the given []string and assigns it to the AcquirerExemption field.
func (o *CardholderAuthenticationData) SetAcquirerExemption(v []string) {
	o.AcquirerExemption = v
}

// GetAuthenticationMethod returns the AuthenticationMethod field value if set, zero value otherwise.
func (o *CardholderAuthenticationData) GetAuthenticationMethod() string {
	if o == nil || IsNil(o.AuthenticationMethod) {
		var ret string
		return ret
	}
	return *o.AuthenticationMethod
}

// GetAuthenticationMethodOk returns a tuple with the AuthenticationMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardholderAuthenticationData) GetAuthenticationMethodOk() (*string, bool) {
	if o == nil || IsNil(o.AuthenticationMethod) {
		return nil, false
	}
	return o.AuthenticationMethod, true
}

// HasAuthenticationMethod returns a boolean if a field has been set.
func (o *CardholderAuthenticationData) HasAuthenticationMethod() bool {
	if o != nil && !IsNil(o.AuthenticationMethod) {
		return true
	}

	return false
}

// SetAuthenticationMethod gets a reference to the given string and assigns it to the AuthenticationMethod field.
func (o *CardholderAuthenticationData) SetAuthenticationMethod(v string) {
	o.AuthenticationMethod = &v
}

// GetAuthenticationStatus returns the AuthenticationStatus field value if set, zero value otherwise.
func (o *CardholderAuthenticationData) GetAuthenticationStatus() string {
	if o == nil || IsNil(o.AuthenticationStatus) {
		var ret string
		return ret
	}
	return *o.AuthenticationStatus
}

// GetAuthenticationStatusOk returns a tuple with the AuthenticationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardholderAuthenticationData) GetAuthenticationStatusOk() (*string, bool) {
	if o == nil || IsNil(o.AuthenticationStatus) {
		return nil, false
	}
	return o.AuthenticationStatus, true
}

// HasAuthenticationStatus returns a boolean if a field has been set.
func (o *CardholderAuthenticationData) HasAuthenticationStatus() bool {
	if o != nil && !IsNil(o.AuthenticationStatus) {
		return true
	}

	return false
}

// SetAuthenticationStatus gets a reference to the given string and assigns it to the AuthenticationStatus field.
func (o *CardholderAuthenticationData) SetAuthenticationStatus(v string) {
	o.AuthenticationStatus = &v
}

// GetCavvAuthenticationAmount returns the CavvAuthenticationAmount field value if set, zero value otherwise.
func (o *CardholderAuthenticationData) GetCavvAuthenticationAmount() string {
	if o == nil || IsNil(o.CavvAuthenticationAmount) {
		var ret string
		return ret
	}
	return *o.CavvAuthenticationAmount
}

// GetCavvAuthenticationAmountOk returns a tuple with the CavvAuthenticationAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardholderAuthenticationData) GetCavvAuthenticationAmountOk() (*string, bool) {
	if o == nil || IsNil(o.CavvAuthenticationAmount) {
		return nil, false
	}
	return o.CavvAuthenticationAmount, true
}

// HasCavvAuthenticationAmount returns a boolean if a field has been set.
func (o *CardholderAuthenticationData) HasCavvAuthenticationAmount() bool {
	if o != nil && !IsNil(o.CavvAuthenticationAmount) {
		return true
	}

	return false
}

// SetCavvAuthenticationAmount gets a reference to the given string and assigns it to the CavvAuthenticationAmount field.
func (o *CardholderAuthenticationData) SetCavvAuthenticationAmount(v string) {
	o.CavvAuthenticationAmount = &v
}

// GetElectronicCommerceIndicator returns the ElectronicCommerceIndicator field value if set, zero value otherwise.
func (o *CardholderAuthenticationData) GetElectronicCommerceIndicator() string {
	if o == nil || IsNil(o.ElectronicCommerceIndicator) {
		var ret string
		return ret
	}
	return *o.ElectronicCommerceIndicator
}

// GetElectronicCommerceIndicatorOk returns a tuple with the ElectronicCommerceIndicator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardholderAuthenticationData) GetElectronicCommerceIndicatorOk() (*string, bool) {
	if o == nil || IsNil(o.ElectronicCommerceIndicator) {
		return nil, false
	}
	return o.ElectronicCommerceIndicator, true
}

// HasElectronicCommerceIndicator returns a boolean if a field has been set.
func (o *CardholderAuthenticationData) HasElectronicCommerceIndicator() bool {
	if o != nil && !IsNil(o.ElectronicCommerceIndicator) {
		return true
	}

	return false
}

// SetElectronicCommerceIndicator gets a reference to the given string and assigns it to the ElectronicCommerceIndicator field.
func (o *CardholderAuthenticationData) SetElectronicCommerceIndicator(v string) {
	o.ElectronicCommerceIndicator = &v
}

// GetIssuerExemption returns the IssuerExemption field value if set, zero value otherwise.
func (o *CardholderAuthenticationData) GetIssuerExemption() string {
	if o == nil || IsNil(o.IssuerExemption) {
		var ret string
		return ret
	}
	return *o.IssuerExemption
}

// GetIssuerExemptionOk returns a tuple with the IssuerExemption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardholderAuthenticationData) GetIssuerExemptionOk() (*string, bool) {
	if o == nil || IsNil(o.IssuerExemption) {
		return nil, false
	}
	return o.IssuerExemption, true
}

// HasIssuerExemption returns a boolean if a field has been set.
func (o *CardholderAuthenticationData) HasIssuerExemption() bool {
	if o != nil && !IsNil(o.IssuerExemption) {
		return true
	}

	return false
}

// SetIssuerExemption gets a reference to the given string and assigns it to the IssuerExemption field.
func (o *CardholderAuthenticationData) SetIssuerExemption(v string) {
	o.IssuerExemption = &v
}

// GetRawCavvData returns the RawCavvData field value if set, zero value otherwise.
func (o *CardholderAuthenticationData) GetRawCavvData() string {
	if o == nil || IsNil(o.RawCavvData) {
		var ret string
		return ret
	}
	return *o.RawCavvData
}

// GetRawCavvDataOk returns a tuple with the RawCavvData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardholderAuthenticationData) GetRawCavvDataOk() (*string, bool) {
	if o == nil || IsNil(o.RawCavvData) {
		return nil, false
	}
	return o.RawCavvData, true
}

// HasRawCavvData returns a boolean if a field has been set.
func (o *CardholderAuthenticationData) HasRawCavvData() bool {
	if o != nil && !IsNil(o.RawCavvData) {
		return true
	}

	return false
}

// SetRawCavvData gets a reference to the given string and assigns it to the RawCavvData field.
func (o *CardholderAuthenticationData) SetRawCavvData(v string) {
	o.RawCavvData = &v
}

// GetThreeDsMessageVersion returns the ThreeDsMessageVersion field value if set, zero value otherwise.
func (o *CardholderAuthenticationData) GetThreeDsMessageVersion() string {
	if o == nil || IsNil(o.ThreeDsMessageVersion) {
		var ret string
		return ret
	}
	return *o.ThreeDsMessageVersion
}

// GetThreeDsMessageVersionOk returns a tuple with the ThreeDsMessageVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardholderAuthenticationData) GetThreeDsMessageVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ThreeDsMessageVersion) {
		return nil, false
	}
	return o.ThreeDsMessageVersion, true
}

// HasThreeDsMessageVersion returns a boolean if a field has been set.
func (o *CardholderAuthenticationData) HasThreeDsMessageVersion() bool {
	if o != nil && !IsNil(o.ThreeDsMessageVersion) {
		return true
	}

	return false
}

// SetThreeDsMessageVersion gets a reference to the given string and assigns it to the ThreeDsMessageVersion field.
func (o *CardholderAuthenticationData) SetThreeDsMessageVersion(v string) {
	o.ThreeDsMessageVersion = &v
}

// GetVerificationResult returns the VerificationResult field value if set, zero value otherwise.
func (o *CardholderAuthenticationData) GetVerificationResult() string {
	if o == nil || IsNil(o.VerificationResult) {
		var ret string
		return ret
	}
	return *o.VerificationResult
}

// GetVerificationResultOk returns a tuple with the VerificationResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardholderAuthenticationData) GetVerificationResultOk() (*string, bool) {
	if o == nil || IsNil(o.VerificationResult) {
		return nil, false
	}
	return o.VerificationResult, true
}

// HasVerificationResult returns a boolean if a field has been set.
func (o *CardholderAuthenticationData) HasVerificationResult() bool {
	if o != nil && !IsNil(o.VerificationResult) {
		return true
	}

	return false
}

// SetVerificationResult gets a reference to the given string and assigns it to the VerificationResult field.
func (o *CardholderAuthenticationData) SetVerificationResult(v string) {
	o.VerificationResult = &v
}

// GetVerificationValueCreatedBy returns the VerificationValueCreatedBy field value if set, zero value otherwise.
func (o *CardholderAuthenticationData) GetVerificationValueCreatedBy() string {
	if o == nil || IsNil(o.VerificationValueCreatedBy) {
		var ret string
		return ret
	}
	return *o.VerificationValueCreatedBy
}

// GetVerificationValueCreatedByOk returns a tuple with the VerificationValueCreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardholderAuthenticationData) GetVerificationValueCreatedByOk() (*string, bool) {
	if o == nil || IsNil(o.VerificationValueCreatedBy) {
		return nil, false
	}
	return o.VerificationValueCreatedBy, true
}

// HasVerificationValueCreatedBy returns a boolean if a field has been set.
func (o *CardholderAuthenticationData) HasVerificationValueCreatedBy() bool {
	if o != nil && !IsNil(o.VerificationValueCreatedBy) {
		return true
	}

	return false
}

// SetVerificationValueCreatedBy gets a reference to the given string and assigns it to the VerificationValueCreatedBy field.
func (o *CardholderAuthenticationData) SetVerificationValueCreatedBy(v string) {
	o.VerificationValueCreatedBy = &v
}

func (o CardholderAuthenticationData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CardholderAuthenticationData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AcquirerExemption) {
		toSerialize["acquirer_exemption"] = o.AcquirerExemption
	}
	if !IsNil(o.AuthenticationMethod) {
		toSerialize["authentication_method"] = o.AuthenticationMethod
	}
	if !IsNil(o.AuthenticationStatus) {
		toSerialize["authentication_status"] = o.AuthenticationStatus
	}
	if !IsNil(o.CavvAuthenticationAmount) {
		toSerialize["cavv_authentication_amount"] = o.CavvAuthenticationAmount
	}
	if !IsNil(o.ElectronicCommerceIndicator) {
		toSerialize["electronic_commerce_indicator"] = o.ElectronicCommerceIndicator
	}
	if !IsNil(o.IssuerExemption) {
		toSerialize["issuer_exemption"] = o.IssuerExemption
	}
	if !IsNil(o.RawCavvData) {
		toSerialize["raw_cavv_data"] = o.RawCavvData
	}
	if !IsNil(o.ThreeDsMessageVersion) {
		toSerialize["three_ds_message_version"] = o.ThreeDsMessageVersion
	}
	if !IsNil(o.VerificationResult) {
		toSerialize["verification_result"] = o.VerificationResult
	}
	if !IsNil(o.VerificationValueCreatedBy) {
		toSerialize["verification_value_created_by"] = o.VerificationValueCreatedBy
	}
	return toSerialize, nil
}

type NullableCardholderAuthenticationData struct {
	value *CardholderAuthenticationData
	isSet bool
}

func (v NullableCardholderAuthenticationData) Get() *CardholderAuthenticationData {
	return v.value
}

func (v *NullableCardholderAuthenticationData) Set(val *CardholderAuthenticationData) {
	v.value = val
	v.isSet = true
}

func (v NullableCardholderAuthenticationData) IsSet() bool {
	return v.isSet
}

func (v *NullableCardholderAuthenticationData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardholderAuthenticationData(val *CardholderAuthenticationData) *NullableCardholderAuthenticationData {
	return &NullableCardholderAuthenticationData{value: val, isSet: true}
}

func (v NullableCardholderAuthenticationData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardholderAuthenticationData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


