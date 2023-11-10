# CardTransitionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardToken** | **string** | Identifies the card whose state will transition. | 
**Channel** | **string** | The mechanism by which the transition was initiated.  * *ADMIN* - Indicates that the card transition was initiated through the Marqeta Dashboard. * *API* - Indicates that the card transition was initiated by you through the Core API. Use this value when creating a card transition with an API &#x60;POST&#x60; request. * *FRAUD* - Indicates that either Marqeta or the card network has determined that the card is fraudulent. * *IVR* - Indicates that the card transition was initiated through your Interactive Voice Response system. * *SYSTEM* - Indicates that the card transition was initiated by Marqeta. For example, Marqeta suspended the card due to excessive failed personal identification number (PIN) entries. | 
**Reason** | Pointer to **string** | Additional information about the state change. | [optional] 
**ReasonCode** | Pointer to **string** | Standard code describing the reason for the transition.  *NOTE:* This field is required if your program uses v2 of the &#x60;user_card_state_version&#x60;, which is a program-specific configuration value that is managed by Marqeta and cannot be accessed via the API. To learn more about the &#x60;user_card_state_version&#x60; program configuration, contact your Marqeta representative.  * *00:* Object activated for the first time * *01:* Requested by you * *02:* Inactivity over time * *03:* This address cannot accept mail or the addressee is unknown * *04:* Negative account balance * *05:* Account under review * *06:* Suspicious activity was identified * *07:* Activity outside the program parameters was identified * *08:* Confirmed fraud was identified * *09:* Matched with an Office of Foreign Assets Control list * *10:* Card was reported lost * *11:* Card information was cloned * *12:* Account or card information was compromised * *13:* Temporary status change while on hold/leave * *14:* Initiated by Marqeta * *15:* Initiated by issuer * *16:* Card expired * *17:* Failed KYC * *18:* Changed to &#x60;ACTIVE&#x60; because information was properly validated * *19:* Changed to &#x60;ACTIVE&#x60; because account activity was properly validated * *20:* Change occurred prior to the normalization of reason codes * *21:* Initiated by a third party, often a digital wallet provider * *22:* PIN retry limit reached * *23:* Card was reported stolen * *24:* Address issue * *25:* Name issue * *26:* SSN issue * *27:* DOB issue * *28:* Email issue * *29:* Phone issue * *30:* Account/fulfillment mismatch * *31:* Other reason | [optional] 
**State** | **string** | Specifies the new state. | [readonly] 
**SyncStateWithDwts** | Pointer to **bool** | Set this field to &#x60;true&#x60; to synchronize the state of the card&#39;s associated token(s) with the card&#39;s new state. The digital wallet tokens must be in a valid starting state for the given transition, which will reflect the card&#39;s state transition. For example, if the card is transitioning from the &#x60;ACTIVE&#x60; state to the &#x60;SUSPENDED&#x60; state, only digital wallet tokens in the &#x60;ACTIVE&#x60; state will be synchronized with the card state transition and therefore be transitioned to the &#x60;SUSPENDED&#x60; state.  Leave this field blank or set it to &#x60;false&#x60; to keep the states of the card and its digital wallet tokens independent. | [optional] [readonly] 
**Token** | Pointer to **string** | Unique identifier of the card transition.  If you do not include a token, the system will generate one automatically. This token is referenced in other API calls, so we recommend that you define a simple string that is easy to remember. This value cannot be updated. | [optional] 
**Validations** | Pointer to [**ValidationsRequest**](ValidationsRequest.md) |  | [optional] 

## Methods

### NewCardTransitionRequest

`func NewCardTransitionRequest(cardToken string, channel string, state string, ) *CardTransitionRequest`

NewCardTransitionRequest instantiates a new CardTransitionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardTransitionRequestWithDefaults

`func NewCardTransitionRequestWithDefaults() *CardTransitionRequest`

NewCardTransitionRequestWithDefaults instantiates a new CardTransitionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardToken

`func (o *CardTransitionRequest) GetCardToken() string`

GetCardToken returns the CardToken field if non-nil, zero value otherwise.

### GetCardTokenOk

`func (o *CardTransitionRequest) GetCardTokenOk() (*string, bool)`

GetCardTokenOk returns a tuple with the CardToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardToken

`func (o *CardTransitionRequest) SetCardToken(v string)`

SetCardToken sets CardToken field to given value.


### GetChannel

`func (o *CardTransitionRequest) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *CardTransitionRequest) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *CardTransitionRequest) SetChannel(v string)`

SetChannel sets Channel field to given value.


### GetReason

`func (o *CardTransitionRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CardTransitionRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CardTransitionRequest) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *CardTransitionRequest) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetReasonCode

`func (o *CardTransitionRequest) GetReasonCode() string`

GetReasonCode returns the ReasonCode field if non-nil, zero value otherwise.

### GetReasonCodeOk

`func (o *CardTransitionRequest) GetReasonCodeOk() (*string, bool)`

GetReasonCodeOk returns a tuple with the ReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonCode

`func (o *CardTransitionRequest) SetReasonCode(v string)`

SetReasonCode sets ReasonCode field to given value.

### HasReasonCode

`func (o *CardTransitionRequest) HasReasonCode() bool`

HasReasonCode returns a boolean if a field has been set.

### GetState

`func (o *CardTransitionRequest) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CardTransitionRequest) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CardTransitionRequest) SetState(v string)`

SetState sets State field to given value.


### GetSyncStateWithDwts

`func (o *CardTransitionRequest) GetSyncStateWithDwts() bool`

GetSyncStateWithDwts returns the SyncStateWithDwts field if non-nil, zero value otherwise.

### GetSyncStateWithDwtsOk

`func (o *CardTransitionRequest) GetSyncStateWithDwtsOk() (*bool, bool)`

GetSyncStateWithDwtsOk returns a tuple with the SyncStateWithDwts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncStateWithDwts

`func (o *CardTransitionRequest) SetSyncStateWithDwts(v bool)`

SetSyncStateWithDwts sets SyncStateWithDwts field to given value.

### HasSyncStateWithDwts

`func (o *CardTransitionRequest) HasSyncStateWithDwts() bool`

HasSyncStateWithDwts returns a boolean if a field has been set.

### GetToken

`func (o *CardTransitionRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CardTransitionRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CardTransitionRequest) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CardTransitionRequest) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetValidations

`func (o *CardTransitionRequest) GetValidations() ValidationsRequest`

GetValidations returns the Validations field if non-nil, zero value otherwise.

### GetValidationsOk

`func (o *CardTransitionRequest) GetValidationsOk() (*ValidationsRequest, bool)`

GetValidationsOk returns a tuple with the Validations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidations

`func (o *CardTransitionRequest) SetValidations(v ValidationsRequest)`

SetValidations sets Validations field to given value.

### HasValidations

`func (o *CardTransitionRequest) HasValidations() bool`

HasValidations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


