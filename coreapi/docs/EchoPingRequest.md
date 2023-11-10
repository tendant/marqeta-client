# EchoPingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Payload** | Pointer to **string** | Payload of the ping request. | [optional] 
**Token** | Pointer to **string** | Unique identifier of the ping request. | [optional] 

## Methods

### NewEchoPingRequest

`func NewEchoPingRequest() *EchoPingRequest`

NewEchoPingRequest instantiates a new EchoPingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEchoPingRequestWithDefaults

`func NewEchoPingRequestWithDefaults() *EchoPingRequest`

NewEchoPingRequestWithDefaults instantiates a new EchoPingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayload

`func (o *EchoPingRequest) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *EchoPingRequest) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *EchoPingRequest) SetPayload(v string)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *EchoPingRequest) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetToken

`func (o *EchoPingRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *EchoPingRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *EchoPingRequest) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *EchoPingRequest) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


