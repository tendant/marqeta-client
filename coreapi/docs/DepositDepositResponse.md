# DepositDepositResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **float32** | Amount being debited or credited. | [optional] 
**BusinessToken** | Pointer to **string** | The unique identifier of the business account holder. | [optional] 
**CompanyDiscretionaryData** | Pointer to **string** | Company-specific data provided by the direct deposit originator. | [optional] 
**CompanyEntryDescription** | Pointer to **string** | Company-specific data provided by the direct deposit originator. | [optional] 
**CompanyIdentification** | Pointer to **string** | Alphanumeric code that identifies the direct deposit originator. | [optional] 
**CompanyName** | Pointer to **string** | Name of the direct deposit originator. | [optional] 
**CreatedTime** | Pointer to **time.Time** | Date and time when the direct deposit account was created. | [optional] 
**DirectDepositAccountToken** | Pointer to **string** | The unique identifier of the direct deposit account. | [optional] 
**IndividualIdentificationNumber** | Pointer to **string** | Accounting number by which the recipient is known to the direct deposit originator. | [optional] 
**IndividualName** | Pointer to **string** | Name of the direct deposit recipient. | [optional] 
**LastModifiedTime** | Pointer to **time.Time** | Date and time when the direct deposit account was last modified. | [optional] 
**SettlementDate** | Pointer to **time.Time** | Date and time when the credit or debit is applied. | [optional] 
**StandardEntryClassCode** | Pointer to **string** | Three-letter code identifying the type of entry. | [optional] 
**State** | Pointer to **string** | Current status of the direct deposit record. | [optional] 
**StateReason** | Pointer to **string** | Explanation for why the direct deposit record is in the current state. | [optional] 
**StateReasonCode** | Pointer to **string** | Standard code describing the reason for the current state. | [optional] 
**Token** | Pointer to **string** | The unique identifier of the direct deposit record. | [optional] 
**Type** | Pointer to **string** | Determines whether funds are being debited from or credited to the account. | [optional] 
**UserToken** | Pointer to **string** | The unique identifier of the user account holder. | [optional] 

## Methods

### NewDepositDepositResponse

`func NewDepositDepositResponse() *DepositDepositResponse`

NewDepositDepositResponse instantiates a new DepositDepositResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDepositDepositResponseWithDefaults

`func NewDepositDepositResponseWithDefaults() *DepositDepositResponse`

NewDepositDepositResponseWithDefaults instantiates a new DepositDepositResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *DepositDepositResponse) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *DepositDepositResponse) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *DepositDepositResponse) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *DepositDepositResponse) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetBusinessToken

`func (o *DepositDepositResponse) GetBusinessToken() string`

GetBusinessToken returns the BusinessToken field if non-nil, zero value otherwise.

### GetBusinessTokenOk

`func (o *DepositDepositResponse) GetBusinessTokenOk() (*string, bool)`

GetBusinessTokenOk returns a tuple with the BusinessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessToken

`func (o *DepositDepositResponse) SetBusinessToken(v string)`

SetBusinessToken sets BusinessToken field to given value.

### HasBusinessToken

`func (o *DepositDepositResponse) HasBusinessToken() bool`

HasBusinessToken returns a boolean if a field has been set.

### GetCompanyDiscretionaryData

`func (o *DepositDepositResponse) GetCompanyDiscretionaryData() string`

GetCompanyDiscretionaryData returns the CompanyDiscretionaryData field if non-nil, zero value otherwise.

### GetCompanyDiscretionaryDataOk

`func (o *DepositDepositResponse) GetCompanyDiscretionaryDataOk() (*string, bool)`

GetCompanyDiscretionaryDataOk returns a tuple with the CompanyDiscretionaryData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyDiscretionaryData

`func (o *DepositDepositResponse) SetCompanyDiscretionaryData(v string)`

SetCompanyDiscretionaryData sets CompanyDiscretionaryData field to given value.

### HasCompanyDiscretionaryData

`func (o *DepositDepositResponse) HasCompanyDiscretionaryData() bool`

HasCompanyDiscretionaryData returns a boolean if a field has been set.

### GetCompanyEntryDescription

`func (o *DepositDepositResponse) GetCompanyEntryDescription() string`

GetCompanyEntryDescription returns the CompanyEntryDescription field if non-nil, zero value otherwise.

### GetCompanyEntryDescriptionOk

`func (o *DepositDepositResponse) GetCompanyEntryDescriptionOk() (*string, bool)`

GetCompanyEntryDescriptionOk returns a tuple with the CompanyEntryDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyEntryDescription

`func (o *DepositDepositResponse) SetCompanyEntryDescription(v string)`

SetCompanyEntryDescription sets CompanyEntryDescription field to given value.

### HasCompanyEntryDescription

`func (o *DepositDepositResponse) HasCompanyEntryDescription() bool`

HasCompanyEntryDescription returns a boolean if a field has been set.

### GetCompanyIdentification

`func (o *DepositDepositResponse) GetCompanyIdentification() string`

GetCompanyIdentification returns the CompanyIdentification field if non-nil, zero value otherwise.

### GetCompanyIdentificationOk

`func (o *DepositDepositResponse) GetCompanyIdentificationOk() (*string, bool)`

GetCompanyIdentificationOk returns a tuple with the CompanyIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyIdentification

`func (o *DepositDepositResponse) SetCompanyIdentification(v string)`

SetCompanyIdentification sets CompanyIdentification field to given value.

### HasCompanyIdentification

`func (o *DepositDepositResponse) HasCompanyIdentification() bool`

HasCompanyIdentification returns a boolean if a field has been set.

### GetCompanyName

`func (o *DepositDepositResponse) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *DepositDepositResponse) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *DepositDepositResponse) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *DepositDepositResponse) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetCreatedTime

`func (o *DepositDepositResponse) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *DepositDepositResponse) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *DepositDepositResponse) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *DepositDepositResponse) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDirectDepositAccountToken

`func (o *DepositDepositResponse) GetDirectDepositAccountToken() string`

GetDirectDepositAccountToken returns the DirectDepositAccountToken field if non-nil, zero value otherwise.

### GetDirectDepositAccountTokenOk

`func (o *DepositDepositResponse) GetDirectDepositAccountTokenOk() (*string, bool)`

GetDirectDepositAccountTokenOk returns a tuple with the DirectDepositAccountToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectDepositAccountToken

`func (o *DepositDepositResponse) SetDirectDepositAccountToken(v string)`

SetDirectDepositAccountToken sets DirectDepositAccountToken field to given value.

### HasDirectDepositAccountToken

`func (o *DepositDepositResponse) HasDirectDepositAccountToken() bool`

HasDirectDepositAccountToken returns a boolean if a field has been set.

### GetIndividualIdentificationNumber

`func (o *DepositDepositResponse) GetIndividualIdentificationNumber() string`

GetIndividualIdentificationNumber returns the IndividualIdentificationNumber field if non-nil, zero value otherwise.

### GetIndividualIdentificationNumberOk

`func (o *DepositDepositResponse) GetIndividualIdentificationNumberOk() (*string, bool)`

GetIndividualIdentificationNumberOk returns a tuple with the IndividualIdentificationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividualIdentificationNumber

`func (o *DepositDepositResponse) SetIndividualIdentificationNumber(v string)`

SetIndividualIdentificationNumber sets IndividualIdentificationNumber field to given value.

### HasIndividualIdentificationNumber

`func (o *DepositDepositResponse) HasIndividualIdentificationNumber() bool`

HasIndividualIdentificationNumber returns a boolean if a field has been set.

### GetIndividualName

`func (o *DepositDepositResponse) GetIndividualName() string`

GetIndividualName returns the IndividualName field if non-nil, zero value otherwise.

### GetIndividualNameOk

`func (o *DepositDepositResponse) GetIndividualNameOk() (*string, bool)`

GetIndividualNameOk returns a tuple with the IndividualName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividualName

`func (o *DepositDepositResponse) SetIndividualName(v string)`

SetIndividualName sets IndividualName field to given value.

### HasIndividualName

`func (o *DepositDepositResponse) HasIndividualName() bool`

HasIndividualName returns a boolean if a field has been set.

### GetLastModifiedTime

`func (o *DepositDepositResponse) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *DepositDepositResponse) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *DepositDepositResponse) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.

### HasLastModifiedTime

`func (o *DepositDepositResponse) HasLastModifiedTime() bool`

HasLastModifiedTime returns a boolean if a field has been set.

### GetSettlementDate

`func (o *DepositDepositResponse) GetSettlementDate() time.Time`

GetSettlementDate returns the SettlementDate field if non-nil, zero value otherwise.

### GetSettlementDateOk

`func (o *DepositDepositResponse) GetSettlementDateOk() (*time.Time, bool)`

GetSettlementDateOk returns a tuple with the SettlementDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementDate

`func (o *DepositDepositResponse) SetSettlementDate(v time.Time)`

SetSettlementDate sets SettlementDate field to given value.

### HasSettlementDate

`func (o *DepositDepositResponse) HasSettlementDate() bool`

HasSettlementDate returns a boolean if a field has been set.

### GetStandardEntryClassCode

`func (o *DepositDepositResponse) GetStandardEntryClassCode() string`

GetStandardEntryClassCode returns the StandardEntryClassCode field if non-nil, zero value otherwise.

### GetStandardEntryClassCodeOk

`func (o *DepositDepositResponse) GetStandardEntryClassCodeOk() (*string, bool)`

GetStandardEntryClassCodeOk returns a tuple with the StandardEntryClassCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardEntryClassCode

`func (o *DepositDepositResponse) SetStandardEntryClassCode(v string)`

SetStandardEntryClassCode sets StandardEntryClassCode field to given value.

### HasStandardEntryClassCode

`func (o *DepositDepositResponse) HasStandardEntryClassCode() bool`

HasStandardEntryClassCode returns a boolean if a field has been set.

### GetState

`func (o *DepositDepositResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DepositDepositResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DepositDepositResponse) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *DepositDepositResponse) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStateReason

`func (o *DepositDepositResponse) GetStateReason() string`

GetStateReason returns the StateReason field if non-nil, zero value otherwise.

### GetStateReasonOk

`func (o *DepositDepositResponse) GetStateReasonOk() (*string, bool)`

GetStateReasonOk returns a tuple with the StateReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateReason

`func (o *DepositDepositResponse) SetStateReason(v string)`

SetStateReason sets StateReason field to given value.

### HasStateReason

`func (o *DepositDepositResponse) HasStateReason() bool`

HasStateReason returns a boolean if a field has been set.

### GetStateReasonCode

`func (o *DepositDepositResponse) GetStateReasonCode() string`

GetStateReasonCode returns the StateReasonCode field if non-nil, zero value otherwise.

### GetStateReasonCodeOk

`func (o *DepositDepositResponse) GetStateReasonCodeOk() (*string, bool)`

GetStateReasonCodeOk returns a tuple with the StateReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateReasonCode

`func (o *DepositDepositResponse) SetStateReasonCode(v string)`

SetStateReasonCode sets StateReasonCode field to given value.

### HasStateReasonCode

`func (o *DepositDepositResponse) HasStateReasonCode() bool`

HasStateReasonCode returns a boolean if a field has been set.

### GetToken

`func (o *DepositDepositResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DepositDepositResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DepositDepositResponse) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DepositDepositResponse) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetType

`func (o *DepositDepositResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DepositDepositResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DepositDepositResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DepositDepositResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUserToken

`func (o *DepositDepositResponse) GetUserToken() string`

GetUserToken returns the UserToken field if non-nil, zero value otherwise.

### GetUserTokenOk

`func (o *DepositDepositResponse) GetUserTokenOk() (*string, bool)`

GetUserTokenOk returns a tuple with the UserToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserToken

`func (o *DepositDepositResponse) SetUserToken(v string)`

SetUserToken sets UserToken field to given value.

### HasUserToken

`func (o *DepositDepositResponse) HasUserToken() bool`

HasUserToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


