# AchVerificationModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Indicates whether the ACH funding source is active. | [optional] [default to false]
**VerifyAmount1** | Pointer to **float32** | Verification amount. | [optional] 
**VerifyAmount2** | Pointer to **float32** | Verification amount. | [optional] 

## Methods

### NewAchVerificationModel

`func NewAchVerificationModel() *AchVerificationModel`

NewAchVerificationModel instantiates a new AchVerificationModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAchVerificationModelWithDefaults

`func NewAchVerificationModelWithDefaults() *AchVerificationModel`

NewAchVerificationModelWithDefaults instantiates a new AchVerificationModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *AchVerificationModel) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *AchVerificationModel) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *AchVerificationModel) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *AchVerificationModel) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetVerifyAmount1

`func (o *AchVerificationModel) GetVerifyAmount1() float32`

GetVerifyAmount1 returns the VerifyAmount1 field if non-nil, zero value otherwise.

### GetVerifyAmount1Ok

`func (o *AchVerificationModel) GetVerifyAmount1Ok() (*float32, bool)`

GetVerifyAmount1Ok returns a tuple with the VerifyAmount1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyAmount1

`func (o *AchVerificationModel) SetVerifyAmount1(v float32)`

SetVerifyAmount1 sets VerifyAmount1 field to given value.

### HasVerifyAmount1

`func (o *AchVerificationModel) HasVerifyAmount1() bool`

HasVerifyAmount1 returns a boolean if a field has been set.

### GetVerifyAmount2

`func (o *AchVerificationModel) GetVerifyAmount2() float32`

GetVerifyAmount2 returns the VerifyAmount2 field if non-nil, zero value otherwise.

### GetVerifyAmount2Ok

`func (o *AchVerificationModel) GetVerifyAmount2Ok() (*float32, bool)`

GetVerifyAmount2Ok returns a tuple with the VerifyAmount2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyAmount2

`func (o *AchVerificationModel) SetVerifyAmount2(v float32)`

SetVerifyAmount2 sets VerifyAmount2 field to given value.

### HasVerifyAmount2

`func (o *AchVerificationModel) HasVerifyAmount2() bool`

HasVerifyAmount2 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


