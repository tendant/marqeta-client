# TokenServiceProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CorrelationId** | Pointer to **string** | Unique value representing a tokenization request (Mastercard only). | [optional] 
**PanReferenceId** | **string** | Unique identifier of the digital wallet token primary account number (PAN) within the card network. | 
**TokenAssuranceLevel** | Pointer to **string** | _(Mastercard only)_ Represents the confidence level in the digital wallet token. | [optional] 
**TokenEligibilityDecision** | Pointer to **string** | Digital wallet&#39;s decision as to whether the digital wallet token should be provisioned. | [optional] 
**TokenExpiration** | Pointer to **string** | Expiration date of the digital wallet token. | [optional] 
**TokenPan** | Pointer to **string** | Primary account number (PAN) of the digital wallet token. | [optional] 
**TokenReferenceId** | Pointer to **string** | Unique identifier of the digital wallet token within the card network. | [optional] 
**TokenRequestorId** | Pointer to **string** | Unique numerical identifier of the token requestor within the card network. These ID numbers map to &#x60;token_requestor_name&#x60; field values as follows:  *Mastercard*  * 50110030273 – &#x60;APPLE_PAY&#x60; * 50120834693 – &#x60;ANDROID_PAY&#x60; * 50139059239 – &#x60;SAMSUNG_PAY&#x60;  *Visa*  * 40010030273 – &#x60;APPLE_PAY&#x60; * 40010075001 – &#x60;ANDROID_PAY&#x60; * 40010043095 – &#x60;SAMSUNG_PAY&#x60; * 40010075196 – &#x60;MICROSOFT_PAY&#x60; * 40010075338 – &#x60;VISA_CHECKOUT&#x60; * 40010075449 – &#x60;FACEBOOK&#x60; * 40010075839 – &#x60;NETFLIX&#x60; * 40010077056 – &#x60;FITBIT_PAY&#x60; * 40010069887 – &#x60;GARMIN_PAY&#x60; | [optional] 
**TokenRequestorName** | Pointer to **string** | Name of the token requestor within the card network.  *NOTE:* The list of example values for this field is maintained by the card networks and is subject to change. | [optional] 
**TokenScore** | Pointer to **string** | Token score assigned by the digital wallet. | [optional] 
**TokenType** | Pointer to **string** | Type of the digital wallet token. | [optional] 

## Methods

### NewTokenServiceProvider

`func NewTokenServiceProvider(panReferenceId string, ) *TokenServiceProvider`

NewTokenServiceProvider instantiates a new TokenServiceProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenServiceProviderWithDefaults

`func NewTokenServiceProviderWithDefaults() *TokenServiceProvider`

NewTokenServiceProviderWithDefaults instantiates a new TokenServiceProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCorrelationId

`func (o *TokenServiceProvider) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *TokenServiceProvider) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *TokenServiceProvider) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.

### HasCorrelationId

`func (o *TokenServiceProvider) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.

### GetPanReferenceId

`func (o *TokenServiceProvider) GetPanReferenceId() string`

GetPanReferenceId returns the PanReferenceId field if non-nil, zero value otherwise.

### GetPanReferenceIdOk

`func (o *TokenServiceProvider) GetPanReferenceIdOk() (*string, bool)`

GetPanReferenceIdOk returns a tuple with the PanReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPanReferenceId

`func (o *TokenServiceProvider) SetPanReferenceId(v string)`

SetPanReferenceId sets PanReferenceId field to given value.


### GetTokenAssuranceLevel

`func (o *TokenServiceProvider) GetTokenAssuranceLevel() string`

GetTokenAssuranceLevel returns the TokenAssuranceLevel field if non-nil, zero value otherwise.

### GetTokenAssuranceLevelOk

`func (o *TokenServiceProvider) GetTokenAssuranceLevelOk() (*string, bool)`

GetTokenAssuranceLevelOk returns a tuple with the TokenAssuranceLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenAssuranceLevel

`func (o *TokenServiceProvider) SetTokenAssuranceLevel(v string)`

SetTokenAssuranceLevel sets TokenAssuranceLevel field to given value.

### HasTokenAssuranceLevel

`func (o *TokenServiceProvider) HasTokenAssuranceLevel() bool`

HasTokenAssuranceLevel returns a boolean if a field has been set.

### GetTokenEligibilityDecision

`func (o *TokenServiceProvider) GetTokenEligibilityDecision() string`

GetTokenEligibilityDecision returns the TokenEligibilityDecision field if non-nil, zero value otherwise.

### GetTokenEligibilityDecisionOk

`func (o *TokenServiceProvider) GetTokenEligibilityDecisionOk() (*string, bool)`

GetTokenEligibilityDecisionOk returns a tuple with the TokenEligibilityDecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenEligibilityDecision

`func (o *TokenServiceProvider) SetTokenEligibilityDecision(v string)`

SetTokenEligibilityDecision sets TokenEligibilityDecision field to given value.

### HasTokenEligibilityDecision

`func (o *TokenServiceProvider) HasTokenEligibilityDecision() bool`

HasTokenEligibilityDecision returns a boolean if a field has been set.

### GetTokenExpiration

`func (o *TokenServiceProvider) GetTokenExpiration() string`

GetTokenExpiration returns the TokenExpiration field if non-nil, zero value otherwise.

### GetTokenExpirationOk

`func (o *TokenServiceProvider) GetTokenExpirationOk() (*string, bool)`

GetTokenExpirationOk returns a tuple with the TokenExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenExpiration

`func (o *TokenServiceProvider) SetTokenExpiration(v string)`

SetTokenExpiration sets TokenExpiration field to given value.

### HasTokenExpiration

`func (o *TokenServiceProvider) HasTokenExpiration() bool`

HasTokenExpiration returns a boolean if a field has been set.

### GetTokenPan

`func (o *TokenServiceProvider) GetTokenPan() string`

GetTokenPan returns the TokenPan field if non-nil, zero value otherwise.

### GetTokenPanOk

`func (o *TokenServiceProvider) GetTokenPanOk() (*string, bool)`

GetTokenPanOk returns a tuple with the TokenPan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenPan

`func (o *TokenServiceProvider) SetTokenPan(v string)`

SetTokenPan sets TokenPan field to given value.

### HasTokenPan

`func (o *TokenServiceProvider) HasTokenPan() bool`

HasTokenPan returns a boolean if a field has been set.

### GetTokenReferenceId

`func (o *TokenServiceProvider) GetTokenReferenceId() string`

GetTokenReferenceId returns the TokenReferenceId field if non-nil, zero value otherwise.

### GetTokenReferenceIdOk

`func (o *TokenServiceProvider) GetTokenReferenceIdOk() (*string, bool)`

GetTokenReferenceIdOk returns a tuple with the TokenReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenReferenceId

`func (o *TokenServiceProvider) SetTokenReferenceId(v string)`

SetTokenReferenceId sets TokenReferenceId field to given value.

### HasTokenReferenceId

`func (o *TokenServiceProvider) HasTokenReferenceId() bool`

HasTokenReferenceId returns a boolean if a field has been set.

### GetTokenRequestorId

`func (o *TokenServiceProvider) GetTokenRequestorId() string`

GetTokenRequestorId returns the TokenRequestorId field if non-nil, zero value otherwise.

### GetTokenRequestorIdOk

`func (o *TokenServiceProvider) GetTokenRequestorIdOk() (*string, bool)`

GetTokenRequestorIdOk returns a tuple with the TokenRequestorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenRequestorId

`func (o *TokenServiceProvider) SetTokenRequestorId(v string)`

SetTokenRequestorId sets TokenRequestorId field to given value.

### HasTokenRequestorId

`func (o *TokenServiceProvider) HasTokenRequestorId() bool`

HasTokenRequestorId returns a boolean if a field has been set.

### GetTokenRequestorName

`func (o *TokenServiceProvider) GetTokenRequestorName() string`

GetTokenRequestorName returns the TokenRequestorName field if non-nil, zero value otherwise.

### GetTokenRequestorNameOk

`func (o *TokenServiceProvider) GetTokenRequestorNameOk() (*string, bool)`

GetTokenRequestorNameOk returns a tuple with the TokenRequestorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenRequestorName

`func (o *TokenServiceProvider) SetTokenRequestorName(v string)`

SetTokenRequestorName sets TokenRequestorName field to given value.

### HasTokenRequestorName

`func (o *TokenServiceProvider) HasTokenRequestorName() bool`

HasTokenRequestorName returns a boolean if a field has been set.

### GetTokenScore

`func (o *TokenServiceProvider) GetTokenScore() string`

GetTokenScore returns the TokenScore field if non-nil, zero value otherwise.

### GetTokenScoreOk

`func (o *TokenServiceProvider) GetTokenScoreOk() (*string, bool)`

GetTokenScoreOk returns a tuple with the TokenScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenScore

`func (o *TokenServiceProvider) SetTokenScore(v string)`

SetTokenScore sets TokenScore field to given value.

### HasTokenScore

`func (o *TokenServiceProvider) HasTokenScore() bool`

HasTokenScore returns a boolean if a field has been set.

### GetTokenType

`func (o *TokenServiceProvider) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *TokenServiceProvider) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *TokenServiceProvider) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.

### HasTokenType

`func (o *TokenServiceProvider) HasTokenType() bool`

HasTokenType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


