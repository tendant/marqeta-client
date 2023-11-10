# BankTransferTransitionResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **float32** |  | [optional] 
**BankTransferToken** | **string** | Unique identifier of the ACH transfer being transitioned. | 
**BatchNumber** | Pointer to **string** | Field required in older versions of the API, but no longer used. | [optional] 
**Channel** | **string** | Mechanism by which the transaction was initiated. | 
**CreatedTime** | Pointer to **time.Time** | Date and time when the ACH transfer was created, in UTC. | [optional] 
**EffectiveDate** | Pointer to **time.Time** |  | [optional] 
**LastModifiedTime** | Pointer to **time.Time** | Date and time when the ACH transfer was last modified. | [optional] 
**ProgramReserveToken** | Pointer to **string** | Not currently used. | [optional] 
**Reason** | Pointer to **string** | Explanation of why the transition is being made, for example \&quot;Created by POST call on API\&quot;. Returned if this information was supplied when the transition was originally created. | [optional] 
**ReturnCode** | Pointer to **string** | Standardized ACH return code for a returned transaction, generally sent by the RDFI.  Transactions can be returned for any of the reasons listed in the &lt;&lt;/developer-guides/ach-origination#_nacha_ach_return_codes, NACHA ACH return codes table&gt;&gt; of the ACH Origination Guide. | [optional] 
**ReturnReason** | Pointer to **string** | Human-readable description correlating to the &#x60;return_code&#x60;, such as &#x60;Insufficient funds&#x60;, if a return code is present in the response. | [optional] 
**ReversalAfter45Days** | Pointer to **bool** |  | [optional] 
**Status** | **string** | New state of the ACH transfer. | 
**Token** | Pointer to **string** | Unique identifier of the bank transfer transition. | [optional] 
**TransactionJitToken** | Pointer to **string** | Transaction token for JIT-related ledger entries for the ACH transfer. | [optional] 
**TransactionToken** | Pointer to **string** | Transaction token for *non*-JIT-related ledger entries for the ACH transfer. | [optional] 

## Methods

### NewBankTransferTransitionResponseModel

`func NewBankTransferTransitionResponseModel(bankTransferToken string, channel string, status string, ) *BankTransferTransitionResponseModel`

NewBankTransferTransitionResponseModel instantiates a new BankTransferTransitionResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankTransferTransitionResponseModelWithDefaults

`func NewBankTransferTransitionResponseModelWithDefaults() *BankTransferTransitionResponseModel`

NewBankTransferTransitionResponseModelWithDefaults instantiates a new BankTransferTransitionResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *BankTransferTransitionResponseModel) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *BankTransferTransitionResponseModel) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *BankTransferTransitionResponseModel) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *BankTransferTransitionResponseModel) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetBankTransferToken

`func (o *BankTransferTransitionResponseModel) GetBankTransferToken() string`

GetBankTransferToken returns the BankTransferToken field if non-nil, zero value otherwise.

### GetBankTransferTokenOk

`func (o *BankTransferTransitionResponseModel) GetBankTransferTokenOk() (*string, bool)`

GetBankTransferTokenOk returns a tuple with the BankTransferToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankTransferToken

`func (o *BankTransferTransitionResponseModel) SetBankTransferToken(v string)`

SetBankTransferToken sets BankTransferToken field to given value.


### GetBatchNumber

`func (o *BankTransferTransitionResponseModel) GetBatchNumber() string`

GetBatchNumber returns the BatchNumber field if non-nil, zero value otherwise.

### GetBatchNumberOk

`func (o *BankTransferTransitionResponseModel) GetBatchNumberOk() (*string, bool)`

GetBatchNumberOk returns a tuple with the BatchNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchNumber

`func (o *BankTransferTransitionResponseModel) SetBatchNumber(v string)`

SetBatchNumber sets BatchNumber field to given value.

### HasBatchNumber

`func (o *BankTransferTransitionResponseModel) HasBatchNumber() bool`

HasBatchNumber returns a boolean if a field has been set.

### GetChannel

`func (o *BankTransferTransitionResponseModel) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *BankTransferTransitionResponseModel) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *BankTransferTransitionResponseModel) SetChannel(v string)`

SetChannel sets Channel field to given value.


### GetCreatedTime

`func (o *BankTransferTransitionResponseModel) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *BankTransferTransitionResponseModel) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *BankTransferTransitionResponseModel) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *BankTransferTransitionResponseModel) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetEffectiveDate

`func (o *BankTransferTransitionResponseModel) GetEffectiveDate() time.Time`

GetEffectiveDate returns the EffectiveDate field if non-nil, zero value otherwise.

### GetEffectiveDateOk

`func (o *BankTransferTransitionResponseModel) GetEffectiveDateOk() (*time.Time, bool)`

GetEffectiveDateOk returns a tuple with the EffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveDate

`func (o *BankTransferTransitionResponseModel) SetEffectiveDate(v time.Time)`

SetEffectiveDate sets EffectiveDate field to given value.

### HasEffectiveDate

`func (o *BankTransferTransitionResponseModel) HasEffectiveDate() bool`

HasEffectiveDate returns a boolean if a field has been set.

### GetLastModifiedTime

`func (o *BankTransferTransitionResponseModel) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *BankTransferTransitionResponseModel) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *BankTransferTransitionResponseModel) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.

### HasLastModifiedTime

`func (o *BankTransferTransitionResponseModel) HasLastModifiedTime() bool`

HasLastModifiedTime returns a boolean if a field has been set.

### GetProgramReserveToken

`func (o *BankTransferTransitionResponseModel) GetProgramReserveToken() string`

GetProgramReserveToken returns the ProgramReserveToken field if non-nil, zero value otherwise.

### GetProgramReserveTokenOk

`func (o *BankTransferTransitionResponseModel) GetProgramReserveTokenOk() (*string, bool)`

GetProgramReserveTokenOk returns a tuple with the ProgramReserveToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramReserveToken

`func (o *BankTransferTransitionResponseModel) SetProgramReserveToken(v string)`

SetProgramReserveToken sets ProgramReserveToken field to given value.

### HasProgramReserveToken

`func (o *BankTransferTransitionResponseModel) HasProgramReserveToken() bool`

HasProgramReserveToken returns a boolean if a field has been set.

### GetReason

`func (o *BankTransferTransitionResponseModel) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *BankTransferTransitionResponseModel) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *BankTransferTransitionResponseModel) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *BankTransferTransitionResponseModel) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetReturnCode

`func (o *BankTransferTransitionResponseModel) GetReturnCode() string`

GetReturnCode returns the ReturnCode field if non-nil, zero value otherwise.

### GetReturnCodeOk

`func (o *BankTransferTransitionResponseModel) GetReturnCodeOk() (*string, bool)`

GetReturnCodeOk returns a tuple with the ReturnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnCode

`func (o *BankTransferTransitionResponseModel) SetReturnCode(v string)`

SetReturnCode sets ReturnCode field to given value.

### HasReturnCode

`func (o *BankTransferTransitionResponseModel) HasReturnCode() bool`

HasReturnCode returns a boolean if a field has been set.

### GetReturnReason

`func (o *BankTransferTransitionResponseModel) GetReturnReason() string`

GetReturnReason returns the ReturnReason field if non-nil, zero value otherwise.

### GetReturnReasonOk

`func (o *BankTransferTransitionResponseModel) GetReturnReasonOk() (*string, bool)`

GetReturnReasonOk returns a tuple with the ReturnReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnReason

`func (o *BankTransferTransitionResponseModel) SetReturnReason(v string)`

SetReturnReason sets ReturnReason field to given value.

### HasReturnReason

`func (o *BankTransferTransitionResponseModel) HasReturnReason() bool`

HasReturnReason returns a boolean if a field has been set.

### GetReversalAfter45Days

`func (o *BankTransferTransitionResponseModel) GetReversalAfter45Days() bool`

GetReversalAfter45Days returns the ReversalAfter45Days field if non-nil, zero value otherwise.

### GetReversalAfter45DaysOk

`func (o *BankTransferTransitionResponseModel) GetReversalAfter45DaysOk() (*bool, bool)`

GetReversalAfter45DaysOk returns a tuple with the ReversalAfter45Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReversalAfter45Days

`func (o *BankTransferTransitionResponseModel) SetReversalAfter45Days(v bool)`

SetReversalAfter45Days sets ReversalAfter45Days field to given value.

### HasReversalAfter45Days

`func (o *BankTransferTransitionResponseModel) HasReversalAfter45Days() bool`

HasReversalAfter45Days returns a boolean if a field has been set.

### GetStatus

`func (o *BankTransferTransitionResponseModel) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BankTransferTransitionResponseModel) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BankTransferTransitionResponseModel) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetToken

`func (o *BankTransferTransitionResponseModel) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *BankTransferTransitionResponseModel) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *BankTransferTransitionResponseModel) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *BankTransferTransitionResponseModel) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTransactionJitToken

`func (o *BankTransferTransitionResponseModel) GetTransactionJitToken() string`

GetTransactionJitToken returns the TransactionJitToken field if non-nil, zero value otherwise.

### GetTransactionJitTokenOk

`func (o *BankTransferTransitionResponseModel) GetTransactionJitTokenOk() (*string, bool)`

GetTransactionJitTokenOk returns a tuple with the TransactionJitToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionJitToken

`func (o *BankTransferTransitionResponseModel) SetTransactionJitToken(v string)`

SetTransactionJitToken sets TransactionJitToken field to given value.

### HasTransactionJitToken

`func (o *BankTransferTransitionResponseModel) HasTransactionJitToken() bool`

HasTransactionJitToken returns a boolean if a field has been set.

### GetTransactionToken

`func (o *BankTransferTransitionResponseModel) GetTransactionToken() string`

GetTransactionToken returns the TransactionToken field if non-nil, zero value otherwise.

### GetTransactionTokenOk

`func (o *BankTransferTransitionResponseModel) GetTransactionTokenOk() (*string, bool)`

GetTransactionTokenOk returns a tuple with the TransactionToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionToken

`func (o *BankTransferTransitionResponseModel) SetTransactionToken(v string)`

SetTransactionToken sets TransactionToken field to given value.

### HasTransactionToken

`func (o *BankTransferTransitionResponseModel) HasTransactionToken() bool`

HasTransactionToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


