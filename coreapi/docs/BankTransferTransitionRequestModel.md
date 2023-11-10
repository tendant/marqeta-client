# BankTransferTransitionRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **float32** |  | [optional] 
**BankTransferToken** | **string** | Token of the ACH transfer you want to transition. | 
**BatchNumber** | Pointer to **string** | Field required in older versions of the API, but no longer used. | [optional] 
**Channel** | **string** | Mechanism by which the transaction was initiated. | 
**EffectiveDate** | Pointer to **time.Time** |  | [optional] 
**ProgramReserveToken** | Pointer to **string** | Not currently used. | [optional] 
**Reason** | Pointer to **string** | Description of why the ACH transfer status was updated. | [optional] 
**ReturnCode** | Pointer to **string** | Standardized ACH return code for a returned transaction, generally sent by the RDFI.  Transactions can be returned for any of the reasons listed in the &lt;&lt;/developer-guides/ach-origination#_nacha_ach_return_codes, NACHA ACH return codes table&gt;&gt; of the ACH Origination Guide. | [optional] 
**ReversalAfter45Days** | Pointer to **bool** |  | [optional] 
**Status** | **string** | New state of the ACH transfer.  *NOTE:* In production environments, the only value to which you are allowed to transition an ACH transfer is &#x60;CANCELLED&#x60;. | 
**Token** | Pointer to **string** | Unique identifier of the bank transfer transition request. | [optional] 

## Methods

### NewBankTransferTransitionRequestModel

`func NewBankTransferTransitionRequestModel(bankTransferToken string, channel string, status string, ) *BankTransferTransitionRequestModel`

NewBankTransferTransitionRequestModel instantiates a new BankTransferTransitionRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankTransferTransitionRequestModelWithDefaults

`func NewBankTransferTransitionRequestModelWithDefaults() *BankTransferTransitionRequestModel`

NewBankTransferTransitionRequestModelWithDefaults instantiates a new BankTransferTransitionRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *BankTransferTransitionRequestModel) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *BankTransferTransitionRequestModel) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *BankTransferTransitionRequestModel) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *BankTransferTransitionRequestModel) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetBankTransferToken

`func (o *BankTransferTransitionRequestModel) GetBankTransferToken() string`

GetBankTransferToken returns the BankTransferToken field if non-nil, zero value otherwise.

### GetBankTransferTokenOk

`func (o *BankTransferTransitionRequestModel) GetBankTransferTokenOk() (*string, bool)`

GetBankTransferTokenOk returns a tuple with the BankTransferToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankTransferToken

`func (o *BankTransferTransitionRequestModel) SetBankTransferToken(v string)`

SetBankTransferToken sets BankTransferToken field to given value.


### GetBatchNumber

`func (o *BankTransferTransitionRequestModel) GetBatchNumber() string`

GetBatchNumber returns the BatchNumber field if non-nil, zero value otherwise.

### GetBatchNumberOk

`func (o *BankTransferTransitionRequestModel) GetBatchNumberOk() (*string, bool)`

GetBatchNumberOk returns a tuple with the BatchNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchNumber

`func (o *BankTransferTransitionRequestModel) SetBatchNumber(v string)`

SetBatchNumber sets BatchNumber field to given value.

### HasBatchNumber

`func (o *BankTransferTransitionRequestModel) HasBatchNumber() bool`

HasBatchNumber returns a boolean if a field has been set.

### GetChannel

`func (o *BankTransferTransitionRequestModel) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *BankTransferTransitionRequestModel) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *BankTransferTransitionRequestModel) SetChannel(v string)`

SetChannel sets Channel field to given value.


### GetEffectiveDate

`func (o *BankTransferTransitionRequestModel) GetEffectiveDate() time.Time`

GetEffectiveDate returns the EffectiveDate field if non-nil, zero value otherwise.

### GetEffectiveDateOk

`func (o *BankTransferTransitionRequestModel) GetEffectiveDateOk() (*time.Time, bool)`

GetEffectiveDateOk returns a tuple with the EffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveDate

`func (o *BankTransferTransitionRequestModel) SetEffectiveDate(v time.Time)`

SetEffectiveDate sets EffectiveDate field to given value.

### HasEffectiveDate

`func (o *BankTransferTransitionRequestModel) HasEffectiveDate() bool`

HasEffectiveDate returns a boolean if a field has been set.

### GetProgramReserveToken

`func (o *BankTransferTransitionRequestModel) GetProgramReserveToken() string`

GetProgramReserveToken returns the ProgramReserveToken field if non-nil, zero value otherwise.

### GetProgramReserveTokenOk

`func (o *BankTransferTransitionRequestModel) GetProgramReserveTokenOk() (*string, bool)`

GetProgramReserveTokenOk returns a tuple with the ProgramReserveToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramReserveToken

`func (o *BankTransferTransitionRequestModel) SetProgramReserveToken(v string)`

SetProgramReserveToken sets ProgramReserveToken field to given value.

### HasProgramReserveToken

`func (o *BankTransferTransitionRequestModel) HasProgramReserveToken() bool`

HasProgramReserveToken returns a boolean if a field has been set.

### GetReason

`func (o *BankTransferTransitionRequestModel) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *BankTransferTransitionRequestModel) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *BankTransferTransitionRequestModel) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *BankTransferTransitionRequestModel) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetReturnCode

`func (o *BankTransferTransitionRequestModel) GetReturnCode() string`

GetReturnCode returns the ReturnCode field if non-nil, zero value otherwise.

### GetReturnCodeOk

`func (o *BankTransferTransitionRequestModel) GetReturnCodeOk() (*string, bool)`

GetReturnCodeOk returns a tuple with the ReturnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnCode

`func (o *BankTransferTransitionRequestModel) SetReturnCode(v string)`

SetReturnCode sets ReturnCode field to given value.

### HasReturnCode

`func (o *BankTransferTransitionRequestModel) HasReturnCode() bool`

HasReturnCode returns a boolean if a field has been set.

### GetReversalAfter45Days

`func (o *BankTransferTransitionRequestModel) GetReversalAfter45Days() bool`

GetReversalAfter45Days returns the ReversalAfter45Days field if non-nil, zero value otherwise.

### GetReversalAfter45DaysOk

`func (o *BankTransferTransitionRequestModel) GetReversalAfter45DaysOk() (*bool, bool)`

GetReversalAfter45DaysOk returns a tuple with the ReversalAfter45Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReversalAfter45Days

`func (o *BankTransferTransitionRequestModel) SetReversalAfter45Days(v bool)`

SetReversalAfter45Days sets ReversalAfter45Days field to given value.

### HasReversalAfter45Days

`func (o *BankTransferTransitionRequestModel) HasReversalAfter45Days() bool`

HasReversalAfter45Days returns a boolean if a field has been set.

### GetStatus

`func (o *BankTransferTransitionRequestModel) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BankTransferTransitionRequestModel) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BankTransferTransitionRequestModel) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetToken

`func (o *BankTransferTransitionRequestModel) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *BankTransferTransitionRequestModel) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *BankTransferTransitionRequestModel) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *BankTransferTransitionRequestModel) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


