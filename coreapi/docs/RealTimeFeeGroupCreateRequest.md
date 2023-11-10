# RealTimeFeeGroupCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Indicates whether the real-time fee group is active. | [optional] [default to true]
**FeeTokens** | Pointer to **[]string** | Unique identifiers of the fees in this real-time fee group. Send a &#x60;GET&#x60; request to &#x60;/fees&#x60; to retrieve fee resource tokens.  No two fees in the group can be applicable to the same transaction type (in other words, each fee must have a different value for its &#x60;real_time_assessment.transaction_type&#x60; field). | [optional] 
**Name** | **string** | Descriptive name for the real-time fee group. | 
**Token** | Pointer to **string** | Unique identifier of the real-time fee group.  If you do not include a token, the system will generate one automatically. This token is necessary for use in other API calls, so we recommend that rather than let the system generate one, you use a simple string that is easy to remember. This value cannot be updated. | [optional] 

## Methods

### NewRealTimeFeeGroupCreateRequest

`func NewRealTimeFeeGroupCreateRequest(name string, ) *RealTimeFeeGroupCreateRequest`

NewRealTimeFeeGroupCreateRequest instantiates a new RealTimeFeeGroupCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRealTimeFeeGroupCreateRequestWithDefaults

`func NewRealTimeFeeGroupCreateRequestWithDefaults() *RealTimeFeeGroupCreateRequest`

NewRealTimeFeeGroupCreateRequestWithDefaults instantiates a new RealTimeFeeGroupCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *RealTimeFeeGroupCreateRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *RealTimeFeeGroupCreateRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *RealTimeFeeGroupCreateRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *RealTimeFeeGroupCreateRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetFeeTokens

`func (o *RealTimeFeeGroupCreateRequest) GetFeeTokens() []string`

GetFeeTokens returns the FeeTokens field if non-nil, zero value otherwise.

### GetFeeTokensOk

`func (o *RealTimeFeeGroupCreateRequest) GetFeeTokensOk() (*[]string, bool)`

GetFeeTokensOk returns a tuple with the FeeTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeTokens

`func (o *RealTimeFeeGroupCreateRequest) SetFeeTokens(v []string)`

SetFeeTokens sets FeeTokens field to given value.

### HasFeeTokens

`func (o *RealTimeFeeGroupCreateRequest) HasFeeTokens() bool`

HasFeeTokens returns a boolean if a field has been set.

### GetName

`func (o *RealTimeFeeGroupCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RealTimeFeeGroupCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RealTimeFeeGroupCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetToken

`func (o *RealTimeFeeGroupCreateRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *RealTimeFeeGroupCreateRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *RealTimeFeeGroupCreateRequest) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *RealTimeFeeGroupCreateRequest) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


