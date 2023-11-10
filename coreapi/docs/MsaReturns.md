# MsaReturns

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | **bool** |  | [default to false]
**AggregatedBalances** | [**MsaAggregatedBalances**](MsaAggregatedBalances.md) |  | 
**Amount** | **float32** |  | 
**BusinessToken** | Pointer to **string** |  | [optional] 
**CampaignToken** | **string** |  | 
**CreatedTime** | **time.Time** | yyyy-MM-ddTHH:mm:ssZ | 
**CurrencyCode** | **string** |  | 
**EndDate** | Pointer to **time.Time** | yyyy-MM-ddThh:mm:ssZ | [optional] 
**Funding** | [**Funding**](Funding.md) |  | 
**LastModifiedTime** | **time.Time** | yyyy-MM-ddTHH:mm:ssZ | 
**LastTransactionDate** | **time.Time** | yyyy-MM-ddThh:mm:ssZ | 
**OrderBalances** | [**MsaBalances**](MsaBalances.md) |  | 
**OriginalOrderToken** | **string** |  | 
**RewardAmount** | **float32** |  | 
**RewardTriggerAmount** | **float32** |  | 
**StartDate** | Pointer to **time.Time** | yyyy-MM-ddThh:mm:ssZ | [optional] 
**Token** | Pointer to **string** |  | [optional] 
**TransactionToken** | **string** |  | 
**UnloadedAmount** | Pointer to **float32** |  | [optional] 
**UserToken** | Pointer to **string** |  | [optional] 

## Methods

### NewMsaReturns

`func NewMsaReturns(active bool, aggregatedBalances MsaAggregatedBalances, amount float32, campaignToken string, createdTime time.Time, currencyCode string, funding Funding, lastModifiedTime time.Time, lastTransactionDate time.Time, orderBalances MsaBalances, originalOrderToken string, rewardAmount float32, rewardTriggerAmount float32, transactionToken string, ) *MsaReturns`

NewMsaReturns instantiates a new MsaReturns object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsaReturnsWithDefaults

`func NewMsaReturnsWithDefaults() *MsaReturns`

NewMsaReturnsWithDefaults instantiates a new MsaReturns object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *MsaReturns) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *MsaReturns) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *MsaReturns) SetActive(v bool)`

SetActive sets Active field to given value.


### GetAggregatedBalances

`func (o *MsaReturns) GetAggregatedBalances() MsaAggregatedBalances`

GetAggregatedBalances returns the AggregatedBalances field if non-nil, zero value otherwise.

### GetAggregatedBalancesOk

`func (o *MsaReturns) GetAggregatedBalancesOk() (*MsaAggregatedBalances, bool)`

GetAggregatedBalancesOk returns a tuple with the AggregatedBalances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregatedBalances

`func (o *MsaReturns) SetAggregatedBalances(v MsaAggregatedBalances)`

SetAggregatedBalances sets AggregatedBalances field to given value.


### GetAmount

`func (o *MsaReturns) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *MsaReturns) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *MsaReturns) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetBusinessToken

`func (o *MsaReturns) GetBusinessToken() string`

GetBusinessToken returns the BusinessToken field if non-nil, zero value otherwise.

### GetBusinessTokenOk

`func (o *MsaReturns) GetBusinessTokenOk() (*string, bool)`

GetBusinessTokenOk returns a tuple with the BusinessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessToken

`func (o *MsaReturns) SetBusinessToken(v string)`

SetBusinessToken sets BusinessToken field to given value.

### HasBusinessToken

`func (o *MsaReturns) HasBusinessToken() bool`

HasBusinessToken returns a boolean if a field has been set.

### GetCampaignToken

`func (o *MsaReturns) GetCampaignToken() string`

GetCampaignToken returns the CampaignToken field if non-nil, zero value otherwise.

### GetCampaignTokenOk

`func (o *MsaReturns) GetCampaignTokenOk() (*string, bool)`

GetCampaignTokenOk returns a tuple with the CampaignToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignToken

`func (o *MsaReturns) SetCampaignToken(v string)`

SetCampaignToken sets CampaignToken field to given value.


### GetCreatedTime

`func (o *MsaReturns) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *MsaReturns) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *MsaReturns) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetCurrencyCode

`func (o *MsaReturns) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *MsaReturns) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *MsaReturns) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.


### GetEndDate

`func (o *MsaReturns) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *MsaReturns) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *MsaReturns) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *MsaReturns) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetFunding

`func (o *MsaReturns) GetFunding() Funding`

GetFunding returns the Funding field if non-nil, zero value otherwise.

### GetFundingOk

`func (o *MsaReturns) GetFundingOk() (*Funding, bool)`

GetFundingOk returns a tuple with the Funding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunding

`func (o *MsaReturns) SetFunding(v Funding)`

SetFunding sets Funding field to given value.


### GetLastModifiedTime

`func (o *MsaReturns) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *MsaReturns) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *MsaReturns) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.


### GetLastTransactionDate

`func (o *MsaReturns) GetLastTransactionDate() time.Time`

GetLastTransactionDate returns the LastTransactionDate field if non-nil, zero value otherwise.

### GetLastTransactionDateOk

`func (o *MsaReturns) GetLastTransactionDateOk() (*time.Time, bool)`

GetLastTransactionDateOk returns a tuple with the LastTransactionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTransactionDate

`func (o *MsaReturns) SetLastTransactionDate(v time.Time)`

SetLastTransactionDate sets LastTransactionDate field to given value.


### GetOrderBalances

`func (o *MsaReturns) GetOrderBalances() MsaBalances`

GetOrderBalances returns the OrderBalances field if non-nil, zero value otherwise.

### GetOrderBalancesOk

`func (o *MsaReturns) GetOrderBalancesOk() (*MsaBalances, bool)`

GetOrderBalancesOk returns a tuple with the OrderBalances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBalances

`func (o *MsaReturns) SetOrderBalances(v MsaBalances)`

SetOrderBalances sets OrderBalances field to given value.


### GetOriginalOrderToken

`func (o *MsaReturns) GetOriginalOrderToken() string`

GetOriginalOrderToken returns the OriginalOrderToken field if non-nil, zero value otherwise.

### GetOriginalOrderTokenOk

`func (o *MsaReturns) GetOriginalOrderTokenOk() (*string, bool)`

GetOriginalOrderTokenOk returns a tuple with the OriginalOrderToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalOrderToken

`func (o *MsaReturns) SetOriginalOrderToken(v string)`

SetOriginalOrderToken sets OriginalOrderToken field to given value.


### GetRewardAmount

`func (o *MsaReturns) GetRewardAmount() float32`

GetRewardAmount returns the RewardAmount field if non-nil, zero value otherwise.

### GetRewardAmountOk

`func (o *MsaReturns) GetRewardAmountOk() (*float32, bool)`

GetRewardAmountOk returns a tuple with the RewardAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardAmount

`func (o *MsaReturns) SetRewardAmount(v float32)`

SetRewardAmount sets RewardAmount field to given value.


### GetRewardTriggerAmount

`func (o *MsaReturns) GetRewardTriggerAmount() float32`

GetRewardTriggerAmount returns the RewardTriggerAmount field if non-nil, zero value otherwise.

### GetRewardTriggerAmountOk

`func (o *MsaReturns) GetRewardTriggerAmountOk() (*float32, bool)`

GetRewardTriggerAmountOk returns a tuple with the RewardTriggerAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardTriggerAmount

`func (o *MsaReturns) SetRewardTriggerAmount(v float32)`

SetRewardTriggerAmount sets RewardTriggerAmount field to given value.


### GetStartDate

`func (o *MsaReturns) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *MsaReturns) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *MsaReturns) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *MsaReturns) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetToken

`func (o *MsaReturns) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *MsaReturns) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *MsaReturns) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *MsaReturns) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTransactionToken

`func (o *MsaReturns) GetTransactionToken() string`

GetTransactionToken returns the TransactionToken field if non-nil, zero value otherwise.

### GetTransactionTokenOk

`func (o *MsaReturns) GetTransactionTokenOk() (*string, bool)`

GetTransactionTokenOk returns a tuple with the TransactionToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionToken

`func (o *MsaReturns) SetTransactionToken(v string)`

SetTransactionToken sets TransactionToken field to given value.


### GetUnloadedAmount

`func (o *MsaReturns) GetUnloadedAmount() float32`

GetUnloadedAmount returns the UnloadedAmount field if non-nil, zero value otherwise.

### GetUnloadedAmountOk

`func (o *MsaReturns) GetUnloadedAmountOk() (*float32, bool)`

GetUnloadedAmountOk returns a tuple with the UnloadedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnloadedAmount

`func (o *MsaReturns) SetUnloadedAmount(v float32)`

SetUnloadedAmount sets UnloadedAmount field to given value.

### HasUnloadedAmount

`func (o *MsaReturns) HasUnloadedAmount() bool`

HasUnloadedAmount returns a boolean if a field has been set.

### GetUserToken

`func (o *MsaReturns) GetUserToken() string`

GetUserToken returns the UserToken field if non-nil, zero value otherwise.

### GetUserTokenOk

`func (o *MsaReturns) GetUserTokenOk() (*string, bool)`

GetUserTokenOk returns a tuple with the UserToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserToken

`func (o *MsaReturns) SetUserToken(v string)`

SetUserToken sets UserToken field to given value.

### HasUserToken

`func (o *MsaReturns) HasUserToken() bool`

HasUserToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


