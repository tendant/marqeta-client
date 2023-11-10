# PaymentScheduleTransitionCreateReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**PaymentScheduleStatus**](PaymentScheduleStatus.md) |  | 
**Token** | Pointer to **string** | Unique identifier of the payment schedule transition. | [optional] 

## Methods

### NewPaymentScheduleTransitionCreateReq

`func NewPaymentScheduleTransitionCreateReq(status PaymentScheduleStatus, ) *PaymentScheduleTransitionCreateReq`

NewPaymentScheduleTransitionCreateReq instantiates a new PaymentScheduleTransitionCreateReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentScheduleTransitionCreateReqWithDefaults

`func NewPaymentScheduleTransitionCreateReqWithDefaults() *PaymentScheduleTransitionCreateReq`

NewPaymentScheduleTransitionCreateReqWithDefaults instantiates a new PaymentScheduleTransitionCreateReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *PaymentScheduleTransitionCreateReq) GetStatus() PaymentScheduleStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentScheduleTransitionCreateReq) GetStatusOk() (*PaymentScheduleStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentScheduleTransitionCreateReq) SetStatus(v PaymentScheduleStatus)`

SetStatus sets Status field to given value.


### GetToken

`func (o *PaymentScheduleTransitionCreateReq) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PaymentScheduleTransitionCreateReq) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PaymentScheduleTransitionCreateReq) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *PaymentScheduleTransitionCreateReq) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


