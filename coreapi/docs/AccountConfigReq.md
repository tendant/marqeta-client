# AccountConfigReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingCycleDay** | Pointer to **int32** | Day of month the billing cycle starts. | [optional] 
**CardLevel** | Pointer to **string** | The level of the credit card. | [optional] [default to "NA"]
**EDisclosureActive** | Pointer to **bool** | A value of &#x60;true&#x60; indicates that the account holder consents to receiving disclosures and statements electronically. | [optional] [default to false]
**Fees** | Pointer to [**[]ConfigFeeScheduleReq**](ConfigFeeScheduleReq.md) | Contains one or more fees associated with the credit account. | [optional] 
**PaymentDueDay** | Pointer to **int32** | Day of month the payment for the previous billing cycle is due. | [optional] 
**PaymentHolds** | Pointer to [**AccountConfigPaymentHolds**](AccountConfigPaymentHolds.md) |  | [optional] 
**Rewards** | Pointer to [**[]AccountReward**](AccountReward.md) | Contains one or more rewards associated with the credit account. | [optional] 

## Methods

### NewAccountConfigReq

`func NewAccountConfigReq() *AccountConfigReq`

NewAccountConfigReq instantiates a new AccountConfigReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountConfigReqWithDefaults

`func NewAccountConfigReqWithDefaults() *AccountConfigReq`

NewAccountConfigReqWithDefaults instantiates a new AccountConfigReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingCycleDay

`func (o *AccountConfigReq) GetBillingCycleDay() int32`

GetBillingCycleDay returns the BillingCycleDay field if non-nil, zero value otherwise.

### GetBillingCycleDayOk

`func (o *AccountConfigReq) GetBillingCycleDayOk() (*int32, bool)`

GetBillingCycleDayOk returns a tuple with the BillingCycleDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCycleDay

`func (o *AccountConfigReq) SetBillingCycleDay(v int32)`

SetBillingCycleDay sets BillingCycleDay field to given value.

### HasBillingCycleDay

`func (o *AccountConfigReq) HasBillingCycleDay() bool`

HasBillingCycleDay returns a boolean if a field has been set.

### GetCardLevel

`func (o *AccountConfigReq) GetCardLevel() string`

GetCardLevel returns the CardLevel field if non-nil, zero value otherwise.

### GetCardLevelOk

`func (o *AccountConfigReq) GetCardLevelOk() (*string, bool)`

GetCardLevelOk returns a tuple with the CardLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardLevel

`func (o *AccountConfigReq) SetCardLevel(v string)`

SetCardLevel sets CardLevel field to given value.

### HasCardLevel

`func (o *AccountConfigReq) HasCardLevel() bool`

HasCardLevel returns a boolean if a field has been set.

### GetEDisclosureActive

`func (o *AccountConfigReq) GetEDisclosureActive() bool`

GetEDisclosureActive returns the EDisclosureActive field if non-nil, zero value otherwise.

### GetEDisclosureActiveOk

`func (o *AccountConfigReq) GetEDisclosureActiveOk() (*bool, bool)`

GetEDisclosureActiveOk returns a tuple with the EDisclosureActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEDisclosureActive

`func (o *AccountConfigReq) SetEDisclosureActive(v bool)`

SetEDisclosureActive sets EDisclosureActive field to given value.

### HasEDisclosureActive

`func (o *AccountConfigReq) HasEDisclosureActive() bool`

HasEDisclosureActive returns a boolean if a field has been set.

### GetFees

`func (o *AccountConfigReq) GetFees() []ConfigFeeScheduleReq`

GetFees returns the Fees field if non-nil, zero value otherwise.

### GetFeesOk

`func (o *AccountConfigReq) GetFeesOk() (*[]ConfigFeeScheduleReq, bool)`

GetFeesOk returns a tuple with the Fees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFees

`func (o *AccountConfigReq) SetFees(v []ConfigFeeScheduleReq)`

SetFees sets Fees field to given value.

### HasFees

`func (o *AccountConfigReq) HasFees() bool`

HasFees returns a boolean if a field has been set.

### GetPaymentDueDay

`func (o *AccountConfigReq) GetPaymentDueDay() int32`

GetPaymentDueDay returns the PaymentDueDay field if non-nil, zero value otherwise.

### GetPaymentDueDayOk

`func (o *AccountConfigReq) GetPaymentDueDayOk() (*int32, bool)`

GetPaymentDueDayOk returns a tuple with the PaymentDueDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentDueDay

`func (o *AccountConfigReq) SetPaymentDueDay(v int32)`

SetPaymentDueDay sets PaymentDueDay field to given value.

### HasPaymentDueDay

`func (o *AccountConfigReq) HasPaymentDueDay() bool`

HasPaymentDueDay returns a boolean if a field has been set.

### GetPaymentHolds

`func (o *AccountConfigReq) GetPaymentHolds() AccountConfigPaymentHolds`

GetPaymentHolds returns the PaymentHolds field if non-nil, zero value otherwise.

### GetPaymentHoldsOk

`func (o *AccountConfigReq) GetPaymentHoldsOk() (*AccountConfigPaymentHolds, bool)`

GetPaymentHoldsOk returns a tuple with the PaymentHolds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentHolds

`func (o *AccountConfigReq) SetPaymentHolds(v AccountConfigPaymentHolds)`

SetPaymentHolds sets PaymentHolds field to given value.

### HasPaymentHolds

`func (o *AccountConfigReq) HasPaymentHolds() bool`

HasPaymentHolds returns a boolean if a field has been set.

### GetRewards

`func (o *AccountConfigReq) GetRewards() []AccountReward`

GetRewards returns the Rewards field if non-nil, zero value otherwise.

### GetRewardsOk

`func (o *AccountConfigReq) GetRewardsOk() (*[]AccountReward, bool)`

GetRewardsOk returns a tuple with the Rewards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewards

`func (o *AccountConfigReq) SetRewards(v []AccountReward)`

SetRewards sets Rewards field to given value.

### HasRewards

`func (o *AccountConfigReq) HasRewards() bool`

HasRewards returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


