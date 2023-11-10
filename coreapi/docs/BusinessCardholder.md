# BusinessCardholder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountHolderGroupToken** | Pointer to **string** | Existing account holder group token that associates the specified account holder group with the business. Send a &#x60;GET&#x60; request to &#x60;/accountholdergroups&#x60; to retrieve account holder group tokens. | [optional] 
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
**BusinessType** | Pointer to **string** | Indicates the type of business, for example B2B (business-to-business) or B2C (business-to-consumer). | [optional] 
**DateEstablished** | Pointer to **time.Time** | Date the business was established.  This field is required for KYC verification (US-based accounts only). | [optional] 
**DunsNumber** | Pointer to **string** | Data Universal Numbering System (DUNS) number of the business. | [optional] 
**GeneralBusinessDescription** | Pointer to **string** | General description of the business. | [optional] 
**History** | Pointer to **string** | History of the business. | [optional] 
**Identifications** | Pointer to [**[]IdentificationRequestModel**](IdentificationRequestModel.md) | One or more objects containing identifications associated with the business. | [optional] 
**InCurrentLocationSince** | Pointer to **time.Time** | Date on which the business office opened in its current location. | [optional] 
**Incorporation** | Pointer to [**BusinessIncorporation**](BusinessIncorporation.md) |  | [optional] 
**InternationalOfficeLocations** | Pointer to **string** | Locations of the business&#39; offices outside the US. | [optional] 
**IpAddress** | Pointer to **string** | IP address of the business. | [optional] 
**Metadata** | Pointer to **map[string]string** | Associates any additional metadata you provide with the business. | [optional] 
**Notes** | Pointer to **string** | Any additional information pertaining to the business. | [optional] 
**OfficeLocation** | Pointer to [**AddressRequestModel**](AddressRequestModel.md) |  | [optional] 
**Password** | Pointer to **string** | Password for the business account on the Marqeta platform. | [optional] 
**Phone** | Pointer to **string** | 10-digit telephone number of business. | [optional] 
**PrimaryContact** | Pointer to [**PrimaryContactInfoModel**](PrimaryContactInfoModel.md) |  | [optional] 
**ProprietorIsBeneficialOwner** | Pointer to **bool** | Indicates that the proprietor or officer of the business is also a beneficial owner.  This field is required for KYC verification if the business proprietor or officer is also a beneficial owner. | [optional] [default to false]
**ProprietorOrOfficer** | Pointer to [**BusinessProprietor**](BusinessProprietor.md) |  | [optional] 
**TaxpayerId** | Pointer to **string** | Taxpayer identifier of the business. | [optional] 
**Token** | Pointer to **string** | Unique identifier of the business resource.  If you do not include a token, the system generates one automatically. This token is necessary for use in other API calls, so we recommend that rather than let the system generate one, you use a simple string that is easy to remember. This value cannot be updated. | [optional] 
**Website** | Pointer to **string** | URL of the business&#39; website. | [optional] 

## Methods

### NewBusinessCardholder

`func NewBusinessCardholder() *BusinessCardholder`

NewBusinessCardholder instantiates a new BusinessCardholder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBusinessCardholderWithDefaults

`func NewBusinessCardholderWithDefaults() *BusinessCardholder`

NewBusinessCardholderWithDefaults instantiates a new BusinessCardholder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountHolderGroupToken

`func (o *BusinessCardholder) GetAccountHolderGroupToken() string`

GetAccountHolderGroupToken returns the AccountHolderGroupToken field if non-nil, zero value otherwise.

### GetAccountHolderGroupTokenOk

`func (o *BusinessCardholder) GetAccountHolderGroupTokenOk() (*string, bool)`

GetAccountHolderGroupTokenOk returns a tuple with the AccountHolderGroupToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountHolderGroupToken

`func (o *BusinessCardholder) SetAccountHolderGroupToken(v string)`

SetAccountHolderGroupToken sets AccountHolderGroupToken field to given value.

### HasAccountHolderGroupToken

`func (o *BusinessCardholder) HasAccountHolderGroupToken() bool`

HasAccountHolderGroupToken returns a boolean if a field has been set.

### GetActive

`func (o *BusinessCardholder) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *BusinessCardholder) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *BusinessCardholder) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *BusinessCardholder) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAttestationConsent

`func (o *BusinessCardholder) GetAttestationConsent() bool`

GetAttestationConsent returns the AttestationConsent field if non-nil, zero value otherwise.

### GetAttestationConsentOk

`func (o *BusinessCardholder) GetAttestationConsentOk() (*bool, bool)`

GetAttestationConsentOk returns a tuple with the AttestationConsent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttestationConsent

`func (o *BusinessCardholder) SetAttestationConsent(v bool)`

SetAttestationConsent sets AttestationConsent field to given value.

### HasAttestationConsent

`func (o *BusinessCardholder) HasAttestationConsent() bool`

HasAttestationConsent returns a boolean if a field has been set.

### GetAttestationDate

`func (o *BusinessCardholder) GetAttestationDate() time.Time`

GetAttestationDate returns the AttestationDate field if non-nil, zero value otherwise.

### GetAttestationDateOk

`func (o *BusinessCardholder) GetAttestationDateOk() (*time.Time, bool)`

GetAttestationDateOk returns a tuple with the AttestationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttestationDate

`func (o *BusinessCardholder) SetAttestationDate(v time.Time)`

SetAttestationDate sets AttestationDate field to given value.

### HasAttestationDate

`func (o *BusinessCardholder) HasAttestationDate() bool`

HasAttestationDate returns a boolean if a field has been set.

### GetAttesterName

`func (o *BusinessCardholder) GetAttesterName() string`

GetAttesterName returns the AttesterName field if non-nil, zero value otherwise.

### GetAttesterNameOk

`func (o *BusinessCardholder) GetAttesterNameOk() (*string, bool)`

GetAttesterNameOk returns a tuple with the AttesterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttesterName

`func (o *BusinessCardholder) SetAttesterName(v string)`

SetAttesterName sets AttesterName field to given value.

### HasAttesterName

`func (o *BusinessCardholder) HasAttesterName() bool`

HasAttesterName returns a boolean if a field has been set.

### GetAttesterTitle

`func (o *BusinessCardholder) GetAttesterTitle() string`

GetAttesterTitle returns the AttesterTitle field if non-nil, zero value otherwise.

### GetAttesterTitleOk

`func (o *BusinessCardholder) GetAttesterTitleOk() (*string, bool)`

GetAttesterTitleOk returns a tuple with the AttesterTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttesterTitle

`func (o *BusinessCardholder) SetAttesterTitle(v string)`

SetAttesterTitle sets AttesterTitle field to given value.

### HasAttesterTitle

`func (o *BusinessCardholder) HasAttesterTitle() bool`

HasAttesterTitle returns a boolean if a field has been set.

### GetBeneficialOwner1

`func (o *BusinessCardholder) GetBeneficialOwner1() BeneficialOwnerRequest`

GetBeneficialOwner1 returns the BeneficialOwner1 field if non-nil, zero value otherwise.

### GetBeneficialOwner1Ok

`func (o *BusinessCardholder) GetBeneficialOwner1Ok() (*BeneficialOwnerRequest, bool)`

GetBeneficialOwner1Ok returns a tuple with the BeneficialOwner1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficialOwner1

`func (o *BusinessCardholder) SetBeneficialOwner1(v BeneficialOwnerRequest)`

SetBeneficialOwner1 sets BeneficialOwner1 field to given value.

### HasBeneficialOwner1

`func (o *BusinessCardholder) HasBeneficialOwner1() bool`

HasBeneficialOwner1 returns a boolean if a field has been set.

### GetBeneficialOwner2

`func (o *BusinessCardholder) GetBeneficialOwner2() BeneficialOwnerRequest`

GetBeneficialOwner2 returns the BeneficialOwner2 field if non-nil, zero value otherwise.

### GetBeneficialOwner2Ok

`func (o *BusinessCardholder) GetBeneficialOwner2Ok() (*BeneficialOwnerRequest, bool)`

GetBeneficialOwner2Ok returns a tuple with the BeneficialOwner2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficialOwner2

`func (o *BusinessCardholder) SetBeneficialOwner2(v BeneficialOwnerRequest)`

SetBeneficialOwner2 sets BeneficialOwner2 field to given value.

### HasBeneficialOwner2

`func (o *BusinessCardholder) HasBeneficialOwner2() bool`

HasBeneficialOwner2 returns a boolean if a field has been set.

### GetBeneficialOwner3

`func (o *BusinessCardholder) GetBeneficialOwner3() BeneficialOwnerRequest`

GetBeneficialOwner3 returns the BeneficialOwner3 field if non-nil, zero value otherwise.

### GetBeneficialOwner3Ok

`func (o *BusinessCardholder) GetBeneficialOwner3Ok() (*BeneficialOwnerRequest, bool)`

GetBeneficialOwner3Ok returns a tuple with the BeneficialOwner3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficialOwner3

`func (o *BusinessCardholder) SetBeneficialOwner3(v BeneficialOwnerRequest)`

SetBeneficialOwner3 sets BeneficialOwner3 field to given value.

### HasBeneficialOwner3

`func (o *BusinessCardholder) HasBeneficialOwner3() bool`

HasBeneficialOwner3 returns a boolean if a field has been set.

### GetBeneficialOwner4

`func (o *BusinessCardholder) GetBeneficialOwner4() BeneficialOwnerRequest`

GetBeneficialOwner4 returns the BeneficialOwner4 field if non-nil, zero value otherwise.

### GetBeneficialOwner4Ok

`func (o *BusinessCardholder) GetBeneficialOwner4Ok() (*BeneficialOwnerRequest, bool)`

GetBeneficialOwner4Ok returns a tuple with the BeneficialOwner4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficialOwner4

`func (o *BusinessCardholder) SetBeneficialOwner4(v BeneficialOwnerRequest)`

SetBeneficialOwner4 sets BeneficialOwner4 field to given value.

### HasBeneficialOwner4

`func (o *BusinessCardholder) HasBeneficialOwner4() bool`

HasBeneficialOwner4 returns a boolean if a field has been set.

### GetBusinessNameDba

`func (o *BusinessCardholder) GetBusinessNameDba() string`

GetBusinessNameDba returns the BusinessNameDba field if non-nil, zero value otherwise.

### GetBusinessNameDbaOk

`func (o *BusinessCardholder) GetBusinessNameDbaOk() (*string, bool)`

GetBusinessNameDbaOk returns a tuple with the BusinessNameDba field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessNameDba

`func (o *BusinessCardholder) SetBusinessNameDba(v string)`

SetBusinessNameDba sets BusinessNameDba field to given value.

### HasBusinessNameDba

`func (o *BusinessCardholder) HasBusinessNameDba() bool`

HasBusinessNameDba returns a boolean if a field has been set.

### GetBusinessNameLegal

`func (o *BusinessCardholder) GetBusinessNameLegal() string`

GetBusinessNameLegal returns the BusinessNameLegal field if non-nil, zero value otherwise.

### GetBusinessNameLegalOk

`func (o *BusinessCardholder) GetBusinessNameLegalOk() (*string, bool)`

GetBusinessNameLegalOk returns a tuple with the BusinessNameLegal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessNameLegal

`func (o *BusinessCardholder) SetBusinessNameLegal(v string)`

SetBusinessNameLegal sets BusinessNameLegal field to given value.

### HasBusinessNameLegal

`func (o *BusinessCardholder) HasBusinessNameLegal() bool`

HasBusinessNameLegal returns a boolean if a field has been set.

### GetBusinessType

`func (o *BusinessCardholder) GetBusinessType() string`

GetBusinessType returns the BusinessType field if non-nil, zero value otherwise.

### GetBusinessTypeOk

`func (o *BusinessCardholder) GetBusinessTypeOk() (*string, bool)`

GetBusinessTypeOk returns a tuple with the BusinessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessType

`func (o *BusinessCardholder) SetBusinessType(v string)`

SetBusinessType sets BusinessType field to given value.

### HasBusinessType

`func (o *BusinessCardholder) HasBusinessType() bool`

HasBusinessType returns a boolean if a field has been set.

### GetDateEstablished

`func (o *BusinessCardholder) GetDateEstablished() time.Time`

GetDateEstablished returns the DateEstablished field if non-nil, zero value otherwise.

### GetDateEstablishedOk

`func (o *BusinessCardholder) GetDateEstablishedOk() (*time.Time, bool)`

GetDateEstablishedOk returns a tuple with the DateEstablished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateEstablished

`func (o *BusinessCardholder) SetDateEstablished(v time.Time)`

SetDateEstablished sets DateEstablished field to given value.

### HasDateEstablished

`func (o *BusinessCardholder) HasDateEstablished() bool`

HasDateEstablished returns a boolean if a field has been set.

### GetDunsNumber

`func (o *BusinessCardholder) GetDunsNumber() string`

GetDunsNumber returns the DunsNumber field if non-nil, zero value otherwise.

### GetDunsNumberOk

`func (o *BusinessCardholder) GetDunsNumberOk() (*string, bool)`

GetDunsNumberOk returns a tuple with the DunsNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDunsNumber

`func (o *BusinessCardholder) SetDunsNumber(v string)`

SetDunsNumber sets DunsNumber field to given value.

### HasDunsNumber

`func (o *BusinessCardholder) HasDunsNumber() bool`

HasDunsNumber returns a boolean if a field has been set.

### GetGeneralBusinessDescription

`func (o *BusinessCardholder) GetGeneralBusinessDescription() string`

GetGeneralBusinessDescription returns the GeneralBusinessDescription field if non-nil, zero value otherwise.

### GetGeneralBusinessDescriptionOk

`func (o *BusinessCardholder) GetGeneralBusinessDescriptionOk() (*string, bool)`

GetGeneralBusinessDescriptionOk returns a tuple with the GeneralBusinessDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneralBusinessDescription

`func (o *BusinessCardholder) SetGeneralBusinessDescription(v string)`

SetGeneralBusinessDescription sets GeneralBusinessDescription field to given value.

### HasGeneralBusinessDescription

`func (o *BusinessCardholder) HasGeneralBusinessDescription() bool`

HasGeneralBusinessDescription returns a boolean if a field has been set.

### GetHistory

`func (o *BusinessCardholder) GetHistory() string`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *BusinessCardholder) GetHistoryOk() (*string, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *BusinessCardholder) SetHistory(v string)`

SetHistory sets History field to given value.

### HasHistory

`func (o *BusinessCardholder) HasHistory() bool`

HasHistory returns a boolean if a field has been set.

### GetIdentifications

`func (o *BusinessCardholder) GetIdentifications() []IdentificationRequestModel`

GetIdentifications returns the Identifications field if non-nil, zero value otherwise.

### GetIdentificationsOk

`func (o *BusinessCardholder) GetIdentificationsOk() (*[]IdentificationRequestModel, bool)`

GetIdentificationsOk returns a tuple with the Identifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifications

`func (o *BusinessCardholder) SetIdentifications(v []IdentificationRequestModel)`

SetIdentifications sets Identifications field to given value.

### HasIdentifications

`func (o *BusinessCardholder) HasIdentifications() bool`

HasIdentifications returns a boolean if a field has been set.

### GetInCurrentLocationSince

`func (o *BusinessCardholder) GetInCurrentLocationSince() time.Time`

GetInCurrentLocationSince returns the InCurrentLocationSince field if non-nil, zero value otherwise.

### GetInCurrentLocationSinceOk

`func (o *BusinessCardholder) GetInCurrentLocationSinceOk() (*time.Time, bool)`

GetInCurrentLocationSinceOk returns a tuple with the InCurrentLocationSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInCurrentLocationSince

`func (o *BusinessCardholder) SetInCurrentLocationSince(v time.Time)`

SetInCurrentLocationSince sets InCurrentLocationSince field to given value.

### HasInCurrentLocationSince

`func (o *BusinessCardholder) HasInCurrentLocationSince() bool`

HasInCurrentLocationSince returns a boolean if a field has been set.

### GetIncorporation

`func (o *BusinessCardholder) GetIncorporation() BusinessIncorporation`

GetIncorporation returns the Incorporation field if non-nil, zero value otherwise.

### GetIncorporationOk

`func (o *BusinessCardholder) GetIncorporationOk() (*BusinessIncorporation, bool)`

GetIncorporationOk returns a tuple with the Incorporation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncorporation

`func (o *BusinessCardholder) SetIncorporation(v BusinessIncorporation)`

SetIncorporation sets Incorporation field to given value.

### HasIncorporation

`func (o *BusinessCardholder) HasIncorporation() bool`

HasIncorporation returns a boolean if a field has been set.

### GetInternationalOfficeLocations

`func (o *BusinessCardholder) GetInternationalOfficeLocations() string`

GetInternationalOfficeLocations returns the InternationalOfficeLocations field if non-nil, zero value otherwise.

### GetInternationalOfficeLocationsOk

`func (o *BusinessCardholder) GetInternationalOfficeLocationsOk() (*string, bool)`

GetInternationalOfficeLocationsOk returns a tuple with the InternationalOfficeLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternationalOfficeLocations

`func (o *BusinessCardholder) SetInternationalOfficeLocations(v string)`

SetInternationalOfficeLocations sets InternationalOfficeLocations field to given value.

### HasInternationalOfficeLocations

`func (o *BusinessCardholder) HasInternationalOfficeLocations() bool`

HasInternationalOfficeLocations returns a boolean if a field has been set.

### GetIpAddress

`func (o *BusinessCardholder) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *BusinessCardholder) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *BusinessCardholder) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *BusinessCardholder) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetMetadata

`func (o *BusinessCardholder) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *BusinessCardholder) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *BusinessCardholder) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *BusinessCardholder) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetNotes

`func (o *BusinessCardholder) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *BusinessCardholder) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *BusinessCardholder) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *BusinessCardholder) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetOfficeLocation

`func (o *BusinessCardholder) GetOfficeLocation() AddressRequestModel`

GetOfficeLocation returns the OfficeLocation field if non-nil, zero value otherwise.

### GetOfficeLocationOk

`func (o *BusinessCardholder) GetOfficeLocationOk() (*AddressRequestModel, bool)`

GetOfficeLocationOk returns a tuple with the OfficeLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficeLocation

`func (o *BusinessCardholder) SetOfficeLocation(v AddressRequestModel)`

SetOfficeLocation sets OfficeLocation field to given value.

### HasOfficeLocation

`func (o *BusinessCardholder) HasOfficeLocation() bool`

HasOfficeLocation returns a boolean if a field has been set.

### GetPassword

`func (o *BusinessCardholder) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *BusinessCardholder) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *BusinessCardholder) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *BusinessCardholder) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPhone

`func (o *BusinessCardholder) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *BusinessCardholder) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *BusinessCardholder) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *BusinessCardholder) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetPrimaryContact

`func (o *BusinessCardholder) GetPrimaryContact() PrimaryContactInfoModel`

GetPrimaryContact returns the PrimaryContact field if non-nil, zero value otherwise.

### GetPrimaryContactOk

`func (o *BusinessCardholder) GetPrimaryContactOk() (*PrimaryContactInfoModel, bool)`

GetPrimaryContactOk returns a tuple with the PrimaryContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryContact

`func (o *BusinessCardholder) SetPrimaryContact(v PrimaryContactInfoModel)`

SetPrimaryContact sets PrimaryContact field to given value.

### HasPrimaryContact

`func (o *BusinessCardholder) HasPrimaryContact() bool`

HasPrimaryContact returns a boolean if a field has been set.

### GetProprietorIsBeneficialOwner

`func (o *BusinessCardholder) GetProprietorIsBeneficialOwner() bool`

GetProprietorIsBeneficialOwner returns the ProprietorIsBeneficialOwner field if non-nil, zero value otherwise.

### GetProprietorIsBeneficialOwnerOk

`func (o *BusinessCardholder) GetProprietorIsBeneficialOwnerOk() (*bool, bool)`

GetProprietorIsBeneficialOwnerOk returns a tuple with the ProprietorIsBeneficialOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProprietorIsBeneficialOwner

`func (o *BusinessCardholder) SetProprietorIsBeneficialOwner(v bool)`

SetProprietorIsBeneficialOwner sets ProprietorIsBeneficialOwner field to given value.

### HasProprietorIsBeneficialOwner

`func (o *BusinessCardholder) HasProprietorIsBeneficialOwner() bool`

HasProprietorIsBeneficialOwner returns a boolean if a field has been set.

### GetProprietorOrOfficer

`func (o *BusinessCardholder) GetProprietorOrOfficer() BusinessProprietor`

GetProprietorOrOfficer returns the ProprietorOrOfficer field if non-nil, zero value otherwise.

### GetProprietorOrOfficerOk

`func (o *BusinessCardholder) GetProprietorOrOfficerOk() (*BusinessProprietor, bool)`

GetProprietorOrOfficerOk returns a tuple with the ProprietorOrOfficer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProprietorOrOfficer

`func (o *BusinessCardholder) SetProprietorOrOfficer(v BusinessProprietor)`

SetProprietorOrOfficer sets ProprietorOrOfficer field to given value.

### HasProprietorOrOfficer

`func (o *BusinessCardholder) HasProprietorOrOfficer() bool`

HasProprietorOrOfficer returns a boolean if a field has been set.

### GetTaxpayerId

`func (o *BusinessCardholder) GetTaxpayerId() string`

GetTaxpayerId returns the TaxpayerId field if non-nil, zero value otherwise.

### GetTaxpayerIdOk

`func (o *BusinessCardholder) GetTaxpayerIdOk() (*string, bool)`

GetTaxpayerIdOk returns a tuple with the TaxpayerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxpayerId

`func (o *BusinessCardholder) SetTaxpayerId(v string)`

SetTaxpayerId sets TaxpayerId field to given value.

### HasTaxpayerId

`func (o *BusinessCardholder) HasTaxpayerId() bool`

HasTaxpayerId returns a boolean if a field has been set.

### GetToken

`func (o *BusinessCardholder) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *BusinessCardholder) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *BusinessCardholder) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *BusinessCardholder) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetWebsite

`func (o *BusinessCardholder) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *BusinessCardholder) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *BusinessCardholder) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *BusinessCardholder) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


