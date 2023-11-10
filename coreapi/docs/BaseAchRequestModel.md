# BaseAchRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountNumber** | **string** | ACH account number. | 
**AccountType** | **string** | Type of account. | 
**BankName** | Pointer to **string** | Name of the bank holding the account. | [optional] 
**IsDefaultAccount** | Pointer to **bool** | If there are multiple funding sources, this field specifies which source is used by default in funding calls. If there is only one funding source, the system ignores this field and always uses that source. | [optional] [default to false]
**NameOnAccount** | **string** | Name on the ACH account. | 
**RoutingNumber** | **string** | Routing number for the ACH account. | [readonly] 
**Token** | Pointer to **string** | Unique identifier of the funding source. If you do not include a token, the system will generate one automatically. This token is necessary for use in other calls, so we recommend that rather than let the system generate one, you use a simple string that is easy to remember. This value cannot be updated. | [optional] 
**VerificationNotes** | Pointer to **string** | Free-form text field for holding notes about verification. This field is returned only if &#x60;verification_override &#x3D; true&#x60;. | [optional] 
**VerificationOverride** | Pointer to **bool** | Allows the ACH funding source to be used regardless of its verification status. | [optional] [default to false]

## Methods

### NewBaseAchRequestModel

`func NewBaseAchRequestModel(accountNumber string, accountType string, nameOnAccount string, routingNumber string, ) *BaseAchRequestModel`

NewBaseAchRequestModel instantiates a new BaseAchRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseAchRequestModelWithDefaults

`func NewBaseAchRequestModelWithDefaults() *BaseAchRequestModel`

NewBaseAchRequestModelWithDefaults instantiates a new BaseAchRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNumber

`func (o *BaseAchRequestModel) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *BaseAchRequestModel) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *BaseAchRequestModel) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.


### GetAccountType

`func (o *BaseAchRequestModel) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *BaseAchRequestModel) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *BaseAchRequestModel) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.


### GetBankName

`func (o *BaseAchRequestModel) GetBankName() string`

GetBankName returns the BankName field if non-nil, zero value otherwise.

### GetBankNameOk

`func (o *BaseAchRequestModel) GetBankNameOk() (*string, bool)`

GetBankNameOk returns a tuple with the BankName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankName

`func (o *BaseAchRequestModel) SetBankName(v string)`

SetBankName sets BankName field to given value.

### HasBankName

`func (o *BaseAchRequestModel) HasBankName() bool`

HasBankName returns a boolean if a field has been set.

### GetIsDefaultAccount

`func (o *BaseAchRequestModel) GetIsDefaultAccount() bool`

GetIsDefaultAccount returns the IsDefaultAccount field if non-nil, zero value otherwise.

### GetIsDefaultAccountOk

`func (o *BaseAchRequestModel) GetIsDefaultAccountOk() (*bool, bool)`

GetIsDefaultAccountOk returns a tuple with the IsDefaultAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefaultAccount

`func (o *BaseAchRequestModel) SetIsDefaultAccount(v bool)`

SetIsDefaultAccount sets IsDefaultAccount field to given value.

### HasIsDefaultAccount

`func (o *BaseAchRequestModel) HasIsDefaultAccount() bool`

HasIsDefaultAccount returns a boolean if a field has been set.

### GetNameOnAccount

`func (o *BaseAchRequestModel) GetNameOnAccount() string`

GetNameOnAccount returns the NameOnAccount field if non-nil, zero value otherwise.

### GetNameOnAccountOk

`func (o *BaseAchRequestModel) GetNameOnAccountOk() (*string, bool)`

GetNameOnAccountOk returns a tuple with the NameOnAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameOnAccount

`func (o *BaseAchRequestModel) SetNameOnAccount(v string)`

SetNameOnAccount sets NameOnAccount field to given value.


### GetRoutingNumber

`func (o *BaseAchRequestModel) GetRoutingNumber() string`

GetRoutingNumber returns the RoutingNumber field if non-nil, zero value otherwise.

### GetRoutingNumberOk

`func (o *BaseAchRequestModel) GetRoutingNumberOk() (*string, bool)`

GetRoutingNumberOk returns a tuple with the RoutingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingNumber

`func (o *BaseAchRequestModel) SetRoutingNumber(v string)`

SetRoutingNumber sets RoutingNumber field to given value.


### GetToken

`func (o *BaseAchRequestModel) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *BaseAchRequestModel) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *BaseAchRequestModel) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *BaseAchRequestModel) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetVerificationNotes

`func (o *BaseAchRequestModel) GetVerificationNotes() string`

GetVerificationNotes returns the VerificationNotes field if non-nil, zero value otherwise.

### GetVerificationNotesOk

`func (o *BaseAchRequestModel) GetVerificationNotesOk() (*string, bool)`

GetVerificationNotesOk returns a tuple with the VerificationNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationNotes

`func (o *BaseAchRequestModel) SetVerificationNotes(v string)`

SetVerificationNotes sets VerificationNotes field to given value.

### HasVerificationNotes

`func (o *BaseAchRequestModel) HasVerificationNotes() bool`

HasVerificationNotes returns a boolean if a field has been set.

### GetVerificationOverride

`func (o *BaseAchRequestModel) GetVerificationOverride() bool`

GetVerificationOverride returns the VerificationOverride field if non-nil, zero value otherwise.

### GetVerificationOverrideOk

`func (o *BaseAchRequestModel) GetVerificationOverrideOk() (*bool, bool)`

GetVerificationOverrideOk returns a tuple with the VerificationOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationOverride

`func (o *BaseAchRequestModel) SetVerificationOverride(v bool)`

SetVerificationOverride sets VerificationOverride field to given value.

### HasVerificationOverride

`func (o *BaseAchRequestModel) HasVerificationOverride() bool`

HasVerificationOverride returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


