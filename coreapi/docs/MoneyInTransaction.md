# MoneyInTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncludeCreditsTypes** | Pointer to **[]string** | Specifies the types of credits to include in the original credit transaction (OCT). | [optional] 
**IncludeNetworkLoadTypes** | Pointer to **[]string** | Indicates whether or not cash and check network load transactions will be subject to velocity control. | [optional] 

## Methods

### NewMoneyInTransaction

`func NewMoneyInTransaction() *MoneyInTransaction`

NewMoneyInTransaction instantiates a new MoneyInTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoneyInTransactionWithDefaults

`func NewMoneyInTransactionWithDefaults() *MoneyInTransaction`

NewMoneyInTransactionWithDefaults instantiates a new MoneyInTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncludeCreditsTypes

`func (o *MoneyInTransaction) GetIncludeCreditsTypes() []string`

GetIncludeCreditsTypes returns the IncludeCreditsTypes field if non-nil, zero value otherwise.

### GetIncludeCreditsTypesOk

`func (o *MoneyInTransaction) GetIncludeCreditsTypesOk() (*[]string, bool)`

GetIncludeCreditsTypesOk returns a tuple with the IncludeCreditsTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeCreditsTypes

`func (o *MoneyInTransaction) SetIncludeCreditsTypes(v []string)`

SetIncludeCreditsTypes sets IncludeCreditsTypes field to given value.

### HasIncludeCreditsTypes

`func (o *MoneyInTransaction) HasIncludeCreditsTypes() bool`

HasIncludeCreditsTypes returns a boolean if a field has been set.

### GetIncludeNetworkLoadTypes

`func (o *MoneyInTransaction) GetIncludeNetworkLoadTypes() []string`

GetIncludeNetworkLoadTypes returns the IncludeNetworkLoadTypes field if non-nil, zero value otherwise.

### GetIncludeNetworkLoadTypesOk

`func (o *MoneyInTransaction) GetIncludeNetworkLoadTypesOk() (*[]string, bool)`

GetIncludeNetworkLoadTypesOk returns a tuple with the IncludeNetworkLoadTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeNetworkLoadTypes

`func (o *MoneyInTransaction) SetIncludeNetworkLoadTypes(v []string)`

SetIncludeNetworkLoadTypes sets IncludeNetworkLoadTypes field to given value.

### HasIncludeNetworkLoadTypes

`func (o *MoneyInTransaction) HasIncludeNetworkLoadTypes() bool`

HasIncludeNetworkLoadTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


