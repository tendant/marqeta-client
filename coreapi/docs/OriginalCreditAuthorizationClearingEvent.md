# OriginalCreditAuthorizationClearingEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **float32** | Amount of the transaction. | 
**CardAcceptor** | Pointer to [**TransactionCardAcceptor**](TransactionCardAcceptor.md) |  | [optional] 

## Methods

### NewOriginalCreditAuthorizationClearingEvent

`func NewOriginalCreditAuthorizationClearingEvent(amount float32, ) *OriginalCreditAuthorizationClearingEvent`

NewOriginalCreditAuthorizationClearingEvent instantiates a new OriginalCreditAuthorizationClearingEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOriginalCreditAuthorizationClearingEventWithDefaults

`func NewOriginalCreditAuthorizationClearingEventWithDefaults() *OriginalCreditAuthorizationClearingEvent`

NewOriginalCreditAuthorizationClearingEventWithDefaults instantiates a new OriginalCreditAuthorizationClearingEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *OriginalCreditAuthorizationClearingEvent) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *OriginalCreditAuthorizationClearingEvent) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *OriginalCreditAuthorizationClearingEvent) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetCardAcceptor

`func (o *OriginalCreditAuthorizationClearingEvent) GetCardAcceptor() TransactionCardAcceptor`

GetCardAcceptor returns the CardAcceptor field if non-nil, zero value otherwise.

### GetCardAcceptorOk

`func (o *OriginalCreditAuthorizationClearingEvent) GetCardAcceptorOk() (*TransactionCardAcceptor, bool)`

GetCardAcceptorOk returns a tuple with the CardAcceptor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardAcceptor

`func (o *OriginalCreditAuthorizationClearingEvent) SetCardAcceptor(v TransactionCardAcceptor)`

SetCardAcceptor sets CardAcceptor field to given value.

### HasCardAcceptor

`func (o *OriginalCreditAuthorizationClearingEvent) HasCardAcceptor() bool`

HasCardAcceptor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


