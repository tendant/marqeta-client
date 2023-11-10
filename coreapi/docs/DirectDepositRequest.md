# DirectDepositRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountNumber** | **string** |  | 
**Amount** | **float32** |  | 
**CompanyDiscretionaryData** | Pointer to **string** |  | [optional] 
**CompanyEntryDescription** | Pointer to **string** |  | [optional] 
**CompanyIdentification** | Pointer to **string** |  | [optional] 
**CompanyName** | Pointer to **string** |  | [optional] 
**EarlyPayEligible** | Pointer to **bool** |  | [optional] [default to false]
**IndividualIdentificationNumber** | Pointer to **string** |  | [optional] 
**IndividualName** | Pointer to **string** |  | [optional] 
**SettlementDate** | **time.Time** |  | 
**StandardEntryClassCode** | Pointer to **string** |  | [optional] 
**Token** | Pointer to **string** |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewDirectDepositRequest

`func NewDirectDepositRequest(accountNumber string, amount float32, settlementDate time.Time, type_ string, ) *DirectDepositRequest`

NewDirectDepositRequest instantiates a new DirectDepositRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectDepositRequestWithDefaults

`func NewDirectDepositRequestWithDefaults() *DirectDepositRequest`

NewDirectDepositRequestWithDefaults instantiates a new DirectDepositRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNumber

`func (o *DirectDepositRequest) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *DirectDepositRequest) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *DirectDepositRequest) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.


### GetAmount

`func (o *DirectDepositRequest) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *DirectDepositRequest) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *DirectDepositRequest) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetCompanyDiscretionaryData

`func (o *DirectDepositRequest) GetCompanyDiscretionaryData() string`

GetCompanyDiscretionaryData returns the CompanyDiscretionaryData field if non-nil, zero value otherwise.

### GetCompanyDiscretionaryDataOk

`func (o *DirectDepositRequest) GetCompanyDiscretionaryDataOk() (*string, bool)`

GetCompanyDiscretionaryDataOk returns a tuple with the CompanyDiscretionaryData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyDiscretionaryData

`func (o *DirectDepositRequest) SetCompanyDiscretionaryData(v string)`

SetCompanyDiscretionaryData sets CompanyDiscretionaryData field to given value.

### HasCompanyDiscretionaryData

`func (o *DirectDepositRequest) HasCompanyDiscretionaryData() bool`

HasCompanyDiscretionaryData returns a boolean if a field has been set.

### GetCompanyEntryDescription

`func (o *DirectDepositRequest) GetCompanyEntryDescription() string`

GetCompanyEntryDescription returns the CompanyEntryDescription field if non-nil, zero value otherwise.

### GetCompanyEntryDescriptionOk

`func (o *DirectDepositRequest) GetCompanyEntryDescriptionOk() (*string, bool)`

GetCompanyEntryDescriptionOk returns a tuple with the CompanyEntryDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyEntryDescription

`func (o *DirectDepositRequest) SetCompanyEntryDescription(v string)`

SetCompanyEntryDescription sets CompanyEntryDescription field to given value.

### HasCompanyEntryDescription

`func (o *DirectDepositRequest) HasCompanyEntryDescription() bool`

HasCompanyEntryDescription returns a boolean if a field has been set.

### GetCompanyIdentification

`func (o *DirectDepositRequest) GetCompanyIdentification() string`

GetCompanyIdentification returns the CompanyIdentification field if non-nil, zero value otherwise.

### GetCompanyIdentificationOk

`func (o *DirectDepositRequest) GetCompanyIdentificationOk() (*string, bool)`

GetCompanyIdentificationOk returns a tuple with the CompanyIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyIdentification

`func (o *DirectDepositRequest) SetCompanyIdentification(v string)`

SetCompanyIdentification sets CompanyIdentification field to given value.

### HasCompanyIdentification

`func (o *DirectDepositRequest) HasCompanyIdentification() bool`

HasCompanyIdentification returns a boolean if a field has been set.

### GetCompanyName

`func (o *DirectDepositRequest) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *DirectDepositRequest) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *DirectDepositRequest) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *DirectDepositRequest) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetEarlyPayEligible

`func (o *DirectDepositRequest) GetEarlyPayEligible() bool`

GetEarlyPayEligible returns the EarlyPayEligible field if non-nil, zero value otherwise.

### GetEarlyPayEligibleOk

`func (o *DirectDepositRequest) GetEarlyPayEligibleOk() (*bool, bool)`

GetEarlyPayEligibleOk returns a tuple with the EarlyPayEligible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarlyPayEligible

`func (o *DirectDepositRequest) SetEarlyPayEligible(v bool)`

SetEarlyPayEligible sets EarlyPayEligible field to given value.

### HasEarlyPayEligible

`func (o *DirectDepositRequest) HasEarlyPayEligible() bool`

HasEarlyPayEligible returns a boolean if a field has been set.

### GetIndividualIdentificationNumber

`func (o *DirectDepositRequest) GetIndividualIdentificationNumber() string`

GetIndividualIdentificationNumber returns the IndividualIdentificationNumber field if non-nil, zero value otherwise.

### GetIndividualIdentificationNumberOk

`func (o *DirectDepositRequest) GetIndividualIdentificationNumberOk() (*string, bool)`

GetIndividualIdentificationNumberOk returns a tuple with the IndividualIdentificationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividualIdentificationNumber

`func (o *DirectDepositRequest) SetIndividualIdentificationNumber(v string)`

SetIndividualIdentificationNumber sets IndividualIdentificationNumber field to given value.

### HasIndividualIdentificationNumber

`func (o *DirectDepositRequest) HasIndividualIdentificationNumber() bool`

HasIndividualIdentificationNumber returns a boolean if a field has been set.

### GetIndividualName

`func (o *DirectDepositRequest) GetIndividualName() string`

GetIndividualName returns the IndividualName field if non-nil, zero value otherwise.

### GetIndividualNameOk

`func (o *DirectDepositRequest) GetIndividualNameOk() (*string, bool)`

GetIndividualNameOk returns a tuple with the IndividualName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividualName

`func (o *DirectDepositRequest) SetIndividualName(v string)`

SetIndividualName sets IndividualName field to given value.

### HasIndividualName

`func (o *DirectDepositRequest) HasIndividualName() bool`

HasIndividualName returns a boolean if a field has been set.

### GetSettlementDate

`func (o *DirectDepositRequest) GetSettlementDate() time.Time`

GetSettlementDate returns the SettlementDate field if non-nil, zero value otherwise.

### GetSettlementDateOk

`func (o *DirectDepositRequest) GetSettlementDateOk() (*time.Time, bool)`

GetSettlementDateOk returns a tuple with the SettlementDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementDate

`func (o *DirectDepositRequest) SetSettlementDate(v time.Time)`

SetSettlementDate sets SettlementDate field to given value.


### GetStandardEntryClassCode

`func (o *DirectDepositRequest) GetStandardEntryClassCode() string`

GetStandardEntryClassCode returns the StandardEntryClassCode field if non-nil, zero value otherwise.

### GetStandardEntryClassCodeOk

`func (o *DirectDepositRequest) GetStandardEntryClassCodeOk() (*string, bool)`

GetStandardEntryClassCodeOk returns a tuple with the StandardEntryClassCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardEntryClassCode

`func (o *DirectDepositRequest) SetStandardEntryClassCode(v string)`

SetStandardEntryClassCode sets StandardEntryClassCode field to given value.

### HasStandardEntryClassCode

`func (o *DirectDepositRequest) HasStandardEntryClassCode() bool`

HasStandardEntryClassCode returns a boolean if a field has been set.

### GetToken

`func (o *DirectDepositRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DirectDepositRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DirectDepositRequest) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DirectDepositRequest) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetType

`func (o *DirectDepositRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DirectDepositRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DirectDepositRequest) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


