# AprScheduleEntryUpdateReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplyNextCycle** | Pointer to **bool** | Whether the APR can be ignored for the current billing cycle and applied on the next. | [optional] [default to false]
**EffectiveDate** | Pointer to **time.Time** | Date and time when the APR goes into effect, in UTC.  If you do not include a date-time value, the system uses the date and time when the API request was received.  *NOTE:* When passing multiple &#x60;schedule&#x60; objects, this field is required in all objects but the first. If you do not include &#x60;effective_date&#x60; in the first &#x60;schedule&#x60;, the system uses the date and time when the API request was received. | [optional] 
**Margin** | Pointer to **float32** | The number of percentage points added to the prime rate, used to calculate a variable value.  Used for variable values only. | [optional] 
**Type** | Pointer to **string** | Indicates whether the APR value is fixed or variable. | [optional] [default to "FIXED"]
**Value** | **float32** | Percentage value of the APR.  If the APR type is &#x60;FIXED&#x60;, this is the value of the fixed rate. If the APR type is &#x60;VARIABLE&#x60;, the value is calculated by adding the margin to the prime rate that was stored on Marqeta&#39;s credit platform when your credit program was created.  When backdating an APR, this value cannot be greater than the value of the effective APR on the backdated date. | 

## Methods

### NewAprScheduleEntryUpdateReq

`func NewAprScheduleEntryUpdateReq(value float32, ) *AprScheduleEntryUpdateReq`

NewAprScheduleEntryUpdateReq instantiates a new AprScheduleEntryUpdateReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAprScheduleEntryUpdateReqWithDefaults

`func NewAprScheduleEntryUpdateReqWithDefaults() *AprScheduleEntryUpdateReq`

NewAprScheduleEntryUpdateReqWithDefaults instantiates a new AprScheduleEntryUpdateReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplyNextCycle

`func (o *AprScheduleEntryUpdateReq) GetApplyNextCycle() bool`

GetApplyNextCycle returns the ApplyNextCycle field if non-nil, zero value otherwise.

### GetApplyNextCycleOk

`func (o *AprScheduleEntryUpdateReq) GetApplyNextCycleOk() (*bool, bool)`

GetApplyNextCycleOk returns a tuple with the ApplyNextCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyNextCycle

`func (o *AprScheduleEntryUpdateReq) SetApplyNextCycle(v bool)`

SetApplyNextCycle sets ApplyNextCycle field to given value.

### HasApplyNextCycle

`func (o *AprScheduleEntryUpdateReq) HasApplyNextCycle() bool`

HasApplyNextCycle returns a boolean if a field has been set.

### GetEffectiveDate

`func (o *AprScheduleEntryUpdateReq) GetEffectiveDate() time.Time`

GetEffectiveDate returns the EffectiveDate field if non-nil, zero value otherwise.

### GetEffectiveDateOk

`func (o *AprScheduleEntryUpdateReq) GetEffectiveDateOk() (*time.Time, bool)`

GetEffectiveDateOk returns a tuple with the EffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveDate

`func (o *AprScheduleEntryUpdateReq) SetEffectiveDate(v time.Time)`

SetEffectiveDate sets EffectiveDate field to given value.

### HasEffectiveDate

`func (o *AprScheduleEntryUpdateReq) HasEffectiveDate() bool`

HasEffectiveDate returns a boolean if a field has been set.

### GetMargin

`func (o *AprScheduleEntryUpdateReq) GetMargin() float32`

GetMargin returns the Margin field if non-nil, zero value otherwise.

### GetMarginOk

`func (o *AprScheduleEntryUpdateReq) GetMarginOk() (*float32, bool)`

GetMarginOk returns a tuple with the Margin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMargin

`func (o *AprScheduleEntryUpdateReq) SetMargin(v float32)`

SetMargin sets Margin field to given value.

### HasMargin

`func (o *AprScheduleEntryUpdateReq) HasMargin() bool`

HasMargin returns a boolean if a field has been set.

### GetType

`func (o *AprScheduleEntryUpdateReq) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AprScheduleEntryUpdateReq) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AprScheduleEntryUpdateReq) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AprScheduleEntryUpdateReq) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *AprScheduleEntryUpdateReq) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AprScheduleEntryUpdateReq) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AprScheduleEntryUpdateReq) SetValue(v float32)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


