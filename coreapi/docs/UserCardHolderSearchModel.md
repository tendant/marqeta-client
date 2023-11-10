# UserCardHolderSearchModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dda** | Pointer to **string** | Performs a match on the specified deposit account number. Send a &#x60;GET&#x60; request to &#x60;/directdeposits/accounts/{user_token}&#x60; to retrieve the deposit account number for a specific cardholder. | [optional] 
**Email** | Pointer to **string** | Performs a non-case-sensitive, exact match on the cardholder&#39;s &#x60;email&#x60; field. | [optional] 
**FirstName** | Pointer to **string** | Performs a non-case-sensitive match on the cardholder&#39;s &#x60;first_name&#x60; field. Matching is partial on the beginning of the name. For example, a match on \&quot;Alex\&quot; returns both \&quot;Alex\&quot; and \&quot;Alexander\&quot;. | [optional] 
**LastName** | Pointer to **string** | Performs a non-case-sensitive match on the cardholder&#39;s &#x60;last_name&#x60; field. Matching is partial on the beginning of the name. For example, a match on \&quot;Smith\&quot; returns both \&quot;Smith\&quot; and \&quot;Smithfield\&quot;. | [optional] 
**Phone** | Pointer to **string** | Performs a match on the cardholder&#39;s &#x60;phone&#x60; field. | [optional] 
**Ssn** | Pointer to **string** | Performs a match on the cardholder&#39;s identification number. | [optional] 

## Methods

### NewUserCardHolderSearchModel

`func NewUserCardHolderSearchModel() *UserCardHolderSearchModel`

NewUserCardHolderSearchModel instantiates a new UserCardHolderSearchModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserCardHolderSearchModelWithDefaults

`func NewUserCardHolderSearchModelWithDefaults() *UserCardHolderSearchModel`

NewUserCardHolderSearchModelWithDefaults instantiates a new UserCardHolderSearchModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDda

`func (o *UserCardHolderSearchModel) GetDda() string`

GetDda returns the Dda field if non-nil, zero value otherwise.

### GetDdaOk

`func (o *UserCardHolderSearchModel) GetDdaOk() (*string, bool)`

GetDdaOk returns a tuple with the Dda field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDda

`func (o *UserCardHolderSearchModel) SetDda(v string)`

SetDda sets Dda field to given value.

### HasDda

`func (o *UserCardHolderSearchModel) HasDda() bool`

HasDda returns a boolean if a field has been set.

### GetEmail

`func (o *UserCardHolderSearchModel) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserCardHolderSearchModel) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserCardHolderSearchModel) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserCardHolderSearchModel) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFirstName

`func (o *UserCardHolderSearchModel) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UserCardHolderSearchModel) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UserCardHolderSearchModel) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *UserCardHolderSearchModel) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *UserCardHolderSearchModel) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UserCardHolderSearchModel) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UserCardHolderSearchModel) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *UserCardHolderSearchModel) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetPhone

`func (o *UserCardHolderSearchModel) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *UserCardHolderSearchModel) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *UserCardHolderSearchModel) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *UserCardHolderSearchModel) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetSsn

`func (o *UserCardHolderSearchModel) GetSsn() string`

GetSsn returns the Ssn field if non-nil, zero value otherwise.

### GetSsnOk

`func (o *UserCardHolderSearchModel) GetSsnOk() (*string, bool)`

GetSsnOk returns a tuple with the Ssn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsn

`func (o *UserCardHolderSearchModel) SetSsn(v string)`

SetSsn sets Ssn field to given value.

### HasSsn

`func (o *UserCardHolderSearchModel) HasSsn() bool`

HasSsn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


