# AccountAdjustmentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountToken** | **string** | Unique identifier of the credit account on which the adjustment was made. | 
**AdjustmentDetailObject** | Pointer to **map[string]interface{}** | Contains the adjustment&#39;s full details.  The fields returned in this object depend on the adjustment type.  Interest returns interest details. For the specific fields returned, see the &#x60;detail_object&#x60; fields marked \&quot;Returned for interest journal entries\&quot; in the &lt;&lt;/core-api/credit-account-journal-entries#getAccountJournalEntry, account journal entry response fields&gt;&gt;.  Disputes return dispute details. For the specific fields returned, see the &lt;&lt;/core-api/credit-disputes#retrieveDispute, dispute response fields&gt;&gt;. | [optional] 
**Amount** | **float32** | Amount of the adjustment. | 
**CreatedTime** | Pointer to **time.Time** | Date and time when the account adjustment was applied, in UTC. | [optional] 
**CurrencyCode** | [**CurrencyCode**](CurrencyCode.md) |  | [default to CURRENCYCODE_USD]
**Description** | **string** | Description of the adjustment. | 
**DetailToken** | Pointer to **NullableString** | Unique identifier of the adjustment detail. For example, the token of the dispute, the interest charge, or the returned payment that prompted the adjustment.  Returned when the system automatically applies an adjustment. | [optional] 
**ExternalAdjustmentId** | Pointer to **NullableString** | Unique identifier you provide of an associated external adjustment that exists outside Marqeta&#39;s credit platform. | [optional] 
**Note** | Pointer to **NullableString** | Additional information on the adjustment. | [optional] 
**OriginalLedgerEntryToken** | Pointer to **string** | Unique identifier of the original journal entry needing the adjustment. | [optional] 
**Reason** | **string** | Reason for the adjustment.  * &#x60;DISPUTE&#x60; - The adjustment occurred because a dispute was initiated. * &#x60;DISPUTE_RESOLUTION&#x60; - The adjustment occurred because of the result of a dispute resolution. * &#x60;RETURNED_OR_CANCELED_PAYMENT&#x60; - The adjustment occurred because a payment was returned or canceled. * &#x60;OTHER&#x60; - Any other reason the adjustment occurred. For example, a waived fee. | 
**RelatedDetailObject** | Pointer to **map[string]interface{}** | Contains full details of the related dispute or returned payment.  The fields returned in this object depend on whether a dispute or returned payment led to the interest adjustment. A dispute returns dispute details; a returned payment returns payment details.  For more on the dispute details returned, see the &lt;&lt;/core-api/credit-disputes#retrieveDispute, dispute response fields&gt;&gt;.  For more on the returned payment details returned, see the &lt;&lt;/core-api/credit-account-payments#retrievePayment, payment response fields&gt;&gt;.  This field is returned for interest adjustments only. | [optional] 
**RelatedDetailToken** | Pointer to **NullableString** | Unique identifier of the dispute or returned payment that prompted the interest adjustment.  This field is returned for interest adjustments only. | [optional] 
**Token** | **string** | Unique identifier of the adjustment.  If in the &#x60;detail_object&#x60;, unique identifier of the detail object. | 
**Type** | **string** | Type of adjustment.  The adjustment is made on its correlating amount (for example, purchase adjustments are made on purchase amounts). You can use general adjustments for standalone adjustments made on the credit account balance itself, which includes account write-offs, credits, and more. | 

## Methods

### NewAccountAdjustmentResponse

`func NewAccountAdjustmentResponse(accountToken string, amount float32, currencyCode CurrencyCode, description string, reason string, token string, type_ string, ) *AccountAdjustmentResponse`

NewAccountAdjustmentResponse instantiates a new AccountAdjustmentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountAdjustmentResponseWithDefaults

`func NewAccountAdjustmentResponseWithDefaults() *AccountAdjustmentResponse`

NewAccountAdjustmentResponseWithDefaults instantiates a new AccountAdjustmentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountToken

`func (o *AccountAdjustmentResponse) GetAccountToken() string`

GetAccountToken returns the AccountToken field if non-nil, zero value otherwise.

### GetAccountTokenOk

`func (o *AccountAdjustmentResponse) GetAccountTokenOk() (*string, bool)`

GetAccountTokenOk returns a tuple with the AccountToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountToken

`func (o *AccountAdjustmentResponse) SetAccountToken(v string)`

SetAccountToken sets AccountToken field to given value.


### GetAdjustmentDetailObject

`func (o *AccountAdjustmentResponse) GetAdjustmentDetailObject() map[string]interface{}`

GetAdjustmentDetailObject returns the AdjustmentDetailObject field if non-nil, zero value otherwise.

### GetAdjustmentDetailObjectOk

`func (o *AccountAdjustmentResponse) GetAdjustmentDetailObjectOk() (*map[string]interface{}, bool)`

GetAdjustmentDetailObjectOk returns a tuple with the AdjustmentDetailObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustmentDetailObject

`func (o *AccountAdjustmentResponse) SetAdjustmentDetailObject(v map[string]interface{})`

SetAdjustmentDetailObject sets AdjustmentDetailObject field to given value.

### HasAdjustmentDetailObject

`func (o *AccountAdjustmentResponse) HasAdjustmentDetailObject() bool`

HasAdjustmentDetailObject returns a boolean if a field has been set.

### SetAdjustmentDetailObjectNil

`func (o *AccountAdjustmentResponse) SetAdjustmentDetailObjectNil(b bool)`

 SetAdjustmentDetailObjectNil sets the value for AdjustmentDetailObject to be an explicit nil

### UnsetAdjustmentDetailObject
`func (o *AccountAdjustmentResponse) UnsetAdjustmentDetailObject()`

UnsetAdjustmentDetailObject ensures that no value is present for AdjustmentDetailObject, not even an explicit nil
### GetAmount

`func (o *AccountAdjustmentResponse) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *AccountAdjustmentResponse) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *AccountAdjustmentResponse) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetCreatedTime

`func (o *AccountAdjustmentResponse) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *AccountAdjustmentResponse) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *AccountAdjustmentResponse) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *AccountAdjustmentResponse) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *AccountAdjustmentResponse) GetCurrencyCode() CurrencyCode`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *AccountAdjustmentResponse) GetCurrencyCodeOk() (*CurrencyCode, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *AccountAdjustmentResponse) SetCurrencyCode(v CurrencyCode)`

SetCurrencyCode sets CurrencyCode field to given value.


### GetDescription

`func (o *AccountAdjustmentResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AccountAdjustmentResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AccountAdjustmentResponse) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDetailToken

`func (o *AccountAdjustmentResponse) GetDetailToken() string`

GetDetailToken returns the DetailToken field if non-nil, zero value otherwise.

### GetDetailTokenOk

`func (o *AccountAdjustmentResponse) GetDetailTokenOk() (*string, bool)`

GetDetailTokenOk returns a tuple with the DetailToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailToken

`func (o *AccountAdjustmentResponse) SetDetailToken(v string)`

SetDetailToken sets DetailToken field to given value.

### HasDetailToken

`func (o *AccountAdjustmentResponse) HasDetailToken() bool`

HasDetailToken returns a boolean if a field has been set.

### SetDetailTokenNil

`func (o *AccountAdjustmentResponse) SetDetailTokenNil(b bool)`

 SetDetailTokenNil sets the value for DetailToken to be an explicit nil

### UnsetDetailToken
`func (o *AccountAdjustmentResponse) UnsetDetailToken()`

UnsetDetailToken ensures that no value is present for DetailToken, not even an explicit nil
### GetExternalAdjustmentId

`func (o *AccountAdjustmentResponse) GetExternalAdjustmentId() string`

GetExternalAdjustmentId returns the ExternalAdjustmentId field if non-nil, zero value otherwise.

### GetExternalAdjustmentIdOk

`func (o *AccountAdjustmentResponse) GetExternalAdjustmentIdOk() (*string, bool)`

GetExternalAdjustmentIdOk returns a tuple with the ExternalAdjustmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalAdjustmentId

`func (o *AccountAdjustmentResponse) SetExternalAdjustmentId(v string)`

SetExternalAdjustmentId sets ExternalAdjustmentId field to given value.

### HasExternalAdjustmentId

`func (o *AccountAdjustmentResponse) HasExternalAdjustmentId() bool`

HasExternalAdjustmentId returns a boolean if a field has been set.

### SetExternalAdjustmentIdNil

`func (o *AccountAdjustmentResponse) SetExternalAdjustmentIdNil(b bool)`

 SetExternalAdjustmentIdNil sets the value for ExternalAdjustmentId to be an explicit nil

### UnsetExternalAdjustmentId
`func (o *AccountAdjustmentResponse) UnsetExternalAdjustmentId()`

UnsetExternalAdjustmentId ensures that no value is present for ExternalAdjustmentId, not even an explicit nil
### GetNote

`func (o *AccountAdjustmentResponse) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *AccountAdjustmentResponse) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *AccountAdjustmentResponse) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *AccountAdjustmentResponse) HasNote() bool`

HasNote returns a boolean if a field has been set.

### SetNoteNil

`func (o *AccountAdjustmentResponse) SetNoteNil(b bool)`

 SetNoteNil sets the value for Note to be an explicit nil

### UnsetNote
`func (o *AccountAdjustmentResponse) UnsetNote()`

UnsetNote ensures that no value is present for Note, not even an explicit nil
### GetOriginalLedgerEntryToken

`func (o *AccountAdjustmentResponse) GetOriginalLedgerEntryToken() string`

GetOriginalLedgerEntryToken returns the OriginalLedgerEntryToken field if non-nil, zero value otherwise.

### GetOriginalLedgerEntryTokenOk

`func (o *AccountAdjustmentResponse) GetOriginalLedgerEntryTokenOk() (*string, bool)`

GetOriginalLedgerEntryTokenOk returns a tuple with the OriginalLedgerEntryToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalLedgerEntryToken

`func (o *AccountAdjustmentResponse) SetOriginalLedgerEntryToken(v string)`

SetOriginalLedgerEntryToken sets OriginalLedgerEntryToken field to given value.

### HasOriginalLedgerEntryToken

`func (o *AccountAdjustmentResponse) HasOriginalLedgerEntryToken() bool`

HasOriginalLedgerEntryToken returns a boolean if a field has been set.

### GetReason

`func (o *AccountAdjustmentResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *AccountAdjustmentResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *AccountAdjustmentResponse) SetReason(v string)`

SetReason sets Reason field to given value.


### GetRelatedDetailObject

`func (o *AccountAdjustmentResponse) GetRelatedDetailObject() map[string]interface{}`

GetRelatedDetailObject returns the RelatedDetailObject field if non-nil, zero value otherwise.

### GetRelatedDetailObjectOk

`func (o *AccountAdjustmentResponse) GetRelatedDetailObjectOk() (*map[string]interface{}, bool)`

GetRelatedDetailObjectOk returns a tuple with the RelatedDetailObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedDetailObject

`func (o *AccountAdjustmentResponse) SetRelatedDetailObject(v map[string]interface{})`

SetRelatedDetailObject sets RelatedDetailObject field to given value.

### HasRelatedDetailObject

`func (o *AccountAdjustmentResponse) HasRelatedDetailObject() bool`

HasRelatedDetailObject returns a boolean if a field has been set.

### SetRelatedDetailObjectNil

`func (o *AccountAdjustmentResponse) SetRelatedDetailObjectNil(b bool)`

 SetRelatedDetailObjectNil sets the value for RelatedDetailObject to be an explicit nil

### UnsetRelatedDetailObject
`func (o *AccountAdjustmentResponse) UnsetRelatedDetailObject()`

UnsetRelatedDetailObject ensures that no value is present for RelatedDetailObject, not even an explicit nil
### GetRelatedDetailToken

`func (o *AccountAdjustmentResponse) GetRelatedDetailToken() string`

GetRelatedDetailToken returns the RelatedDetailToken field if non-nil, zero value otherwise.

### GetRelatedDetailTokenOk

`func (o *AccountAdjustmentResponse) GetRelatedDetailTokenOk() (*string, bool)`

GetRelatedDetailTokenOk returns a tuple with the RelatedDetailToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedDetailToken

`func (o *AccountAdjustmentResponse) SetRelatedDetailToken(v string)`

SetRelatedDetailToken sets RelatedDetailToken field to given value.

### HasRelatedDetailToken

`func (o *AccountAdjustmentResponse) HasRelatedDetailToken() bool`

HasRelatedDetailToken returns a boolean if a field has been set.

### SetRelatedDetailTokenNil

`func (o *AccountAdjustmentResponse) SetRelatedDetailTokenNil(b bool)`

 SetRelatedDetailTokenNil sets the value for RelatedDetailToken to be an explicit nil

### UnsetRelatedDetailToken
`func (o *AccountAdjustmentResponse) UnsetRelatedDetailToken()`

UnsetRelatedDetailToken ensures that no value is present for RelatedDetailToken, not even an explicit nil
### GetToken

`func (o *AccountAdjustmentResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AccountAdjustmentResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AccountAdjustmentResponse) SetToken(v string)`

SetToken sets Token field to given value.


### GetType

`func (o *AccountAdjustmentResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AccountAdjustmentResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AccountAdjustmentResponse) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


