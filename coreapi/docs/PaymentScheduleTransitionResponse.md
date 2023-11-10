# PaymentScheduleTransitionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountToken** | Pointer to **string** | Unique identifier of the credit account on which to transition a payment schedule.  Send a &#x60;GET&#x60; request to &#x60;/credit/accounts&#x60; to retrieve existing credit account tokens. | [optional] 
**CreatedTime** | Pointer to **time.Time** | Date and time when the payment schedule transition was created on Marqeta&#39;s credit platform, in UTC. | [optional] 
**PaymentScheduleToken** | Pointer to **string** | The unique identifier of the payment schedule whose status is to transition.  Send a &#x60;GET&#x60; request to &#x60;/credit/accounts/{account_token}/paymentschedules&#x60; to retrieve existing payment schedule tokens. | [optional] 
**Status** | Pointer to [**PaymentScheduleStatus**](PaymentScheduleStatus.md) |  | [optional] 
**Token** | Pointer to **string** | Unique identifier of the payment schedule transition. | [optional] 
**UpdatedTime** | Pointer to **time.Time** | Date and time when the payment schedule transition was last updated on Marqeta&#39;s credit platform, in UTC. | [optional] 

## Methods

### NewPaymentScheduleTransitionResponse

`func NewPaymentScheduleTransitionResponse() *PaymentScheduleTransitionResponse`

NewPaymentScheduleTransitionResponse instantiates a new PaymentScheduleTransitionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentScheduleTransitionResponseWithDefaults

`func NewPaymentScheduleTransitionResponseWithDefaults() *PaymentScheduleTransitionResponse`

NewPaymentScheduleTransitionResponseWithDefaults instantiates a new PaymentScheduleTransitionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountToken

`func (o *PaymentScheduleTransitionResponse) GetAccountToken() string`

GetAccountToken returns the AccountToken field if non-nil, zero value otherwise.

### GetAccountTokenOk

`func (o *PaymentScheduleTransitionResponse) GetAccountTokenOk() (*string, bool)`

GetAccountTokenOk returns a tuple with the AccountToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountToken

`func (o *PaymentScheduleTransitionResponse) SetAccountToken(v string)`

SetAccountToken sets AccountToken field to given value.

### HasAccountToken

`func (o *PaymentScheduleTransitionResponse) HasAccountToken() bool`

HasAccountToken returns a boolean if a field has been set.

### GetCreatedTime

`func (o *PaymentScheduleTransitionResponse) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *PaymentScheduleTransitionResponse) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *PaymentScheduleTransitionResponse) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *PaymentScheduleTransitionResponse) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetPaymentScheduleToken

`func (o *PaymentScheduleTransitionResponse) GetPaymentScheduleToken() string`

GetPaymentScheduleToken returns the PaymentScheduleToken field if non-nil, zero value otherwise.

### GetPaymentScheduleTokenOk

`func (o *PaymentScheduleTransitionResponse) GetPaymentScheduleTokenOk() (*string, bool)`

GetPaymentScheduleTokenOk returns a tuple with the PaymentScheduleToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentScheduleToken

`func (o *PaymentScheduleTransitionResponse) SetPaymentScheduleToken(v string)`

SetPaymentScheduleToken sets PaymentScheduleToken field to given value.

### HasPaymentScheduleToken

`func (o *PaymentScheduleTransitionResponse) HasPaymentScheduleToken() bool`

HasPaymentScheduleToken returns a boolean if a field has been set.

### GetStatus

`func (o *PaymentScheduleTransitionResponse) GetStatus() PaymentScheduleStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentScheduleTransitionResponse) GetStatusOk() (*PaymentScheduleStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentScheduleTransitionResponse) SetStatus(v PaymentScheduleStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PaymentScheduleTransitionResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetToken

`func (o *PaymentScheduleTransitionResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PaymentScheduleTransitionResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PaymentScheduleTransitionResponse) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *PaymentScheduleTransitionResponse) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUpdatedTime

`func (o *PaymentScheduleTransitionResponse) GetUpdatedTime() time.Time`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *PaymentScheduleTransitionResponse) GetUpdatedTimeOk() (*time.Time, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *PaymentScheduleTransitionResponse) SetUpdatedTime(v time.Time)`

SetUpdatedTime sets UpdatedTime field to given value.

### HasUpdatedTime

`func (o *PaymentScheduleTransitionResponse) HasUpdatedTime() bool`

HasUpdatedTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


