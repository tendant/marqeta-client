# FeeAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | Pointer to **string** | Describes the reason for the fee. | [optional] 
**Region** | Pointer to **string** | Describes the region in which the fee is assessed. | [optional] 
**Status** | Pointer to **string** | Describes the status of the fee. | [optional] 
**TransactionType** | Pointer to **string** | Specifies the transaction type in which the fee was assessed. | [optional] 

## Methods

### NewFeeAttributes

`func NewFeeAttributes() *FeeAttributes`

NewFeeAttributes instantiates a new FeeAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeeAttributesWithDefaults

`func NewFeeAttributesWithDefaults() *FeeAttributes`

NewFeeAttributesWithDefaults instantiates a new FeeAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *FeeAttributes) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *FeeAttributes) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *FeeAttributes) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *FeeAttributes) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetRegion

`func (o *FeeAttributes) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *FeeAttributes) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *FeeAttributes) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *FeeAttributes) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetStatus

`func (o *FeeAttributes) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FeeAttributes) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FeeAttributes) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FeeAttributes) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTransactionType

`func (o *FeeAttributes) GetTransactionType() string`

GetTransactionType returns the TransactionType field if non-nil, zero value otherwise.

### GetTransactionTypeOk

`func (o *FeeAttributes) GetTransactionTypeOk() (*string, bool)`

GetTransactionTypeOk returns a tuple with the TransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionType

`func (o *FeeAttributes) SetTransactionType(v string)`

SetTransactionType sets TransactionType field to given value.

### HasTransactionType

`func (o *FeeAttributes) HasTransactionType() bool`

HasTransactionType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


