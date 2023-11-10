# PaymentCreateReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **float32** | Amount of the payment. | 
**CurrencyCode** | [**CurrencyCode**](CurrencyCode.md) |  | [default to CURRENCYCODE_USD]
**Description** | **string** | Description of the payment. | 
**ImpactTime** | Pointer to **time.Time** | Date and time when the payment impacts the account balance and fee calculations. | [optional] 
**Metadata** | Pointer to **string** | Customer-defined additional information about the payment (for example, a check number). | [optional] 
**Method** | **string** | Method of payment. | 
**PaymentSourceToken** | Pointer to **string** | Unique identifier of the payment source. Required for ACH payments. | [optional] 
**Token** | Pointer to **string** | Unique identifier of the payment. | [optional] 

## Methods

### NewPaymentCreateReq

`func NewPaymentCreateReq(amount float32, currencyCode CurrencyCode, description string, method string, ) *PaymentCreateReq`

NewPaymentCreateReq instantiates a new PaymentCreateReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentCreateReqWithDefaults

`func NewPaymentCreateReqWithDefaults() *PaymentCreateReq`

NewPaymentCreateReqWithDefaults instantiates a new PaymentCreateReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *PaymentCreateReq) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PaymentCreateReq) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PaymentCreateReq) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetCurrencyCode

`func (o *PaymentCreateReq) GetCurrencyCode() CurrencyCode`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *PaymentCreateReq) GetCurrencyCodeOk() (*CurrencyCode, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *PaymentCreateReq) SetCurrencyCode(v CurrencyCode)`

SetCurrencyCode sets CurrencyCode field to given value.


### GetDescription

`func (o *PaymentCreateReq) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PaymentCreateReq) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PaymentCreateReq) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetImpactTime

`func (o *PaymentCreateReq) GetImpactTime() time.Time`

GetImpactTime returns the ImpactTime field if non-nil, zero value otherwise.

### GetImpactTimeOk

`func (o *PaymentCreateReq) GetImpactTimeOk() (*time.Time, bool)`

GetImpactTimeOk returns a tuple with the ImpactTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpactTime

`func (o *PaymentCreateReq) SetImpactTime(v time.Time)`

SetImpactTime sets ImpactTime field to given value.

### HasImpactTime

`func (o *PaymentCreateReq) HasImpactTime() bool`

HasImpactTime returns a boolean if a field has been set.

### GetMetadata

`func (o *PaymentCreateReq) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PaymentCreateReq) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PaymentCreateReq) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PaymentCreateReq) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMethod

`func (o *PaymentCreateReq) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *PaymentCreateReq) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *PaymentCreateReq) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetPaymentSourceToken

`func (o *PaymentCreateReq) GetPaymentSourceToken() string`

GetPaymentSourceToken returns the PaymentSourceToken field if non-nil, zero value otherwise.

### GetPaymentSourceTokenOk

`func (o *PaymentCreateReq) GetPaymentSourceTokenOk() (*string, bool)`

GetPaymentSourceTokenOk returns a tuple with the PaymentSourceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentSourceToken

`func (o *PaymentCreateReq) SetPaymentSourceToken(v string)`

SetPaymentSourceToken sets PaymentSourceToken field to given value.

### HasPaymentSourceToken

`func (o *PaymentCreateReq) HasPaymentSourceToken() bool`

HasPaymentSourceToken returns a boolean if a field has been set.

### GetToken

`func (o *PaymentCreateReq) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PaymentCreateReq) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PaymentCreateReq) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *PaymentCreateReq) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


