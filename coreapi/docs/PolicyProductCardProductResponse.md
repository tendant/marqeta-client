# PolicyProductCardProductResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Level** | Pointer to [**PolicyProductCardProductLevel**](PolicyProductCardProductLevel.md) |  | [optional] 
**Network** | Pointer to **string** | Name of the card network. | [optional] 
**Token** | Pointer to **string** | Unique identifier of the card product. | [optional] 

## Methods

### NewPolicyProductCardProductResponse

`func NewPolicyProductCardProductResponse() *PolicyProductCardProductResponse`

NewPolicyProductCardProductResponse instantiates a new PolicyProductCardProductResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyProductCardProductResponseWithDefaults

`func NewPolicyProductCardProductResponseWithDefaults() *PolicyProductCardProductResponse`

NewPolicyProductCardProductResponseWithDefaults instantiates a new PolicyProductCardProductResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLevel

`func (o *PolicyProductCardProductResponse) GetLevel() PolicyProductCardProductLevel`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *PolicyProductCardProductResponse) GetLevelOk() (*PolicyProductCardProductLevel, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *PolicyProductCardProductResponse) SetLevel(v PolicyProductCardProductLevel)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *PolicyProductCardProductResponse) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetNetwork

`func (o *PolicyProductCardProductResponse) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *PolicyProductCardProductResponse) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *PolicyProductCardProductResponse) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *PolicyProductCardProductResponse) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetToken

`func (o *PolicyProductCardProductResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PolicyProductCardProductResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PolicyProductCardProductResponse) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *PolicyProductCardProductResponse) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


