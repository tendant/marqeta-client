# FraudView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IssuerProcessor** | Pointer to [**IssuerFraudView**](IssuerFraudView.md) |  | [optional] 
**Network** | Pointer to [**NetworkFraudView**](NetworkFraudView.md) |  | [optional] 
**NetworkAccountIntelligenceScore** | Pointer to [**NetworkAccountIntelligenceScore**](NetworkAccountIntelligenceScore.md) |  | [optional] 

## Methods

### NewFraudView

`func NewFraudView() *FraudView`

NewFraudView instantiates a new FraudView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFraudViewWithDefaults

`func NewFraudViewWithDefaults() *FraudView`

NewFraudViewWithDefaults instantiates a new FraudView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssuerProcessor

`func (o *FraudView) GetIssuerProcessor() IssuerFraudView`

GetIssuerProcessor returns the IssuerProcessor field if non-nil, zero value otherwise.

### GetIssuerProcessorOk

`func (o *FraudView) GetIssuerProcessorOk() (*IssuerFraudView, bool)`

GetIssuerProcessorOk returns a tuple with the IssuerProcessor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerProcessor

`func (o *FraudView) SetIssuerProcessor(v IssuerFraudView)`

SetIssuerProcessor sets IssuerProcessor field to given value.

### HasIssuerProcessor

`func (o *FraudView) HasIssuerProcessor() bool`

HasIssuerProcessor returns a boolean if a field has been set.

### GetNetwork

`func (o *FraudView) GetNetwork() NetworkFraudView`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *FraudView) GetNetworkOk() (*NetworkFraudView, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *FraudView) SetNetwork(v NetworkFraudView)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *FraudView) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetNetworkAccountIntelligenceScore

`func (o *FraudView) GetNetworkAccountIntelligenceScore() NetworkAccountIntelligenceScore`

GetNetworkAccountIntelligenceScore returns the NetworkAccountIntelligenceScore field if non-nil, zero value otherwise.

### GetNetworkAccountIntelligenceScoreOk

`func (o *FraudView) GetNetworkAccountIntelligenceScoreOk() (*NetworkAccountIntelligenceScore, bool)`

GetNetworkAccountIntelligenceScoreOk returns a tuple with the NetworkAccountIntelligenceScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkAccountIntelligenceScore

`func (o *FraudView) SetNetworkAccountIntelligenceScore(v NetworkAccountIntelligenceScore)`

SetNetworkAccountIntelligenceScore sets NetworkAccountIntelligenceScore field to given value.

### HasNetworkAccountIntelligenceScore

`func (o *FraudView) HasNetworkAccountIntelligenceScore() bool`

HasNetworkAccountIntelligenceScore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


