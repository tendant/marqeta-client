# GatewayResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | Code received from the gateway. | 
**Data** | Pointer to [**JitProgramResponse**](JitProgramResponse.md) |  | [optional] 

## Methods

### NewGatewayResponse

`func NewGatewayResponse(code string, ) *GatewayResponse`

NewGatewayResponse instantiates a new GatewayResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayResponseWithDefaults

`func NewGatewayResponseWithDefaults() *GatewayResponse`

NewGatewayResponseWithDefaults instantiates a new GatewayResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *GatewayResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GatewayResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GatewayResponse) SetCode(v string)`

SetCode sets Code field to given value.


### GetData

`func (o *GatewayResponse) GetData() JitProgramResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GatewayResponse) GetDataOk() (*JitProgramResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GatewayResponse) SetData(v JitProgramResponse)`

SetData sets Data field to given value.

### HasData

`func (o *GatewayResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


