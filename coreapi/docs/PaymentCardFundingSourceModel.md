# PaymentCardFundingSourceModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSuffix** | **string** |  | 
**AccountType** | **string** |  | 
**BusinessToken** | Pointer to **string** | Required if &#39;user_token&#39; is null | [optional] 
**ExpDate** | **string** |  | 
**UserToken** | Pointer to **string** | Required if &#39;business_token&#39; is null | [optional] 

## Methods

### NewPaymentCardFundingSourceModel

`func NewPaymentCardFundingSourceModel(accountSuffix string, accountType string, expDate string, ) *PaymentCardFundingSourceModel`

NewPaymentCardFundingSourceModel instantiates a new PaymentCardFundingSourceModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentCardFundingSourceModelWithDefaults

`func NewPaymentCardFundingSourceModelWithDefaults() *PaymentCardFundingSourceModel`

NewPaymentCardFundingSourceModelWithDefaults instantiates a new PaymentCardFundingSourceModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSuffix

`func (o *PaymentCardFundingSourceModel) GetAccountSuffix() string`

GetAccountSuffix returns the AccountSuffix field if non-nil, zero value otherwise.

### GetAccountSuffixOk

`func (o *PaymentCardFundingSourceModel) GetAccountSuffixOk() (*string, bool)`

GetAccountSuffixOk returns a tuple with the AccountSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSuffix

`func (o *PaymentCardFundingSourceModel) SetAccountSuffix(v string)`

SetAccountSuffix sets AccountSuffix field to given value.


### GetAccountType

`func (o *PaymentCardFundingSourceModel) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *PaymentCardFundingSourceModel) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *PaymentCardFundingSourceModel) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.


### GetBusinessToken

`func (o *PaymentCardFundingSourceModel) GetBusinessToken() string`

GetBusinessToken returns the BusinessToken field if non-nil, zero value otherwise.

### GetBusinessTokenOk

`func (o *PaymentCardFundingSourceModel) GetBusinessTokenOk() (*string, bool)`

GetBusinessTokenOk returns a tuple with the BusinessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessToken

`func (o *PaymentCardFundingSourceModel) SetBusinessToken(v string)`

SetBusinessToken sets BusinessToken field to given value.

### HasBusinessToken

`func (o *PaymentCardFundingSourceModel) HasBusinessToken() bool`

HasBusinessToken returns a boolean if a field has been set.

### GetExpDate

`func (o *PaymentCardFundingSourceModel) GetExpDate() string`

GetExpDate returns the ExpDate field if non-nil, zero value otherwise.

### GetExpDateOk

`func (o *PaymentCardFundingSourceModel) GetExpDateOk() (*string, bool)`

GetExpDateOk returns a tuple with the ExpDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpDate

`func (o *PaymentCardFundingSourceModel) SetExpDate(v string)`

SetExpDate sets ExpDate field to given value.


### GetUserToken

`func (o *PaymentCardFundingSourceModel) GetUserToken() string`

GetUserToken returns the UserToken field if non-nil, zero value otherwise.

### GetUserTokenOk

`func (o *PaymentCardFundingSourceModel) GetUserTokenOk() (*string, bool)`

GetUserTokenOk returns a tuple with the UserToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserToken

`func (o *PaymentCardFundingSourceModel) SetUserToken(v string)`

SetUserToken sets UserToken field to given value.

### HasUserToken

`func (o *PaymentCardFundingSourceModel) HasUserToken() bool`

HasUserToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


