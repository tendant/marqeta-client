# CardTransitionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Barcode** | **string** | The barcode printed on the card, expressed as digits. | 
**BulkIssuanceToken** | Pointer to **string** | The unique identifier of the bulk card order. | [optional] 
**Card** | Pointer to [**CardMetadata**](CardMetadata.md) |  | [optional] 
**CardProductToken** | **string** | Unique identifier of the card product. | 
**CardToken** | **string** | Unique identifier of the card. | 
**Channel** | **string** | The mechanism by which the transition was initiated.  * *ADMIN* - Indicates that the card transition was initiated through the Marqeta Dashboard. * *API* - Indicates that the card transition was initiated by you through the Core API. Use this value when creating a card transition with an API &#x60;POST&#x60; request. * *FRAUD* - Indicates that either Marqeta or the card network has determined that the card is fraudulent. * *IVR* - Indicates that the card transition was initiated through your Interactive Voice Response system. * *SYSTEM* - Indicates that the card transition was initiated by Marqeta. For example, Marqeta suspended the card due to excessive failed personal identification number (PIN) entries. | 
**CreatedTime** | Pointer to **time.Time** | Date and time when the resource was created, in UTC. | [optional] 
**Expedite** | Pointer to **bool** | A value of &#x60;true&#x60; indicates that you requested expedited processing of the card from your card fulfillment provider. | [optional] [default to false]
**Expiration** | **string** | Expiration date in &#x60;MMyy&#x60; format. | 
**ExpirationTime** | **string** | Expiration date and time in UTC format. | 
**Fulfillment** | Pointer to [**CardFulfillmentRequest**](CardFulfillmentRequest.md) |  | [optional] 
**FulfillmentStatus** | **string** | Provides status information about the card related to order and delivery.  The possible fulfillment states are:  * *ISSUED:* Initial state of all newly created/issued cards * *ORDERED:* Card ordered through card fulfillment provider * *REJECTED:* Card rejected by card fulfillment provider * *SHIPPED:* Card shipped by card fulfillment provider * *DELIVERED:* Card delivered by the card fulfillment provider. * *DIGITALLY_PRESENTED:* Card digitally presented using the &#x60;/cards/{token}/showpan&#x60; endpoint; does not affect the delivery of physical cards | 
**LastFour** | **string** | Last four digits of the card primary account number (PAN). | 
**NewPanFromCardToken** | Pointer to **string** | Reissues the specified (\&quot;source\&quot;) card with a new primary account number (PAN). | [optional] 
**Pan** | **string** | Primary account number (PAN) of the card. | 
**PinIsSet** | **bool** | Specifies if the personal identification number (PIN) has been set for the card. | [default to false]
**Reason** | Pointer to **string** | Additional information about the state change. | [optional] 
**ReasonCode** | Pointer to **string** | A standard code describing the reason for the transition:  * *00:* Object activated for the first time * *01:* Requested by you * *02:* Inactivity over time * *03:* This address cannot accept mail or the addressee is unknown * *04:* Negative account balance * *05:* Account under review * *06:* Suspicious activity was identified * *07:* Activity outside the program parameters was identified * *08:* Confirmed fraud was identified * *09:* Matched with an Office of Foreign Assets Control list * *10:* Card was reported lost * *11:* Card information was cloned * *12:* Account or card information was compromised * *13:* Temporary status change while on hold/leave * *14:* Initiated by Marqeta * *15:* Initiated by issuer * *16:* Card expired * *17:* Failed KYC * *18:* Changed to &#x60;ACTIVE&#x60; because information was properly validated * *19:* Changed to &#x60;ACTIVE&#x60; because account activity was properly validated * *20:* Change occurred prior to the normalization of reason codes * *21:* Initiated by a third party, often a digital wallet provider * *22:* PIN retry limit reached * *23:* Card was reported stolen * *24:* Address issue * *25:* Name issue * *26:* SSN issue * *27:* DOB issue * *28:* Email issue * *29:* Phone issue * *30:* Account/fulfillment mismatch * *31:* Other reason | [optional] 
**ReissuePanFromCardToken** | Pointer to **string** | Reissues the specified (\&quot;source\&quot;) card. | [optional] 
**State** | **string** | Indicates the state of the card. | 
**Token** | **string** | Unique identifier of the card transition. | 
**Type** | **string** | This field cannot be set directly using the &#x60;/cardtransitions&#x60; endpoint. A card transition&#39;s &#x60;type&#x60; is managed by the Marqeta platform, based on the before and after state of the transition, as specified in the request&#39;s &#x60;state&#x60; field.  This field appears only when populated by the card fulfillment provider. The &#x60;type&#x60; field&#39;s possible values are:  * *fulfillment.delivered:* Card was delivered by the card fulfillment provider * *fulfillment.digitally_presented:* Card was digitally presented using the &#x60;/cards/{token}/showpan&#x60; endpoint; does not affect the delivery of physical cards * *fulfillment.issued:* Card was created/issued * *fulfillment.ordered:* Card was ordered from the card fulfillment provider * *fulfillment.rejected:* Card was rejected by the card fulfillment provider * *fulfillment.shipped:* Card was shipped by the card fulfillment provider * *state.activated:* Card was activated * *state.limited:* Card was limited * *state.reinstated:* Card was reinstated from a suspended state * *state.suspended:* Card was suspended * *state.terminated:* Card was terminated  //Also appears in /core-api/event-types#_card_transition_events | [readonly] 
**User** | Pointer to [**CardholderMetadata**](CardholderMetadata.md) |  | [optional] 
**UserToken** | **string** | Unique identifier of the cardholder. | 
**Validations** | Pointer to [**ValidationsResponse**](ValidationsResponse.md) |  | [optional] 

## Methods

### NewCardTransitionResponse

`func NewCardTransitionResponse(barcode string, cardProductToken string, cardToken string, channel string, expiration string, expirationTime string, fulfillmentStatus string, lastFour string, pan string, pinIsSet bool, state string, token string, type_ string, userToken string, ) *CardTransitionResponse`

NewCardTransitionResponse instantiates a new CardTransitionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardTransitionResponseWithDefaults

`func NewCardTransitionResponseWithDefaults() *CardTransitionResponse`

NewCardTransitionResponseWithDefaults instantiates a new CardTransitionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBarcode

`func (o *CardTransitionResponse) GetBarcode() string`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *CardTransitionResponse) GetBarcodeOk() (*string, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode

`func (o *CardTransitionResponse) SetBarcode(v string)`

SetBarcode sets Barcode field to given value.


### GetBulkIssuanceToken

`func (o *CardTransitionResponse) GetBulkIssuanceToken() string`

GetBulkIssuanceToken returns the BulkIssuanceToken field if non-nil, zero value otherwise.

### GetBulkIssuanceTokenOk

`func (o *CardTransitionResponse) GetBulkIssuanceTokenOk() (*string, bool)`

GetBulkIssuanceTokenOk returns a tuple with the BulkIssuanceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulkIssuanceToken

`func (o *CardTransitionResponse) SetBulkIssuanceToken(v string)`

SetBulkIssuanceToken sets BulkIssuanceToken field to given value.

### HasBulkIssuanceToken

`func (o *CardTransitionResponse) HasBulkIssuanceToken() bool`

HasBulkIssuanceToken returns a boolean if a field has been set.

### GetCard

`func (o *CardTransitionResponse) GetCard() CardMetadata`

GetCard returns the Card field if non-nil, zero value otherwise.

### GetCardOk

`func (o *CardTransitionResponse) GetCardOk() (*CardMetadata, bool)`

GetCardOk returns a tuple with the Card field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCard

`func (o *CardTransitionResponse) SetCard(v CardMetadata)`

SetCard sets Card field to given value.

### HasCard

`func (o *CardTransitionResponse) HasCard() bool`

HasCard returns a boolean if a field has been set.

### GetCardProductToken

`func (o *CardTransitionResponse) GetCardProductToken() string`

GetCardProductToken returns the CardProductToken field if non-nil, zero value otherwise.

### GetCardProductTokenOk

`func (o *CardTransitionResponse) GetCardProductTokenOk() (*string, bool)`

GetCardProductTokenOk returns a tuple with the CardProductToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardProductToken

`func (o *CardTransitionResponse) SetCardProductToken(v string)`

SetCardProductToken sets CardProductToken field to given value.


### GetCardToken

`func (o *CardTransitionResponse) GetCardToken() string`

GetCardToken returns the CardToken field if non-nil, zero value otherwise.

### GetCardTokenOk

`func (o *CardTransitionResponse) GetCardTokenOk() (*string, bool)`

GetCardTokenOk returns a tuple with the CardToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardToken

`func (o *CardTransitionResponse) SetCardToken(v string)`

SetCardToken sets CardToken field to given value.


### GetChannel

`func (o *CardTransitionResponse) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *CardTransitionResponse) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *CardTransitionResponse) SetChannel(v string)`

SetChannel sets Channel field to given value.


### GetCreatedTime

`func (o *CardTransitionResponse) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *CardTransitionResponse) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *CardTransitionResponse) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *CardTransitionResponse) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetExpedite

`func (o *CardTransitionResponse) GetExpedite() bool`

GetExpedite returns the Expedite field if non-nil, zero value otherwise.

### GetExpediteOk

`func (o *CardTransitionResponse) GetExpediteOk() (*bool, bool)`

GetExpediteOk returns a tuple with the Expedite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpedite

`func (o *CardTransitionResponse) SetExpedite(v bool)`

SetExpedite sets Expedite field to given value.

### HasExpedite

`func (o *CardTransitionResponse) HasExpedite() bool`

HasExpedite returns a boolean if a field has been set.

### GetExpiration

`func (o *CardTransitionResponse) GetExpiration() string`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *CardTransitionResponse) GetExpirationOk() (*string, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *CardTransitionResponse) SetExpiration(v string)`

SetExpiration sets Expiration field to given value.


### GetExpirationTime

`func (o *CardTransitionResponse) GetExpirationTime() string`

GetExpirationTime returns the ExpirationTime field if non-nil, zero value otherwise.

### GetExpirationTimeOk

`func (o *CardTransitionResponse) GetExpirationTimeOk() (*string, bool)`

GetExpirationTimeOk returns a tuple with the ExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTime

`func (o *CardTransitionResponse) SetExpirationTime(v string)`

SetExpirationTime sets ExpirationTime field to given value.


### GetFulfillment

`func (o *CardTransitionResponse) GetFulfillment() CardFulfillmentRequest`

GetFulfillment returns the Fulfillment field if non-nil, zero value otherwise.

### GetFulfillmentOk

`func (o *CardTransitionResponse) GetFulfillmentOk() (*CardFulfillmentRequest, bool)`

GetFulfillmentOk returns a tuple with the Fulfillment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillment

`func (o *CardTransitionResponse) SetFulfillment(v CardFulfillmentRequest)`

SetFulfillment sets Fulfillment field to given value.

### HasFulfillment

`func (o *CardTransitionResponse) HasFulfillment() bool`

HasFulfillment returns a boolean if a field has been set.

### GetFulfillmentStatus

`func (o *CardTransitionResponse) GetFulfillmentStatus() string`

GetFulfillmentStatus returns the FulfillmentStatus field if non-nil, zero value otherwise.

### GetFulfillmentStatusOk

`func (o *CardTransitionResponse) GetFulfillmentStatusOk() (*string, bool)`

GetFulfillmentStatusOk returns a tuple with the FulfillmentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentStatus

`func (o *CardTransitionResponse) SetFulfillmentStatus(v string)`

SetFulfillmentStatus sets FulfillmentStatus field to given value.


### GetLastFour

`func (o *CardTransitionResponse) GetLastFour() string`

GetLastFour returns the LastFour field if non-nil, zero value otherwise.

### GetLastFourOk

`func (o *CardTransitionResponse) GetLastFourOk() (*string, bool)`

GetLastFourOk returns a tuple with the LastFour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFour

`func (o *CardTransitionResponse) SetLastFour(v string)`

SetLastFour sets LastFour field to given value.


### GetNewPanFromCardToken

`func (o *CardTransitionResponse) GetNewPanFromCardToken() string`

GetNewPanFromCardToken returns the NewPanFromCardToken field if non-nil, zero value otherwise.

### GetNewPanFromCardTokenOk

`func (o *CardTransitionResponse) GetNewPanFromCardTokenOk() (*string, bool)`

GetNewPanFromCardTokenOk returns a tuple with the NewPanFromCardToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPanFromCardToken

`func (o *CardTransitionResponse) SetNewPanFromCardToken(v string)`

SetNewPanFromCardToken sets NewPanFromCardToken field to given value.

### HasNewPanFromCardToken

`func (o *CardTransitionResponse) HasNewPanFromCardToken() bool`

HasNewPanFromCardToken returns a boolean if a field has been set.

### GetPan

`func (o *CardTransitionResponse) GetPan() string`

GetPan returns the Pan field if non-nil, zero value otherwise.

### GetPanOk

`func (o *CardTransitionResponse) GetPanOk() (*string, bool)`

GetPanOk returns a tuple with the Pan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPan

`func (o *CardTransitionResponse) SetPan(v string)`

SetPan sets Pan field to given value.


### GetPinIsSet

`func (o *CardTransitionResponse) GetPinIsSet() bool`

GetPinIsSet returns the PinIsSet field if non-nil, zero value otherwise.

### GetPinIsSetOk

`func (o *CardTransitionResponse) GetPinIsSetOk() (*bool, bool)`

GetPinIsSetOk returns a tuple with the PinIsSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinIsSet

`func (o *CardTransitionResponse) SetPinIsSet(v bool)`

SetPinIsSet sets PinIsSet field to given value.


### GetReason

`func (o *CardTransitionResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CardTransitionResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CardTransitionResponse) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *CardTransitionResponse) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetReasonCode

`func (o *CardTransitionResponse) GetReasonCode() string`

GetReasonCode returns the ReasonCode field if non-nil, zero value otherwise.

### GetReasonCodeOk

`func (o *CardTransitionResponse) GetReasonCodeOk() (*string, bool)`

GetReasonCodeOk returns a tuple with the ReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonCode

`func (o *CardTransitionResponse) SetReasonCode(v string)`

SetReasonCode sets ReasonCode field to given value.

### HasReasonCode

`func (o *CardTransitionResponse) HasReasonCode() bool`

HasReasonCode returns a boolean if a field has been set.

### GetReissuePanFromCardToken

`func (o *CardTransitionResponse) GetReissuePanFromCardToken() string`

GetReissuePanFromCardToken returns the ReissuePanFromCardToken field if non-nil, zero value otherwise.

### GetReissuePanFromCardTokenOk

`func (o *CardTransitionResponse) GetReissuePanFromCardTokenOk() (*string, bool)`

GetReissuePanFromCardTokenOk returns a tuple with the ReissuePanFromCardToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReissuePanFromCardToken

`func (o *CardTransitionResponse) SetReissuePanFromCardToken(v string)`

SetReissuePanFromCardToken sets ReissuePanFromCardToken field to given value.

### HasReissuePanFromCardToken

`func (o *CardTransitionResponse) HasReissuePanFromCardToken() bool`

HasReissuePanFromCardToken returns a boolean if a field has been set.

### GetState

`func (o *CardTransitionResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CardTransitionResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CardTransitionResponse) SetState(v string)`

SetState sets State field to given value.


### GetToken

`func (o *CardTransitionResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CardTransitionResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CardTransitionResponse) SetToken(v string)`

SetToken sets Token field to given value.


### GetType

`func (o *CardTransitionResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CardTransitionResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CardTransitionResponse) SetType(v string)`

SetType sets Type field to given value.


### GetUser

`func (o *CardTransitionResponse) GetUser() CardholderMetadata`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *CardTransitionResponse) GetUserOk() (*CardholderMetadata, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *CardTransitionResponse) SetUser(v CardholderMetadata)`

SetUser sets User field to given value.

### HasUser

`func (o *CardTransitionResponse) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetUserToken

`func (o *CardTransitionResponse) GetUserToken() string`

GetUserToken returns the UserToken field if non-nil, zero value otherwise.

### GetUserTokenOk

`func (o *CardTransitionResponse) GetUserTokenOk() (*string, bool)`

GetUserTokenOk returns a tuple with the UserToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserToken

`func (o *CardTransitionResponse) SetUserToken(v string)`

SetUserToken sets UserToken field to given value.


### GetValidations

`func (o *CardTransitionResponse) GetValidations() ValidationsResponse`

GetValidations returns the Validations field if non-nil, zero value otherwise.

### GetValidationsOk

`func (o *CardTransitionResponse) GetValidationsOk() (*ValidationsResponse, bool)`

GetValidationsOk returns a tuple with the Validations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidations

`func (o *CardTransitionResponse) SetValidations(v ValidationsResponse)`

SetValidations sets Validations field to given value.

### HasValidations

`func (o *CardTransitionResponse) HasValidations() bool`

HasValidations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


