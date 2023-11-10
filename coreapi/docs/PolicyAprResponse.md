# PolicyAprResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedTime** | Pointer to **time.Time** | Date and time when the APR policy was created on Marqeta&#39;s credit platform, in UTC. | [optional] 
**Description** | Pointer to **string** | Description of the APR policy. | [optional] 
**EffectiveDate** | Pointer to **string** | Date the APR goes into effect, in UTC. | [optional] 
**Name** | Pointer to **string** | Name of the APR policy. | [optional] 
**Purchases** | Pointer to [**PolicyAprPurchaseResponse**](PolicyAprPurchaseResponse.md) |  | [optional] 
**Token** | Pointer to **string** | Unique identifier of the APR policy. | [optional] 
**UpdatedTime** | Pointer to **time.Time** | Date and time when the APR policy was last updated on Marqeta&#39;s credit platform, in UTC. | [optional] 

## Methods

### NewPolicyAprResponse

`func NewPolicyAprResponse() *PolicyAprResponse`

NewPolicyAprResponse instantiates a new PolicyAprResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyAprResponseWithDefaults

`func NewPolicyAprResponseWithDefaults() *PolicyAprResponse`

NewPolicyAprResponseWithDefaults instantiates a new PolicyAprResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedTime

`func (o *PolicyAprResponse) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *PolicyAprResponse) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *PolicyAprResponse) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *PolicyAprResponse) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDescription

`func (o *PolicyAprResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PolicyAprResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PolicyAprResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PolicyAprResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEffectiveDate

`func (o *PolicyAprResponse) GetEffectiveDate() string`

GetEffectiveDate returns the EffectiveDate field if non-nil, zero value otherwise.

### GetEffectiveDateOk

`func (o *PolicyAprResponse) GetEffectiveDateOk() (*string, bool)`

GetEffectiveDateOk returns a tuple with the EffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveDate

`func (o *PolicyAprResponse) SetEffectiveDate(v string)`

SetEffectiveDate sets EffectiveDate field to given value.

### HasEffectiveDate

`func (o *PolicyAprResponse) HasEffectiveDate() bool`

HasEffectiveDate returns a boolean if a field has been set.

### GetName

`func (o *PolicyAprResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PolicyAprResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PolicyAprResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PolicyAprResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPurchases

`func (o *PolicyAprResponse) GetPurchases() PolicyAprPurchaseResponse`

GetPurchases returns the Purchases field if non-nil, zero value otherwise.

### GetPurchasesOk

`func (o *PolicyAprResponse) GetPurchasesOk() (*PolicyAprPurchaseResponse, bool)`

GetPurchasesOk returns a tuple with the Purchases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchases

`func (o *PolicyAprResponse) SetPurchases(v PolicyAprPurchaseResponse)`

SetPurchases sets Purchases field to given value.

### HasPurchases

`func (o *PolicyAprResponse) HasPurchases() bool`

HasPurchases returns a boolean if a field has been set.

### GetToken

`func (o *PolicyAprResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PolicyAprResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PolicyAprResponse) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *PolicyAprResponse) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUpdatedTime

`func (o *PolicyAprResponse) GetUpdatedTime() time.Time`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *PolicyAprResponse) GetUpdatedTimeOk() (*time.Time, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *PolicyAprResponse) SetUpdatedTime(v time.Time)`

SetUpdatedTime sets UpdatedTime field to given value.

### HasUpdatedTime

`func (o *PolicyAprResponse) HasUpdatedTime() bool`

HasUpdatedTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


