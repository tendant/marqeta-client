# MerchantGroupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Indicates if the merchant group is active or not. | [optional] [default to false]
**CreatedTime** | Pointer to **time.Time** | Date and time when the resource was created, in UTC. | [optional] 
**LastModifiedTime** | Pointer to **time.Time** | Date and time when the resource was last modified, in UTC. | [optional] 
**Mids** | Pointer to **[]string** | Comma-separated list of alphanumeric merchant identifiers. | [optional] 
**Name** | Pointer to **string** | Name of the merchant group. | [optional] 
**Token** | Pointer to **string** | Unique identifier of the merchant group. | [optional] 

## Methods

### NewMerchantGroupResponse

`func NewMerchantGroupResponse() *MerchantGroupResponse`

NewMerchantGroupResponse instantiates a new MerchantGroupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantGroupResponseWithDefaults

`func NewMerchantGroupResponseWithDefaults() *MerchantGroupResponse`

NewMerchantGroupResponseWithDefaults instantiates a new MerchantGroupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *MerchantGroupResponse) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *MerchantGroupResponse) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *MerchantGroupResponse) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *MerchantGroupResponse) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCreatedTime

`func (o *MerchantGroupResponse) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *MerchantGroupResponse) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *MerchantGroupResponse) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *MerchantGroupResponse) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetLastModifiedTime

`func (o *MerchantGroupResponse) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *MerchantGroupResponse) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *MerchantGroupResponse) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.

### HasLastModifiedTime

`func (o *MerchantGroupResponse) HasLastModifiedTime() bool`

HasLastModifiedTime returns a boolean if a field has been set.

### GetMids

`func (o *MerchantGroupResponse) GetMids() []string`

GetMids returns the Mids field if non-nil, zero value otherwise.

### GetMidsOk

`func (o *MerchantGroupResponse) GetMidsOk() (*[]string, bool)`

GetMidsOk returns a tuple with the Mids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMids

`func (o *MerchantGroupResponse) SetMids(v []string)`

SetMids sets Mids field to given value.

### HasMids

`func (o *MerchantGroupResponse) HasMids() bool`

HasMids returns a boolean if a field has been set.

### GetName

`func (o *MerchantGroupResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MerchantGroupResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MerchantGroupResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MerchantGroupResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetToken

`func (o *MerchantGroupResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *MerchantGroupResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *MerchantGroupResponse) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *MerchantGroupResponse) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


