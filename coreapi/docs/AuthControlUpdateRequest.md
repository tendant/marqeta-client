# AuthControlUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Indicates whether the authorization control is active. | [optional] [default to true]
**Association** | Pointer to [**SpendControlAssociation**](SpendControlAssociation.md) |  | [optional] 
**EndTime** | Pointer to **time.Time** | Date and time when the exception ends, in UTC. | [optional] 
**MerchantScope** | Pointer to [**MerchantScope**](MerchantScope.md) |  | [optional] 
**Name** | Pointer to **string** | Name of the authorization control. | [optional] 
**StartTime** | Pointer to **time.Time** | Date and time when the exception goes into effect, in UTC. | [optional] 
**Token** | **string** | Unique identifier of the authorization control. | 

## Methods

### NewAuthControlUpdateRequest

`func NewAuthControlUpdateRequest(token string, ) *AuthControlUpdateRequest`

NewAuthControlUpdateRequest instantiates a new AuthControlUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthControlUpdateRequestWithDefaults

`func NewAuthControlUpdateRequestWithDefaults() *AuthControlUpdateRequest`

NewAuthControlUpdateRequestWithDefaults instantiates a new AuthControlUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *AuthControlUpdateRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *AuthControlUpdateRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *AuthControlUpdateRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *AuthControlUpdateRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAssociation

`func (o *AuthControlUpdateRequest) GetAssociation() SpendControlAssociation`

GetAssociation returns the Association field if non-nil, zero value otherwise.

### GetAssociationOk

`func (o *AuthControlUpdateRequest) GetAssociationOk() (*SpendControlAssociation, bool)`

GetAssociationOk returns a tuple with the Association field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociation

`func (o *AuthControlUpdateRequest) SetAssociation(v SpendControlAssociation)`

SetAssociation sets Association field to given value.

### HasAssociation

`func (o *AuthControlUpdateRequest) HasAssociation() bool`

HasAssociation returns a boolean if a field has been set.

### GetEndTime

`func (o *AuthControlUpdateRequest) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *AuthControlUpdateRequest) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *AuthControlUpdateRequest) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *AuthControlUpdateRequest) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetMerchantScope

`func (o *AuthControlUpdateRequest) GetMerchantScope() MerchantScope`

GetMerchantScope returns the MerchantScope field if non-nil, zero value otherwise.

### GetMerchantScopeOk

`func (o *AuthControlUpdateRequest) GetMerchantScopeOk() (*MerchantScope, bool)`

GetMerchantScopeOk returns a tuple with the MerchantScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantScope

`func (o *AuthControlUpdateRequest) SetMerchantScope(v MerchantScope)`

SetMerchantScope sets MerchantScope field to given value.

### HasMerchantScope

`func (o *AuthControlUpdateRequest) HasMerchantScope() bool`

HasMerchantScope returns a boolean if a field has been set.

### GetName

`func (o *AuthControlUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthControlUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthControlUpdateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AuthControlUpdateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStartTime

`func (o *AuthControlUpdateRequest) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *AuthControlUpdateRequest) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *AuthControlUpdateRequest) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *AuthControlUpdateRequest) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetToken

`func (o *AuthControlUpdateRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AuthControlUpdateRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AuthControlUpdateRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


