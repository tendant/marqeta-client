# BankAccountFundingSourceModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSuffix** | **string** |  | 
**AccountType** | **string** |  | 
**BusinessToken** | Pointer to **string** | Required if &#39;user_token&#39; is null | [optional] 
**NameOnAccount** | **string** |  | 
**RoutingNumber** | **string** |  | 
**UserToken** | Pointer to **string** | Required if &#39;business_token&#39; is null | [optional] 
**VerificationStatus** | **string** |  | 

## Methods

### NewBankAccountFundingSourceModel

`func NewBankAccountFundingSourceModel(accountSuffix string, accountType string, nameOnAccount string, routingNumber string, verificationStatus string, ) *BankAccountFundingSourceModel`

NewBankAccountFundingSourceModel instantiates a new BankAccountFundingSourceModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankAccountFundingSourceModelWithDefaults

`func NewBankAccountFundingSourceModelWithDefaults() *BankAccountFundingSourceModel`

NewBankAccountFundingSourceModelWithDefaults instantiates a new BankAccountFundingSourceModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSuffix

`func (o *BankAccountFundingSourceModel) GetAccountSuffix() string`

GetAccountSuffix returns the AccountSuffix field if non-nil, zero value otherwise.

### GetAccountSuffixOk

`func (o *BankAccountFundingSourceModel) GetAccountSuffixOk() (*string, bool)`

GetAccountSuffixOk returns a tuple with the AccountSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSuffix

`func (o *BankAccountFundingSourceModel) SetAccountSuffix(v string)`

SetAccountSuffix sets AccountSuffix field to given value.


### GetAccountType

`func (o *BankAccountFundingSourceModel) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *BankAccountFundingSourceModel) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *BankAccountFundingSourceModel) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.


### GetBusinessToken

`func (o *BankAccountFundingSourceModel) GetBusinessToken() string`

GetBusinessToken returns the BusinessToken field if non-nil, zero value otherwise.

### GetBusinessTokenOk

`func (o *BankAccountFundingSourceModel) GetBusinessTokenOk() (*string, bool)`

GetBusinessTokenOk returns a tuple with the BusinessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessToken

`func (o *BankAccountFundingSourceModel) SetBusinessToken(v string)`

SetBusinessToken sets BusinessToken field to given value.

### HasBusinessToken

`func (o *BankAccountFundingSourceModel) HasBusinessToken() bool`

HasBusinessToken returns a boolean if a field has been set.

### GetNameOnAccount

`func (o *BankAccountFundingSourceModel) GetNameOnAccount() string`

GetNameOnAccount returns the NameOnAccount field if non-nil, zero value otherwise.

### GetNameOnAccountOk

`func (o *BankAccountFundingSourceModel) GetNameOnAccountOk() (*string, bool)`

GetNameOnAccountOk returns a tuple with the NameOnAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameOnAccount

`func (o *BankAccountFundingSourceModel) SetNameOnAccount(v string)`

SetNameOnAccount sets NameOnAccount field to given value.


### GetRoutingNumber

`func (o *BankAccountFundingSourceModel) GetRoutingNumber() string`

GetRoutingNumber returns the RoutingNumber field if non-nil, zero value otherwise.

### GetRoutingNumberOk

`func (o *BankAccountFundingSourceModel) GetRoutingNumberOk() (*string, bool)`

GetRoutingNumberOk returns a tuple with the RoutingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingNumber

`func (o *BankAccountFundingSourceModel) SetRoutingNumber(v string)`

SetRoutingNumber sets RoutingNumber field to given value.


### GetUserToken

`func (o *BankAccountFundingSourceModel) GetUserToken() string`

GetUserToken returns the UserToken field if non-nil, zero value otherwise.

### GetUserTokenOk

`func (o *BankAccountFundingSourceModel) GetUserTokenOk() (*string, bool)`

GetUserTokenOk returns a tuple with the UserToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserToken

`func (o *BankAccountFundingSourceModel) SetUserToken(v string)`

SetUserToken sets UserToken field to given value.

### HasUserToken

`func (o *BankAccountFundingSourceModel) HasUserToken() bool`

HasUserToken returns a boolean if a field has been set.

### GetVerificationStatus

`func (o *BankAccountFundingSourceModel) GetVerificationStatus() string`

GetVerificationStatus returns the VerificationStatus field if non-nil, zero value otherwise.

### GetVerificationStatusOk

`func (o *BankAccountFundingSourceModel) GetVerificationStatusOk() (*string, bool)`

GetVerificationStatusOk returns a tuple with the VerificationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationStatus

`func (o *BankAccountFundingSourceModel) SetVerificationStatus(v string)`

SetVerificationStatus sets VerificationStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


