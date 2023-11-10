/*
Core API

Marqeta's Core API endpoints, conveniently annotated to enable code generation (including SDKs), test cases, and documentation. Currently in beta.

API version: 3.0.11
Contact: support@marqeta.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package marqeta_coreapi_client

import (
	"encoding/json"
)

// checks if the CardHolderModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CardHolderModel{}

// CardHolderModel Contains information about a cardholder.
type CardHolderModel struct {
	// Associates the specified account holder group with the cardholder.  Send a `GET` request to `/accountholdergroups` to retrieve account holder group tokens.
	AccountHolderGroupToken *string `json:"account_holder_group_token,omitempty"`
	// Specifies if the cardholder is in the `ACTIVE` state on the Marqeta platform.  *NOTE:* Do not set the value of the `user.active` field directly. Instead, use the `/usertransitions` endpoints to transition user resources between statuses. For more information on status changes, see <</core-api/user-transitions#postUsertransitions, Create User Transition>>.
	Active *bool `json:"active,omitempty"`
	// Cardholder's address.  *NOTE:* Required for KYC verification (US-based cardholders only). Cannot perform KYC if set to a PO Box.
	Address1 *string `json:"address1,omitempty"`
	// Additional address information for the cardholder.  *NOTE:* Cannot perform KYC if set to a PO Box.
	Address2 *string `json:"address2,omitempty"`
	// Cardholder's date of birth.  *NOTE:* Required for KYC verification (US-based cardholders only).
	BirthDate *string `json:"birth_date,omitempty"`
	// City where the cardholder resides.  *NOTE:* Required for KYC verification (US-based cardholders only).
	City *string `json:"city,omitempty"`
	// Company name.
	Company *string `json:"company,omitempty"`
	// Specifies if the cardholder holds a corporate card.
	CorporateCardHolder *bool `json:"corporate_card_holder,omitempty"`
	// Country where the cardholder resides.  *NOTE:* Required for KYC verification (US-based cardholders only).
	Country *string `json:"country,omitempty"`
	// Valid email address of the cardholder.  This value must be unique among users.  *NOTE:* Required for KYC verification by certain banks (US-based cardholders only). To determine if you must provide an email address, contact your Marqeta representative.
	Email *string `json:"email,omitempty"`
	// Cardholder's first name.  *NOTE:* Required for KYC verification (US-based cardholders only).
	FirstName *string `json:"first_name,omitempty"`
	// Gender of the cardholder.
	Gender *string `json:"gender,omitempty"`
	// Cardholder's title or prefix: Dr., Miss, Mr., Ms., and so on.
	Honorific *string `json:"honorific,omitempty"`
	// Expiration date of the cardholder's identification card.
	IdCardExpirationDate *string `json:"id_card_expiration_date,omitempty"`
	// Cardholder's identification card number.
	IdCardNumber *string `json:"id_card_number,omitempty"`
	// One or more objects containing identifications associated with the cardholder.
	Identifications []IdentificationRequestModel `json:"identifications,omitempty"`
	// Cardholder's IP address.
	IpAddress *string `json:"ip_address,omitempty"`
	// Cardholder's last name.  *NOTE:* Required for KYC verification (US-based cardholders only).
	LastName *string `json:"last_name,omitempty"`
	// Associates any additional metadata you provide with the cardholder.
	Metadata *map[string]string `json:"metadata,omitempty"`
	// Cardholder's middle name.
	MiddleName *string `json:"middle_name,omitempty"`
	// Cardholder's nationality.
	Nationality *string `json:"nationality,omitempty"`
	// Any additional information pertaining to the cardholder.
	Notes *string `json:"notes,omitempty"`
	// Unique identifier of the parent user or business resource. Send a `GET` request to `/users` to retrieve user resource tokens or to `/businesses` to retrieve business resource tokens.  Required if `uses_parent_account = true`. This user or business is configured as the parent of the current user.
	ParentToken *string `json:"parent_token,omitempty"`
	// Expiration date of the cardholder's passport.
	PassportExpirationDate *string `json:"passport_expiration_date,omitempty"`
	// Cardholder's passport number.
	PassportNumber *string `json:"passport_number,omitempty"`
	// Password to the cardholder's user account on the Marqeta platform.  * Must contain at least one numeral + * Must contain at least one lowercase letter + * Must contain at least one uppercase letter + * Must contain at least one of these symbols: `@ # $ % ! ^ & * ( ) \\ _ + ~ ` - = [ ] { } , ; : ' \" , . / < > ?`
	Password *string `json:"password,omitempty"`
	// Telephone number of the cardholder (including area code), prepended by the `+` symbol and the 1- to 3-digit country calling code. Do not include hyphens, spaces, or parentheses.  *NOTE:* Required for KYC verification by certain banks (US-based cardholders only). To determine if you must provide a phone number, contact your Marqeta representative.
	Phone *string `json:"phone,omitempty"`
	// Postal code of the cardholder's address.  *NOTE:* Required for KYC verification (US-based cardholders only).
	PostalCode *string `json:"postal_code,omitempty"`
	// Cardholder's Social Security Number (SSN).
	Ssn *string `json:"ssn,omitempty"`
	// State or province where the cardholder resides.  *NOTE:* <</core-api/kyc-verification#_valid_state_provincial_and_territorial_abbreviations, Valid two-character abbreviation>> required for KYC verification (US-based cardholders only).
	State *string `json:"state,omitempty"`
	// Unique identifier of the cardholder. If you do not include a token, the system generates one automatically. This token is necessary for use in other API calls, so we recommend that rather than let the system generate one, you use a simple string that is easy to remember. This value cannot be updated.
	Token *string `json:"token,omitempty"`
	// Indicates whether the child shares balances with the parent (`true`), or the child's balances are independent of the parent (`false`).  If set to `true`, you must also include a `parent_token` in the request. This value cannot be updated.
	UsesParentAccount *bool `json:"uses_parent_account,omitempty"`
}

// NewCardHolderModel instantiates a new CardHolderModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardHolderModel() *CardHolderModel {
	this := CardHolderModel{}
	var active bool = true
	this.Active = &active
	var corporateCardHolder bool = false
	this.CorporateCardHolder = &corporateCardHolder
	var usesParentAccount bool = false
	this.UsesParentAccount = &usesParentAccount
	return &this
}

// NewCardHolderModelWithDefaults instantiates a new CardHolderModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardHolderModelWithDefaults() *CardHolderModel {
	this := CardHolderModel{}
	var active bool = true
	this.Active = &active
	var corporateCardHolder bool = false
	this.CorporateCardHolder = &corporateCardHolder
	var usesParentAccount bool = false
	this.UsesParentAccount = &usesParentAccount
	return &this
}

// GetAccountHolderGroupToken returns the AccountHolderGroupToken field value if set, zero value otherwise.
func (o *CardHolderModel) GetAccountHolderGroupToken() string {
	if o == nil || IsNil(o.AccountHolderGroupToken) {
		var ret string
		return ret
	}
	return *o.AccountHolderGroupToken
}

// GetAccountHolderGroupTokenOk returns a tuple with the AccountHolderGroupToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardHolderModel) GetAccountHolderGroupTokenOk() (*string, bool) {
	if o == nil || IsNil(o.AccountHolderGroupToken) {
		return nil, false
	}
	return o.AccountHolderGroupToken, true
}

// HasAccountHolderGroupToken returns a boolean if a field has been set.
func (o *CardHolderModel) HasAccountHolderGroupToken() bool {
	if o != nil && !IsNil(o.AccountHolderGroupToken) {
		return true
	}

	return false
}

// SetAccountHolderGroupToken gets a reference to the given string and assigns it to the AccountHolderGroupToken field.
func (o *CardHolderModel) SetAccountHolderGroupToken(v string) {
	o.AccountHolderGroupToken = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *CardHolderModel) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardHolderModel) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *CardHolderModel) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *CardHolderModel) SetActive(v bool) {
	o.Active = &v
}

// GetAddress1 returns the Address1 field value if set, zero value otherwise.
func (o *CardHolderModel) GetAddress1() string {
	if o == nil || IsNil(o.Address1) {
		var ret string
		return ret
	}
	return *o.Address1
}

// GetAddress1Ok returns a tuple with the Address1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardHolderModel) GetAddress1Ok() (*string, bool) {
	if o == nil || IsNil(o.Address1) {
		return nil, false
	}
	return o.Address1, true
}

// HasAddress1 returns a boolean if a field has been set.
func (o *CardHolderModel) HasAddress1() bool {
	if o != nil && !IsNil(o.Address1) {
		return true
	}

	return false
}

// SetAddress1 gets a reference to the given string and assigns it to the Address1 field.
func (o *CardHolderModel) SetAddress1(v string) {
	o.Address1 = &v
}

// GetAddress2 returns the Address2 field value if set, zero value otherwise.
func (o *CardHolderModel) GetAddress2() string {
	if o == nil || IsNil(o.Address2) {
		var ret string
		return ret
	}
	return *o.Address2
}

// GetAddress2Ok returns a tuple with the Address2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardHolderModel) GetAddress2Ok() (*string, bool) {
	if o == nil || IsNil(o.Address2) {
		return nil, false
	}
	return o.Address2, true
}

// HasAddress2 returns a boolean if a field has been set.
func (o *CardHolderModel) HasAddress2() bool {
	if o != nil && !IsNil(o.Address2) {
		return true
	}

	return false
}

// SetAddress2 gets a reference to the given string and assigns it to the Address2 field.
func (o *CardHolderModel) SetAddress2(v string) {
	o.Address2 = &v
}

// GetBirthDate returns the BirthDate field value if set, zero value otherwise.
func (o *CardHolderModel) GetBirthDate() string {
	if o == nil || IsNil(o.BirthDate) {
		var ret string
		return ret
	}
	return *o.BirthDate
}

// GetBirthDateOk returns a tuple with the BirthDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardHolderModel) GetBirthDateOk() (*string, bool) {
	if o == nil || IsNil(o.BirthDate) {
		return nil, false
	}
	return o.BirthDate, true
}

// HasBirthDate returns a boolean if a field has been set.
func (o *CardHolderModel) HasBirthDate() bool {
	if o != nil && !IsNil(o.BirthDate) {
		return true
	}

	return false
}

// SetBirthDate gets a reference to the given string and assigns it to the BirthDate field.
func (o *CardHolderModel) SetBirthDate(v string) {
	o.BirthDate = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *CardHolderModel) GetCity() string {
	if o == nil || IsNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardHolderModel) GetCityOk() (*string, bool) {
	if o == nil || IsNil(o.City) {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *CardHolderModel) HasCity() bool {
	if o != nil && !IsNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *CardHolderModel) SetCity(v string) {
	o.City = &v
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *CardHolderModel) GetCompany() string {
	if o == nil || IsNil(o.Company) {
		var ret string
		return ret
	}
	return *o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardHolderModel) GetCompanyOk() (*string, bool) {
	if o == nil || IsNil(o.Company) {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *CardHolderModel) HasCompany() bool {
	if o != nil && !IsNil(o.Company) {
		return true
	}

	return false
}

// SetCompany gets a reference to the given string and assigns it to the Company field.
func (o *CardHolderModel) SetCompany(v string) {
	o.Company = &v
}

// GetCorporateCardHolder returns the CorporateCardHolder field value if set, zero value otherwise.
func (o *CardHolderModel) GetCorporateCardHolder() bool {
	if o == nil || IsNil(o.CorporateCardHolder) {
		var ret bool
		return ret
	}
	return *o.CorporateCardHolder
}

// GetCorporateCardHolderOk returns a tuple with the CorporateCardHolder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardHolderModel) GetCorporateCardHolderOk() (*bool, bool) {
	if o == nil || IsNil(o.CorporateCardHolder) {
		return nil, false
	}
	return o.CorporateCardHolder, true
}

// HasCorporateCardHolder returns a boolean if a field has been set.
func (o *CardHolderModel) HasCorporateCardHolder() bool {
	if o != nil && !IsNil(o.CorporateCardHolder) {
		return true
	}

	return false
}

// SetCorporateCardHolder gets a reference to the given bool and assigns it to the CorporateCardHolder field.
func (o *CardHolderModel) SetCorporateCardHolder(v bool) {
	o.CorporateCardHolder = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *CardHolderModel) GetCountry() string {
	if o == nil || IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardHolderModel) GetCountryOk() (*string, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *CardHolderModel) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *CardHolderModel) SetCountry(v string) {
	o.Country = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *CardHolderModel) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardHolderModel) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *CardHolderModel) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *CardHolderModel) SetEmail(v string) {
	o.Email = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *CardHolderModel) GetFirstName() string {
	if o == nil || IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardHolderModel) GetFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *CardHolderModel) HasFirstName() bool {
	if o != nil && !IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *CardHolderModel) SetFirstName(v string) {
	o.FirstName = &v
}

// GetGender returns the Gender field value if set, zero value otherwise.
func (o *CardHolderModel) GetGender() string {
	if o == nil || IsNil(o.Gender) {
		var ret string
		return ret
	}
	return *o.Gender
}

// GetGenderOk returns a tuple with the Gender field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardHolderModel) GetGenderOk() (*string, bool) {
	if o == nil || IsNil(o.Gender) {
		return nil, false
	}
	return o.Gender, true
}

// HasGender returns a boolean if a field has been set.
func (o *CardHolderModel) HasGender() bool {
	if o != nil && !IsNil(o.Gender) {
		return true
	}

	return false
}

// SetGender gets a reference to the given string and assigns it to the Gender field.
func (o *CardHolderModel) SetGender(v string) {
	o.Gender = &v
}

// GetHonorific returns the Honorific field value if set, zero value otherwise.
func (o *CardHolderModel) GetHonorific() string {
	if o == nil || IsNil(o.Honorific) {
		var ret string
		return ret
	}
	return *o.Honorific
}

// GetHonorificOk returns a tuple with the Honorific field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardHolderModel) GetHonorificOk() (*string, bool) {
	if o == nil || IsNil(o.Honorific) {
		return nil, false
	}
	return o.Honorific, true
}

// HasHonorific returns a boolean if a field has been set.
func (o *CardHolderModel) HasHonorific() bool {
	if o != nil && !IsNil(o.Honorific) {
		return true
	}

	return false
}

// SetHonorific gets a reference to the given string and assigns it to the Honorific field.
func (o *CardHolderModel) SetHonorific(v string) {
	o.Honorific = &v
}

// GetIdCardExpirationDate returns the IdCardExpirationDate field value if set, zero value otherwise.
func (o *CardHolderModel) GetIdCardExpirationDate() string {
	if o == nil || IsNil(o.IdCardExpirationDate) {
		var ret string
		return ret
	}
	return *o.IdCardExpirationDate
}

// GetIdCardExpirationDateOk returns a tuple with the IdCardExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardHolderModel) GetIdCardExpirationDateOk() (*string, bool) {
	if o == nil || IsNil(o.IdCardExpirationDate) {
		return nil, false
	}
	return o.IdCardExpirationDate, true
}

// HasIdCardExpirationDate returns a boolean if a field has been set.
func (o *CardHolderModel) HasIdCardExpirationDate() bool {
	if o != nil && !IsNil(o.IdCardExpirationDate) {
		return true
	}

	return false
}

// SetIdCardExpirationDate gets a reference to the given string and assigns it to the IdCardExpirationDate field.
func (o *CardHolderModel) SetIdCardExpirationDate(v string) {
	o.IdCardExpirationDate = &v
}

// GetIdCardNumber returns the IdCardNumber field value if set, zero value otherwise.
func (o *CardHolderModel) GetIdCardNumber() string {
	if o == nil || IsNil(o.IdCardNumber) {
		var ret string
		return ret
	}
	return *o.IdCardNumber
}

// GetIdCardNumberOk returns a tuple with the IdCardNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardHolderModel) GetIdCardNumberOk() (*string, bool) {
	if o == nil || IsNil(o.IdCardNumber) {
		return nil, false
	}
	return o.IdCardNumber, true
}

// HasIdCardNumber returns a boolean if a field has been set.
func (o *CardHolderModel) HasIdCardNumber() bool {
	if o != nil && !IsNil(o.IdCardNumber) {
		return true
	}

	return false
}

// SetIdCardNumber gets a reference to the given string and assigns it to the IdCardNumber field.
func (o *CardHolderModel) SetIdCardNumber(v string) {
	o.IdCardNumber = &v
}

// GetIdentifications returns the Identifications field value if set, zero value otherwise.
func (o *CardHolderModel) GetIdentifications() []IdentificationRequestModel {
	if o == nil || IsNil(o.Identifications) {
		var ret []IdentificationRequestModel
		return ret
	}
	return o.Identifications
}

// GetIdentificationsOk returns a tuple with the Identifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardHolderModel) GetIdentificationsOk() ([]IdentificationRequestModel, bool) {
	if o == nil || IsNil(o.Identifications) {
		return nil, false
	}
	return o.Identifications, true
}

// HasIdentifications returns a boolean if a field has been set.
func (o *CardHolderModel) HasIdentifications() bool {
	if o != nil && !IsNil(o.Identifications) {
		return true
	}

	return false
}

// SetIdentifications gets a reference to the given []IdentificationRequestModel and assigns it to the Identifications field.
func (o *CardHolderModel) SetIdentifications(v []IdentificationRequestModel) {
	o.Identifications = v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *CardHolderModel) GetIpAddress() string {
	if o == nil || IsNil(o.IpAddress) {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardHolderModel) GetIpAddressOk() (*string, bool) {
	if o == nil || IsNil(o.IpAddress) {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *CardHolderModel) HasIpAddress() bool {
	if o != nil && !IsNil(o.IpAddress) {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *CardHolderModel) SetIpAddress(v string) {
	o.IpAddress = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *CardHolderModel) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardHolderModel) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *CardHolderModel) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *CardHolderModel) SetLastName(v string) {
	o.LastName = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *CardHolderModel) GetMetadata() map[string]string {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardHolderModel) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *CardHolderModel) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *CardHolderModel) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

// GetMiddleName returns the MiddleName field value if set, zero value otherwise.
func (o *CardHolderModel) GetMiddleName() string {
	if o == nil || IsNil(o.MiddleName) {
		var ret string
		return ret
	}
	return *o.MiddleName
}

// GetMiddleNameOk returns a tuple with the MiddleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardHolderModel) GetMiddleNameOk() (*string, bool) {
	if o == nil || IsNil(o.MiddleName) {
		return nil, false
	}
	return o.MiddleName, true
}

// HasMiddleName returns a boolean if a field has been set.
func (o *CardHolderModel) HasMiddleName() bool {
	if o != nil && !IsNil(o.MiddleName) {
		return true
	}

	return false
}

// SetMiddleName gets a reference to the given string and assigns it to the MiddleName field.
func (o *CardHolderModel) SetMiddleName(v string) {
	o.MiddleName = &v
}

// GetNationality returns the Nationality field value if set, zero value otherwise.
func (o *CardHolderModel) GetNationality() string {
	if o == nil || IsNil(o.Nationality) {
		var ret string
		return ret
	}
	return *o.Nationality
}

// GetNationalityOk returns a tuple with the Nationality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardHolderModel) GetNationalityOk() (*string, bool) {
	if o == nil || IsNil(o.Nationality) {
		return nil, false
	}
	return o.Nationality, true
}

// HasNationality returns a boolean if a field has been set.
func (o *CardHolderModel) HasNationality() bool {
	if o != nil && !IsNil(o.Nationality) {
		return true
	}

	return false
}

// SetNationality gets a reference to the given string and assigns it to the Nationality field.
func (o *CardHolderModel) SetNationality(v string) {
	o.Nationality = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *CardHolderModel) GetNotes() string {
	if o == nil || IsNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardHolderModel) GetNotesOk() (*string, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *CardHolderModel) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *CardHolderModel) SetNotes(v string) {
	o.Notes = &v
}

// GetParentToken returns the ParentToken field value if set, zero value otherwise.
func (o *CardHolderModel) GetParentToken() string {
	if o == nil || IsNil(o.ParentToken) {
		var ret string
		return ret
	}
	return *o.ParentToken
}

// GetParentTokenOk returns a tuple with the ParentToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardHolderModel) GetParentTokenOk() (*string, bool) {
	if o == nil || IsNil(o.ParentToken) {
		return nil, false
	}
	return o.ParentToken, true
}

// HasParentToken returns a boolean if a field has been set.
func (o *CardHolderModel) HasParentToken() bool {
	if o != nil && !IsNil(o.ParentToken) {
		return true
	}

	return false
}

// SetParentToken gets a reference to the given string and assigns it to the ParentToken field.
func (o *CardHolderModel) SetParentToken(v string) {
	o.ParentToken = &v
}

// GetPassportExpirationDate returns the PassportExpirationDate field value if set, zero value otherwise.
func (o *CardHolderModel) GetPassportExpirationDate() string {
	if o == nil || IsNil(o.PassportExpirationDate) {
		var ret string
		return ret
	}
	return *o.PassportExpirationDate
}

// GetPassportExpirationDateOk returns a tuple with the PassportExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardHolderModel) GetPassportExpirationDateOk() (*string, bool) {
	if o == nil || IsNil(o.PassportExpirationDate) {
		return nil, false
	}
	return o.PassportExpirationDate, true
}

// HasPassportExpirationDate returns a boolean if a field has been set.
func (o *CardHolderModel) HasPassportExpirationDate() bool {
	if o != nil && !IsNil(o.PassportExpirationDate) {
		return true
	}

	return false
}

// SetPassportExpirationDate gets a reference to the given string and assigns it to the PassportExpirationDate field.
func (o *CardHolderModel) SetPassportExpirationDate(v string) {
	o.PassportExpirationDate = &v
}

// GetPassportNumber returns the PassportNumber field value if set, zero value otherwise.
func (o *CardHolderModel) GetPassportNumber() string {
	if o == nil || IsNil(o.PassportNumber) {
		var ret string
		return ret
	}
	return *o.PassportNumber
}

// GetPassportNumberOk returns a tuple with the PassportNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardHolderModel) GetPassportNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PassportNumber) {
		return nil, false
	}
	return o.PassportNumber, true
}

// HasPassportNumber returns a boolean if a field has been set.
func (o *CardHolderModel) HasPassportNumber() bool {
	if o != nil && !IsNil(o.PassportNumber) {
		return true
	}

	return false
}

// SetPassportNumber gets a reference to the given string and assigns it to the PassportNumber field.
func (o *CardHolderModel) SetPassportNumber(v string) {
	o.PassportNumber = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *CardHolderModel) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardHolderModel) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *CardHolderModel) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *CardHolderModel) SetPassword(v string) {
	o.Password = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *CardHolderModel) GetPhone() string {
	if o == nil || IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardHolderModel) GetPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *CardHolderModel) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *CardHolderModel) SetPhone(v string) {
	o.Phone = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *CardHolderModel) GetPostalCode() string {
	if o == nil || IsNil(o.PostalCode) {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardHolderModel) GetPostalCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PostalCode) {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *CardHolderModel) HasPostalCode() bool {
	if o != nil && !IsNil(o.PostalCode) {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *CardHolderModel) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetSsn returns the Ssn field value if set, zero value otherwise.
func (o *CardHolderModel) GetSsn() string {
	if o == nil || IsNil(o.Ssn) {
		var ret string
		return ret
	}
	return *o.Ssn
}

// GetSsnOk returns a tuple with the Ssn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardHolderModel) GetSsnOk() (*string, bool) {
	if o == nil || IsNil(o.Ssn) {
		return nil, false
	}
	return o.Ssn, true
}

// HasSsn returns a boolean if a field has been set.
func (o *CardHolderModel) HasSsn() bool {
	if o != nil && !IsNil(o.Ssn) {
		return true
	}

	return false
}

// SetSsn gets a reference to the given string and assigns it to the Ssn field.
func (o *CardHolderModel) SetSsn(v string) {
	o.Ssn = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *CardHolderModel) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardHolderModel) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *CardHolderModel) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *CardHolderModel) SetState(v string) {
	o.State = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *CardHolderModel) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardHolderModel) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *CardHolderModel) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *CardHolderModel) SetToken(v string) {
	o.Token = &v
}

// GetUsesParentAccount returns the UsesParentAccount field value if set, zero value otherwise.
func (o *CardHolderModel) GetUsesParentAccount() bool {
	if o == nil || IsNil(o.UsesParentAccount) {
		var ret bool
		return ret
	}
	return *o.UsesParentAccount
}

// GetUsesParentAccountOk returns a tuple with the UsesParentAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardHolderModel) GetUsesParentAccountOk() (*bool, bool) {
	if o == nil || IsNil(o.UsesParentAccount) {
		return nil, false
	}
	return o.UsesParentAccount, true
}

// HasUsesParentAccount returns a boolean if a field has been set.
func (o *CardHolderModel) HasUsesParentAccount() bool {
	if o != nil && !IsNil(o.UsesParentAccount) {
		return true
	}

	return false
}

// SetUsesParentAccount gets a reference to the given bool and assigns it to the UsesParentAccount field.
func (o *CardHolderModel) SetUsesParentAccount(v bool) {
	o.UsesParentAccount = &v
}

func (o CardHolderModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CardHolderModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountHolderGroupToken) {
		toSerialize["account_holder_group_token"] = o.AccountHolderGroupToken
	}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.Address1) {
		toSerialize["address1"] = o.Address1
	}
	if !IsNil(o.Address2) {
		toSerialize["address2"] = o.Address2
	}
	if !IsNil(o.BirthDate) {
		toSerialize["birth_date"] = o.BirthDate
	}
	if !IsNil(o.City) {
		toSerialize["city"] = o.City
	}
	if !IsNil(o.Company) {
		toSerialize["company"] = o.Company
	}
	if !IsNil(o.CorporateCardHolder) {
		toSerialize["corporate_card_holder"] = o.CorporateCardHolder
	}
	if !IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.FirstName) {
		toSerialize["first_name"] = o.FirstName
	}
	if !IsNil(o.Gender) {
		toSerialize["gender"] = o.Gender
	}
	if !IsNil(o.Honorific) {
		toSerialize["honorific"] = o.Honorific
	}
	if !IsNil(o.IdCardExpirationDate) {
		toSerialize["id_card_expiration_date"] = o.IdCardExpirationDate
	}
	if !IsNil(o.IdCardNumber) {
		toSerialize["id_card_number"] = o.IdCardNumber
	}
	if !IsNil(o.Identifications) {
		toSerialize["identifications"] = o.Identifications
	}
	if !IsNil(o.IpAddress) {
		toSerialize["ip_address"] = o.IpAddress
	}
	if !IsNil(o.LastName) {
		toSerialize["last_name"] = o.LastName
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.MiddleName) {
		toSerialize["middle_name"] = o.MiddleName
	}
	if !IsNil(o.Nationality) {
		toSerialize["nationality"] = o.Nationality
	}
	if !IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	if !IsNil(o.ParentToken) {
		toSerialize["parent_token"] = o.ParentToken
	}
	if !IsNil(o.PassportExpirationDate) {
		toSerialize["passport_expiration_date"] = o.PassportExpirationDate
	}
	if !IsNil(o.PassportNumber) {
		toSerialize["passport_number"] = o.PassportNumber
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	if !IsNil(o.PostalCode) {
		toSerialize["postal_code"] = o.PostalCode
	}
	if !IsNil(o.Ssn) {
		toSerialize["ssn"] = o.Ssn
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	if !IsNil(o.UsesParentAccount) {
		toSerialize["uses_parent_account"] = o.UsesParentAccount
	}
	return toSerialize, nil
}

type NullableCardHolderModel struct {
	value *CardHolderModel
	isSet bool
}

func (v NullableCardHolderModel) Get() *CardHolderModel {
	return v.value
}

func (v *NullableCardHolderModel) Set(val *CardHolderModel) {
	v.value = val
	v.isSet = true
}

func (v NullableCardHolderModel) IsSet() bool {
	return v.isSet
}

func (v *NullableCardHolderModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardHolderModel(val *CardHolderModel) *NullableCardHolderModel {
	return &NullableCardHolderModel{value: val, isSet: true}
}

func (v NullableCardHolderModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardHolderModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


