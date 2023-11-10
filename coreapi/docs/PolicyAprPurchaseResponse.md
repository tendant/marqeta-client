# PolicyAprPurchaseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalToken** | Pointer to **string** | Unique identifier of the pricing strategy on a credit program. | [optional] 
**Name** | Pointer to **string** | Name of the pricing strategy. | [optional] 
**Tiers** | Pointer to [**[]PolicyAprTierResponse**](PolicyAprTierResponse.md) | One or more risk tiers for a pricing strategy. | [optional] 

## Methods

### NewPolicyAprPurchaseResponse

`func NewPolicyAprPurchaseResponse() *PolicyAprPurchaseResponse`

NewPolicyAprPurchaseResponse instantiates a new PolicyAprPurchaseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyAprPurchaseResponseWithDefaults

`func NewPolicyAprPurchaseResponseWithDefaults() *PolicyAprPurchaseResponse`

NewPolicyAprPurchaseResponseWithDefaults instantiates a new PolicyAprPurchaseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalToken

`func (o *PolicyAprPurchaseResponse) GetExternalToken() string`

GetExternalToken returns the ExternalToken field if non-nil, zero value otherwise.

### GetExternalTokenOk

`func (o *PolicyAprPurchaseResponse) GetExternalTokenOk() (*string, bool)`

GetExternalTokenOk returns a tuple with the ExternalToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalToken

`func (o *PolicyAprPurchaseResponse) SetExternalToken(v string)`

SetExternalToken sets ExternalToken field to given value.

### HasExternalToken

`func (o *PolicyAprPurchaseResponse) HasExternalToken() bool`

HasExternalToken returns a boolean if a field has been set.

### GetName

`func (o *PolicyAprPurchaseResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PolicyAprPurchaseResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PolicyAprPurchaseResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PolicyAprPurchaseResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTiers

`func (o *PolicyAprPurchaseResponse) GetTiers() []PolicyAprTierResponse`

GetTiers returns the Tiers field if non-nil, zero value otherwise.

### GetTiersOk

`func (o *PolicyAprPurchaseResponse) GetTiersOk() (*[]PolicyAprTierResponse, bool)`

GetTiersOk returns a tuple with the Tiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiers

`func (o *PolicyAprPurchaseResponse) SetTiers(v []PolicyAprTierResponse)`

SetTiers sets Tiers field to given value.

### HasTiers

`func (o *PolicyAprPurchaseResponse) HasTiers() bool`

HasTiers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


