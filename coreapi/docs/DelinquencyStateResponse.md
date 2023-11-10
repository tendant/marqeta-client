# DelinquencyStateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountToken** | **string** | Unique identifier of the credit account. | 
**Buckets** | Pointer to [**[]DelinquencyBucketResponse**](DelinquencyBucketResponse.md) | One or more delinquency buckets for an account. Each delinquency bucket represents a billing cycle during which the account was delinquent. | [optional] 
**CurrentDue** | **float32** | Amount that is due for the current billing cycle. | 
**DateAccountCurrent** | Pointer to **NullableTime** | Date and time when the account was last made current on the Marqeta platform, in UTC.  If the account was never delinquent, this field returns the date and time the account was created on the Marqeta platform, in UTC.  If &#x60;is_delinquent&#x60; is &#x60;true&#x60;, a null value is returned. | [optional] 
**DateAccountDelinquent** | Pointer to **NullableTime** | Date and time when the account last fell delinquent on the Marqeta platform, in UTC.  If &#x60;is_delinquent&#x60; is &#x60;false&#x60;, a null value is returned. | [optional] 
**IsDelinquent** | **bool** | A value of &#x60;true&#x60; indicates that the account is currently delinquent. | 
**TotalDaysPastDue** | **int32** | Total number of days that the account is past due. | 
**TotalDue** | **float32** | Total amount that is due for the current billing cycle; the sum of &#x60;total_past_due_amount&#x60; and &#x60;current_due_amount&#x60;. | 
**TotalPastDue** | **float32** | Total amount that is past due. | 

## Methods

### NewDelinquencyStateResponse

`func NewDelinquencyStateResponse(accountToken string, currentDue float32, isDelinquent bool, totalDaysPastDue int32, totalDue float32, totalPastDue float32, ) *DelinquencyStateResponse`

NewDelinquencyStateResponse instantiates a new DelinquencyStateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDelinquencyStateResponseWithDefaults

`func NewDelinquencyStateResponseWithDefaults() *DelinquencyStateResponse`

NewDelinquencyStateResponseWithDefaults instantiates a new DelinquencyStateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountToken

`func (o *DelinquencyStateResponse) GetAccountToken() string`

GetAccountToken returns the AccountToken field if non-nil, zero value otherwise.

### GetAccountTokenOk

`func (o *DelinquencyStateResponse) GetAccountTokenOk() (*string, bool)`

GetAccountTokenOk returns a tuple with the AccountToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountToken

`func (o *DelinquencyStateResponse) SetAccountToken(v string)`

SetAccountToken sets AccountToken field to given value.


### GetBuckets

`func (o *DelinquencyStateResponse) GetBuckets() []DelinquencyBucketResponse`

GetBuckets returns the Buckets field if non-nil, zero value otherwise.

### GetBucketsOk

`func (o *DelinquencyStateResponse) GetBucketsOk() (*[]DelinquencyBucketResponse, bool)`

GetBucketsOk returns a tuple with the Buckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuckets

`func (o *DelinquencyStateResponse) SetBuckets(v []DelinquencyBucketResponse)`

SetBuckets sets Buckets field to given value.

### HasBuckets

`func (o *DelinquencyStateResponse) HasBuckets() bool`

HasBuckets returns a boolean if a field has been set.

### GetCurrentDue

`func (o *DelinquencyStateResponse) GetCurrentDue() float32`

GetCurrentDue returns the CurrentDue field if non-nil, zero value otherwise.

### GetCurrentDueOk

`func (o *DelinquencyStateResponse) GetCurrentDueOk() (*float32, bool)`

GetCurrentDueOk returns a tuple with the CurrentDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentDue

`func (o *DelinquencyStateResponse) SetCurrentDue(v float32)`

SetCurrentDue sets CurrentDue field to given value.


### GetDateAccountCurrent

`func (o *DelinquencyStateResponse) GetDateAccountCurrent() time.Time`

GetDateAccountCurrent returns the DateAccountCurrent field if non-nil, zero value otherwise.

### GetDateAccountCurrentOk

`func (o *DelinquencyStateResponse) GetDateAccountCurrentOk() (*time.Time, bool)`

GetDateAccountCurrentOk returns a tuple with the DateAccountCurrent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAccountCurrent

`func (o *DelinquencyStateResponse) SetDateAccountCurrent(v time.Time)`

SetDateAccountCurrent sets DateAccountCurrent field to given value.

### HasDateAccountCurrent

`func (o *DelinquencyStateResponse) HasDateAccountCurrent() bool`

HasDateAccountCurrent returns a boolean if a field has been set.

### SetDateAccountCurrentNil

`func (o *DelinquencyStateResponse) SetDateAccountCurrentNil(b bool)`

 SetDateAccountCurrentNil sets the value for DateAccountCurrent to be an explicit nil

### UnsetDateAccountCurrent
`func (o *DelinquencyStateResponse) UnsetDateAccountCurrent()`

UnsetDateAccountCurrent ensures that no value is present for DateAccountCurrent, not even an explicit nil
### GetDateAccountDelinquent

`func (o *DelinquencyStateResponse) GetDateAccountDelinquent() time.Time`

GetDateAccountDelinquent returns the DateAccountDelinquent field if non-nil, zero value otherwise.

### GetDateAccountDelinquentOk

`func (o *DelinquencyStateResponse) GetDateAccountDelinquentOk() (*time.Time, bool)`

GetDateAccountDelinquentOk returns a tuple with the DateAccountDelinquent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAccountDelinquent

`func (o *DelinquencyStateResponse) SetDateAccountDelinquent(v time.Time)`

SetDateAccountDelinquent sets DateAccountDelinquent field to given value.

### HasDateAccountDelinquent

`func (o *DelinquencyStateResponse) HasDateAccountDelinquent() bool`

HasDateAccountDelinquent returns a boolean if a field has been set.

### SetDateAccountDelinquentNil

`func (o *DelinquencyStateResponse) SetDateAccountDelinquentNil(b bool)`

 SetDateAccountDelinquentNil sets the value for DateAccountDelinquent to be an explicit nil

### UnsetDateAccountDelinquent
`func (o *DelinquencyStateResponse) UnsetDateAccountDelinquent()`

UnsetDateAccountDelinquent ensures that no value is present for DateAccountDelinquent, not even an explicit nil
### GetIsDelinquent

`func (o *DelinquencyStateResponse) GetIsDelinquent() bool`

GetIsDelinquent returns the IsDelinquent field if non-nil, zero value otherwise.

### GetIsDelinquentOk

`func (o *DelinquencyStateResponse) GetIsDelinquentOk() (*bool, bool)`

GetIsDelinquentOk returns a tuple with the IsDelinquent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDelinquent

`func (o *DelinquencyStateResponse) SetIsDelinquent(v bool)`

SetIsDelinquent sets IsDelinquent field to given value.


### GetTotalDaysPastDue

`func (o *DelinquencyStateResponse) GetTotalDaysPastDue() int32`

GetTotalDaysPastDue returns the TotalDaysPastDue field if non-nil, zero value otherwise.

### GetTotalDaysPastDueOk

`func (o *DelinquencyStateResponse) GetTotalDaysPastDueOk() (*int32, bool)`

GetTotalDaysPastDueOk returns a tuple with the TotalDaysPastDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDaysPastDue

`func (o *DelinquencyStateResponse) SetTotalDaysPastDue(v int32)`

SetTotalDaysPastDue sets TotalDaysPastDue field to given value.


### GetTotalDue

`func (o *DelinquencyStateResponse) GetTotalDue() float32`

GetTotalDue returns the TotalDue field if non-nil, zero value otherwise.

### GetTotalDueOk

`func (o *DelinquencyStateResponse) GetTotalDueOk() (*float32, bool)`

GetTotalDueOk returns a tuple with the TotalDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDue

`func (o *DelinquencyStateResponse) SetTotalDue(v float32)`

SetTotalDue sets TotalDue field to given value.


### GetTotalPastDue

`func (o *DelinquencyStateResponse) GetTotalPastDue() float32`

GetTotalPastDue returns the TotalPastDue field if non-nil, zero value otherwise.

### GetTotalPastDueOk

`func (o *DelinquencyStateResponse) GetTotalPastDueOk() (*float32, bool)`

GetTotalPastDueOk returns a tuple with the TotalPastDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPastDue

`func (o *DelinquencyStateResponse) SetTotalPastDue(v float32)`

SetTotalPastDue sets TotalPastDue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


