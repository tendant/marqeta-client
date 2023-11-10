# PindebitQuasiCashEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardToken** | **string** | Unique identifier of the card. Useful when a single account holder has multiple cards. | 
**Amount** | **float32** | Amount of the transaction. | 
**CardAcceptor** | Pointer to [**TransactionCardAcceptor**](transaction_card_acceptor.md) |  | [optional] 

## Methods

### NewPindebitQuasiCashEvent

`func NewPindebitQuasiCashEvent(cardToken string, amount float32, ) *PindebitQuasiCashEvent`

NewPindebitQuasiCashEvent instantiates a new PindebitQuasiCashEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPindebitQuasiCashEventWithDefaults

`func NewPindebitQuasiCashEventWithDefaults() *PindebitQuasiCashEvent`

NewPindebitQuasiCashEventWithDefaults instantiates a new PindebitQuasiCashEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardToken

`func (o *PindebitQuasiCashEvent) GetCardToken() string`

GetCardToken returns the CardToken field if non-nil, zero value otherwise.

### GetCardTokenOk

`func (o *PindebitQuasiCashEvent) GetCardTokenOk() (*string, bool)`

GetCardTokenOk returns a tuple with the CardToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardToken

`func (o *PindebitQuasiCashEvent) SetCardToken(v string)`

SetCardToken sets CardToken field to given value.


### GetAmount

`func (o *PindebitQuasiCashEvent) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PindebitQuasiCashEvent) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PindebitQuasiCashEvent) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetCardAcceptor

`func (o *PindebitQuasiCashEvent) GetCardAcceptor() TransactionCardAcceptor`

GetCardAcceptor returns the CardAcceptor field if non-nil, zero value otherwise.

### GetCardAcceptorOk

`func (o *PindebitQuasiCashEvent) GetCardAcceptorOk() (*TransactionCardAcceptor, bool)`

GetCardAcceptorOk returns a tuple with the CardAcceptor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardAcceptor

`func (o *PindebitQuasiCashEvent) SetCardAcceptor(v TransactionCardAcceptor)`

SetCardAcceptor sets CardAcceptor field to given value.

### HasCardAcceptor

`func (o *PindebitQuasiCashEvent) HasCardAcceptor() bool`

HasCardAcceptor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


