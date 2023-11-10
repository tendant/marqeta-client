# PolicyProductCardProductReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Level** | [**PolicyProductCardProductLevel**](PolicyProductCardProductLevel.md) |  | 
**Token** | **string** | Unique identifier of the card product. | 

## Methods

### NewPolicyProductCardProductReq

`func NewPolicyProductCardProductReq(level PolicyProductCardProductLevel, token string, ) *PolicyProductCardProductReq`

NewPolicyProductCardProductReq instantiates a new PolicyProductCardProductReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyProductCardProductReqWithDefaults

`func NewPolicyProductCardProductReqWithDefaults() *PolicyProductCardProductReq`

NewPolicyProductCardProductReqWithDefaults instantiates a new PolicyProductCardProductReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLevel

`func (o *PolicyProductCardProductReq) GetLevel() PolicyProductCardProductLevel`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *PolicyProductCardProductReq) GetLevelOk() (*PolicyProductCardProductLevel, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *PolicyProductCardProductReq) SetLevel(v PolicyProductCardProductLevel)`

SetLevel sets Level field to given value.


### GetToken

`func (o *PolicyProductCardProductReq) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PolicyProductCardProductReq) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PolicyProductCardProductReq) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


