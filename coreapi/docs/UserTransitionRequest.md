# UserTransitionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | **string** | Mechanism by which the transaction was initiated. | 
**IdempotentHash** | Pointer to **string** | Unique hashed value that identifies subsequent submissions of the user transition request. | [optional] 
**Reason** | Pointer to **string** | Additional information about the status change. | [optional] 
**ReasonCode** | **string** | Identifies the standardized reason for the transition:  *00:* Object activated for the first time.  *01:* Requested by you.  *02:* Inactivity over time.  *03:* This address cannot accept mail or the addressee is unknown.  *04:* Negative account balance.  *05:* Account under review.  *06:* Suspicious activity was identified.  *07:* Activity outside the program parameters was identified.  *08:* Confirmed fraud was identified.  *09:* Matched with an Office of Foreign Assets Control list.  *10:* Card was reported lost.  *11:* Card information was cloned.  *12:* Account or card information was compromised.  *13:* Temporary status change while on hold/leave.  *14:* Initiated by Marqeta.  *15:* Initiated by issuer.  *16:* Card expired.  *17:* Failed KYC.  *18:* Changed to &#x60;ACTIVE&#x60; because information was properly validated.  *19:* Changed to &#x60;ACTIVE&#x60; because account activity was properly validated.  *20:* Change occurred prior to the normalization of reason codes.  *21:* Initiated by a third party, often a digital wallet provider.  *22:* PIN retry limit reached.  *23:* Card was reported stolen.  *24:* Address issue.  *25:* Name issue.  *26:* SSN issue.  *27:* DOB issue.  *28:* Email issue.  *29:* Phone issue.  *30:* Account/fulfillment mismatch.  *31:* Other reason. | 
**Status** | **string** | Specifies the new status of the user. | 
**Token** | Pointer to **string** | Unique identifier of the user transition.  If you do not include a token, the system generates one automatically. This token is referenced in other API calls, so we recommend that you define a simple string that is easy to remember. This value cannot be updated. | [optional] 
**UserToken** | **string** | Unique identifier of the user whose status you want to transition. | 

## Methods

### NewUserTransitionRequest

`func NewUserTransitionRequest(channel string, reasonCode string, status string, userToken string, ) *UserTransitionRequest`

NewUserTransitionRequest instantiates a new UserTransitionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserTransitionRequestWithDefaults

`func NewUserTransitionRequestWithDefaults() *UserTransitionRequest`

NewUserTransitionRequestWithDefaults instantiates a new UserTransitionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *UserTransitionRequest) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *UserTransitionRequest) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *UserTransitionRequest) SetChannel(v string)`

SetChannel sets Channel field to given value.


### GetIdempotentHash

`func (o *UserTransitionRequest) GetIdempotentHash() string`

GetIdempotentHash returns the IdempotentHash field if non-nil, zero value otherwise.

### GetIdempotentHashOk

`func (o *UserTransitionRequest) GetIdempotentHashOk() (*string, bool)`

GetIdempotentHashOk returns a tuple with the IdempotentHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotentHash

`func (o *UserTransitionRequest) SetIdempotentHash(v string)`

SetIdempotentHash sets IdempotentHash field to given value.

### HasIdempotentHash

`func (o *UserTransitionRequest) HasIdempotentHash() bool`

HasIdempotentHash returns a boolean if a field has been set.

### GetReason

`func (o *UserTransitionRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *UserTransitionRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *UserTransitionRequest) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *UserTransitionRequest) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetReasonCode

`func (o *UserTransitionRequest) GetReasonCode() string`

GetReasonCode returns the ReasonCode field if non-nil, zero value otherwise.

### GetReasonCodeOk

`func (o *UserTransitionRequest) GetReasonCodeOk() (*string, bool)`

GetReasonCodeOk returns a tuple with the ReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonCode

`func (o *UserTransitionRequest) SetReasonCode(v string)`

SetReasonCode sets ReasonCode field to given value.


### GetStatus

`func (o *UserTransitionRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UserTransitionRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UserTransitionRequest) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetToken

`func (o *UserTransitionRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UserTransitionRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UserTransitionRequest) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *UserTransitionRequest) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUserToken

`func (o *UserTransitionRequest) GetUserToken() string`

GetUserToken returns the UserToken field if non-nil, zero value otherwise.

### GetUserTokenOk

`func (o *UserTransitionRequest) GetUserTokenOk() (*string, bool)`

GetUserTokenOk returns a tuple with the UserToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserToken

`func (o *UserTransitionRequest) SetUserToken(v string)`

SetUserToken sets UserToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


