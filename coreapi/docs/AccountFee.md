# AccountFee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | [**Method**](Method.md) |  | 
**Type** | [**FeeType**](FeeType.md) |  | 
**Value** | Pointer to **float32** | Value of the fee, either a flat fee amount or percentage value. | [optional] 

## Methods

### NewAccountFee

`func NewAccountFee(method Method, type_ FeeType, ) *AccountFee`

NewAccountFee instantiates a new AccountFee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountFeeWithDefaults

`func NewAccountFeeWithDefaults() *AccountFee`

NewAccountFeeWithDefaults instantiates a new AccountFee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *AccountFee) GetMethod() Method`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *AccountFee) GetMethodOk() (*Method, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *AccountFee) SetMethod(v Method)`

SetMethod sets Method field to given value.


### GetType

`func (o *AccountFee) GetType() FeeType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AccountFee) GetTypeOk() (*FeeType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AccountFee) SetType(v FeeType)`

SetType sets Type field to given value.


### GetValue

`func (o *AccountFee) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AccountFee) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AccountFee) SetValue(v float32)`

SetValue sets Value field to given value.

### HasValue

`func (o *AccountFee) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


