# AccountConfigUpdateReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EDisclosureActive** | Pointer to **bool** | A value of &#x60;true&#x60; indicates that the account holder consents to receiving disclosures and statements electronically. | [optional] [default to false]
**Fees** | Pointer to [**[]ConfigFeeScheduleReq**](ConfigFeeScheduleReq.md) | Contains one or more fees associated with the credit account. | [optional] 
**MinPayment** | Pointer to [**AccountConfigMinPayment**](AccountConfigMinPayment.md) |  | [optional] 
**PaymentHolds** | Pointer to [**AccountConfigPaymentHolds**](AccountConfigPaymentHolds.md) |  | [optional] 

## Methods

### NewAccountConfigUpdateReq

`func NewAccountConfigUpdateReq() *AccountConfigUpdateReq`

NewAccountConfigUpdateReq instantiates a new AccountConfigUpdateReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountConfigUpdateReqWithDefaults

`func NewAccountConfigUpdateReqWithDefaults() *AccountConfigUpdateReq`

NewAccountConfigUpdateReqWithDefaults instantiates a new AccountConfigUpdateReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEDisclosureActive

`func (o *AccountConfigUpdateReq) GetEDisclosureActive() bool`

GetEDisclosureActive returns the EDisclosureActive field if non-nil, zero value otherwise.

### GetEDisclosureActiveOk

`func (o *AccountConfigUpdateReq) GetEDisclosureActiveOk() (*bool, bool)`

GetEDisclosureActiveOk returns a tuple with the EDisclosureActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEDisclosureActive

`func (o *AccountConfigUpdateReq) SetEDisclosureActive(v bool)`

SetEDisclosureActive sets EDisclosureActive field to given value.

### HasEDisclosureActive

`func (o *AccountConfigUpdateReq) HasEDisclosureActive() bool`

HasEDisclosureActive returns a boolean if a field has been set.

### GetFees

`func (o *AccountConfigUpdateReq) GetFees() []ConfigFeeScheduleReq`

GetFees returns the Fees field if non-nil, zero value otherwise.

### GetFeesOk

`func (o *AccountConfigUpdateReq) GetFeesOk() (*[]ConfigFeeScheduleReq, bool)`

GetFeesOk returns a tuple with the Fees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFees

`func (o *AccountConfigUpdateReq) SetFees(v []ConfigFeeScheduleReq)`

SetFees sets Fees field to given value.

### HasFees

`func (o *AccountConfigUpdateReq) HasFees() bool`

HasFees returns a boolean if a field has been set.

### GetMinPayment

`func (o *AccountConfigUpdateReq) GetMinPayment() AccountConfigMinPayment`

GetMinPayment returns the MinPayment field if non-nil, zero value otherwise.

### GetMinPaymentOk

`func (o *AccountConfigUpdateReq) GetMinPaymentOk() (*AccountConfigMinPayment, bool)`

GetMinPaymentOk returns a tuple with the MinPayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinPayment

`func (o *AccountConfigUpdateReq) SetMinPayment(v AccountConfigMinPayment)`

SetMinPayment sets MinPayment field to given value.

### HasMinPayment

`func (o *AccountConfigUpdateReq) HasMinPayment() bool`

HasMinPayment returns a boolean if a field has been set.

### GetPaymentHolds

`func (o *AccountConfigUpdateReq) GetPaymentHolds() AccountConfigPaymentHolds`

GetPaymentHolds returns the PaymentHolds field if non-nil, zero value otherwise.

### GetPaymentHoldsOk

`func (o *AccountConfigUpdateReq) GetPaymentHoldsOk() (*AccountConfigPaymentHolds, bool)`

GetPaymentHoldsOk returns a tuple with the PaymentHolds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentHolds

`func (o *AccountConfigUpdateReq) SetPaymentHolds(v AccountConfigPaymentHolds)`

SetPaymentHolds sets PaymentHolds field to given value.

### HasPaymentHolds

`func (o *AccountConfigUpdateReq) HasPaymentHolds() bool`

HasPaymentHolds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


