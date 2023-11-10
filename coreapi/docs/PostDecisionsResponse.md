# PostDecisionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationToken** | **string** | Unique identifier of the application. | 
**DecisionId** | **string** | Unique identifier of the decision made based on the decision model.  See &#x60;decision_model.status&#x60; for the current state of the application. | 
**Status** | **string** | Status of the decision on the application. | 
**SubmittedDateTime** | **time.Time** | Date and time when the decision was submitted, in UTC. | 
**Token** | **string** | Unique identifier of the decision model.  See &#x60;decision_model.status&#x60; for the current state of the application. | 
**UserToken** | **string** | Unique identifier of the applicant, the user applying for a credit account. | 

## Methods

### NewPostDecisionsResponse

`func NewPostDecisionsResponse(applicationToken string, decisionId string, status string, submittedDateTime time.Time, token string, userToken string, ) *PostDecisionsResponse`

NewPostDecisionsResponse instantiates a new PostDecisionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostDecisionsResponseWithDefaults

`func NewPostDecisionsResponseWithDefaults() *PostDecisionsResponse`

NewPostDecisionsResponseWithDefaults instantiates a new PostDecisionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationToken

`func (o *PostDecisionsResponse) GetApplicationToken() string`

GetApplicationToken returns the ApplicationToken field if non-nil, zero value otherwise.

### GetApplicationTokenOk

`func (o *PostDecisionsResponse) GetApplicationTokenOk() (*string, bool)`

GetApplicationTokenOk returns a tuple with the ApplicationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationToken

`func (o *PostDecisionsResponse) SetApplicationToken(v string)`

SetApplicationToken sets ApplicationToken field to given value.


### GetDecisionId

`func (o *PostDecisionsResponse) GetDecisionId() string`

GetDecisionId returns the DecisionId field if non-nil, zero value otherwise.

### GetDecisionIdOk

`func (o *PostDecisionsResponse) GetDecisionIdOk() (*string, bool)`

GetDecisionIdOk returns a tuple with the DecisionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecisionId

`func (o *PostDecisionsResponse) SetDecisionId(v string)`

SetDecisionId sets DecisionId field to given value.


### GetStatus

`func (o *PostDecisionsResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PostDecisionsResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PostDecisionsResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSubmittedDateTime

`func (o *PostDecisionsResponse) GetSubmittedDateTime() time.Time`

GetSubmittedDateTime returns the SubmittedDateTime field if non-nil, zero value otherwise.

### GetSubmittedDateTimeOk

`func (o *PostDecisionsResponse) GetSubmittedDateTimeOk() (*time.Time, bool)`

GetSubmittedDateTimeOk returns a tuple with the SubmittedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmittedDateTime

`func (o *PostDecisionsResponse) SetSubmittedDateTime(v time.Time)`

SetSubmittedDateTime sets SubmittedDateTime field to given value.


### GetToken

`func (o *PostDecisionsResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PostDecisionsResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PostDecisionsResponse) SetToken(v string)`

SetToken sets Token field to given value.


### GetUserToken

`func (o *PostDecisionsResponse) GetUserToken() string`

GetUserToken returns the UserToken field if non-nil, zero value otherwise.

### GetUserTokenOk

`func (o *PostDecisionsResponse) GetUserTokenOk() (*string, bool)`

GetUserTokenOk returns a tuple with the UserToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserToken

`func (o *PostDecisionsResponse) SetUserToken(v string)`

SetUserToken sets UserToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


