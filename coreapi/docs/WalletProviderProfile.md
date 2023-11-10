# WalletProviderProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to [**Account**](Account.md) |  | [optional] 
**DeviceScore** | Pointer to **string** | Score from the device. | [optional] 
**PanSource** | Pointer to **string** | Source from which the digital wallet provider obtained the card primary account number (PAN). | [optional] 
**ReasonCode** | Pointer to **string** | Reason for the wallet provider&#39;s provisioning decision.  * *01*  Cardholder&#39;s wallet account is too new relative to launch. * *02*  Cardholder&#39;s wallet account is too new relative to provisioning request. * *03*  Cardholder&#39;s wallet account/card pair is newer than date threshold. * *04*  Changes made to account data within the account threshold. * *05*  Suspicious transactions linked to this account. * *06*  Account has not had activity in the last year. * *07*  Suspended cards in the secure element. * *08*  Device was put in lost mode in the last seven days for longer than the duration threshold. * *09*  The number of provisioning attempts on this device in 24 hours exceeds threshold. * *0A*  There have been more than the threshold number of different cards attempted at provisioning to this phone in 24 hours. * *0B*  The card provisioning attempt contains a distinct name in excess of the threshold. * *0C*  The device score is less than 3. * *0D*  The account score is less than 4. * *0E*  Device provisioning location outside of the cardholder&#39;s wallet account home country. * *0G*  Suspect fraud. | [optional] 
**RecommendationReasons** | Pointer to **[]string** | Array of recommendation reasons from the digital wallet provider. | [optional] 
**RiskAssessment** | Pointer to [**RiskAssessment**](RiskAssessment.md) |  | [optional] 

## Methods

### NewWalletProviderProfile

`func NewWalletProviderProfile() *WalletProviderProfile`

NewWalletProviderProfile instantiates a new WalletProviderProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWalletProviderProfileWithDefaults

`func NewWalletProviderProfileWithDefaults() *WalletProviderProfile`

NewWalletProviderProfileWithDefaults instantiates a new WalletProviderProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *WalletProviderProfile) GetAccount() Account`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *WalletProviderProfile) GetAccountOk() (*Account, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *WalletProviderProfile) SetAccount(v Account)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *WalletProviderProfile) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetDeviceScore

`func (o *WalletProviderProfile) GetDeviceScore() string`

GetDeviceScore returns the DeviceScore field if non-nil, zero value otherwise.

### GetDeviceScoreOk

`func (o *WalletProviderProfile) GetDeviceScoreOk() (*string, bool)`

GetDeviceScoreOk returns a tuple with the DeviceScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceScore

`func (o *WalletProviderProfile) SetDeviceScore(v string)`

SetDeviceScore sets DeviceScore field to given value.

### HasDeviceScore

`func (o *WalletProviderProfile) HasDeviceScore() bool`

HasDeviceScore returns a boolean if a field has been set.

### GetPanSource

`func (o *WalletProviderProfile) GetPanSource() string`

GetPanSource returns the PanSource field if non-nil, zero value otherwise.

### GetPanSourceOk

`func (o *WalletProviderProfile) GetPanSourceOk() (*string, bool)`

GetPanSourceOk returns a tuple with the PanSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPanSource

`func (o *WalletProviderProfile) SetPanSource(v string)`

SetPanSource sets PanSource field to given value.

### HasPanSource

`func (o *WalletProviderProfile) HasPanSource() bool`

HasPanSource returns a boolean if a field has been set.

### GetReasonCode

`func (o *WalletProviderProfile) GetReasonCode() string`

GetReasonCode returns the ReasonCode field if non-nil, zero value otherwise.

### GetReasonCodeOk

`func (o *WalletProviderProfile) GetReasonCodeOk() (*string, bool)`

GetReasonCodeOk returns a tuple with the ReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonCode

`func (o *WalletProviderProfile) SetReasonCode(v string)`

SetReasonCode sets ReasonCode field to given value.

### HasReasonCode

`func (o *WalletProviderProfile) HasReasonCode() bool`

HasReasonCode returns a boolean if a field has been set.

### GetRecommendationReasons

`func (o *WalletProviderProfile) GetRecommendationReasons() []string`

GetRecommendationReasons returns the RecommendationReasons field if non-nil, zero value otherwise.

### GetRecommendationReasonsOk

`func (o *WalletProviderProfile) GetRecommendationReasonsOk() (*[]string, bool)`

GetRecommendationReasonsOk returns a tuple with the RecommendationReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendationReasons

`func (o *WalletProviderProfile) SetRecommendationReasons(v []string)`

SetRecommendationReasons sets RecommendationReasons field to given value.

### HasRecommendationReasons

`func (o *WalletProviderProfile) HasRecommendationReasons() bool`

HasRecommendationReasons returns a boolean if a field has been set.

### GetRiskAssessment

`func (o *WalletProviderProfile) GetRiskAssessment() RiskAssessment`

GetRiskAssessment returns the RiskAssessment field if non-nil, zero value otherwise.

### GetRiskAssessmentOk

`func (o *WalletProviderProfile) GetRiskAssessmentOk() (*RiskAssessment, bool)`

GetRiskAssessmentOk returns a tuple with the RiskAssessment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskAssessment

`func (o *WalletProviderProfile) SetRiskAssessment(v RiskAssessment)`

SetRiskAssessment sets RiskAssessment field to given value.

### HasRiskAssessment

`func (o *WalletProviderProfile) HasRiskAssessment() bool`

HasRiskAssessment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


