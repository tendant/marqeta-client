# ApplicationTransitionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationState** | [**ApplicationResourceState**](ApplicationResourceState.md) |  | 
**BenefitsDisclosureTrackingToken** | Pointer to **string** | The tracking token of the Benefits Disclosure.  This is the &#x60;tracking_token&#x60; returned in the &#x60;BENEFITS_DISCLOSURE&#x60; object when sending a &#x60;GET&#x60; request to the &#x60;/credit/applications/files&#x60; endpoint before the user makes a decision on their approved application.  Required if transitioning application state to &#x60;ACCEPTED&#x60; | [optional] 
**CardMemberAgreementTrackingToken** | Pointer to **string** | The tracking token of the Card Member Agreement.  This is the &#x60;tracking_token&#x60; returned in the &#x60;CARD_MEMBER_AGREEMENT&#x60; object when sending a &#x60;GET&#x60; request to the &#x60;/credit/applications/files&#x60; endpoint before the user makes a decision on their approved application.  Required if transitioning application state to &#x60;ACCEPTED&#x60; | [optional] 
**RewardsDisclosurePostTermsTrackingToken** | Pointer to **string** | The tracking token of the Rewards Disclosure Post-terms.  This is the &#x60;tracking_token&#x60; returned in the &#x60;REWARDS_DISCLOSURE_POST_TERMS&#x60; object when sending a &#x60;GET&#x60; request to the &#x60;/credit/applications/files&#x60; endpoint before the user makes a decision on their approved application.  Required if transitioning application state to &#x60;ACCEPTED&#x60; | [optional] 
**TermsScheduleTrackingToken** | Pointer to **string** | The tracking token of the Terms Schedule.  This is the &#x60;tracking_token&#x60; returned in the &#x60;TERMS_SCHEDULE&#x60; object when sending a &#x60;GET&#x60; request to the &#x60;/credit/applications/files&#x60; endpoint before the user makes a decision on their approved application.  Required if transitioning application state to &#x60;ACCEPTED&#x60; | [optional] 
**Token** | Pointer to **string** | Main identifier of the resource, also used as the idempotency key on the request. | [optional] 

## Methods

### NewApplicationTransitionRequest

`func NewApplicationTransitionRequest(applicationState ApplicationResourceState, ) *ApplicationTransitionRequest`

NewApplicationTransitionRequest instantiates a new ApplicationTransitionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationTransitionRequestWithDefaults

`func NewApplicationTransitionRequestWithDefaults() *ApplicationTransitionRequest`

NewApplicationTransitionRequestWithDefaults instantiates a new ApplicationTransitionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationState

`func (o *ApplicationTransitionRequest) GetApplicationState() ApplicationResourceState`

GetApplicationState returns the ApplicationState field if non-nil, zero value otherwise.

### GetApplicationStateOk

`func (o *ApplicationTransitionRequest) GetApplicationStateOk() (*ApplicationResourceState, bool)`

GetApplicationStateOk returns a tuple with the ApplicationState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationState

`func (o *ApplicationTransitionRequest) SetApplicationState(v ApplicationResourceState)`

SetApplicationState sets ApplicationState field to given value.


### GetBenefitsDisclosureTrackingToken

`func (o *ApplicationTransitionRequest) GetBenefitsDisclosureTrackingToken() string`

GetBenefitsDisclosureTrackingToken returns the BenefitsDisclosureTrackingToken field if non-nil, zero value otherwise.

### GetBenefitsDisclosureTrackingTokenOk

`func (o *ApplicationTransitionRequest) GetBenefitsDisclosureTrackingTokenOk() (*string, bool)`

GetBenefitsDisclosureTrackingTokenOk returns a tuple with the BenefitsDisclosureTrackingToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBenefitsDisclosureTrackingToken

`func (o *ApplicationTransitionRequest) SetBenefitsDisclosureTrackingToken(v string)`

SetBenefitsDisclosureTrackingToken sets BenefitsDisclosureTrackingToken field to given value.

### HasBenefitsDisclosureTrackingToken

`func (o *ApplicationTransitionRequest) HasBenefitsDisclosureTrackingToken() bool`

HasBenefitsDisclosureTrackingToken returns a boolean if a field has been set.

### GetCardMemberAgreementTrackingToken

`func (o *ApplicationTransitionRequest) GetCardMemberAgreementTrackingToken() string`

GetCardMemberAgreementTrackingToken returns the CardMemberAgreementTrackingToken field if non-nil, zero value otherwise.

### GetCardMemberAgreementTrackingTokenOk

`func (o *ApplicationTransitionRequest) GetCardMemberAgreementTrackingTokenOk() (*string, bool)`

GetCardMemberAgreementTrackingTokenOk returns a tuple with the CardMemberAgreementTrackingToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardMemberAgreementTrackingToken

`func (o *ApplicationTransitionRequest) SetCardMemberAgreementTrackingToken(v string)`

SetCardMemberAgreementTrackingToken sets CardMemberAgreementTrackingToken field to given value.

### HasCardMemberAgreementTrackingToken

`func (o *ApplicationTransitionRequest) HasCardMemberAgreementTrackingToken() bool`

HasCardMemberAgreementTrackingToken returns a boolean if a field has been set.

### GetRewardsDisclosurePostTermsTrackingToken

`func (o *ApplicationTransitionRequest) GetRewardsDisclosurePostTermsTrackingToken() string`

GetRewardsDisclosurePostTermsTrackingToken returns the RewardsDisclosurePostTermsTrackingToken field if non-nil, zero value otherwise.

### GetRewardsDisclosurePostTermsTrackingTokenOk

`func (o *ApplicationTransitionRequest) GetRewardsDisclosurePostTermsTrackingTokenOk() (*string, bool)`

GetRewardsDisclosurePostTermsTrackingTokenOk returns a tuple with the RewardsDisclosurePostTermsTrackingToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardsDisclosurePostTermsTrackingToken

`func (o *ApplicationTransitionRequest) SetRewardsDisclosurePostTermsTrackingToken(v string)`

SetRewardsDisclosurePostTermsTrackingToken sets RewardsDisclosurePostTermsTrackingToken field to given value.

### HasRewardsDisclosurePostTermsTrackingToken

`func (o *ApplicationTransitionRequest) HasRewardsDisclosurePostTermsTrackingToken() bool`

HasRewardsDisclosurePostTermsTrackingToken returns a boolean if a field has been set.

### GetTermsScheduleTrackingToken

`func (o *ApplicationTransitionRequest) GetTermsScheduleTrackingToken() string`

GetTermsScheduleTrackingToken returns the TermsScheduleTrackingToken field if non-nil, zero value otherwise.

### GetTermsScheduleTrackingTokenOk

`func (o *ApplicationTransitionRequest) GetTermsScheduleTrackingTokenOk() (*string, bool)`

GetTermsScheduleTrackingTokenOk returns a tuple with the TermsScheduleTrackingToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsScheduleTrackingToken

`func (o *ApplicationTransitionRequest) SetTermsScheduleTrackingToken(v string)`

SetTermsScheduleTrackingToken sets TermsScheduleTrackingToken field to given value.

### HasTermsScheduleTrackingToken

`func (o *ApplicationTransitionRequest) HasTermsScheduleTrackingToken() bool`

HasTermsScheduleTrackingToken returns a boolean if a field has been set.

### GetToken

`func (o *ApplicationTransitionRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ApplicationTransitionRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ApplicationTransitionRequest) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *ApplicationTransitionRequest) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


