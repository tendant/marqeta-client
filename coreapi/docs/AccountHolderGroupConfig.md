# AccountHolderGroupConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsReloadable** | Pointer to **bool** | If set to &#x60;false&#x60;, this control prohibits an account holder&#39;s account from being reloaded with funds after the initial load.  This restriction applies to GPA orders, peer transfers, and direct deposits, but does not apply to operator adjustments. | [optional] [default to false]
**KycRequired** | Pointer to **string** | If set to &#x60;ALWAYS&#x60;, new account holders are created in an &#x60;UNVERIFIED&#x60; status and must pass identity verification (KYC) before they can be active; if set to &#x60;CONDITIONAL&#x60;, new account holders begin in a &#x60;LIMITED&#x60; status and have limited actions available before passing identity verification; if set to &#x60;NEVER&#x60;, new account holders are created in an active state. | [optional] 
**PreKycControls** | Pointer to [**PreKycControls**](PreKycControls.md) |  | [optional] 
**RealTimeFeeGroupToken** | Pointer to **string** | Associates the specified real-time fee group with the members of the account holder group. | [optional] 

## Methods

### NewAccountHolderGroupConfig

`func NewAccountHolderGroupConfig() *AccountHolderGroupConfig`

NewAccountHolderGroupConfig instantiates a new AccountHolderGroupConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountHolderGroupConfigWithDefaults

`func NewAccountHolderGroupConfigWithDefaults() *AccountHolderGroupConfig`

NewAccountHolderGroupConfigWithDefaults instantiates a new AccountHolderGroupConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsReloadable

`func (o *AccountHolderGroupConfig) GetIsReloadable() bool`

GetIsReloadable returns the IsReloadable field if non-nil, zero value otherwise.

### GetIsReloadableOk

`func (o *AccountHolderGroupConfig) GetIsReloadableOk() (*bool, bool)`

GetIsReloadableOk returns a tuple with the IsReloadable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReloadable

`func (o *AccountHolderGroupConfig) SetIsReloadable(v bool)`

SetIsReloadable sets IsReloadable field to given value.

### HasIsReloadable

`func (o *AccountHolderGroupConfig) HasIsReloadable() bool`

HasIsReloadable returns a boolean if a field has been set.

### GetKycRequired

`func (o *AccountHolderGroupConfig) GetKycRequired() string`

GetKycRequired returns the KycRequired field if non-nil, zero value otherwise.

### GetKycRequiredOk

`func (o *AccountHolderGroupConfig) GetKycRequiredOk() (*string, bool)`

GetKycRequiredOk returns a tuple with the KycRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKycRequired

`func (o *AccountHolderGroupConfig) SetKycRequired(v string)`

SetKycRequired sets KycRequired field to given value.

### HasKycRequired

`func (o *AccountHolderGroupConfig) HasKycRequired() bool`

HasKycRequired returns a boolean if a field has been set.

### GetPreKycControls

`func (o *AccountHolderGroupConfig) GetPreKycControls() PreKycControls`

GetPreKycControls returns the PreKycControls field if non-nil, zero value otherwise.

### GetPreKycControlsOk

`func (o *AccountHolderGroupConfig) GetPreKycControlsOk() (*PreKycControls, bool)`

GetPreKycControlsOk returns a tuple with the PreKycControls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreKycControls

`func (o *AccountHolderGroupConfig) SetPreKycControls(v PreKycControls)`

SetPreKycControls sets PreKycControls field to given value.

### HasPreKycControls

`func (o *AccountHolderGroupConfig) HasPreKycControls() bool`

HasPreKycControls returns a boolean if a field has been set.

### GetRealTimeFeeGroupToken

`func (o *AccountHolderGroupConfig) GetRealTimeFeeGroupToken() string`

GetRealTimeFeeGroupToken returns the RealTimeFeeGroupToken field if non-nil, zero value otherwise.

### GetRealTimeFeeGroupTokenOk

`func (o *AccountHolderGroupConfig) GetRealTimeFeeGroupTokenOk() (*string, bool)`

GetRealTimeFeeGroupTokenOk returns a tuple with the RealTimeFeeGroupToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealTimeFeeGroupToken

`func (o *AccountHolderGroupConfig) SetRealTimeFeeGroupToken(v string)`

SetRealTimeFeeGroupToken sets RealTimeFeeGroupToken field to given value.

### HasRealTimeFeeGroupToken

`func (o *AccountHolderGroupConfig) HasRealTimeFeeGroupToken() bool`

HasRealTimeFeeGroupToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


