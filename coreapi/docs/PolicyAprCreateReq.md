# PolicyAprCreateReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description of the APR policy. | [optional] 
**Name** | **string** | Name of the APR policy. | 
**Purchases** | [**PolicyAprPurchaseReq**](PolicyAprPurchaseReq.md) |  | 
**Token** | Pointer to **string** | Unique identifier of the APR policy. | [optional] 

## Methods

### NewPolicyAprCreateReq

`func NewPolicyAprCreateReq(name string, purchases PolicyAprPurchaseReq, ) *PolicyAprCreateReq`

NewPolicyAprCreateReq instantiates a new PolicyAprCreateReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyAprCreateReqWithDefaults

`func NewPolicyAprCreateReqWithDefaults() *PolicyAprCreateReq`

NewPolicyAprCreateReqWithDefaults instantiates a new PolicyAprCreateReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *PolicyAprCreateReq) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PolicyAprCreateReq) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PolicyAprCreateReq) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PolicyAprCreateReq) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *PolicyAprCreateReq) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PolicyAprCreateReq) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PolicyAprCreateReq) SetName(v string)`

SetName sets Name field to given value.


### GetPurchases

`func (o *PolicyAprCreateReq) GetPurchases() PolicyAprPurchaseReq`

GetPurchases returns the Purchases field if non-nil, zero value otherwise.

### GetPurchasesOk

`func (o *PolicyAprCreateReq) GetPurchasesOk() (*PolicyAprPurchaseReq, bool)`

GetPurchasesOk returns a tuple with the Purchases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchases

`func (o *PolicyAprCreateReq) SetPurchases(v PolicyAprPurchaseReq)`

SetPurchases sets Purchases field to given value.


### GetToken

`func (o *PolicyAprCreateReq) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PolicyAprCreateReq) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PolicyAprCreateReq) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *PolicyAprCreateReq) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


