# RealTimeFeeAssessment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomesticEnabled** | Pointer to **bool** | Enables fee assessments where the origin of the transaction acquirer is inside the US. | [optional] [default to false]
**InternationalEnabled** | Pointer to **bool** | Enables fee assessments where the origin of the transaction acquirer is outside the US. | [optional] [default to false]
**TransactionType** | Pointer to **string** | Indicates the type of transactions on which the fee is assessed. | [optional] 

## Methods

### NewRealTimeFeeAssessment

`func NewRealTimeFeeAssessment() *RealTimeFeeAssessment`

NewRealTimeFeeAssessment instantiates a new RealTimeFeeAssessment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRealTimeFeeAssessmentWithDefaults

`func NewRealTimeFeeAssessmentWithDefaults() *RealTimeFeeAssessment`

NewRealTimeFeeAssessmentWithDefaults instantiates a new RealTimeFeeAssessment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomesticEnabled

`func (o *RealTimeFeeAssessment) GetDomesticEnabled() bool`

GetDomesticEnabled returns the DomesticEnabled field if non-nil, zero value otherwise.

### GetDomesticEnabledOk

`func (o *RealTimeFeeAssessment) GetDomesticEnabledOk() (*bool, bool)`

GetDomesticEnabledOk returns a tuple with the DomesticEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomesticEnabled

`func (o *RealTimeFeeAssessment) SetDomesticEnabled(v bool)`

SetDomesticEnabled sets DomesticEnabled field to given value.

### HasDomesticEnabled

`func (o *RealTimeFeeAssessment) HasDomesticEnabled() bool`

HasDomesticEnabled returns a boolean if a field has been set.

### GetInternationalEnabled

`func (o *RealTimeFeeAssessment) GetInternationalEnabled() bool`

GetInternationalEnabled returns the InternationalEnabled field if non-nil, zero value otherwise.

### GetInternationalEnabledOk

`func (o *RealTimeFeeAssessment) GetInternationalEnabledOk() (*bool, bool)`

GetInternationalEnabledOk returns a tuple with the InternationalEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternationalEnabled

`func (o *RealTimeFeeAssessment) SetInternationalEnabled(v bool)`

SetInternationalEnabled sets InternationalEnabled field to given value.

### HasInternationalEnabled

`func (o *RealTimeFeeAssessment) HasInternationalEnabled() bool`

HasInternationalEnabled returns a boolean if a field has been set.

### GetTransactionType

`func (o *RealTimeFeeAssessment) GetTransactionType() string`

GetTransactionType returns the TransactionType field if non-nil, zero value otherwise.

### GetTransactionTypeOk

`func (o *RealTimeFeeAssessment) GetTransactionTypeOk() (*string, bool)`

GetTransactionTypeOk returns a tuple with the TransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionType

`func (o *RealTimeFeeAssessment) SetTransactionType(v string)`

SetTransactionType sets TransactionType field to given value.

### HasTransactionType

`func (o *RealTimeFeeAssessment) HasTransactionType() bool`

HasTransactionType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


