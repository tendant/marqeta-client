# DelinquencyTransitionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountToken** | **string** | Unique identifier of the credit account. | 
**BucketCount** | Pointer to **float32** | Number of buckets for the account after the triggering event occurred. | [optional] 
**CreatedTime** | **time.Time** | Date and time when the delinquency state transition was created on Marqeta&#39;s credit platform, in UTC. | 
**CurrentDue** | Pointer to **float32** | Current amount that is due after the triggering event occurred.  Equivalent to &#x60;current_due&#x60; for the account&#39;s most recent delinquency bucket. To retrieve delinquency buckets for an account, send a &#x60;GET&#x60; request to &#x60;/credit/accounts/{account_token}/delinquencystate&#x60;. | [optional] 
**ImpactTime** | **time.Time** | Date and time when the triggering event impacted the account, in UTC. | 
**IsRolledBack** | **bool** | A value of &#x60;true&#x60; indicates that the system invalidated and rolled back the delinquency transition.  This is a temporary field that allows Marqeta to handle occasional cases of out-of-order processing. This can occur when two delinquency state transition webhooks are sent near-simultaneously.  For example, if a credit and a payment that bring an account current are made around the same time, two delinquency state transitions are sent very close together. In these cases, one of the transitions is rolled back and invalidated. For the transition that is rolled back, &#x60;is_rolled_back&#x60; is &#x60;true&#x60; and the transition should be ignored.  This field is temporary and to be deprecated when out-of-order processing is addressed in a future release. | 
**OldestPaymentDueDate** | Pointer to **time.Time** | The payment due date of the account&#39;s oldest delinquency bucket, in UTC.  Useful when used with the delinquency state transition&#39;s &#x60;created_time&#x60; to determine the total number of days a payment is past due. | [optional] 
**OriginalStatus** | [**DelinquencyStatus**](DelinquencyStatus.md) |  | 
**Status** | [**DelinquencyStatus**](DelinquencyStatus.md) |  | 
**Token** | **string** | Unique identifier of the delinquency state transition. | 
**TotalDue** | Pointer to **float32** | Total amount that is due after the triggering event occurred; the sum of &#x60;total_past_due&#x60; and &#x60;current_due&#x60;.  Equivalent to &#x60;total_due&#x60; for the account&#39;s most recent delinquency bucket. To retrieve delinquency buckets for an account, send a &#x60;GET&#x60; request to &#x60;/credit/accounts/{account_token}/delinquencystate&#x60;. | [optional] 
**TotalPastDue** | Pointer to **float32** | Total amount that is past due after the triggering event occurred.  Equivalent to &#x60;past_due_carried_forward&#x60; for the account&#39;s most recent delinquency bucket. To retrieve delinquency buckets for an account, send a &#x60;GET&#x60; request to &#x60;/credit/accounts/{account_token}/delinquencystate&#x60;. | [optional] 
**TransitionTriggerReason** | [**DelinquencyTransitionTriggerReason**](DelinquencyTransitionTriggerReason.md) |  | 
**TransitionTriggerTime** | **time.Time** | Date and time when the triggering event caused the account&#39;s delinquency state to transition, in UTC.  For &lt;&lt;/core-api/credit-account-journal-entries, journal entries&gt;&gt;, equivalent to &#x60;request_time&#x60;. For &lt;&lt;/core-api/credit-account-statements#listStatementJournalEntries, statement journal entries&gt;&gt;, equivalent to &#x60;impact_time&#x60;, | 
**UpdatedTime** | Pointer to **time.Time** | Date and time when the delinquency state transition was last updated on Marqeta&#39;s credit platform, in UTC. | [optional] 

## Methods

### NewDelinquencyTransitionResponse

`func NewDelinquencyTransitionResponse(accountToken string, createdTime time.Time, impactTime time.Time, isRolledBack bool, originalStatus DelinquencyStatus, status DelinquencyStatus, token string, transitionTriggerReason DelinquencyTransitionTriggerReason, transitionTriggerTime time.Time, ) *DelinquencyTransitionResponse`

NewDelinquencyTransitionResponse instantiates a new DelinquencyTransitionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDelinquencyTransitionResponseWithDefaults

`func NewDelinquencyTransitionResponseWithDefaults() *DelinquencyTransitionResponse`

NewDelinquencyTransitionResponseWithDefaults instantiates a new DelinquencyTransitionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountToken

`func (o *DelinquencyTransitionResponse) GetAccountToken() string`

GetAccountToken returns the AccountToken field if non-nil, zero value otherwise.

### GetAccountTokenOk

`func (o *DelinquencyTransitionResponse) GetAccountTokenOk() (*string, bool)`

GetAccountTokenOk returns a tuple with the AccountToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountToken

`func (o *DelinquencyTransitionResponse) SetAccountToken(v string)`

SetAccountToken sets AccountToken field to given value.


### GetBucketCount

`func (o *DelinquencyTransitionResponse) GetBucketCount() float32`

GetBucketCount returns the BucketCount field if non-nil, zero value otherwise.

### GetBucketCountOk

`func (o *DelinquencyTransitionResponse) GetBucketCountOk() (*float32, bool)`

GetBucketCountOk returns a tuple with the BucketCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketCount

`func (o *DelinquencyTransitionResponse) SetBucketCount(v float32)`

SetBucketCount sets BucketCount field to given value.

### HasBucketCount

`func (o *DelinquencyTransitionResponse) HasBucketCount() bool`

HasBucketCount returns a boolean if a field has been set.

### GetCreatedTime

`func (o *DelinquencyTransitionResponse) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *DelinquencyTransitionResponse) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *DelinquencyTransitionResponse) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetCurrentDue

`func (o *DelinquencyTransitionResponse) GetCurrentDue() float32`

GetCurrentDue returns the CurrentDue field if non-nil, zero value otherwise.

### GetCurrentDueOk

`func (o *DelinquencyTransitionResponse) GetCurrentDueOk() (*float32, bool)`

GetCurrentDueOk returns a tuple with the CurrentDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentDue

`func (o *DelinquencyTransitionResponse) SetCurrentDue(v float32)`

SetCurrentDue sets CurrentDue field to given value.

### HasCurrentDue

`func (o *DelinquencyTransitionResponse) HasCurrentDue() bool`

HasCurrentDue returns a boolean if a field has been set.

### GetImpactTime

`func (o *DelinquencyTransitionResponse) GetImpactTime() time.Time`

GetImpactTime returns the ImpactTime field if non-nil, zero value otherwise.

### GetImpactTimeOk

`func (o *DelinquencyTransitionResponse) GetImpactTimeOk() (*time.Time, bool)`

GetImpactTimeOk returns a tuple with the ImpactTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpactTime

`func (o *DelinquencyTransitionResponse) SetImpactTime(v time.Time)`

SetImpactTime sets ImpactTime field to given value.


### GetIsRolledBack

`func (o *DelinquencyTransitionResponse) GetIsRolledBack() bool`

GetIsRolledBack returns the IsRolledBack field if non-nil, zero value otherwise.

### GetIsRolledBackOk

`func (o *DelinquencyTransitionResponse) GetIsRolledBackOk() (*bool, bool)`

GetIsRolledBackOk returns a tuple with the IsRolledBack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRolledBack

`func (o *DelinquencyTransitionResponse) SetIsRolledBack(v bool)`

SetIsRolledBack sets IsRolledBack field to given value.


### GetOldestPaymentDueDate

`func (o *DelinquencyTransitionResponse) GetOldestPaymentDueDate() time.Time`

GetOldestPaymentDueDate returns the OldestPaymentDueDate field if non-nil, zero value otherwise.

### GetOldestPaymentDueDateOk

`func (o *DelinquencyTransitionResponse) GetOldestPaymentDueDateOk() (*time.Time, bool)`

GetOldestPaymentDueDateOk returns a tuple with the OldestPaymentDueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldestPaymentDueDate

`func (o *DelinquencyTransitionResponse) SetOldestPaymentDueDate(v time.Time)`

SetOldestPaymentDueDate sets OldestPaymentDueDate field to given value.

### HasOldestPaymentDueDate

`func (o *DelinquencyTransitionResponse) HasOldestPaymentDueDate() bool`

HasOldestPaymentDueDate returns a boolean if a field has been set.

### GetOriginalStatus

`func (o *DelinquencyTransitionResponse) GetOriginalStatus() DelinquencyStatus`

GetOriginalStatus returns the OriginalStatus field if non-nil, zero value otherwise.

### GetOriginalStatusOk

`func (o *DelinquencyTransitionResponse) GetOriginalStatusOk() (*DelinquencyStatus, bool)`

GetOriginalStatusOk returns a tuple with the OriginalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalStatus

`func (o *DelinquencyTransitionResponse) SetOriginalStatus(v DelinquencyStatus)`

SetOriginalStatus sets OriginalStatus field to given value.


### GetStatus

`func (o *DelinquencyTransitionResponse) GetStatus() DelinquencyStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DelinquencyTransitionResponse) GetStatusOk() (*DelinquencyStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DelinquencyTransitionResponse) SetStatus(v DelinquencyStatus)`

SetStatus sets Status field to given value.


### GetToken

`func (o *DelinquencyTransitionResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DelinquencyTransitionResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DelinquencyTransitionResponse) SetToken(v string)`

SetToken sets Token field to given value.


### GetTotalDue

`func (o *DelinquencyTransitionResponse) GetTotalDue() float32`

GetTotalDue returns the TotalDue field if non-nil, zero value otherwise.

### GetTotalDueOk

`func (o *DelinquencyTransitionResponse) GetTotalDueOk() (*float32, bool)`

GetTotalDueOk returns a tuple with the TotalDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDue

`func (o *DelinquencyTransitionResponse) SetTotalDue(v float32)`

SetTotalDue sets TotalDue field to given value.

### HasTotalDue

`func (o *DelinquencyTransitionResponse) HasTotalDue() bool`

HasTotalDue returns a boolean if a field has been set.

### GetTotalPastDue

`func (o *DelinquencyTransitionResponse) GetTotalPastDue() float32`

GetTotalPastDue returns the TotalPastDue field if non-nil, zero value otherwise.

### GetTotalPastDueOk

`func (o *DelinquencyTransitionResponse) GetTotalPastDueOk() (*float32, bool)`

GetTotalPastDueOk returns a tuple with the TotalPastDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPastDue

`func (o *DelinquencyTransitionResponse) SetTotalPastDue(v float32)`

SetTotalPastDue sets TotalPastDue field to given value.

### HasTotalPastDue

`func (o *DelinquencyTransitionResponse) HasTotalPastDue() bool`

HasTotalPastDue returns a boolean if a field has been set.

### GetTransitionTriggerReason

`func (o *DelinquencyTransitionResponse) GetTransitionTriggerReason() DelinquencyTransitionTriggerReason`

GetTransitionTriggerReason returns the TransitionTriggerReason field if non-nil, zero value otherwise.

### GetTransitionTriggerReasonOk

`func (o *DelinquencyTransitionResponse) GetTransitionTriggerReasonOk() (*DelinquencyTransitionTriggerReason, bool)`

GetTransitionTriggerReasonOk returns a tuple with the TransitionTriggerReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitionTriggerReason

`func (o *DelinquencyTransitionResponse) SetTransitionTriggerReason(v DelinquencyTransitionTriggerReason)`

SetTransitionTriggerReason sets TransitionTriggerReason field to given value.


### GetTransitionTriggerTime

`func (o *DelinquencyTransitionResponse) GetTransitionTriggerTime() time.Time`

GetTransitionTriggerTime returns the TransitionTriggerTime field if non-nil, zero value otherwise.

### GetTransitionTriggerTimeOk

`func (o *DelinquencyTransitionResponse) GetTransitionTriggerTimeOk() (*time.Time, bool)`

GetTransitionTriggerTimeOk returns a tuple with the TransitionTriggerTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitionTriggerTime

`func (o *DelinquencyTransitionResponse) SetTransitionTriggerTime(v time.Time)`

SetTransitionTriggerTime sets TransitionTriggerTime field to given value.


### GetUpdatedTime

`func (o *DelinquencyTransitionResponse) GetUpdatedTime() time.Time`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *DelinquencyTransitionResponse) GetUpdatedTimeOk() (*time.Time, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *DelinquencyTransitionResponse) SetUpdatedTime(v time.Time)`

SetUpdatedTime sets UpdatedTime field to given value.

### HasUpdatedTime

`func (o *DelinquencyTransitionResponse) HasUpdatedTime() bool`

HasUpdatedTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


