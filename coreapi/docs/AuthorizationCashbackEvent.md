# AuthorizationCashbackEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardToken** | **string** | Unique identifier of the card. Useful when a single account holder has multiple cards. | 
**Amount** | **float32** | Amount of the transaction. | 
**CardAcceptor** | Pointer to [**TransactionCardAcceptor**](transaction_card_acceptor.md) |  | [optional] 

## Methods

### NewAuthorizationCashbackEvent

`func NewAuthorizationCashbackEvent(cardToken string, amount float32, ) *AuthorizationCashbackEvent`

NewAuthorizationCashbackEvent instantiates a new AuthorizationCashbackEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizationCashbackEventWithDefaults

`func NewAuthorizationCashbackEventWithDefaults() *AuthorizationCashbackEvent`

NewAuthorizationCashbackEventWithDefaults instantiates a new AuthorizationCashbackEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardToken

`func (o *AuthorizationCashbackEvent) GetCardToken() string`

GetCardToken returns the CardToken field if non-nil, zero value otherwise.

### GetCardTokenOk

`func (o *AuthorizationCashbackEvent) GetCardTokenOk() (*string, bool)`

GetCardTokenOk returns a tuple with the CardToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardToken

`func (o *AuthorizationCashbackEvent) SetCardToken(v string)`

SetCardToken sets CardToken field to given value.


### GetAmount

`func (o *AuthorizationCashbackEvent) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *AuthorizationCashbackEvent) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *AuthorizationCashbackEvent) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetCardAcceptor

`func (o *AuthorizationCashbackEvent) GetCardAcceptor() TransactionCardAcceptor`

GetCardAcceptor returns the CardAcceptor field if non-nil, zero value otherwise.

### GetCardAcceptorOk

`func (o *AuthorizationCashbackEvent) GetCardAcceptorOk() (*TransactionCardAcceptor, bool)`

GetCardAcceptorOk returns a tuple with the CardAcceptor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardAcceptor

`func (o *AuthorizationCashbackEvent) SetCardAcceptor(v TransactionCardAcceptor)`

SetCardAcceptor sets CardAcceptor field to given value.

### HasCardAcceptor

`func (o *AuthorizationCashbackEvent) HasCardAcceptor() bool`

HasCardAcceptor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


