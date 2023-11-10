# DirectDepositAccountRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowImmediateCredit** | Pointer to **bool** |  | [optional] [default to false]
**BusinessToken** | Pointer to **string** | Required if &#39;user_token&#39; is null | [optional] 
**CustomerDueDiligence** | Pointer to [**[]CustomerDueDiligenceRequest**](CustomerDueDiligenceRequest.md) | Required if account type &#x3D; Checking | [optional] 
**Token** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UserToken** | Pointer to **string** | Required if &#39;business_token&#39; is null | [optional] 

## Methods

### NewDirectDepositAccountRequest

`func NewDirectDepositAccountRequest() *DirectDepositAccountRequest`

NewDirectDepositAccountRequest instantiates a new DirectDepositAccountRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectDepositAccountRequestWithDefaults

`func NewDirectDepositAccountRequestWithDefaults() *DirectDepositAccountRequest`

NewDirectDepositAccountRequestWithDefaults instantiates a new DirectDepositAccountRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowImmediateCredit

`func (o *DirectDepositAccountRequest) GetAllowImmediateCredit() bool`

GetAllowImmediateCredit returns the AllowImmediateCredit field if non-nil, zero value otherwise.

### GetAllowImmediateCreditOk

`func (o *DirectDepositAccountRequest) GetAllowImmediateCreditOk() (*bool, bool)`

GetAllowImmediateCreditOk returns a tuple with the AllowImmediateCredit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowImmediateCredit

`func (o *DirectDepositAccountRequest) SetAllowImmediateCredit(v bool)`

SetAllowImmediateCredit sets AllowImmediateCredit field to given value.

### HasAllowImmediateCredit

`func (o *DirectDepositAccountRequest) HasAllowImmediateCredit() bool`

HasAllowImmediateCredit returns a boolean if a field has been set.

### GetBusinessToken

`func (o *DirectDepositAccountRequest) GetBusinessToken() string`

GetBusinessToken returns the BusinessToken field if non-nil, zero value otherwise.

### GetBusinessTokenOk

`func (o *DirectDepositAccountRequest) GetBusinessTokenOk() (*string, bool)`

GetBusinessTokenOk returns a tuple with the BusinessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessToken

`func (o *DirectDepositAccountRequest) SetBusinessToken(v string)`

SetBusinessToken sets BusinessToken field to given value.

### HasBusinessToken

`func (o *DirectDepositAccountRequest) HasBusinessToken() bool`

HasBusinessToken returns a boolean if a field has been set.

### GetCustomerDueDiligence

`func (o *DirectDepositAccountRequest) GetCustomerDueDiligence() []CustomerDueDiligenceRequest`

GetCustomerDueDiligence returns the CustomerDueDiligence field if non-nil, zero value otherwise.

### GetCustomerDueDiligenceOk

`func (o *DirectDepositAccountRequest) GetCustomerDueDiligenceOk() (*[]CustomerDueDiligenceRequest, bool)`

GetCustomerDueDiligenceOk returns a tuple with the CustomerDueDiligence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerDueDiligence

`func (o *DirectDepositAccountRequest) SetCustomerDueDiligence(v []CustomerDueDiligenceRequest)`

SetCustomerDueDiligence sets CustomerDueDiligence field to given value.

### HasCustomerDueDiligence

`func (o *DirectDepositAccountRequest) HasCustomerDueDiligence() bool`

HasCustomerDueDiligence returns a boolean if a field has been set.

### GetToken

`func (o *DirectDepositAccountRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DirectDepositAccountRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DirectDepositAccountRequest) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DirectDepositAccountRequest) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetType

`func (o *DirectDepositAccountRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DirectDepositAccountRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DirectDepositAccountRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DirectDepositAccountRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUserToken

`func (o *DirectDepositAccountRequest) GetUserToken() string`

GetUserToken returns the UserToken field if non-nil, zero value otherwise.

### GetUserTokenOk

`func (o *DirectDepositAccountRequest) GetUserTokenOk() (*string, bool)`

GetUserTokenOk returns a tuple with the UserToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserToken

`func (o *DirectDepositAccountRequest) SetUserToken(v string)`

SetUserToken sets UserToken field to given value.

### HasUserToken

`func (o *DirectDepositAccountRequest) HasUserToken() bool`

HasUserToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


