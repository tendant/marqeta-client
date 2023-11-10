# AuthControlRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Indicates whether the authorization control is active. | [optional] [default to true]
**Association** | Pointer to [**SpendControlAssociation**](SpendControlAssociation.md) |  | [optional] 
**EndTime** | Pointer to **time.Time** | Date and time when the exception ends, in UTC. | [optional] 
**MerchantScope** | Pointer to [**AuthControlMerchantScope**](AuthControlMerchantScope.md) |  | [optional] 
**Name** | **string** | Name of the authorization control. | 
**StartTime** | Pointer to **time.Time** | Date and time when the exception goes into effect, in UTC. | [optional] 
**Token** | Pointer to **string** | Unique identifier of the authorization control.  If you do not include a token, the system will generate one automatically. This token is necessary for use in other API calls, so we recommend that rather than let the system generate one, you use a simple string that is easy to remember. This value cannot be updated. | [optional] 

## Methods

### NewAuthControlRequest

`func NewAuthControlRequest(name string, ) *AuthControlRequest`

NewAuthControlRequest instantiates a new AuthControlRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthControlRequestWithDefaults

`func NewAuthControlRequestWithDefaults() *AuthControlRequest`

NewAuthControlRequestWithDefaults instantiates a new AuthControlRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *AuthControlRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *AuthControlRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *AuthControlRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *AuthControlRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAssociation

`func (o *AuthControlRequest) GetAssociation() SpendControlAssociation`

GetAssociation returns the Association field if non-nil, zero value otherwise.

### GetAssociationOk

`func (o *AuthControlRequest) GetAssociationOk() (*SpendControlAssociation, bool)`

GetAssociationOk returns a tuple with the Association field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociation

`func (o *AuthControlRequest) SetAssociation(v SpendControlAssociation)`

SetAssociation sets Association field to given value.

### HasAssociation

`func (o *AuthControlRequest) HasAssociation() bool`

HasAssociation returns a boolean if a field has been set.

### GetEndTime

`func (o *AuthControlRequest) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *AuthControlRequest) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *AuthControlRequest) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *AuthControlRequest) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetMerchantScope

`func (o *AuthControlRequest) GetMerchantScope() AuthControlMerchantScope`

GetMerchantScope returns the MerchantScope field if non-nil, zero value otherwise.

### GetMerchantScopeOk

`func (o *AuthControlRequest) GetMerchantScopeOk() (*AuthControlMerchantScope, bool)`

GetMerchantScopeOk returns a tuple with the MerchantScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantScope

`func (o *AuthControlRequest) SetMerchantScope(v AuthControlMerchantScope)`

SetMerchantScope sets MerchantScope field to given value.

### HasMerchantScope

`func (o *AuthControlRequest) HasMerchantScope() bool`

HasMerchantScope returns a boolean if a field has been set.

### GetName

`func (o *AuthControlRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthControlRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthControlRequest) SetName(v string)`

SetName sets Name field to given value.


### GetStartTime

`func (o *AuthControlRequest) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *AuthControlRequest) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *AuthControlRequest) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *AuthControlRequest) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetToken

`func (o *AuthControlRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AuthControlRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AuthControlRequest) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *AuthControlRequest) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


