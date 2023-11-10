# RedemptionTransitionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedTime** | **time.Time** | Date and time when the reward redemption transition was created on Marqeta&#39;s credit platform. | 
**ExternalSettlementDateTime** | Pointer to **time.Time** | Date and time when the reward redemption was settled on your external platform.  This field is returned if you handled the reward redemption outside of Marqeta&#39;s credit platform. | [optional] 
**InitialStatus** | [**RedemptionStatus**](RedemptionStatus.md) |  | 
**NewStatus** | [**RedemptionStatus**](RedemptionStatus.md) |  | 
**RedemptionToken** | **string** | Unique identifier of the redemption whose status was transitioned. | 
**Token** | **string** | Unique identifier of the reward redemption transition. | 

## Methods

### NewRedemptionTransitionsResponse

`func NewRedemptionTransitionsResponse(createdTime time.Time, initialStatus RedemptionStatus, newStatus RedemptionStatus, redemptionToken string, token string, ) *RedemptionTransitionsResponse`

NewRedemptionTransitionsResponse instantiates a new RedemptionTransitionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRedemptionTransitionsResponseWithDefaults

`func NewRedemptionTransitionsResponseWithDefaults() *RedemptionTransitionsResponse`

NewRedemptionTransitionsResponseWithDefaults instantiates a new RedemptionTransitionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedTime

`func (o *RedemptionTransitionsResponse) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *RedemptionTransitionsResponse) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *RedemptionTransitionsResponse) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetExternalSettlementDateTime

`func (o *RedemptionTransitionsResponse) GetExternalSettlementDateTime() time.Time`

GetExternalSettlementDateTime returns the ExternalSettlementDateTime field if non-nil, zero value otherwise.

### GetExternalSettlementDateTimeOk

`func (o *RedemptionTransitionsResponse) GetExternalSettlementDateTimeOk() (*time.Time, bool)`

GetExternalSettlementDateTimeOk returns a tuple with the ExternalSettlementDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSettlementDateTime

`func (o *RedemptionTransitionsResponse) SetExternalSettlementDateTime(v time.Time)`

SetExternalSettlementDateTime sets ExternalSettlementDateTime field to given value.

### HasExternalSettlementDateTime

`func (o *RedemptionTransitionsResponse) HasExternalSettlementDateTime() bool`

HasExternalSettlementDateTime returns a boolean if a field has been set.

### GetInitialStatus

`func (o *RedemptionTransitionsResponse) GetInitialStatus() RedemptionStatus`

GetInitialStatus returns the InitialStatus field if non-nil, zero value otherwise.

### GetInitialStatusOk

`func (o *RedemptionTransitionsResponse) GetInitialStatusOk() (*RedemptionStatus, bool)`

GetInitialStatusOk returns a tuple with the InitialStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialStatus

`func (o *RedemptionTransitionsResponse) SetInitialStatus(v RedemptionStatus)`

SetInitialStatus sets InitialStatus field to given value.


### GetNewStatus

`func (o *RedemptionTransitionsResponse) GetNewStatus() RedemptionStatus`

GetNewStatus returns the NewStatus field if non-nil, zero value otherwise.

### GetNewStatusOk

`func (o *RedemptionTransitionsResponse) GetNewStatusOk() (*RedemptionStatus, bool)`

GetNewStatusOk returns a tuple with the NewStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewStatus

`func (o *RedemptionTransitionsResponse) SetNewStatus(v RedemptionStatus)`

SetNewStatus sets NewStatus field to given value.


### GetRedemptionToken

`func (o *RedemptionTransitionsResponse) GetRedemptionToken() string`

GetRedemptionToken returns the RedemptionToken field if non-nil, zero value otherwise.

### GetRedemptionTokenOk

`func (o *RedemptionTransitionsResponse) GetRedemptionTokenOk() (*string, bool)`

GetRedemptionTokenOk returns a tuple with the RedemptionToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedemptionToken

`func (o *RedemptionTransitionsResponse) SetRedemptionToken(v string)`

SetRedemptionToken sets RedemptionToken field to given value.


### GetToken

`func (o *RedemptionTransitionsResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *RedemptionTransitionsResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *RedemptionTransitionsResponse) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


