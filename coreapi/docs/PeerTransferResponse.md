# PeerTransferResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **float32** | Amount of the transfer. | 
**CreatedTime** | **time.Time** |  | 
**CurrencyCode** | **string** | Three-digit ISO 4217 currency code. | 
**Memo** | Pointer to **string** | Additional descriptive text about the peer transfer. | [optional] 
**RecipientBusinessToken** | Pointer to **string** | Specifies the business account holder that receives funds. | [optional] 
**RecipientUserToken** | Pointer to **string** | Specifies the user account holder that receives funds. | [optional] 
**SenderBusinessToken** | Pointer to **string** | Specifies the business account holder that sends funds. | [optional] 
**SenderUserToken** | Pointer to **string** | Specifies the user account holder that sends funds. | [optional] 
**Tags** | Pointer to **string** | Metadata about the peer transfer. | [optional] 
**Token** | **string** | Unique identifier of the peer transfer request. | 

## Methods

### NewPeerTransferResponse

`func NewPeerTransferResponse(amount float32, createdTime time.Time, currencyCode string, token string, ) *PeerTransferResponse`

NewPeerTransferResponse instantiates a new PeerTransferResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPeerTransferResponseWithDefaults

`func NewPeerTransferResponseWithDefaults() *PeerTransferResponse`

NewPeerTransferResponseWithDefaults instantiates a new PeerTransferResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *PeerTransferResponse) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PeerTransferResponse) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PeerTransferResponse) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetCreatedTime

`func (o *PeerTransferResponse) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *PeerTransferResponse) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *PeerTransferResponse) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetCurrencyCode

`func (o *PeerTransferResponse) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *PeerTransferResponse) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *PeerTransferResponse) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.


### GetMemo

`func (o *PeerTransferResponse) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *PeerTransferResponse) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *PeerTransferResponse) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *PeerTransferResponse) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetRecipientBusinessToken

`func (o *PeerTransferResponse) GetRecipientBusinessToken() string`

GetRecipientBusinessToken returns the RecipientBusinessToken field if non-nil, zero value otherwise.

### GetRecipientBusinessTokenOk

`func (o *PeerTransferResponse) GetRecipientBusinessTokenOk() (*string, bool)`

GetRecipientBusinessTokenOk returns a tuple with the RecipientBusinessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientBusinessToken

`func (o *PeerTransferResponse) SetRecipientBusinessToken(v string)`

SetRecipientBusinessToken sets RecipientBusinessToken field to given value.

### HasRecipientBusinessToken

`func (o *PeerTransferResponse) HasRecipientBusinessToken() bool`

HasRecipientBusinessToken returns a boolean if a field has been set.

### GetRecipientUserToken

`func (o *PeerTransferResponse) GetRecipientUserToken() string`

GetRecipientUserToken returns the RecipientUserToken field if non-nil, zero value otherwise.

### GetRecipientUserTokenOk

`func (o *PeerTransferResponse) GetRecipientUserTokenOk() (*string, bool)`

GetRecipientUserTokenOk returns a tuple with the RecipientUserToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientUserToken

`func (o *PeerTransferResponse) SetRecipientUserToken(v string)`

SetRecipientUserToken sets RecipientUserToken field to given value.

### HasRecipientUserToken

`func (o *PeerTransferResponse) HasRecipientUserToken() bool`

HasRecipientUserToken returns a boolean if a field has been set.

### GetSenderBusinessToken

`func (o *PeerTransferResponse) GetSenderBusinessToken() string`

GetSenderBusinessToken returns the SenderBusinessToken field if non-nil, zero value otherwise.

### GetSenderBusinessTokenOk

`func (o *PeerTransferResponse) GetSenderBusinessTokenOk() (*string, bool)`

GetSenderBusinessTokenOk returns a tuple with the SenderBusinessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderBusinessToken

`func (o *PeerTransferResponse) SetSenderBusinessToken(v string)`

SetSenderBusinessToken sets SenderBusinessToken field to given value.

### HasSenderBusinessToken

`func (o *PeerTransferResponse) HasSenderBusinessToken() bool`

HasSenderBusinessToken returns a boolean if a field has been set.

### GetSenderUserToken

`func (o *PeerTransferResponse) GetSenderUserToken() string`

GetSenderUserToken returns the SenderUserToken field if non-nil, zero value otherwise.

### GetSenderUserTokenOk

`func (o *PeerTransferResponse) GetSenderUserTokenOk() (*string, bool)`

GetSenderUserTokenOk returns a tuple with the SenderUserToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderUserToken

`func (o *PeerTransferResponse) SetSenderUserToken(v string)`

SetSenderUserToken sets SenderUserToken field to given value.

### HasSenderUserToken

`func (o *PeerTransferResponse) HasSenderUserToken() bool`

HasSenderUserToken returns a boolean if a field has been set.

### GetTags

`func (o *PeerTransferResponse) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PeerTransferResponse) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PeerTransferResponse) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PeerTransferResponse) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetToken

`func (o *PeerTransferResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PeerTransferResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PeerTransferResponse) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


