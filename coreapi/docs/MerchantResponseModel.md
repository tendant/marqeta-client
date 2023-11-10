# MerchantResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** |  | [optional] [default to true]
**Address1** | Pointer to **string** |  | [optional] 
**Address2** | Pointer to **string** |  | [optional] 
**City** | Pointer to **string** |  | [optional] 
**Contact** | Pointer to **string** |  | [optional] 
**ContactEmail** | Pointer to **string** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**CreatedTime** | **time.Time** | yyyy-MM-ddTHH:mm:ssZ | 
**LastModifiedTime** | **time.Time** | yyyy-MM-ddTHH:mm:ssZ | 
**Latitude** | Pointer to **float32** |  | [optional] 
**Longitude** | Pointer to **float32** |  | [optional] 
**Name** | **string** |  | 
**PartialAuthFlag** | Pointer to **bool** |  | [optional] [default to true]
**Phone** | Pointer to **string** |  | [optional] 
**Province** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Token** | Pointer to **string** | The unique identifier of the merchant | [optional] 
**Zip** | Pointer to **string** |  | [optional] 

## Methods

### NewMerchantResponseModel

`func NewMerchantResponseModel(createdTime time.Time, lastModifiedTime time.Time, name string, ) *MerchantResponseModel`

NewMerchantResponseModel instantiates a new MerchantResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantResponseModelWithDefaults

`func NewMerchantResponseModelWithDefaults() *MerchantResponseModel`

NewMerchantResponseModelWithDefaults instantiates a new MerchantResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *MerchantResponseModel) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *MerchantResponseModel) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *MerchantResponseModel) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *MerchantResponseModel) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAddress1

`func (o *MerchantResponseModel) GetAddress1() string`

GetAddress1 returns the Address1 field if non-nil, zero value otherwise.

### GetAddress1Ok

`func (o *MerchantResponseModel) GetAddress1Ok() (*string, bool)`

GetAddress1Ok returns a tuple with the Address1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress1

`func (o *MerchantResponseModel) SetAddress1(v string)`

SetAddress1 sets Address1 field to given value.

### HasAddress1

`func (o *MerchantResponseModel) HasAddress1() bool`

HasAddress1 returns a boolean if a field has been set.

### GetAddress2

`func (o *MerchantResponseModel) GetAddress2() string`

GetAddress2 returns the Address2 field if non-nil, zero value otherwise.

### GetAddress2Ok

`func (o *MerchantResponseModel) GetAddress2Ok() (*string, bool)`

GetAddress2Ok returns a tuple with the Address2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress2

`func (o *MerchantResponseModel) SetAddress2(v string)`

SetAddress2 sets Address2 field to given value.

### HasAddress2

`func (o *MerchantResponseModel) HasAddress2() bool`

HasAddress2 returns a boolean if a field has been set.

### GetCity

`func (o *MerchantResponseModel) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *MerchantResponseModel) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *MerchantResponseModel) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *MerchantResponseModel) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetContact

`func (o *MerchantResponseModel) GetContact() string`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *MerchantResponseModel) GetContactOk() (*string, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *MerchantResponseModel) SetContact(v string)`

SetContact sets Contact field to given value.

### HasContact

`func (o *MerchantResponseModel) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetContactEmail

`func (o *MerchantResponseModel) GetContactEmail() string`

GetContactEmail returns the ContactEmail field if non-nil, zero value otherwise.

### GetContactEmailOk

`func (o *MerchantResponseModel) GetContactEmailOk() (*string, bool)`

GetContactEmailOk returns a tuple with the ContactEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactEmail

`func (o *MerchantResponseModel) SetContactEmail(v string)`

SetContactEmail sets ContactEmail field to given value.

### HasContactEmail

`func (o *MerchantResponseModel) HasContactEmail() bool`

HasContactEmail returns a boolean if a field has been set.

### GetCountry

`func (o *MerchantResponseModel) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *MerchantResponseModel) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *MerchantResponseModel) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *MerchantResponseModel) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCreatedTime

`func (o *MerchantResponseModel) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *MerchantResponseModel) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *MerchantResponseModel) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetLastModifiedTime

`func (o *MerchantResponseModel) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *MerchantResponseModel) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *MerchantResponseModel) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.


### GetLatitude

`func (o *MerchantResponseModel) GetLatitude() float32`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *MerchantResponseModel) GetLatitudeOk() (*float32, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *MerchantResponseModel) SetLatitude(v float32)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *MerchantResponseModel) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### GetLongitude

`func (o *MerchantResponseModel) GetLongitude() float32`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *MerchantResponseModel) GetLongitudeOk() (*float32, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *MerchantResponseModel) SetLongitude(v float32)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *MerchantResponseModel) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### GetName

`func (o *MerchantResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MerchantResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MerchantResponseModel) SetName(v string)`

SetName sets Name field to given value.


### GetPartialAuthFlag

`func (o *MerchantResponseModel) GetPartialAuthFlag() bool`

GetPartialAuthFlag returns the PartialAuthFlag field if non-nil, zero value otherwise.

### GetPartialAuthFlagOk

`func (o *MerchantResponseModel) GetPartialAuthFlagOk() (*bool, bool)`

GetPartialAuthFlagOk returns a tuple with the PartialAuthFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartialAuthFlag

`func (o *MerchantResponseModel) SetPartialAuthFlag(v bool)`

SetPartialAuthFlag sets PartialAuthFlag field to given value.

### HasPartialAuthFlag

`func (o *MerchantResponseModel) HasPartialAuthFlag() bool`

HasPartialAuthFlag returns a boolean if a field has been set.

### GetPhone

`func (o *MerchantResponseModel) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *MerchantResponseModel) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *MerchantResponseModel) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *MerchantResponseModel) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetProvince

`func (o *MerchantResponseModel) GetProvince() string`

GetProvince returns the Province field if non-nil, zero value otherwise.

### GetProvinceOk

`func (o *MerchantResponseModel) GetProvinceOk() (*string, bool)`

GetProvinceOk returns a tuple with the Province field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvince

`func (o *MerchantResponseModel) SetProvince(v string)`

SetProvince sets Province field to given value.

### HasProvince

`func (o *MerchantResponseModel) HasProvince() bool`

HasProvince returns a boolean if a field has been set.

### GetState

`func (o *MerchantResponseModel) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MerchantResponseModel) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *MerchantResponseModel) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *MerchantResponseModel) HasState() bool`

HasState returns a boolean if a field has been set.

### GetToken

`func (o *MerchantResponseModel) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *MerchantResponseModel) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *MerchantResponseModel) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *MerchantResponseModel) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetZip

`func (o *MerchantResponseModel) GetZip() string`

GetZip returns the Zip field if non-nil, zero value otherwise.

### GetZipOk

`func (o *MerchantResponseModel) GetZipOk() (*string, bool)`

GetZipOk returns a tuple with the Zip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZip

`func (o *MerchantResponseModel) SetZip(v string)`

SetZip sets Zip field to given value.

### HasZip

`func (o *MerchantResponseModel) HasZip() bool`

HasZip returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


