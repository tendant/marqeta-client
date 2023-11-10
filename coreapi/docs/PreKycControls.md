# PreKycControls

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BalanceMax** | Pointer to **float32** | Specifies the maximum ledger balance allowed for members of the account holder group. | [optional] 
**CashAccessEnabled** | Pointer to **bool** | If set to &#x60;false&#x60;, this control prohibits an account holder&#39;s cards from being used at an ATM.  *NOTE:* If a card product&#39;s &#x60;config.poi.atm&#x60; field is set to &#x60;false&#x60;, associated cards are prohibited from being used at an ATM regardless of this control&#39;s setting. | [optional] [default to false]
**EnableNonProgramLoads** | Pointer to **bool** | If set to &#x60;true&#x60;, funds can only be loaded from a program funding source.  This restriction applies to GPA orders, peer transfers, and direct deposits, but does not apply to operator adjustments. | [optional] [default to false]
**InternationalEnabled** | Pointer to **bool** | If set to &#x60;false&#x60;, this control prohibits an account holder from conducting transactions with a non-domestic country code.  *NOTE:* If a card product is configured to prohibit non-domestic transactions, its associated cards are prohibited from such transactions regardless of this control&#39;s setting. | [optional] [default to false]
**IsReloadablePreKyc** | Pointer to **bool** | If set to &#x60;false&#x60;, this control prohibits an account holder&#39;s account from being reloaded with funds after an initial load.  This restriction applies to GPA orders, peer transfers, and direct deposits, but does not apply to operator adjustments. | [optional] [default to false]

## Methods

### NewPreKycControls

`func NewPreKycControls() *PreKycControls`

NewPreKycControls instantiates a new PreKycControls object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPreKycControlsWithDefaults

`func NewPreKycControlsWithDefaults() *PreKycControls`

NewPreKycControlsWithDefaults instantiates a new PreKycControls object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalanceMax

`func (o *PreKycControls) GetBalanceMax() float32`

GetBalanceMax returns the BalanceMax field if non-nil, zero value otherwise.

### GetBalanceMaxOk

`func (o *PreKycControls) GetBalanceMaxOk() (*float32, bool)`

GetBalanceMaxOk returns a tuple with the BalanceMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceMax

`func (o *PreKycControls) SetBalanceMax(v float32)`

SetBalanceMax sets BalanceMax field to given value.

### HasBalanceMax

`func (o *PreKycControls) HasBalanceMax() bool`

HasBalanceMax returns a boolean if a field has been set.

### GetCashAccessEnabled

`func (o *PreKycControls) GetCashAccessEnabled() bool`

GetCashAccessEnabled returns the CashAccessEnabled field if non-nil, zero value otherwise.

### GetCashAccessEnabledOk

`func (o *PreKycControls) GetCashAccessEnabledOk() (*bool, bool)`

GetCashAccessEnabledOk returns a tuple with the CashAccessEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashAccessEnabled

`func (o *PreKycControls) SetCashAccessEnabled(v bool)`

SetCashAccessEnabled sets CashAccessEnabled field to given value.

### HasCashAccessEnabled

`func (o *PreKycControls) HasCashAccessEnabled() bool`

HasCashAccessEnabled returns a boolean if a field has been set.

### GetEnableNonProgramLoads

`func (o *PreKycControls) GetEnableNonProgramLoads() bool`

GetEnableNonProgramLoads returns the EnableNonProgramLoads field if non-nil, zero value otherwise.

### GetEnableNonProgramLoadsOk

`func (o *PreKycControls) GetEnableNonProgramLoadsOk() (*bool, bool)`

GetEnableNonProgramLoadsOk returns a tuple with the EnableNonProgramLoads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNonProgramLoads

`func (o *PreKycControls) SetEnableNonProgramLoads(v bool)`

SetEnableNonProgramLoads sets EnableNonProgramLoads field to given value.

### HasEnableNonProgramLoads

`func (o *PreKycControls) HasEnableNonProgramLoads() bool`

HasEnableNonProgramLoads returns a boolean if a field has been set.

### GetInternationalEnabled

`func (o *PreKycControls) GetInternationalEnabled() bool`

GetInternationalEnabled returns the InternationalEnabled field if non-nil, zero value otherwise.

### GetInternationalEnabledOk

`func (o *PreKycControls) GetInternationalEnabledOk() (*bool, bool)`

GetInternationalEnabledOk returns a tuple with the InternationalEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternationalEnabled

`func (o *PreKycControls) SetInternationalEnabled(v bool)`

SetInternationalEnabled sets InternationalEnabled field to given value.

### HasInternationalEnabled

`func (o *PreKycControls) HasInternationalEnabled() bool`

HasInternationalEnabled returns a boolean if a field has been set.

### GetIsReloadablePreKyc

`func (o *PreKycControls) GetIsReloadablePreKyc() bool`

GetIsReloadablePreKyc returns the IsReloadablePreKyc field if non-nil, zero value otherwise.

### GetIsReloadablePreKycOk

`func (o *PreKycControls) GetIsReloadablePreKycOk() (*bool, bool)`

GetIsReloadablePreKycOk returns a tuple with the IsReloadablePreKyc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReloadablePreKyc

`func (o *PreKycControls) SetIsReloadablePreKyc(v bool)`

SetIsReloadablePreKyc sets IsReloadablePreKyc field to given value.

### HasIsReloadablePreKyc

`func (o *PreKycControls) HasIsReloadablePreKyc() bool`

HasIsReloadablePreKyc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


