# ActivationActions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SwapDigitalWalletTokensFromCardToken** | Pointer to **string** | Moves all digital wallet tokens from the specified card to the new card.  Not relevant when &#x60;reissue_pan_from_card_token&#x60; is set.  Send a &#x60;GET&#x60; request to &#x60;/cards/user/{token}&#x60; to retrieve card tokens for a particular user. | [optional] 
**TerminateReissuedSourceCard** | Pointer to **bool** | If you are reissuing a card, the source card is terminated by default. To prevent the source card from being terminated, set this field to &#x60;false&#x60;.  Only relevant when &#x60;reissue_pan_from_card_token&#x60; is set. | [optional] [default to true]

## Methods

### NewActivationActions

`func NewActivationActions() *ActivationActions`

NewActivationActions instantiates a new ActivationActions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivationActionsWithDefaults

`func NewActivationActionsWithDefaults() *ActivationActions`

NewActivationActionsWithDefaults instantiates a new ActivationActions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSwapDigitalWalletTokensFromCardToken

`func (o *ActivationActions) GetSwapDigitalWalletTokensFromCardToken() string`

GetSwapDigitalWalletTokensFromCardToken returns the SwapDigitalWalletTokensFromCardToken field if non-nil, zero value otherwise.

### GetSwapDigitalWalletTokensFromCardTokenOk

`func (o *ActivationActions) GetSwapDigitalWalletTokensFromCardTokenOk() (*string, bool)`

GetSwapDigitalWalletTokensFromCardTokenOk returns a tuple with the SwapDigitalWalletTokensFromCardToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwapDigitalWalletTokensFromCardToken

`func (o *ActivationActions) SetSwapDigitalWalletTokensFromCardToken(v string)`

SetSwapDigitalWalletTokensFromCardToken sets SwapDigitalWalletTokensFromCardToken field to given value.

### HasSwapDigitalWalletTokensFromCardToken

`func (o *ActivationActions) HasSwapDigitalWalletTokensFromCardToken() bool`

HasSwapDigitalWalletTokensFromCardToken returns a boolean if a field has been set.

### GetTerminateReissuedSourceCard

`func (o *ActivationActions) GetTerminateReissuedSourceCard() bool`

GetTerminateReissuedSourceCard returns the TerminateReissuedSourceCard field if non-nil, zero value otherwise.

### GetTerminateReissuedSourceCardOk

`func (o *ActivationActions) GetTerminateReissuedSourceCardOk() (*bool, bool)`

GetTerminateReissuedSourceCardOk returns a tuple with the TerminateReissuedSourceCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminateReissuedSourceCard

`func (o *ActivationActions) SetTerminateReissuedSourceCard(v bool)`

SetTerminateReissuedSourceCard sets TerminateReissuedSourceCard field to given value.

### HasTerminateReissuedSourceCard

`func (o *ActivationActions) HasTerminateReissuedSourceCard() bool`

HasTerminateReissuedSourceCard returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


