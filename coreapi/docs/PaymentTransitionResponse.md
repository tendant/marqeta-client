# PaymentTransitionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountToken** | **string** | Unique identifier of the credit account on which you want to transition a payment status. | 
**CreatedTime** | Pointer to **time.Time** | Date and time when the payment transition was created on Marqeta&#39;s credit platform, in UTC. | [optional] 
**PaymentToken** | **string** | Unique identifier of the payment whose status you want to transition. | 
**Status** | [**PaymentStatus**](PaymentStatus.md) |  | 
**Token** | **string** | Unique identifier of the payment status transition. | 

## Methods

### NewPaymentTransitionResponse

`func NewPaymentTransitionResponse(accountToken string, paymentToken string, status PaymentStatus, token string, ) *PaymentTransitionResponse`

NewPaymentTransitionResponse instantiates a new PaymentTransitionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentTransitionResponseWithDefaults

`func NewPaymentTransitionResponseWithDefaults() *PaymentTransitionResponse`

NewPaymentTransitionResponseWithDefaults instantiates a new PaymentTransitionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountToken

`func (o *PaymentTransitionResponse) GetAccountToken() string`

GetAccountToken returns the AccountToken field if non-nil, zero value otherwise.

### GetAccountTokenOk

`func (o *PaymentTransitionResponse) GetAccountTokenOk() (*string, bool)`

GetAccountTokenOk returns a tuple with the AccountToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountToken

`func (o *PaymentTransitionResponse) SetAccountToken(v string)`

SetAccountToken sets AccountToken field to given value.


### GetCreatedTime

`func (o *PaymentTransitionResponse) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *PaymentTransitionResponse) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *PaymentTransitionResponse) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *PaymentTransitionResponse) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetPaymentToken

`func (o *PaymentTransitionResponse) GetPaymentToken() string`

GetPaymentToken returns the PaymentToken field if non-nil, zero value otherwise.

### GetPaymentTokenOk

`func (o *PaymentTransitionResponse) GetPaymentTokenOk() (*string, bool)`

GetPaymentTokenOk returns a tuple with the PaymentToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentToken

`func (o *PaymentTransitionResponse) SetPaymentToken(v string)`

SetPaymentToken sets PaymentToken field to given value.


### GetStatus

`func (o *PaymentTransitionResponse) GetStatus() PaymentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentTransitionResponse) GetStatusOk() (*PaymentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentTransitionResponse) SetStatus(v PaymentStatus)`

SetStatus sets Status field to given value.


### GetToken

`func (o *PaymentTransitionResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PaymentTransitionResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PaymentTransitionResponse) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


