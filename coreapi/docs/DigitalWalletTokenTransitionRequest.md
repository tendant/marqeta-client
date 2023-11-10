# DigitalWalletTokenTransitionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to **string** | Mechanism by which the transition was initiated. | [optional] 
**DigitalWalletToken** | [**DigitalWalletTokenHash**](DigitalWalletTokenHash.md) |  | 
**Reason** | Pointer to **string** | The reason for the transition. | [optional] 
**ReasonCode** | Pointer to **string** | Standard code describing the reason for the transition.  *NOTE:* This field is required if your program uses v2 of the &#x60;user_card_state_version&#x60;, which is a program-specific configuration value that is managed by Marqeta and cannot be accessed via the API. To learn more about the &#x60;user_card_state_version&#x60; program configuration, contact your Marqeta representative.  * *00:* Object activated for the first time * *01:* Requested by you * *02:* Inactivity over time * *03:* This address cannot accept mail or the addressee is unknown * *04:* Negative account balance * *05:* Account under review * *06:* Suspicious activity was identified * *07:* Activity outside the program parameters was identified * *08:* Confirmed fraud was identified * *09:* Matched with an Office of Foreign Assets Control list * *10:* Card was reported lost * *11:* Card information was cloned * *12:* Account or card information was compromised * *13:* Temporary status change while on hold/leave * *14:* Initiated by Marqeta * *15:* Initiated by issuer * *16:* Card expired * *17:* Failed KYC * *18:* Changed to &#x60;ACTIVE&#x60; because information was properly validated * *19:* Changed to &#x60;ACTIVE&#x60; because account activity was properly validated * *20:* Change occurred prior to the normalization of reason codes * *21:* Initiated by a third party, often a digital wallet provider * *22:* PIN retry limit reached * *23:* Card was reported stolen * *24:* Address issue * *25:* Name issue * *26:* SSN issue * *27:* DOB issue * *28:* Email issue * *29:* Phone issue * *30:* Account/fulfillment mismatch * *31:* Other reason | [optional] 
**State** | **string** | Specifies the state to which the digital wallet token will transition.  The original state is &#x60;REQUESTED&#x60;. You cannot modify the state if its current value is either &#x60;REQUEST_DECLINED&#x60; or &#x60;TERMINATED&#x60;. | 
**Token** | Pointer to **string** | The unique identifier of the digital wallet token transition (not the identifier of the digital wallet token itself).  If you do not include a value for the &#x60;token&#x60; field, the system will generate one automatically. This value is necessary for use in other API calls, so we recommend that rather than let the system generate one, you use a simple string that is easy to remember. This value cannot be updated. | [optional] 
**TokenReferenceId** | Pointer to **string** | The unique identifier of the digital wallet token within the card network. The &#x60;token_reference_id&#x60; is unique at the card network level. | [optional] 

## Methods

### NewDigitalWalletTokenTransitionRequest

`func NewDigitalWalletTokenTransitionRequest(digitalWalletToken DigitalWalletTokenHash, state string, ) *DigitalWalletTokenTransitionRequest`

NewDigitalWalletTokenTransitionRequest instantiates a new DigitalWalletTokenTransitionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDigitalWalletTokenTransitionRequestWithDefaults

`func NewDigitalWalletTokenTransitionRequestWithDefaults() *DigitalWalletTokenTransitionRequest`

NewDigitalWalletTokenTransitionRequestWithDefaults instantiates a new DigitalWalletTokenTransitionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *DigitalWalletTokenTransitionRequest) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *DigitalWalletTokenTransitionRequest) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *DigitalWalletTokenTransitionRequest) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *DigitalWalletTokenTransitionRequest) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetDigitalWalletToken

`func (o *DigitalWalletTokenTransitionRequest) GetDigitalWalletToken() DigitalWalletTokenHash`

GetDigitalWalletToken returns the DigitalWalletToken field if non-nil, zero value otherwise.

### GetDigitalWalletTokenOk

`func (o *DigitalWalletTokenTransitionRequest) GetDigitalWalletTokenOk() (*DigitalWalletTokenHash, bool)`

GetDigitalWalletTokenOk returns a tuple with the DigitalWalletToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigitalWalletToken

`func (o *DigitalWalletTokenTransitionRequest) SetDigitalWalletToken(v DigitalWalletTokenHash)`

SetDigitalWalletToken sets DigitalWalletToken field to given value.


### GetReason

`func (o *DigitalWalletTokenTransitionRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *DigitalWalletTokenTransitionRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *DigitalWalletTokenTransitionRequest) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *DigitalWalletTokenTransitionRequest) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetReasonCode

`func (o *DigitalWalletTokenTransitionRequest) GetReasonCode() string`

GetReasonCode returns the ReasonCode field if non-nil, zero value otherwise.

### GetReasonCodeOk

`func (o *DigitalWalletTokenTransitionRequest) GetReasonCodeOk() (*string, bool)`

GetReasonCodeOk returns a tuple with the ReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonCode

`func (o *DigitalWalletTokenTransitionRequest) SetReasonCode(v string)`

SetReasonCode sets ReasonCode field to given value.

### HasReasonCode

`func (o *DigitalWalletTokenTransitionRequest) HasReasonCode() bool`

HasReasonCode returns a boolean if a field has been set.

### GetState

`func (o *DigitalWalletTokenTransitionRequest) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DigitalWalletTokenTransitionRequest) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DigitalWalletTokenTransitionRequest) SetState(v string)`

SetState sets State field to given value.


### GetToken

`func (o *DigitalWalletTokenTransitionRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DigitalWalletTokenTransitionRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DigitalWalletTokenTransitionRequest) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DigitalWalletTokenTransitionRequest) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTokenReferenceId

`func (o *DigitalWalletTokenTransitionRequest) GetTokenReferenceId() string`

GetTokenReferenceId returns the TokenReferenceId field if non-nil, zero value otherwise.

### GetTokenReferenceIdOk

`func (o *DigitalWalletTokenTransitionRequest) GetTokenReferenceIdOk() (*string, bool)`

GetTokenReferenceIdOk returns a tuple with the TokenReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenReferenceId

`func (o *DigitalWalletTokenTransitionRequest) SetTokenReferenceId(v string)`

SetTokenReferenceId sets TokenReferenceId field to given value.

### HasTokenReferenceId

`func (o *DigitalWalletTokenTransitionRequest) HasTokenReferenceId() bool`

HasTokenReferenceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


