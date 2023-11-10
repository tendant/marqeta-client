# Gpa

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReloadAmount** | **float32** | Available balance on the card after the reload has completed.  This value must be greater than or equal to the value of &#x60;trigger_amount&#x60;. Note that this is not the same as the amount added to the card, which will vary from reload to reload. | 
**TriggerAmount** | **float32** | Threshold that determines when the reload happens.  The reload is triggered when the card balance falls below this amount. | 

## Methods

### NewGpa

`func NewGpa(reloadAmount float32, triggerAmount float32, ) *Gpa`

NewGpa instantiates a new Gpa object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGpaWithDefaults

`func NewGpaWithDefaults() *Gpa`

NewGpaWithDefaults instantiates a new Gpa object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReloadAmount

`func (o *Gpa) GetReloadAmount() float32`

GetReloadAmount returns the ReloadAmount field if non-nil, zero value otherwise.

### GetReloadAmountOk

`func (o *Gpa) GetReloadAmountOk() (*float32, bool)`

GetReloadAmountOk returns a tuple with the ReloadAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReloadAmount

`func (o *Gpa) SetReloadAmount(v float32)`

SetReloadAmount sets ReloadAmount field to given value.


### GetTriggerAmount

`func (o *Gpa) GetTriggerAmount() float32`

GetTriggerAmount returns the TriggerAmount field if non-nil, zero value otherwise.

### GetTriggerAmountOk

`func (o *Gpa) GetTriggerAmountOk() (*float32, bool)`

GetTriggerAmountOk returns a tuple with the TriggerAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerAmount

`func (o *Gpa) SetTriggerAmount(v float32)`

SetTriggerAmount sets TriggerAmount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


