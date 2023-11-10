# AvsControls

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthMessages** | Pointer to [**AvsControlOptions**](AvsControlOptions.md) |  | [optional] 
**AvMessages** | Pointer to [**AvsControlOptions**](AvsControlOptions.md) |  | [optional] 

## Methods

### NewAvsControls

`func NewAvsControls() *AvsControls`

NewAvsControls instantiates a new AvsControls object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAvsControlsWithDefaults

`func NewAvsControlsWithDefaults() *AvsControls`

NewAvsControlsWithDefaults instantiates a new AvsControls object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthMessages

`func (o *AvsControls) GetAuthMessages() AvsControlOptions`

GetAuthMessages returns the AuthMessages field if non-nil, zero value otherwise.

### GetAuthMessagesOk

`func (o *AvsControls) GetAuthMessagesOk() (*AvsControlOptions, bool)`

GetAuthMessagesOk returns a tuple with the AuthMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMessages

`func (o *AvsControls) SetAuthMessages(v AvsControlOptions)`

SetAuthMessages sets AuthMessages field to given value.

### HasAuthMessages

`func (o *AvsControls) HasAuthMessages() bool`

HasAuthMessages returns a boolean if a field has been set.

### GetAvMessages

`func (o *AvsControls) GetAvMessages() AvsControlOptions`

GetAvMessages returns the AvMessages field if non-nil, zero value otherwise.

### GetAvMessagesOk

`func (o *AvsControls) GetAvMessagesOk() (*AvsControlOptions, bool)`

GetAvMessagesOk returns a tuple with the AvMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvMessages

`func (o *AvsControls) SetAvMessages(v AvsControlOptions)`

SetAvMessages sets AvMessages field to given value.

### HasAvMessages

`func (o *AvsControls) HasAvMessages() bool`

HasAvMessages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


