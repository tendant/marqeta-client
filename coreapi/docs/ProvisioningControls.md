# ProvisioningControls

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DwtUseCardStatusDuringAuth** | Pointer to **bool** |  | [optional] 
**DwtVerifyAtcDuringAuth** | Pointer to **bool** |  | [optional] 
**ForceYellowPathForCardProduct** | Pointer to **bool** | A value of &#x60;true&#x60; requires identity verification set-up for all digital wallets at the card product level. | [optional] 
**InAppProvisioning** | Pointer to [**InAppProvisioning**](InAppProvisioning.md) |  | [optional] 
**ManualEntry** | Pointer to [**ManualEntry**](ManualEntry.md) |  | [optional] 
**WalletProviderCardOnFile** | Pointer to [**WalletProviderCardOnFile**](WalletProviderCardOnFile.md) |  | [optional] 
**WebPushProvisioning** | Pointer to [**WebPushProvisioning**](WebPushProvisioning.md) |  | [optional] 

## Methods

### NewProvisioningControls

`func NewProvisioningControls() *ProvisioningControls`

NewProvisioningControls instantiates a new ProvisioningControls object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisioningControlsWithDefaults

`func NewProvisioningControlsWithDefaults() *ProvisioningControls`

NewProvisioningControlsWithDefaults instantiates a new ProvisioningControls object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDwtUseCardStatusDuringAuth

`func (o *ProvisioningControls) GetDwtUseCardStatusDuringAuth() bool`

GetDwtUseCardStatusDuringAuth returns the DwtUseCardStatusDuringAuth field if non-nil, zero value otherwise.

### GetDwtUseCardStatusDuringAuthOk

`func (o *ProvisioningControls) GetDwtUseCardStatusDuringAuthOk() (*bool, bool)`

GetDwtUseCardStatusDuringAuthOk returns a tuple with the DwtUseCardStatusDuringAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDwtUseCardStatusDuringAuth

`func (o *ProvisioningControls) SetDwtUseCardStatusDuringAuth(v bool)`

SetDwtUseCardStatusDuringAuth sets DwtUseCardStatusDuringAuth field to given value.

### HasDwtUseCardStatusDuringAuth

`func (o *ProvisioningControls) HasDwtUseCardStatusDuringAuth() bool`

HasDwtUseCardStatusDuringAuth returns a boolean if a field has been set.

### GetDwtVerifyAtcDuringAuth

`func (o *ProvisioningControls) GetDwtVerifyAtcDuringAuth() bool`

GetDwtVerifyAtcDuringAuth returns the DwtVerifyAtcDuringAuth field if non-nil, zero value otherwise.

### GetDwtVerifyAtcDuringAuthOk

`func (o *ProvisioningControls) GetDwtVerifyAtcDuringAuthOk() (*bool, bool)`

GetDwtVerifyAtcDuringAuthOk returns a tuple with the DwtVerifyAtcDuringAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDwtVerifyAtcDuringAuth

`func (o *ProvisioningControls) SetDwtVerifyAtcDuringAuth(v bool)`

SetDwtVerifyAtcDuringAuth sets DwtVerifyAtcDuringAuth field to given value.

### HasDwtVerifyAtcDuringAuth

`func (o *ProvisioningControls) HasDwtVerifyAtcDuringAuth() bool`

HasDwtVerifyAtcDuringAuth returns a boolean if a field has been set.

### GetForceYellowPathForCardProduct

`func (o *ProvisioningControls) GetForceYellowPathForCardProduct() bool`

GetForceYellowPathForCardProduct returns the ForceYellowPathForCardProduct field if non-nil, zero value otherwise.

### GetForceYellowPathForCardProductOk

`func (o *ProvisioningControls) GetForceYellowPathForCardProductOk() (*bool, bool)`

GetForceYellowPathForCardProductOk returns a tuple with the ForceYellowPathForCardProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceYellowPathForCardProduct

`func (o *ProvisioningControls) SetForceYellowPathForCardProduct(v bool)`

SetForceYellowPathForCardProduct sets ForceYellowPathForCardProduct field to given value.

### HasForceYellowPathForCardProduct

`func (o *ProvisioningControls) HasForceYellowPathForCardProduct() bool`

HasForceYellowPathForCardProduct returns a boolean if a field has been set.

### GetInAppProvisioning

`func (o *ProvisioningControls) GetInAppProvisioning() InAppProvisioning`

GetInAppProvisioning returns the InAppProvisioning field if non-nil, zero value otherwise.

### GetInAppProvisioningOk

`func (o *ProvisioningControls) GetInAppProvisioningOk() (*InAppProvisioning, bool)`

GetInAppProvisioningOk returns a tuple with the InAppProvisioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInAppProvisioning

`func (o *ProvisioningControls) SetInAppProvisioning(v InAppProvisioning)`

SetInAppProvisioning sets InAppProvisioning field to given value.

### HasInAppProvisioning

`func (o *ProvisioningControls) HasInAppProvisioning() bool`

HasInAppProvisioning returns a boolean if a field has been set.

### GetManualEntry

`func (o *ProvisioningControls) GetManualEntry() ManualEntry`

GetManualEntry returns the ManualEntry field if non-nil, zero value otherwise.

### GetManualEntryOk

`func (o *ProvisioningControls) GetManualEntryOk() (*ManualEntry, bool)`

GetManualEntryOk returns a tuple with the ManualEntry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualEntry

`func (o *ProvisioningControls) SetManualEntry(v ManualEntry)`

SetManualEntry sets ManualEntry field to given value.

### HasManualEntry

`func (o *ProvisioningControls) HasManualEntry() bool`

HasManualEntry returns a boolean if a field has been set.

### GetWalletProviderCardOnFile

`func (o *ProvisioningControls) GetWalletProviderCardOnFile() WalletProviderCardOnFile`

GetWalletProviderCardOnFile returns the WalletProviderCardOnFile field if non-nil, zero value otherwise.

### GetWalletProviderCardOnFileOk

`func (o *ProvisioningControls) GetWalletProviderCardOnFileOk() (*WalletProviderCardOnFile, bool)`

GetWalletProviderCardOnFileOk returns a tuple with the WalletProviderCardOnFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletProviderCardOnFile

`func (o *ProvisioningControls) SetWalletProviderCardOnFile(v WalletProviderCardOnFile)`

SetWalletProviderCardOnFile sets WalletProviderCardOnFile field to given value.

### HasWalletProviderCardOnFile

`func (o *ProvisioningControls) HasWalletProviderCardOnFile() bool`

HasWalletProviderCardOnFile returns a boolean if a field has been set.

### GetWebPushProvisioning

`func (o *ProvisioningControls) GetWebPushProvisioning() WebPushProvisioning`

GetWebPushProvisioning returns the WebPushProvisioning field if non-nil, zero value otherwise.

### GetWebPushProvisioningOk

`func (o *ProvisioningControls) GetWebPushProvisioningOk() (*WebPushProvisioning, bool)`

GetWebPushProvisioningOk returns a tuple with the WebPushProvisioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebPushProvisioning

`func (o *ProvisioningControls) SetWebPushProvisioning(v WebPushProvisioning)`

SetWebPushProvisioning sets WebPushProvisioning field to given value.

### HasWebPushProvisioning

`func (o *ProvisioningControls) HasWebPushProvisioning() bool`

HasWebPushProvisioning returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


