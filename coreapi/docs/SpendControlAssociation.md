# SpendControlAssociation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardProductToken** | Pointer to **string** | Unique identifier of the card product.  Pass either &#x60;card_product_token&#x60; or &#x60;user_token&#x60;, not both. | [optional] 
**UserToken** | Pointer to **string** | Unique identifier of the cardholder.  Pass either &#x60;card_product_token&#x60; or &#x60;user_token&#x60;, not both. | [optional] 

## Methods

### NewSpendControlAssociation

`func NewSpendControlAssociation() *SpendControlAssociation`

NewSpendControlAssociation instantiates a new SpendControlAssociation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpendControlAssociationWithDefaults

`func NewSpendControlAssociationWithDefaults() *SpendControlAssociation`

NewSpendControlAssociationWithDefaults instantiates a new SpendControlAssociation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardProductToken

`func (o *SpendControlAssociation) GetCardProductToken() string`

GetCardProductToken returns the CardProductToken field if non-nil, zero value otherwise.

### GetCardProductTokenOk

`func (o *SpendControlAssociation) GetCardProductTokenOk() (*string, bool)`

GetCardProductTokenOk returns a tuple with the CardProductToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardProductToken

`func (o *SpendControlAssociation) SetCardProductToken(v string)`

SetCardProductToken sets CardProductToken field to given value.

### HasCardProductToken

`func (o *SpendControlAssociation) HasCardProductToken() bool`

HasCardProductToken returns a boolean if a field has been set.

### GetUserToken

`func (o *SpendControlAssociation) GetUserToken() string`

GetUserToken returns the UserToken field if non-nil, zero value otherwise.

### GetUserTokenOk

`func (o *SpendControlAssociation) GetUserTokenOk() (*string, bool)`

GetUserTokenOk returns a tuple with the UserToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserToken

`func (o *SpendControlAssociation) SetUserToken(v string)`

SetUserToken sets UserToken field to given value.

### HasUserToken

`func (o *SpendControlAssociation) HasUserToken() bool`

HasUserToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


