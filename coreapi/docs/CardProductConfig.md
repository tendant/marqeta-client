# CardProductConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardLifeCycle** | Pointer to [**CardLifeCycle**](CardLifeCycle.md) |  | [optional] 
**ClearingAndSettlement** | Pointer to [**ClearingAndSettlement**](ClearingAndSettlement.md) |  | [optional] 
**DigitalWalletTokenization** | Pointer to [**DigitalWalletTokenization**](DigitalWalletTokenization.md) |  | [optional] 
**Fulfillment** | Pointer to [**CardProductFulfillment**](CardProductFulfillment.md) |  | [optional] 
**JitFunding** | Pointer to [**JitFunding**](JitFunding.md) |  | [optional] 
**Poi** | Pointer to [**Poi**](Poi.md) |  | [optional] 
**SelectiveAuth** | Pointer to [**SelectiveAuth**](SelectiveAuth.md) |  | [optional] 
**Special** | Pointer to [**Special**](Special.md) |  | [optional] 
**TransactionControls** | Pointer to [**TransactionControls**](TransactionControls.md) |  | [optional] 

## Methods

### NewCardProductConfig

`func NewCardProductConfig() *CardProductConfig`

NewCardProductConfig instantiates a new CardProductConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardProductConfigWithDefaults

`func NewCardProductConfigWithDefaults() *CardProductConfig`

NewCardProductConfigWithDefaults instantiates a new CardProductConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardLifeCycle

`func (o *CardProductConfig) GetCardLifeCycle() CardLifeCycle`

GetCardLifeCycle returns the CardLifeCycle field if non-nil, zero value otherwise.

### GetCardLifeCycleOk

`func (o *CardProductConfig) GetCardLifeCycleOk() (*CardLifeCycle, bool)`

GetCardLifeCycleOk returns a tuple with the CardLifeCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardLifeCycle

`func (o *CardProductConfig) SetCardLifeCycle(v CardLifeCycle)`

SetCardLifeCycle sets CardLifeCycle field to given value.

### HasCardLifeCycle

`func (o *CardProductConfig) HasCardLifeCycle() bool`

HasCardLifeCycle returns a boolean if a field has been set.

### GetClearingAndSettlement

`func (o *CardProductConfig) GetClearingAndSettlement() ClearingAndSettlement`

GetClearingAndSettlement returns the ClearingAndSettlement field if non-nil, zero value otherwise.

### GetClearingAndSettlementOk

`func (o *CardProductConfig) GetClearingAndSettlementOk() (*ClearingAndSettlement, bool)`

GetClearingAndSettlementOk returns a tuple with the ClearingAndSettlement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClearingAndSettlement

`func (o *CardProductConfig) SetClearingAndSettlement(v ClearingAndSettlement)`

SetClearingAndSettlement sets ClearingAndSettlement field to given value.

### HasClearingAndSettlement

`func (o *CardProductConfig) HasClearingAndSettlement() bool`

HasClearingAndSettlement returns a boolean if a field has been set.

### GetDigitalWalletTokenization

`func (o *CardProductConfig) GetDigitalWalletTokenization() DigitalWalletTokenization`

GetDigitalWalletTokenization returns the DigitalWalletTokenization field if non-nil, zero value otherwise.

### GetDigitalWalletTokenizationOk

`func (o *CardProductConfig) GetDigitalWalletTokenizationOk() (*DigitalWalletTokenization, bool)`

GetDigitalWalletTokenizationOk returns a tuple with the DigitalWalletTokenization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigitalWalletTokenization

`func (o *CardProductConfig) SetDigitalWalletTokenization(v DigitalWalletTokenization)`

SetDigitalWalletTokenization sets DigitalWalletTokenization field to given value.

### HasDigitalWalletTokenization

`func (o *CardProductConfig) HasDigitalWalletTokenization() bool`

HasDigitalWalletTokenization returns a boolean if a field has been set.

### GetFulfillment

`func (o *CardProductConfig) GetFulfillment() CardProductFulfillment`

GetFulfillment returns the Fulfillment field if non-nil, zero value otherwise.

### GetFulfillmentOk

`func (o *CardProductConfig) GetFulfillmentOk() (*CardProductFulfillment, bool)`

GetFulfillmentOk returns a tuple with the Fulfillment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillment

`func (o *CardProductConfig) SetFulfillment(v CardProductFulfillment)`

SetFulfillment sets Fulfillment field to given value.

### HasFulfillment

`func (o *CardProductConfig) HasFulfillment() bool`

HasFulfillment returns a boolean if a field has been set.

### GetJitFunding

`func (o *CardProductConfig) GetJitFunding() JitFunding`

GetJitFunding returns the JitFunding field if non-nil, zero value otherwise.

### GetJitFundingOk

`func (o *CardProductConfig) GetJitFundingOk() (*JitFunding, bool)`

GetJitFundingOk returns a tuple with the JitFunding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJitFunding

`func (o *CardProductConfig) SetJitFunding(v JitFunding)`

SetJitFunding sets JitFunding field to given value.

### HasJitFunding

`func (o *CardProductConfig) HasJitFunding() bool`

HasJitFunding returns a boolean if a field has been set.

### GetPoi

`func (o *CardProductConfig) GetPoi() Poi`

GetPoi returns the Poi field if non-nil, zero value otherwise.

### GetPoiOk

`func (o *CardProductConfig) GetPoiOk() (*Poi, bool)`

GetPoiOk returns a tuple with the Poi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoi

`func (o *CardProductConfig) SetPoi(v Poi)`

SetPoi sets Poi field to given value.

### HasPoi

`func (o *CardProductConfig) HasPoi() bool`

HasPoi returns a boolean if a field has been set.

### GetSelectiveAuth

`func (o *CardProductConfig) GetSelectiveAuth() SelectiveAuth`

GetSelectiveAuth returns the SelectiveAuth field if non-nil, zero value otherwise.

### GetSelectiveAuthOk

`func (o *CardProductConfig) GetSelectiveAuthOk() (*SelectiveAuth, bool)`

GetSelectiveAuthOk returns a tuple with the SelectiveAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectiveAuth

`func (o *CardProductConfig) SetSelectiveAuth(v SelectiveAuth)`

SetSelectiveAuth sets SelectiveAuth field to given value.

### HasSelectiveAuth

`func (o *CardProductConfig) HasSelectiveAuth() bool`

HasSelectiveAuth returns a boolean if a field has been set.

### GetSpecial

`func (o *CardProductConfig) GetSpecial() Special`

GetSpecial returns the Special field if non-nil, zero value otherwise.

### GetSpecialOk

`func (o *CardProductConfig) GetSpecialOk() (*Special, bool)`

GetSpecialOk returns a tuple with the Special field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecial

`func (o *CardProductConfig) SetSpecial(v Special)`

SetSpecial sets Special field to given value.

### HasSpecial

`func (o *CardProductConfig) HasSpecial() bool`

HasSpecial returns a boolean if a field has been set.

### GetTransactionControls

`func (o *CardProductConfig) GetTransactionControls() TransactionControls`

GetTransactionControls returns the TransactionControls field if non-nil, zero value otherwise.

### GetTransactionControlsOk

`func (o *CardProductConfig) GetTransactionControlsOk() (*TransactionControls, bool)`

GetTransactionControlsOk returns a tuple with the TransactionControls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionControls

`func (o *CardProductConfig) SetTransactionControls(v TransactionControls)`

SetTransactionControls sets TransactionControls field to given value.

### HasTransactionControls

`func (o *CardProductConfig) HasTransactionControls() bool`

HasTransactionControls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


