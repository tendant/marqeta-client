# ProductMinPaymentPercentage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncludeFeesCharged** | [**[]ProductFeeType**](ProductFeeType.md) | One or more fee types to include when calculating the minimum payment. | 
**PercentageOfBalance** | **float32** | Minimum payment, expressed as a percentage of the total statement balance, due on the payment due day. | 

## Methods

### NewProductMinPaymentPercentage

`func NewProductMinPaymentPercentage(includeFeesCharged []ProductFeeType, percentageOfBalance float32, ) *ProductMinPaymentPercentage`

NewProductMinPaymentPercentage instantiates a new ProductMinPaymentPercentage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductMinPaymentPercentageWithDefaults

`func NewProductMinPaymentPercentageWithDefaults() *ProductMinPaymentPercentage`

NewProductMinPaymentPercentageWithDefaults instantiates a new ProductMinPaymentPercentage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncludeFeesCharged

`func (o *ProductMinPaymentPercentage) GetIncludeFeesCharged() []ProductFeeType`

GetIncludeFeesCharged returns the IncludeFeesCharged field if non-nil, zero value otherwise.

### GetIncludeFeesChargedOk

`func (o *ProductMinPaymentPercentage) GetIncludeFeesChargedOk() (*[]ProductFeeType, bool)`

GetIncludeFeesChargedOk returns a tuple with the IncludeFeesCharged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeFeesCharged

`func (o *ProductMinPaymentPercentage) SetIncludeFeesCharged(v []ProductFeeType)`

SetIncludeFeesCharged sets IncludeFeesCharged field to given value.


### GetPercentageOfBalance

`func (o *ProductMinPaymentPercentage) GetPercentageOfBalance() float32`

GetPercentageOfBalance returns the PercentageOfBalance field if non-nil, zero value otherwise.

### GetPercentageOfBalanceOk

`func (o *ProductMinPaymentPercentage) GetPercentageOfBalanceOk() (*float32, bool)`

GetPercentageOfBalanceOk returns a tuple with the PercentageOfBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageOfBalance

`func (o *ProductMinPaymentPercentage) SetPercentageOfBalance(v float32)`

SetPercentageOfBalance sets PercentageOfBalance field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


