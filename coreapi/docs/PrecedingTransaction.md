# PrecedingTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **float32** | Amount of the preceding transaction. | [optional] 
**Token** | Pointer to **string** | Unique identifier of the preceding transaction. | [optional] 

## Methods

### NewPrecedingTransaction

`func NewPrecedingTransaction() *PrecedingTransaction`

NewPrecedingTransaction instantiates a new PrecedingTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrecedingTransactionWithDefaults

`func NewPrecedingTransactionWithDefaults() *PrecedingTransaction`

NewPrecedingTransactionWithDefaults instantiates a new PrecedingTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *PrecedingTransaction) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PrecedingTransaction) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PrecedingTransaction) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *PrecedingTransaction) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetToken

`func (o *PrecedingTransaction) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PrecedingTransaction) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PrecedingTransaction) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *PrecedingTransaction) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


