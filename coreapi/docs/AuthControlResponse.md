# AuthControlResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Indicates whether the authorization control is active.  This field is returned if it exists in the resource. | [optional] [default to true]
**Association** | Pointer to [**SpendControlAssociation**](SpendControlAssociation.md) |  | [optional] 
**EndTime** | Pointer to **time.Time** | Date and time when the exception ends, in UTC.  This field is returned if it exists in the resource. | [optional] 
**MerchantScope** | Pointer to [**AuthControlMerchantScope**](AuthControlMerchantScope.md) |  | [optional] 
**Name** | **string** | Name of the authorization control. | 
**StartTime** | Pointer to **time.Time** | Date and time when the exception goes into effect, in UTC.  This field is returned if it exists in the resource. | [optional] 
**Token** | Pointer to **string** | Unique identifier of the authorization control.  This field is returned if it exists in the resource. | [optional] 

## Methods

### NewAuthControlResponse

`func NewAuthControlResponse(name string, ) *AuthControlResponse`

NewAuthControlResponse instantiates a new AuthControlResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthControlResponseWithDefaults

`func NewAuthControlResponseWithDefaults() *AuthControlResponse`

NewAuthControlResponseWithDefaults instantiates a new AuthControlResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *AuthControlResponse) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *AuthControlResponse) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *AuthControlResponse) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *AuthControlResponse) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAssociation

`func (o *AuthControlResponse) GetAssociation() SpendControlAssociation`

GetAssociation returns the Association field if non-nil, zero value otherwise.

### GetAssociationOk

`func (o *AuthControlResponse) GetAssociationOk() (*SpendControlAssociation, bool)`

GetAssociationOk returns a tuple with the Association field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociation

`func (o *AuthControlResponse) SetAssociation(v SpendControlAssociation)`

SetAssociation sets Association field to given value.

### HasAssociation

`func (o *AuthControlResponse) HasAssociation() bool`

HasAssociation returns a boolean if a field has been set.

### GetEndTime

`func (o *AuthControlResponse) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *AuthControlResponse) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *AuthControlResponse) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *AuthControlResponse) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetMerchantScope

`func (o *AuthControlResponse) GetMerchantScope() AuthControlMerchantScope`

GetMerchantScope returns the MerchantScope field if non-nil, zero value otherwise.

### GetMerchantScopeOk

`func (o *AuthControlResponse) GetMerchantScopeOk() (*AuthControlMerchantScope, bool)`

GetMerchantScopeOk returns a tuple with the MerchantScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantScope

`func (o *AuthControlResponse) SetMerchantScope(v AuthControlMerchantScope)`

SetMerchantScope sets MerchantScope field to given value.

### HasMerchantScope

`func (o *AuthControlResponse) HasMerchantScope() bool`

HasMerchantScope returns a boolean if a field has been set.

### GetName

`func (o *AuthControlResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthControlResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthControlResponse) SetName(v string)`

SetName sets Name field to given value.


### GetStartTime

`func (o *AuthControlResponse) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *AuthControlResponse) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *AuthControlResponse) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *AuthControlResponse) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetToken

`func (o *AuthControlResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AuthControlResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AuthControlResponse) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *AuthControlResponse) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


