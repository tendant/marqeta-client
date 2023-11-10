# BusinessTransitionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessToken** | **string** | Unique identifier of the business whose status you want to transition. | 
**Channel** | **string** | Mechanism by which the transaction was initiated. | 
**IdempotentHash** | Pointer to **string** | Unique hashed value that identifies subsequent submissions of the business transition request. | [optional] 
**Reason** | Pointer to **string** | Additional information about the status change. | [optional] 
**ReasonCode** | **string** | Identifies the standardized reason for the transition:  *00:* Object activated for the first time.  *01:* Requested by you.  *02:* Inactivity over time.  *03:* This address cannot accept mail or the addressee is unknown.  *04:* Negative account balance.  *05:* Account under review.  *06:* Suspicious activity was identified.  *07:* Activity outside the program parameters was identified.  *08:* Confirmed fraud was identified.  *09:* Matched with an Office of Foreign Assets Control list.  *10:* Card was reported lost.  *11:* Card information was cloned.  *12:* Account or card information was compromised.  *13:* Temporary status change while on hold/leave.  *14:* Initiated by Marqeta.  *15:* Initiated by issuer.  *16:* Card expired.  *17:* Failed KYC.  *18:* Changed to &#x60;ACTIVE&#x60; because information was properly validated.  *19:* Changed to &#x60;ACTIVE&#x60; because account activity was properly validated.  *20:* Change occurred prior to the normalization of reason codes.  *21:* Initiated by a third party, often a digital wallet provider.  *22:* PIN retry limit reached.  *23:* Card was reported stolen.  *24:* Address issue.  *25:* Name issue.  *26:* SSN issue.  *27:* DOB issue.  *28:* Email issue.  *29:* Phone issue.  *30:* Account/fulfillment mismatch.  *31:* Other reason. | 
**Status** | **string** | Specifies the new status of the business. | 
**Token** | Pointer to **string** | Unique identifier of the business transition.  If you do not include a token, the system generates one automatically. This token is referenced in other API calls, so we recommend that you define a simple string that is easy to remember. This value cannot be updated. | [optional] 

## Methods

### NewBusinessTransitionRequest

`func NewBusinessTransitionRequest(businessToken string, channel string, reasonCode string, status string, ) *BusinessTransitionRequest`

NewBusinessTransitionRequest instantiates a new BusinessTransitionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBusinessTransitionRequestWithDefaults

`func NewBusinessTransitionRequestWithDefaults() *BusinessTransitionRequest`

NewBusinessTransitionRequestWithDefaults instantiates a new BusinessTransitionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessToken

`func (o *BusinessTransitionRequest) GetBusinessToken() string`

GetBusinessToken returns the BusinessToken field if non-nil, zero value otherwise.

### GetBusinessTokenOk

`func (o *BusinessTransitionRequest) GetBusinessTokenOk() (*string, bool)`

GetBusinessTokenOk returns a tuple with the BusinessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessToken

`func (o *BusinessTransitionRequest) SetBusinessToken(v string)`

SetBusinessToken sets BusinessToken field to given value.


### GetChannel

`func (o *BusinessTransitionRequest) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *BusinessTransitionRequest) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *BusinessTransitionRequest) SetChannel(v string)`

SetChannel sets Channel field to given value.


### GetIdempotentHash

`func (o *BusinessTransitionRequest) GetIdempotentHash() string`

GetIdempotentHash returns the IdempotentHash field if non-nil, zero value otherwise.

### GetIdempotentHashOk

`func (o *BusinessTransitionRequest) GetIdempotentHashOk() (*string, bool)`

GetIdempotentHashOk returns a tuple with the IdempotentHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotentHash

`func (o *BusinessTransitionRequest) SetIdempotentHash(v string)`

SetIdempotentHash sets IdempotentHash field to given value.

### HasIdempotentHash

`func (o *BusinessTransitionRequest) HasIdempotentHash() bool`

HasIdempotentHash returns a boolean if a field has been set.

### GetReason

`func (o *BusinessTransitionRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *BusinessTransitionRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *BusinessTransitionRequest) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *BusinessTransitionRequest) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetReasonCode

`func (o *BusinessTransitionRequest) GetReasonCode() string`

GetReasonCode returns the ReasonCode field if non-nil, zero value otherwise.

### GetReasonCodeOk

`func (o *BusinessTransitionRequest) GetReasonCodeOk() (*string, bool)`

GetReasonCodeOk returns a tuple with the ReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonCode

`func (o *BusinessTransitionRequest) SetReasonCode(v string)`

SetReasonCode sets ReasonCode field to given value.


### GetStatus

`func (o *BusinessTransitionRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BusinessTransitionRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BusinessTransitionRequest) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetToken

`func (o *BusinessTransitionRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *BusinessTransitionRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *BusinessTransitionRequest) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *BusinessTransitionRequest) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


