# AccountCreateReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationToken** | Pointer to **string** | Unique identifier of the associated credit account application. | [optional] 
**BundleToken** | Pointer to **string** | Unique identifier of the associated bundle.  You must pass either &#x60;bundle_token&#x60; or both &#x60;credit_product_token&#x60; and &#x60;external_offer_id&#x60;. | [optional] 
**Config** | Pointer to [**AccountConfigReq**](AccountConfigReq.md) |  | [optional] 
**CreditLimit** | **float32** | Maximum balance the credit account can carry. | 
**CreditProductToken** | Pointer to **string** | Unique identifier of the associated credit product.  This field is required if passing &#x60;external_offer_id&#x60;.  You must pass either both &#x60;credit_product_token&#x60; and &#x60;external_offer_id&#x60; or &#x60;bundle_token&#x60;. | [optional] 
**Description** | Pointer to **string** | Description for the credit account. | [optional] 
**ExternalOfferId** | Pointer to **string** | Unique identifier you provide of the associated external credit offer.  This field is required if passing &#x60;credit_product_token&#x60;.  You must pass either both &#x60;external_offer_id&#x60; and &#x60;credit_product_token&#x60; or &#x60;bundle_token&#x60;. | [optional] 
**Name** | Pointer to **string** | Name of the credit account. | [optional] 
**Token** | Pointer to **string** | Unique identifier of the credit account. | [optional] 
**Usages** | [**[]AccountUsageCreateReq**](AccountUsageCreateReq.md) | Contains one or more &#x60;usages&#x60; objects that contain information on how a credit account is used and what types of balances are permitted on the account.  You can pass only one &#x60;usages&#x60; object per &#x60;usages.type&#x60;. | 
**UserToken** | **string** | Unique identifier of the primary account holder. | 

## Methods

### NewAccountCreateReq

`func NewAccountCreateReq(creditLimit float32, usages []AccountUsageCreateReq, userToken string, ) *AccountCreateReq`

NewAccountCreateReq instantiates a new AccountCreateReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountCreateReqWithDefaults

`func NewAccountCreateReqWithDefaults() *AccountCreateReq`

NewAccountCreateReqWithDefaults instantiates a new AccountCreateReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationToken

`func (o *AccountCreateReq) GetApplicationToken() string`

GetApplicationToken returns the ApplicationToken field if non-nil, zero value otherwise.

### GetApplicationTokenOk

`func (o *AccountCreateReq) GetApplicationTokenOk() (*string, bool)`

GetApplicationTokenOk returns a tuple with the ApplicationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationToken

`func (o *AccountCreateReq) SetApplicationToken(v string)`

SetApplicationToken sets ApplicationToken field to given value.

### HasApplicationToken

`func (o *AccountCreateReq) HasApplicationToken() bool`

HasApplicationToken returns a boolean if a field has been set.

### GetBundleToken

`func (o *AccountCreateReq) GetBundleToken() string`

GetBundleToken returns the BundleToken field if non-nil, zero value otherwise.

### GetBundleTokenOk

`func (o *AccountCreateReq) GetBundleTokenOk() (*string, bool)`

GetBundleTokenOk returns a tuple with the BundleToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleToken

`func (o *AccountCreateReq) SetBundleToken(v string)`

SetBundleToken sets BundleToken field to given value.

### HasBundleToken

`func (o *AccountCreateReq) HasBundleToken() bool`

HasBundleToken returns a boolean if a field has been set.

### GetConfig

`func (o *AccountCreateReq) GetConfig() AccountConfigReq`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AccountCreateReq) GetConfigOk() (*AccountConfigReq, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AccountCreateReq) SetConfig(v AccountConfigReq)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *AccountCreateReq) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCreditLimit

`func (o *AccountCreateReq) GetCreditLimit() float32`

GetCreditLimit returns the CreditLimit field if non-nil, zero value otherwise.

### GetCreditLimitOk

`func (o *AccountCreateReq) GetCreditLimitOk() (*float32, bool)`

GetCreditLimitOk returns a tuple with the CreditLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditLimit

`func (o *AccountCreateReq) SetCreditLimit(v float32)`

SetCreditLimit sets CreditLimit field to given value.


### GetCreditProductToken

`func (o *AccountCreateReq) GetCreditProductToken() string`

GetCreditProductToken returns the CreditProductToken field if non-nil, zero value otherwise.

### GetCreditProductTokenOk

`func (o *AccountCreateReq) GetCreditProductTokenOk() (*string, bool)`

GetCreditProductTokenOk returns a tuple with the CreditProductToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditProductToken

`func (o *AccountCreateReq) SetCreditProductToken(v string)`

SetCreditProductToken sets CreditProductToken field to given value.

### HasCreditProductToken

`func (o *AccountCreateReq) HasCreditProductToken() bool`

HasCreditProductToken returns a boolean if a field has been set.

### GetDescription

`func (o *AccountCreateReq) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AccountCreateReq) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AccountCreateReq) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AccountCreateReq) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExternalOfferId

`func (o *AccountCreateReq) GetExternalOfferId() string`

GetExternalOfferId returns the ExternalOfferId field if non-nil, zero value otherwise.

### GetExternalOfferIdOk

`func (o *AccountCreateReq) GetExternalOfferIdOk() (*string, bool)`

GetExternalOfferIdOk returns a tuple with the ExternalOfferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalOfferId

`func (o *AccountCreateReq) SetExternalOfferId(v string)`

SetExternalOfferId sets ExternalOfferId field to given value.

### HasExternalOfferId

`func (o *AccountCreateReq) HasExternalOfferId() bool`

HasExternalOfferId returns a boolean if a field has been set.

### GetName

`func (o *AccountCreateReq) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountCreateReq) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountCreateReq) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AccountCreateReq) HasName() bool`

HasName returns a boolean if a field has been set.

### GetToken

`func (o *AccountCreateReq) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AccountCreateReq) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AccountCreateReq) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *AccountCreateReq) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUsages

`func (o *AccountCreateReq) GetUsages() []AccountUsageCreateReq`

GetUsages returns the Usages field if non-nil, zero value otherwise.

### GetUsagesOk

`func (o *AccountCreateReq) GetUsagesOk() (*[]AccountUsageCreateReq, bool)`

GetUsagesOk returns a tuple with the Usages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsages

`func (o *AccountCreateReq) SetUsages(v []AccountUsageCreateReq)`

SetUsages sets Usages field to given value.


### GetUserToken

`func (o *AccountCreateReq) GetUserToken() string`

GetUserToken returns the UserToken field if non-nil, zero value otherwise.

### GetUserTokenOk

`func (o *AccountCreateReq) GetUserTokenOk() (*string, bool)`

GetUserTokenOk returns a tuple with the UserToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserToken

`func (o *AccountCreateReq) SetUserToken(v string)`

SetUserToken sets UserToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


