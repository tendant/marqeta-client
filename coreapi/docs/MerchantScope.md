# MerchantScope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mcc** | Pointer to **string** | Merchant Category Code (MCC). Identifies the type of products or services provided by the merchant.  Enter a value to control spending on a particular type of product or service. | [optional] 
**MccGroup** | Pointer to **string** | Token identifying a group of MCCs. Enter a value to control spending on a group of product or service types.  Send a &#x60;GET&#x60; request to &#x60;/mccgroups&#x60; to retrieve MCC group tokens. | [optional] 
**Mid** | Pointer to **string** | Merchant identifier (MID). Identifies a specific merchant.  Enter a value to control spending with a particular merchant. | [optional] 

## Methods

### NewMerchantScope

`func NewMerchantScope() *MerchantScope`

NewMerchantScope instantiates a new MerchantScope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantScopeWithDefaults

`func NewMerchantScopeWithDefaults() *MerchantScope`

NewMerchantScopeWithDefaults instantiates a new MerchantScope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMcc

`func (o *MerchantScope) GetMcc() string`

GetMcc returns the Mcc field if non-nil, zero value otherwise.

### GetMccOk

`func (o *MerchantScope) GetMccOk() (*string, bool)`

GetMccOk returns a tuple with the Mcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcc

`func (o *MerchantScope) SetMcc(v string)`

SetMcc sets Mcc field to given value.

### HasMcc

`func (o *MerchantScope) HasMcc() bool`

HasMcc returns a boolean if a field has been set.

### GetMccGroup

`func (o *MerchantScope) GetMccGroup() string`

GetMccGroup returns the MccGroup field if non-nil, zero value otherwise.

### GetMccGroupOk

`func (o *MerchantScope) GetMccGroupOk() (*string, bool)`

GetMccGroupOk returns a tuple with the MccGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMccGroup

`func (o *MerchantScope) SetMccGroup(v string)`

SetMccGroup sets MccGroup field to given value.

### HasMccGroup

`func (o *MerchantScope) HasMccGroup() bool`

HasMccGroup returns a boolean if a field has been set.

### GetMid

`func (o *MerchantScope) GetMid() string`

GetMid returns the Mid field if non-nil, zero value otherwise.

### GetMidOk

`func (o *MerchantScope) GetMidOk() (*string, bool)`

GetMidOk returns a tuple with the Mid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMid

`func (o *MerchantScope) SetMid(v string)`

SetMid sets Mid field to given value.

### HasMid

`func (o *MerchantScope) HasMid() bool`

HasMid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


