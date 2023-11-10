# BusinessCardHolderUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountHolderGroupToken** | Pointer to **string** | Associates the specified account holder group with the business. | [optional] 
**Active** | Pointer to **bool** | Specifies if the business is in the &#x60;ACTIVE&#x60; state on the Marqeta platform. | [optional] [default to true]
**AttestationConsent** | Pointer to **bool** | Indicates that the attester agrees that the information provided is correct and truthful.  This field is required for KYC verification (US-based accounts only). | [optional] [default to false]
**AttestationDate** | Pointer to **time.Time** | Timestamp of the attestation.  This field is required for KYC verification (US-based accounts only). | [optional] 
**AttesterName** | Pointer to **string** | Name of the attester for KYC verification.  This field is required for KYC verification (US-based accounts only). | [optional] 
**AttesterTitle** | Pointer to **string** | Title of the attester for KYC verification. | [optional] 
**BeneficialOwner1** | Pointer to [**BeneficialOwnerRequest**](BeneficialOwnerRequest.md) |  | [optional] 
**BeneficialOwner2** | Pointer to [**BeneficialOwnerRequest**](BeneficialOwnerRequest.md) |  | [optional] 
**BeneficialOwner3** | Pointer to [**BeneficialOwnerRequest**](BeneficialOwnerRequest.md) |  | [optional] 
**BeneficialOwner4** | Pointer to [**BeneficialOwnerRequest**](BeneficialOwnerRequest.md) |  | [optional] 
**BusinessNameDba** | Pointer to **string** | Fictitious business name (\&quot;Doing Business As\&quot; or DBA).  This field is required for KYC verification (US-based accounts only). If your business does not use a fictitious business name, enter your legal business name again in this field. | [optional] 
**BusinessNameLegal** | Pointer to **string** | Legal name of business.  This field is required for KYC verification (US-based accounts only). | [optional] 
**BusinessType** | Pointer to **string** | Indicates the type of business, for example business-to-business (B2B) or business-to-consumer (B2C). | [optional] 
**DateEstablished** | Pointer to **time.Time** | Date the business was established.  This field is required for KYC verification (US-based accounts only). | [optional] 
**DunsNumber** | Pointer to **string** | Data Universal Numbering System (DUNS) number of the business. | [optional] 
**GeneralBusinessDescription** | Pointer to **string** | General description of the business. | [optional] 
**History** | Pointer to **string** | History of the business. | [optional] 
**Identifications** | Pointer to [**[]IdentificationRequestModel**](IdentificationRequestModel.md) | One or more objects containing identifications associated with the business. | [optional] 
**InCurrentLocationSince** | Pointer to **time.Time** | Date on which the business office opened in its current location. | [optional] 
**Incorporation** | Pointer to [**BusinessIncorporation**](BusinessIncorporation.md) |  | [optional] 
**InternationalOfficeLocations** | Pointer to **string** | Locations of the business&#39; offices outside the United States. | [optional] 
**IpAddress** | Pointer to **string** | IP address of the business. | [optional] 
**Metadata** | Pointer to **map[string]string** | Associates any additional metadata you provide with the business. | [optional] 
**Notes** | Pointer to **string** | Any additional information pertaining to the business. | [optional] 
**OfficeLocation** | Pointer to [**AddressRequestModel**](AddressRequestModel.md) |  | [optional] 
**Password** | Pointer to **string** | Password for the business account on the Marqeta platform. | [optional] 
**Phone** | Pointer to **string** | 10-digit telephone number of business. | [optional] 
**PrimaryContact** | Pointer to [**PrimaryContactInfoModel**](PrimaryContactInfoModel.md) |  | [optional] 
**ProprietorIsBeneficialOwner** | Pointer to **bool** | Indicates that the proprietor or officer of the business is also a beneficial owner.  This field is required for KYC verification in the United States if the business proprietor or officer is also a beneficial owner. If the proprietor or business officer is a beneficial owner, use this field to indicate their beneficial ownership. Do not include information about the proprietor or business officer in a &#x60;beneficial_owner&#x60; object. | [optional] [default to false]
**ProprietorOrOfficer** | Pointer to [**BusinessProprietor**](BusinessProprietor.md) |  | [optional] 
**TaxpayerId** | Pointer to **string** | Taxpayer identifier of the business. | [optional] 
**Token** | Pointer to **string** | Unique identifier of the business.  Send a &#x60;GET&#x60; request to &#x60;/businesses&#x60; to retrieve business tokens. | [optional] 
**Website** | Pointer to **string** | URL of the business&#39; website. | [optional] 

## Methods

### NewBusinessCardHolderUpdate

`func NewBusinessCardHolderUpdate() *BusinessCardHolderUpdate`

NewBusinessCardHolderUpdate instantiates a new BusinessCardHolderUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBusinessCardHolderUpdateWithDefaults

`func NewBusinessCardHolderUpdateWithDefaults() *BusinessCardHolderUpdate`

NewBusinessCardHolderUpdateWithDefaults instantiates a new BusinessCardHolderUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountHolderGroupToken

`func (o *BusinessCardHolderUpdate) GetAccountHolderGroupToken() string`

GetAccountHolderGroupToken returns the AccountHolderGroupToken field if non-nil, zero value otherwise.

### GetAccountHolderGroupTokenOk

`func (o *BusinessCardHolderUpdate) GetAccountHolderGroupTokenOk() (*string, bool)`

GetAccountHolderGroupTokenOk returns a tuple with the AccountHolderGroupToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountHolderGroupToken

`func (o *BusinessCardHolderUpdate) SetAccountHolderGroupToken(v string)`

SetAccountHolderGroupToken sets AccountHolderGroupToken field to given value.

### HasAccountHolderGroupToken

`func (o *BusinessCardHolderUpdate) HasAccountHolderGroupToken() bool`

HasAccountHolderGroupToken returns a boolean if a field has been set.

### GetActive

`func (o *BusinessCardHolderUpdate) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *BusinessCardHolderUpdate) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *BusinessCardHolderUpdate) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *BusinessCardHolderUpdate) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAttestationConsent

`func (o *BusinessCardHolderUpdate) GetAttestationConsent() bool`

GetAttestationConsent returns the AttestationConsent field if non-nil, zero value otherwise.

### GetAttestationConsentOk

`func (o *BusinessCardHolderUpdate) GetAttestationConsentOk() (*bool, bool)`

GetAttestationConsentOk returns a tuple with the AttestationConsent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttestationConsent

`func (o *BusinessCardHolderUpdate) SetAttestationConsent(v bool)`

SetAttestationConsent sets AttestationConsent field to given value.

### HasAttestationConsent

`func (o *BusinessCardHolderUpdate) HasAttestationConsent() bool`

HasAttestationConsent returns a boolean if a field has been set.

### GetAttestationDate

`func (o *BusinessCardHolderUpdate) GetAttestationDate() time.Time`

GetAttestationDate returns the AttestationDate field if non-nil, zero value otherwise.

### GetAttestationDateOk

`func (o *BusinessCardHolderUpdate) GetAttestationDateOk() (*time.Time, bool)`

GetAttestationDateOk returns a tuple with the AttestationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttestationDate

`func (o *BusinessCardHolderUpdate) SetAttestationDate(v time.Time)`

SetAttestationDate sets AttestationDate field to given value.

### HasAttestationDate

`func (o *BusinessCardHolderUpdate) HasAttestationDate() bool`

HasAttestationDate returns a boolean if a field has been set.

### GetAttesterName

`func (o *BusinessCardHolderUpdate) GetAttesterName() string`

GetAttesterName returns the AttesterName field if non-nil, zero value otherwise.

### GetAttesterNameOk

`func (o *BusinessCardHolderUpdate) GetAttesterNameOk() (*string, bool)`

GetAttesterNameOk returns a tuple with the AttesterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttesterName

`func (o *BusinessCardHolderUpdate) SetAttesterName(v string)`

SetAttesterName sets AttesterName field to given value.

### HasAttesterName

`func (o *BusinessCardHolderUpdate) HasAttesterName() bool`

HasAttesterName returns a boolean if a field has been set.

### GetAttesterTitle

`func (o *BusinessCardHolderUpdate) GetAttesterTitle() string`

GetAttesterTitle returns the AttesterTitle field if non-nil, zero value otherwise.

### GetAttesterTitleOk

`func (o *BusinessCardHolderUpdate) GetAttesterTitleOk() (*string, bool)`

GetAttesterTitleOk returns a tuple with the AttesterTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttesterTitle

`func (o *BusinessCardHolderUpdate) SetAttesterTitle(v string)`

SetAttesterTitle sets AttesterTitle field to given value.

### HasAttesterTitle

`func (o *BusinessCardHolderUpdate) HasAttesterTitle() bool`

HasAttesterTitle returns a boolean if a field has been set.

### GetBeneficialOwner1

`func (o *BusinessCardHolderUpdate) GetBeneficialOwner1() BeneficialOwnerRequest`

GetBeneficialOwner1 returns the BeneficialOwner1 field if non-nil, zero value otherwise.

### GetBeneficialOwner1Ok

`func (o *BusinessCardHolderUpdate) GetBeneficialOwner1Ok() (*BeneficialOwnerRequest, bool)`

GetBeneficialOwner1Ok returns a tuple with the BeneficialOwner1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficialOwner1

`func (o *BusinessCardHolderUpdate) SetBeneficialOwner1(v BeneficialOwnerRequest)`

SetBeneficialOwner1 sets BeneficialOwner1 field to given value.

### HasBeneficialOwner1

`func (o *BusinessCardHolderUpdate) HasBeneficialOwner1() bool`

HasBeneficialOwner1 returns a boolean if a field has been set.

### GetBeneficialOwner2

`func (o *BusinessCardHolderUpdate) GetBeneficialOwner2() BeneficialOwnerRequest`

GetBeneficialOwner2 returns the BeneficialOwner2 field if non-nil, zero value otherwise.

### GetBeneficialOwner2Ok

`func (o *BusinessCardHolderUpdate) GetBeneficialOwner2Ok() (*BeneficialOwnerRequest, bool)`

GetBeneficialOwner2Ok returns a tuple with the BeneficialOwner2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficialOwner2

`func (o *BusinessCardHolderUpdate) SetBeneficialOwner2(v BeneficialOwnerRequest)`

SetBeneficialOwner2 sets BeneficialOwner2 field to given value.

### HasBeneficialOwner2

`func (o *BusinessCardHolderUpdate) HasBeneficialOwner2() bool`

HasBeneficialOwner2 returns a boolean if a field has been set.

### GetBeneficialOwner3

`func (o *BusinessCardHolderUpdate) GetBeneficialOwner3() BeneficialOwnerRequest`

GetBeneficialOwner3 returns the BeneficialOwner3 field if non-nil, zero value otherwise.

### GetBeneficialOwner3Ok

`func (o *BusinessCardHolderUpdate) GetBeneficialOwner3Ok() (*BeneficialOwnerRequest, bool)`

GetBeneficialOwner3Ok returns a tuple with the BeneficialOwner3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficialOwner3

`func (o *BusinessCardHolderUpdate) SetBeneficialOwner3(v BeneficialOwnerRequest)`

SetBeneficialOwner3 sets BeneficialOwner3 field to given value.

### HasBeneficialOwner3

`func (o *BusinessCardHolderUpdate) HasBeneficialOwner3() bool`

HasBeneficialOwner3 returns a boolean if a field has been set.

### GetBeneficialOwner4

`func (o *BusinessCardHolderUpdate) GetBeneficialOwner4() BeneficialOwnerRequest`

GetBeneficialOwner4 returns the BeneficialOwner4 field if non-nil, zero value otherwise.

### GetBeneficialOwner4Ok

`func (o *BusinessCardHolderUpdate) GetBeneficialOwner4Ok() (*BeneficialOwnerRequest, bool)`

GetBeneficialOwner4Ok returns a tuple with the BeneficialOwner4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficialOwner4

`func (o *BusinessCardHolderUpdate) SetBeneficialOwner4(v BeneficialOwnerRequest)`

SetBeneficialOwner4 sets BeneficialOwner4 field to given value.

### HasBeneficialOwner4

`func (o *BusinessCardHolderUpdate) HasBeneficialOwner4() bool`

HasBeneficialOwner4 returns a boolean if a field has been set.

### GetBusinessNameDba

`func (o *BusinessCardHolderUpdate) GetBusinessNameDba() string`

GetBusinessNameDba returns the BusinessNameDba field if non-nil, zero value otherwise.

### GetBusinessNameDbaOk

`func (o *BusinessCardHolderUpdate) GetBusinessNameDbaOk() (*string, bool)`

GetBusinessNameDbaOk returns a tuple with the BusinessNameDba field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessNameDba

`func (o *BusinessCardHolderUpdate) SetBusinessNameDba(v string)`

SetBusinessNameDba sets BusinessNameDba field to given value.

### HasBusinessNameDba

`func (o *BusinessCardHolderUpdate) HasBusinessNameDba() bool`

HasBusinessNameDba returns a boolean if a field has been set.

### GetBusinessNameLegal

`func (o *BusinessCardHolderUpdate) GetBusinessNameLegal() string`

GetBusinessNameLegal returns the BusinessNameLegal field if non-nil, zero value otherwise.

### GetBusinessNameLegalOk

`func (o *BusinessCardHolderUpdate) GetBusinessNameLegalOk() (*string, bool)`

GetBusinessNameLegalOk returns a tuple with the BusinessNameLegal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessNameLegal

`func (o *BusinessCardHolderUpdate) SetBusinessNameLegal(v string)`

SetBusinessNameLegal sets BusinessNameLegal field to given value.

### HasBusinessNameLegal

`func (o *BusinessCardHolderUpdate) HasBusinessNameLegal() bool`

HasBusinessNameLegal returns a boolean if a field has been set.

### GetBusinessType

`func (o *BusinessCardHolderUpdate) GetBusinessType() string`

GetBusinessType returns the BusinessType field if non-nil, zero value otherwise.

### GetBusinessTypeOk

`func (o *BusinessCardHolderUpdate) GetBusinessTypeOk() (*string, bool)`

GetBusinessTypeOk returns a tuple with the BusinessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessType

`func (o *BusinessCardHolderUpdate) SetBusinessType(v string)`

SetBusinessType sets BusinessType field to given value.

### HasBusinessType

`func (o *BusinessCardHolderUpdate) HasBusinessType() bool`

HasBusinessType returns a boolean if a field has been set.

### GetDateEstablished

`func (o *BusinessCardHolderUpdate) GetDateEstablished() time.Time`

GetDateEstablished returns the DateEstablished field if non-nil, zero value otherwise.

### GetDateEstablishedOk

`func (o *BusinessCardHolderUpdate) GetDateEstablishedOk() (*time.Time, bool)`

GetDateEstablishedOk returns a tuple with the DateEstablished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateEstablished

`func (o *BusinessCardHolderUpdate) SetDateEstablished(v time.Time)`

SetDateEstablished sets DateEstablished field to given value.

### HasDateEstablished

`func (o *BusinessCardHolderUpdate) HasDateEstablished() bool`

HasDateEstablished returns a boolean if a field has been set.

### GetDunsNumber

`func (o *BusinessCardHolderUpdate) GetDunsNumber() string`

GetDunsNumber returns the DunsNumber field if non-nil, zero value otherwise.

### GetDunsNumberOk

`func (o *BusinessCardHolderUpdate) GetDunsNumberOk() (*string, bool)`

GetDunsNumberOk returns a tuple with the DunsNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDunsNumber

`func (o *BusinessCardHolderUpdate) SetDunsNumber(v string)`

SetDunsNumber sets DunsNumber field to given value.

### HasDunsNumber

`func (o *BusinessCardHolderUpdate) HasDunsNumber() bool`

HasDunsNumber returns a boolean if a field has been set.

### GetGeneralBusinessDescription

`func (o *BusinessCardHolderUpdate) GetGeneralBusinessDescription() string`

GetGeneralBusinessDescription returns the GeneralBusinessDescription field if non-nil, zero value otherwise.

### GetGeneralBusinessDescriptionOk

`func (o *BusinessCardHolderUpdate) GetGeneralBusinessDescriptionOk() (*string, bool)`

GetGeneralBusinessDescriptionOk returns a tuple with the GeneralBusinessDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneralBusinessDescription

`func (o *BusinessCardHolderUpdate) SetGeneralBusinessDescription(v string)`

SetGeneralBusinessDescription sets GeneralBusinessDescription field to given value.

### HasGeneralBusinessDescription

`func (o *BusinessCardHolderUpdate) HasGeneralBusinessDescription() bool`

HasGeneralBusinessDescription returns a boolean if a field has been set.

### GetHistory

`func (o *BusinessCardHolderUpdate) GetHistory() string`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *BusinessCardHolderUpdate) GetHistoryOk() (*string, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *BusinessCardHolderUpdate) SetHistory(v string)`

SetHistory sets History field to given value.

### HasHistory

`func (o *BusinessCardHolderUpdate) HasHistory() bool`

HasHistory returns a boolean if a field has been set.

### GetIdentifications

`func (o *BusinessCardHolderUpdate) GetIdentifications() []IdentificationRequestModel`

GetIdentifications returns the Identifications field if non-nil, zero value otherwise.

### GetIdentificationsOk

`func (o *BusinessCardHolderUpdate) GetIdentificationsOk() (*[]IdentificationRequestModel, bool)`

GetIdentificationsOk returns a tuple with the Identifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifications

`func (o *BusinessCardHolderUpdate) SetIdentifications(v []IdentificationRequestModel)`

SetIdentifications sets Identifications field to given value.

### HasIdentifications

`func (o *BusinessCardHolderUpdate) HasIdentifications() bool`

HasIdentifications returns a boolean if a field has been set.

### GetInCurrentLocationSince

`func (o *BusinessCardHolderUpdate) GetInCurrentLocationSince() time.Time`

GetInCurrentLocationSince returns the InCurrentLocationSince field if non-nil, zero value otherwise.

### GetInCurrentLocationSinceOk

`func (o *BusinessCardHolderUpdate) GetInCurrentLocationSinceOk() (*time.Time, bool)`

GetInCurrentLocationSinceOk returns a tuple with the InCurrentLocationSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInCurrentLocationSince

`func (o *BusinessCardHolderUpdate) SetInCurrentLocationSince(v time.Time)`

SetInCurrentLocationSince sets InCurrentLocationSince field to given value.

### HasInCurrentLocationSince

`func (o *BusinessCardHolderUpdate) HasInCurrentLocationSince() bool`

HasInCurrentLocationSince returns a boolean if a field has been set.

### GetIncorporation

`func (o *BusinessCardHolderUpdate) GetIncorporation() BusinessIncorporation`

GetIncorporation returns the Incorporation field if non-nil, zero value otherwise.

### GetIncorporationOk

`func (o *BusinessCardHolderUpdate) GetIncorporationOk() (*BusinessIncorporation, bool)`

GetIncorporationOk returns a tuple with the Incorporation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncorporation

`func (o *BusinessCardHolderUpdate) SetIncorporation(v BusinessIncorporation)`

SetIncorporation sets Incorporation field to given value.

### HasIncorporation

`func (o *BusinessCardHolderUpdate) HasIncorporation() bool`

HasIncorporation returns a boolean if a field has been set.

### GetInternationalOfficeLocations

`func (o *BusinessCardHolderUpdate) GetInternationalOfficeLocations() string`

GetInternationalOfficeLocations returns the InternationalOfficeLocations field if non-nil, zero value otherwise.

### GetInternationalOfficeLocationsOk

`func (o *BusinessCardHolderUpdate) GetInternationalOfficeLocationsOk() (*string, bool)`

GetInternationalOfficeLocationsOk returns a tuple with the InternationalOfficeLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternationalOfficeLocations

`func (o *BusinessCardHolderUpdate) SetInternationalOfficeLocations(v string)`

SetInternationalOfficeLocations sets InternationalOfficeLocations field to given value.

### HasInternationalOfficeLocations

`func (o *BusinessCardHolderUpdate) HasInternationalOfficeLocations() bool`

HasInternationalOfficeLocations returns a boolean if a field has been set.

### GetIpAddress

`func (o *BusinessCardHolderUpdate) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *BusinessCardHolderUpdate) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *BusinessCardHolderUpdate) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *BusinessCardHolderUpdate) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetMetadata

`func (o *BusinessCardHolderUpdate) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *BusinessCardHolderUpdate) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *BusinessCardHolderUpdate) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *BusinessCardHolderUpdate) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetNotes

`func (o *BusinessCardHolderUpdate) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *BusinessCardHolderUpdate) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *BusinessCardHolderUpdate) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *BusinessCardHolderUpdate) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetOfficeLocation

`func (o *BusinessCardHolderUpdate) GetOfficeLocation() AddressRequestModel`

GetOfficeLocation returns the OfficeLocation field if non-nil, zero value otherwise.

### GetOfficeLocationOk

`func (o *BusinessCardHolderUpdate) GetOfficeLocationOk() (*AddressRequestModel, bool)`

GetOfficeLocationOk returns a tuple with the OfficeLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficeLocation

`func (o *BusinessCardHolderUpdate) SetOfficeLocation(v AddressRequestModel)`

SetOfficeLocation sets OfficeLocation field to given value.

### HasOfficeLocation

`func (o *BusinessCardHolderUpdate) HasOfficeLocation() bool`

HasOfficeLocation returns a boolean if a field has been set.

### GetPassword

`func (o *BusinessCardHolderUpdate) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *BusinessCardHolderUpdate) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *BusinessCardHolderUpdate) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *BusinessCardHolderUpdate) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPhone

`func (o *BusinessCardHolderUpdate) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *BusinessCardHolderUpdate) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *BusinessCardHolderUpdate) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *BusinessCardHolderUpdate) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetPrimaryContact

`func (o *BusinessCardHolderUpdate) GetPrimaryContact() PrimaryContactInfoModel`

GetPrimaryContact returns the PrimaryContact field if non-nil, zero value otherwise.

### GetPrimaryContactOk

`func (o *BusinessCardHolderUpdate) GetPrimaryContactOk() (*PrimaryContactInfoModel, bool)`

GetPrimaryContactOk returns a tuple with the PrimaryContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryContact

`func (o *BusinessCardHolderUpdate) SetPrimaryContact(v PrimaryContactInfoModel)`

SetPrimaryContact sets PrimaryContact field to given value.

### HasPrimaryContact

`func (o *BusinessCardHolderUpdate) HasPrimaryContact() bool`

HasPrimaryContact returns a boolean if a field has been set.

### GetProprietorIsBeneficialOwner

`func (o *BusinessCardHolderUpdate) GetProprietorIsBeneficialOwner() bool`

GetProprietorIsBeneficialOwner returns the ProprietorIsBeneficialOwner field if non-nil, zero value otherwise.

### GetProprietorIsBeneficialOwnerOk

`func (o *BusinessCardHolderUpdate) GetProprietorIsBeneficialOwnerOk() (*bool, bool)`

GetProprietorIsBeneficialOwnerOk returns a tuple with the ProprietorIsBeneficialOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProprietorIsBeneficialOwner

`func (o *BusinessCardHolderUpdate) SetProprietorIsBeneficialOwner(v bool)`

SetProprietorIsBeneficialOwner sets ProprietorIsBeneficialOwner field to given value.

### HasProprietorIsBeneficialOwner

`func (o *BusinessCardHolderUpdate) HasProprietorIsBeneficialOwner() bool`

HasProprietorIsBeneficialOwner returns a boolean if a field has been set.

### GetProprietorOrOfficer

`func (o *BusinessCardHolderUpdate) GetProprietorOrOfficer() BusinessProprietor`

GetProprietorOrOfficer returns the ProprietorOrOfficer field if non-nil, zero value otherwise.

### GetProprietorOrOfficerOk

`func (o *BusinessCardHolderUpdate) GetProprietorOrOfficerOk() (*BusinessProprietor, bool)`

GetProprietorOrOfficerOk returns a tuple with the ProprietorOrOfficer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProprietorOrOfficer

`func (o *BusinessCardHolderUpdate) SetProprietorOrOfficer(v BusinessProprietor)`

SetProprietorOrOfficer sets ProprietorOrOfficer field to given value.

### HasProprietorOrOfficer

`func (o *BusinessCardHolderUpdate) HasProprietorOrOfficer() bool`

HasProprietorOrOfficer returns a boolean if a field has been set.

### GetTaxpayerId

`func (o *BusinessCardHolderUpdate) GetTaxpayerId() string`

GetTaxpayerId returns the TaxpayerId field if non-nil, zero value otherwise.

### GetTaxpayerIdOk

`func (o *BusinessCardHolderUpdate) GetTaxpayerIdOk() (*string, bool)`

GetTaxpayerIdOk returns a tuple with the TaxpayerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxpayerId

`func (o *BusinessCardHolderUpdate) SetTaxpayerId(v string)`

SetTaxpayerId sets TaxpayerId field to given value.

### HasTaxpayerId

`func (o *BusinessCardHolderUpdate) HasTaxpayerId() bool`

HasTaxpayerId returns a boolean if a field has been set.

### GetToken

`func (o *BusinessCardHolderUpdate) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *BusinessCardHolderUpdate) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *BusinessCardHolderUpdate) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *BusinessCardHolderUpdate) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetWebsite

`func (o *BusinessCardHolderUpdate) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *BusinessCardHolderUpdate) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *BusinessCardHolderUpdate) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *BusinessCardHolderUpdate) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


