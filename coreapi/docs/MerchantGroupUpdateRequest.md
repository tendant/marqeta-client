# MerchantGroupUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Indicates if the merchant group is active or not. | [optional] [default to false]
**Mids** | Pointer to **[]string** | Comma-separated list of alphanumeric merchant identifiers. You can include merchant identifiers in multiple merchant groups. | [optional] 
**Name** | Pointer to **string** | Name of the merchant group. | [optional] 

## Methods

### NewMerchantGroupUpdateRequest

`func NewMerchantGroupUpdateRequest() *MerchantGroupUpdateRequest`

NewMerchantGroupUpdateRequest instantiates a new MerchantGroupUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantGroupUpdateRequestWithDefaults

`func NewMerchantGroupUpdateRequestWithDefaults() *MerchantGroupUpdateRequest`

NewMerchantGroupUpdateRequestWithDefaults instantiates a new MerchantGroupUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *MerchantGroupUpdateRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *MerchantGroupUpdateRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *MerchantGroupUpdateRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *MerchantGroupUpdateRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetMids

`func (o *MerchantGroupUpdateRequest) GetMids() []string`

GetMids returns the Mids field if non-nil, zero value otherwise.

### GetMidsOk

`func (o *MerchantGroupUpdateRequest) GetMidsOk() (*[]string, bool)`

GetMidsOk returns a tuple with the Mids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMids

`func (o *MerchantGroupUpdateRequest) SetMids(v []string)`

SetMids sets Mids field to given value.

### HasMids

`func (o *MerchantGroupUpdateRequest) HasMids() bool`

HasMids returns a boolean if a field has been set.

### GetName

`func (o *MerchantGroupUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MerchantGroupUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MerchantGroupUpdateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MerchantGroupUpdateRequest) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


