# PeerTransferRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **float32** | Amount of the transfer. | 
**CurrencyCode** | **string** | Three-digit ISO 4217 currency code. | 
**Memo** | Pointer to **string** | Additional descriptive text about the transfer. | [optional] 
**RecipientBusinessToken** | Pointer to **string** | Specifies the business account holder that receives funds.  Send a &#x60;GET&#x60; request to &#x60;/businesses&#x60; to retrieve business tokens. | [optional] 
**RecipientUserToken** | Pointer to **string** | Specifies the user account holder that receives funds.  Send a &#x60;GET&#x60; request to &#x60;/users&#x60; to retrieve user tokens. | [optional] 
**SenderBusinessToken** | Pointer to **string** | Specifies the business account holder that sends funds.  Send a &#x60;GET&#x60; request to &#x60;/businesses&#x60; to retrieve business tokens. | [optional] 
**SenderUserToken** | Pointer to **string** | Specifies the user account holder that sends funds.  Send a &#x60;GET&#x60; request to &#x60;/users&#x60; to retrieve user tokens. | [optional] 
**Tags** | Pointer to **string** | Metadata about the peer transfer. | [optional] 
**Token** | Pointer to **string** | Unique identifier of the peer transfer request.  If you do not include a token, the system will generate one automatically. This token is necessary for use in other API calls, so we recommend that rather than let the system generate one, you use a simple string that is easy to remember. This value cannot be updated. | [optional] 

## Methods

### NewPeerTransferRequest

`func NewPeerTransferRequest(amount float32, currencyCode string, ) *PeerTransferRequest`

NewPeerTransferRequest instantiates a new PeerTransferRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPeerTransferRequestWithDefaults

`func NewPeerTransferRequestWithDefaults() *PeerTransferRequest`

NewPeerTransferRequestWithDefaults instantiates a new PeerTransferRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *PeerTransferRequest) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PeerTransferRequest) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PeerTransferRequest) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetCurrencyCode

`func (o *PeerTransferRequest) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *PeerTransferRequest) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *PeerTransferRequest) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.


### GetMemo

`func (o *PeerTransferRequest) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *PeerTransferRequest) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *PeerTransferRequest) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *PeerTransferRequest) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetRecipientBusinessToken

`func (o *PeerTransferRequest) GetRecipientBusinessToken() string`

GetRecipientBusinessToken returns the RecipientBusinessToken field if non-nil, zero value otherwise.

### GetRecipientBusinessTokenOk

`func (o *PeerTransferRequest) GetRecipientBusinessTokenOk() (*string, bool)`

GetRecipientBusinessTokenOk returns a tuple with the RecipientBusinessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientBusinessToken

`func (o *PeerTransferRequest) SetRecipientBusinessToken(v string)`

SetRecipientBusinessToken sets RecipientBusinessToken field to given value.

### HasRecipientBusinessToken

`func (o *PeerTransferRequest) HasRecipientBusinessToken() bool`

HasRecipientBusinessToken returns a boolean if a field has been set.

### GetRecipientUserToken

`func (o *PeerTransferRequest) GetRecipientUserToken() string`

GetRecipientUserToken returns the RecipientUserToken field if non-nil, zero value otherwise.

### GetRecipientUserTokenOk

`func (o *PeerTransferRequest) GetRecipientUserTokenOk() (*string, bool)`

GetRecipientUserTokenOk returns a tuple with the RecipientUserToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientUserToken

`func (o *PeerTransferRequest) SetRecipientUserToken(v string)`

SetRecipientUserToken sets RecipientUserToken field to given value.

### HasRecipientUserToken

`func (o *PeerTransferRequest) HasRecipientUserToken() bool`

HasRecipientUserToken returns a boolean if a field has been set.

### GetSenderBusinessToken

`func (o *PeerTransferRequest) GetSenderBusinessToken() string`

GetSenderBusinessToken returns the SenderBusinessToken field if non-nil, zero value otherwise.

### GetSenderBusinessTokenOk

`func (o *PeerTransferRequest) GetSenderBusinessTokenOk() (*string, bool)`

GetSenderBusinessTokenOk returns a tuple with the SenderBusinessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderBusinessToken

`func (o *PeerTransferRequest) SetSenderBusinessToken(v string)`

SetSenderBusinessToken sets SenderBusinessToken field to given value.

### HasSenderBusinessToken

`func (o *PeerTransferRequest) HasSenderBusinessToken() bool`

HasSenderBusinessToken returns a boolean if a field has been set.

### GetSenderUserToken

`func (o *PeerTransferRequest) GetSenderUserToken() string`

GetSenderUserToken returns the SenderUserToken field if non-nil, zero value otherwise.

### GetSenderUserTokenOk

`func (o *PeerTransferRequest) GetSenderUserTokenOk() (*string, bool)`

GetSenderUserTokenOk returns a tuple with the SenderUserToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderUserToken

`func (o *PeerTransferRequest) SetSenderUserToken(v string)`

SetSenderUserToken sets SenderUserToken field to given value.

### HasSenderUserToken

`func (o *PeerTransferRequest) HasSenderUserToken() bool`

HasSenderUserToken returns a boolean if a field has been set.

### GetTags

`func (o *PeerTransferRequest) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PeerTransferRequest) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PeerTransferRequest) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PeerTransferRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetToken

`func (o *PeerTransferRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PeerTransferRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PeerTransferRequest) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *PeerTransferRequest) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


