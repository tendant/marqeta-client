# JitFundingProgramFundingSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Specifies whether JIT Funding is enabled or disabled for the program funding source. A value of &#x60;true&#x60; indicates that the program funding source is enabled and will be debited when swipes occur. | [optional] [default to false]
**FundingSourceToken** | Pointer to **string** | Unique identifier of the already existing funding source. Required if JIT Funding is enabled. | [optional] 
**RefundsDestination** | Pointer to **string** | Specifies the return destination for refunds in the case of a transaction reversal. &#x60;PROGRAM_FUNDING_SOURCE&#x60; returns funds to the program funding source. &#x60;GPA&#x60; returns the funds to the user&#39;s GPA. | [optional] 

## Methods

### NewJitFundingProgramFundingSource

`func NewJitFundingProgramFundingSource() *JitFundingProgramFundingSource`

NewJitFundingProgramFundingSource instantiates a new JitFundingProgramFundingSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJitFundingProgramFundingSourceWithDefaults

`func NewJitFundingProgramFundingSourceWithDefaults() *JitFundingProgramFundingSource`

NewJitFundingProgramFundingSourceWithDefaults instantiates a new JitFundingProgramFundingSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *JitFundingProgramFundingSource) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *JitFundingProgramFundingSource) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *JitFundingProgramFundingSource) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *JitFundingProgramFundingSource) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFundingSourceToken

`func (o *JitFundingProgramFundingSource) GetFundingSourceToken() string`

GetFundingSourceToken returns the FundingSourceToken field if non-nil, zero value otherwise.

### GetFundingSourceTokenOk

`func (o *JitFundingProgramFundingSource) GetFundingSourceTokenOk() (*string, bool)`

GetFundingSourceTokenOk returns a tuple with the FundingSourceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingSourceToken

`func (o *JitFundingProgramFundingSource) SetFundingSourceToken(v string)`

SetFundingSourceToken sets FundingSourceToken field to given value.

### HasFundingSourceToken

`func (o *JitFundingProgramFundingSource) HasFundingSourceToken() bool`

HasFundingSourceToken returns a boolean if a field has been set.

### GetRefundsDestination

`func (o *JitFundingProgramFundingSource) GetRefundsDestination() string`

GetRefundsDestination returns the RefundsDestination field if non-nil, zero value otherwise.

### GetRefundsDestinationOk

`func (o *JitFundingProgramFundingSource) GetRefundsDestinationOk() (*string, bool)`

GetRefundsDestinationOk returns a tuple with the RefundsDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundsDestination

`func (o *JitFundingProgramFundingSource) SetRefundsDestination(v string)`

SetRefundsDestination sets RefundsDestination field to given value.

### HasRefundsDestination

`func (o *JitFundingProgramFundingSource) HasRefundsDestination() bool`

HasRefundsDestination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


