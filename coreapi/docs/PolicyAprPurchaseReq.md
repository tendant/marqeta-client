# PolicyAprPurchaseReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalToken** | Pointer to **string** | Unique identifier of the external pricing strategy for the credit program. | [optional] 
**Name** | Pointer to **string** | Name of the pricing strategy. | [optional] 
**Tiers** | [**[]PolicyAprTierReq**](PolicyAprTierReq.md) | One or more risk tiers for a pricing strategy. | 

## Methods

### NewPolicyAprPurchaseReq

`func NewPolicyAprPurchaseReq(tiers []PolicyAprTierReq, ) *PolicyAprPurchaseReq`

NewPolicyAprPurchaseReq instantiates a new PolicyAprPurchaseReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyAprPurchaseReqWithDefaults

`func NewPolicyAprPurchaseReqWithDefaults() *PolicyAprPurchaseReq`

NewPolicyAprPurchaseReqWithDefaults instantiates a new PolicyAprPurchaseReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalToken

`func (o *PolicyAprPurchaseReq) GetExternalToken() string`

GetExternalToken returns the ExternalToken field if non-nil, zero value otherwise.

### GetExternalTokenOk

`func (o *PolicyAprPurchaseReq) GetExternalTokenOk() (*string, bool)`

GetExternalTokenOk returns a tuple with the ExternalToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalToken

`func (o *PolicyAprPurchaseReq) SetExternalToken(v string)`

SetExternalToken sets ExternalToken field to given value.

### HasExternalToken

`func (o *PolicyAprPurchaseReq) HasExternalToken() bool`

HasExternalToken returns a boolean if a field has been set.

### GetName

`func (o *PolicyAprPurchaseReq) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PolicyAprPurchaseReq) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PolicyAprPurchaseReq) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PolicyAprPurchaseReq) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTiers

`func (o *PolicyAprPurchaseReq) GetTiers() []PolicyAprTierReq`

GetTiers returns the Tiers field if non-nil, zero value otherwise.

### GetTiersOk

`func (o *PolicyAprPurchaseReq) GetTiersOk() (*[]PolicyAprTierReq, bool)`

GetTiersOk returns a tuple with the Tiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiers

`func (o *PolicyAprPurchaseReq) SetTiers(v []PolicyAprTierReq)`

SetTiers sets Tiers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


