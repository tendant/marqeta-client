# PushToCardDisburseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **float32** |  | 
**CurrencyCode** | **string** |  | 
**Memo** | Pointer to **string** |  | [optional] 
**PaymentInstrumentToken** | **string** |  | 
**SoftDescriptor** | Pointer to [**PTCSoftDescriptor**](PTCSoftDescriptor.md) |  | [optional] 
**Tags** | Pointer to **string** |  | [optional] 
**Token** | Pointer to **string** |  | [optional] 

## Methods

### NewPushToCardDisburseRequest

`func NewPushToCardDisburseRequest(amount float32, currencyCode string, paymentInstrumentToken string, ) *PushToCardDisburseRequest`

NewPushToCardDisburseRequest instantiates a new PushToCardDisburseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPushToCardDisburseRequestWithDefaults

`func NewPushToCardDisburseRequestWithDefaults() *PushToCardDisburseRequest`

NewPushToCardDisburseRequestWithDefaults instantiates a new PushToCardDisburseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *PushToCardDisburseRequest) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PushToCardDisburseRequest) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PushToCardDisburseRequest) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetCurrencyCode

`func (o *PushToCardDisburseRequest) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *PushToCardDisburseRequest) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *PushToCardDisburseRequest) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.


### GetMemo

`func (o *PushToCardDisburseRequest) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *PushToCardDisburseRequest) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *PushToCardDisburseRequest) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *PushToCardDisburseRequest) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetPaymentInstrumentToken

`func (o *PushToCardDisburseRequest) GetPaymentInstrumentToken() string`

GetPaymentInstrumentToken returns the PaymentInstrumentToken field if non-nil, zero value otherwise.

### GetPaymentInstrumentTokenOk

`func (o *PushToCardDisburseRequest) GetPaymentInstrumentTokenOk() (*string, bool)`

GetPaymentInstrumentTokenOk returns a tuple with the PaymentInstrumentToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentInstrumentToken

`func (o *PushToCardDisburseRequest) SetPaymentInstrumentToken(v string)`

SetPaymentInstrumentToken sets PaymentInstrumentToken field to given value.


### GetSoftDescriptor

`func (o *PushToCardDisburseRequest) GetSoftDescriptor() PTCSoftDescriptor`

GetSoftDescriptor returns the SoftDescriptor field if non-nil, zero value otherwise.

### GetSoftDescriptorOk

`func (o *PushToCardDisburseRequest) GetSoftDescriptorOk() (*PTCSoftDescriptor, bool)`

GetSoftDescriptorOk returns a tuple with the SoftDescriptor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftDescriptor

`func (o *PushToCardDisburseRequest) SetSoftDescriptor(v PTCSoftDescriptor)`

SetSoftDescriptor sets SoftDescriptor field to given value.

### HasSoftDescriptor

`func (o *PushToCardDisburseRequest) HasSoftDescriptor() bool`

HasSoftDescriptor returns a boolean if a field has been set.

### GetTags

`func (o *PushToCardDisburseRequest) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PushToCardDisburseRequest) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PushToCardDisburseRequest) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PushToCardDisburseRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetToken

`func (o *PushToCardDisburseRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PushToCardDisburseRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PushToCardDisburseRequest) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *PushToCardDisburseRequest) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


