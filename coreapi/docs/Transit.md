# Transit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionType** | Pointer to **string** | Type of transit transaction. | [optional] 
**TransportationMode** | Pointer to **string** | Mode of transportation. | [optional] 

## Methods

### NewTransit

`func NewTransit() *Transit`

NewTransit instantiates a new Transit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransitWithDefaults

`func NewTransitWithDefaults() *Transit`

NewTransitWithDefaults instantiates a new Transit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionType

`func (o *Transit) GetTransactionType() string`

GetTransactionType returns the TransactionType field if non-nil, zero value otherwise.

### GetTransactionTypeOk

`func (o *Transit) GetTransactionTypeOk() (*string, bool)`

GetTransactionTypeOk returns a tuple with the TransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionType

`func (o *Transit) SetTransactionType(v string)`

SetTransactionType sets TransactionType field to given value.

### HasTransactionType

`func (o *Transit) HasTransactionType() bool`

HasTransactionType returns a boolean if a field has been set.

### GetTransportationMode

`func (o *Transit) GetTransportationMode() string`

GetTransportationMode returns the TransportationMode field if non-nil, zero value otherwise.

### GetTransportationModeOk

`func (o *Transit) GetTransportationModeOk() (*string, bool)`

GetTransportationModeOk returns a tuple with the TransportationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportationMode

`func (o *Transit) SetTransportationMode(v string)`

SetTransportationMode sets TransportationMode field to given value.

### HasTransportationMode

`func (o *Transit) HasTransportationMode() bool`

HasTransportationMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


