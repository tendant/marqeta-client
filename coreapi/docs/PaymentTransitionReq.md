# PaymentTransitionReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RefundDetails** | Pointer to [**RefundDetails**](RefundDetails.md) |  | [optional] 
**Status** | [**PaymentStatus**](PaymentStatus.md) |  | 
**Token** | Pointer to **string** | Unique identifier of the payment status transition. | [optional] 

## Methods

### NewPaymentTransitionReq

`func NewPaymentTransitionReq(status PaymentStatus, ) *PaymentTransitionReq`

NewPaymentTransitionReq instantiates a new PaymentTransitionReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentTransitionReqWithDefaults

`func NewPaymentTransitionReqWithDefaults() *PaymentTransitionReq`

NewPaymentTransitionReqWithDefaults instantiates a new PaymentTransitionReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRefundDetails

`func (o *PaymentTransitionReq) GetRefundDetails() RefundDetails`

GetRefundDetails returns the RefundDetails field if non-nil, zero value otherwise.

### GetRefundDetailsOk

`func (o *PaymentTransitionReq) GetRefundDetailsOk() (*RefundDetails, bool)`

GetRefundDetailsOk returns a tuple with the RefundDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundDetails

`func (o *PaymentTransitionReq) SetRefundDetails(v RefundDetails)`

SetRefundDetails sets RefundDetails field to given value.

### HasRefundDetails

`func (o *PaymentTransitionReq) HasRefundDetails() bool`

HasRefundDetails returns a boolean if a field has been set.

### GetStatus

`func (o *PaymentTransitionReq) GetStatus() PaymentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentTransitionReq) GetStatusOk() (*PaymentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentTransitionReq) SetStatus(v PaymentStatus)`

SetStatus sets Status field to given value.


### GetToken

`func (o *PaymentTransitionReq) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PaymentTransitionReq) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PaymentTransitionReq) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *PaymentTransitionReq) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


