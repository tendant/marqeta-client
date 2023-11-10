# BusinessCardHolderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountHolderGroupToken** | Pointer to **string** | Associates the specified account holder group with the business.  This field is returned if it exists in the resource. | [optional] 
**Active** | Pointer to **bool** | Specifies if the business is in the &#x60;ACTIVE&#x60; state on the Marqeta platform.  This field is returned if it exists in the resource. | [optional] [default to false]
**AttestationConsent** | Pointer to **bool** | Indicates that the attester agrees that the information provided is correct and truthful.  This field is returned if it exists in the resource. | [optional] [default to false]
**AttestationDate** | Pointer to **time.Time** | Timestamp of the attestation.  This field is returned if it exists in the resource. | [optional] 
**AttesterName** | Pointer to **string** | Name of the attester for KYC verification.  This field is returned if it exists in the resource. | [optional] 
**AttesterTitle** | Pointer to **string** | Title of the attester for KYC verification.  This field is returned if it exists in the resource. | [optional] 
**Authentication** | Pointer to [**Authentication**](Authentication.md) |  | [optional] 
**BeneficialOwner1** | Pointer to [**BeneficialOwnerResponse**](BeneficialOwnerResponse.md) |  | [optional] 
**BeneficialOwner2** | Pointer to [**BeneficialOwnerResponse**](BeneficialOwnerResponse.md) |  | [optional] 
**BeneficialOwner3** | Pointer to [**BeneficialOwnerResponse**](BeneficialOwnerResponse.md) |  | [optional] 
**BeneficialOwner4** | Pointer to [**BeneficialOwnerResponse**](BeneficialOwnerResponse.md) |  | [optional] 
**BusinessNameDba** | Pointer to **string** | Fictitious business name (\&quot;Doing Business As\&quot; or DBA).  This field is returned if it exists in the resource. | [optional] 
**BusinessNameLegal** | Pointer to **string** | Legal name of the business.  This field is returned if it exists in the resource. | [optional] 
**BusinessType** | Pointer to **string** | Indicates the type of business, for example B2B (business-to-business) or B2C (business-to-consumer).  This field is returned if it exists in the resource. | [optional] 
**CreatedTime** | **time.Time** | Date and time when the business was created, in UTC. | 
**DateEstablished** | Pointer to **time.Time** | Date and time when the business was established.  This field is returned if it exists in the resource. | [optional] 
**DunsNumber** | Pointer to **string** | Data Universal Numbering System (DUNS) number of the business.  This field is returned if it exists in the resource. | [optional] 
**GeneralBusinessDescription** | Pointer to **string** | General description of the business.  This field is returned if it exists in the resource. | [optional] 
**History** | Pointer to **string** | History of the business.  This field is returned if it exists in the resource. | [optional] 
**Identifications** | Pointer to [**[]IdentificationResponseModel**](IdentificationResponseModel.md) | One or more objects containing identifications associated with the business.  Objects are returned if they exist in the resource. | [optional] 
**InCurrentLocationSince** | Pointer to **time.Time** | Date on which the business office opened in its current location.  This field is returned if it exists in the resource. | [optional] 
**Incorporation** | Pointer to [**BusinessIncorporationResponse**](BusinessIncorporationResponse.md) |  | [optional] 
**InternationalOfficeLocations** | Pointer to **string** | Locations of the business&#39; offices outside the US.  This field is returned if it exists in the resource. | [optional] 
**IpAddress** | Pointer to **string** | IP address of the business.  This field is returned if it exists in the resource. | [optional] 
**LastModifiedTime** | **time.Time** | Date and time when the business was last modified, in UTC. | 
**Metadata** | Pointer to **map[string]string** | Associates any additional metadata you provide with the business.  Metadata is returned if it exists in the resource. | [optional] 
**Notes** | Pointer to **string** | Any additional information pertaining to the business.  This field is returned if it exists in the resource. | [optional] 
**OfficeLocation** | Pointer to [**AddressResponseModel**](AddressResponseModel.md) |  | [optional] 
**Password** | Pointer to **string** | Password for the business account on the Marqeta platform.  This field is returned if it exists in the resource. | [optional] 
**Phone** | Pointer to **string** | 10-digit telephone number of the business.  This field is returned if it exists in the resource. | [optional] 
**PrimaryContact** | Pointer to [**PrimaryContactInfoModel**](PrimaryContactInfoModel.md) |  | [optional] 
**ProprietorIsBeneficialOwner** | Pointer to **bool** | Indicates that the proprietor or officer of the business is also a beneficial owner.  This field is returned if it exists in the resource. | [optional] [default to false]
**ProprietorOrOfficer** | Pointer to [**BusinessProprietorResponse**](BusinessProprietorResponse.md) |  | [optional] 
**Status** | Pointer to **string** | Specifies the state of the business on the Marqeta platform.  This field is returned if it exists in the resource. | [optional] 
**TaxpayerId** | Pointer to **string** | Taxpayer identifier of the business.  This field is returned if it exists in the resource. | [optional] 
**Token** | Pointer to **string** | Unique identifier of the business resource.  This field is always returned. | [optional] 
**Website** | Pointer to **string** | URL of the business&#39; website.  This field is returned if it exists in the resource. | [optional] 

## Methods

### NewBusinessCardHolderResponse

`func NewBusinessCardHolderResponse(createdTime time.Time, lastModifiedTime time.Time, ) *BusinessCardHolderResponse`

NewBusinessCardHolderResponse instantiates a new BusinessCardHolderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBusinessCardHolderResponseWithDefaults

`func NewBusinessCardHolderResponseWithDefaults() *BusinessCardHolderResponse`

NewBusinessCardHolderResponseWithDefaults instantiates a new BusinessCardHolderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountHolderGroupToken

`func (o *BusinessCardHolderResponse) GetAccountHolderGroupToken() string`

GetAccountHolderGroupToken returns the AccountHolderGroupToken field if non-nil, zero value otherwise.

### GetAccountHolderGroupTokenOk

`func (o *BusinessCardHolderResponse) GetAccountHolderGroupTokenOk() (*string, bool)`

GetAccountHolderGroupTokenOk returns a tuple with the AccountHolderGroupToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountHolderGroupToken

`func (o *BusinessCardHolderResponse) SetAccountHolderGroupToken(v string)`

SetAccountHolderGroupToken sets AccountHolderGroupToken field to given value.

### HasAccountHolderGroupToken

`func (o *BusinessCardHolderResponse) HasAccountHolderGroupToken() bool`

HasAccountHolderGroupToken returns a boolean if a field has been set.

### GetActive

`func (o *BusinessCardHolderResponse) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *BusinessCardHolderResponse) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *BusinessCardHolderResponse) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *BusinessCardHolderResponse) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAttestationConsent

`func (o *BusinessCardHolderResponse) GetAttestationConsent() bool`

GetAttestationConsent returns the AttestationConsent field if non-nil, zero value otherwise.

### GetAttestationConsentOk

`func (o *BusinessCardHolderResponse) GetAttestationConsentOk() (*bool, bool)`

GetAttestationConsentOk returns a tuple with the AttestationConsent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttestationConsent

`func (o *BusinessCardHolderResponse) SetAttestationConsent(v bool)`

SetAttestationConsent sets AttestationConsent field to given value.

### HasAttestationConsent

`func (o *BusinessCardHolderResponse) HasAttestationConsent() bool`

HasAttestationConsent returns a boolean if a field has been set.

### GetAttestationDate

`func (o *BusinessCardHolderResponse) GetAttestationDate() time.Time`

GetAttestationDate returns the AttestationDate field if non-nil, zero value otherwise.

### GetAttestationDateOk

`func (o *BusinessCardHolderResponse) GetAttestationDateOk() (*time.Time, bool)`

GetAttestationDateOk returns a tuple with the AttestationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttestationDate

`func (o *BusinessCardHolderResponse) SetAttestationDate(v time.Time)`

SetAttestationDate sets AttestationDate field to given value.

### HasAttestationDate

`func (o *BusinessCardHolderResponse) HasAttestationDate() bool`

HasAttestationDate returns a boolean if a field has been set.

### GetAttesterName

`func (o *BusinessCardHolderResponse) GetAttesterName() string`

GetAttesterName returns the AttesterName field if non-nil, zero value otherwise.

### GetAttesterNameOk

`func (o *BusinessCardHolderResponse) GetAttesterNameOk() (*string, bool)`

GetAttesterNameOk returns a tuple with the AttesterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttesterName

`func (o *BusinessCardHolderResponse) SetAttesterName(v string)`

SetAttesterName sets AttesterName field to given value.

### HasAttesterName

`func (o *BusinessCardHolderResponse) HasAttesterName() bool`

HasAttesterName returns a boolean if a field has been set.

### GetAttesterTitle

`func (o *BusinessCardHolderResponse) GetAttesterTitle() string`

GetAttesterTitle returns the AttesterTitle field if non-nil, zero value otherwise.

### GetAttesterTitleOk

`func (o *BusinessCardHolderResponse) GetAttesterTitleOk() (*string, bool)`

GetAttesterTitleOk returns a tuple with the AttesterTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttesterTitle

`func (o *BusinessCardHolderResponse) SetAttesterTitle(v string)`

SetAttesterTitle sets AttesterTitle field to given value.

### HasAttesterTitle

`func (o *BusinessCardHolderResponse) HasAttesterTitle() bool`

HasAttesterTitle returns a boolean if a field has been set.

### GetAuthentication

`func (o *BusinessCardHolderResponse) GetAuthentication() Authentication`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *BusinessCardHolderResponse) GetAuthenticationOk() (*Authentication, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *BusinessCardHolderResponse) SetAuthentication(v Authentication)`

SetAuthentication sets Authentication field to given value.

### HasAuthentication

`func (o *BusinessCardHolderResponse) HasAuthentication() bool`

HasAuthentication returns a boolean if a field has been set.

### GetBeneficialOwner1

`func (o *BusinessCardHolderResponse) GetBeneficialOwner1() BeneficialOwnerResponse`

GetBeneficialOwner1 returns the BeneficialOwner1 field if non-nil, zero value otherwise.

### GetBeneficialOwner1Ok

`func (o *BusinessCardHolderResponse) GetBeneficialOwner1Ok() (*BeneficialOwnerResponse, bool)`

GetBeneficialOwner1Ok returns a tuple with the BeneficialOwner1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficialOwner1

`func (o *BusinessCardHolderResponse) SetBeneficialOwner1(v BeneficialOwnerResponse)`

SetBeneficialOwner1 sets BeneficialOwner1 field to given value.

### HasBeneficialOwner1

`func (o *BusinessCardHolderResponse) HasBeneficialOwner1() bool`

HasBeneficialOwner1 returns a boolean if a field has been set.

### GetBeneficialOwner2

`func (o *BusinessCardHolderResponse) GetBeneficialOwner2() BeneficialOwnerResponse`

GetBeneficialOwner2 returns the BeneficialOwner2 field if non-nil, zero value otherwise.

### GetBeneficialOwner2Ok

`func (o *BusinessCardHolderResponse) GetBeneficialOwner2Ok() (*BeneficialOwnerResponse, bool)`

GetBeneficialOwner2Ok returns a tuple with the BeneficialOwner2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficialOwner2

`func (o *BusinessCardHolderResponse) SetBeneficialOwner2(v BeneficialOwnerResponse)`

SetBeneficialOwner2 sets BeneficialOwner2 field to given value.

### HasBeneficialOwner2

`func (o *BusinessCardHolderResponse) HasBeneficialOwner2() bool`

HasBeneficialOwner2 returns a boolean if a field has been set.

### GetBeneficialOwner3

`func (o *BusinessCardHolderResponse) GetBeneficialOwner3() BeneficialOwnerResponse`

GetBeneficialOwner3 returns the BeneficialOwner3 field if non-nil, zero value otherwise.

### GetBeneficialOwner3Ok

`func (o *BusinessCardHolderResponse) GetBeneficialOwner3Ok() (*BeneficialOwnerResponse, bool)`

GetBeneficialOwner3Ok returns a tuple with the BeneficialOwner3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficialOwner3

`func (o *BusinessCardHolderResponse) SetBeneficialOwner3(v BeneficialOwnerResponse)`

SetBeneficialOwner3 sets BeneficialOwner3 field to given value.

### HasBeneficialOwner3

`func (o *BusinessCardHolderResponse) HasBeneficialOwner3() bool`

HasBeneficialOwner3 returns a boolean if a field has been set.

### GetBeneficialOwner4

`func (o *BusinessCardHolderResponse) GetBeneficialOwner4() BeneficialOwnerResponse`

GetBeneficialOwner4 returns the BeneficialOwner4 field if non-nil, zero value otherwise.

### GetBeneficialOwner4Ok

`func (o *BusinessCardHolderResponse) GetBeneficialOwner4Ok() (*BeneficialOwnerResponse, bool)`

GetBeneficialOwner4Ok returns a tuple with the BeneficialOwner4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficialOwner4

`func (o *BusinessCardHolderResponse) SetBeneficialOwner4(v BeneficialOwnerResponse)`

SetBeneficialOwner4 sets BeneficialOwner4 field to given value.

### HasBeneficialOwner4

`func (o *BusinessCardHolderResponse) HasBeneficialOwner4() bool`

HasBeneficialOwner4 returns a boolean if a field has been set.

### GetBusinessNameDba

`func (o *BusinessCardHolderResponse) GetBusinessNameDba() string`

GetBusinessNameDba returns the BusinessNameDba field if non-nil, zero value otherwise.

### GetBusinessNameDbaOk

`func (o *BusinessCardHolderResponse) GetBusinessNameDbaOk() (*string, bool)`

GetBusinessNameDbaOk returns a tuple with the BusinessNameDba field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessNameDba

`func (o *BusinessCardHolderResponse) SetBusinessNameDba(v string)`

SetBusinessNameDba sets BusinessNameDba field to given value.

### HasBusinessNameDba

`func (o *BusinessCardHolderResponse) HasBusinessNameDba() bool`

HasBusinessNameDba returns a boolean if a field has been set.

### GetBusinessNameLegal

`func (o *BusinessCardHolderResponse) GetBusinessNameLegal() string`

GetBusinessNameLegal returns the BusinessNameLegal field if non-nil, zero value otherwise.

### GetBusinessNameLegalOk

`func (o *BusinessCardHolderResponse) GetBusinessNameLegalOk() (*string, bool)`

GetBusinessNameLegalOk returns a tuple with the BusinessNameLegal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessNameLegal

`func (o *BusinessCardHolderResponse) SetBusinessNameLegal(v string)`

SetBusinessNameLegal sets BusinessNameLegal field to given value.

### HasBusinessNameLegal

`func (o *BusinessCardHolderResponse) HasBusinessNameLegal() bool`

HasBusinessNameLegal returns a boolean if a field has been set.

### GetBusinessType

`func (o *BusinessCardHolderResponse) GetBusinessType() string`

GetBusinessType returns the BusinessType field if non-nil, zero value otherwise.

### GetBusinessTypeOk

`func (o *BusinessCardHolderResponse) GetBusinessTypeOk() (*string, bool)`

GetBusinessTypeOk returns a tuple with the BusinessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessType

`func (o *BusinessCardHolderResponse) SetBusinessType(v string)`

SetBusinessType sets BusinessType field to given value.

### HasBusinessType

`func (o *BusinessCardHolderResponse) HasBusinessType() bool`

HasBusinessType returns a boolean if a field has been set.

### GetCreatedTime

`func (o *BusinessCardHolderResponse) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *BusinessCardHolderResponse) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *BusinessCardHolderResponse) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetDateEstablished

`func (o *BusinessCardHolderResponse) GetDateEstablished() time.Time`

GetDateEstablished returns the DateEstablished field if non-nil, zero value otherwise.

### GetDateEstablishedOk

`func (o *BusinessCardHolderResponse) GetDateEstablishedOk() (*time.Time, bool)`

GetDateEstablishedOk returns a tuple with the DateEstablished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateEstablished

`func (o *BusinessCardHolderResponse) SetDateEstablished(v time.Time)`

SetDateEstablished sets DateEstablished field to given value.

### HasDateEstablished

`func (o *BusinessCardHolderResponse) HasDateEstablished() bool`

HasDateEstablished returns a boolean if a field has been set.

### GetDunsNumber

`func (o *BusinessCardHolderResponse) GetDunsNumber() string`

GetDunsNumber returns the DunsNumber field if non-nil, zero value otherwise.

### GetDunsNumberOk

`func (o *BusinessCardHolderResponse) GetDunsNumberOk() (*string, bool)`

GetDunsNumberOk returns a tuple with the DunsNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDunsNumber

`func (o *BusinessCardHolderResponse) SetDunsNumber(v string)`

SetDunsNumber sets DunsNumber field to given value.

### HasDunsNumber

`func (o *BusinessCardHolderResponse) HasDunsNumber() bool`

HasDunsNumber returns a boolean if a field has been set.

### GetGeneralBusinessDescription

`func (o *BusinessCardHolderResponse) GetGeneralBusinessDescription() string`

GetGeneralBusinessDescription returns the GeneralBusinessDescription field if non-nil, zero value otherwise.

### GetGeneralBusinessDescriptionOk

`func (o *BusinessCardHolderResponse) GetGeneralBusinessDescriptionOk() (*string, bool)`

GetGeneralBusinessDescriptionOk returns a tuple with the GeneralBusinessDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneralBusinessDescription

`func (o *BusinessCardHolderResponse) SetGeneralBusinessDescription(v string)`

SetGeneralBusinessDescription sets GeneralBusinessDescription field to given value.

### HasGeneralBusinessDescription

`func (o *BusinessCardHolderResponse) HasGeneralBusinessDescription() bool`

HasGeneralBusinessDescription returns a boolean if a field has been set.

### GetHistory

`func (o *BusinessCardHolderResponse) GetHistory() string`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *BusinessCardHolderResponse) GetHistoryOk() (*string, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *BusinessCardHolderResponse) SetHistory(v string)`

SetHistory sets History field to given value.

### HasHistory

`func (o *BusinessCardHolderResponse) HasHistory() bool`

HasHistory returns a boolean if a field has been set.

### GetIdentifications

`func (o *BusinessCardHolderResponse) GetIdentifications() []IdentificationResponseModel`

GetIdentifications returns the Identifications field if non-nil, zero value otherwise.

### GetIdentificationsOk

`func (o *BusinessCardHolderResponse) GetIdentificationsOk() (*[]IdentificationResponseModel, bool)`

GetIdentificationsOk returns a tuple with the Identifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifications

`func (o *BusinessCardHolderResponse) SetIdentifications(v []IdentificationResponseModel)`

SetIdentifications sets Identifications field to given value.

### HasIdentifications

`func (o *BusinessCardHolderResponse) HasIdentifications() bool`

HasIdentifications returns a boolean if a field has been set.

### GetInCurrentLocationSince

`func (o *BusinessCardHolderResponse) GetInCurrentLocationSince() time.Time`

GetInCurrentLocationSince returns the InCurrentLocationSince field if non-nil, zero value otherwise.

### GetInCurrentLocationSinceOk

`func (o *BusinessCardHolderResponse) GetInCurrentLocationSinceOk() (*time.Time, bool)`

GetInCurrentLocationSinceOk returns a tuple with the InCurrentLocationSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInCurrentLocationSince

`func (o *BusinessCardHolderResponse) SetInCurrentLocationSince(v time.Time)`

SetInCurrentLocationSince sets InCurrentLocationSince field to given value.

### HasInCurrentLocationSince

`func (o *BusinessCardHolderResponse) HasInCurrentLocationSince() bool`

HasInCurrentLocationSince returns a boolean if a field has been set.

### GetIncorporation

`func (o *BusinessCardHolderResponse) GetIncorporation() BusinessIncorporationResponse`

GetIncorporation returns the Incorporation field if non-nil, zero value otherwise.

### GetIncorporationOk

`func (o *BusinessCardHolderResponse) GetIncorporationOk() (*BusinessIncorporationResponse, bool)`

GetIncorporationOk returns a tuple with the Incorporation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncorporation

`func (o *BusinessCardHolderResponse) SetIncorporation(v BusinessIncorporationResponse)`

SetIncorporation sets Incorporation field to given value.

### HasIncorporation

`func (o *BusinessCardHolderResponse) HasIncorporation() bool`

HasIncorporation returns a boolean if a field has been set.

### GetInternationalOfficeLocations

`func (o *BusinessCardHolderResponse) GetInternationalOfficeLocations() string`

GetInternationalOfficeLocations returns the InternationalOfficeLocations field if non-nil, zero value otherwise.

### GetInternationalOfficeLocationsOk

`func (o *BusinessCardHolderResponse) GetInternationalOfficeLocationsOk() (*string, bool)`

GetInternationalOfficeLocationsOk returns a tuple with the InternationalOfficeLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternationalOfficeLocations

`func (o *BusinessCardHolderResponse) SetInternationalOfficeLocations(v string)`

SetInternationalOfficeLocations sets InternationalOfficeLocations field to given value.

### HasInternationalOfficeLocations

`func (o *BusinessCardHolderResponse) HasInternationalOfficeLocations() bool`

HasInternationalOfficeLocations returns a boolean if a field has been set.

### GetIpAddress

`func (o *BusinessCardHolderResponse) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *BusinessCardHolderResponse) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *BusinessCardHolderResponse) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *BusinessCardHolderResponse) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetLastModifiedTime

`func (o *BusinessCardHolderResponse) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *BusinessCardHolderResponse) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *BusinessCardHolderResponse) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.


### GetMetadata

`func (o *BusinessCardHolderResponse) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *BusinessCardHolderResponse) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *BusinessCardHolderResponse) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *BusinessCardHolderResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetNotes

`func (o *BusinessCardHolderResponse) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *BusinessCardHolderResponse) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *BusinessCardHolderResponse) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *BusinessCardHolderResponse) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetOfficeLocation

`func (o *BusinessCardHolderResponse) GetOfficeLocation() AddressResponseModel`

GetOfficeLocation returns the OfficeLocation field if non-nil, zero value otherwise.

### GetOfficeLocationOk

`func (o *BusinessCardHolderResponse) GetOfficeLocationOk() (*AddressResponseModel, bool)`

GetOfficeLocationOk returns a tuple with the OfficeLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficeLocation

`func (o *BusinessCardHolderResponse) SetOfficeLocation(v AddressResponseModel)`

SetOfficeLocation sets OfficeLocation field to given value.

### HasOfficeLocation

`func (o *BusinessCardHolderResponse) HasOfficeLocation() bool`

HasOfficeLocation returns a boolean if a field has been set.

### GetPassword

`func (o *BusinessCardHolderResponse) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *BusinessCardHolderResponse) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *BusinessCardHolderResponse) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *BusinessCardHolderResponse) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPhone

`func (o *BusinessCardHolderResponse) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *BusinessCardHolderResponse) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *BusinessCardHolderResponse) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *BusinessCardHolderResponse) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetPrimaryContact

`func (o *BusinessCardHolderResponse) GetPrimaryContact() PrimaryContactInfoModel`

GetPrimaryContact returns the PrimaryContact field if non-nil, zero value otherwise.

### GetPrimaryContactOk

`func (o *BusinessCardHolderResponse) GetPrimaryContactOk() (*PrimaryContactInfoModel, bool)`

GetPrimaryContactOk returns a tuple with the PrimaryContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryContact

`func (o *BusinessCardHolderResponse) SetPrimaryContact(v PrimaryContactInfoModel)`

SetPrimaryContact sets PrimaryContact field to given value.

### HasPrimaryContact

`func (o *BusinessCardHolderResponse) HasPrimaryContact() bool`

HasPrimaryContact returns a boolean if a field has been set.

### GetProprietorIsBeneficialOwner

`func (o *BusinessCardHolderResponse) GetProprietorIsBeneficialOwner() bool`

GetProprietorIsBeneficialOwner returns the ProprietorIsBeneficialOwner field if non-nil, zero value otherwise.

### GetProprietorIsBeneficialOwnerOk

`func (o *BusinessCardHolderResponse) GetProprietorIsBeneficialOwnerOk() (*bool, bool)`

GetProprietorIsBeneficialOwnerOk returns a tuple with the ProprietorIsBeneficialOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProprietorIsBeneficialOwner

`func (o *BusinessCardHolderResponse) SetProprietorIsBeneficialOwner(v bool)`

SetProprietorIsBeneficialOwner sets ProprietorIsBeneficialOwner field to given value.

### HasProprietorIsBeneficialOwner

`func (o *BusinessCardHolderResponse) HasProprietorIsBeneficialOwner() bool`

HasProprietorIsBeneficialOwner returns a boolean if a field has been set.

### GetProprietorOrOfficer

`func (o *BusinessCardHolderResponse) GetProprietorOrOfficer() BusinessProprietorResponse`

GetProprietorOrOfficer returns the ProprietorOrOfficer field if non-nil, zero value otherwise.

### GetProprietorOrOfficerOk

`func (o *BusinessCardHolderResponse) GetProprietorOrOfficerOk() (*BusinessProprietorResponse, bool)`

GetProprietorOrOfficerOk returns a tuple with the ProprietorOrOfficer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProprietorOrOfficer

`func (o *BusinessCardHolderResponse) SetProprietorOrOfficer(v BusinessProprietorResponse)`

SetProprietorOrOfficer sets ProprietorOrOfficer field to given value.

### HasProprietorOrOfficer

`func (o *BusinessCardHolderResponse) HasProprietorOrOfficer() bool`

HasProprietorOrOfficer returns a boolean if a field has been set.

### GetStatus

`func (o *BusinessCardHolderResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BusinessCardHolderResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BusinessCardHolderResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BusinessCardHolderResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTaxpayerId

`func (o *BusinessCardHolderResponse) GetTaxpayerId() string`

GetTaxpayerId returns the TaxpayerId field if non-nil, zero value otherwise.

### GetTaxpayerIdOk

`func (o *BusinessCardHolderResponse) GetTaxpayerIdOk() (*string, bool)`

GetTaxpayerIdOk returns a tuple with the TaxpayerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxpayerId

`func (o *BusinessCardHolderResponse) SetTaxpayerId(v string)`

SetTaxpayerId sets TaxpayerId field to given value.

### HasTaxpayerId

`func (o *BusinessCardHolderResponse) HasTaxpayerId() bool`

HasTaxpayerId returns a boolean if a field has been set.

### GetToken

`func (o *BusinessCardHolderResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *BusinessCardHolderResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *BusinessCardHolderResponse) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *BusinessCardHolderResponse) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetWebsite

`func (o *BusinessCardHolderResponse) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *BusinessCardHolderResponse) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *BusinessCardHolderResponse) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *BusinessCardHolderResponse) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


