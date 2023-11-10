# CardRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivationActions** | Pointer to [**ActivationActions**](ActivationActions.md) |  | [optional] 
**BulkIssuanceToken** | Pointer to **string** | Associates the card with the specified bulk card order. This field cannot be updated. | [optional] 
**CardProductToken** | **string** | Unique identifier of the card product. | 
**Expedite** | Pointer to **bool** | Set to &#x60;true&#x60; to request expedited processing of the card by your card fulfillment provider.  This expedited service is available for cards fulfilled by link:http://perfectplastic.com/[Perfect Plastic Printing, window&#x3D;\&quot;_blank\&quot;], link:http://www.idemia.com[IDEMIA, window&#x3D;\&quot;_blank\&quot;], and link:https://www.arroweye.com/[Arroweye Solutions, window&#x3D;\&quot;_blank\&quot;].  *NOTE:* Contact your Marqeta representative for information regarding the cost of expedited service. | [optional] [default to false]
**ExpirationOffset** | Pointer to [**ExpirationOffset**](ExpirationOffset.md) |  | [optional] 
**Fulfillment** | Pointer to [**CardFulfillmentRequest**](CardFulfillmentRequest.md) |  | [optional] 
**Metadata** | Pointer to **map[string]string** | Associates customer-provided metadata with the card. | [optional] 
**NewPanFromCardToken** | Pointer to **string** | Reissues the specified card (known as the \&quot;source\&quot; card) with a new primary account number (PAN).  This field reissues a card with a new PAN from the specified source card. The source card is automatically terminated when the card is reissued with the new PAN. Use this field when reissuing a lost or stolen card.  Send a &#x60;GET&#x60; request to &#x60;/cards/user/{token}&#x60; to retrieve card tokens for a particular user. | [optional] 
**ReissuePanFromCardToken** | Pointer to **string** | Reissues the specified card (known as the \&quot;source\&quot; card).  This field reissues a card by copying the primary account number (PAN) and personal identification number (PIN) from the specified source card to the newly created card. The reissued card has the same PAN and PIN as the source card but a new expiration date and CVV2 number.  Send a &#x60;GET&#x60; request to &#x60;/cards/user/{token}&#x60; to retrieve card tokens for a particular user.  *NOTE:* By default, the source card is automatically terminated when the reissued card is activated. However, if your program is configured for multiple active cards, you can prevent the source card from being automatically terminated by setting the &#x60;activation_actions.terminate_reissued_source_card&#x60; field to &#x60;false&#x60;. | [optional] 
**Token** | Pointer to **string** | Unique identifier of the card.  If you do not include a token, the system will generate one automatically. Other API calls will require this token, so we recommend creating a token that is easy to remember rather than letting the system generate one. This value cannot be updated. | [optional] 
**TranslatePinFromCardToken** | Pointer to **string** | Copies the PIN from the specified card to the newly created card.  Both cards must belong to the same user. Populating this field will raise an error if &#x60;reissue_pan_from_card_token&#x60; is also set.  Send a &#x60;GET&#x60; request to &#x60;/cards/user/{token}&#x60; to retrieve card tokens for a particular user. | [optional] 
**UserToken** | **string** | Unique identifier of the authorized user of the card. | 

## Methods

### NewCardRequest

`func NewCardRequest(cardProductToken string, userToken string, ) *CardRequest`

NewCardRequest instantiates a new CardRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardRequestWithDefaults

`func NewCardRequestWithDefaults() *CardRequest`

NewCardRequestWithDefaults instantiates a new CardRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivationActions

`func (o *CardRequest) GetActivationActions() ActivationActions`

GetActivationActions returns the ActivationActions field if non-nil, zero value otherwise.

### GetActivationActionsOk

`func (o *CardRequest) GetActivationActionsOk() (*ActivationActions, bool)`

GetActivationActionsOk returns a tuple with the ActivationActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationActions

`func (o *CardRequest) SetActivationActions(v ActivationActions)`

SetActivationActions sets ActivationActions field to given value.

### HasActivationActions

`func (o *CardRequest) HasActivationActions() bool`

HasActivationActions returns a boolean if a field has been set.

### GetBulkIssuanceToken

`func (o *CardRequest) GetBulkIssuanceToken() string`

GetBulkIssuanceToken returns the BulkIssuanceToken field if non-nil, zero value otherwise.

### GetBulkIssuanceTokenOk

`func (o *CardRequest) GetBulkIssuanceTokenOk() (*string, bool)`

GetBulkIssuanceTokenOk returns a tuple with the BulkIssuanceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulkIssuanceToken

`func (o *CardRequest) SetBulkIssuanceToken(v string)`

SetBulkIssuanceToken sets BulkIssuanceToken field to given value.

### HasBulkIssuanceToken

`func (o *CardRequest) HasBulkIssuanceToken() bool`

HasBulkIssuanceToken returns a boolean if a field has been set.

### GetCardProductToken

`func (o *CardRequest) GetCardProductToken() string`

GetCardProductToken returns the CardProductToken field if non-nil, zero value otherwise.

### GetCardProductTokenOk

`func (o *CardRequest) GetCardProductTokenOk() (*string, bool)`

GetCardProductTokenOk returns a tuple with the CardProductToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardProductToken

`func (o *CardRequest) SetCardProductToken(v string)`

SetCardProductToken sets CardProductToken field to given value.


### GetExpedite

`func (o *CardRequest) GetExpedite() bool`

GetExpedite returns the Expedite field if non-nil, zero value otherwise.

### GetExpediteOk

`func (o *CardRequest) GetExpediteOk() (*bool, bool)`

GetExpediteOk returns a tuple with the Expedite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpedite

`func (o *CardRequest) SetExpedite(v bool)`

SetExpedite sets Expedite field to given value.

### HasExpedite

`func (o *CardRequest) HasExpedite() bool`

HasExpedite returns a boolean if a field has been set.

### GetExpirationOffset

`func (o *CardRequest) GetExpirationOffset() ExpirationOffset`

GetExpirationOffset returns the ExpirationOffset field if non-nil, zero value otherwise.

### GetExpirationOffsetOk

`func (o *CardRequest) GetExpirationOffsetOk() (*ExpirationOffset, bool)`

GetExpirationOffsetOk returns a tuple with the ExpirationOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationOffset

`func (o *CardRequest) SetExpirationOffset(v ExpirationOffset)`

SetExpirationOffset sets ExpirationOffset field to given value.

### HasExpirationOffset

`func (o *CardRequest) HasExpirationOffset() bool`

HasExpirationOffset returns a boolean if a field has been set.

### GetFulfillment

`func (o *CardRequest) GetFulfillment() CardFulfillmentRequest`

GetFulfillment returns the Fulfillment field if non-nil, zero value otherwise.

### GetFulfillmentOk

`func (o *CardRequest) GetFulfillmentOk() (*CardFulfillmentRequest, bool)`

GetFulfillmentOk returns a tuple with the Fulfillment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillment

`func (o *CardRequest) SetFulfillment(v CardFulfillmentRequest)`

SetFulfillment sets Fulfillment field to given value.

### HasFulfillment

`func (o *CardRequest) HasFulfillment() bool`

HasFulfillment returns a boolean if a field has been set.

### GetMetadata

`func (o *CardRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CardRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CardRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CardRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetNewPanFromCardToken

`func (o *CardRequest) GetNewPanFromCardToken() string`

GetNewPanFromCardToken returns the NewPanFromCardToken field if non-nil, zero value otherwise.

### GetNewPanFromCardTokenOk

`func (o *CardRequest) GetNewPanFromCardTokenOk() (*string, bool)`

GetNewPanFromCardTokenOk returns a tuple with the NewPanFromCardToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPanFromCardToken

`func (o *CardRequest) SetNewPanFromCardToken(v string)`

SetNewPanFromCardToken sets NewPanFromCardToken field to given value.

### HasNewPanFromCardToken

`func (o *CardRequest) HasNewPanFromCardToken() bool`

HasNewPanFromCardToken returns a boolean if a field has been set.

### GetReissuePanFromCardToken

`func (o *CardRequest) GetReissuePanFromCardToken() string`

GetReissuePanFromCardToken returns the ReissuePanFromCardToken field if non-nil, zero value otherwise.

### GetReissuePanFromCardTokenOk

`func (o *CardRequest) GetReissuePanFromCardTokenOk() (*string, bool)`

GetReissuePanFromCardTokenOk returns a tuple with the ReissuePanFromCardToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReissuePanFromCardToken

`func (o *CardRequest) SetReissuePanFromCardToken(v string)`

SetReissuePanFromCardToken sets ReissuePanFromCardToken field to given value.

### HasReissuePanFromCardToken

`func (o *CardRequest) HasReissuePanFromCardToken() bool`

HasReissuePanFromCardToken returns a boolean if a field has been set.

### GetToken

`func (o *CardRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CardRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CardRequest) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CardRequest) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTranslatePinFromCardToken

`func (o *CardRequest) GetTranslatePinFromCardToken() string`

GetTranslatePinFromCardToken returns the TranslatePinFromCardToken field if non-nil, zero value otherwise.

### GetTranslatePinFromCardTokenOk

`func (o *CardRequest) GetTranslatePinFromCardTokenOk() (*string, bool)`

GetTranslatePinFromCardTokenOk returns a tuple with the TranslatePinFromCardToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslatePinFromCardToken

`func (o *CardRequest) SetTranslatePinFromCardToken(v string)`

SetTranslatePinFromCardToken sets TranslatePinFromCardToken field to given value.

### HasTranslatePinFromCardToken

`func (o *CardRequest) HasTranslatePinFromCardToken() bool`

HasTranslatePinFromCardToken returns a boolean if a field has been set.

### GetUserToken

`func (o *CardRequest) GetUserToken() string`

GetUserToken returns the UserToken field if non-nil, zero value otherwise.

### GetUserTokenOk

`func (o *CardRequest) GetUserTokenOk() (*string, bool)`

GetUserTokenOk returns a tuple with the UserToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserToken

`func (o *CardRequest) SetUserToken(v string)`

SetUserToken sets UserToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


