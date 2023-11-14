# DecisionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdverseActionTemplateCode** | Pointer to **string** | Indicates the version of the Notice of Adverse Action (NOAA) template used. Can have these possible values:  * &#x60;AA0&#x60; - score denial with score disclosure * &#x60;AA1&#x60; - individual reason with score disclosure * &#x60;AA2&#x60; - individual reason without score disclosure * &#x60;AA3&#x60; - locked and frozen * &#x60;AA4&#x60; - fraud related | [optional] 
**CardProductLevel** | Pointer to **string** | The tier of the card product. | [optional] 
**CreatedDate** | Pointer to **time.Time** | Date and time when the decision model was created on the Marqeta platform, in UTC. | [optional] 
**CreditLimit** | Pointer to **int32** | The maximum line of credit extended to the user, also the maximum balance the credit account can carry. | [optional] 
**CreditScore** | Pointer to **int32** | The user&#39;s credit score. | [optional] 
**CreditScoreDate** | Pointer to **string** | Date and time when the credit score went into effect. | [optional] 
**DecisionDate** | Pointer to **time.Time** | Date and time when the decision on the application was rendered, in UTC. | [optional] 
**DenialReasons** | Pointer to **[]string** | An array of reasons that explain why the application was declined. | [optional] 
**ExpireDate** | Pointer to **string** | Date when the decision model expires. | [optional] 
**Margin** | Pointer to **float32** | Number of percentage points added to the prime rate, used to calculate the purchase APR. | [optional] 
**PrimeRate** | Pointer to **float32** | The current prime rate set by the Fed. | [optional] 
**PurchaseApr** | Pointer to **float32** | The purchase APR approved for the user. | [optional] 
**ReceivedBestRate** | Pointer to **bool** | A value of &#x60;true&#x60; indicates that the user received the credit product&#39;s best APR.  If &#x60;false&#x60;, you must display to the user the following: &#x60;score_factors&#x60;, &#x60;credit_score&#x60;, &#x60;credit_score_date&#x60;, &#x60;credit_bureau&#x60;, &#x60;score_range&#x60;. | [optional] 
**ScoreFactors** | Pointer to **[]string** | Factors that the bureau used to determine the user&#39;s credit score. | [optional] 
**ScoreRange** | Pointer to **string** | The range in which the user&#39;s credit score falls. | [optional] 
**ApplicationToken** | **string** | Unique identifier of the application. | 
**DecisionId** | **string** | Unique identifier of the decision made based on the decision model.  See &#x60;decision_model.status&#x60; for the current state of the application. | 
**Status** | **string** | Status of the decision on the application. | 
**SubmittedDateTime** | **time.Time** | Date and time when the decision was submitted, in UTC. | 
**Token** | **string** | Unique identifier of the decision model.  See &#x60;decision_model.status&#x60; for the current state of the application. | 
**UserToken** | **string** | Unique identifier of the applicant, the user applying for a credit account. | 

## Methods

### NewDecisionsResponse

`func NewDecisionsResponse(applicationToken string, decisionId string, status string, submittedDateTime time.Time, token string, userToken string, ) *DecisionsResponse`

NewDecisionsResponse instantiates a new DecisionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDecisionsResponseWithDefaults

`func NewDecisionsResponseWithDefaults() *DecisionsResponse`

NewDecisionsResponseWithDefaults instantiates a new DecisionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdverseActionTemplateCode

`func (o *DecisionsResponse) GetAdverseActionTemplateCode() string`

GetAdverseActionTemplateCode returns the AdverseActionTemplateCode field if non-nil, zero value otherwise.

### GetAdverseActionTemplateCodeOk

`func (o *DecisionsResponse) GetAdverseActionTemplateCodeOk() (*string, bool)`

GetAdverseActionTemplateCodeOk returns a tuple with the AdverseActionTemplateCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdverseActionTemplateCode

`func (o *DecisionsResponse) SetAdverseActionTemplateCode(v string)`

SetAdverseActionTemplateCode sets AdverseActionTemplateCode field to given value.

### HasAdverseActionTemplateCode

`func (o *DecisionsResponse) HasAdverseActionTemplateCode() bool`

HasAdverseActionTemplateCode returns a boolean if a field has been set.

### GetCardProductLevel

`func (o *DecisionsResponse) GetCardProductLevel() string`

GetCardProductLevel returns the CardProductLevel field if non-nil, zero value otherwise.

### GetCardProductLevelOk

`func (o *DecisionsResponse) GetCardProductLevelOk() (*string, bool)`

GetCardProductLevelOk returns a tuple with the CardProductLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardProductLevel

`func (o *DecisionsResponse) SetCardProductLevel(v string)`

SetCardProductLevel sets CardProductLevel field to given value.

### HasCardProductLevel

`func (o *DecisionsResponse) HasCardProductLevel() bool`

HasCardProductLevel returns a boolean if a field has been set.

### GetCreatedDate

`func (o *DecisionsResponse) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *DecisionsResponse) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *DecisionsResponse) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *DecisionsResponse) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetCreditLimit

`func (o *DecisionsResponse) GetCreditLimit() int32`

GetCreditLimit returns the CreditLimit field if non-nil, zero value otherwise.

### GetCreditLimitOk

`func (o *DecisionsResponse) GetCreditLimitOk() (*int32, bool)`

GetCreditLimitOk returns a tuple with the CreditLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditLimit

`func (o *DecisionsResponse) SetCreditLimit(v int32)`

SetCreditLimit sets CreditLimit field to given value.

### HasCreditLimit

`func (o *DecisionsResponse) HasCreditLimit() bool`

HasCreditLimit returns a boolean if a field has been set.

### GetCreditScore

`func (o *DecisionsResponse) GetCreditScore() int32`

GetCreditScore returns the CreditScore field if non-nil, zero value otherwise.

### GetCreditScoreOk

`func (o *DecisionsResponse) GetCreditScoreOk() (*int32, bool)`

GetCreditScoreOk returns a tuple with the CreditScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditScore

`func (o *DecisionsResponse) SetCreditScore(v int32)`

SetCreditScore sets CreditScore field to given value.

### HasCreditScore

`func (o *DecisionsResponse) HasCreditScore() bool`

HasCreditScore returns a boolean if a field has been set.

### GetCreditScoreDate

`func (o *DecisionsResponse) GetCreditScoreDate() string`

GetCreditScoreDate returns the CreditScoreDate field if non-nil, zero value otherwise.

### GetCreditScoreDateOk

`func (o *DecisionsResponse) GetCreditScoreDateOk() (*string, bool)`

GetCreditScoreDateOk returns a tuple with the CreditScoreDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditScoreDate

`func (o *DecisionsResponse) SetCreditScoreDate(v string)`

SetCreditScoreDate sets CreditScoreDate field to given value.

### HasCreditScoreDate

`func (o *DecisionsResponse) HasCreditScoreDate() bool`

HasCreditScoreDate returns a boolean if a field has been set.

### GetDecisionDate

`func (o *DecisionsResponse) GetDecisionDate() time.Time`

GetDecisionDate returns the DecisionDate field if non-nil, zero value otherwise.

### GetDecisionDateOk

`func (o *DecisionsResponse) GetDecisionDateOk() (*time.Time, bool)`

GetDecisionDateOk returns a tuple with the DecisionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecisionDate

`func (o *DecisionsResponse) SetDecisionDate(v time.Time)`

SetDecisionDate sets DecisionDate field to given value.

### HasDecisionDate

`func (o *DecisionsResponse) HasDecisionDate() bool`

HasDecisionDate returns a boolean if a field has been set.

### GetDenialReasons

`func (o *DecisionsResponse) GetDenialReasons() []string`

GetDenialReasons returns the DenialReasons field if non-nil, zero value otherwise.

### GetDenialReasonsOk

`func (o *DecisionsResponse) GetDenialReasonsOk() (*[]string, bool)`

GetDenialReasonsOk returns a tuple with the DenialReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenialReasons

`func (o *DecisionsResponse) SetDenialReasons(v []string)`

SetDenialReasons sets DenialReasons field to given value.

### HasDenialReasons

`func (o *DecisionsResponse) HasDenialReasons() bool`

HasDenialReasons returns a boolean if a field has been set.

### GetExpireDate

`func (o *DecisionsResponse) GetExpireDate() string`

GetExpireDate returns the ExpireDate field if non-nil, zero value otherwise.

### GetExpireDateOk

`func (o *DecisionsResponse) GetExpireDateOk() (*string, bool)`

GetExpireDateOk returns a tuple with the ExpireDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireDate

`func (o *DecisionsResponse) SetExpireDate(v string)`

SetExpireDate sets ExpireDate field to given value.

### HasExpireDate

`func (o *DecisionsResponse) HasExpireDate() bool`

HasExpireDate returns a boolean if a field has been set.

### GetMargin

`func (o *DecisionsResponse) GetMargin() float32`

GetMargin returns the Margin field if non-nil, zero value otherwise.

### GetMarginOk

`func (o *DecisionsResponse) GetMarginOk() (*float32, bool)`

GetMarginOk returns a tuple with the Margin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMargin

`func (o *DecisionsResponse) SetMargin(v float32)`

SetMargin sets Margin field to given value.

### HasMargin

`func (o *DecisionsResponse) HasMargin() bool`

HasMargin returns a boolean if a field has been set.

### GetPrimeRate

`func (o *DecisionsResponse) GetPrimeRate() float32`

GetPrimeRate returns the PrimeRate field if non-nil, zero value otherwise.

### GetPrimeRateOk

`func (o *DecisionsResponse) GetPrimeRateOk() (*float32, bool)`

GetPrimeRateOk returns a tuple with the PrimeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimeRate

`func (o *DecisionsResponse) SetPrimeRate(v float32)`

SetPrimeRate sets PrimeRate field to given value.

### HasPrimeRate

`func (o *DecisionsResponse) HasPrimeRate() bool`

HasPrimeRate returns a boolean if a field has been set.

### GetPurchaseApr

`func (o *DecisionsResponse) GetPurchaseApr() float32`

GetPurchaseApr returns the PurchaseApr field if non-nil, zero value otherwise.

### GetPurchaseAprOk

`func (o *DecisionsResponse) GetPurchaseAprOk() (*float32, bool)`

GetPurchaseAprOk returns a tuple with the PurchaseApr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseApr

`func (o *DecisionsResponse) SetPurchaseApr(v float32)`

SetPurchaseApr sets PurchaseApr field to given value.

### HasPurchaseApr

`func (o *DecisionsResponse) HasPurchaseApr() bool`

HasPurchaseApr returns a boolean if a field has been set.

### GetReceivedBestRate

`func (o *DecisionsResponse) GetReceivedBestRate() bool`

GetReceivedBestRate returns the ReceivedBestRate field if non-nil, zero value otherwise.

### GetReceivedBestRateOk

`func (o *DecisionsResponse) GetReceivedBestRateOk() (*bool, bool)`

GetReceivedBestRateOk returns a tuple with the ReceivedBestRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedBestRate

`func (o *DecisionsResponse) SetReceivedBestRate(v bool)`

SetReceivedBestRate sets ReceivedBestRate field to given value.

### HasReceivedBestRate

`func (o *DecisionsResponse) HasReceivedBestRate() bool`

HasReceivedBestRate returns a boolean if a field has been set.

### GetScoreFactors

`func (o *DecisionsResponse) GetScoreFactors() []string`

GetScoreFactors returns the ScoreFactors field if non-nil, zero value otherwise.

### GetScoreFactorsOk

`func (o *DecisionsResponse) GetScoreFactorsOk() (*[]string, bool)`

GetScoreFactorsOk returns a tuple with the ScoreFactors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoreFactors

`func (o *DecisionsResponse) SetScoreFactors(v []string)`

SetScoreFactors sets ScoreFactors field to given value.

### HasScoreFactors

`func (o *DecisionsResponse) HasScoreFactors() bool`

HasScoreFactors returns a boolean if a field has been set.

### GetScoreRange

`func (o *DecisionsResponse) GetScoreRange() string`

GetScoreRange returns the ScoreRange field if non-nil, zero value otherwise.

### GetScoreRangeOk

`func (o *DecisionsResponse) GetScoreRangeOk() (*string, bool)`

GetScoreRangeOk returns a tuple with the ScoreRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoreRange

`func (o *DecisionsResponse) SetScoreRange(v string)`

SetScoreRange sets ScoreRange field to given value.

### HasScoreRange

`func (o *DecisionsResponse) HasScoreRange() bool`

HasScoreRange returns a boolean if a field has been set.

### GetApplicationToken

`func (o *DecisionsResponse) GetApplicationToken() string`

GetApplicationToken returns the ApplicationToken field if non-nil, zero value otherwise.

### GetApplicationTokenOk

`func (o *DecisionsResponse) GetApplicationTokenOk() (*string, bool)`

GetApplicationTokenOk returns a tuple with the ApplicationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationToken

`func (o *DecisionsResponse) SetApplicationToken(v string)`

SetApplicationToken sets ApplicationToken field to given value.


### GetDecisionId

`func (o *DecisionsResponse) GetDecisionId() string`

GetDecisionId returns the DecisionId field if non-nil, zero value otherwise.

### GetDecisionIdOk

`func (o *DecisionsResponse) GetDecisionIdOk() (*string, bool)`

GetDecisionIdOk returns a tuple with the DecisionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecisionId

`func (o *DecisionsResponse) SetDecisionId(v string)`

SetDecisionId sets DecisionId field to given value.


### GetStatus

`func (o *DecisionsResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DecisionsResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DecisionsResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSubmittedDateTime

`func (o *DecisionsResponse) GetSubmittedDateTime() time.Time`

GetSubmittedDateTime returns the SubmittedDateTime field if non-nil, zero value otherwise.

### GetSubmittedDateTimeOk

`func (o *DecisionsResponse) GetSubmittedDateTimeOk() (*time.Time, bool)`

GetSubmittedDateTimeOk returns a tuple with the SubmittedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmittedDateTime

`func (o *DecisionsResponse) SetSubmittedDateTime(v time.Time)`

SetSubmittedDateTime sets SubmittedDateTime field to given value.


### GetToken

`func (o *DecisionsResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DecisionsResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DecisionsResponse) SetToken(v string)`

SetToken sets Token field to given value.


### GetUserToken

`func (o *DecisionsResponse) GetUserToken() string`

GetUserToken returns the UserToken field if non-nil, zero value otherwise.

### GetUserTokenOk

`func (o *DecisionsResponse) GetUserTokenOk() (*string, bool)`

GetUserTokenOk returns a tuple with the UserToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserToken

`func (o *DecisionsResponse) SetUserToken(v string)`

SetUserToken sets UserToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


