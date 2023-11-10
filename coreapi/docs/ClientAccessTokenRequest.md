# ClientAccessTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationToken** | Pointer to **string** | Unique identifier of the &#x60;application&#x60; object. | [optional] 
**CardToken** | **string** | Unique identifier of the card whose sensitive information you want to display. | 

## Methods

### NewClientAccessTokenRequest

`func NewClientAccessTokenRequest(cardToken string, ) *ClientAccessTokenRequest`

NewClientAccessTokenRequest instantiates a new ClientAccessTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientAccessTokenRequestWithDefaults

`func NewClientAccessTokenRequestWithDefaults() *ClientAccessTokenRequest`

NewClientAccessTokenRequestWithDefaults instantiates a new ClientAccessTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationToken

`func (o *ClientAccessTokenRequest) GetApplicationToken() string`

GetApplicationToken returns the ApplicationToken field if non-nil, zero value otherwise.

### GetApplicationTokenOk

`func (o *ClientAccessTokenRequest) GetApplicationTokenOk() (*string, bool)`

GetApplicationTokenOk returns a tuple with the ApplicationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationToken

`func (o *ClientAccessTokenRequest) SetApplicationToken(v string)`

SetApplicationToken sets ApplicationToken field to given value.

### HasApplicationToken

`func (o *ClientAccessTokenRequest) HasApplicationToken() bool`

HasApplicationToken returns a boolean if a field has been set.

### GetCardToken

`func (o *ClientAccessTokenRequest) GetCardToken() string`

GetCardToken returns the CardToken field if non-nil, zero value otherwise.

### GetCardTokenOk

`func (o *ClientAccessTokenRequest) GetCardTokenOk() (*string, bool)`

GetCardTokenOk returns a tuple with the CardToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardToken

`func (o *ClientAccessTokenRequest) SetCardToken(v string)`

SetCardToken sets CardToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


