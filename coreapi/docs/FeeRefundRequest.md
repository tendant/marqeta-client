# FeeRefundRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OriginalFeeTransactionToken** | Pointer to **string** | Unique identifier of the fee to be refunded.   You can find this token in the response of the original &#x60;/feecharges&#x60; or &#x60;/gpaorders&#x60; request used to assess the standalone fee or from the webhook corresponding to the original request. You can also send a &#x60;GET&#x60; request to &#x60;/transactions?type&#x3D;fee.charges&#x60; to retrieve a list of fee transactions. | [optional] 
**Tags** | Pointer to **string** | Descriptive metadata about the fee. | [optional] 

## Methods

### NewFeeRefundRequest

`func NewFeeRefundRequest() *FeeRefundRequest`

NewFeeRefundRequest instantiates a new FeeRefundRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeeRefundRequestWithDefaults

`func NewFeeRefundRequestWithDefaults() *FeeRefundRequest`

NewFeeRefundRequestWithDefaults instantiates a new FeeRefundRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOriginalFeeTransactionToken

`func (o *FeeRefundRequest) GetOriginalFeeTransactionToken() string`

GetOriginalFeeTransactionToken returns the OriginalFeeTransactionToken field if non-nil, zero value otherwise.

### GetOriginalFeeTransactionTokenOk

`func (o *FeeRefundRequest) GetOriginalFeeTransactionTokenOk() (*string, bool)`

GetOriginalFeeTransactionTokenOk returns a tuple with the OriginalFeeTransactionToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalFeeTransactionToken

`func (o *FeeRefundRequest) SetOriginalFeeTransactionToken(v string)`

SetOriginalFeeTransactionToken sets OriginalFeeTransactionToken field to given value.

### HasOriginalFeeTransactionToken

`func (o *FeeRefundRequest) HasOriginalFeeTransactionToken() bool`

HasOriginalFeeTransactionToken returns a boolean if a field has been set.

### GetTags

`func (o *FeeRefundRequest) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FeeRefundRequest) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FeeRefundRequest) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *FeeRefundRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


