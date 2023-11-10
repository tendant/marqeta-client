# Airline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepartDate** | Pointer to **time.Time** | The date and time of departure. | [optional] 
**OriginationCity** | Pointer to **string** | The city where the flight originates. | [optional] 
**PassengerName** | Pointer to **string** | The name of the passenger. | [optional] 

## Methods

### NewAirline

`func NewAirline() *Airline`

NewAirline instantiates a new Airline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAirlineWithDefaults

`func NewAirlineWithDefaults() *Airline`

NewAirlineWithDefaults instantiates a new Airline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepartDate

`func (o *Airline) GetDepartDate() time.Time`

GetDepartDate returns the DepartDate field if non-nil, zero value otherwise.

### GetDepartDateOk

`func (o *Airline) GetDepartDateOk() (*time.Time, bool)`

GetDepartDateOk returns a tuple with the DepartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartDate

`func (o *Airline) SetDepartDate(v time.Time)`

SetDepartDate sets DepartDate field to given value.

### HasDepartDate

`func (o *Airline) HasDepartDate() bool`

HasDepartDate returns a boolean if a field has been set.

### GetOriginationCity

`func (o *Airline) GetOriginationCity() string`

GetOriginationCity returns the OriginationCity field if non-nil, zero value otherwise.

### GetOriginationCityOk

`func (o *Airline) GetOriginationCityOk() (*string, bool)`

GetOriginationCityOk returns a tuple with the OriginationCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginationCity

`func (o *Airline) SetOriginationCity(v string)`

SetOriginationCity sets OriginationCity field to given value.

### HasOriginationCity

`func (o *Airline) HasOriginationCity() bool`

HasOriginationCity returns a boolean if a field has been set.

### GetPassengerName

`func (o *Airline) GetPassengerName() string`

GetPassengerName returns the PassengerName field if non-nil, zero value otherwise.

### GetPassengerNameOk

`func (o *Airline) GetPassengerNameOk() (*string, bool)`

GetPassengerNameOk returns a tuple with the PassengerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassengerName

`func (o *Airline) SetPassengerName(v string)`

SetPassengerName sets PassengerName field to given value.

### HasPassengerName

`func (o *Airline) HasPassengerName() bool`

HasPassengerName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


