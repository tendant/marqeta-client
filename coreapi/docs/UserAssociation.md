# UserAssociation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SingleInventoryUser** | Pointer to **bool** | Set to &#x60;true&#x60; to associate all cards with the same user. Set to &#x60;false&#x60; to associate each card with a different user. When set to &#x60;false&#x60;, users are generated automatically and associated with the cards. | [optional] [default to false]
**SingleInventoryUserToken** | Pointer to **string** | If &#x60;single_inventory_user&#x3D;true&#x60;, use this field to specify the token of an existing user. All cards in the order will be associated with this user. | [optional] [default to "false"]

## Methods

### NewUserAssociation

`func NewUserAssociation() *UserAssociation`

NewUserAssociation instantiates a new UserAssociation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserAssociationWithDefaults

`func NewUserAssociationWithDefaults() *UserAssociation`

NewUserAssociationWithDefaults instantiates a new UserAssociation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSingleInventoryUser

`func (o *UserAssociation) GetSingleInventoryUser() bool`

GetSingleInventoryUser returns the SingleInventoryUser field if non-nil, zero value otherwise.

### GetSingleInventoryUserOk

`func (o *UserAssociation) GetSingleInventoryUserOk() (*bool, bool)`

GetSingleInventoryUserOk returns a tuple with the SingleInventoryUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleInventoryUser

`func (o *UserAssociation) SetSingleInventoryUser(v bool)`

SetSingleInventoryUser sets SingleInventoryUser field to given value.

### HasSingleInventoryUser

`func (o *UserAssociation) HasSingleInventoryUser() bool`

HasSingleInventoryUser returns a boolean if a field has been set.

### GetSingleInventoryUserToken

`func (o *UserAssociation) GetSingleInventoryUserToken() string`

GetSingleInventoryUserToken returns the SingleInventoryUserToken field if non-nil, zero value otherwise.

### GetSingleInventoryUserTokenOk

`func (o *UserAssociation) GetSingleInventoryUserTokenOk() (*string, bool)`

GetSingleInventoryUserTokenOk returns a tuple with the SingleInventoryUserToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleInventoryUserToken

`func (o *UserAssociation) SetSingleInventoryUserToken(v string)`

SetSingleInventoryUserToken sets SingleInventoryUserToken field to given value.

### HasSingleInventoryUserToken

`func (o *UserAssociation) HasSingleInventoryUserToken() bool`

HasSingleInventoryUserToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


