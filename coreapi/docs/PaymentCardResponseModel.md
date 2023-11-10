# PaymentCardResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSuffix** | **string** | Account identifier appended to the payment card number. | 
**AccountType** | **string** | Type of payment card account. | 
**Active** | **bool** | Specifies whether the account is active. | [default to false]
**BusinessToken** | Pointer to **string** | Unique identifier of the business account holder. This token is returned if a &#x60;user_token&#x60; is not specified. | [optional] 
**CreatedTime** | **time.Time** | Date and time when the resource was created, in UTC. | 
**ExpDate** | **string** | Payment card expiration date. | 
**IsDefaultAccount** | **bool** | If there are multiple funding sources, this field specifies which source is used by default in funding calls. If there is only one funding source, the system ignores this field and always uses that source. | [default to false]
**LastModifiedTime** | **time.Time** | Date and time when the resource was last modified, in UTC. | 
**Token** | **string** | Unique identifier of the funding source. | 
**Type** | **string** | Funding source type. | 
**UserToken** | Pointer to **string** | Unique identifier of the user account holder. This token is returned if a &#x60;business_token&#x60; is not specified. | [optional] 

## Methods

### NewPaymentCardResponseModel

`func NewPaymentCardResponseModel(accountSuffix string, accountType string, active bool, createdTime time.Time, expDate string, isDefaultAccount bool, lastModifiedTime time.Time, token string, type_ string, ) *PaymentCardResponseModel`

NewPaymentCardResponseModel instantiates a new PaymentCardResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentCardResponseModelWithDefaults

`func NewPaymentCardResponseModelWithDefaults() *PaymentCardResponseModel`

NewPaymentCardResponseModelWithDefaults instantiates a new PaymentCardResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSuffix

`func (o *PaymentCardResponseModel) GetAccountSuffix() string`

GetAccountSuffix returns the AccountSuffix field if non-nil, zero value otherwise.

### GetAccountSuffixOk

`func (o *PaymentCardResponseModel) GetAccountSuffixOk() (*string, bool)`

GetAccountSuffixOk returns a tuple with the AccountSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSuffix

`func (o *PaymentCardResponseModel) SetAccountSuffix(v string)`

SetAccountSuffix sets AccountSuffix field to given value.


### GetAccountType

`func (o *PaymentCardResponseModel) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *PaymentCardResponseModel) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *PaymentCardResponseModel) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.


### GetActive

`func (o *PaymentCardResponseModel) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *PaymentCardResponseModel) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *PaymentCardResponseModel) SetActive(v bool)`

SetActive sets Active field to given value.


### GetBusinessToken

`func (o *PaymentCardResponseModel) GetBusinessToken() string`

GetBusinessToken returns the BusinessToken field if non-nil, zero value otherwise.

### GetBusinessTokenOk

`func (o *PaymentCardResponseModel) GetBusinessTokenOk() (*string, bool)`

GetBusinessTokenOk returns a tuple with the BusinessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessToken

`func (o *PaymentCardResponseModel) SetBusinessToken(v string)`

SetBusinessToken sets BusinessToken field to given value.

### HasBusinessToken

`func (o *PaymentCardResponseModel) HasBusinessToken() bool`

HasBusinessToken returns a boolean if a field has been set.

### GetCreatedTime

`func (o *PaymentCardResponseModel) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *PaymentCardResponseModel) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *PaymentCardResponseModel) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetExpDate

`func (o *PaymentCardResponseModel) GetExpDate() string`

GetExpDate returns the ExpDate field if non-nil, zero value otherwise.

### GetExpDateOk

`func (o *PaymentCardResponseModel) GetExpDateOk() (*string, bool)`

GetExpDateOk returns a tuple with the ExpDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpDate

`func (o *PaymentCardResponseModel) SetExpDate(v string)`

SetExpDate sets ExpDate field to given value.


### GetIsDefaultAccount

`func (o *PaymentCardResponseModel) GetIsDefaultAccount() bool`

GetIsDefaultAccount returns the IsDefaultAccount field if non-nil, zero value otherwise.

### GetIsDefaultAccountOk

`func (o *PaymentCardResponseModel) GetIsDefaultAccountOk() (*bool, bool)`

GetIsDefaultAccountOk returns a tuple with the IsDefaultAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefaultAccount

`func (o *PaymentCardResponseModel) SetIsDefaultAccount(v bool)`

SetIsDefaultAccount sets IsDefaultAccount field to given value.


### GetLastModifiedTime

`func (o *PaymentCardResponseModel) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *PaymentCardResponseModel) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *PaymentCardResponseModel) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.


### GetToken

`func (o *PaymentCardResponseModel) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PaymentCardResponseModel) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PaymentCardResponseModel) SetToken(v string)`

SetToken sets Token field to given value.


### GetType

`func (o *PaymentCardResponseModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaymentCardResponseModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaymentCardResponseModel) SetType(v string)`

SetType sets Type field to given value.


### GetUserToken

`func (o *PaymentCardResponseModel) GetUserToken() string`

GetUserToken returns the UserToken field if non-nil, zero value otherwise.

### GetUserTokenOk

`func (o *PaymentCardResponseModel) GetUserTokenOk() (*string, bool)`

GetUserTokenOk returns a tuple with the UserToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserToken

`func (o *PaymentCardResponseModel) SetUserToken(v string)`

SetUserToken sets UserToken field to given value.

### HasUserToken

`func (o *PaymentCardResponseModel) HasUserToken() bool`

HasUserToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


