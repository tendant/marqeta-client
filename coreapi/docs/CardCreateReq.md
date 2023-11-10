# CardCreateReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivationActions** | Pointer to [**ActivationActions**](ActivationActions.md) |  | [optional] 
**CardProductToken** | **string** | Unique identifier of the associated card product. | 
**ExpirationOffset** | Pointer to [**ExpirationOffset**](ExpirationOffset.md) |  | [optional] 
**NewPanFromCardToken** | Pointer to **string** | Reissues the specified card (known as the \&quot;source\&quot; card) with a new primary account number (PAN).  This field reissues a card with a new PAN from the specified source card. The source card is automatically terminated when the card is reissued with the new PAN. Use this field when reissuing a lost or stolen card. | [optional] 
**ReissuePanFromCardToken** | Pointer to **string** | Reissues the specified card (known as the \&quot;source\&quot; card).  This field reissues a card by copying the PAN and PIN from the specified source card to the newly created card. The reissued card has the same PAN and PIN as the source card but a new expiration date and CVV2 number.  *NOTE:* By default, the source card is automatically terminated when the reissued card is activated. However, if your program is configured for multiple active cards, you can prevent the source card from being automatically terminated by setting the &#x60;activation_actions.terminate_reissued_source_card&#x60; field to &#x60;false&#x60;. | [optional] 
**Token** | Pointer to **string** | Unique identifier of the credit card. | [optional] 
**TranslatePinFromCardToken** | Pointer to **string** | Copies the PIN from the specified card to the newly created card.  Both cards must belong to the same user. This field is not allowed if &#x60;reissue_pan_from_card_token&#x60; is set.  Send a &#x60;GET&#x60; request to &#x60;/cards/user/{token}&#x60; to retrieve card tokens for a particular user. | [optional] 
**UserToken** | **string** | Unique identifier of the credit cardholder. | 

## Methods

### NewCardCreateReq

`func NewCardCreateReq(cardProductToken string, userToken string, ) *CardCreateReq`

NewCardCreateReq instantiates a new CardCreateReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardCreateReqWithDefaults

`func NewCardCreateReqWithDefaults() *CardCreateReq`

NewCardCreateReqWithDefaults instantiates a new CardCreateReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivationActions

`func (o *CardCreateReq) GetActivationActions() ActivationActions`

GetActivationActions returns the ActivationActions field if non-nil, zero value otherwise.

### GetActivationActionsOk

`func (o *CardCreateReq) GetActivationActionsOk() (*ActivationActions, bool)`

GetActivationActionsOk returns a tuple with the ActivationActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationActions

`func (o *CardCreateReq) SetActivationActions(v ActivationActions)`

SetActivationActions sets ActivationActions field to given value.

### HasActivationActions

`func (o *CardCreateReq) HasActivationActions() bool`

HasActivationActions returns a boolean if a field has been set.

### GetCardProductToken

`func (o *CardCreateReq) GetCardProductToken() string`

GetCardProductToken returns the CardProductToken field if non-nil, zero value otherwise.

### GetCardProductTokenOk

`func (o *CardCreateReq) GetCardProductTokenOk() (*string, bool)`

GetCardProductTokenOk returns a tuple with the CardProductToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardProductToken

`func (o *CardCreateReq) SetCardProductToken(v string)`

SetCardProductToken sets CardProductToken field to given value.


### GetExpirationOffset

`func (o *CardCreateReq) GetExpirationOffset() ExpirationOffset`

GetExpirationOffset returns the ExpirationOffset field if non-nil, zero value otherwise.

### GetExpirationOffsetOk

`func (o *CardCreateReq) GetExpirationOffsetOk() (*ExpirationOffset, bool)`

GetExpirationOffsetOk returns a tuple with the ExpirationOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationOffset

`func (o *CardCreateReq) SetExpirationOffset(v ExpirationOffset)`

SetExpirationOffset sets ExpirationOffset field to given value.

### HasExpirationOffset

`func (o *CardCreateReq) HasExpirationOffset() bool`

HasExpirationOffset returns a boolean if a field has been set.

### GetNewPanFromCardToken

`func (o *CardCreateReq) GetNewPanFromCardToken() string`

GetNewPanFromCardToken returns the NewPanFromCardToken field if non-nil, zero value otherwise.

### GetNewPanFromCardTokenOk

`func (o *CardCreateReq) GetNewPanFromCardTokenOk() (*string, bool)`

GetNewPanFromCardTokenOk returns a tuple with the NewPanFromCardToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPanFromCardToken

`func (o *CardCreateReq) SetNewPanFromCardToken(v string)`

SetNewPanFromCardToken sets NewPanFromCardToken field to given value.

### HasNewPanFromCardToken

`func (o *CardCreateReq) HasNewPanFromCardToken() bool`

HasNewPanFromCardToken returns a boolean if a field has been set.

### GetReissuePanFromCardToken

`func (o *CardCreateReq) GetReissuePanFromCardToken() string`

GetReissuePanFromCardToken returns the ReissuePanFromCardToken field if non-nil, zero value otherwise.

### GetReissuePanFromCardTokenOk

`func (o *CardCreateReq) GetReissuePanFromCardTokenOk() (*string, bool)`

GetReissuePanFromCardTokenOk returns a tuple with the ReissuePanFromCardToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReissuePanFromCardToken

`func (o *CardCreateReq) SetReissuePanFromCardToken(v string)`

SetReissuePanFromCardToken sets ReissuePanFromCardToken field to given value.

### HasReissuePanFromCardToken

`func (o *CardCreateReq) HasReissuePanFromCardToken() bool`

HasReissuePanFromCardToken returns a boolean if a field has been set.

### GetToken

`func (o *CardCreateReq) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CardCreateReq) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CardCreateReq) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CardCreateReq) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTranslatePinFromCardToken

`func (o *CardCreateReq) GetTranslatePinFromCardToken() string`

GetTranslatePinFromCardToken returns the TranslatePinFromCardToken field if non-nil, zero value otherwise.

### GetTranslatePinFromCardTokenOk

`func (o *CardCreateReq) GetTranslatePinFromCardTokenOk() (*string, bool)`

GetTranslatePinFromCardTokenOk returns a tuple with the TranslatePinFromCardToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslatePinFromCardToken

`func (o *CardCreateReq) SetTranslatePinFromCardToken(v string)`

SetTranslatePinFromCardToken sets TranslatePinFromCardToken field to given value.

### HasTranslatePinFromCardToken

`func (o *CardCreateReq) HasTranslatePinFromCardToken() bool`

HasTranslatePinFromCardToken returns a boolean if a field has been set.

### GetUserToken

`func (o *CardCreateReq) GetUserToken() string`

GetUserToken returns the UserToken field if non-nil, zero value otherwise.

### GetUserTokenOk

`func (o *CardCreateReq) GetUserTokenOk() (*string, bool)`

GetUserTokenOk returns a tuple with the UserToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserToken

`func (o *CardCreateReq) SetUserToken(v string)`

SetUserToken sets UserToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


