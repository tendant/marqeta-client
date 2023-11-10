# PushToCardDisbursementResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **float32** |  | [optional] 
**CreatedTime** | **time.Time** | yyyy-MM-ddTHH:mm:ssZ | 
**CurrencyCode** | Pointer to **string** |  | [optional] 
**LastModifiedTime** | **time.Time** | yyyy-MM-ddTHH:mm:ssZ | 
**Memo** | Pointer to **string** |  | [optional] 
**PaymentInstrumentToken** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **string** |  | [optional] 
**Token** | Pointer to **string** |  | [optional] 

## Methods

### NewPushToCardDisbursementResponse

`func NewPushToCardDisbursementResponse(createdTime time.Time, lastModifiedTime time.Time, ) *PushToCardDisbursementResponse`

NewPushToCardDisbursementResponse instantiates a new PushToCardDisbursementResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPushToCardDisbursementResponseWithDefaults

`func NewPushToCardDisbursementResponseWithDefaults() *PushToCardDisbursementResponse`

NewPushToCardDisbursementResponseWithDefaults instantiates a new PushToCardDisbursementResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *PushToCardDisbursementResponse) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PushToCardDisbursementResponse) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PushToCardDisbursementResponse) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *PushToCardDisbursementResponse) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCreatedTime

`func (o *PushToCardDisbursementResponse) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *PushToCardDisbursementResponse) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *PushToCardDisbursementResponse) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetCurrencyCode

`func (o *PushToCardDisbursementResponse) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *PushToCardDisbursementResponse) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *PushToCardDisbursementResponse) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *PushToCardDisbursementResponse) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetLastModifiedTime

`func (o *PushToCardDisbursementResponse) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *PushToCardDisbursementResponse) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *PushToCardDisbursementResponse) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.


### GetMemo

`func (o *PushToCardDisbursementResponse) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *PushToCardDisbursementResponse) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *PushToCardDisbursementResponse) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *PushToCardDisbursementResponse) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetPaymentInstrumentToken

`func (o *PushToCardDisbursementResponse) GetPaymentInstrumentToken() string`

GetPaymentInstrumentToken returns the PaymentInstrumentToken field if non-nil, zero value otherwise.

### GetPaymentInstrumentTokenOk

`func (o *PushToCardDisbursementResponse) GetPaymentInstrumentTokenOk() (*string, bool)`

GetPaymentInstrumentTokenOk returns a tuple with the PaymentInstrumentToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentInstrumentToken

`func (o *PushToCardDisbursementResponse) SetPaymentInstrumentToken(v string)`

SetPaymentInstrumentToken sets PaymentInstrumentToken field to given value.

### HasPaymentInstrumentToken

`func (o *PushToCardDisbursementResponse) HasPaymentInstrumentToken() bool`

HasPaymentInstrumentToken returns a boolean if a field has been set.

### GetStatus

`func (o *PushToCardDisbursementResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PushToCardDisbursementResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PushToCardDisbursementResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PushToCardDisbursementResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *PushToCardDisbursementResponse) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PushToCardDisbursementResponse) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PushToCardDisbursementResponse) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PushToCardDisbursementResponse) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetToken

`func (o *PushToCardDisbursementResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PushToCardDisbursementResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PushToCardDisbursementResponse) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *PushToCardDisbursementResponse) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


