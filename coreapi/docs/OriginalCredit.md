# OriginalCredit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeferredHoldBy** | Pointer to **string** |  | [optional] 
**FastFundsEnabled** | Pointer to **bool** | Indicates that Fast Funds are enabled for dual-message original credit transactions. If the value of this field is &#x60;true&#x60;, you must make the funds available to your cardholder within 30 minutes of the transaction. | [optional] 
**FundingSource** | Pointer to **string** | Sender&#39;s account from which the OCT draws funds. | [optional] 
**ScreeningScore** | Pointer to **string** | Sanctions screening score to assist with meeting Anti-Money Laundering (AML) obligations.  Higher scores indicate that the sender&#39;s data more closely resembles an entry on the regulatory watchlist.  A value of 999 means that no screening score is available. | [optional] 
**SenderAccountType** | Pointer to **string** | The type of account from which the OCT draws funds. | [optional] 
**SenderAddress** | Pointer to **string** | Sender&#39;s street address. | [optional] 
**SenderCity** | Pointer to **string** | Sender&#39;s city. | [optional] 
**SenderCountry** | Pointer to **string** | Sender&#39;s country. | [optional] 
**SenderName** | Pointer to **string** | Full name of the sender. | [optional] 
**SenderState** | Pointer to **string** | Sender&#39;s state. | [optional] 
**TransactionPurpose** | Pointer to **string** | The purpose of the original credit transaction. | [optional] 
**TransactionType** | Pointer to **string** | Type of original credit transaction. | [optional] 

## Methods

### NewOriginalCredit

`func NewOriginalCredit() *OriginalCredit`

NewOriginalCredit instantiates a new OriginalCredit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOriginalCreditWithDefaults

`func NewOriginalCreditWithDefaults() *OriginalCredit`

NewOriginalCreditWithDefaults instantiates a new OriginalCredit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeferredHoldBy

`func (o *OriginalCredit) GetDeferredHoldBy() string`

GetDeferredHoldBy returns the DeferredHoldBy field if non-nil, zero value otherwise.

### GetDeferredHoldByOk

`func (o *OriginalCredit) GetDeferredHoldByOk() (*string, bool)`

GetDeferredHoldByOk returns a tuple with the DeferredHoldBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeferredHoldBy

`func (o *OriginalCredit) SetDeferredHoldBy(v string)`

SetDeferredHoldBy sets DeferredHoldBy field to given value.

### HasDeferredHoldBy

`func (o *OriginalCredit) HasDeferredHoldBy() bool`

HasDeferredHoldBy returns a boolean if a field has been set.

### GetFastFundsEnabled

`func (o *OriginalCredit) GetFastFundsEnabled() bool`

GetFastFundsEnabled returns the FastFundsEnabled field if non-nil, zero value otherwise.

### GetFastFundsEnabledOk

`func (o *OriginalCredit) GetFastFundsEnabledOk() (*bool, bool)`

GetFastFundsEnabledOk returns a tuple with the FastFundsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFastFundsEnabled

`func (o *OriginalCredit) SetFastFundsEnabled(v bool)`

SetFastFundsEnabled sets FastFundsEnabled field to given value.

### HasFastFundsEnabled

`func (o *OriginalCredit) HasFastFundsEnabled() bool`

HasFastFundsEnabled returns a boolean if a field has been set.

### GetFundingSource

`func (o *OriginalCredit) GetFundingSource() string`

GetFundingSource returns the FundingSource field if non-nil, zero value otherwise.

### GetFundingSourceOk

`func (o *OriginalCredit) GetFundingSourceOk() (*string, bool)`

GetFundingSourceOk returns a tuple with the FundingSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingSource

`func (o *OriginalCredit) SetFundingSource(v string)`

SetFundingSource sets FundingSource field to given value.

### HasFundingSource

`func (o *OriginalCredit) HasFundingSource() bool`

HasFundingSource returns a boolean if a field has been set.

### GetScreeningScore

`func (o *OriginalCredit) GetScreeningScore() string`

GetScreeningScore returns the ScreeningScore field if non-nil, zero value otherwise.

### GetScreeningScoreOk

`func (o *OriginalCredit) GetScreeningScoreOk() (*string, bool)`

GetScreeningScoreOk returns a tuple with the ScreeningScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreeningScore

`func (o *OriginalCredit) SetScreeningScore(v string)`

SetScreeningScore sets ScreeningScore field to given value.

### HasScreeningScore

`func (o *OriginalCredit) HasScreeningScore() bool`

HasScreeningScore returns a boolean if a field has been set.

### GetSenderAccountType

`func (o *OriginalCredit) GetSenderAccountType() string`

GetSenderAccountType returns the SenderAccountType field if non-nil, zero value otherwise.

### GetSenderAccountTypeOk

`func (o *OriginalCredit) GetSenderAccountTypeOk() (*string, bool)`

GetSenderAccountTypeOk returns a tuple with the SenderAccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderAccountType

`func (o *OriginalCredit) SetSenderAccountType(v string)`

SetSenderAccountType sets SenderAccountType field to given value.

### HasSenderAccountType

`func (o *OriginalCredit) HasSenderAccountType() bool`

HasSenderAccountType returns a boolean if a field has been set.

### GetSenderAddress

`func (o *OriginalCredit) GetSenderAddress() string`

GetSenderAddress returns the SenderAddress field if non-nil, zero value otherwise.

### GetSenderAddressOk

`func (o *OriginalCredit) GetSenderAddressOk() (*string, bool)`

GetSenderAddressOk returns a tuple with the SenderAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderAddress

`func (o *OriginalCredit) SetSenderAddress(v string)`

SetSenderAddress sets SenderAddress field to given value.

### HasSenderAddress

`func (o *OriginalCredit) HasSenderAddress() bool`

HasSenderAddress returns a boolean if a field has been set.

### GetSenderCity

`func (o *OriginalCredit) GetSenderCity() string`

GetSenderCity returns the SenderCity field if non-nil, zero value otherwise.

### GetSenderCityOk

`func (o *OriginalCredit) GetSenderCityOk() (*string, bool)`

GetSenderCityOk returns a tuple with the SenderCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderCity

`func (o *OriginalCredit) SetSenderCity(v string)`

SetSenderCity sets SenderCity field to given value.

### HasSenderCity

`func (o *OriginalCredit) HasSenderCity() bool`

HasSenderCity returns a boolean if a field has been set.

### GetSenderCountry

`func (o *OriginalCredit) GetSenderCountry() string`

GetSenderCountry returns the SenderCountry field if non-nil, zero value otherwise.

### GetSenderCountryOk

`func (o *OriginalCredit) GetSenderCountryOk() (*string, bool)`

GetSenderCountryOk returns a tuple with the SenderCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderCountry

`func (o *OriginalCredit) SetSenderCountry(v string)`

SetSenderCountry sets SenderCountry field to given value.

### HasSenderCountry

`func (o *OriginalCredit) HasSenderCountry() bool`

HasSenderCountry returns a boolean if a field has been set.

### GetSenderName

`func (o *OriginalCredit) GetSenderName() string`

GetSenderName returns the SenderName field if non-nil, zero value otherwise.

### GetSenderNameOk

`func (o *OriginalCredit) GetSenderNameOk() (*string, bool)`

GetSenderNameOk returns a tuple with the SenderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderName

`func (o *OriginalCredit) SetSenderName(v string)`

SetSenderName sets SenderName field to given value.

### HasSenderName

`func (o *OriginalCredit) HasSenderName() bool`

HasSenderName returns a boolean if a field has been set.

### GetSenderState

`func (o *OriginalCredit) GetSenderState() string`

GetSenderState returns the SenderState field if non-nil, zero value otherwise.

### GetSenderStateOk

`func (o *OriginalCredit) GetSenderStateOk() (*string, bool)`

GetSenderStateOk returns a tuple with the SenderState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderState

`func (o *OriginalCredit) SetSenderState(v string)`

SetSenderState sets SenderState field to given value.

### HasSenderState

`func (o *OriginalCredit) HasSenderState() bool`

HasSenderState returns a boolean if a field has been set.

### GetTransactionPurpose

`func (o *OriginalCredit) GetTransactionPurpose() string`

GetTransactionPurpose returns the TransactionPurpose field if non-nil, zero value otherwise.

### GetTransactionPurposeOk

`func (o *OriginalCredit) GetTransactionPurposeOk() (*string, bool)`

GetTransactionPurposeOk returns a tuple with the TransactionPurpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionPurpose

`func (o *OriginalCredit) SetTransactionPurpose(v string)`

SetTransactionPurpose sets TransactionPurpose field to given value.

### HasTransactionPurpose

`func (o *OriginalCredit) HasTransactionPurpose() bool`

HasTransactionPurpose returns a boolean if a field has been set.

### GetTransactionType

`func (o *OriginalCredit) GetTransactionType() string`

GetTransactionType returns the TransactionType field if non-nil, zero value otherwise.

### GetTransactionTypeOk

`func (o *OriginalCredit) GetTransactionTypeOk() (*string, bool)`

GetTransactionTypeOk returns a tuple with the TransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionType

`func (o *OriginalCredit) SetTransactionType(v string)`

SetTransactionType sets TransactionType field to given value.

### HasTransactionType

`func (o *OriginalCredit) HasTransactionType() bool`

HasTransactionType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


