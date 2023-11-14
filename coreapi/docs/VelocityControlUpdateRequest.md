# VelocityControlUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Indicates whether the velocity control is active. | [optional] 
**AmountLimit** | Pointer to **float32** | Maximum monetary sum that can be cleared within the time period defined by the &#x60;velocity_window&#x60; field. | [optional] 
**ApprovalsOnly** | Pointer to **bool** | If set to &#x60;true&#x60;, only approved transactions are subject to control. | [optional] 
**Association** | Pointer to [**SpendControlAssociation**](SpendControlAssociation.md) |  | [optional] 
**CurrencyCode** | Pointer to **string** | Three-character ISO 4217 currency code. | [optional] 
**IncludeCashback** | Pointer to **bool** | If set to &#x60;true&#x60;, cashbacks from a point of sale are subject to control. | [optional] 
**IncludeCredits** | Pointer to **bool** | If set to &#x60;true&#x60;, original credit transactions are subject to control. Your request can contain either a &#x60;money_in_transaction&#x60; object or the &#x60;include_credits&#x60; field, not both. | [optional] 
**IncludePurchases** | Pointer to **bool** | If set to &#x60;true&#x60;, purchases are subject to control. | [optional] 
**IncludeTransfers** | Pointer to **bool** | If set to &#x60;true&#x60;, transfers are subject to control. | [optional] 
**IncludeWithdrawals** | Pointer to **bool** | If set to &#x60;true&#x60;, ATM withdrawals are subject to control. | [optional] 
**MerchantScope** | Pointer to [**MerchantScope**](MerchantScope.md) |  | [optional] 
**MoneyInTransaction** | Pointer to [**MoneyInTransaction**](MoneyInTransaction.md) |  | [optional] 
**Name** | Pointer to **string** | Description of how the velocity control restricts spending. For example, \&quot;Max spend of $500 per day\&quot; or \&quot;Max spend of $5000 per month for non-exempt employees\&quot;.  Ensure that the description you provide here adequately captures the kind of restriction exerted by this velocity control, because the Marqeta platform will send this information to you in a webhook in the event that the transaction authorization attempt is declined by the velocity control.  *NOTE:* This field is very important. If your program has multiple velocity controls in place, it is not always clear which one prevented the transaction from being approved. Adding specific details to this field gives you more contextual information when debugging or monitoring declined authorization attempts. | [optional] 
**Token** | **string** | Unique identifier of the velocity control resource. | 
**UsageLimit** | Pointer to **int32** | Maximum number of times a card can be used within the time period defined by the &#x60;velocity_window&#x60; field.  If &#x60;velocity_window&#x60; is set to &#x60;TRANSACTION&#x60;, do not include a &#x60;usage_limit&#x60; in your request. | [optional] 
**VelocityWindow** | Pointer to **string** | Defines the time period to which the &#x60;amount_limit&#x60; and &#x60;usage_limit&#x60; fields apply:  * *DAY*  one day; days begin at 00:00:00 UTC. * *WEEK*  one week; weeks begin Sundays at 00:00:00 UTC. * *MONTH*  one month; months begin on the first day of month at 00:00:00 UTC. * *QUARTER* - three months; quarters begin on January 1, April 1, July 1, and October 1 at 00:00:00 UTC. * *YEAR* - twelve months; years begin on January 1 at 00:00:00 UTC. * *LIFETIME*  forever; time period never expires. * *TRANSACTION*  a single transaction.  // (2023-05-03): This statement was validated by Processing, as part of a customer inquiry. *NOTE:* If set to &#x60;DAY&#x60;, &#x60;WEEK&#x60;, or &#x60;MONTH&#x60;, the velocity control takes effect retroactively from the beginning of the specified period. The amount and usage data already collected within the first period is counted toward the limits.  // (2023-05-03): Commenting this note out, as it is untrue in testing as reported by customers and confirmed by transaction engine team //_*NOTE:* Editing any of the following fields on a velocity control resets its usage and amount count to 0:  //_* &#x60;merchant_scope.mcc&#x60; //_* &#x60;merchant_scope.mid&#x60; //_* &#x60;merchant_scope.mcc_group&#x60; //_* &#x60;association.user_token&#x60; //_* &#x60;association.card_product_token&#x60; | [optional] 

## Methods

### NewVelocityControlUpdateRequest

`func NewVelocityControlUpdateRequest(token string, ) *VelocityControlUpdateRequest`

NewVelocityControlUpdateRequest instantiates a new VelocityControlUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVelocityControlUpdateRequestWithDefaults

`func NewVelocityControlUpdateRequestWithDefaults() *VelocityControlUpdateRequest`

NewVelocityControlUpdateRequestWithDefaults instantiates a new VelocityControlUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *VelocityControlUpdateRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *VelocityControlUpdateRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *VelocityControlUpdateRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *VelocityControlUpdateRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAmountLimit

`func (o *VelocityControlUpdateRequest) GetAmountLimit() float32`

GetAmountLimit returns the AmountLimit field if non-nil, zero value otherwise.

### GetAmountLimitOk

`func (o *VelocityControlUpdateRequest) GetAmountLimitOk() (*float32, bool)`

GetAmountLimitOk returns a tuple with the AmountLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountLimit

`func (o *VelocityControlUpdateRequest) SetAmountLimit(v float32)`

SetAmountLimit sets AmountLimit field to given value.

### HasAmountLimit

`func (o *VelocityControlUpdateRequest) HasAmountLimit() bool`

HasAmountLimit returns a boolean if a field has been set.

### GetApprovalsOnly

`func (o *VelocityControlUpdateRequest) GetApprovalsOnly() bool`

GetApprovalsOnly returns the ApprovalsOnly field if non-nil, zero value otherwise.

### GetApprovalsOnlyOk

`func (o *VelocityControlUpdateRequest) GetApprovalsOnlyOk() (*bool, bool)`

GetApprovalsOnlyOk returns a tuple with the ApprovalsOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalsOnly

`func (o *VelocityControlUpdateRequest) SetApprovalsOnly(v bool)`

SetApprovalsOnly sets ApprovalsOnly field to given value.

### HasApprovalsOnly

`func (o *VelocityControlUpdateRequest) HasApprovalsOnly() bool`

HasApprovalsOnly returns a boolean if a field has been set.

### GetAssociation

`func (o *VelocityControlUpdateRequest) GetAssociation() SpendControlAssociation`

GetAssociation returns the Association field if non-nil, zero value otherwise.

### GetAssociationOk

`func (o *VelocityControlUpdateRequest) GetAssociationOk() (*SpendControlAssociation, bool)`

GetAssociationOk returns a tuple with the Association field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociation

`func (o *VelocityControlUpdateRequest) SetAssociation(v SpendControlAssociation)`

SetAssociation sets Association field to given value.

### HasAssociation

`func (o *VelocityControlUpdateRequest) HasAssociation() bool`

HasAssociation returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *VelocityControlUpdateRequest) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *VelocityControlUpdateRequest) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *VelocityControlUpdateRequest) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *VelocityControlUpdateRequest) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetIncludeCashback

`func (o *VelocityControlUpdateRequest) GetIncludeCashback() bool`

GetIncludeCashback returns the IncludeCashback field if non-nil, zero value otherwise.

### GetIncludeCashbackOk

`func (o *VelocityControlUpdateRequest) GetIncludeCashbackOk() (*bool, bool)`

GetIncludeCashbackOk returns a tuple with the IncludeCashback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeCashback

`func (o *VelocityControlUpdateRequest) SetIncludeCashback(v bool)`

SetIncludeCashback sets IncludeCashback field to given value.

### HasIncludeCashback

`func (o *VelocityControlUpdateRequest) HasIncludeCashback() bool`

HasIncludeCashback returns a boolean if a field has been set.

### GetIncludeCredits

`func (o *VelocityControlUpdateRequest) GetIncludeCredits() bool`

GetIncludeCredits returns the IncludeCredits field if non-nil, zero value otherwise.

### GetIncludeCreditsOk

`func (o *VelocityControlUpdateRequest) GetIncludeCreditsOk() (*bool, bool)`

GetIncludeCreditsOk returns a tuple with the IncludeCredits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeCredits

`func (o *VelocityControlUpdateRequest) SetIncludeCredits(v bool)`

SetIncludeCredits sets IncludeCredits field to given value.

### HasIncludeCredits

`func (o *VelocityControlUpdateRequest) HasIncludeCredits() bool`

HasIncludeCredits returns a boolean if a field has been set.

### GetIncludePurchases

`func (o *VelocityControlUpdateRequest) GetIncludePurchases() bool`

GetIncludePurchases returns the IncludePurchases field if non-nil, zero value otherwise.

### GetIncludePurchasesOk

`func (o *VelocityControlUpdateRequest) GetIncludePurchasesOk() (*bool, bool)`

GetIncludePurchasesOk returns a tuple with the IncludePurchases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludePurchases

`func (o *VelocityControlUpdateRequest) SetIncludePurchases(v bool)`

SetIncludePurchases sets IncludePurchases field to given value.

### HasIncludePurchases

`func (o *VelocityControlUpdateRequest) HasIncludePurchases() bool`

HasIncludePurchases returns a boolean if a field has been set.

### GetIncludeTransfers

`func (o *VelocityControlUpdateRequest) GetIncludeTransfers() bool`

GetIncludeTransfers returns the IncludeTransfers field if non-nil, zero value otherwise.

### GetIncludeTransfersOk

`func (o *VelocityControlUpdateRequest) GetIncludeTransfersOk() (*bool, bool)`

GetIncludeTransfersOk returns a tuple with the IncludeTransfers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeTransfers

`func (o *VelocityControlUpdateRequest) SetIncludeTransfers(v bool)`

SetIncludeTransfers sets IncludeTransfers field to given value.

### HasIncludeTransfers

`func (o *VelocityControlUpdateRequest) HasIncludeTransfers() bool`

HasIncludeTransfers returns a boolean if a field has been set.

### GetIncludeWithdrawals

`func (o *VelocityControlUpdateRequest) GetIncludeWithdrawals() bool`

GetIncludeWithdrawals returns the IncludeWithdrawals field if non-nil, zero value otherwise.

### GetIncludeWithdrawalsOk

`func (o *VelocityControlUpdateRequest) GetIncludeWithdrawalsOk() (*bool, bool)`

GetIncludeWithdrawalsOk returns a tuple with the IncludeWithdrawals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeWithdrawals

`func (o *VelocityControlUpdateRequest) SetIncludeWithdrawals(v bool)`

SetIncludeWithdrawals sets IncludeWithdrawals field to given value.

### HasIncludeWithdrawals

`func (o *VelocityControlUpdateRequest) HasIncludeWithdrawals() bool`

HasIncludeWithdrawals returns a boolean if a field has been set.

### GetMerchantScope

`func (o *VelocityControlUpdateRequest) GetMerchantScope() MerchantScope`

GetMerchantScope returns the MerchantScope field if non-nil, zero value otherwise.

### GetMerchantScopeOk

`func (o *VelocityControlUpdateRequest) GetMerchantScopeOk() (*MerchantScope, bool)`

GetMerchantScopeOk returns a tuple with the MerchantScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantScope

`func (o *VelocityControlUpdateRequest) SetMerchantScope(v MerchantScope)`

SetMerchantScope sets MerchantScope field to given value.

### HasMerchantScope

`func (o *VelocityControlUpdateRequest) HasMerchantScope() bool`

HasMerchantScope returns a boolean if a field has been set.

### GetMoneyInTransaction

`func (o *VelocityControlUpdateRequest) GetMoneyInTransaction() MoneyInTransaction`

GetMoneyInTransaction returns the MoneyInTransaction field if non-nil, zero value otherwise.

### GetMoneyInTransactionOk

`func (o *VelocityControlUpdateRequest) GetMoneyInTransactionOk() (*MoneyInTransaction, bool)`

GetMoneyInTransactionOk returns a tuple with the MoneyInTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoneyInTransaction

`func (o *VelocityControlUpdateRequest) SetMoneyInTransaction(v MoneyInTransaction)`

SetMoneyInTransaction sets MoneyInTransaction field to given value.

### HasMoneyInTransaction

`func (o *VelocityControlUpdateRequest) HasMoneyInTransaction() bool`

HasMoneyInTransaction returns a boolean if a field has been set.

### GetName

`func (o *VelocityControlUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VelocityControlUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VelocityControlUpdateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VelocityControlUpdateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetToken

`func (o *VelocityControlUpdateRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *VelocityControlUpdateRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *VelocityControlUpdateRequest) SetToken(v string)`

SetToken sets Token field to given value.


### GetUsageLimit

`func (o *VelocityControlUpdateRequest) GetUsageLimit() int32`

GetUsageLimit returns the UsageLimit field if non-nil, zero value otherwise.

### GetUsageLimitOk

`func (o *VelocityControlUpdateRequest) GetUsageLimitOk() (*int32, bool)`

GetUsageLimitOk returns a tuple with the UsageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageLimit

`func (o *VelocityControlUpdateRequest) SetUsageLimit(v int32)`

SetUsageLimit sets UsageLimit field to given value.

### HasUsageLimit

`func (o *VelocityControlUpdateRequest) HasUsageLimit() bool`

HasUsageLimit returns a boolean if a field has been set.

### GetVelocityWindow

`func (o *VelocityControlUpdateRequest) GetVelocityWindow() string`

GetVelocityWindow returns the VelocityWindow field if non-nil, zero value otherwise.

### GetVelocityWindowOk

`func (o *VelocityControlUpdateRequest) GetVelocityWindowOk() (*string, bool)`

GetVelocityWindowOk returns a tuple with the VelocityWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVelocityWindow

`func (o *VelocityControlUpdateRequest) SetVelocityWindow(v string)`

SetVelocityWindow sets VelocityWindow field to given value.

### HasVelocityWindow

`func (o *VelocityControlUpdateRequest) HasVelocityWindow() bool`

HasVelocityWindow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


