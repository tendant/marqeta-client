# Account

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailAddress** | Pointer to **string** | Digital wallet provider&#39;s email address for the cardholder. | [optional] 
**Id** | Pointer to **string** | Digital wallet provider&#39;s identity number for the cardholder. | [optional] 
**Score** | Pointer to **string** | Score from the digital wallet provider. | [optional] 

## Methods

### NewAccount

`func NewAccount() *Account`

NewAccount instantiates a new Account object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountWithDefaults

`func NewAccountWithDefaults() *Account`

NewAccountWithDefaults instantiates a new Account object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailAddress

`func (o *Account) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *Account) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *Account) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *Account) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### GetId

`func (o *Account) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Account) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Account) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Account) HasId() bool`

HasId returns a boolean if a field has been set.

### GetScore

`func (o *Account) GetScore() string`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *Account) GetScoreOk() (*string, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *Account) SetScore(v string)`

SetScore sets Score field to given value.

### HasScore

`func (o *Account) HasScore() bool`

HasScore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


