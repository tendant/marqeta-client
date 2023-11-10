# RealTimeStandinCriteria

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | If set to &#x60;true&#x60;, Commando Mode is automatically enabled by events defined in the &#x60;real_time_standin_criteria&#x60; object. If set to &#x60;false&#x60;, Auto Commando Mode is not enabled. | [optional] [default to false]
**IncludeApplicationErrors** | Pointer to **bool** | If set to &#x60;true&#x60;, an application error (any non-connection, non-timeout error) automatically enables Commando Mode when &#x60;real_time_standin_criteria.enabled&#x60; is also &#x60;true&#x60;. | [optional] [default to false]
**IncludeConnectionErrors** | Pointer to **bool** | If set to &#x60;true&#x60;, a non-timeout connection error automatically enables Commando Mode when &#x60;real_time_standin_criteria.enabled&#x60; is also &#x60;true&#x60;. | [optional] [default to false]
**IncludeResponseTimeouts** | Pointer to **bool** | If set to &#x60;true&#x60;, a gateway response slower than 3000ms automatically enables Commando Mode when &#x60;real_time_standin_criteria.enabled&#x60; is also &#x60;true&#x60;. | [optional] [default to false]

## Methods

### NewRealTimeStandinCriteria

`func NewRealTimeStandinCriteria() *RealTimeStandinCriteria`

NewRealTimeStandinCriteria instantiates a new RealTimeStandinCriteria object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRealTimeStandinCriteriaWithDefaults

`func NewRealTimeStandinCriteriaWithDefaults() *RealTimeStandinCriteria`

NewRealTimeStandinCriteriaWithDefaults instantiates a new RealTimeStandinCriteria object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *RealTimeStandinCriteria) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *RealTimeStandinCriteria) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *RealTimeStandinCriteria) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *RealTimeStandinCriteria) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetIncludeApplicationErrors

`func (o *RealTimeStandinCriteria) GetIncludeApplicationErrors() bool`

GetIncludeApplicationErrors returns the IncludeApplicationErrors field if non-nil, zero value otherwise.

### GetIncludeApplicationErrorsOk

`func (o *RealTimeStandinCriteria) GetIncludeApplicationErrorsOk() (*bool, bool)`

GetIncludeApplicationErrorsOk returns a tuple with the IncludeApplicationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeApplicationErrors

`func (o *RealTimeStandinCriteria) SetIncludeApplicationErrors(v bool)`

SetIncludeApplicationErrors sets IncludeApplicationErrors field to given value.

### HasIncludeApplicationErrors

`func (o *RealTimeStandinCriteria) HasIncludeApplicationErrors() bool`

HasIncludeApplicationErrors returns a boolean if a field has been set.

### GetIncludeConnectionErrors

`func (o *RealTimeStandinCriteria) GetIncludeConnectionErrors() bool`

GetIncludeConnectionErrors returns the IncludeConnectionErrors field if non-nil, zero value otherwise.

### GetIncludeConnectionErrorsOk

`func (o *RealTimeStandinCriteria) GetIncludeConnectionErrorsOk() (*bool, bool)`

GetIncludeConnectionErrorsOk returns a tuple with the IncludeConnectionErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeConnectionErrors

`func (o *RealTimeStandinCriteria) SetIncludeConnectionErrors(v bool)`

SetIncludeConnectionErrors sets IncludeConnectionErrors field to given value.

### HasIncludeConnectionErrors

`func (o *RealTimeStandinCriteria) HasIncludeConnectionErrors() bool`

HasIncludeConnectionErrors returns a boolean if a field has been set.

### GetIncludeResponseTimeouts

`func (o *RealTimeStandinCriteria) GetIncludeResponseTimeouts() bool`

GetIncludeResponseTimeouts returns the IncludeResponseTimeouts field if non-nil, zero value otherwise.

### GetIncludeResponseTimeoutsOk

`func (o *RealTimeStandinCriteria) GetIncludeResponseTimeoutsOk() (*bool, bool)`

GetIncludeResponseTimeoutsOk returns a tuple with the IncludeResponseTimeouts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeResponseTimeouts

`func (o *RealTimeStandinCriteria) SetIncludeResponseTimeouts(v bool)`

SetIncludeResponseTimeouts sets IncludeResponseTimeouts field to given value.

### HasIncludeResponseTimeouts

`func (o *RealTimeStandinCriteria) HasIncludeResponseTimeouts() bool`

HasIncludeResponseTimeouts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


