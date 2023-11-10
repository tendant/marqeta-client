# AchModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountNumber** | **string** | ACH account number. | 
**AccountType** | **string** | Type of account. | 
**BankName** | Pointer to **string** | Name of the financial institution where the ACH account is held. | [optional] 
**BusinessToken** | Pointer to **string** | Unique identifier of the business account holder. This token is required if a &#x60;user_token&#x60; is not specified. | [optional] 
**IsDefaultAccount** | Pointer to **bool** | If there are multiple funding sources, this field specifies which source is used by default in funding calls. If there is only one funding source, the system ignores this field and always uses that source. | [optional] [default to false]
**NameOnAccount** | **string** | Name on the ACH account. | 
**RoutingNumber** | **string** | Routing number for the ACH account. | [readonly] 
**Token** | Pointer to **string** | Unique identifier of the funding source. If you do not include a token, the system will generate one automatically. This token is necessary for use in other calls, so we recommend that rather than let the system generate one, you use a simple string that is easy to remember. This value cannot be updated. | [optional] 
**UserToken** | Pointer to **string** | Unique identifier of the user account holder. This token is required if a &#x60;business_token&#x60; is not specified. | [optional] 
**VerificationNotes** | Pointer to **string** | Free-form text field for holding notes about verification. This field is returned only if &#x60;verification_override &#x3D; true&#x60;. | [optional] 
**VerificationOverride** | Pointer to **bool** | Allows the ACH funding source to be used regardless of its verification status. | [optional] [default to false]

## Methods

### NewAchModel

`func NewAchModel(accountNumber string, accountType string, nameOnAccount string, routingNumber string, ) *AchModel`

NewAchModel instantiates a new AchModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAchModelWithDefaults

`func NewAchModelWithDefaults() *AchModel`

NewAchModelWithDefaults instantiates a new AchModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNumber

`func (o *AchModel) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *AchModel) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *AchModel) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.


### GetAccountType

`func (o *AchModel) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *AchModel) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *AchModel) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.


### GetBankName

`func (o *AchModel) GetBankName() string`

GetBankName returns the BankName field if non-nil, zero value otherwise.

### GetBankNameOk

`func (o *AchModel) GetBankNameOk() (*string, bool)`

GetBankNameOk returns a tuple with the BankName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankName

`func (o *AchModel) SetBankName(v string)`

SetBankName sets BankName field to given value.

### HasBankName

`func (o *AchModel) HasBankName() bool`

HasBankName returns a boolean if a field has been set.

### GetBusinessToken

`func (o *AchModel) GetBusinessToken() string`

GetBusinessToken returns the BusinessToken field if non-nil, zero value otherwise.

### GetBusinessTokenOk

`func (o *AchModel) GetBusinessTokenOk() (*string, bool)`

GetBusinessTokenOk returns a tuple with the BusinessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessToken

`func (o *AchModel) SetBusinessToken(v string)`

SetBusinessToken sets BusinessToken field to given value.

### HasBusinessToken

`func (o *AchModel) HasBusinessToken() bool`

HasBusinessToken returns a boolean if a field has been set.

### GetIsDefaultAccount

`func (o *AchModel) GetIsDefaultAccount() bool`

GetIsDefaultAccount returns the IsDefaultAccount field if non-nil, zero value otherwise.

### GetIsDefaultAccountOk

`func (o *AchModel) GetIsDefaultAccountOk() (*bool, bool)`

GetIsDefaultAccountOk returns a tuple with the IsDefaultAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefaultAccount

`func (o *AchModel) SetIsDefaultAccount(v bool)`

SetIsDefaultAccount sets IsDefaultAccount field to given value.

### HasIsDefaultAccount

`func (o *AchModel) HasIsDefaultAccount() bool`

HasIsDefaultAccount returns a boolean if a field has been set.

### GetNameOnAccount

`func (o *AchModel) GetNameOnAccount() string`

GetNameOnAccount returns the NameOnAccount field if non-nil, zero value otherwise.

### GetNameOnAccountOk

`func (o *AchModel) GetNameOnAccountOk() (*string, bool)`

GetNameOnAccountOk returns a tuple with the NameOnAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameOnAccount

`func (o *AchModel) SetNameOnAccount(v string)`

SetNameOnAccount sets NameOnAccount field to given value.


### GetRoutingNumber

`func (o *AchModel) GetRoutingNumber() string`

GetRoutingNumber returns the RoutingNumber field if non-nil, zero value otherwise.

### GetRoutingNumberOk

`func (o *AchModel) GetRoutingNumberOk() (*string, bool)`

GetRoutingNumberOk returns a tuple with the RoutingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingNumber

`func (o *AchModel) SetRoutingNumber(v string)`

SetRoutingNumber sets RoutingNumber field to given value.


### GetToken

`func (o *AchModel) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AchModel) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AchModel) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *AchModel) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUserToken

`func (o *AchModel) GetUserToken() string`

GetUserToken returns the UserToken field if non-nil, zero value otherwise.

### GetUserTokenOk

`func (o *AchModel) GetUserTokenOk() (*string, bool)`

GetUserTokenOk returns a tuple with the UserToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserToken

`func (o *AchModel) SetUserToken(v string)`

SetUserToken sets UserToken field to given value.

### HasUserToken

`func (o *AchModel) HasUserToken() bool`

HasUserToken returns a boolean if a field has been set.

### GetVerificationNotes

`func (o *AchModel) GetVerificationNotes() string`

GetVerificationNotes returns the VerificationNotes field if non-nil, zero value otherwise.

### GetVerificationNotesOk

`func (o *AchModel) GetVerificationNotesOk() (*string, bool)`

GetVerificationNotesOk returns a tuple with the VerificationNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationNotes

`func (o *AchModel) SetVerificationNotes(v string)`

SetVerificationNotes sets VerificationNotes field to given value.

### HasVerificationNotes

`func (o *AchModel) HasVerificationNotes() bool`

HasVerificationNotes returns a boolean if a field has been set.

### GetVerificationOverride

`func (o *AchModel) GetVerificationOverride() bool`

GetVerificationOverride returns the VerificationOverride field if non-nil, zero value otherwise.

### GetVerificationOverrideOk

`func (o *AchModel) GetVerificationOverrideOk() (*bool, bool)`

GetVerificationOverrideOk returns a tuple with the VerificationOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationOverride

`func (o *AchModel) SetVerificationOverride(v bool)`

SetVerificationOverride sets VerificationOverride field to given value.

### HasVerificationOverride

`func (o *AchModel) HasVerificationOverride() bool`

HasVerificationOverride returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


