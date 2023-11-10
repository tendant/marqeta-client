# PaymentScheduleCreateReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **float32** | Amount of the payment.  Required if &#x60;amount_category&#x60; is &#x60;FIXED&#x60;. | [optional] 
**AmountCategory** | [**PaymentScheduleAmountCategory**](PaymentScheduleAmountCategory.md) |  | 
**CurrencyCode** | [**CurrencyCode**](CurrencyCode.md) |  | [default to CURRENCYCODE_USD]
**Description** | Pointer to **string** | Description of the payment schedule. | [optional] 
**Frequency** | [**PaymentScheduleFrequency**](PaymentScheduleFrequency.md) |  | 
**NextPaymentImpactDate** | Pointer to **string** | The date to make a one-time payment.  Required if frequency is &#x60;ONCE&#x60;. | [optional] 
**PaymentDay** | Pointer to **string** | Day on which monthly payments are made.  Required if &#x60;frequency&#x60; is &#x60;MONTHLY&#x60;. | [optional] 
**PaymentSourceToken** | **string** | Unique identifier of the payment source. | 
**Token** | Pointer to **string** | Unique identifier of the payment schedule. | [optional] 

## Methods

### NewPaymentScheduleCreateReq

`func NewPaymentScheduleCreateReq(amountCategory PaymentScheduleAmountCategory, currencyCode CurrencyCode, frequency PaymentScheduleFrequency, paymentSourceToken string, ) *PaymentScheduleCreateReq`

NewPaymentScheduleCreateReq instantiates a new PaymentScheduleCreateReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentScheduleCreateReqWithDefaults

`func NewPaymentScheduleCreateReqWithDefaults() *PaymentScheduleCreateReq`

NewPaymentScheduleCreateReqWithDefaults instantiates a new PaymentScheduleCreateReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *PaymentScheduleCreateReq) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PaymentScheduleCreateReq) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PaymentScheduleCreateReq) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *PaymentScheduleCreateReq) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetAmountCategory

`func (o *PaymentScheduleCreateReq) GetAmountCategory() PaymentScheduleAmountCategory`

GetAmountCategory returns the AmountCategory field if non-nil, zero value otherwise.

### GetAmountCategoryOk

`func (o *PaymentScheduleCreateReq) GetAmountCategoryOk() (*PaymentScheduleAmountCategory, bool)`

GetAmountCategoryOk returns a tuple with the AmountCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCategory

`func (o *PaymentScheduleCreateReq) SetAmountCategory(v PaymentScheduleAmountCategory)`

SetAmountCategory sets AmountCategory field to given value.


### GetCurrencyCode

`func (o *PaymentScheduleCreateReq) GetCurrencyCode() CurrencyCode`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *PaymentScheduleCreateReq) GetCurrencyCodeOk() (*CurrencyCode, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *PaymentScheduleCreateReq) SetCurrencyCode(v CurrencyCode)`

SetCurrencyCode sets CurrencyCode field to given value.


### GetDescription

`func (o *PaymentScheduleCreateReq) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PaymentScheduleCreateReq) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PaymentScheduleCreateReq) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PaymentScheduleCreateReq) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFrequency

`func (o *PaymentScheduleCreateReq) GetFrequency() PaymentScheduleFrequency`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *PaymentScheduleCreateReq) GetFrequencyOk() (*PaymentScheduleFrequency, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *PaymentScheduleCreateReq) SetFrequency(v PaymentScheduleFrequency)`

SetFrequency sets Frequency field to given value.


### GetNextPaymentImpactDate

`func (o *PaymentScheduleCreateReq) GetNextPaymentImpactDate() string`

GetNextPaymentImpactDate returns the NextPaymentImpactDate field if non-nil, zero value otherwise.

### GetNextPaymentImpactDateOk

`func (o *PaymentScheduleCreateReq) GetNextPaymentImpactDateOk() (*string, bool)`

GetNextPaymentImpactDateOk returns a tuple with the NextPaymentImpactDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPaymentImpactDate

`func (o *PaymentScheduleCreateReq) SetNextPaymentImpactDate(v string)`

SetNextPaymentImpactDate sets NextPaymentImpactDate field to given value.

### HasNextPaymentImpactDate

`func (o *PaymentScheduleCreateReq) HasNextPaymentImpactDate() bool`

HasNextPaymentImpactDate returns a boolean if a field has been set.

### GetPaymentDay

`func (o *PaymentScheduleCreateReq) GetPaymentDay() string`

GetPaymentDay returns the PaymentDay field if non-nil, zero value otherwise.

### GetPaymentDayOk

`func (o *PaymentScheduleCreateReq) GetPaymentDayOk() (*string, bool)`

GetPaymentDayOk returns a tuple with the PaymentDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentDay

`func (o *PaymentScheduleCreateReq) SetPaymentDay(v string)`

SetPaymentDay sets PaymentDay field to given value.

### HasPaymentDay

`func (o *PaymentScheduleCreateReq) HasPaymentDay() bool`

HasPaymentDay returns a boolean if a field has been set.

### GetPaymentSourceToken

`func (o *PaymentScheduleCreateReq) GetPaymentSourceToken() string`

GetPaymentSourceToken returns the PaymentSourceToken field if non-nil, zero value otherwise.

### GetPaymentSourceTokenOk

`func (o *PaymentScheduleCreateReq) GetPaymentSourceTokenOk() (*string, bool)`

GetPaymentSourceTokenOk returns a tuple with the PaymentSourceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentSourceToken

`func (o *PaymentScheduleCreateReq) SetPaymentSourceToken(v string)`

SetPaymentSourceToken sets PaymentSourceToken field to given value.


### GetToken

`func (o *PaymentScheduleCreateReq) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PaymentScheduleCreateReq) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PaymentScheduleCreateReq) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *PaymentScheduleCreateReq) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


