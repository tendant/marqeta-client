# BusinessTransitionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessToken** | Pointer to **string** | Unique identifier of the business whose status you want to transition. | [optional] 
**Channel** | **string** | Mechanism by which the transaction was initiated. | 
**CreatedTime** | Pointer to **time.Time** | Date and time when the resource was created, in UTC. | [optional] 
**LastModifiedTime** | Pointer to **time.Time** | Date and time when the resource was last updated, in UTC. | [optional] 
**Reason** | Pointer to **string** | Additional information about the status change. | [optional] 
**ReasonCode** | **string** | Identifies the standardized reason for the transition:  *00:* Object activated for the first time.  *01:* Requested by you.  *02:* Inactivity over time.  *03:* This address cannot accept mail or the addressee is unknown.  *04:* Negative account balance.  *05:* Account under review.  *06:* Suspicious activity was identified.  *07:* Activity outside the program parameters was identified.  *08:* Confirmed fraud was identified.  *09:* Matched with an Office of Foreign Assets Control list.  *10:* Card was reported lost.  *11:* Card information was cloned.  *12:* Account or card information was compromised.  *13:* Temporary status change while on hold/leave.  *14:* Initiated by Marqeta.  *15:* Initiated by issuer.  *16:* Card expired.  *17:* Failed KYC.  *18:* Changed to &#x60;ACTIVE&#x60; because information was properly validated.  *19:* Changed to &#x60;ACTIVE&#x60; because account activity was properly validated.  *20:* Change occurred prior to the normalization of reason codes.  *21:* Initiated by a third party, often a digital wallet provider.  *22:* PIN retry limit reached.  *23:* Card was reported stolen.  *24:* Address issue.  *25:* Name issue.  *26:* SSN issue.  *27:* DOB issue.  *28:* Email issue.  *29:* Phone issue.  *30:* Account/fulfillment mismatch.  *31:* Other reason. | 
**Status** | **string** | Specifies the status of the business. | 
**Token** | **string** | Unique identifier of the business transition. | 

## Methods

### NewBusinessTransitionResponse

`func NewBusinessTransitionResponse(channel string, reasonCode string, status string, token string, ) *BusinessTransitionResponse`

NewBusinessTransitionResponse instantiates a new BusinessTransitionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBusinessTransitionResponseWithDefaults

`func NewBusinessTransitionResponseWithDefaults() *BusinessTransitionResponse`

NewBusinessTransitionResponseWithDefaults instantiates a new BusinessTransitionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessToken

`func (o *BusinessTransitionResponse) GetBusinessToken() string`

GetBusinessToken returns the BusinessToken field if non-nil, zero value otherwise.

### GetBusinessTokenOk

`func (o *BusinessTransitionResponse) GetBusinessTokenOk() (*string, bool)`

GetBusinessTokenOk returns a tuple with the BusinessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessToken

`func (o *BusinessTransitionResponse) SetBusinessToken(v string)`

SetBusinessToken sets BusinessToken field to given value.

### HasBusinessToken

`func (o *BusinessTransitionResponse) HasBusinessToken() bool`

HasBusinessToken returns a boolean if a field has been set.

### GetChannel

`func (o *BusinessTransitionResponse) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *BusinessTransitionResponse) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *BusinessTransitionResponse) SetChannel(v string)`

SetChannel sets Channel field to given value.


### GetCreatedTime

`func (o *BusinessTransitionResponse) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *BusinessTransitionResponse) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *BusinessTransitionResponse) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *BusinessTransitionResponse) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetLastModifiedTime

`func (o *BusinessTransitionResponse) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *BusinessTransitionResponse) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *BusinessTransitionResponse) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.

### HasLastModifiedTime

`func (o *BusinessTransitionResponse) HasLastModifiedTime() bool`

HasLastModifiedTime returns a boolean if a field has been set.

### GetReason

`func (o *BusinessTransitionResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *BusinessTransitionResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *BusinessTransitionResponse) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *BusinessTransitionResponse) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetReasonCode

`func (o *BusinessTransitionResponse) GetReasonCode() string`

GetReasonCode returns the ReasonCode field if non-nil, zero value otherwise.

### GetReasonCodeOk

`func (o *BusinessTransitionResponse) GetReasonCodeOk() (*string, bool)`

GetReasonCodeOk returns a tuple with the ReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonCode

`func (o *BusinessTransitionResponse) SetReasonCode(v string)`

SetReasonCode sets ReasonCode field to given value.


### GetStatus

`func (o *BusinessTransitionResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BusinessTransitionResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BusinessTransitionResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetToken

`func (o *BusinessTransitionResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *BusinessTransitionResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *BusinessTransitionResponse) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


