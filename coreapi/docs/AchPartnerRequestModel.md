# AchPartnerRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessToken** | Pointer to **string** | Required if &#39;user_token&#39; is null | [optional] 
**IdempotentHash** | Pointer to **string** |  | [optional] 
**IsDefaultAccount** | Pointer to **bool** | If there are multiple funding sources, this field specifies which source is used by default in funding calls. If there is only one funding source, the system ignores this field and always uses that source. | [optional] [default to false]
**Partner** | **string** | The name of the partner who validated the account holder. Returned when a third-party partner was used for account validation. | 
**PartnerAccountLinkReferenceToken** | **string** | Supplied by the account validation partner, this value is a reference to the account holder&#39;s details, such as the account number and routing number. Returned when a third-party partner was used for account validation. | 
**Token** | Pointer to **string** | Unique identifier of the funding source. If you do not include a token, the system will generate one automatically. This token is necessary for use in other calls, so we recommend that rather than let the system generate one, you use a simple string that is easy to remember. This value cannot be updated. | [optional] 
**UserToken** | Pointer to **string** | Supplied by the account validation partner, this value is a reference to the account holder&#39;s details, such as the account number and routing number. This token is required if a &#x60;business_token&#x60; is not specified.  Send a &#x60;GET&#x60; request to &#x60;/users&#x60; to retrieve user tokens. | [optional] 

## Methods

### NewAchPartnerRequestModel

`func NewAchPartnerRequestModel(partner string, partnerAccountLinkReferenceToken string, ) *AchPartnerRequestModel`

NewAchPartnerRequestModel instantiates a new AchPartnerRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAchPartnerRequestModelWithDefaults

`func NewAchPartnerRequestModelWithDefaults() *AchPartnerRequestModel`

NewAchPartnerRequestModelWithDefaults instantiates a new AchPartnerRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessToken

`func (o *AchPartnerRequestModel) GetBusinessToken() string`

GetBusinessToken returns the BusinessToken field if non-nil, zero value otherwise.

### GetBusinessTokenOk

`func (o *AchPartnerRequestModel) GetBusinessTokenOk() (*string, bool)`

GetBusinessTokenOk returns a tuple with the BusinessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessToken

`func (o *AchPartnerRequestModel) SetBusinessToken(v string)`

SetBusinessToken sets BusinessToken field to given value.

### HasBusinessToken

`func (o *AchPartnerRequestModel) HasBusinessToken() bool`

HasBusinessToken returns a boolean if a field has been set.

### GetIdempotentHash

`func (o *AchPartnerRequestModel) GetIdempotentHash() string`

GetIdempotentHash returns the IdempotentHash field if non-nil, zero value otherwise.

### GetIdempotentHashOk

`func (o *AchPartnerRequestModel) GetIdempotentHashOk() (*string, bool)`

GetIdempotentHashOk returns a tuple with the IdempotentHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotentHash

`func (o *AchPartnerRequestModel) SetIdempotentHash(v string)`

SetIdempotentHash sets IdempotentHash field to given value.

### HasIdempotentHash

`func (o *AchPartnerRequestModel) HasIdempotentHash() bool`

HasIdempotentHash returns a boolean if a field has been set.

### GetIsDefaultAccount

`func (o *AchPartnerRequestModel) GetIsDefaultAccount() bool`

GetIsDefaultAccount returns the IsDefaultAccount field if non-nil, zero value otherwise.

### GetIsDefaultAccountOk

`func (o *AchPartnerRequestModel) GetIsDefaultAccountOk() (*bool, bool)`

GetIsDefaultAccountOk returns a tuple with the IsDefaultAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefaultAccount

`func (o *AchPartnerRequestModel) SetIsDefaultAccount(v bool)`

SetIsDefaultAccount sets IsDefaultAccount field to given value.

### HasIsDefaultAccount

`func (o *AchPartnerRequestModel) HasIsDefaultAccount() bool`

HasIsDefaultAccount returns a boolean if a field has been set.

### GetPartner

`func (o *AchPartnerRequestModel) GetPartner() string`

GetPartner returns the Partner field if non-nil, zero value otherwise.

### GetPartnerOk

`func (o *AchPartnerRequestModel) GetPartnerOk() (*string, bool)`

GetPartnerOk returns a tuple with the Partner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartner

`func (o *AchPartnerRequestModel) SetPartner(v string)`

SetPartner sets Partner field to given value.


### GetPartnerAccountLinkReferenceToken

`func (o *AchPartnerRequestModel) GetPartnerAccountLinkReferenceToken() string`

GetPartnerAccountLinkReferenceToken returns the PartnerAccountLinkReferenceToken field if non-nil, zero value otherwise.

### GetPartnerAccountLinkReferenceTokenOk

`func (o *AchPartnerRequestModel) GetPartnerAccountLinkReferenceTokenOk() (*string, bool)`

GetPartnerAccountLinkReferenceTokenOk returns a tuple with the PartnerAccountLinkReferenceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerAccountLinkReferenceToken

`func (o *AchPartnerRequestModel) SetPartnerAccountLinkReferenceToken(v string)`

SetPartnerAccountLinkReferenceToken sets PartnerAccountLinkReferenceToken field to given value.


### GetToken

`func (o *AchPartnerRequestModel) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AchPartnerRequestModel) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AchPartnerRequestModel) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *AchPartnerRequestModel) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUserToken

`func (o *AchPartnerRequestModel) GetUserToken() string`

GetUserToken returns the UserToken field if non-nil, zero value otherwise.

### GetUserTokenOk

`func (o *AchPartnerRequestModel) GetUserTokenOk() (*string, bool)`

GetUserTokenOk returns a tuple with the UserToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserToken

`func (o *AchPartnerRequestModel) SetUserToken(v string)`

SetUserToken sets UserToken field to given value.

### HasUserToken

`func (o *AchPartnerRequestModel) HasUserToken() bool`

HasUserToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


