# DigitalWalletTokenTransitionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardSwap** | Pointer to [**CardSwapHash**](CardSwapHash.md) |  | [optional] 
**Channel** | **string** | Mechanism by which the transition was initiated. | 
**CreatedTime** | Pointer to **time.Time** | Date and time when the digital wallet provisioning request was created, in UTC. | [optional] 
**DigitalWalletToken** | [**DigitalWalletTokenHash**](DigitalWalletTokenHash.md) |  | 
**FulfillmentStatus** | **string** | Provisioning status of the digital wallet token. | 
**Reason** | Pointer to **string** | Reason for the transition. | [optional] 
**ReasonCode** | Pointer to **string** | Standard code describing the reason for the transition:  * *00:* Object activated for the first time * *01:* Requested by you * *02:* Inactivity over time * *03:* This address cannot accept mail or the addressee is unknown * *04:* Negative account balance * *05:* Account under review * *06:* Suspicious activity was identified * *07:* Activity outside the program parameters was identified * *08:* Confirmed fraud was identified * *09:* Matched with an Office of Foreign Assets Control list * *10:* Card was reported lost * *11:* Card information was cloned * *12:* Account or card information was compromised * *13:* Temporary status change while on hold/leave * *14:* Initiated by Marqeta * *15:* Initiated by issuer * *16:* Card expired * *17:* Failed KYC * *18:* Changed to &#x60;ACTIVE&#x60; because information was properly validated * *19:* Changed to &#x60;ACTIVE&#x60; because account activity was properly validated * *20:* Change occurred prior to the normalization of reason codes * *21:* Initiated by a third party, often a digital wallet provider * *22:* PIN retry limit reached * *23:* Card was reported stolen * *24:* Address issue * *25:* Name issue * *26:* SSN issue * *27:* DOB issue * *28:* Email issue * *29:* Phone issue * *30:* Account/fulfillment mismatch * *31:* Other reason | [optional] 
**State** | **string** | Specifies the state to which the digital wallet token is transitioning. | 
**Token** | **string** | Unique identifier of the digital wallet token transition, and not the identifier of the digital wallet token itself. | 
**Type** | **string** | Type of digital wallet token transition. &#x60;state.activated&#x60;, for example. | [readonly] 

## Methods

### NewDigitalWalletTokenTransitionResponse

`func NewDigitalWalletTokenTransitionResponse(channel string, digitalWalletToken DigitalWalletTokenHash, fulfillmentStatus string, state string, token string, type_ string, ) *DigitalWalletTokenTransitionResponse`

NewDigitalWalletTokenTransitionResponse instantiates a new DigitalWalletTokenTransitionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDigitalWalletTokenTransitionResponseWithDefaults

`func NewDigitalWalletTokenTransitionResponseWithDefaults() *DigitalWalletTokenTransitionResponse`

NewDigitalWalletTokenTransitionResponseWithDefaults instantiates a new DigitalWalletTokenTransitionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardSwap

`func (o *DigitalWalletTokenTransitionResponse) GetCardSwap() CardSwapHash`

GetCardSwap returns the CardSwap field if non-nil, zero value otherwise.

### GetCardSwapOk

`func (o *DigitalWalletTokenTransitionResponse) GetCardSwapOk() (*CardSwapHash, bool)`

GetCardSwapOk returns a tuple with the CardSwap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardSwap

`func (o *DigitalWalletTokenTransitionResponse) SetCardSwap(v CardSwapHash)`

SetCardSwap sets CardSwap field to given value.

### HasCardSwap

`func (o *DigitalWalletTokenTransitionResponse) HasCardSwap() bool`

HasCardSwap returns a boolean if a field has been set.

### GetChannel

`func (o *DigitalWalletTokenTransitionResponse) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *DigitalWalletTokenTransitionResponse) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *DigitalWalletTokenTransitionResponse) SetChannel(v string)`

SetChannel sets Channel field to given value.


### GetCreatedTime

`func (o *DigitalWalletTokenTransitionResponse) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *DigitalWalletTokenTransitionResponse) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *DigitalWalletTokenTransitionResponse) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *DigitalWalletTokenTransitionResponse) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDigitalWalletToken

`func (o *DigitalWalletTokenTransitionResponse) GetDigitalWalletToken() DigitalWalletTokenHash`

GetDigitalWalletToken returns the DigitalWalletToken field if non-nil, zero value otherwise.

### GetDigitalWalletTokenOk

`func (o *DigitalWalletTokenTransitionResponse) GetDigitalWalletTokenOk() (*DigitalWalletTokenHash, bool)`

GetDigitalWalletTokenOk returns a tuple with the DigitalWalletToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigitalWalletToken

`func (o *DigitalWalletTokenTransitionResponse) SetDigitalWalletToken(v DigitalWalletTokenHash)`

SetDigitalWalletToken sets DigitalWalletToken field to given value.


### GetFulfillmentStatus

`func (o *DigitalWalletTokenTransitionResponse) GetFulfillmentStatus() string`

GetFulfillmentStatus returns the FulfillmentStatus field if non-nil, zero value otherwise.

### GetFulfillmentStatusOk

`func (o *DigitalWalletTokenTransitionResponse) GetFulfillmentStatusOk() (*string, bool)`

GetFulfillmentStatusOk returns a tuple with the FulfillmentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentStatus

`func (o *DigitalWalletTokenTransitionResponse) SetFulfillmentStatus(v string)`

SetFulfillmentStatus sets FulfillmentStatus field to given value.


### GetReason

`func (o *DigitalWalletTokenTransitionResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *DigitalWalletTokenTransitionResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *DigitalWalletTokenTransitionResponse) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *DigitalWalletTokenTransitionResponse) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetReasonCode

`func (o *DigitalWalletTokenTransitionResponse) GetReasonCode() string`

GetReasonCode returns the ReasonCode field if non-nil, zero value otherwise.

### GetReasonCodeOk

`func (o *DigitalWalletTokenTransitionResponse) GetReasonCodeOk() (*string, bool)`

GetReasonCodeOk returns a tuple with the ReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonCode

`func (o *DigitalWalletTokenTransitionResponse) SetReasonCode(v string)`

SetReasonCode sets ReasonCode field to given value.

### HasReasonCode

`func (o *DigitalWalletTokenTransitionResponse) HasReasonCode() bool`

HasReasonCode returns a boolean if a field has been set.

### GetState

`func (o *DigitalWalletTokenTransitionResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DigitalWalletTokenTransitionResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DigitalWalletTokenTransitionResponse) SetState(v string)`

SetState sets State field to given value.


### GetToken

`func (o *DigitalWalletTokenTransitionResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DigitalWalletTokenTransitionResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DigitalWalletTokenTransitionResponse) SetToken(v string)`

SetToken sets Token field to given value.


### GetType

`func (o *DigitalWalletTokenTransitionResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DigitalWalletTokenTransitionResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DigitalWalletTokenTransitionResponse) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


