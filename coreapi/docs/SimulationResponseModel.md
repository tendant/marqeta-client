# SimulationResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RawIso8583** | Pointer to **map[string]map[string]interface{}** |  | [optional] [readonly] 
**Transaction** | Pointer to [**TransactionModel**](TransactionModel.md) |  | [optional] 

## Methods

### NewSimulationResponseModel

`func NewSimulationResponseModel() *SimulationResponseModel`

NewSimulationResponseModel instantiates a new SimulationResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimulationResponseModelWithDefaults

`func NewSimulationResponseModelWithDefaults() *SimulationResponseModel`

NewSimulationResponseModelWithDefaults instantiates a new SimulationResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRawIso8583

`func (o *SimulationResponseModel) GetRawIso8583() map[string]map[string]interface{}`

GetRawIso8583 returns the RawIso8583 field if non-nil, zero value otherwise.

### GetRawIso8583Ok

`func (o *SimulationResponseModel) GetRawIso8583Ok() (*map[string]map[string]interface{}, bool)`

GetRawIso8583Ok returns a tuple with the RawIso8583 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawIso8583

`func (o *SimulationResponseModel) SetRawIso8583(v map[string]map[string]interface{})`

SetRawIso8583 sets RawIso8583 field to given value.

### HasRawIso8583

`func (o *SimulationResponseModel) HasRawIso8583() bool`

HasRawIso8583 returns a boolean if a field has been set.

### GetTransaction

`func (o *SimulationResponseModel) GetTransaction() TransactionModel`

GetTransaction returns the Transaction field if non-nil, zero value otherwise.

### GetTransactionOk

`func (o *SimulationResponseModel) GetTransactionOk() (*TransactionModel, bool)`

GetTransactionOk returns a tuple with the Transaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransaction

`func (o *SimulationResponseModel) SetTransaction(v TransactionModel)`

SetTransaction sets Transaction field to given value.

### HasTransaction

`func (o *SimulationResponseModel) HasTransaction() bool`

HasTransaction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


