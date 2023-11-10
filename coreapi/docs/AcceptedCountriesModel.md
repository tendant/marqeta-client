# AcceptedCountriesModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountryCodes** | **[]string** | Comma-delimited list of accepted countries by ISO 3166 three-digit country code. | 
**CreatedTime** | Pointer to **time.Time** | Date and time when the accepted countries object was created, in UTC. This field is returned when included in your query. | [optional] 
**IsWhitelist** | **bool** | Specifies if the list of accepted countries in this object is an allow list. If set to &#x60;true&#x60;, transactions are accepted for all countries included in the object&#39;s &#x60;country_codes&#x60; array. If set to &#x60;false&#x60;, transactions are prohibited for all countries included in the object&#39;s &#x60;country_codes&#x60; array. | [default to false]
**LastModifiedTime** | Pointer to **time.Time** | Date and time when the accepted countries object was last updated, in UTC. This field is returned when included in your query. | [optional] 
**Name** | **string** | Name of the &#x60;acceptedcountries&#x60; object. | 
**Token** | Pointer to **string** | The unique identifier of the &#x60;acceptedcountries&#x60; object.  This field is always returned. | [optional] 

## Methods

### NewAcceptedCountriesModel

`func NewAcceptedCountriesModel(countryCodes []string, isWhitelist bool, name string, ) *AcceptedCountriesModel`

NewAcceptedCountriesModel instantiates a new AcceptedCountriesModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAcceptedCountriesModelWithDefaults

`func NewAcceptedCountriesModelWithDefaults() *AcceptedCountriesModel`

NewAcceptedCountriesModelWithDefaults instantiates a new AcceptedCountriesModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountryCodes

`func (o *AcceptedCountriesModel) GetCountryCodes() []string`

GetCountryCodes returns the CountryCodes field if non-nil, zero value otherwise.

### GetCountryCodesOk

`func (o *AcceptedCountriesModel) GetCountryCodesOk() (*[]string, bool)`

GetCountryCodesOk returns a tuple with the CountryCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCodes

`func (o *AcceptedCountriesModel) SetCountryCodes(v []string)`

SetCountryCodes sets CountryCodes field to given value.


### GetCreatedTime

`func (o *AcceptedCountriesModel) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *AcceptedCountriesModel) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *AcceptedCountriesModel) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *AcceptedCountriesModel) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetIsWhitelist

`func (o *AcceptedCountriesModel) GetIsWhitelist() bool`

GetIsWhitelist returns the IsWhitelist field if non-nil, zero value otherwise.

### GetIsWhitelistOk

`func (o *AcceptedCountriesModel) GetIsWhitelistOk() (*bool, bool)`

GetIsWhitelistOk returns a tuple with the IsWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWhitelist

`func (o *AcceptedCountriesModel) SetIsWhitelist(v bool)`

SetIsWhitelist sets IsWhitelist field to given value.


### GetLastModifiedTime

`func (o *AcceptedCountriesModel) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *AcceptedCountriesModel) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *AcceptedCountriesModel) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.

### HasLastModifiedTime

`func (o *AcceptedCountriesModel) HasLastModifiedTime() bool`

HasLastModifiedTime returns a boolean if a field has been set.

### GetName

`func (o *AcceptedCountriesModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AcceptedCountriesModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AcceptedCountriesModel) SetName(v string)`

SetName sets Name field to given value.


### GetToken

`func (o *AcceptedCountriesModel) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AcceptedCountriesModel) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AcceptedCountriesModel) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *AcceptedCountriesModel) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


