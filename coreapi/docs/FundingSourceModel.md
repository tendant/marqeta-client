# FundingSourceModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | **bool** | Whether the funding source is active. | [default to false]
**CreatedTime** | **time.Time** | Date and time when the funding source was created, in UTC. | 
**IsDefaultAccount** | **bool** | Whether the GPA order unload&#39;s funding source is the default funding account. | [default to false]
**LastModifiedTime** | **time.Time** | Date and time when the funding source was last modified, in UTC. | 
**Token** | **string** | Unique identifier of the funding source. | 
**Type** | **string** | Funding type of the funding source. | 

## Methods

### NewFundingSourceModel

`func NewFundingSourceModel(active bool, createdTime time.Time, isDefaultAccount bool, lastModifiedTime time.Time, token string, type_ string, ) *FundingSourceModel`

NewFundingSourceModel instantiates a new FundingSourceModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFundingSourceModelWithDefaults

`func NewFundingSourceModelWithDefaults() *FundingSourceModel`

NewFundingSourceModelWithDefaults instantiates a new FundingSourceModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *FundingSourceModel) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *FundingSourceModel) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *FundingSourceModel) SetActive(v bool)`

SetActive sets Active field to given value.


### GetCreatedTime

`func (o *FundingSourceModel) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *FundingSourceModel) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *FundingSourceModel) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetIsDefaultAccount

`func (o *FundingSourceModel) GetIsDefaultAccount() bool`

GetIsDefaultAccount returns the IsDefaultAccount field if non-nil, zero value otherwise.

### GetIsDefaultAccountOk

`func (o *FundingSourceModel) GetIsDefaultAccountOk() (*bool, bool)`

GetIsDefaultAccountOk returns a tuple with the IsDefaultAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefaultAccount

`func (o *FundingSourceModel) SetIsDefaultAccount(v bool)`

SetIsDefaultAccount sets IsDefaultAccount field to given value.


### GetLastModifiedTime

`func (o *FundingSourceModel) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *FundingSourceModel) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *FundingSourceModel) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.


### GetToken

`func (o *FundingSourceModel) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *FundingSourceModel) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *FundingSourceModel) SetToken(v string)`

SetToken sets Token field to given value.


### GetType

`func (o *FundingSourceModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FundingSourceModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FundingSourceModel) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


