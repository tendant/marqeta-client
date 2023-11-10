# RealTimeFeeGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Indicates whether the real-time fee group is active. | [optional] [default to true]
**FeeTokens** | Pointer to **[]string** | Specifies the fees in this real-time fee group. Send a &#x60;GET&#x60; request to &#x60;/fees&#x60; to retrieve fee resources tokens.  No two fees in the group can be applicable to the same transaction type (in other words, each fee must have a different value for its &#x60;real_time_assessment.transaction_type&#x60; field). | [optional] 
**Name** | Pointer to **string** | Descriptive name for the real-time fee group. | [optional] 

## Methods

### NewRealTimeFeeGroupRequest

`func NewRealTimeFeeGroupRequest() *RealTimeFeeGroupRequest`

NewRealTimeFeeGroupRequest instantiates a new RealTimeFeeGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRealTimeFeeGroupRequestWithDefaults

`func NewRealTimeFeeGroupRequestWithDefaults() *RealTimeFeeGroupRequest`

NewRealTimeFeeGroupRequestWithDefaults instantiates a new RealTimeFeeGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *RealTimeFeeGroupRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *RealTimeFeeGroupRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *RealTimeFeeGroupRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *RealTimeFeeGroupRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetFeeTokens

`func (o *RealTimeFeeGroupRequest) GetFeeTokens() []string`

GetFeeTokens returns the FeeTokens field if non-nil, zero value otherwise.

### GetFeeTokensOk

`func (o *RealTimeFeeGroupRequest) GetFeeTokensOk() (*[]string, bool)`

GetFeeTokensOk returns a tuple with the FeeTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeTokens

`func (o *RealTimeFeeGroupRequest) SetFeeTokens(v []string)`

SetFeeTokens sets FeeTokens field to given value.

### HasFeeTokens

`func (o *RealTimeFeeGroupRequest) HasFeeTokens() bool`

HasFeeTokens returns a boolean if a field has been set.

### GetName

`func (o *RealTimeFeeGroupRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RealTimeFeeGroupRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RealTimeFeeGroupRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RealTimeFeeGroupRequest) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


