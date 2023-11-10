# DisputeCreateReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **float32** | Amount of the dispute. Max value is equal to the value of the original transaction. | 
**Category** | [**DisputeCategory**](DisputeCategory.md) |  | 
**LedgerEntryToken** | **string** | Unique identifier of the journal entry (&#x60;authorization.clearing&#x60; type only) in dispute. | 
**Notes** | Pointer to **string** | Additional information on the dispute (for example, a reason for the dispute). | [optional] 
**Token** | Pointer to **string** | Unique identifier of the dispute. | [optional] 

## Methods

### NewDisputeCreateReq

`func NewDisputeCreateReq(amount float32, category DisputeCategory, ledgerEntryToken string, ) *DisputeCreateReq`

NewDisputeCreateReq instantiates a new DisputeCreateReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDisputeCreateReqWithDefaults

`func NewDisputeCreateReqWithDefaults() *DisputeCreateReq`

NewDisputeCreateReqWithDefaults instantiates a new DisputeCreateReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *DisputeCreateReq) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *DisputeCreateReq) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *DisputeCreateReq) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetCategory

`func (o *DisputeCreateReq) GetCategory() DisputeCategory`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *DisputeCreateReq) GetCategoryOk() (*DisputeCategory, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *DisputeCreateReq) SetCategory(v DisputeCategory)`

SetCategory sets Category field to given value.


### GetLedgerEntryToken

`func (o *DisputeCreateReq) GetLedgerEntryToken() string`

GetLedgerEntryToken returns the LedgerEntryToken field if non-nil, zero value otherwise.

### GetLedgerEntryTokenOk

`func (o *DisputeCreateReq) GetLedgerEntryTokenOk() (*string, bool)`

GetLedgerEntryTokenOk returns a tuple with the LedgerEntryToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedgerEntryToken

`func (o *DisputeCreateReq) SetLedgerEntryToken(v string)`

SetLedgerEntryToken sets LedgerEntryToken field to given value.


### GetNotes

`func (o *DisputeCreateReq) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *DisputeCreateReq) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *DisputeCreateReq) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *DisputeCreateReq) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetToken

`func (o *DisputeCreateReq) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DisputeCreateReq) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DisputeCreateReq) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DisputeCreateReq) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


