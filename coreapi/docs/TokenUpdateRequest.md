# TokenUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Indicates whether the card funding source is active. | [optional] [default to true]
**ExpDate** | **string** | Expiration date for the payment card. | 
**IsDefaultAccount** | Pointer to **bool** | If there are multiple funding sources, this field specifies which source is used by default in funding calls. If there is only one funding source, the system ignores this field and always uses that source. | [optional] [default to false]

## Methods

### NewTokenUpdateRequest

`func NewTokenUpdateRequest(expDate string, ) *TokenUpdateRequest`

NewTokenUpdateRequest instantiates a new TokenUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenUpdateRequestWithDefaults

`func NewTokenUpdateRequestWithDefaults() *TokenUpdateRequest`

NewTokenUpdateRequestWithDefaults instantiates a new TokenUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *TokenUpdateRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *TokenUpdateRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *TokenUpdateRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *TokenUpdateRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetExpDate

`func (o *TokenUpdateRequest) GetExpDate() string`

GetExpDate returns the ExpDate field if non-nil, zero value otherwise.

### GetExpDateOk

`func (o *TokenUpdateRequest) GetExpDateOk() (*string, bool)`

GetExpDateOk returns a tuple with the ExpDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpDate

`func (o *TokenUpdateRequest) SetExpDate(v string)`

SetExpDate sets ExpDate field to given value.


### GetIsDefaultAccount

`func (o *TokenUpdateRequest) GetIsDefaultAccount() bool`

GetIsDefaultAccount returns the IsDefaultAccount field if non-nil, zero value otherwise.

### GetIsDefaultAccountOk

`func (o *TokenUpdateRequest) GetIsDefaultAccountOk() (*bool, bool)`

GetIsDefaultAccountOk returns a tuple with the IsDefaultAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefaultAccount

`func (o *TokenUpdateRequest) SetIsDefaultAccount(v bool)`

SetIsDefaultAccount sets IsDefaultAccount field to given value.

### HasIsDefaultAccount

`func (o *TokenUpdateRequest) HasIsDefaultAccount() bool`

HasIsDefaultAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


