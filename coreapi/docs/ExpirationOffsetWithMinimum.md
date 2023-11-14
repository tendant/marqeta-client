# ExpirationOffsetWithMinimum

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinOffset** | Pointer to [**MinOffset**](MinOffset.md) |  | [optional] 
**Unit** | Pointer to **string** | Specifies the units for the &#x60;value&#x60; field. | [optional] 
**Value** | Pointer to **int32** | Specifies the number of time units (as defined by the &#x60;unit&#x60; field in this object) for which cards of this card product type are valid. In other words, cards expire &#x60;value&#x60; x &#x60;unit&#x60; after the date of issue.  This number is rounded as follows:  * *YEARS*  Rounds up to the last second of the last day of the month of expiration. For example, if the issue date is 1 Jan 2021 and &#x60;value &#x3D; 1&#x60;, the cards expire on the last day of Jan 2022.  * *MONTHS*  Rounds up to the last second of the last day of the month of expiration. For example, if the issue date is 1 May 2022 and &#x60;value &#x3D; 1&#x60;, the cards expire on the last day of June 2022.  * *DAYS*  Rounds up to the last second of the day of expiration.  * *HOURS*, *MINUTES*, *SECONDS*  No rounding. | [optional] 

## Methods

### NewExpirationOffsetWithMinimum

`func NewExpirationOffsetWithMinimum() *ExpirationOffsetWithMinimum`

NewExpirationOffsetWithMinimum instantiates a new ExpirationOffsetWithMinimum object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpirationOffsetWithMinimumWithDefaults

`func NewExpirationOffsetWithMinimumWithDefaults() *ExpirationOffsetWithMinimum`

NewExpirationOffsetWithMinimumWithDefaults instantiates a new ExpirationOffsetWithMinimum object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinOffset

`func (o *ExpirationOffsetWithMinimum) GetMinOffset() MinOffset`

GetMinOffset returns the MinOffset field if non-nil, zero value otherwise.

### GetMinOffsetOk

`func (o *ExpirationOffsetWithMinimum) GetMinOffsetOk() (*MinOffset, bool)`

GetMinOffsetOk returns a tuple with the MinOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinOffset

`func (o *ExpirationOffsetWithMinimum) SetMinOffset(v MinOffset)`

SetMinOffset sets MinOffset field to given value.

### HasMinOffset

`func (o *ExpirationOffsetWithMinimum) HasMinOffset() bool`

HasMinOffset returns a boolean if a field has been set.

### GetUnit

`func (o *ExpirationOffsetWithMinimum) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *ExpirationOffsetWithMinimum) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *ExpirationOffsetWithMinimum) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *ExpirationOffsetWithMinimum) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetValue

`func (o *ExpirationOffsetWithMinimum) GetValue() int32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ExpirationOffsetWithMinimum) GetValueOk() (*int32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ExpirationOffsetWithMinimum) SetValue(v int32)`

SetValue sets Value field to given value.

### HasValue

`func (o *ExpirationOffsetWithMinimum) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


