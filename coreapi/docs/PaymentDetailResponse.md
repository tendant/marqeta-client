# PaymentDetailResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountToken** | **string** | Unique identifier of the credit account on which the payment is made. | 
**Amount** | **float32** | Total amount of the payment. | 
**CreatedTime** | **time.Time** | Date and time when the payment was created on Marqeta&#39;s credit platform, in UTC. | 
**CurrencyCode** | [**CurrencyCode**](CurrencyCode.md) |  | [default to CURRENCYCODE_USD]
**Description** | **string** | Description of the payment. | 
**HoldDays** | **int32** | After a payment completes, the number of days to hold the available credit on the account before increasing it. | [default to 0]
**HoldEndTime** | Pointer to **time.Time** | Date and time when the available credit hold is released. | [optional] 
**IsManuallyReleased** | Pointer to **bool** | Whether the available credit hold was manually released for this payment. | [optional] [default to false]
**Metadata** | Pointer to **string** | Customer-defined additional information about the payment (for example, a check number). | [optional] 
**Method** | **string** | Method of payment. | 
**OnHold** | Pointer to **bool** | Whether the available credit is on hold for this payment. | [optional] [default to false]
**PaymentScheduleToken** | Pointer to **string** | Unique identifier of the payment schedule. | [optional] 
**PaymentSourceToken** | Pointer to **string** | Unique identifier of the payment source. Required for ACH payments. | [optional] 
**RefundDetails** | Pointer to [**NullableRefundDetailsResponse**](RefundDetailsResponse.md) |  | [optional] 
**ReturnedDetails** | Pointer to [**NullableReturnedDetails**](ReturnedDetails.md) |  | [optional] 
**Status** | [**PaymentStatus**](PaymentStatus.md) |  | 
**Token** | **string** | Unique identifier of the payment.  If in the &#x60;detail_object&#x60;, unique identifier of the detail object. | 
**Transitions** | [**[]PaymentTransitionResponse**](PaymentTransitionResponse.md) | Contains one or more &#x60;transitions&#x60; objects, which contain information on a payment status transition. | 
**UpdatedTime** | **time.Time** | Date and time when the payment was last updated on Marqeta&#39;s credit platform, in UTC. | 

## Methods

### NewPaymentDetailResponse

`func NewPaymentDetailResponse(accountToken string, amount float32, createdTime time.Time, currencyCode CurrencyCode, description string, holdDays int32, method string, status PaymentStatus, token string, transitions []PaymentTransitionResponse, updatedTime time.Time, ) *PaymentDetailResponse`

NewPaymentDetailResponse instantiates a new PaymentDetailResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentDetailResponseWithDefaults

`func NewPaymentDetailResponseWithDefaults() *PaymentDetailResponse`

NewPaymentDetailResponseWithDefaults instantiates a new PaymentDetailResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountToken

`func (o *PaymentDetailResponse) GetAccountToken() string`

GetAccountToken returns the AccountToken field if non-nil, zero value otherwise.

### GetAccountTokenOk

`func (o *PaymentDetailResponse) GetAccountTokenOk() (*string, bool)`

GetAccountTokenOk returns a tuple with the AccountToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountToken

`func (o *PaymentDetailResponse) SetAccountToken(v string)`

SetAccountToken sets AccountToken field to given value.


### GetAmount

`func (o *PaymentDetailResponse) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PaymentDetailResponse) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PaymentDetailResponse) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetCreatedTime

`func (o *PaymentDetailResponse) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *PaymentDetailResponse) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *PaymentDetailResponse) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetCurrencyCode

`func (o *PaymentDetailResponse) GetCurrencyCode() CurrencyCode`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *PaymentDetailResponse) GetCurrencyCodeOk() (*CurrencyCode, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *PaymentDetailResponse) SetCurrencyCode(v CurrencyCode)`

SetCurrencyCode sets CurrencyCode field to given value.


### GetDescription

`func (o *PaymentDetailResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PaymentDetailResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PaymentDetailResponse) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetHoldDays

`func (o *PaymentDetailResponse) GetHoldDays() int32`

GetHoldDays returns the HoldDays field if non-nil, zero value otherwise.

### GetHoldDaysOk

`func (o *PaymentDetailResponse) GetHoldDaysOk() (*int32, bool)`

GetHoldDaysOk returns a tuple with the HoldDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldDays

`func (o *PaymentDetailResponse) SetHoldDays(v int32)`

SetHoldDays sets HoldDays field to given value.


### GetHoldEndTime

`func (o *PaymentDetailResponse) GetHoldEndTime() time.Time`

GetHoldEndTime returns the HoldEndTime field if non-nil, zero value otherwise.

### GetHoldEndTimeOk

`func (o *PaymentDetailResponse) GetHoldEndTimeOk() (*time.Time, bool)`

GetHoldEndTimeOk returns a tuple with the HoldEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldEndTime

`func (o *PaymentDetailResponse) SetHoldEndTime(v time.Time)`

SetHoldEndTime sets HoldEndTime field to given value.

### HasHoldEndTime

`func (o *PaymentDetailResponse) HasHoldEndTime() bool`

HasHoldEndTime returns a boolean if a field has been set.

### GetIsManuallyReleased

`func (o *PaymentDetailResponse) GetIsManuallyReleased() bool`

GetIsManuallyReleased returns the IsManuallyReleased field if non-nil, zero value otherwise.

### GetIsManuallyReleasedOk

`func (o *PaymentDetailResponse) GetIsManuallyReleasedOk() (*bool, bool)`

GetIsManuallyReleasedOk returns a tuple with the IsManuallyReleased field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsManuallyReleased

`func (o *PaymentDetailResponse) SetIsManuallyReleased(v bool)`

SetIsManuallyReleased sets IsManuallyReleased field to given value.

### HasIsManuallyReleased

`func (o *PaymentDetailResponse) HasIsManuallyReleased() bool`

HasIsManuallyReleased returns a boolean if a field has been set.

### GetMetadata

`func (o *PaymentDetailResponse) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PaymentDetailResponse) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PaymentDetailResponse) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PaymentDetailResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMethod

`func (o *PaymentDetailResponse) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *PaymentDetailResponse) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *PaymentDetailResponse) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetOnHold

`func (o *PaymentDetailResponse) GetOnHold() bool`

GetOnHold returns the OnHold field if non-nil, zero value otherwise.

### GetOnHoldOk

`func (o *PaymentDetailResponse) GetOnHoldOk() (*bool, bool)`

GetOnHoldOk returns a tuple with the OnHold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnHold

`func (o *PaymentDetailResponse) SetOnHold(v bool)`

SetOnHold sets OnHold field to given value.

### HasOnHold

`func (o *PaymentDetailResponse) HasOnHold() bool`

HasOnHold returns a boolean if a field has been set.

### GetPaymentScheduleToken

`func (o *PaymentDetailResponse) GetPaymentScheduleToken() string`

GetPaymentScheduleToken returns the PaymentScheduleToken field if non-nil, zero value otherwise.

### GetPaymentScheduleTokenOk

`func (o *PaymentDetailResponse) GetPaymentScheduleTokenOk() (*string, bool)`

GetPaymentScheduleTokenOk returns a tuple with the PaymentScheduleToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentScheduleToken

`func (o *PaymentDetailResponse) SetPaymentScheduleToken(v string)`

SetPaymentScheduleToken sets PaymentScheduleToken field to given value.

### HasPaymentScheduleToken

`func (o *PaymentDetailResponse) HasPaymentScheduleToken() bool`

HasPaymentScheduleToken returns a boolean if a field has been set.

### GetPaymentSourceToken

`func (o *PaymentDetailResponse) GetPaymentSourceToken() string`

GetPaymentSourceToken returns the PaymentSourceToken field if non-nil, zero value otherwise.

### GetPaymentSourceTokenOk

`func (o *PaymentDetailResponse) GetPaymentSourceTokenOk() (*string, bool)`

GetPaymentSourceTokenOk returns a tuple with the PaymentSourceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentSourceToken

`func (o *PaymentDetailResponse) SetPaymentSourceToken(v string)`

SetPaymentSourceToken sets PaymentSourceToken field to given value.

### HasPaymentSourceToken

`func (o *PaymentDetailResponse) HasPaymentSourceToken() bool`

HasPaymentSourceToken returns a boolean if a field has been set.

### GetRefundDetails

`func (o *PaymentDetailResponse) GetRefundDetails() RefundDetailsResponse`

GetRefundDetails returns the RefundDetails field if non-nil, zero value otherwise.

### GetRefundDetailsOk

`func (o *PaymentDetailResponse) GetRefundDetailsOk() (*RefundDetailsResponse, bool)`

GetRefundDetailsOk returns a tuple with the RefundDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundDetails

`func (o *PaymentDetailResponse) SetRefundDetails(v RefundDetailsResponse)`

SetRefundDetails sets RefundDetails field to given value.

### HasRefundDetails

`func (o *PaymentDetailResponse) HasRefundDetails() bool`

HasRefundDetails returns a boolean if a field has been set.

### SetRefundDetailsNil

`func (o *PaymentDetailResponse) SetRefundDetailsNil(b bool)`

 SetRefundDetailsNil sets the value for RefundDetails to be an explicit nil

### UnsetRefundDetails
`func (o *PaymentDetailResponse) UnsetRefundDetails()`

UnsetRefundDetails ensures that no value is present for RefundDetails, not even an explicit nil
### GetReturnedDetails

`func (o *PaymentDetailResponse) GetReturnedDetails() ReturnedDetails`

GetReturnedDetails returns the ReturnedDetails field if non-nil, zero value otherwise.

### GetReturnedDetailsOk

`func (o *PaymentDetailResponse) GetReturnedDetailsOk() (*ReturnedDetails, bool)`

GetReturnedDetailsOk returns a tuple with the ReturnedDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnedDetails

`func (o *PaymentDetailResponse) SetReturnedDetails(v ReturnedDetails)`

SetReturnedDetails sets ReturnedDetails field to given value.

### HasReturnedDetails

`func (o *PaymentDetailResponse) HasReturnedDetails() bool`

HasReturnedDetails returns a boolean if a field has been set.

### SetReturnedDetailsNil

`func (o *PaymentDetailResponse) SetReturnedDetailsNil(b bool)`

 SetReturnedDetailsNil sets the value for ReturnedDetails to be an explicit nil

### UnsetReturnedDetails
`func (o *PaymentDetailResponse) UnsetReturnedDetails()`

UnsetReturnedDetails ensures that no value is present for ReturnedDetails, not even an explicit nil
### GetStatus

`func (o *PaymentDetailResponse) GetStatus() PaymentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentDetailResponse) GetStatusOk() (*PaymentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentDetailResponse) SetStatus(v PaymentStatus)`

SetStatus sets Status field to given value.


### GetToken

`func (o *PaymentDetailResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PaymentDetailResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PaymentDetailResponse) SetToken(v string)`

SetToken sets Token field to given value.


### GetTransitions

`func (o *PaymentDetailResponse) GetTransitions() []PaymentTransitionResponse`

GetTransitions returns the Transitions field if non-nil, zero value otherwise.

### GetTransitionsOk

`func (o *PaymentDetailResponse) GetTransitionsOk() (*[]PaymentTransitionResponse, bool)`

GetTransitionsOk returns a tuple with the Transitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitions

`func (o *PaymentDetailResponse) SetTransitions(v []PaymentTransitionResponse)`

SetTransitions sets Transitions field to given value.


### GetUpdatedTime

`func (o *PaymentDetailResponse) GetUpdatedTime() time.Time`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *PaymentDetailResponse) GetUpdatedTimeOk() (*time.Time, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *PaymentDetailResponse) SetUpdatedTime(v time.Time)`

SetUpdatedTime sets UpdatedTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


