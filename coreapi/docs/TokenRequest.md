# TokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountNumber** | **string** | Payment card account number. | 
**BusinessToken** | Pointer to **string** | Unique identifier of the business account holder. This token is required if the &#x60;user_token&#x60; is not included.  Send a &#x60;GET&#x60; request to &#x60;/businesses&#x60; to retrieve business tokens. | [optional] 
**CvvNumber** | **string** | Payment card CVV2 number. | 
**ExpDate** | **string** | Payment card expiration date. | 
**IsDefaultAccount** | Pointer to **bool** | If there are multiple funding sources, this field specifies which source is used by default in funding calls. If there is only one funding source, the system ignores this field and always uses that source. | [optional] [default to false]
**PostalCode** | Pointer to **string** | Postal code of the account holder (user or business). | [optional] 
**Token** | Pointer to **string** | Unique identifier of the funding account. If you do not include a token, the system will generate one automatically. As this token is necessary for use in other calls, we recommend that you define a simple and easy to remember string rather than letting the system generate a token for you. This value cannot be updated. | [optional] 
**UserToken** | Pointer to **string** | Unique identifier of the user account holder. This token is required if the &#x60;business_token&#x60; is not included.  Send a &#x60;GET&#x60; request to &#x60;/users&#x60; to retrieve user tokens. | [optional] 
**Zip** | Pointer to **string** | United States ZIP code of the account holder (user or business). | [optional] 

## Methods

### NewTokenRequest

`func NewTokenRequest(accountNumber string, cvvNumber string, expDate string, ) *TokenRequest`

NewTokenRequest instantiates a new TokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenRequestWithDefaults

`func NewTokenRequestWithDefaults() *TokenRequest`

NewTokenRequestWithDefaults instantiates a new TokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNumber

`func (o *TokenRequest) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *TokenRequest) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *TokenRequest) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.


### GetBusinessToken

`func (o *TokenRequest) GetBusinessToken() string`

GetBusinessToken returns the BusinessToken field if non-nil, zero value otherwise.

### GetBusinessTokenOk

`func (o *TokenRequest) GetBusinessTokenOk() (*string, bool)`

GetBusinessTokenOk returns a tuple with the BusinessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessToken

`func (o *TokenRequest) SetBusinessToken(v string)`

SetBusinessToken sets BusinessToken field to given value.

### HasBusinessToken

`func (o *TokenRequest) HasBusinessToken() bool`

HasBusinessToken returns a boolean if a field has been set.

### GetCvvNumber

`func (o *TokenRequest) GetCvvNumber() string`

GetCvvNumber returns the CvvNumber field if non-nil, zero value otherwise.

### GetCvvNumberOk

`func (o *TokenRequest) GetCvvNumberOk() (*string, bool)`

GetCvvNumberOk returns a tuple with the CvvNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvvNumber

`func (o *TokenRequest) SetCvvNumber(v string)`

SetCvvNumber sets CvvNumber field to given value.


### GetExpDate

`func (o *TokenRequest) GetExpDate() string`

GetExpDate returns the ExpDate field if non-nil, zero value otherwise.

### GetExpDateOk

`func (o *TokenRequest) GetExpDateOk() (*string, bool)`

GetExpDateOk returns a tuple with the ExpDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpDate

`func (o *TokenRequest) SetExpDate(v string)`

SetExpDate sets ExpDate field to given value.


### GetIsDefaultAccount

`func (o *TokenRequest) GetIsDefaultAccount() bool`

GetIsDefaultAccount returns the IsDefaultAccount field if non-nil, zero value otherwise.

### GetIsDefaultAccountOk

`func (o *TokenRequest) GetIsDefaultAccountOk() (*bool, bool)`

GetIsDefaultAccountOk returns a tuple with the IsDefaultAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefaultAccount

`func (o *TokenRequest) SetIsDefaultAccount(v bool)`

SetIsDefaultAccount sets IsDefaultAccount field to given value.

### HasIsDefaultAccount

`func (o *TokenRequest) HasIsDefaultAccount() bool`

HasIsDefaultAccount returns a boolean if a field has been set.

### GetPostalCode

`func (o *TokenRequest) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *TokenRequest) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *TokenRequest) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *TokenRequest) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetToken

`func (o *TokenRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TokenRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TokenRequest) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *TokenRequest) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUserToken

`func (o *TokenRequest) GetUserToken() string`

GetUserToken returns the UserToken field if non-nil, zero value otherwise.

### GetUserTokenOk

`func (o *TokenRequest) GetUserTokenOk() (*string, bool)`

GetUserTokenOk returns a tuple with the UserToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserToken

`func (o *TokenRequest) SetUserToken(v string)`

SetUserToken sets UserToken field to given value.

### HasUserToken

`func (o *TokenRequest) HasUserToken() bool`

HasUserToken returns a boolean if a field has been set.

### GetZip

`func (o *TokenRequest) GetZip() string`

GetZip returns the Zip field if non-nil, zero value otherwise.

### GetZipOk

`func (o *TokenRequest) GetZipOk() (*string, bool)`

GetZipOk returns a tuple with the Zip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZip

`func (o *TokenRequest) SetZip(v string)`

SetZip sets Zip field to given value.

### HasZip

`func (o *TokenRequest) HasZip() bool`

HasZip returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


