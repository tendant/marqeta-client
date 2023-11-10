# DirectDepositSimulationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountNumber** | **string** | The account number of the user to debit or credit. | 
**Amount** | **float32** | Amount of the transaction. | 
**CompanyDiscretionaryData** | Pointer to **string** | Company-specific data provided by the direct deposit originator. | [optional] 
**CompanyEntryDescription** | Pointer to **string** | Company-specific data provided by the direct deposit originator. | [optional] 
**CompanyIdentification** | Pointer to **string** | Alphanumeric code that identifies the direct deposit originator. | [optional] 
**CompanyName** | Pointer to **string** | Name of the direct deposit originator. | [optional] 
**EarlyPayEligible** | Pointer to **bool** | Indicates whether the transaction is eligible for early pay. | [optional] [default to false]
**IndividualIdentificationNumber** | Pointer to **string** | Accounting number by which the recipient is known to the direct deposit originator. | [optional] 
**IndividualName** | Pointer to **string** | Identity of the direct deposit recipient. | [optional] 
**SettlementDate** | **time.Time** | Date when the credit or debit is applied. | 
**StandardEntryClassCode** | Pointer to **string** | Three-letter code identifying the type of entry. | [optional] 
**Token** | Pointer to **string** | Unique identifier of this direct deposit transition. If you do not include a token, the system generates one automatically. This token is used in other API calls, so rather than let the system generate a string, enter a string that you can remember. This value cannot be updated. | [optional] 
**Type** | Pointer to **string** | Whether funds are being debited from or credited to the account. | [optional] [readonly] 

## Methods

### NewDirectDepositSimulationRequest

`func NewDirectDepositSimulationRequest(accountNumber string, amount float32, settlementDate time.Time, ) *DirectDepositSimulationRequest`

NewDirectDepositSimulationRequest instantiates a new DirectDepositSimulationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectDepositSimulationRequestWithDefaults

`func NewDirectDepositSimulationRequestWithDefaults() *DirectDepositSimulationRequest`

NewDirectDepositSimulationRequestWithDefaults instantiates a new DirectDepositSimulationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNumber

`func (o *DirectDepositSimulationRequest) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *DirectDepositSimulationRequest) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *DirectDepositSimulationRequest) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.


### GetAmount

`func (o *DirectDepositSimulationRequest) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *DirectDepositSimulationRequest) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *DirectDepositSimulationRequest) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetCompanyDiscretionaryData

`func (o *DirectDepositSimulationRequest) GetCompanyDiscretionaryData() string`

GetCompanyDiscretionaryData returns the CompanyDiscretionaryData field if non-nil, zero value otherwise.

### GetCompanyDiscretionaryDataOk

`func (o *DirectDepositSimulationRequest) GetCompanyDiscretionaryDataOk() (*string, bool)`

GetCompanyDiscretionaryDataOk returns a tuple with the CompanyDiscretionaryData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyDiscretionaryData

`func (o *DirectDepositSimulationRequest) SetCompanyDiscretionaryData(v string)`

SetCompanyDiscretionaryData sets CompanyDiscretionaryData field to given value.

### HasCompanyDiscretionaryData

`func (o *DirectDepositSimulationRequest) HasCompanyDiscretionaryData() bool`

HasCompanyDiscretionaryData returns a boolean if a field has been set.

### GetCompanyEntryDescription

`func (o *DirectDepositSimulationRequest) GetCompanyEntryDescription() string`

GetCompanyEntryDescription returns the CompanyEntryDescription field if non-nil, zero value otherwise.

### GetCompanyEntryDescriptionOk

`func (o *DirectDepositSimulationRequest) GetCompanyEntryDescriptionOk() (*string, bool)`

GetCompanyEntryDescriptionOk returns a tuple with the CompanyEntryDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyEntryDescription

`func (o *DirectDepositSimulationRequest) SetCompanyEntryDescription(v string)`

SetCompanyEntryDescription sets CompanyEntryDescription field to given value.

### HasCompanyEntryDescription

`func (o *DirectDepositSimulationRequest) HasCompanyEntryDescription() bool`

HasCompanyEntryDescription returns a boolean if a field has been set.

### GetCompanyIdentification

`func (o *DirectDepositSimulationRequest) GetCompanyIdentification() string`

GetCompanyIdentification returns the CompanyIdentification field if non-nil, zero value otherwise.

### GetCompanyIdentificationOk

`func (o *DirectDepositSimulationRequest) GetCompanyIdentificationOk() (*string, bool)`

GetCompanyIdentificationOk returns a tuple with the CompanyIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyIdentification

`func (o *DirectDepositSimulationRequest) SetCompanyIdentification(v string)`

SetCompanyIdentification sets CompanyIdentification field to given value.

### HasCompanyIdentification

`func (o *DirectDepositSimulationRequest) HasCompanyIdentification() bool`

HasCompanyIdentification returns a boolean if a field has been set.

### GetCompanyName

`func (o *DirectDepositSimulationRequest) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *DirectDepositSimulationRequest) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *DirectDepositSimulationRequest) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *DirectDepositSimulationRequest) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetEarlyPayEligible

`func (o *DirectDepositSimulationRequest) GetEarlyPayEligible() bool`

GetEarlyPayEligible returns the EarlyPayEligible field if non-nil, zero value otherwise.

### GetEarlyPayEligibleOk

`func (o *DirectDepositSimulationRequest) GetEarlyPayEligibleOk() (*bool, bool)`

GetEarlyPayEligibleOk returns a tuple with the EarlyPayEligible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarlyPayEligible

`func (o *DirectDepositSimulationRequest) SetEarlyPayEligible(v bool)`

SetEarlyPayEligible sets EarlyPayEligible field to given value.

### HasEarlyPayEligible

`func (o *DirectDepositSimulationRequest) HasEarlyPayEligible() bool`

HasEarlyPayEligible returns a boolean if a field has been set.

### GetIndividualIdentificationNumber

`func (o *DirectDepositSimulationRequest) GetIndividualIdentificationNumber() string`

GetIndividualIdentificationNumber returns the IndividualIdentificationNumber field if non-nil, zero value otherwise.

### GetIndividualIdentificationNumberOk

`func (o *DirectDepositSimulationRequest) GetIndividualIdentificationNumberOk() (*string, bool)`

GetIndividualIdentificationNumberOk returns a tuple with the IndividualIdentificationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividualIdentificationNumber

`func (o *DirectDepositSimulationRequest) SetIndividualIdentificationNumber(v string)`

SetIndividualIdentificationNumber sets IndividualIdentificationNumber field to given value.

### HasIndividualIdentificationNumber

`func (o *DirectDepositSimulationRequest) HasIndividualIdentificationNumber() bool`

HasIndividualIdentificationNumber returns a boolean if a field has been set.

### GetIndividualName

`func (o *DirectDepositSimulationRequest) GetIndividualName() string`

GetIndividualName returns the IndividualName field if non-nil, zero value otherwise.

### GetIndividualNameOk

`func (o *DirectDepositSimulationRequest) GetIndividualNameOk() (*string, bool)`

GetIndividualNameOk returns a tuple with the IndividualName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividualName

`func (o *DirectDepositSimulationRequest) SetIndividualName(v string)`

SetIndividualName sets IndividualName field to given value.

### HasIndividualName

`func (o *DirectDepositSimulationRequest) HasIndividualName() bool`

HasIndividualName returns a boolean if a field has been set.

### GetSettlementDate

`func (o *DirectDepositSimulationRequest) GetSettlementDate() time.Time`

GetSettlementDate returns the SettlementDate field if non-nil, zero value otherwise.

### GetSettlementDateOk

`func (o *DirectDepositSimulationRequest) GetSettlementDateOk() (*time.Time, bool)`

GetSettlementDateOk returns a tuple with the SettlementDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementDate

`func (o *DirectDepositSimulationRequest) SetSettlementDate(v time.Time)`

SetSettlementDate sets SettlementDate field to given value.


### GetStandardEntryClassCode

`func (o *DirectDepositSimulationRequest) GetStandardEntryClassCode() string`

GetStandardEntryClassCode returns the StandardEntryClassCode field if non-nil, zero value otherwise.

### GetStandardEntryClassCodeOk

`func (o *DirectDepositSimulationRequest) GetStandardEntryClassCodeOk() (*string, bool)`

GetStandardEntryClassCodeOk returns a tuple with the StandardEntryClassCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardEntryClassCode

`func (o *DirectDepositSimulationRequest) SetStandardEntryClassCode(v string)`

SetStandardEntryClassCode sets StandardEntryClassCode field to given value.

### HasStandardEntryClassCode

`func (o *DirectDepositSimulationRequest) HasStandardEntryClassCode() bool`

HasStandardEntryClassCode returns a boolean if a field has been set.

### GetToken

`func (o *DirectDepositSimulationRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DirectDepositSimulationRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DirectDepositSimulationRequest) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DirectDepositSimulationRequest) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetType

`func (o *DirectDepositSimulationRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DirectDepositSimulationRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DirectDepositSimulationRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DirectDepositSimulationRequest) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


