# AuthControlMerchantScope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mcc** | Pointer to **string** | Merchant Category Code (MCC). Identifies the type of goods or services provided by the merchant.  Enter a value to control access to a particular type of product or service.  See &lt;&lt;/developer-guides/controlling-spending, Controlling Spending&gt;&gt; for links to more information about merchant category codes. | [optional] 
**MccGroup** | Pointer to **string** | Token identifying a group of MCCs.  Enter a value to control access to a group of product or service types. | [optional] 
**MerchantGroupToken** | Pointer to **string** | Unique identifier of a merchant group.  Enter a value to control access to a group of merchants. | [optional] 
**Mid** | Pointer to **string** | Merchant identifier (MID). Identifies a specific merchant.  Enter a value to control access to a particular merchant. | [optional] 

## Methods

### NewAuthControlMerchantScope

`func NewAuthControlMerchantScope() *AuthControlMerchantScope`

NewAuthControlMerchantScope instantiates a new AuthControlMerchantScope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthControlMerchantScopeWithDefaults

`func NewAuthControlMerchantScopeWithDefaults() *AuthControlMerchantScope`

NewAuthControlMerchantScopeWithDefaults instantiates a new AuthControlMerchantScope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMcc

`func (o *AuthControlMerchantScope) GetMcc() string`

GetMcc returns the Mcc field if non-nil, zero value otherwise.

### GetMccOk

`func (o *AuthControlMerchantScope) GetMccOk() (*string, bool)`

GetMccOk returns a tuple with the Mcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcc

`func (o *AuthControlMerchantScope) SetMcc(v string)`

SetMcc sets Mcc field to given value.

### HasMcc

`func (o *AuthControlMerchantScope) HasMcc() bool`

HasMcc returns a boolean if a field has been set.

### GetMccGroup

`func (o *AuthControlMerchantScope) GetMccGroup() string`

GetMccGroup returns the MccGroup field if non-nil, zero value otherwise.

### GetMccGroupOk

`func (o *AuthControlMerchantScope) GetMccGroupOk() (*string, bool)`

GetMccGroupOk returns a tuple with the MccGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMccGroup

`func (o *AuthControlMerchantScope) SetMccGroup(v string)`

SetMccGroup sets MccGroup field to given value.

### HasMccGroup

`func (o *AuthControlMerchantScope) HasMccGroup() bool`

HasMccGroup returns a boolean if a field has been set.

### GetMerchantGroupToken

`func (o *AuthControlMerchantScope) GetMerchantGroupToken() string`

GetMerchantGroupToken returns the MerchantGroupToken field if non-nil, zero value otherwise.

### GetMerchantGroupTokenOk

`func (o *AuthControlMerchantScope) GetMerchantGroupTokenOk() (*string, bool)`

GetMerchantGroupTokenOk returns a tuple with the MerchantGroupToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantGroupToken

`func (o *AuthControlMerchantScope) SetMerchantGroupToken(v string)`

SetMerchantGroupToken sets MerchantGroupToken field to given value.

### HasMerchantGroupToken

`func (o *AuthControlMerchantScope) HasMerchantGroupToken() bool`

HasMerchantGroupToken returns a boolean if a field has been set.

### GetMid

`func (o *AuthControlMerchantScope) GetMid() string`

GetMid returns the Mid field if non-nil, zero value otherwise.

### GetMidOk

`func (o *AuthControlMerchantScope) GetMidOk() (*string, bool)`

GetMidOk returns a tuple with the Mid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMid

`func (o *AuthControlMerchantScope) SetMid(v string)`

SetMid sets Mid field to given value.

### HasMid

`func (o *AuthControlMerchantScope) HasMid() bool`

HasMid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


