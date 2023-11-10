# InterestCalculation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationOfCredits** | [**ApplicationOfCredits**](ApplicationOfCredits.md) |  | 
**DayCount** | **string** | Day-count convention. | 
**ExcludeTranTypes** | Pointer to **[]string** | One or more transactions that are excluded from current billing period&#39;s interest charge, but included in next. | [optional] [default to ["ANNUAL_FEE","MONTHLY_FEE","LATE_PAYMENT_FEE","CASH_BACK_STATEMENT_CREDIT"]]
**GraceDaysApplication** | **string** | Determines the last day of grace period based on which interest charges are calculated. | 
**InterestApplication** | **[]string** | One or more balance types on which interest is applied. | [default to ["PRINCIPAL","FEES"]]
**InterestOnGraceReactivation** | [**InterestOnGraceReactivationEnum**](InterestOnGraceReactivationEnum.md) |  | 
**Method** | **string** | Method of interest calculation. | 
**MinimumInterest** | **float32** | When interest is applied, this value determines the minimum amount of interest that can be charged. | 

## Methods

### NewInterestCalculation

`func NewInterestCalculation(applicationOfCredits ApplicationOfCredits, dayCount string, graceDaysApplication string, interestApplication []string, interestOnGraceReactivation InterestOnGraceReactivationEnum, method string, minimumInterest float32, ) *InterestCalculation`

NewInterestCalculation instantiates a new InterestCalculation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInterestCalculationWithDefaults

`func NewInterestCalculationWithDefaults() *InterestCalculation`

NewInterestCalculationWithDefaults instantiates a new InterestCalculation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationOfCredits

`func (o *InterestCalculation) GetApplicationOfCredits() ApplicationOfCredits`

GetApplicationOfCredits returns the ApplicationOfCredits field if non-nil, zero value otherwise.

### GetApplicationOfCreditsOk

`func (o *InterestCalculation) GetApplicationOfCreditsOk() (*ApplicationOfCredits, bool)`

GetApplicationOfCreditsOk returns a tuple with the ApplicationOfCredits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationOfCredits

`func (o *InterestCalculation) SetApplicationOfCredits(v ApplicationOfCredits)`

SetApplicationOfCredits sets ApplicationOfCredits field to given value.


### GetDayCount

`func (o *InterestCalculation) GetDayCount() string`

GetDayCount returns the DayCount field if non-nil, zero value otherwise.

### GetDayCountOk

`func (o *InterestCalculation) GetDayCountOk() (*string, bool)`

GetDayCountOk returns a tuple with the DayCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayCount

`func (o *InterestCalculation) SetDayCount(v string)`

SetDayCount sets DayCount field to given value.


### GetExcludeTranTypes

`func (o *InterestCalculation) GetExcludeTranTypes() []string`

GetExcludeTranTypes returns the ExcludeTranTypes field if non-nil, zero value otherwise.

### GetExcludeTranTypesOk

`func (o *InterestCalculation) GetExcludeTranTypesOk() (*[]string, bool)`

GetExcludeTranTypesOk returns a tuple with the ExcludeTranTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeTranTypes

`func (o *InterestCalculation) SetExcludeTranTypes(v []string)`

SetExcludeTranTypes sets ExcludeTranTypes field to given value.

### HasExcludeTranTypes

`func (o *InterestCalculation) HasExcludeTranTypes() bool`

HasExcludeTranTypes returns a boolean if a field has been set.

### GetGraceDaysApplication

`func (o *InterestCalculation) GetGraceDaysApplication() string`

GetGraceDaysApplication returns the GraceDaysApplication field if non-nil, zero value otherwise.

### GetGraceDaysApplicationOk

`func (o *InterestCalculation) GetGraceDaysApplicationOk() (*string, bool)`

GetGraceDaysApplicationOk returns a tuple with the GraceDaysApplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraceDaysApplication

`func (o *InterestCalculation) SetGraceDaysApplication(v string)`

SetGraceDaysApplication sets GraceDaysApplication field to given value.


### GetInterestApplication

`func (o *InterestCalculation) GetInterestApplication() []string`

GetInterestApplication returns the InterestApplication field if non-nil, zero value otherwise.

### GetInterestApplicationOk

`func (o *InterestCalculation) GetInterestApplicationOk() (*[]string, bool)`

GetInterestApplicationOk returns a tuple with the InterestApplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestApplication

`func (o *InterestCalculation) SetInterestApplication(v []string)`

SetInterestApplication sets InterestApplication field to given value.


### GetInterestOnGraceReactivation

`func (o *InterestCalculation) GetInterestOnGraceReactivation() InterestOnGraceReactivationEnum`

GetInterestOnGraceReactivation returns the InterestOnGraceReactivation field if non-nil, zero value otherwise.

### GetInterestOnGraceReactivationOk

`func (o *InterestCalculation) GetInterestOnGraceReactivationOk() (*InterestOnGraceReactivationEnum, bool)`

GetInterestOnGraceReactivationOk returns a tuple with the InterestOnGraceReactivation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestOnGraceReactivation

`func (o *InterestCalculation) SetInterestOnGraceReactivation(v InterestOnGraceReactivationEnum)`

SetInterestOnGraceReactivation sets InterestOnGraceReactivation field to given value.


### GetMethod

`func (o *InterestCalculation) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *InterestCalculation) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *InterestCalculation) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetMinimumInterest

`func (o *InterestCalculation) GetMinimumInterest() float32`

GetMinimumInterest returns the MinimumInterest field if non-nil, zero value otherwise.

### GetMinimumInterestOk

`func (o *InterestCalculation) GetMinimumInterestOk() (*float32, bool)`

GetMinimumInterestOk returns a tuple with the MinimumInterest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumInterest

`func (o *InterestCalculation) SetMinimumInterest(v float32)`

SetMinimumInterest sets MinimumInterest field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


