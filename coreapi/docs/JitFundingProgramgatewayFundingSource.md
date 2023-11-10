# JitFundingProgramgatewayFundingSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlwaysFund** | Pointer to **bool** | If set to &#x60;true&#x60;, this card product is always funded from this program gateway funding source. | [optional] [default to false]
**Enabled** | Pointer to **bool** | Specifies whether JIT Funding is enabled or disabled for the program gateway funding source. A value of &#x60;true&#x60; indicates that the program gateway funding source is enabled and will be debited when swipes occur. | [optional] [default to false]
**FundingSourceToken** | Pointer to **string** | Unique identifier of the already existing funding source. Required if JIT Funding is enabled. | [optional] 
**RefundsDestination** | Pointer to **string** | Specifies the return destination for refunds in the case of a transaction reversal. In most cases, you should set the value to &#x60;GATEWAY&#x60;, which returns funds to the program gateway funding source. Setting to &#x60;GPA&#x60; returns the funds to the user&#39;s GPA, which creates a positive account balance and introduces the potential of a transaction being authorized without a JIT Funding request being sent to the gateway. | [optional] 

## Methods

### NewJitFundingProgramgatewayFundingSource

`func NewJitFundingProgramgatewayFundingSource() *JitFundingProgramgatewayFundingSource`

NewJitFundingProgramgatewayFundingSource instantiates a new JitFundingProgramgatewayFundingSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJitFundingProgramgatewayFundingSourceWithDefaults

`func NewJitFundingProgramgatewayFundingSourceWithDefaults() *JitFundingProgramgatewayFundingSource`

NewJitFundingProgramgatewayFundingSourceWithDefaults instantiates a new JitFundingProgramgatewayFundingSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlwaysFund

`func (o *JitFundingProgramgatewayFundingSource) GetAlwaysFund() bool`

GetAlwaysFund returns the AlwaysFund field if non-nil, zero value otherwise.

### GetAlwaysFundOk

`func (o *JitFundingProgramgatewayFundingSource) GetAlwaysFundOk() (*bool, bool)`

GetAlwaysFundOk returns a tuple with the AlwaysFund field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlwaysFund

`func (o *JitFundingProgramgatewayFundingSource) SetAlwaysFund(v bool)`

SetAlwaysFund sets AlwaysFund field to given value.

### HasAlwaysFund

`func (o *JitFundingProgramgatewayFundingSource) HasAlwaysFund() bool`

HasAlwaysFund returns a boolean if a field has been set.

### GetEnabled

`func (o *JitFundingProgramgatewayFundingSource) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *JitFundingProgramgatewayFundingSource) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *JitFundingProgramgatewayFundingSource) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *JitFundingProgramgatewayFundingSource) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFundingSourceToken

`func (o *JitFundingProgramgatewayFundingSource) GetFundingSourceToken() string`

GetFundingSourceToken returns the FundingSourceToken field if non-nil, zero value otherwise.

### GetFundingSourceTokenOk

`func (o *JitFundingProgramgatewayFundingSource) GetFundingSourceTokenOk() (*string, bool)`

GetFundingSourceTokenOk returns a tuple with the FundingSourceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingSourceToken

`func (o *JitFundingProgramgatewayFundingSource) SetFundingSourceToken(v string)`

SetFundingSourceToken sets FundingSourceToken field to given value.

### HasFundingSourceToken

`func (o *JitFundingProgramgatewayFundingSource) HasFundingSourceToken() bool`

HasFundingSourceToken returns a boolean if a field has been set.

### GetRefundsDestination

`func (o *JitFundingProgramgatewayFundingSource) GetRefundsDestination() string`

GetRefundsDestination returns the RefundsDestination field if non-nil, zero value otherwise.

### GetRefundsDestinationOk

`func (o *JitFundingProgramgatewayFundingSource) GetRefundsDestinationOk() (*string, bool)`

GetRefundsDestinationOk returns a tuple with the RefundsDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundsDestination

`func (o *JitFundingProgramgatewayFundingSource) SetRefundsDestination(v string)`

SetRefundsDestination sets RefundsDestination field to given value.

### HasRefundsDestination

`func (o *JitFundingProgramgatewayFundingSource) HasRefundsDestination() bool`

HasRefundsDestination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


