# AccountAdjustmentReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **float32** | Amount of the adjustment.  Value must be negative if &#x60;original_ledger_entry_token&#x60; is not passed. | 
**CurrencyCode** | [**CurrencyCode**](CurrencyCode.md) |  | [default to CURRENCYCODE_USD]
**Description** | **string** | Description of the adjustment. | 
**ExternalAdjustmentId** | Pointer to **string** | Unique identifier you provide of an associated external adjustment that exists outside Marqeta&#39;s credit platform. | [optional] 
**Note** | Pointer to **string** | Additional information on the adjustment. | [optional] 
**OriginalLedgerEntryToken** | Pointer to **string** | Unique identifier of the original journal entry needing the adjustment.  Required when adjusting an existing journal entry. | [optional] 
**Reason** | Pointer to **string** | Reason for the adjustment.  * &#x60;DISPUTE&#x60; - The adjustment occurred because a dispute was initiated. * &#x60;DISPUTE_RESOLUTION&#x60; - The adjustment occurred because of the result of a dispute resolution. * &#x60;RETURNED_OR_CANCELED_PAYMENT&#x60; - The adjustment occurred because a payment was returned or canceled. * &#x60;OTHER&#x60; - Any other reason the adjustment occurred. For example, a waived fee or account write-off. | [optional] 
**Token** | Pointer to **string** | Unique identifier of the adjustment. | [optional] 

## Methods

### NewAccountAdjustmentReq

`func NewAccountAdjustmentReq(amount float32, currencyCode CurrencyCode, description string, ) *AccountAdjustmentReq`

NewAccountAdjustmentReq instantiates a new AccountAdjustmentReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountAdjustmentReqWithDefaults

`func NewAccountAdjustmentReqWithDefaults() *AccountAdjustmentReq`

NewAccountAdjustmentReqWithDefaults instantiates a new AccountAdjustmentReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *AccountAdjustmentReq) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *AccountAdjustmentReq) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *AccountAdjustmentReq) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetCurrencyCode

`func (o *AccountAdjustmentReq) GetCurrencyCode() CurrencyCode`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *AccountAdjustmentReq) GetCurrencyCodeOk() (*CurrencyCode, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *AccountAdjustmentReq) SetCurrencyCode(v CurrencyCode)`

SetCurrencyCode sets CurrencyCode field to given value.


### GetDescription

`func (o *AccountAdjustmentReq) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AccountAdjustmentReq) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AccountAdjustmentReq) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetExternalAdjustmentId

`func (o *AccountAdjustmentReq) GetExternalAdjustmentId() string`

GetExternalAdjustmentId returns the ExternalAdjustmentId field if non-nil, zero value otherwise.

### GetExternalAdjustmentIdOk

`func (o *AccountAdjustmentReq) GetExternalAdjustmentIdOk() (*string, bool)`

GetExternalAdjustmentIdOk returns a tuple with the ExternalAdjustmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalAdjustmentId

`func (o *AccountAdjustmentReq) SetExternalAdjustmentId(v string)`

SetExternalAdjustmentId sets ExternalAdjustmentId field to given value.

### HasExternalAdjustmentId

`func (o *AccountAdjustmentReq) HasExternalAdjustmentId() bool`

HasExternalAdjustmentId returns a boolean if a field has been set.

### GetNote

`func (o *AccountAdjustmentReq) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *AccountAdjustmentReq) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *AccountAdjustmentReq) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *AccountAdjustmentReq) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetOriginalLedgerEntryToken

`func (o *AccountAdjustmentReq) GetOriginalLedgerEntryToken() string`

GetOriginalLedgerEntryToken returns the OriginalLedgerEntryToken field if non-nil, zero value otherwise.

### GetOriginalLedgerEntryTokenOk

`func (o *AccountAdjustmentReq) GetOriginalLedgerEntryTokenOk() (*string, bool)`

GetOriginalLedgerEntryTokenOk returns a tuple with the OriginalLedgerEntryToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalLedgerEntryToken

`func (o *AccountAdjustmentReq) SetOriginalLedgerEntryToken(v string)`

SetOriginalLedgerEntryToken sets OriginalLedgerEntryToken field to given value.

### HasOriginalLedgerEntryToken

`func (o *AccountAdjustmentReq) HasOriginalLedgerEntryToken() bool`

HasOriginalLedgerEntryToken returns a boolean if a field has been set.

### GetReason

`func (o *AccountAdjustmentReq) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *AccountAdjustmentReq) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *AccountAdjustmentReq) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *AccountAdjustmentReq) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetToken

`func (o *AccountAdjustmentReq) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AccountAdjustmentReq) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AccountAdjustmentReq) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *AccountAdjustmentReq) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


