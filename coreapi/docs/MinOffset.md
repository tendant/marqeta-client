# MinOffset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Unit** | Pointer to **string** | Specifies the time unit of the &#x60;value&#x60; field. | [optional] 
**Value** | Pointer to **int32** | Specifies the number of time units (as defined by the &#x60;unit&#x60; field) for which cards of this card product type are valid. Cards expire &#x60;value&#x60; x &#x60;unit&#x60; after the date of issue.  This number is rounded as follows:  * *YEARS* – Rounds up to the last second of the last day of the month of expiration. For example, if the issue date is 1 Jan 2021 and &#x60;value &#x3D; 1&#x60;, the cards expire on the last day of Jan 2022.  * *MONTHS* – Rounds up to the last second of the last day of the month of expiration. For example, if the issue date is 1 May 2022 and &#x60;value &#x3D; 1&#x60;, the cards expire on the last day of June 2022.  * *DAYS* – Rounds up to the last second of the day of expiration.  * *HOURS*, *MINUTES*, *SECONDS* – No rounding. | [optional] 

## Methods

### NewMinOffset

`func NewMinOffset() *MinOffset`

NewMinOffset instantiates a new MinOffset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMinOffsetWithDefaults

`func NewMinOffsetWithDefaults() *MinOffset`

NewMinOffsetWithDefaults instantiates a new MinOffset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnit

`func (o *MinOffset) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *MinOffset) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *MinOffset) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *MinOffset) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetValue

`func (o *MinOffset) GetValue() int32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *MinOffset) GetValueOk() (*int32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *MinOffset) SetValue(v int32)`

SetValue sets Value field to given value.

### HasValue

`func (o *MinOffset) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


