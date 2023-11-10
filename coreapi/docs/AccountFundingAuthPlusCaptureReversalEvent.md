# AccountFundingAuthPlusCaptureReversalEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrecedingRelatedTransactionToken** | **string** | Unique identifier of the card. Useful when a single account holder has multiple cards. | 
**Amount** | **float32** | Amount of the transaction. | 
**CardAcceptor** | Pointer to [**TransactionCardAcceptor**](TransactionCardAcceptor.md) |  | [optional] 

## Methods

### NewAccountFundingAuthPlusCaptureReversalEvent

`func NewAccountFundingAuthPlusCaptureReversalEvent(precedingRelatedTransactionToken string, amount float32, ) *AccountFundingAuthPlusCaptureReversalEvent`

NewAccountFundingAuthPlusCaptureReversalEvent instantiates a new AccountFundingAuthPlusCaptureReversalEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountFundingAuthPlusCaptureReversalEventWithDefaults

`func NewAccountFundingAuthPlusCaptureReversalEventWithDefaults() *AccountFundingAuthPlusCaptureReversalEvent`

NewAccountFundingAuthPlusCaptureReversalEventWithDefaults instantiates a new AccountFundingAuthPlusCaptureReversalEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrecedingRelatedTransactionToken

`func (o *AccountFundingAuthPlusCaptureReversalEvent) GetPrecedingRelatedTransactionToken() string`

GetPrecedingRelatedTransactionToken returns the PrecedingRelatedTransactionToken field if non-nil, zero value otherwise.

### GetPrecedingRelatedTransactionTokenOk

`func (o *AccountFundingAuthPlusCaptureReversalEvent) GetPrecedingRelatedTransactionTokenOk() (*string, bool)`

GetPrecedingRelatedTransactionTokenOk returns a tuple with the PrecedingRelatedTransactionToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecedingRelatedTransactionToken

`func (o *AccountFundingAuthPlusCaptureReversalEvent) SetPrecedingRelatedTransactionToken(v string)`

SetPrecedingRelatedTransactionToken sets PrecedingRelatedTransactionToken field to given value.


### GetAmount

`func (o *AccountFundingAuthPlusCaptureReversalEvent) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *AccountFundingAuthPlusCaptureReversalEvent) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *AccountFundingAuthPlusCaptureReversalEvent) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetCardAcceptor

`func (o *AccountFundingAuthPlusCaptureReversalEvent) GetCardAcceptor() TransactionCardAcceptor`

GetCardAcceptor returns the CardAcceptor field if non-nil, zero value otherwise.

### GetCardAcceptorOk

`func (o *AccountFundingAuthPlusCaptureReversalEvent) GetCardAcceptorOk() (*TransactionCardAcceptor, bool)`

GetCardAcceptorOk returns a tuple with the CardAcceptor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardAcceptor

`func (o *AccountFundingAuthPlusCaptureReversalEvent) SetCardAcceptor(v TransactionCardAcceptor)`

SetCardAcceptor sets CardAcceptor field to given value.

### HasCardAcceptor

`func (o *AccountFundingAuthPlusCaptureReversalEvent) HasCardAcceptor() bool`

HasCardAcceptor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


