# VelocityControlRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Indicates whether the velocity control is active. | [optional] 
**AmountLimit** | **float32** | Maximum monetary sum that can be cleared within the time period defined by the &#x60;velocity_window&#x60; field. | 
**ApprovalsOnly** | Pointer to **bool** | If set to &#x60;true&#x60;, only approved transactions are subject to control. | [optional] 
**Association** | Pointer to [**SpendControlAssociation**](SpendControlAssociation.md) |  | [optional] 
**CurrencyCode** | **string** | Three-character ISO 4217 currency code. | 
**IncludeCashback** | Pointer to **bool** | If set to &#x60;true&#x60;, cashbacks from a point of sale are subject to control. | [optional] 
**IncludeCredits** | Pointer to **bool** | If set to &#x60;true&#x60;, original credit transactions (OCT) are subject to control. Your request can contain either a &#x60;money_in_transaction&#x60; object or the &#x60;include_credits&#x60; field, not both. | [optional] 
**IncludePurchases** | Pointer to **bool** | If set to &#x60;true&#x60;, purchases are subject to control. | [optional] 
**IncludeTransfers** | Pointer to **bool** | If set to &#x60;true&#x60;, transfers are subject to control. | [optional] 
**IncludeWithdrawals** | Pointer to **bool** | If set to &#x60;true&#x60;, ATM withdrawals are subject to control. | [optional] 
**MerchantScope** | Pointer to [**MerchantScope**](MerchantScope.md) |  | [optional] 
**MoneyInTransaction** | Pointer to [**MoneyInTransaction**](MoneyInTransaction.md) |  | [optional] 
**Name** | Pointer to **string** | Description of how the velocity control restricts spending, for example, \&quot;Max spend of $500 per day\&quot; or \&quot;Max spend of $5000 per month for non-exempt employees\&quot;.  Ensure that the description you provide here adequately captures the kind of restriction exerted by this velocity control, because the Marqeta platform will send this information to you in a webhook in the event that the transaction authorization attempt is declined by the velocity control.  *NOTE:* This field is very important. If your program has multiple velocity controls in place, it is not always clear which one prevented the transaction from being approved. Adding specific details to this field gives you more contextual information when debugging or monitoring declined authorization attempts. | [optional] 
**Token** | Pointer to **string** | Unique identifier of the velocity control.  If you do not include a token, the system will generate one automatically. This token is necessary for use in other API calls, so we recommend that rather than let the system generate one, you use a simple string that is easy to remember. This value cannot be updated. | [optional] 
**UsageLimit** | Pointer to **int32** | Maximum number of times a card can be used within the time period defined by the &#x60;velocity_window&#x60; field.  If &#x60;velocity_window&#x60; is set to &#x60;TRANSACTION&#x60;, do not include a &#x60;usage_limit&#x60; in your request.  Set to &#x60;-1&#x60; to indicate that the card can be used an unlimited number of times. | [optional] 
**VelocityWindow** | **string** | Defines the time period to which the &#x60;amount_limit&#x60; and &#x60;usage_limit&#x60; fields apply:  * *DAY*  one day; days begin at 00:00:00 UTC. * *WEEK*  one week; weeks begin Sundays at 00:00:00 UTC. * *MONTH*  one month; months begin on the first day of month at 00:00:00 UTC. * *LIFETIME*  forever; time period never expires. * *TRANSACTION*  a single transaction.  // (2023-05-03): This statement was validated by Processing, as part of a customer inquiry. *NOTE:* If set to &#x60;DAY&#x60;, &#x60;WEEK&#x60;, or &#x60;MONTH&#x60;, the velocity control takes effect retroactively from the beginning of the specified period. The amount and usage data already collected within the first period is counted toward the limits.  // (2023-05-03): Commenting this note out, as it is untrue in testing as reported by customers and confirmed by transaction engine team //_*NOTE:* Editing any of the following fields on a velocity control resets its usage and amount count to 0:  //_* &#x60;merchant_scope.mcc&#x60; //_* &#x60;merchant_scope.mid&#x60; //_* &#x60;merchant_scope.mcc_group&#x60; //_* &#x60;association.user_token&#x60; //_* &#x60;association.card_product_token&#x60; | 

## Methods

### NewVelocityControlRequest

`func NewVelocityControlRequest(amountLimit float32, currencyCode string, velocityWindow string, ) *VelocityControlRequest`

NewVelocityControlRequest instantiates a new VelocityControlRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVelocityControlRequestWithDefaults

`func NewVelocityControlRequestWithDefaults() *VelocityControlRequest`

NewVelocityControlRequestWithDefaults instantiates a new VelocityControlRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *VelocityControlRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *VelocityControlRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *VelocityControlRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *VelocityControlRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAmountLimit

`func (o *VelocityControlRequest) GetAmountLimit() float32`

GetAmountLimit returns the AmountLimit field if non-nil, zero value otherwise.

### GetAmountLimitOk

`func (o *VelocityControlRequest) GetAmountLimitOk() (*float32, bool)`

GetAmountLimitOk returns a tuple with the AmountLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountLimit

`func (o *VelocityControlRequest) SetAmountLimit(v float32)`

SetAmountLimit sets AmountLimit field to given value.


### GetApprovalsOnly

`func (o *VelocityControlRequest) GetApprovalsOnly() bool`

GetApprovalsOnly returns the ApprovalsOnly field if non-nil, zero value otherwise.

### GetApprovalsOnlyOk

`func (o *VelocityControlRequest) GetApprovalsOnlyOk() (*bool, bool)`

GetApprovalsOnlyOk returns a tuple with the ApprovalsOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalsOnly

`func (o *VelocityControlRequest) SetApprovalsOnly(v bool)`

SetApprovalsOnly sets ApprovalsOnly field to given value.

### HasApprovalsOnly

`func (o *VelocityControlRequest) HasApprovalsOnly() bool`

HasApprovalsOnly returns a boolean if a field has been set.

### GetAssociation

`func (o *VelocityControlRequest) GetAssociation() SpendControlAssociation`

GetAssociation returns the Association field if non-nil, zero value otherwise.

### GetAssociationOk

`func (o *VelocityControlRequest) GetAssociationOk() (*SpendControlAssociation, bool)`

GetAssociationOk returns a tuple with the Association field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociation

`func (o *VelocityControlRequest) SetAssociation(v SpendControlAssociation)`

SetAssociation sets Association field to given value.

### HasAssociation

`func (o *VelocityControlRequest) HasAssociation() bool`

HasAssociation returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *VelocityControlRequest) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *VelocityControlRequest) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *VelocityControlRequest) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.


### GetIncludeCashback

`func (o *VelocityControlRequest) GetIncludeCashback() bool`

GetIncludeCashback returns the IncludeCashback field if non-nil, zero value otherwise.

### GetIncludeCashbackOk

`func (o *VelocityControlRequest) GetIncludeCashbackOk() (*bool, bool)`

GetIncludeCashbackOk returns a tuple with the IncludeCashback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeCashback

`func (o *VelocityControlRequest) SetIncludeCashback(v bool)`

SetIncludeCashback sets IncludeCashback field to given value.

### HasIncludeCashback

`func (o *VelocityControlRequest) HasIncludeCashback() bool`

HasIncludeCashback returns a boolean if a field has been set.

### GetIncludeCredits

`func (o *VelocityControlRequest) GetIncludeCredits() bool`

GetIncludeCredits returns the IncludeCredits field if non-nil, zero value otherwise.

### GetIncludeCreditsOk

`func (o *VelocityControlRequest) GetIncludeCreditsOk() (*bool, bool)`

GetIncludeCreditsOk returns a tuple with the IncludeCredits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeCredits

`func (o *VelocityControlRequest) SetIncludeCredits(v bool)`

SetIncludeCredits sets IncludeCredits field to given value.

### HasIncludeCredits

`func (o *VelocityControlRequest) HasIncludeCredits() bool`

HasIncludeCredits returns a boolean if a field has been set.

### GetIncludePurchases

`func (o *VelocityControlRequest) GetIncludePurchases() bool`

GetIncludePurchases returns the IncludePurchases field if non-nil, zero value otherwise.

### GetIncludePurchasesOk

`func (o *VelocityControlRequest) GetIncludePurchasesOk() (*bool, bool)`

GetIncludePurchasesOk returns a tuple with the IncludePurchases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludePurchases

`func (o *VelocityControlRequest) SetIncludePurchases(v bool)`

SetIncludePurchases sets IncludePurchases field to given value.

### HasIncludePurchases

`func (o *VelocityControlRequest) HasIncludePurchases() bool`

HasIncludePurchases returns a boolean if a field has been set.

### GetIncludeTransfers

`func (o *VelocityControlRequest) GetIncludeTransfers() bool`

GetIncludeTransfers returns the IncludeTransfers field if non-nil, zero value otherwise.

### GetIncludeTransfersOk

`func (o *VelocityControlRequest) GetIncludeTransfersOk() (*bool, bool)`

GetIncludeTransfersOk returns a tuple with the IncludeTransfers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeTransfers

`func (o *VelocityControlRequest) SetIncludeTransfers(v bool)`

SetIncludeTransfers sets IncludeTransfers field to given value.

### HasIncludeTransfers

`func (o *VelocityControlRequest) HasIncludeTransfers() bool`

HasIncludeTransfers returns a boolean if a field has been set.

### GetIncludeWithdrawals

`func (o *VelocityControlRequest) GetIncludeWithdrawals() bool`

GetIncludeWithdrawals returns the IncludeWithdrawals field if non-nil, zero value otherwise.

### GetIncludeWithdrawalsOk

`func (o *VelocityControlRequest) GetIncludeWithdrawalsOk() (*bool, bool)`

GetIncludeWithdrawalsOk returns a tuple with the IncludeWithdrawals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeWithdrawals

`func (o *VelocityControlRequest) SetIncludeWithdrawals(v bool)`

SetIncludeWithdrawals sets IncludeWithdrawals field to given value.

### HasIncludeWithdrawals

`func (o *VelocityControlRequest) HasIncludeWithdrawals() bool`

HasIncludeWithdrawals returns a boolean if a field has been set.

### GetMerchantScope

`func (o *VelocityControlRequest) GetMerchantScope() MerchantScope`

GetMerchantScope returns the MerchantScope field if non-nil, zero value otherwise.

### GetMerchantScopeOk

`func (o *VelocityControlRequest) GetMerchantScopeOk() (*MerchantScope, bool)`

GetMerchantScopeOk returns a tuple with the MerchantScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantScope

`func (o *VelocityControlRequest) SetMerchantScope(v MerchantScope)`

SetMerchantScope sets MerchantScope field to given value.

### HasMerchantScope

`func (o *VelocityControlRequest) HasMerchantScope() bool`

HasMerchantScope returns a boolean if a field has been set.

### GetMoneyInTransaction

`func (o *VelocityControlRequest) GetMoneyInTransaction() MoneyInTransaction`

GetMoneyInTransaction returns the MoneyInTransaction field if non-nil, zero value otherwise.

### GetMoneyInTransactionOk

`func (o *VelocityControlRequest) GetMoneyInTransactionOk() (*MoneyInTransaction, bool)`

GetMoneyInTransactionOk returns a tuple with the MoneyInTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoneyInTransaction

`func (o *VelocityControlRequest) SetMoneyInTransaction(v MoneyInTransaction)`

SetMoneyInTransaction sets MoneyInTransaction field to given value.

### HasMoneyInTransaction

`func (o *VelocityControlRequest) HasMoneyInTransaction() bool`

HasMoneyInTransaction returns a boolean if a field has been set.

### GetName

`func (o *VelocityControlRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VelocityControlRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VelocityControlRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VelocityControlRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetToken

`func (o *VelocityControlRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *VelocityControlRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *VelocityControlRequest) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *VelocityControlRequest) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUsageLimit

`func (o *VelocityControlRequest) GetUsageLimit() int32`

GetUsageLimit returns the UsageLimit field if non-nil, zero value otherwise.

### GetUsageLimitOk

`func (o *VelocityControlRequest) GetUsageLimitOk() (*int32, bool)`

GetUsageLimitOk returns a tuple with the UsageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageLimit

`func (o *VelocityControlRequest) SetUsageLimit(v int32)`

SetUsageLimit sets UsageLimit field to given value.

### HasUsageLimit

`func (o *VelocityControlRequest) HasUsageLimit() bool`

HasUsageLimit returns a boolean if a field has been set.

### GetVelocityWindow

`func (o *VelocityControlRequest) GetVelocityWindow() string`

GetVelocityWindow returns the VelocityWindow field if non-nil, zero value otherwise.

### GetVelocityWindowOk

`func (o *VelocityControlRequest) GetVelocityWindowOk() (*string, bool)`

GetVelocityWindowOk returns a tuple with the VelocityWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVelocityWindow

`func (o *VelocityControlRequest) SetVelocityWindow(v string)`

SetVelocityWindow sets VelocityWindow field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


