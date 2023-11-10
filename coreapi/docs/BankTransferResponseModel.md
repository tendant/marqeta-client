# BankTransferResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **float32** | Amount to push or pull. | 
**BatchNumber** | Pointer to **string** | Field required in older versions of the API, but no longer used. | [optional] 
**Channel** | Pointer to **string** | default &#x3D; API | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**CreatedTime** | Pointer to **time.Time** | Date and time when the ACH transfer was created, in UTC. | [optional] 
**CurrencyCode** | Pointer to **string** | Valid alpha-3 link:https://www.iso.org/iso-4217-currency-codes.html[ISO 4217 currency code, window&#x3D;\&quot;_blank\&quot;] | [optional] 
**FundingSourceToken** | **string** | ACH funding source token for the external account. | 
**LastModifiedTime** | Pointer to **time.Time** | Date and time when the ACH transfer was last modified, in UTC. | [optional] 
**Memo** | Pointer to **string** | Additional text describing the ACH transfer. | [optional] 
**ReturnCode** | Pointer to **string** | Standardized ACH return code for a returned transaction, generally sent by the RDFI.  Transactions can be returned for any of the reasons listed in the &lt;&lt;/developer-guides/ach-origination#_nacha_ach_return_codes, NACHA ACH return codes table&gt;&gt; of the ACH Origination Guide. | [optional] 
**ReturnReason** | Pointer to **string** | Human-readable description correlating to the &#x60;return_code&#x60;, such as &#x60;Insufficient funds&#x60;, if a return code is present in the response. | [optional] 
**StandardEntryClassCode** | Pointer to **string** | Three-letter code identifying the type of entry.  * *WEB* — An internet-initiated entry * *PPD* — Prearranged Payment and Deposit * *CCD* — Cash Concentration and Disbursement | [optional] 
**StatementDescriptor** | Pointer to **string** | Description of the transaction, as it will appear on the receiver&#39;s bank statement. | [optional] 
**Status** | Pointer to **string** | New state of the ACH transfer. | [optional] 
**Token** | Pointer to **string** | Unique identifier of the ACH transfer to retrieve. | [optional] 
**TransferSpeed** | Pointer to **string** | Specifies how quickly to initiate the ACH transfer.  *NOTE:* Same-day transfers are limited to a maximum amount of $100,000. | [optional] 
**Transitions** | Pointer to [**[]BankTransferTransitionResponseModel**](BankTransferTransitionResponseModel.md) | Array of ACH transfer transition objects. | [optional] 
**Type** | **string** | Specifies whether the ACH transfer is a push (credit) or pull (debit). | 

## Methods

### NewBankTransferResponseModel

`func NewBankTransferResponseModel(amount float32, fundingSourceToken string, type_ string, ) *BankTransferResponseModel`

NewBankTransferResponseModel instantiates a new BankTransferResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankTransferResponseModelWithDefaults

`func NewBankTransferResponseModelWithDefaults() *BankTransferResponseModel`

NewBankTransferResponseModelWithDefaults instantiates a new BankTransferResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *BankTransferResponseModel) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *BankTransferResponseModel) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *BankTransferResponseModel) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetBatchNumber

`func (o *BankTransferResponseModel) GetBatchNumber() string`

GetBatchNumber returns the BatchNumber field if non-nil, zero value otherwise.

### GetBatchNumberOk

`func (o *BankTransferResponseModel) GetBatchNumberOk() (*string, bool)`

GetBatchNumberOk returns a tuple with the BatchNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchNumber

`func (o *BankTransferResponseModel) SetBatchNumber(v string)`

SetBatchNumber sets BatchNumber field to given value.

### HasBatchNumber

`func (o *BankTransferResponseModel) HasBatchNumber() bool`

HasBatchNumber returns a boolean if a field has been set.

### GetChannel

`func (o *BankTransferResponseModel) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *BankTransferResponseModel) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *BankTransferResponseModel) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *BankTransferResponseModel) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetCreatedBy

`func (o *BankTransferResponseModel) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *BankTransferResponseModel) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *BankTransferResponseModel) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *BankTransferResponseModel) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedTime

`func (o *BankTransferResponseModel) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *BankTransferResponseModel) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *BankTransferResponseModel) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *BankTransferResponseModel) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *BankTransferResponseModel) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *BankTransferResponseModel) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *BankTransferResponseModel) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *BankTransferResponseModel) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetFundingSourceToken

`func (o *BankTransferResponseModel) GetFundingSourceToken() string`

GetFundingSourceToken returns the FundingSourceToken field if non-nil, zero value otherwise.

### GetFundingSourceTokenOk

`func (o *BankTransferResponseModel) GetFundingSourceTokenOk() (*string, bool)`

GetFundingSourceTokenOk returns a tuple with the FundingSourceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingSourceToken

`func (o *BankTransferResponseModel) SetFundingSourceToken(v string)`

SetFundingSourceToken sets FundingSourceToken field to given value.


### GetLastModifiedTime

`func (o *BankTransferResponseModel) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *BankTransferResponseModel) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *BankTransferResponseModel) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.

### HasLastModifiedTime

`func (o *BankTransferResponseModel) HasLastModifiedTime() bool`

HasLastModifiedTime returns a boolean if a field has been set.

### GetMemo

`func (o *BankTransferResponseModel) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *BankTransferResponseModel) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *BankTransferResponseModel) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *BankTransferResponseModel) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetReturnCode

`func (o *BankTransferResponseModel) GetReturnCode() string`

GetReturnCode returns the ReturnCode field if non-nil, zero value otherwise.

### GetReturnCodeOk

`func (o *BankTransferResponseModel) GetReturnCodeOk() (*string, bool)`

GetReturnCodeOk returns a tuple with the ReturnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnCode

`func (o *BankTransferResponseModel) SetReturnCode(v string)`

SetReturnCode sets ReturnCode field to given value.

### HasReturnCode

`func (o *BankTransferResponseModel) HasReturnCode() bool`

HasReturnCode returns a boolean if a field has been set.

### GetReturnReason

`func (o *BankTransferResponseModel) GetReturnReason() string`

GetReturnReason returns the ReturnReason field if non-nil, zero value otherwise.

### GetReturnReasonOk

`func (o *BankTransferResponseModel) GetReturnReasonOk() (*string, bool)`

GetReturnReasonOk returns a tuple with the ReturnReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnReason

`func (o *BankTransferResponseModel) SetReturnReason(v string)`

SetReturnReason sets ReturnReason field to given value.

### HasReturnReason

`func (o *BankTransferResponseModel) HasReturnReason() bool`

HasReturnReason returns a boolean if a field has been set.

### GetStandardEntryClassCode

`func (o *BankTransferResponseModel) GetStandardEntryClassCode() string`

GetStandardEntryClassCode returns the StandardEntryClassCode field if non-nil, zero value otherwise.

### GetStandardEntryClassCodeOk

`func (o *BankTransferResponseModel) GetStandardEntryClassCodeOk() (*string, bool)`

GetStandardEntryClassCodeOk returns a tuple with the StandardEntryClassCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardEntryClassCode

`func (o *BankTransferResponseModel) SetStandardEntryClassCode(v string)`

SetStandardEntryClassCode sets StandardEntryClassCode field to given value.

### HasStandardEntryClassCode

`func (o *BankTransferResponseModel) HasStandardEntryClassCode() bool`

HasStandardEntryClassCode returns a boolean if a field has been set.

### GetStatementDescriptor

`func (o *BankTransferResponseModel) GetStatementDescriptor() string`

GetStatementDescriptor returns the StatementDescriptor field if non-nil, zero value otherwise.

### GetStatementDescriptorOk

`func (o *BankTransferResponseModel) GetStatementDescriptorOk() (*string, bool)`

GetStatementDescriptorOk returns a tuple with the StatementDescriptor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatementDescriptor

`func (o *BankTransferResponseModel) SetStatementDescriptor(v string)`

SetStatementDescriptor sets StatementDescriptor field to given value.

### HasStatementDescriptor

`func (o *BankTransferResponseModel) HasStatementDescriptor() bool`

HasStatementDescriptor returns a boolean if a field has been set.

### GetStatus

`func (o *BankTransferResponseModel) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BankTransferResponseModel) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BankTransferResponseModel) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BankTransferResponseModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetToken

`func (o *BankTransferResponseModel) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *BankTransferResponseModel) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *BankTransferResponseModel) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *BankTransferResponseModel) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTransferSpeed

`func (o *BankTransferResponseModel) GetTransferSpeed() string`

GetTransferSpeed returns the TransferSpeed field if non-nil, zero value otherwise.

### GetTransferSpeedOk

`func (o *BankTransferResponseModel) GetTransferSpeedOk() (*string, bool)`

GetTransferSpeedOk returns a tuple with the TransferSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferSpeed

`func (o *BankTransferResponseModel) SetTransferSpeed(v string)`

SetTransferSpeed sets TransferSpeed field to given value.

### HasTransferSpeed

`func (o *BankTransferResponseModel) HasTransferSpeed() bool`

HasTransferSpeed returns a boolean if a field has been set.

### GetTransitions

`func (o *BankTransferResponseModel) GetTransitions() []BankTransferTransitionResponseModel`

GetTransitions returns the Transitions field if non-nil, zero value otherwise.

### GetTransitionsOk

`func (o *BankTransferResponseModel) GetTransitionsOk() (*[]BankTransferTransitionResponseModel, bool)`

GetTransitionsOk returns a tuple with the Transitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitions

`func (o *BankTransferResponseModel) SetTransitions(v []BankTransferTransitionResponseModel)`

SetTransitions sets Transitions field to given value.

### HasTransitions

`func (o *BankTransferResponseModel) HasTransitions() bool`

HasTransitions returns a boolean if a field has been set.

### GetType

`func (o *BankTransferResponseModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BankTransferResponseModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BankTransferResponseModel) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


