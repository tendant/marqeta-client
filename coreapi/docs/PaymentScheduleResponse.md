# PaymentScheduleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountToken** | **string** | Unique identifier of the credit account on which the payment schedule is made. | 
**Amount** | Pointer to **NullableFloat32** | Amount of the payment.  Returned if the &#x60;amount_category&#x60; is &#x60;FIXED&#x60;. | [optional] 
**AmountCategory** | [**PaymentScheduleAmountCategory**](PaymentScheduleAmountCategory.md) |  | 
**CreatedTime** | Pointer to **time.Time** | Date and time when the payment schedule was created on Marqeta&#39;s credit platform, in UTC. | [optional] 
**CurrencyCode** | [**CurrencyCode**](CurrencyCode.md) |  | [default to CURRENCYCODE_USD]
**Description** | Pointer to **string** | Description of the payment schedule. | [optional] 
**Frequency** | [**PaymentScheduleFrequency**](PaymentScheduleFrequency.md) |  | 
**NextPaymentImpactDate** | Pointer to **string** | The date to make a one-time payment.  Returned if &#x60;frequency&#x60; is &#x60;ONCE&#x60;. | [optional] 
**PaymentDay** | Pointer to **string** | Day on which monthly payments are made.  Returned if the &#x60;frequency&#x60; is &#x60;MONTHLY&#x60;. | [optional] 
**PaymentSourceToken** | **string** | Unique identifier of a payment source. | 
**Status** | [**PaymentScheduleStatus**](PaymentScheduleStatus.md) |  | 
**Token** | **string** | Unique identifier of the payment schedule. | 
**UpdatedTime** | Pointer to **time.Time** | Date and time when the payment schedule was last updated on Marqeta&#39;s credit platform, in UTC. | [optional] 

## Methods

### NewPaymentScheduleResponse

`func NewPaymentScheduleResponse(accountToken string, amountCategory PaymentScheduleAmountCategory, currencyCode CurrencyCode, frequency PaymentScheduleFrequency, paymentSourceToken string, status PaymentScheduleStatus, token string, ) *PaymentScheduleResponse`

NewPaymentScheduleResponse instantiates a new PaymentScheduleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentScheduleResponseWithDefaults

`func NewPaymentScheduleResponseWithDefaults() *PaymentScheduleResponse`

NewPaymentScheduleResponseWithDefaults instantiates a new PaymentScheduleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountToken

`func (o *PaymentScheduleResponse) GetAccountToken() string`

GetAccountToken returns the AccountToken field if non-nil, zero value otherwise.

### GetAccountTokenOk

`func (o *PaymentScheduleResponse) GetAccountTokenOk() (*string, bool)`

GetAccountTokenOk returns a tuple with the AccountToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountToken

`func (o *PaymentScheduleResponse) SetAccountToken(v string)`

SetAccountToken sets AccountToken field to given value.


### GetAmount

`func (o *PaymentScheduleResponse) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PaymentScheduleResponse) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PaymentScheduleResponse) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *PaymentScheduleResponse) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### SetAmountNil

`func (o *PaymentScheduleResponse) SetAmountNil(b bool)`

 SetAmountNil sets the value for Amount to be an explicit nil

### UnsetAmount
`func (o *PaymentScheduleResponse) UnsetAmount()`

UnsetAmount ensures that no value is present for Amount, not even an explicit nil
### GetAmountCategory

`func (o *PaymentScheduleResponse) GetAmountCategory() PaymentScheduleAmountCategory`

GetAmountCategory returns the AmountCategory field if non-nil, zero value otherwise.

### GetAmountCategoryOk

`func (o *PaymentScheduleResponse) GetAmountCategoryOk() (*PaymentScheduleAmountCategory, bool)`

GetAmountCategoryOk returns a tuple with the AmountCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCategory

`func (o *PaymentScheduleResponse) SetAmountCategory(v PaymentScheduleAmountCategory)`

SetAmountCategory sets AmountCategory field to given value.


### GetCreatedTime

`func (o *PaymentScheduleResponse) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *PaymentScheduleResponse) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *PaymentScheduleResponse) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *PaymentScheduleResponse) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *PaymentScheduleResponse) GetCurrencyCode() CurrencyCode`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *PaymentScheduleResponse) GetCurrencyCodeOk() (*CurrencyCode, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *PaymentScheduleResponse) SetCurrencyCode(v CurrencyCode)`

SetCurrencyCode sets CurrencyCode field to given value.


### GetDescription

`func (o *PaymentScheduleResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PaymentScheduleResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PaymentScheduleResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PaymentScheduleResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFrequency

`func (o *PaymentScheduleResponse) GetFrequency() PaymentScheduleFrequency`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *PaymentScheduleResponse) GetFrequencyOk() (*PaymentScheduleFrequency, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *PaymentScheduleResponse) SetFrequency(v PaymentScheduleFrequency)`

SetFrequency sets Frequency field to given value.


### GetNextPaymentImpactDate

`func (o *PaymentScheduleResponse) GetNextPaymentImpactDate() string`

GetNextPaymentImpactDate returns the NextPaymentImpactDate field if non-nil, zero value otherwise.

### GetNextPaymentImpactDateOk

`func (o *PaymentScheduleResponse) GetNextPaymentImpactDateOk() (*string, bool)`

GetNextPaymentImpactDateOk returns a tuple with the NextPaymentImpactDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPaymentImpactDate

`func (o *PaymentScheduleResponse) SetNextPaymentImpactDate(v string)`

SetNextPaymentImpactDate sets NextPaymentImpactDate field to given value.

### HasNextPaymentImpactDate

`func (o *PaymentScheduleResponse) HasNextPaymentImpactDate() bool`

HasNextPaymentImpactDate returns a boolean if a field has been set.

### GetPaymentDay

`func (o *PaymentScheduleResponse) GetPaymentDay() string`

GetPaymentDay returns the PaymentDay field if non-nil, zero value otherwise.

### GetPaymentDayOk

`func (o *PaymentScheduleResponse) GetPaymentDayOk() (*string, bool)`

GetPaymentDayOk returns a tuple with the PaymentDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentDay

`func (o *PaymentScheduleResponse) SetPaymentDay(v string)`

SetPaymentDay sets PaymentDay field to given value.

### HasPaymentDay

`func (o *PaymentScheduleResponse) HasPaymentDay() bool`

HasPaymentDay returns a boolean if a field has been set.

### GetPaymentSourceToken

`func (o *PaymentScheduleResponse) GetPaymentSourceToken() string`

GetPaymentSourceToken returns the PaymentSourceToken field if non-nil, zero value otherwise.

### GetPaymentSourceTokenOk

`func (o *PaymentScheduleResponse) GetPaymentSourceTokenOk() (*string, bool)`

GetPaymentSourceTokenOk returns a tuple with the PaymentSourceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentSourceToken

`func (o *PaymentScheduleResponse) SetPaymentSourceToken(v string)`

SetPaymentSourceToken sets PaymentSourceToken field to given value.


### GetStatus

`func (o *PaymentScheduleResponse) GetStatus() PaymentScheduleStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentScheduleResponse) GetStatusOk() (*PaymentScheduleStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentScheduleResponse) SetStatus(v PaymentScheduleStatus)`

SetStatus sets Status field to given value.


### GetToken

`func (o *PaymentScheduleResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PaymentScheduleResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PaymentScheduleResponse) SetToken(v string)`

SetToken sets Token field to given value.


### GetUpdatedTime

`func (o *PaymentScheduleResponse) GetUpdatedTime() time.Time`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *PaymentScheduleResponse) GetUpdatedTimeOk() (*time.Time, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *PaymentScheduleResponse) SetUpdatedTime(v time.Time)`

SetUpdatedTime sets UpdatedTime field to given value.

### HasUpdatedTime

`func (o *PaymentScheduleResponse) HasUpdatedTime() bool`

HasUpdatedTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


