# CardProductResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Indicates whether the card product is active.  This field is returned if it exists in the resource. | [optional] [default to false]
**Config** | Pointer to [**CardProductConfig**](CardProductConfig.md) |  | [optional] 
**CreatedTime** | **time.Time** | Date and time when the resource was created, in UTC. | 
**EndDate** | Pointer to **string** | End date of the range over which the card product can be active.  This field is returned if it exists in the resource. | [optional] 
**LastModifiedTime** | **time.Time** | Date and time when the resource was last updated, in UTC. | 
**Name** | **string** | Name of the card product. | 
**StartDate** | **string** | Date when the card product becomes active. | 
**Token** | Pointer to **string** | Unique identifier of the card product. | [optional] 

## Methods

### NewCardProductResponse

`func NewCardProductResponse(createdTime time.Time, lastModifiedTime time.Time, name string, startDate string, ) *CardProductResponse`

NewCardProductResponse instantiates a new CardProductResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardProductResponseWithDefaults

`func NewCardProductResponseWithDefaults() *CardProductResponse`

NewCardProductResponseWithDefaults instantiates a new CardProductResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *CardProductResponse) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CardProductResponse) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CardProductResponse) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *CardProductResponse) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetConfig

`func (o *CardProductResponse) GetConfig() CardProductConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *CardProductResponse) GetConfigOk() (*CardProductConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *CardProductResponse) SetConfig(v CardProductConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *CardProductResponse) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCreatedTime

`func (o *CardProductResponse) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *CardProductResponse) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *CardProductResponse) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetEndDate

`func (o *CardProductResponse) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *CardProductResponse) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *CardProductResponse) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *CardProductResponse) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetLastModifiedTime

`func (o *CardProductResponse) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *CardProductResponse) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *CardProductResponse) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.


### GetName

`func (o *CardProductResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CardProductResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CardProductResponse) SetName(v string)`

SetName sets Name field to given value.


### GetStartDate

`func (o *CardProductResponse) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CardProductResponse) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *CardProductResponse) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.


### GetToken

`func (o *CardProductResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CardProductResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CardProductResponse) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CardProductResponse) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


