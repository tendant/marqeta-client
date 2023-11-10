# AutoReloadAssociation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessToken** | Pointer to **string** | Unique identifier of the business for which the auto reload is configured.  Send a &#x60;GET&#x60; request to &#x60;/businesses&#x60; to retrieve business tokens. | [optional] 
**CardProductToken** | Pointer to **string** | Unique identifier of the card product for which the auto reload is configured.  Send a &#x60;GET&#x60; request to &#x60;/cardproducts&#x60; to retrieve card product tokens. | [optional] 
**UserToken** | Pointer to **string** | Unique identifier of the user for which the auto reload is configured.  Send a &#x60;GET&#x60; request to &#x60;/users&#x60; to retrieve user tokens. | [optional] 

## Methods

### NewAutoReloadAssociation

`func NewAutoReloadAssociation() *AutoReloadAssociation`

NewAutoReloadAssociation instantiates a new AutoReloadAssociation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoReloadAssociationWithDefaults

`func NewAutoReloadAssociationWithDefaults() *AutoReloadAssociation`

NewAutoReloadAssociationWithDefaults instantiates a new AutoReloadAssociation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessToken

`func (o *AutoReloadAssociation) GetBusinessToken() string`

GetBusinessToken returns the BusinessToken field if non-nil, zero value otherwise.

### GetBusinessTokenOk

`func (o *AutoReloadAssociation) GetBusinessTokenOk() (*string, bool)`

GetBusinessTokenOk returns a tuple with the BusinessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessToken

`func (o *AutoReloadAssociation) SetBusinessToken(v string)`

SetBusinessToken sets BusinessToken field to given value.

### HasBusinessToken

`func (o *AutoReloadAssociation) HasBusinessToken() bool`

HasBusinessToken returns a boolean if a field has been set.

### GetCardProductToken

`func (o *AutoReloadAssociation) GetCardProductToken() string`

GetCardProductToken returns the CardProductToken field if non-nil, zero value otherwise.

### GetCardProductTokenOk

`func (o *AutoReloadAssociation) GetCardProductTokenOk() (*string, bool)`

GetCardProductTokenOk returns a tuple with the CardProductToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardProductToken

`func (o *AutoReloadAssociation) SetCardProductToken(v string)`

SetCardProductToken sets CardProductToken field to given value.

### HasCardProductToken

`func (o *AutoReloadAssociation) HasCardProductToken() bool`

HasCardProductToken returns a boolean if a field has been set.

### GetUserToken

`func (o *AutoReloadAssociation) GetUserToken() string`

GetUserToken returns the UserToken field if non-nil, zero value otherwise.

### GetUserTokenOk

`func (o *AutoReloadAssociation) GetUserTokenOk() (*string, bool)`

GetUserTokenOk returns a tuple with the UserToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserToken

`func (o *AutoReloadAssociation) SetUserToken(v string)`

SetUserToken sets UserToken field to given value.

### HasUserToken

`func (o *AutoReloadAssociation) HasUserToken() bool`

HasUserToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


