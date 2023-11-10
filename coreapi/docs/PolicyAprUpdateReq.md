# PolicyAprUpdateReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description of the APR policy. | [optional] 
**Name** | **string** | Name of the APR policy. | 
**Purchases** | Pointer to [**PolicyAprPurchaseReq**](PolicyAprPurchaseReq.md) |  | [optional] 
**Token** | Pointer to **string** | Unique identifier of the APR policy. | [optional] 

## Methods

### NewPolicyAprUpdateReq

`func NewPolicyAprUpdateReq(name string, ) *PolicyAprUpdateReq`

NewPolicyAprUpdateReq instantiates a new PolicyAprUpdateReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyAprUpdateReqWithDefaults

`func NewPolicyAprUpdateReqWithDefaults() *PolicyAprUpdateReq`

NewPolicyAprUpdateReqWithDefaults instantiates a new PolicyAprUpdateReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *PolicyAprUpdateReq) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PolicyAprUpdateReq) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PolicyAprUpdateReq) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PolicyAprUpdateReq) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *PolicyAprUpdateReq) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PolicyAprUpdateReq) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PolicyAprUpdateReq) SetName(v string)`

SetName sets Name field to given value.


### GetPurchases

`func (o *PolicyAprUpdateReq) GetPurchases() PolicyAprPurchaseReq`

GetPurchases returns the Purchases field if non-nil, zero value otherwise.

### GetPurchasesOk

`func (o *PolicyAprUpdateReq) GetPurchasesOk() (*PolicyAprPurchaseReq, bool)`

GetPurchasesOk returns a tuple with the Purchases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchases

`func (o *PolicyAprUpdateReq) SetPurchases(v PolicyAprPurchaseReq)`

SetPurchases sets Purchases field to given value.

### HasPurchases

`func (o *PolicyAprUpdateReq) HasPurchases() bool`

HasPurchases returns a boolean if a field has been set.

### GetToken

`func (o *PolicyAprUpdateReq) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PolicyAprUpdateReq) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PolicyAprUpdateReq) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *PolicyAprUpdateReq) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


