# AccountResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivationTime** | Pointer to **time.Time** | Date and time when the credit account was activated on Marqeta&#39;s credit platform, in UTC. | [optional] 
**AvailableCredit** | **float32** | Amount of credit available for use on the credit account. | 
**BundleToken** | Pointer to **string** | Unique identifier of the associated bundle product. | [optional] 
**Config** | [**AccountConfigResponse**](AccountConfigResponse.md) |  | 
**CreatedTime** | **time.Time** | Date and time when the credit account was created on Marqeta&#39;s credit platform, in UTC. | 
**CreditLimit** | **float32** | Maximum balance the credit account can carry. | 
**CreditProductToken** | **string** | Unique identifier of the associated credit product. | 
**CurrencyCode** | [**CurrencyCode**](CurrencyCode.md) |  | [default to CURRENCYCODE_USD]
**CurrentBalance** | **float32** | Current purchase balance on the credit account. | 
**Description** | Pointer to **string** | Description for the credit account. | [optional] 
**ExternalOfferId** | Pointer to **string** | Unique identifier you provide of the associated external credit offer. | [optional] 
**LatestStatementCycleType** | [**CycleType**](CycleType.md) |  | 
**Name** | Pointer to **string** | Name of the credit account. | [optional] 
**RemainingMinPaymentDue** | **float32** | Amount remaining on the latest statement&#39;s minimum payment, after it&#39;s adjusted for payments, returned payments, and applicable credits that occurred after the latest statement&#39;s closing date. | 
**RemainingStatementBalance** | **float32** | Amount remaining on the latest statement&#39;s balance, after it&#39;s adjusted for payments, returned payments, and applicable credits that occurred after the latest statement&#39;s closing date. | 
**Status** | [**AccountStatusEnum**](AccountStatusEnum.md) |  | 
**Token** | **string** | Unique identifier of the credit account. | 
**Type** | Pointer to [**AccountType**](AccountType.md) |  | [optional] 
**UpdatedTime** | **time.Time** | Date and time when the credit account was last updated on Marqeta&#39;s credit platform, in UTC. | 
**Usages** | [**[]AccountUsageResponse**](AccountUsageResponse.md) | Contains one or more &#x60;usages&#x60; objects that contain information on how a credit account is used and what types of balances are permitted on the account. | 
**UserToken** | **string** | Unique identifier of the primary account holder. | 

## Methods

### NewAccountResponse

`func NewAccountResponse(availableCredit float32, config AccountConfigResponse, createdTime time.Time, creditLimit float32, creditProductToken string, currencyCode CurrencyCode, currentBalance float32, latestStatementCycleType CycleType, remainingMinPaymentDue float32, remainingStatementBalance float32, status AccountStatusEnum, token string, updatedTime time.Time, usages []AccountUsageResponse, userToken string, ) *AccountResponse`

NewAccountResponse instantiates a new AccountResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountResponseWithDefaults

`func NewAccountResponseWithDefaults() *AccountResponse`

NewAccountResponseWithDefaults instantiates a new AccountResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivationTime

`func (o *AccountResponse) GetActivationTime() time.Time`

GetActivationTime returns the ActivationTime field if non-nil, zero value otherwise.

### GetActivationTimeOk

`func (o *AccountResponse) GetActivationTimeOk() (*time.Time, bool)`

GetActivationTimeOk returns a tuple with the ActivationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationTime

`func (o *AccountResponse) SetActivationTime(v time.Time)`

SetActivationTime sets ActivationTime field to given value.

### HasActivationTime

`func (o *AccountResponse) HasActivationTime() bool`

HasActivationTime returns a boolean if a field has been set.

### GetAvailableCredit

`func (o *AccountResponse) GetAvailableCredit() float32`

GetAvailableCredit returns the AvailableCredit field if non-nil, zero value otherwise.

### GetAvailableCreditOk

`func (o *AccountResponse) GetAvailableCreditOk() (*float32, bool)`

GetAvailableCreditOk returns a tuple with the AvailableCredit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableCredit

`func (o *AccountResponse) SetAvailableCredit(v float32)`

SetAvailableCredit sets AvailableCredit field to given value.


### GetBundleToken

`func (o *AccountResponse) GetBundleToken() string`

GetBundleToken returns the BundleToken field if non-nil, zero value otherwise.

### GetBundleTokenOk

`func (o *AccountResponse) GetBundleTokenOk() (*string, bool)`

GetBundleTokenOk returns a tuple with the BundleToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleToken

`func (o *AccountResponse) SetBundleToken(v string)`

SetBundleToken sets BundleToken field to given value.

### HasBundleToken

`func (o *AccountResponse) HasBundleToken() bool`

HasBundleToken returns a boolean if a field has been set.

### GetConfig

`func (o *AccountResponse) GetConfig() AccountConfigResponse`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AccountResponse) GetConfigOk() (*AccountConfigResponse, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AccountResponse) SetConfig(v AccountConfigResponse)`

SetConfig sets Config field to given value.


### GetCreatedTime

`func (o *AccountResponse) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *AccountResponse) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *AccountResponse) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetCreditLimit

`func (o *AccountResponse) GetCreditLimit() float32`

GetCreditLimit returns the CreditLimit field if non-nil, zero value otherwise.

### GetCreditLimitOk

`func (o *AccountResponse) GetCreditLimitOk() (*float32, bool)`

GetCreditLimitOk returns a tuple with the CreditLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditLimit

`func (o *AccountResponse) SetCreditLimit(v float32)`

SetCreditLimit sets CreditLimit field to given value.


### GetCreditProductToken

`func (o *AccountResponse) GetCreditProductToken() string`

GetCreditProductToken returns the CreditProductToken field if non-nil, zero value otherwise.

### GetCreditProductTokenOk

`func (o *AccountResponse) GetCreditProductTokenOk() (*string, bool)`

GetCreditProductTokenOk returns a tuple with the CreditProductToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditProductToken

`func (o *AccountResponse) SetCreditProductToken(v string)`

SetCreditProductToken sets CreditProductToken field to given value.


### GetCurrencyCode

`func (o *AccountResponse) GetCurrencyCode() CurrencyCode`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *AccountResponse) GetCurrencyCodeOk() (*CurrencyCode, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *AccountResponse) SetCurrencyCode(v CurrencyCode)`

SetCurrencyCode sets CurrencyCode field to given value.


### GetCurrentBalance

`func (o *AccountResponse) GetCurrentBalance() float32`

GetCurrentBalance returns the CurrentBalance field if non-nil, zero value otherwise.

### GetCurrentBalanceOk

`func (o *AccountResponse) GetCurrentBalanceOk() (*float32, bool)`

GetCurrentBalanceOk returns a tuple with the CurrentBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentBalance

`func (o *AccountResponse) SetCurrentBalance(v float32)`

SetCurrentBalance sets CurrentBalance field to given value.


### GetDescription

`func (o *AccountResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AccountResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AccountResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AccountResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExternalOfferId

`func (o *AccountResponse) GetExternalOfferId() string`

GetExternalOfferId returns the ExternalOfferId field if non-nil, zero value otherwise.

### GetExternalOfferIdOk

`func (o *AccountResponse) GetExternalOfferIdOk() (*string, bool)`

GetExternalOfferIdOk returns a tuple with the ExternalOfferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalOfferId

`func (o *AccountResponse) SetExternalOfferId(v string)`

SetExternalOfferId sets ExternalOfferId field to given value.

### HasExternalOfferId

`func (o *AccountResponse) HasExternalOfferId() bool`

HasExternalOfferId returns a boolean if a field has been set.

### GetLatestStatementCycleType

`func (o *AccountResponse) GetLatestStatementCycleType() CycleType`

GetLatestStatementCycleType returns the LatestStatementCycleType field if non-nil, zero value otherwise.

### GetLatestStatementCycleTypeOk

`func (o *AccountResponse) GetLatestStatementCycleTypeOk() (*CycleType, bool)`

GetLatestStatementCycleTypeOk returns a tuple with the LatestStatementCycleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestStatementCycleType

`func (o *AccountResponse) SetLatestStatementCycleType(v CycleType)`

SetLatestStatementCycleType sets LatestStatementCycleType field to given value.


### GetName

`func (o *AccountResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AccountResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRemainingMinPaymentDue

`func (o *AccountResponse) GetRemainingMinPaymentDue() float32`

GetRemainingMinPaymentDue returns the RemainingMinPaymentDue field if non-nil, zero value otherwise.

### GetRemainingMinPaymentDueOk

`func (o *AccountResponse) GetRemainingMinPaymentDueOk() (*float32, bool)`

GetRemainingMinPaymentDueOk returns a tuple with the RemainingMinPaymentDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingMinPaymentDue

`func (o *AccountResponse) SetRemainingMinPaymentDue(v float32)`

SetRemainingMinPaymentDue sets RemainingMinPaymentDue field to given value.


### GetRemainingStatementBalance

`func (o *AccountResponse) GetRemainingStatementBalance() float32`

GetRemainingStatementBalance returns the RemainingStatementBalance field if non-nil, zero value otherwise.

### GetRemainingStatementBalanceOk

`func (o *AccountResponse) GetRemainingStatementBalanceOk() (*float32, bool)`

GetRemainingStatementBalanceOk returns a tuple with the RemainingStatementBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingStatementBalance

`func (o *AccountResponse) SetRemainingStatementBalance(v float32)`

SetRemainingStatementBalance sets RemainingStatementBalance field to given value.


### GetStatus

`func (o *AccountResponse) GetStatus() AccountStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AccountResponse) GetStatusOk() (*AccountStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AccountResponse) SetStatus(v AccountStatusEnum)`

SetStatus sets Status field to given value.


### GetToken

`func (o *AccountResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AccountResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AccountResponse) SetToken(v string)`

SetToken sets Token field to given value.


### GetType

`func (o *AccountResponse) GetType() AccountType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AccountResponse) GetTypeOk() (*AccountType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AccountResponse) SetType(v AccountType)`

SetType sets Type field to given value.

### HasType

`func (o *AccountResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedTime

`func (o *AccountResponse) GetUpdatedTime() time.Time`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *AccountResponse) GetUpdatedTimeOk() (*time.Time, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *AccountResponse) SetUpdatedTime(v time.Time)`

SetUpdatedTime sets UpdatedTime field to given value.


### GetUsages

`func (o *AccountResponse) GetUsages() []AccountUsageResponse`

GetUsages returns the Usages field if non-nil, zero value otherwise.

### GetUsagesOk

`func (o *AccountResponse) GetUsagesOk() (*[]AccountUsageResponse, bool)`

GetUsagesOk returns a tuple with the Usages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsages

`func (o *AccountResponse) SetUsages(v []AccountUsageResponse)`

SetUsages sets Usages field to given value.


### GetUserToken

`func (o *AccountResponse) GetUserToken() string`

GetUserToken returns the UserToken field if non-nil, zero value otherwise.

### GetUserTokenOk

`func (o *AccountResponse) GetUserTokenOk() (*string, bool)`

GetUserTokenOk returns a tuple with the UserToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserToken

`func (o *AccountResponse) SetUserToken(v string)`

SetUserToken sets UserToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


