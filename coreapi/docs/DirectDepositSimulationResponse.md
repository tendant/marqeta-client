# DirectDepositSimulationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Whether funds are being debited from or credited to the account. | [optional] 
**State** | Pointer to **string** | Current status of the direct deposit record. | [optional] 
**StateReason** | Pointer to **string** | Explanation of why the direct deposit record is in the current state. | [optional] 
**StateReasonCode** | Pointer to **string** | A standard code describing the reason for the current state. | [optional] 
**Token** | Pointer to **string** | Unique identifier of the direct deposit record. | [optional] 
**UserToken** | Pointer to **string** | Unique identifier of the user associated with the direct deposit record. | [optional] 
**Amount** | Pointer to **float32** | Amount being debited or credited. | [optional] 
**BusinessToken** | Pointer to **string** | Unique identifier of the business associated with the direct deposit record. | [optional] 
**CompanyDiscretionaryData** | Pointer to **string** | Company-specific data provided by the direct deposit originator. | [optional] 
**CompanyEntryDescription** | Pointer to **string** | Description of the purpose of the direct deposit. | [optional] 
**CompanyIdentification** | Pointer to **string** | Alphanumeric code that identifies the direct deposit originator. | [optional] 
**CompanyName** | Pointer to **string** | Name of the direct deposit originator. | [optional] 
**CreatedTime** | Pointer to **time.Time** | Date and time when the deposit account was created. | [optional] 
**DirectDepositAccountToken** | Pointer to **string** | Unique identifier of the affected deposit account. | [optional] 
**IndividualIdentificationNumber** | Pointer to **string** | Accounting number by which the recipient is known to the direct deposit originator. | [optional] 
**IndividualName** | Pointer to **string** | Name of the direct deposit recipient. | [optional] 
**LastModifiedTime** | Pointer to **time.Time** | Date and time when the direct deposit account was last modified. | [optional] 
**SettlementDate** | Pointer to **time.Time** | Date and time when the credit or debit is applied. | [optional] 
**StandardEntryClassCode** | Pointer to **string** | Three-letter code identifying the type of entry. | [optional] 

## Methods

### NewDirectDepositSimulationResponse

`func NewDirectDepositSimulationResponse() *DirectDepositSimulationResponse`

NewDirectDepositSimulationResponse instantiates a new DirectDepositSimulationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectDepositSimulationResponseWithDefaults

`func NewDirectDepositSimulationResponseWithDefaults() *DirectDepositSimulationResponse`

NewDirectDepositSimulationResponseWithDefaults instantiates a new DirectDepositSimulationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DirectDepositSimulationResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DirectDepositSimulationResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DirectDepositSimulationResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DirectDepositSimulationResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetState

`func (o *DirectDepositSimulationResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DirectDepositSimulationResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DirectDepositSimulationResponse) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *DirectDepositSimulationResponse) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStateReason

`func (o *DirectDepositSimulationResponse) GetStateReason() string`

GetStateReason returns the StateReason field if non-nil, zero value otherwise.

### GetStateReasonOk

`func (o *DirectDepositSimulationResponse) GetStateReasonOk() (*string, bool)`

GetStateReasonOk returns a tuple with the StateReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateReason

`func (o *DirectDepositSimulationResponse) SetStateReason(v string)`

SetStateReason sets StateReason field to given value.

### HasStateReason

`func (o *DirectDepositSimulationResponse) HasStateReason() bool`

HasStateReason returns a boolean if a field has been set.

### GetStateReasonCode

`func (o *DirectDepositSimulationResponse) GetStateReasonCode() string`

GetStateReasonCode returns the StateReasonCode field if non-nil, zero value otherwise.

### GetStateReasonCodeOk

`func (o *DirectDepositSimulationResponse) GetStateReasonCodeOk() (*string, bool)`

GetStateReasonCodeOk returns a tuple with the StateReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateReasonCode

`func (o *DirectDepositSimulationResponse) SetStateReasonCode(v string)`

SetStateReasonCode sets StateReasonCode field to given value.

### HasStateReasonCode

`func (o *DirectDepositSimulationResponse) HasStateReasonCode() bool`

HasStateReasonCode returns a boolean if a field has been set.

### GetToken

`func (o *DirectDepositSimulationResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DirectDepositSimulationResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DirectDepositSimulationResponse) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DirectDepositSimulationResponse) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUserToken

`func (o *DirectDepositSimulationResponse) GetUserToken() string`

GetUserToken returns the UserToken field if non-nil, zero value otherwise.

### GetUserTokenOk

`func (o *DirectDepositSimulationResponse) GetUserTokenOk() (*string, bool)`

GetUserTokenOk returns a tuple with the UserToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserToken

`func (o *DirectDepositSimulationResponse) SetUserToken(v string)`

SetUserToken sets UserToken field to given value.

### HasUserToken

`func (o *DirectDepositSimulationResponse) HasUserToken() bool`

HasUserToken returns a boolean if a field has been set.

### GetAmount

`func (o *DirectDepositSimulationResponse) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *DirectDepositSimulationResponse) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *DirectDepositSimulationResponse) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *DirectDepositSimulationResponse) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetBusinessToken

`func (o *DirectDepositSimulationResponse) GetBusinessToken() string`

GetBusinessToken returns the BusinessToken field if non-nil, zero value otherwise.

### GetBusinessTokenOk

`func (o *DirectDepositSimulationResponse) GetBusinessTokenOk() (*string, bool)`

GetBusinessTokenOk returns a tuple with the BusinessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessToken

`func (o *DirectDepositSimulationResponse) SetBusinessToken(v string)`

SetBusinessToken sets BusinessToken field to given value.

### HasBusinessToken

`func (o *DirectDepositSimulationResponse) HasBusinessToken() bool`

HasBusinessToken returns a boolean if a field has been set.

### GetCompanyDiscretionaryData

`func (o *DirectDepositSimulationResponse) GetCompanyDiscretionaryData() string`

GetCompanyDiscretionaryData returns the CompanyDiscretionaryData field if non-nil, zero value otherwise.

### GetCompanyDiscretionaryDataOk

`func (o *DirectDepositSimulationResponse) GetCompanyDiscretionaryDataOk() (*string, bool)`

GetCompanyDiscretionaryDataOk returns a tuple with the CompanyDiscretionaryData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyDiscretionaryData

`func (o *DirectDepositSimulationResponse) SetCompanyDiscretionaryData(v string)`

SetCompanyDiscretionaryData sets CompanyDiscretionaryData field to given value.

### HasCompanyDiscretionaryData

`func (o *DirectDepositSimulationResponse) HasCompanyDiscretionaryData() bool`

HasCompanyDiscretionaryData returns a boolean if a field has been set.

### GetCompanyEntryDescription

`func (o *DirectDepositSimulationResponse) GetCompanyEntryDescription() string`

GetCompanyEntryDescription returns the CompanyEntryDescription field if non-nil, zero value otherwise.

### GetCompanyEntryDescriptionOk

`func (o *DirectDepositSimulationResponse) GetCompanyEntryDescriptionOk() (*string, bool)`

GetCompanyEntryDescriptionOk returns a tuple with the CompanyEntryDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyEntryDescription

`func (o *DirectDepositSimulationResponse) SetCompanyEntryDescription(v string)`

SetCompanyEntryDescription sets CompanyEntryDescription field to given value.

### HasCompanyEntryDescription

`func (o *DirectDepositSimulationResponse) HasCompanyEntryDescription() bool`

HasCompanyEntryDescription returns a boolean if a field has been set.

### GetCompanyIdentification

`func (o *DirectDepositSimulationResponse) GetCompanyIdentification() string`

GetCompanyIdentification returns the CompanyIdentification field if non-nil, zero value otherwise.

### GetCompanyIdentificationOk

`func (o *DirectDepositSimulationResponse) GetCompanyIdentificationOk() (*string, bool)`

GetCompanyIdentificationOk returns a tuple with the CompanyIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyIdentification

`func (o *DirectDepositSimulationResponse) SetCompanyIdentification(v string)`

SetCompanyIdentification sets CompanyIdentification field to given value.

### HasCompanyIdentification

`func (o *DirectDepositSimulationResponse) HasCompanyIdentification() bool`

HasCompanyIdentification returns a boolean if a field has been set.

### GetCompanyName

`func (o *DirectDepositSimulationResponse) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *DirectDepositSimulationResponse) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *DirectDepositSimulationResponse) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *DirectDepositSimulationResponse) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetCreatedTime

`func (o *DirectDepositSimulationResponse) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *DirectDepositSimulationResponse) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *DirectDepositSimulationResponse) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *DirectDepositSimulationResponse) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDirectDepositAccountToken

`func (o *DirectDepositSimulationResponse) GetDirectDepositAccountToken() string`

GetDirectDepositAccountToken returns the DirectDepositAccountToken field if non-nil, zero value otherwise.

### GetDirectDepositAccountTokenOk

`func (o *DirectDepositSimulationResponse) GetDirectDepositAccountTokenOk() (*string, bool)`

GetDirectDepositAccountTokenOk returns a tuple with the DirectDepositAccountToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectDepositAccountToken

`func (o *DirectDepositSimulationResponse) SetDirectDepositAccountToken(v string)`

SetDirectDepositAccountToken sets DirectDepositAccountToken field to given value.

### HasDirectDepositAccountToken

`func (o *DirectDepositSimulationResponse) HasDirectDepositAccountToken() bool`

HasDirectDepositAccountToken returns a boolean if a field has been set.

### GetIndividualIdentificationNumber

`func (o *DirectDepositSimulationResponse) GetIndividualIdentificationNumber() string`

GetIndividualIdentificationNumber returns the IndividualIdentificationNumber field if non-nil, zero value otherwise.

### GetIndividualIdentificationNumberOk

`func (o *DirectDepositSimulationResponse) GetIndividualIdentificationNumberOk() (*string, bool)`

GetIndividualIdentificationNumberOk returns a tuple with the IndividualIdentificationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividualIdentificationNumber

`func (o *DirectDepositSimulationResponse) SetIndividualIdentificationNumber(v string)`

SetIndividualIdentificationNumber sets IndividualIdentificationNumber field to given value.

### HasIndividualIdentificationNumber

`func (o *DirectDepositSimulationResponse) HasIndividualIdentificationNumber() bool`

HasIndividualIdentificationNumber returns a boolean if a field has been set.

### GetIndividualName

`func (o *DirectDepositSimulationResponse) GetIndividualName() string`

GetIndividualName returns the IndividualName field if non-nil, zero value otherwise.

### GetIndividualNameOk

`func (o *DirectDepositSimulationResponse) GetIndividualNameOk() (*string, bool)`

GetIndividualNameOk returns a tuple with the IndividualName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividualName

`func (o *DirectDepositSimulationResponse) SetIndividualName(v string)`

SetIndividualName sets IndividualName field to given value.

### HasIndividualName

`func (o *DirectDepositSimulationResponse) HasIndividualName() bool`

HasIndividualName returns a boolean if a field has been set.

### GetLastModifiedTime

`func (o *DirectDepositSimulationResponse) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *DirectDepositSimulationResponse) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *DirectDepositSimulationResponse) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.

### HasLastModifiedTime

`func (o *DirectDepositSimulationResponse) HasLastModifiedTime() bool`

HasLastModifiedTime returns a boolean if a field has been set.

### GetSettlementDate

`func (o *DirectDepositSimulationResponse) GetSettlementDate() time.Time`

GetSettlementDate returns the SettlementDate field if non-nil, zero value otherwise.

### GetSettlementDateOk

`func (o *DirectDepositSimulationResponse) GetSettlementDateOk() (*time.Time, bool)`

GetSettlementDateOk returns a tuple with the SettlementDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementDate

`func (o *DirectDepositSimulationResponse) SetSettlementDate(v time.Time)`

SetSettlementDate sets SettlementDate field to given value.

### HasSettlementDate

`func (o *DirectDepositSimulationResponse) HasSettlementDate() bool`

HasSettlementDate returns a boolean if a field has been set.

### GetStandardEntryClassCode

`func (o *DirectDepositSimulationResponse) GetStandardEntryClassCode() string`

GetStandardEntryClassCode returns the StandardEntryClassCode field if non-nil, zero value otherwise.

### GetStandardEntryClassCodeOk

`func (o *DirectDepositSimulationResponse) GetStandardEntryClassCodeOk() (*string, bool)`

GetStandardEntryClassCodeOk returns a tuple with the StandardEntryClassCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardEntryClassCode

`func (o *DirectDepositSimulationResponse) SetStandardEntryClassCode(v string)`

SetStandardEntryClassCode sets StandardEntryClassCode field to given value.

### HasStandardEntryClassCode

`func (o *DirectDepositSimulationResponse) HasStandardEntryClassCode() bool`

HasStandardEntryClassCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


