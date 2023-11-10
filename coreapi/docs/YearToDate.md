# YearToDate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountToken** | **string** | Unique identifier of the associated credit account. | 
**CreatedTime** | Pointer to **time.Time** | Date and time when the year-to-date total was created on Marqeta&#39;s credit platform, in UTC. | [optional] 
**StatementToken** | **string** | Unique identifier of the statement summary from which to retrieve year-to-date totals. | 
**Token** | Pointer to **string** | Unique identifier of the year-to-date total. | [optional] 
**TotalFees** | **float32** | Total fees charged year-to-date. | 
**TotalInterest** | **float32** | Total interest charged year-to-date. | 

## Methods

### NewYearToDate

`func NewYearToDate(accountToken string, statementToken string, totalFees float32, totalInterest float32, ) *YearToDate`

NewYearToDate instantiates a new YearToDate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewYearToDateWithDefaults

`func NewYearToDateWithDefaults() *YearToDate`

NewYearToDateWithDefaults instantiates a new YearToDate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountToken

`func (o *YearToDate) GetAccountToken() string`

GetAccountToken returns the AccountToken field if non-nil, zero value otherwise.

### GetAccountTokenOk

`func (o *YearToDate) GetAccountTokenOk() (*string, bool)`

GetAccountTokenOk returns a tuple with the AccountToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountToken

`func (o *YearToDate) SetAccountToken(v string)`

SetAccountToken sets AccountToken field to given value.


### GetCreatedTime

`func (o *YearToDate) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *YearToDate) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *YearToDate) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *YearToDate) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetStatementToken

`func (o *YearToDate) GetStatementToken() string`

GetStatementToken returns the StatementToken field if non-nil, zero value otherwise.

### GetStatementTokenOk

`func (o *YearToDate) GetStatementTokenOk() (*string, bool)`

GetStatementTokenOk returns a tuple with the StatementToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatementToken

`func (o *YearToDate) SetStatementToken(v string)`

SetStatementToken sets StatementToken field to given value.


### GetToken

`func (o *YearToDate) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *YearToDate) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *YearToDate) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *YearToDate) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTotalFees

`func (o *YearToDate) GetTotalFees() float32`

GetTotalFees returns the TotalFees field if non-nil, zero value otherwise.

### GetTotalFeesOk

`func (o *YearToDate) GetTotalFeesOk() (*float32, bool)`

GetTotalFeesOk returns a tuple with the TotalFees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalFees

`func (o *YearToDate) SetTotalFees(v float32)`

SetTotalFees sets TotalFees field to given value.


### GetTotalInterest

`func (o *YearToDate) GetTotalInterest() float32`

GetTotalInterest returns the TotalInterest field if non-nil, zero value otherwise.

### GetTotalInterestOk

`func (o *YearToDate) GetTotalInterestOk() (*float32, bool)`

GetTotalInterestOk returns a tuple with the TotalInterest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalInterest

`func (o *YearToDate) SetTotalInterest(v float32)`

SetTotalInterest sets TotalInterest field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


