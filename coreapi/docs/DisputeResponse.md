# DisputeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountToken** | **string** | Unique identifier of the credit account on which the dispute was created. | 
**Amount** | **float32** | Amount of the dispute. | 
**Category** | [**DisputeCategory**](DisputeCategory.md) |  | 
**CreatedTime** | **time.Time** | Date and time when the dispute was created on Marqeta&#39;s credit platform, in UTC. | 
**LedgerEntryToken** | **string** | Unique identifier of the journal entry (&#x60;authorization.clearing&#x60; type only) in dispute. | 
**Notes** | Pointer to **string** | Additional information on the dispute (for example, a reason for the dispute). | [optional] 
**ResolvedAt** | Pointer to **time.Time** | Date and time when the dispute was resolved and no longer in &#x60;ACTIVE&#x60; status. | [optional] 
**Status** | [**DisputeStatus**](DisputeStatus.md) |  | 
**Token** | **string** | Unique identifier of the dispute. | 
**UpdatedTime** | **time.Time** | Date and time when the dispute was last updated on Marqeta&#39;s credit platform, in UTC. | 

## Methods

### NewDisputeResponse

`func NewDisputeResponse(accountToken string, amount float32, category DisputeCategory, createdTime time.Time, ledgerEntryToken string, status DisputeStatus, token string, updatedTime time.Time, ) *DisputeResponse`

NewDisputeResponse instantiates a new DisputeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDisputeResponseWithDefaults

`func NewDisputeResponseWithDefaults() *DisputeResponse`

NewDisputeResponseWithDefaults instantiates a new DisputeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountToken

`func (o *DisputeResponse) GetAccountToken() string`

GetAccountToken returns the AccountToken field if non-nil, zero value otherwise.

### GetAccountTokenOk

`func (o *DisputeResponse) GetAccountTokenOk() (*string, bool)`

GetAccountTokenOk returns a tuple with the AccountToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountToken

`func (o *DisputeResponse) SetAccountToken(v string)`

SetAccountToken sets AccountToken field to given value.


### GetAmount

`func (o *DisputeResponse) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *DisputeResponse) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *DisputeResponse) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetCategory

`func (o *DisputeResponse) GetCategory() DisputeCategory`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *DisputeResponse) GetCategoryOk() (*DisputeCategory, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *DisputeResponse) SetCategory(v DisputeCategory)`

SetCategory sets Category field to given value.


### GetCreatedTime

`func (o *DisputeResponse) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *DisputeResponse) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *DisputeResponse) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetLedgerEntryToken

`func (o *DisputeResponse) GetLedgerEntryToken() string`

GetLedgerEntryToken returns the LedgerEntryToken field if non-nil, zero value otherwise.

### GetLedgerEntryTokenOk

`func (o *DisputeResponse) GetLedgerEntryTokenOk() (*string, bool)`

GetLedgerEntryTokenOk returns a tuple with the LedgerEntryToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedgerEntryToken

`func (o *DisputeResponse) SetLedgerEntryToken(v string)`

SetLedgerEntryToken sets LedgerEntryToken field to given value.


### GetNotes

`func (o *DisputeResponse) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *DisputeResponse) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *DisputeResponse) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *DisputeResponse) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetResolvedAt

`func (o *DisputeResponse) GetResolvedAt() time.Time`

GetResolvedAt returns the ResolvedAt field if non-nil, zero value otherwise.

### GetResolvedAtOk

`func (o *DisputeResponse) GetResolvedAtOk() (*time.Time, bool)`

GetResolvedAtOk returns a tuple with the ResolvedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedAt

`func (o *DisputeResponse) SetResolvedAt(v time.Time)`

SetResolvedAt sets ResolvedAt field to given value.

### HasResolvedAt

`func (o *DisputeResponse) HasResolvedAt() bool`

HasResolvedAt returns a boolean if a field has been set.

### GetStatus

`func (o *DisputeResponse) GetStatus() DisputeStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DisputeResponse) GetStatusOk() (*DisputeStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DisputeResponse) SetStatus(v DisputeStatus)`

SetStatus sets Status field to given value.


### GetToken

`func (o *DisputeResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DisputeResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DisputeResponse) SetToken(v string)`

SetToken sets Token field to given value.


### GetUpdatedTime

`func (o *DisputeResponse) GetUpdatedTime() time.Time`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *DisputeResponse) GetUpdatedTimeOk() (*time.Time, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *DisputeResponse) SetUpdatedTime(v time.Time)`

SetUpdatedTime sets UpdatedTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


