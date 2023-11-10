# VelocityControlResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Indicates whether the velocity control is active. | [optional] 
**AmountLimit** | **float32** | Maximum monetary sum that can be cleared within the time period defined by the &#x60;velocity_window&#x60; field. | 
**ApprovalsOnly** | Pointer to **bool** | If set to &#x60;true&#x60;, only approved transactions are subject to control. | [optional] 
**Association** | Pointer to [**SpendControlAssociation**](SpendControlAssociation.md) |  | [optional] 
**CurrencyCode** | **string** | Three-character ISO 4217 currency code. | 
**IncludeCashback** | Pointer to **bool** | If set to &#x60;true&#x60;, cashbacks from a point of sale are subject to control. | [optional] 
**IncludeCredits** | Pointer to **bool** | If set to &#x60;true&#x60;, original credit transactions (OCT) are subject to control. | [optional] 
**IncludePurchases** | Pointer to **bool** | If set to &#x60;true&#x60;, purchases are subject to control. | [optional] 
**IncludeTransfers** | Pointer to **bool** | If set to &#x60;true&#x60;, transfers are subject to control. | [optional] 
**IncludeWithdrawals** | Pointer to **bool** | If set to &#x60;true&#x60;, ATM withdrawals are subject to control. | [optional] 
**MerchantScope** | Pointer to [**MerchantScope**](MerchantScope.md) |  | [optional] 
**MoneyInTransaction** | Pointer to [**MoneyInTransaction**](MoneyInTransaction.md) |  | [optional] 
**Name** | Pointer to **string** | Description of how the velocity control restricts spending. For example, \&quot;Max spend of $500 per day\&quot; or \&quot;Max spend of $5000 per month for non-exempt employees\&quot;. | [optional] 
**Token** | Pointer to **string** | Unique identifier of the velocity control. | [optional] 
**UsageLimit** | Pointer to **int32** | Maximum number of times a card can be used within the time period defined by the &#x60;velocity_window&#x60; field. | [optional] 
**VelocityWindow** | **string** | Defines the time period to which the &#x60;amount_limit&#x60; and &#x60;usage_limit&#x60; fields apply:  * *DAY* – one day; days begin at 00:00:00 UTC. * *WEEK* – one week; weeks begin Sundays at 00:00:00 UTC. * *MONTH* – one month; months begin on the first day of month at 00:00:00 UTC. * *LIFETIME* – forever; time period never expires. * *TRANSACTION* – a single transaction. | 

## Methods

### NewVelocityControlResponse

`func NewVelocityControlResponse(amountLimit float32, currencyCode string, velocityWindow string, ) *VelocityControlResponse`

NewVelocityControlResponse instantiates a new VelocityControlResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVelocityControlResponseWithDefaults

`func NewVelocityControlResponseWithDefaults() *VelocityControlResponse`

NewVelocityControlResponseWithDefaults instantiates a new VelocityControlResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *VelocityControlResponse) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *VelocityControlResponse) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *VelocityControlResponse) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *VelocityControlResponse) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAmountLimit

`func (o *VelocityControlResponse) GetAmountLimit() float32`

GetAmountLimit returns the AmountLimit field if non-nil, zero value otherwise.

### GetAmountLimitOk

`func (o *VelocityControlResponse) GetAmountLimitOk() (*float32, bool)`

GetAmountLimitOk returns a tuple with the AmountLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountLimit

`func (o *VelocityControlResponse) SetAmountLimit(v float32)`

SetAmountLimit sets AmountLimit field to given value.


### GetApprovalsOnly

`func (o *VelocityControlResponse) GetApprovalsOnly() bool`

GetApprovalsOnly returns the ApprovalsOnly field if non-nil, zero value otherwise.

### GetApprovalsOnlyOk

`func (o *VelocityControlResponse) GetApprovalsOnlyOk() (*bool, bool)`

GetApprovalsOnlyOk returns a tuple with the ApprovalsOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalsOnly

`func (o *VelocityControlResponse) SetApprovalsOnly(v bool)`

SetApprovalsOnly sets ApprovalsOnly field to given value.

### HasApprovalsOnly

`func (o *VelocityControlResponse) HasApprovalsOnly() bool`

HasApprovalsOnly returns a boolean if a field has been set.

### GetAssociation

`func (o *VelocityControlResponse) GetAssociation() SpendControlAssociation`

GetAssociation returns the Association field if non-nil, zero value otherwise.

### GetAssociationOk

`func (o *VelocityControlResponse) GetAssociationOk() (*SpendControlAssociation, bool)`

GetAssociationOk returns a tuple with the Association field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociation

`func (o *VelocityControlResponse) SetAssociation(v SpendControlAssociation)`

SetAssociation sets Association field to given value.

### HasAssociation

`func (o *VelocityControlResponse) HasAssociation() bool`

HasAssociation returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *VelocityControlResponse) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *VelocityControlResponse) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *VelocityControlResponse) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.


### GetIncludeCashback

`func (o *VelocityControlResponse) GetIncludeCashback() bool`

GetIncludeCashback returns the IncludeCashback field if non-nil, zero value otherwise.

### GetIncludeCashbackOk

`func (o *VelocityControlResponse) GetIncludeCashbackOk() (*bool, bool)`

GetIncludeCashbackOk returns a tuple with the IncludeCashback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeCashback

`func (o *VelocityControlResponse) SetIncludeCashback(v bool)`

SetIncludeCashback sets IncludeCashback field to given value.

### HasIncludeCashback

`func (o *VelocityControlResponse) HasIncludeCashback() bool`

HasIncludeCashback returns a boolean if a field has been set.

### GetIncludeCredits

`func (o *VelocityControlResponse) GetIncludeCredits() bool`

GetIncludeCredits returns the IncludeCredits field if non-nil, zero value otherwise.

### GetIncludeCreditsOk

`func (o *VelocityControlResponse) GetIncludeCreditsOk() (*bool, bool)`

GetIncludeCreditsOk returns a tuple with the IncludeCredits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeCredits

`func (o *VelocityControlResponse) SetIncludeCredits(v bool)`

SetIncludeCredits sets IncludeCredits field to given value.

### HasIncludeCredits

`func (o *VelocityControlResponse) HasIncludeCredits() bool`

HasIncludeCredits returns a boolean if a field has been set.

### GetIncludePurchases

`func (o *VelocityControlResponse) GetIncludePurchases() bool`

GetIncludePurchases returns the IncludePurchases field if non-nil, zero value otherwise.

### GetIncludePurchasesOk

`func (o *VelocityControlResponse) GetIncludePurchasesOk() (*bool, bool)`

GetIncludePurchasesOk returns a tuple with the IncludePurchases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludePurchases

`func (o *VelocityControlResponse) SetIncludePurchases(v bool)`

SetIncludePurchases sets IncludePurchases field to given value.

### HasIncludePurchases

`func (o *VelocityControlResponse) HasIncludePurchases() bool`

HasIncludePurchases returns a boolean if a field has been set.

### GetIncludeTransfers

`func (o *VelocityControlResponse) GetIncludeTransfers() bool`

GetIncludeTransfers returns the IncludeTransfers field if non-nil, zero value otherwise.

### GetIncludeTransfersOk

`func (o *VelocityControlResponse) GetIncludeTransfersOk() (*bool, bool)`

GetIncludeTransfersOk returns a tuple with the IncludeTransfers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeTransfers

`func (o *VelocityControlResponse) SetIncludeTransfers(v bool)`

SetIncludeTransfers sets IncludeTransfers field to given value.

### HasIncludeTransfers

`func (o *VelocityControlResponse) HasIncludeTransfers() bool`

HasIncludeTransfers returns a boolean if a field has been set.

### GetIncludeWithdrawals

`func (o *VelocityControlResponse) GetIncludeWithdrawals() bool`

GetIncludeWithdrawals returns the IncludeWithdrawals field if non-nil, zero value otherwise.

### GetIncludeWithdrawalsOk

`func (o *VelocityControlResponse) GetIncludeWithdrawalsOk() (*bool, bool)`

GetIncludeWithdrawalsOk returns a tuple with the IncludeWithdrawals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeWithdrawals

`func (o *VelocityControlResponse) SetIncludeWithdrawals(v bool)`

SetIncludeWithdrawals sets IncludeWithdrawals field to given value.

### HasIncludeWithdrawals

`func (o *VelocityControlResponse) HasIncludeWithdrawals() bool`

HasIncludeWithdrawals returns a boolean if a field has been set.

### GetMerchantScope

`func (o *VelocityControlResponse) GetMerchantScope() MerchantScope`

GetMerchantScope returns the MerchantScope field if non-nil, zero value otherwise.

### GetMerchantScopeOk

`func (o *VelocityControlResponse) GetMerchantScopeOk() (*MerchantScope, bool)`

GetMerchantScopeOk returns a tuple with the MerchantScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantScope

`func (o *VelocityControlResponse) SetMerchantScope(v MerchantScope)`

SetMerchantScope sets MerchantScope field to given value.

### HasMerchantScope

`func (o *VelocityControlResponse) HasMerchantScope() bool`

HasMerchantScope returns a boolean if a field has been set.

### GetMoneyInTransaction

`func (o *VelocityControlResponse) GetMoneyInTransaction() MoneyInTransaction`

GetMoneyInTransaction returns the MoneyInTransaction field if non-nil, zero value otherwise.

### GetMoneyInTransactionOk

`func (o *VelocityControlResponse) GetMoneyInTransactionOk() (*MoneyInTransaction, bool)`

GetMoneyInTransactionOk returns a tuple with the MoneyInTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoneyInTransaction

`func (o *VelocityControlResponse) SetMoneyInTransaction(v MoneyInTransaction)`

SetMoneyInTransaction sets MoneyInTransaction field to given value.

### HasMoneyInTransaction

`func (o *VelocityControlResponse) HasMoneyInTransaction() bool`

HasMoneyInTransaction returns a boolean if a field has been set.

### GetName

`func (o *VelocityControlResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VelocityControlResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VelocityControlResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VelocityControlResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetToken

`func (o *VelocityControlResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *VelocityControlResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *VelocityControlResponse) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *VelocityControlResponse) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUsageLimit

`func (o *VelocityControlResponse) GetUsageLimit() int32`

GetUsageLimit returns the UsageLimit field if non-nil, zero value otherwise.

### GetUsageLimitOk

`func (o *VelocityControlResponse) GetUsageLimitOk() (*int32, bool)`

GetUsageLimitOk returns a tuple with the UsageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageLimit

`func (o *VelocityControlResponse) SetUsageLimit(v int32)`

SetUsageLimit sets UsageLimit field to given value.

### HasUsageLimit

`func (o *VelocityControlResponse) HasUsageLimit() bool`

HasUsageLimit returns a boolean if a field has been set.

### GetVelocityWindow

`func (o *VelocityControlResponse) GetVelocityWindow() string`

GetVelocityWindow returns the VelocityWindow field if non-nil, zero value otherwise.

### GetVelocityWindowOk

`func (o *VelocityControlResponse) GetVelocityWindowOk() (*string, bool)`

GetVelocityWindowOk returns a tuple with the VelocityWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVelocityWindow

`func (o *VelocityControlResponse) SetVelocityWindow(v string)`

SetVelocityWindow sets VelocityWindow field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


