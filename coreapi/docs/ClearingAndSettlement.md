# ClearingAndSettlement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OverdraftDestination** | Pointer to **string** | Specifies the destination for overdraft funds.  This field does not apply if JIT Funding is enabled. | [optional] [default to "GPA"]

## Methods

### NewClearingAndSettlement

`func NewClearingAndSettlement() *ClearingAndSettlement`

NewClearingAndSettlement instantiates a new ClearingAndSettlement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClearingAndSettlementWithDefaults

`func NewClearingAndSettlementWithDefaults() *ClearingAndSettlement`

NewClearingAndSettlementWithDefaults instantiates a new ClearingAndSettlement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOverdraftDestination

`func (o *ClearingAndSettlement) GetOverdraftDestination() string`

GetOverdraftDestination returns the OverdraftDestination field if non-nil, zero value otherwise.

### GetOverdraftDestinationOk

`func (o *ClearingAndSettlement) GetOverdraftDestinationOk() (*string, bool)`

GetOverdraftDestinationOk returns a tuple with the OverdraftDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverdraftDestination

`func (o *ClearingAndSettlement) SetOverdraftDestination(v string)`

SetOverdraftDestination sets OverdraftDestination field to given value.

### HasOverdraftDestination

`func (o *ClearingAndSettlement) HasOverdraftDestination() bool`

HasOverdraftDestination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


