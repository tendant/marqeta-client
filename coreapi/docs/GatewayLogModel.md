# GatewayLogModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | Pointer to **int64** | Length of time in milliseconds that the gateway took to respond to a funding request. | [optional] 
**Message** | **string** | Message about the status of the funding request. Useful for determining whether it was approved and completed successfully, declined by the gateway, or timed out. | 
**OrderNumber** | **string** | Customer order number, same value as &#x60;transaction.token&#x60;. | 
**Response** | Pointer to [**GatewayResponse**](GatewayResponse.md) |  | [optional] 
**TimedOut** | Pointer to **bool** | Whether the gateway sent a response (&#x60;true&#x60;) or timed out (&#x60;false&#x60;). | [optional] [default to false]
**TransactionId** | **string** | Customer-defined identifier for the transaction. | 

## Methods

### NewGatewayLogModel

`func NewGatewayLogModel(message string, orderNumber string, transactionId string, ) *GatewayLogModel`

NewGatewayLogModel instantiates a new GatewayLogModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayLogModelWithDefaults

`func NewGatewayLogModelWithDefaults() *GatewayLogModel`

NewGatewayLogModelWithDefaults instantiates a new GatewayLogModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *GatewayLogModel) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *GatewayLogModel) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *GatewayLogModel) SetDuration(v int64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *GatewayLogModel) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetMessage

`func (o *GatewayLogModel) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GatewayLogModel) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GatewayLogModel) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetOrderNumber

`func (o *GatewayLogModel) GetOrderNumber() string`

GetOrderNumber returns the OrderNumber field if non-nil, zero value otherwise.

### GetOrderNumberOk

`func (o *GatewayLogModel) GetOrderNumberOk() (*string, bool)`

GetOrderNumberOk returns a tuple with the OrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderNumber

`func (o *GatewayLogModel) SetOrderNumber(v string)`

SetOrderNumber sets OrderNumber field to given value.


### GetResponse

`func (o *GatewayLogModel) GetResponse() GatewayResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *GatewayLogModel) GetResponseOk() (*GatewayResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *GatewayLogModel) SetResponse(v GatewayResponse)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *GatewayLogModel) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetTimedOut

`func (o *GatewayLogModel) GetTimedOut() bool`

GetTimedOut returns the TimedOut field if non-nil, zero value otherwise.

### GetTimedOutOk

`func (o *GatewayLogModel) GetTimedOutOk() (*bool, bool)`

GetTimedOutOk returns a tuple with the TimedOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimedOut

`func (o *GatewayLogModel) SetTimedOut(v bool)`

SetTimedOut sets TimedOut field to given value.

### HasTimedOut

`func (o *GatewayLogModel) HasTimedOut() bool`

HasTimedOut returns a boolean if a field has been set.

### GetTransactionId

`func (o *GatewayLogModel) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *GatewayLogModel) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *GatewayLogModel) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


