# CardholderBalances

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gpa** | [**CardholderBalance**](CardholderBalance.md) |  | 
**Links** | [**[]Link**](Link.md) | Array of links to balances of related user GPAs. | 

## Methods

### NewCardholderBalances

`func NewCardholderBalances(gpa CardholderBalance, links []Link, ) *CardholderBalances`

NewCardholderBalances instantiates a new CardholderBalances object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardholderBalancesWithDefaults

`func NewCardholderBalancesWithDefaults() *CardholderBalances`

NewCardholderBalancesWithDefaults instantiates a new CardholderBalances object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGpa

`func (o *CardholderBalances) GetGpa() CardholderBalance`

GetGpa returns the Gpa field if non-nil, zero value otherwise.

### GetGpaOk

`func (o *CardholderBalances) GetGpaOk() (*CardholderBalance, bool)`

GetGpaOk returns a tuple with the Gpa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpa

`func (o *CardholderBalances) SetGpa(v CardholderBalance)`

SetGpa sets Gpa field to given value.


### GetLinks

`func (o *CardholderBalances) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CardholderBalances) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CardholderBalances) SetLinks(v []Link)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


