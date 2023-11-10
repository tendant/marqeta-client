# CardTransactionSimulation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | Pointer to **string** |  | [optional] [readonly] 
**CashBackAmount** | Pointer to **float32** | The cashback amount requested. | [optional] 
**IsPreauthorization** | Pointer to **bool** | Indicates if the transaction is a pre-authorization. Set to &#x60;true&#x60; to mark the amount as an authorization only. | [optional] [default to false]
**ForcePost** | Pointer to **bool** | Set to &#x60;true&#x60; to simulate a forced clearing transaction. *NOTE:* If you set this field to &#x60;true&#x60;, you must also set the &#x60;original_transaction_token&#x60; field to an existing transaction token (the transaction does not need to be related). | [optional] [default to false]
**Fees** | Pointer to [**[]NetworkFeeModel**](NetworkFeeModel.md) | List of fees associated with the transaction. This array is returned if it exists in the resource. | [optional] 
**CardOptions** | Pointer to [**CardOptions**](CardOptions.md) |  | [optional] 
**Pos** | Pointer to [**Pos**](Pos.md) |  | [optional] 
**Network** | Pointer to **string** | Indicates which card network was used to complete the transaction. | [optional] 
**SubNetwork** | Pointer to **string** | Indicates which subnetwork used to complete the transaction. | [optional] 
**CurrencyConversion** | Pointer to [**CurrencyConversion**](CurrencyConversion.md) |  | [optional] 
**CurrencyCode** | Pointer to **string** | Currency type of the transaction. | [optional] 
**AccountFunding** | Pointer to [**AccountFundingRequest**](AccountFundingRequest.md) |  | [optional] 
**OriginalCredit** | Pointer to [**OriginalCredit**](OriginalCredit.md) |  | [optional] 

## Methods

### NewCardTransactionSimulation

`func NewCardTransactionSimulation() *CardTransactionSimulation`

NewCardTransactionSimulation instantiates a new CardTransactionSimulation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardTransactionSimulationWithDefaults

`func NewCardTransactionSimulationWithDefaults() *CardTransactionSimulation`

NewCardTransactionSimulationWithDefaults instantiates a new CardTransactionSimulation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventType

`func (o *CardTransactionSimulation) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *CardTransactionSimulation) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *CardTransactionSimulation) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *CardTransactionSimulation) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetCashBackAmount

`func (o *CardTransactionSimulation) GetCashBackAmount() float32`

GetCashBackAmount returns the CashBackAmount field if non-nil, zero value otherwise.

### GetCashBackAmountOk

`func (o *CardTransactionSimulation) GetCashBackAmountOk() (*float32, bool)`

GetCashBackAmountOk returns a tuple with the CashBackAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashBackAmount

`func (o *CardTransactionSimulation) SetCashBackAmount(v float32)`

SetCashBackAmount sets CashBackAmount field to given value.

### HasCashBackAmount

`func (o *CardTransactionSimulation) HasCashBackAmount() bool`

HasCashBackAmount returns a boolean if a field has been set.

### GetIsPreauthorization

`func (o *CardTransactionSimulation) GetIsPreauthorization() bool`

GetIsPreauthorization returns the IsPreauthorization field if non-nil, zero value otherwise.

### GetIsPreauthorizationOk

`func (o *CardTransactionSimulation) GetIsPreauthorizationOk() (*bool, bool)`

GetIsPreauthorizationOk returns a tuple with the IsPreauthorization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPreauthorization

`func (o *CardTransactionSimulation) SetIsPreauthorization(v bool)`

SetIsPreauthorization sets IsPreauthorization field to given value.

### HasIsPreauthorization

`func (o *CardTransactionSimulation) HasIsPreauthorization() bool`

HasIsPreauthorization returns a boolean if a field has been set.

### GetForcePost

`func (o *CardTransactionSimulation) GetForcePost() bool`

GetForcePost returns the ForcePost field if non-nil, zero value otherwise.

### GetForcePostOk

`func (o *CardTransactionSimulation) GetForcePostOk() (*bool, bool)`

GetForcePostOk returns a tuple with the ForcePost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcePost

`func (o *CardTransactionSimulation) SetForcePost(v bool)`

SetForcePost sets ForcePost field to given value.

### HasForcePost

`func (o *CardTransactionSimulation) HasForcePost() bool`

HasForcePost returns a boolean if a field has been set.

### GetFees

`func (o *CardTransactionSimulation) GetFees() []NetworkFeeModel`

GetFees returns the Fees field if non-nil, zero value otherwise.

### GetFeesOk

`func (o *CardTransactionSimulation) GetFeesOk() (*[]NetworkFeeModel, bool)`

GetFeesOk returns a tuple with the Fees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFees

`func (o *CardTransactionSimulation) SetFees(v []NetworkFeeModel)`

SetFees sets Fees field to given value.

### HasFees

`func (o *CardTransactionSimulation) HasFees() bool`

HasFees returns a boolean if a field has been set.

### GetCardOptions

`func (o *CardTransactionSimulation) GetCardOptions() CardOptions`

GetCardOptions returns the CardOptions field if non-nil, zero value otherwise.

### GetCardOptionsOk

`func (o *CardTransactionSimulation) GetCardOptionsOk() (*CardOptions, bool)`

GetCardOptionsOk returns a tuple with the CardOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardOptions

`func (o *CardTransactionSimulation) SetCardOptions(v CardOptions)`

SetCardOptions sets CardOptions field to given value.

### HasCardOptions

`func (o *CardTransactionSimulation) HasCardOptions() bool`

HasCardOptions returns a boolean if a field has been set.

### GetPos

`func (o *CardTransactionSimulation) GetPos() Pos`

GetPos returns the Pos field if non-nil, zero value otherwise.

### GetPosOk

`func (o *CardTransactionSimulation) GetPosOk() (*Pos, bool)`

GetPosOk returns a tuple with the Pos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPos

`func (o *CardTransactionSimulation) SetPos(v Pos)`

SetPos sets Pos field to given value.

### HasPos

`func (o *CardTransactionSimulation) HasPos() bool`

HasPos returns a boolean if a field has been set.

### GetNetwork

`func (o *CardTransactionSimulation) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *CardTransactionSimulation) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *CardTransactionSimulation) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *CardTransactionSimulation) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetSubNetwork

`func (o *CardTransactionSimulation) GetSubNetwork() string`

GetSubNetwork returns the SubNetwork field if non-nil, zero value otherwise.

### GetSubNetworkOk

`func (o *CardTransactionSimulation) GetSubNetworkOk() (*string, bool)`

GetSubNetworkOk returns a tuple with the SubNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubNetwork

`func (o *CardTransactionSimulation) SetSubNetwork(v string)`

SetSubNetwork sets SubNetwork field to given value.

### HasSubNetwork

`func (o *CardTransactionSimulation) HasSubNetwork() bool`

HasSubNetwork returns a boolean if a field has been set.

### GetCurrencyConversion

`func (o *CardTransactionSimulation) GetCurrencyConversion() CurrencyConversion`

GetCurrencyConversion returns the CurrencyConversion field if non-nil, zero value otherwise.

### GetCurrencyConversionOk

`func (o *CardTransactionSimulation) GetCurrencyConversionOk() (*CurrencyConversion, bool)`

GetCurrencyConversionOk returns a tuple with the CurrencyConversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyConversion

`func (o *CardTransactionSimulation) SetCurrencyConversion(v CurrencyConversion)`

SetCurrencyConversion sets CurrencyConversion field to given value.

### HasCurrencyConversion

`func (o *CardTransactionSimulation) HasCurrencyConversion() bool`

HasCurrencyConversion returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *CardTransactionSimulation) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *CardTransactionSimulation) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *CardTransactionSimulation) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *CardTransactionSimulation) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetAccountFunding

`func (o *CardTransactionSimulation) GetAccountFunding() AccountFundingRequest`

GetAccountFunding returns the AccountFunding field if non-nil, zero value otherwise.

### GetAccountFundingOk

`func (o *CardTransactionSimulation) GetAccountFundingOk() (*AccountFundingRequest, bool)`

GetAccountFundingOk returns a tuple with the AccountFunding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountFunding

`func (o *CardTransactionSimulation) SetAccountFunding(v AccountFundingRequest)`

SetAccountFunding sets AccountFunding field to given value.

### HasAccountFunding

`func (o *CardTransactionSimulation) HasAccountFunding() bool`

HasAccountFunding returns a boolean if a field has been set.

### GetOriginalCredit

`func (o *CardTransactionSimulation) GetOriginalCredit() OriginalCredit`

GetOriginalCredit returns the OriginalCredit field if non-nil, zero value otherwise.

### GetOriginalCreditOk

`func (o *CardTransactionSimulation) GetOriginalCreditOk() (*OriginalCredit, bool)`

GetOriginalCreditOk returns a tuple with the OriginalCredit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalCredit

`func (o *CardTransactionSimulation) SetOriginalCredit(v OriginalCredit)`

SetOriginalCredit sets OriginalCredit field to given value.

### HasOriginalCredit

`func (o *CardTransactionSimulation) HasOriginalCredit() bool`

HasOriginalCredit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


