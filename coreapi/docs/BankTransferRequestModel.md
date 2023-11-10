# BankTransferRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **float32** | Amount to push or pull. | 
**Channel** | Pointer to **string** | default &#x3D; API | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**CurrencyCode** | Pointer to **string** | Currency of the push or pull. | [optional] 
**FundingSourceToken** | **string** | ACH funding source token for the external account. | 
**Memo** | Pointer to **string** | Additional text describing the ACH transfer. | [optional] 
**StandardEntryClassCode** | Pointer to **string** | Three-letter code identifying the type of entry.  * *WEB* — An internet-initiated entry * *PPD* — Prearranged Payment and Deposit * *CCD* — Cash Concentration and Disbursement | [optional] 
**StatementDescriptor** | Pointer to **string** | Description of the transaction, as it will appear on the receiver&#39;s bank statement. | [optional] 
**Token** | Pointer to **string** | Unique identifier of the ACH transfer to retrieve. | [optional] 
**TransferSpeed** | Pointer to **string** | Specifies how quickly to initiate the ACH transfer.  *NOTE:* Same-day transfers are limited to a maximum amount of $100,000. | [optional] 
**Type** | **string** | Specifies whether the ACH transfer is a push (credit) or pull (debit). | 

## Methods

### NewBankTransferRequestModel

`func NewBankTransferRequestModel(amount float32, fundingSourceToken string, type_ string, ) *BankTransferRequestModel`

NewBankTransferRequestModel instantiates a new BankTransferRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankTransferRequestModelWithDefaults

`func NewBankTransferRequestModelWithDefaults() *BankTransferRequestModel`

NewBankTransferRequestModelWithDefaults instantiates a new BankTransferRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *BankTransferRequestModel) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *BankTransferRequestModel) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *BankTransferRequestModel) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetChannel

`func (o *BankTransferRequestModel) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *BankTransferRequestModel) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *BankTransferRequestModel) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *BankTransferRequestModel) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetCreatedBy

`func (o *BankTransferRequestModel) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *BankTransferRequestModel) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *BankTransferRequestModel) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *BankTransferRequestModel) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *BankTransferRequestModel) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *BankTransferRequestModel) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *BankTransferRequestModel) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *BankTransferRequestModel) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetFundingSourceToken

`func (o *BankTransferRequestModel) GetFundingSourceToken() string`

GetFundingSourceToken returns the FundingSourceToken field if non-nil, zero value otherwise.

### GetFundingSourceTokenOk

`func (o *BankTransferRequestModel) GetFundingSourceTokenOk() (*string, bool)`

GetFundingSourceTokenOk returns a tuple with the FundingSourceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingSourceToken

`func (o *BankTransferRequestModel) SetFundingSourceToken(v string)`

SetFundingSourceToken sets FundingSourceToken field to given value.


### GetMemo

`func (o *BankTransferRequestModel) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *BankTransferRequestModel) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *BankTransferRequestModel) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *BankTransferRequestModel) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetStandardEntryClassCode

`func (o *BankTransferRequestModel) GetStandardEntryClassCode() string`

GetStandardEntryClassCode returns the StandardEntryClassCode field if non-nil, zero value otherwise.

### GetStandardEntryClassCodeOk

`func (o *BankTransferRequestModel) GetStandardEntryClassCodeOk() (*string, bool)`

GetStandardEntryClassCodeOk returns a tuple with the StandardEntryClassCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardEntryClassCode

`func (o *BankTransferRequestModel) SetStandardEntryClassCode(v string)`

SetStandardEntryClassCode sets StandardEntryClassCode field to given value.

### HasStandardEntryClassCode

`func (o *BankTransferRequestModel) HasStandardEntryClassCode() bool`

HasStandardEntryClassCode returns a boolean if a field has been set.

### GetStatementDescriptor

`func (o *BankTransferRequestModel) GetStatementDescriptor() string`

GetStatementDescriptor returns the StatementDescriptor field if non-nil, zero value otherwise.

### GetStatementDescriptorOk

`func (o *BankTransferRequestModel) GetStatementDescriptorOk() (*string, bool)`

GetStatementDescriptorOk returns a tuple with the StatementDescriptor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatementDescriptor

`func (o *BankTransferRequestModel) SetStatementDescriptor(v string)`

SetStatementDescriptor sets StatementDescriptor field to given value.

### HasStatementDescriptor

`func (o *BankTransferRequestModel) HasStatementDescriptor() bool`

HasStatementDescriptor returns a boolean if a field has been set.

### GetToken

`func (o *BankTransferRequestModel) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *BankTransferRequestModel) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *BankTransferRequestModel) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *BankTransferRequestModel) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTransferSpeed

`func (o *BankTransferRequestModel) GetTransferSpeed() string`

GetTransferSpeed returns the TransferSpeed field if non-nil, zero value otherwise.

### GetTransferSpeedOk

`func (o *BankTransferRequestModel) GetTransferSpeedOk() (*string, bool)`

GetTransferSpeedOk returns a tuple with the TransferSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferSpeed

`func (o *BankTransferRequestModel) SetTransferSpeed(v string)`

SetTransferSpeed sets TransferSpeed field to given value.

### HasTransferSpeed

`func (o *BankTransferRequestModel) HasTransferSpeed() bool`

HasTransferSpeed returns a boolean if a field has been set.

### GetType

`func (o *BankTransferRequestModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BankTransferRequestModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BankTransferRequestModel) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


