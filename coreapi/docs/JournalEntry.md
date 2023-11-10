# JournalEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountToken** | **string** | Unique identifier of the credit account associated with the credit card used to make the journal entry. | 
**Amount** | **float32** | Amount of the journal entry. | 
**CardToken** | **string** | Unique identifier of the credit card used to make the journal entry. | 
**CreatedTime** | **time.Time** | Date and time when the journal entry was created on Marqeta&#39;s credit platform, in UTC. | 
**CurrencyCode** | [**CurrencyCode**](CurrencyCode.md) |  | [default to CURRENCYCODE_USD]
**DetailObject** | Pointer to **map[string]interface{}** | Contains the journal entry&#39;s full details. The fields returned in this object vary based on the journal entry group.  The following lists each journal entry group and the specific fields returned for each group.  * Purchases and refunds: see the &lt;&lt;/core-api/transactions#getTransactions, transactions&gt;&gt; response fields.  * Disputes: see the &lt;&lt;/core-api/credit-disputes#retrieveDispute, account disputes response fields.&gt;&gt;  * Original credit transaction (OCT): see the &lt;&lt;/core-api/push-to-card-payments#_create_push_to_card_disbursement, Push-to-Card disbursement&gt;&gt; fields.  * Rewards: see the &lt;&lt;/core-api/credit-account-rewards#createReward, account reward&gt;&gt; response fields.  * Payments: see the &lt;&lt;/core-api/credit-account-payments#retrievePayment, account payment&gt;&gt; response fields.  * Balance refunds: see the &lt;&lt;/core-api/credit-balance-refunds#createBalanceRefund, balance refund&gt;&gt; response fields.  * Adjustments: see the &lt;&lt;/core-api/credit-account-adjustments#retrieveAdjustment, account adjustment&gt;&gt; response fields.  * Interest and fees: see fields below. | [optional] 
**DetailToken** | **string** | Unique identifier of the journal entry&#39;s full details. | 
**DisputeToken** | Pointer to **string** | Unique identifier of the dispute, if the journal entry is disputed. | [optional] 
**Group** | **string** | Group to which the journal entry belongs. | 
**Id** | **string** | Eight-digit numeric identifier of the journal entry, an alternate identifier to the UUID that is useful for remembering and referencing. | 
**ImpactTime** | **time.Time** | Date and time when the journal entry impacts the account balance.  For purchases, this is the time of the authorization.  For purchase authorization clearings, this is the time when the transaction is settled. | 
**Memo** | **string** | Merchant name or description for the journal entry. | 
**RelatedToken** | Pointer to **string** | Unique identifier of the original journal entry. If the current journal entry is the original, this field is returned empty. | [optional] 
**RequestTime** | **time.Time** | For purchases, the date and time of the authorization, which is when the user initiates the journal entry.  For other journal entry groups, equivalent to &#x60;impact_time&#x60;. | 
**RootToken** | Pointer to **string** | Unique identifier of the root journal entry. If the current journal entry is the root, this field is returned empty. | [optional] 
**Status** | **string** | Status of the journal entry when it was initially recorded and had an impact on the balance, either &#x60;PENDING&#x60; or &#x60;POSTED&#x60;. This field is immutable and may not represent the current status.  To view the current status of purchases, refunds, OCTs, and payments, see the &#x60;detail_object.state&#x60; field. These journal entries start in &#x60;PENDING&#x60; and can transition to &#x60;CLEARED&#x60;, &#x60;DECLINED&#x60;, or &#x60;ERROR&#x60;. This transition of status is only sent through webhook event notifications.  Journal entries that are final transactions, such as clearings, start and remain in a &#x60;POSTED&#x60; state.  *NOTE*: &#x60;CLEARED&#x60;, &#x60;DECLINED&#x60;, and &#x60;ERROR&#x60; are special statuses that do not have an impact on the account balance, and are not recorded in the journal. For these special statuses, &#x60;impact_time&#x60;, &#x60;request_time&#x60;, &#x60;created_time&#x60;, &#x60;token&#x60;, and &#x60;id&#x60; are returned blank. | 
**Token** | **string** | Unique identifier of the journal entry. | 
**Type** | **string** | The journal entry event type. | 
**UserToken** | **string** | Unique identifier of the credit user. | 

## Methods

### NewJournalEntry

`func NewJournalEntry(accountToken string, amount float32, cardToken string, createdTime time.Time, currencyCode CurrencyCode, detailToken string, group string, id string, impactTime time.Time, memo string, requestTime time.Time, status string, token string, type_ string, userToken string, ) *JournalEntry`

NewJournalEntry instantiates a new JournalEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJournalEntryWithDefaults

`func NewJournalEntryWithDefaults() *JournalEntry`

NewJournalEntryWithDefaults instantiates a new JournalEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountToken

`func (o *JournalEntry) GetAccountToken() string`

GetAccountToken returns the AccountToken field if non-nil, zero value otherwise.

### GetAccountTokenOk

`func (o *JournalEntry) GetAccountTokenOk() (*string, bool)`

GetAccountTokenOk returns a tuple with the AccountToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountToken

`func (o *JournalEntry) SetAccountToken(v string)`

SetAccountToken sets AccountToken field to given value.


### GetAmount

`func (o *JournalEntry) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *JournalEntry) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *JournalEntry) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetCardToken

`func (o *JournalEntry) GetCardToken() string`

GetCardToken returns the CardToken field if non-nil, zero value otherwise.

### GetCardTokenOk

`func (o *JournalEntry) GetCardTokenOk() (*string, bool)`

GetCardTokenOk returns a tuple with the CardToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardToken

`func (o *JournalEntry) SetCardToken(v string)`

SetCardToken sets CardToken field to given value.


### GetCreatedTime

`func (o *JournalEntry) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *JournalEntry) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *JournalEntry) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetCurrencyCode

`func (o *JournalEntry) GetCurrencyCode() CurrencyCode`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *JournalEntry) GetCurrencyCodeOk() (*CurrencyCode, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *JournalEntry) SetCurrencyCode(v CurrencyCode)`

SetCurrencyCode sets CurrencyCode field to given value.


### GetDetailObject

`func (o *JournalEntry) GetDetailObject() map[string]interface{}`

GetDetailObject returns the DetailObject field if non-nil, zero value otherwise.

### GetDetailObjectOk

`func (o *JournalEntry) GetDetailObjectOk() (*map[string]interface{}, bool)`

GetDetailObjectOk returns a tuple with the DetailObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailObject

`func (o *JournalEntry) SetDetailObject(v map[string]interface{})`

SetDetailObject sets DetailObject field to given value.

### HasDetailObject

`func (o *JournalEntry) HasDetailObject() bool`

HasDetailObject returns a boolean if a field has been set.

### GetDetailToken

`func (o *JournalEntry) GetDetailToken() string`

GetDetailToken returns the DetailToken field if non-nil, zero value otherwise.

### GetDetailTokenOk

`func (o *JournalEntry) GetDetailTokenOk() (*string, bool)`

GetDetailTokenOk returns a tuple with the DetailToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailToken

`func (o *JournalEntry) SetDetailToken(v string)`

SetDetailToken sets DetailToken field to given value.


### GetDisputeToken

`func (o *JournalEntry) GetDisputeToken() string`

GetDisputeToken returns the DisputeToken field if non-nil, zero value otherwise.

### GetDisputeTokenOk

`func (o *JournalEntry) GetDisputeTokenOk() (*string, bool)`

GetDisputeTokenOk returns a tuple with the DisputeToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisputeToken

`func (o *JournalEntry) SetDisputeToken(v string)`

SetDisputeToken sets DisputeToken field to given value.

### HasDisputeToken

`func (o *JournalEntry) HasDisputeToken() bool`

HasDisputeToken returns a boolean if a field has been set.

### GetGroup

`func (o *JournalEntry) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *JournalEntry) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *JournalEntry) SetGroup(v string)`

SetGroup sets Group field to given value.


### GetId

`func (o *JournalEntry) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JournalEntry) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JournalEntry) SetId(v string)`

SetId sets Id field to given value.


### GetImpactTime

`func (o *JournalEntry) GetImpactTime() time.Time`

GetImpactTime returns the ImpactTime field if non-nil, zero value otherwise.

### GetImpactTimeOk

`func (o *JournalEntry) GetImpactTimeOk() (*time.Time, bool)`

GetImpactTimeOk returns a tuple with the ImpactTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpactTime

`func (o *JournalEntry) SetImpactTime(v time.Time)`

SetImpactTime sets ImpactTime field to given value.


### GetMemo

`func (o *JournalEntry) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *JournalEntry) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *JournalEntry) SetMemo(v string)`

SetMemo sets Memo field to given value.


### GetRelatedToken

`func (o *JournalEntry) GetRelatedToken() string`

GetRelatedToken returns the RelatedToken field if non-nil, zero value otherwise.

### GetRelatedTokenOk

`func (o *JournalEntry) GetRelatedTokenOk() (*string, bool)`

GetRelatedTokenOk returns a tuple with the RelatedToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedToken

`func (o *JournalEntry) SetRelatedToken(v string)`

SetRelatedToken sets RelatedToken field to given value.

### HasRelatedToken

`func (o *JournalEntry) HasRelatedToken() bool`

HasRelatedToken returns a boolean if a field has been set.

### GetRequestTime

`func (o *JournalEntry) GetRequestTime() time.Time`

GetRequestTime returns the RequestTime field if non-nil, zero value otherwise.

### GetRequestTimeOk

`func (o *JournalEntry) GetRequestTimeOk() (*time.Time, bool)`

GetRequestTimeOk returns a tuple with the RequestTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTime

`func (o *JournalEntry) SetRequestTime(v time.Time)`

SetRequestTime sets RequestTime field to given value.


### GetRootToken

`func (o *JournalEntry) GetRootToken() string`

GetRootToken returns the RootToken field if non-nil, zero value otherwise.

### GetRootTokenOk

`func (o *JournalEntry) GetRootTokenOk() (*string, bool)`

GetRootTokenOk returns a tuple with the RootToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootToken

`func (o *JournalEntry) SetRootToken(v string)`

SetRootToken sets RootToken field to given value.

### HasRootToken

`func (o *JournalEntry) HasRootToken() bool`

HasRootToken returns a boolean if a field has been set.

### GetStatus

`func (o *JournalEntry) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *JournalEntry) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *JournalEntry) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetToken

`func (o *JournalEntry) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *JournalEntry) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *JournalEntry) SetToken(v string)`

SetToken sets Token field to given value.


### GetType

`func (o *JournalEntry) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *JournalEntry) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *JournalEntry) SetType(v string)`

SetType sets Type field to given value.


### GetUserToken

`func (o *JournalEntry) GetUserToken() string`

GetUserToken returns the UserToken field if non-nil, zero value otherwise.

### GetUserTokenOk

`func (o *JournalEntry) GetUserTokenOk() (*string, bool)`

GetUserTokenOk returns a tuple with the UserToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserToken

`func (o *JournalEntry) SetUserToken(v string)`

SetUserToken sets UserToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


