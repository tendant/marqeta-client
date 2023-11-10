# EchoPingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Payload** | Pointer to **string** |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewEchoPingResponse

`func NewEchoPingResponse() *EchoPingResponse`

NewEchoPingResponse instantiates a new EchoPingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEchoPingResponseWithDefaults

`func NewEchoPingResponseWithDefaults() *EchoPingResponse`

NewEchoPingResponseWithDefaults instantiates a new EchoPingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EchoPingResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EchoPingResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EchoPingResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EchoPingResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPayload

`func (o *EchoPingResponse) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *EchoPingResponse) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *EchoPingResponse) SetPayload(v string)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *EchoPingResponse) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetSuccess

`func (o *EchoPingResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *EchoPingResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *EchoPingResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *EchoPingResponse) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


