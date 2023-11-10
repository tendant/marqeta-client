# CreateRedemptionTransitionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalSettlementDateTime** | Pointer to **time.Time** | Date and time when the reward redemption was settled on your external platform.  Pass this field if you handle the reward redemption outside of Marqeta&#39;s credit platform. | [optional] 
**NewState** | [**RedemptionStatus**](RedemptionStatus.md) |  | 

## Methods

### NewCreateRedemptionTransitionsRequest

`func NewCreateRedemptionTransitionsRequest(newState RedemptionStatus, ) *CreateRedemptionTransitionsRequest`

NewCreateRedemptionTransitionsRequest instantiates a new CreateRedemptionTransitionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRedemptionTransitionsRequestWithDefaults

`func NewCreateRedemptionTransitionsRequestWithDefaults() *CreateRedemptionTransitionsRequest`

NewCreateRedemptionTransitionsRequestWithDefaults instantiates a new CreateRedemptionTransitionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalSettlementDateTime

`func (o *CreateRedemptionTransitionsRequest) GetExternalSettlementDateTime() time.Time`

GetExternalSettlementDateTime returns the ExternalSettlementDateTime field if non-nil, zero value otherwise.

### GetExternalSettlementDateTimeOk

`func (o *CreateRedemptionTransitionsRequest) GetExternalSettlementDateTimeOk() (*time.Time, bool)`

GetExternalSettlementDateTimeOk returns a tuple with the ExternalSettlementDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSettlementDateTime

`func (o *CreateRedemptionTransitionsRequest) SetExternalSettlementDateTime(v time.Time)`

SetExternalSettlementDateTime sets ExternalSettlementDateTime field to given value.

### HasExternalSettlementDateTime

`func (o *CreateRedemptionTransitionsRequest) HasExternalSettlementDateTime() bool`

HasExternalSettlementDateTime returns a boolean if a field has been set.

### GetNewState

`func (o *CreateRedemptionTransitionsRequest) GetNewState() RedemptionStatus`

GetNewState returns the NewState field if non-nil, zero value otherwise.

### GetNewStateOk

`func (o *CreateRedemptionTransitionsRequest) GetNewStateOk() (*RedemptionStatus, bool)`

GetNewStateOk returns a tuple with the NewState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewState

`func (o *CreateRedemptionTransitionsRequest) SetNewState(v RedemptionStatus)`

SetNewState sets NewState field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


