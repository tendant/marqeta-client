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

// checks if the ProvisioningControls type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProvisioningControls{}

// ProvisioningControls struct for ProvisioningControls
type ProvisioningControls struct {
	DwtUseCardStatusDuringAuth *bool `json:"dwt_use_card_status_during_auth,omitempty"`
	DwtVerifyAtcDuringAuth *bool `json:"dwt_verify_atc_during_auth,omitempty"`
	// A value of `true` requires identity verification set-up for all digital wallets at the card product level.
	ForceYellowPathForCardProduct *bool `json:"force_yellow_path_for_card_product,omitempty"`
	InAppProvisioning *InAppProvisioning `json:"in_app_provisioning,omitempty"`
	ManualEntry *ManualEntry `json:"manual_entry,omitempty"`
	WalletProviderCardOnFile *WalletProviderCardOnFile `json:"wallet_provider_card_on_file,omitempty"`
	WebPushProvisioning *WebPushProvisioning `json:"web_push_provisioning,omitempty"`
}

// NewProvisioningControls instantiates a new ProvisioningControls object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProvisioningControls() *ProvisioningControls {
	this := ProvisioningControls{}
	return &this
}

// NewProvisioningControlsWithDefaults instantiates a new ProvisioningControls object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProvisioningControlsWithDefaults() *ProvisioningControls {
	this := ProvisioningControls{}
	return &this
}

// GetDwtUseCardStatusDuringAuth returns the DwtUseCardStatusDuringAuth field value if set, zero value otherwise.
func (o *ProvisioningControls) GetDwtUseCardStatusDuringAuth() bool {
	if o == nil || IsNil(o.DwtUseCardStatusDuringAuth) {
		var ret bool
		return ret
	}
	return *o.DwtUseCardStatusDuringAuth
}

// GetDwtUseCardStatusDuringAuthOk returns a tuple with the DwtUseCardStatusDuringAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningControls) GetDwtUseCardStatusDuringAuthOk() (*bool, bool) {
	if o == nil || IsNil(o.DwtUseCardStatusDuringAuth) {
		return nil, false
	}
	return o.DwtUseCardStatusDuringAuth, true
}

// HasDwtUseCardStatusDuringAuth returns a boolean if a field has been set.
func (o *ProvisioningControls) HasDwtUseCardStatusDuringAuth() bool {
	if o != nil && !IsNil(o.DwtUseCardStatusDuringAuth) {
		return true
	}

	return false
}

// SetDwtUseCardStatusDuringAuth gets a reference to the given bool and assigns it to the DwtUseCardStatusDuringAuth field.
func (o *ProvisioningControls) SetDwtUseCardStatusDuringAuth(v bool) {
	o.DwtUseCardStatusDuringAuth = &v
}

// GetDwtVerifyAtcDuringAuth returns the DwtVerifyAtcDuringAuth field value if set, zero value otherwise.
func (o *ProvisioningControls) GetDwtVerifyAtcDuringAuth() bool {
	if o == nil || IsNil(o.DwtVerifyAtcDuringAuth) {
		var ret bool
		return ret
	}
	return *o.DwtVerifyAtcDuringAuth
}

// GetDwtVerifyAtcDuringAuthOk returns a tuple with the DwtVerifyAtcDuringAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningControls) GetDwtVerifyAtcDuringAuthOk() (*bool, bool) {
	if o == nil || IsNil(o.DwtVerifyAtcDuringAuth) {
		return nil, false
	}
	return o.DwtVerifyAtcDuringAuth, true
}

// HasDwtVerifyAtcDuringAuth returns a boolean if a field has been set.
func (o *ProvisioningControls) HasDwtVerifyAtcDuringAuth() bool {
	if o != nil && !IsNil(o.DwtVerifyAtcDuringAuth) {
		return true
	}

	return false
}

// SetDwtVerifyAtcDuringAuth gets a reference to the given bool and assigns it to the DwtVerifyAtcDuringAuth field.
func (o *ProvisioningControls) SetDwtVerifyAtcDuringAuth(v bool) {
	o.DwtVerifyAtcDuringAuth = &v
}

// GetForceYellowPathForCardProduct returns the ForceYellowPathForCardProduct field value if set, zero value otherwise.
func (o *ProvisioningControls) GetForceYellowPathForCardProduct() bool {
	if o == nil || IsNil(o.ForceYellowPathForCardProduct) {
		var ret bool
		return ret
	}
	return *o.ForceYellowPathForCardProduct
}

// GetForceYellowPathForCardProductOk returns a tuple with the ForceYellowPathForCardProduct field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningControls) GetForceYellowPathForCardProductOk() (*bool, bool) {
	if o == nil || IsNil(o.ForceYellowPathForCardProduct) {
		return nil, false
	}
	return o.ForceYellowPathForCardProduct, true
}

// HasForceYellowPathForCardProduct returns a boolean if a field has been set.
func (o *ProvisioningControls) HasForceYellowPathForCardProduct() bool {
	if o != nil && !IsNil(o.ForceYellowPathForCardProduct) {
		return true
	}

	return false
}

// SetForceYellowPathForCardProduct gets a reference to the given bool and assigns it to the ForceYellowPathForCardProduct field.
func (o *ProvisioningControls) SetForceYellowPathForCardProduct(v bool) {
	o.ForceYellowPathForCardProduct = &v
}

// GetInAppProvisioning returns the InAppProvisioning field value if set, zero value otherwise.
func (o *ProvisioningControls) GetInAppProvisioning() InAppProvisioning {
	if o == nil || IsNil(o.InAppProvisioning) {
		var ret InAppProvisioning
		return ret
	}
	return *o.InAppProvisioning
}

// GetInAppProvisioningOk returns a tuple with the InAppProvisioning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningControls) GetInAppProvisioningOk() (*InAppProvisioning, bool) {
	if o == nil || IsNil(o.InAppProvisioning) {
		return nil, false
	}
	return o.InAppProvisioning, true
}

// HasInAppProvisioning returns a boolean if a field has been set.
func (o *ProvisioningControls) HasInAppProvisioning() bool {
	if o != nil && !IsNil(o.InAppProvisioning) {
		return true
	}

	return false
}

// SetInAppProvisioning gets a reference to the given InAppProvisioning and assigns it to the InAppProvisioning field.
func (o *ProvisioningControls) SetInAppProvisioning(v InAppProvisioning) {
	o.InAppProvisioning = &v
}

// GetManualEntry returns the ManualEntry field value if set, zero value otherwise.
func (o *ProvisioningControls) GetManualEntry() ManualEntry {
	if o == nil || IsNil(o.ManualEntry) {
		var ret ManualEntry
		return ret
	}
	return *o.ManualEntry
}

// GetManualEntryOk returns a tuple with the ManualEntry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningControls) GetManualEntryOk() (*ManualEntry, bool) {
	if o == nil || IsNil(o.ManualEntry) {
		return nil, false
	}
	return o.ManualEntry, true
}

// HasManualEntry returns a boolean if a field has been set.
func (o *ProvisioningControls) HasManualEntry() bool {
	if o != nil && !IsNil(o.ManualEntry) {
		return true
	}

	return false
}

// SetManualEntry gets a reference to the given ManualEntry and assigns it to the ManualEntry field.
func (o *ProvisioningControls) SetManualEntry(v ManualEntry) {
	o.ManualEntry = &v
}

// GetWalletProviderCardOnFile returns the WalletProviderCardOnFile field value if set, zero value otherwise.
func (o *ProvisioningControls) GetWalletProviderCardOnFile() WalletProviderCardOnFile {
	if o == nil || IsNil(o.WalletProviderCardOnFile) {
		var ret WalletProviderCardOnFile
		return ret
	}
	return *o.WalletProviderCardOnFile
}

// GetWalletProviderCardOnFileOk returns a tuple with the WalletProviderCardOnFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningControls) GetWalletProviderCardOnFileOk() (*WalletProviderCardOnFile, bool) {
	if o == nil || IsNil(o.WalletProviderCardOnFile) {
		return nil, false
	}
	return o.WalletProviderCardOnFile, true
}

// HasWalletProviderCardOnFile returns a boolean if a field has been set.
func (o *ProvisioningControls) HasWalletProviderCardOnFile() bool {
	if o != nil && !IsNil(o.WalletProviderCardOnFile) {
		return true
	}

	return false
}

// SetWalletProviderCardOnFile gets a reference to the given WalletProviderCardOnFile and assigns it to the WalletProviderCardOnFile field.
func (o *ProvisioningControls) SetWalletProviderCardOnFile(v WalletProviderCardOnFile) {
	o.WalletProviderCardOnFile = &v
}

// GetWebPushProvisioning returns the WebPushProvisioning field value if set, zero value otherwise.
func (o *ProvisioningControls) GetWebPushProvisioning() WebPushProvisioning {
	if o == nil || IsNil(o.WebPushProvisioning) {
		var ret WebPushProvisioning
		return ret
	}
	return *o.WebPushProvisioning
}

// GetWebPushProvisioningOk returns a tuple with the WebPushProvisioning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningControls) GetWebPushProvisioningOk() (*WebPushProvisioning, bool) {
	if o == nil || IsNil(o.WebPushProvisioning) {
		return nil, false
	}
	return o.WebPushProvisioning, true
}

// HasWebPushProvisioning returns a boolean if a field has been set.
func (o *ProvisioningControls) HasWebPushProvisioning() bool {
	if o != nil && !IsNil(o.WebPushProvisioning) {
		return true
	}

	return false
}

// SetWebPushProvisioning gets a reference to the given WebPushProvisioning and assigns it to the WebPushProvisioning field.
func (o *ProvisioningControls) SetWebPushProvisioning(v WebPushProvisioning) {
	o.WebPushProvisioning = &v
}

func (o ProvisioningControls) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProvisioningControls) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DwtUseCardStatusDuringAuth) {
		toSerialize["dwt_use_card_status_during_auth"] = o.DwtUseCardStatusDuringAuth
	}
	if !IsNil(o.DwtVerifyAtcDuringAuth) {
		toSerialize["dwt_verify_atc_during_auth"] = o.DwtVerifyAtcDuringAuth
	}
	if !IsNil(o.ForceYellowPathForCardProduct) {
		toSerialize["force_yellow_path_for_card_product"] = o.ForceYellowPathForCardProduct
	}
	if !IsNil(o.InAppProvisioning) {
		toSerialize["in_app_provisioning"] = o.InAppProvisioning
	}
	if !IsNil(o.ManualEntry) {
		toSerialize["manual_entry"] = o.ManualEntry
	}
	if !IsNil(o.WalletProviderCardOnFile) {
		toSerialize["wallet_provider_card_on_file"] = o.WalletProviderCardOnFile
	}
	if !IsNil(o.WebPushProvisioning) {
		toSerialize["web_push_provisioning"] = o.WebPushProvisioning
	}
	return toSerialize, nil
}

type NullableProvisioningControls struct {
	value *ProvisioningControls
	isSet bool
}

func (v NullableProvisioningControls) Get() *ProvisioningControls {
	return v.value
}

func (v *NullableProvisioningControls) Set(val *ProvisioningControls) {
	v.value = val
	v.isSet = true
}

func (v NullableProvisioningControls) IsSet() bool {
	return v.isSet
}

func (v *NullableProvisioningControls) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvisioningControls(val *ProvisioningControls) *NullableProvisioningControls {
	return &NullableProvisioningControls{value: val, isSet: true}
}

func (v NullableProvisioningControls) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvisioningControls) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


