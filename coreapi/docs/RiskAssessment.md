# RiskAssessment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Score** | Pointer to **string** | Wallet provider&#39;s decision as to whether the digital wallet token should be provisioned. | [optional] 
**Version** | Pointer to **string** | Wallet provider&#39;s risk assessment version. | [optional] 

## Methods

### NewRiskAssessment

`func NewRiskAssessment() *RiskAssessment`

NewRiskAssessment instantiates a new RiskAssessment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskAssessmentWithDefaults

`func NewRiskAssessmentWithDefaults() *RiskAssessment`

NewRiskAssessmentWithDefaults instantiates a new RiskAssessment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScore

`func (o *RiskAssessment) GetScore() string`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *RiskAssessment) GetScoreOk() (*string, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *RiskAssessment) SetScore(v string)`

SetScore sets Score field to given value.

### HasScore

`func (o *RiskAssessment) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetVersion

`func (o *RiskAssessment) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *RiskAssessment) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *RiskAssessment) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *RiskAssessment) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


