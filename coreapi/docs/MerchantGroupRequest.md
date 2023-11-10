# MerchantGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Indicates if the merchant group is active or not. | [optional] [default to false]
**Mids** | Pointer to **[]string** | Comma-separated list of alphanumeric merchant identifiers. You can include merchant identifiers in multiple merchant groups. | [optional] 
**Name** | **string** | Name of the merchant group. | 
**Token** | Pointer to **string** | Unique identifier of the group.  If you do not include a token, the system will generate one automatically. This token is necessary for use in other API calls, so we recommend that rather than let the system generate one, you use a simple string that is easy to remember. This value cannot be updated. | [optional] 

## Methods

### NewMerchantGroupRequest

`func NewMerchantGroupRequest(name string, ) *MerchantGroupRequest`

NewMerchantGroupRequest instantiates a new MerchantGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantGroupRequestWithDefaults

`func NewMerchantGroupRequestWithDefaults() *MerchantGroupRequest`

NewMerchantGroupRequestWithDefaults instantiates a new MerchantGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *MerchantGroupRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *MerchantGroupRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *MerchantGroupRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *MerchantGroupRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetMids

`func (o *MerchantGroupRequest) GetMids() []string`

GetMids returns the Mids field if non-nil, zero value otherwise.

### GetMidsOk

`func (o *MerchantGroupRequest) GetMidsOk() (*[]string, bool)`

GetMidsOk returns a tuple with the Mids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMids

`func (o *MerchantGroupRequest) SetMids(v []string)`

SetMids sets Mids field to given value.

### HasMids

`func (o *MerchantGroupRequest) HasMids() bool`

HasMids returns a boolean if a field has been set.

### GetName

`func (o *MerchantGroupRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MerchantGroupRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MerchantGroupRequest) SetName(v string)`

SetName sets Name field to given value.


### GetToken

`func (o *MerchantGroupRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *MerchantGroupRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *MerchantGroupRequest) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *MerchantGroupRequest) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


