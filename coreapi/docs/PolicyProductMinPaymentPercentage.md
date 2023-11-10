# PolicyProductMinPaymentPercentage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncludeFeesCharged** | Pointer to **[]string** | One or more fee types to include when calculating the minimum payment. | [optional] 
**IncludeInterestCharged** | **bool** | Whether to include the amount of interest charged when calculating the minimum payment. | 
**PercentageOfBalance** | **float32** | Minimum payment, expressed as a percentage of the total statement balance, due on the payment due day. | 

## Methods

### NewPolicyProductMinPaymentPercentage

`func NewPolicyProductMinPaymentPercentage(includeInterestCharged bool, percentageOfBalance float32, ) *PolicyProductMinPaymentPercentage`

NewPolicyProductMinPaymentPercentage instantiates a new PolicyProductMinPaymentPercentage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyProductMinPaymentPercentageWithDefaults

`func NewPolicyProductMinPaymentPercentageWithDefaults() *PolicyProductMinPaymentPercentage`

NewPolicyProductMinPaymentPercentageWithDefaults instantiates a new PolicyProductMinPaymentPercentage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncludeFeesCharged

`func (o *PolicyProductMinPaymentPercentage) GetIncludeFeesCharged() []string`

GetIncludeFeesCharged returns the IncludeFeesCharged field if non-nil, zero value otherwise.

### GetIncludeFeesChargedOk

`func (o *PolicyProductMinPaymentPercentage) GetIncludeFeesChargedOk() (*[]string, bool)`

GetIncludeFeesChargedOk returns a tuple with the IncludeFeesCharged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeFeesCharged

`func (o *PolicyProductMinPaymentPercentage) SetIncludeFeesCharged(v []string)`

SetIncludeFeesCharged sets IncludeFeesCharged field to given value.

### HasIncludeFeesCharged

`func (o *PolicyProductMinPaymentPercentage) HasIncludeFeesCharged() bool`

HasIncludeFeesCharged returns a boolean if a field has been set.

### GetIncludeInterestCharged

`func (o *PolicyProductMinPaymentPercentage) GetIncludeInterestCharged() bool`

GetIncludeInterestCharged returns the IncludeInterestCharged field if non-nil, zero value otherwise.

### GetIncludeInterestChargedOk

`func (o *PolicyProductMinPaymentPercentage) GetIncludeInterestChargedOk() (*bool, bool)`

GetIncludeInterestChargedOk returns a tuple with the IncludeInterestCharged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeInterestCharged

`func (o *PolicyProductMinPaymentPercentage) SetIncludeInterestCharged(v bool)`

SetIncludeInterestCharged sets IncludeInterestCharged field to given value.


### GetPercentageOfBalance

`func (o *PolicyProductMinPaymentPercentage) GetPercentageOfBalance() float32`

GetPercentageOfBalance returns the PercentageOfBalance field if non-nil, zero value otherwise.

### GetPercentageOfBalanceOk

`func (o *PolicyProductMinPaymentPercentage) GetPercentageOfBalanceOk() (*float32, bool)`

GetPercentageOfBalanceOk returns a tuple with the PercentageOfBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageOfBalance

`func (o *PolicyProductMinPaymentPercentage) SetPercentageOfBalance(v float32)`

SetPercentageOfBalance sets PercentageOfBalance field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


