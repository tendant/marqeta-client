# StoreResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** |  | [optional] [default to true]
**Address1** | **string** |  | 
**Address2** | Pointer to **string** |  | [optional] 
**City** | **string** |  | 
**Contact** | Pointer to **string** |  | [optional] 
**ContactEmail** | Pointer to **string** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**CreatedTime** | **time.Time** | yyyy-MM-ddTHH:mm:ssZ | 
**KeyedAuthCvvEnforced** | Pointer to **bool** |  | [optional] [default to false]
**LastModifiedTime** | **time.Time** | yyyy-MM-ddTHH:mm:ssZ | 
**Latitude** | Pointer to **float32** |  | [optional] 
**Longitude** | Pointer to **float32** |  | [optional] 
**MerchantToken** | **string** |  | 
**Mid** | **string** |  | 
**Name** | **string** |  | 
**NetworkMid** | Pointer to **string** |  | [optional] 
**PartialApprovalCapable** | Pointer to **bool** |  | [optional] [default to false]
**PartialAuthFlag** | Pointer to **bool** | 1 char max | [optional] [default to true]
**Phone** | Pointer to **string** |  | [optional] 
**PostalCode** | Pointer to **string** |  | [optional] 
**Province** | Pointer to **string** |  | [optional] 
**State** | **string** |  | 
**Token** | Pointer to **string** | The unique identifier of the merchant | [optional] 
**Zip** | Pointer to **string** |  | [optional] 

## Methods

### NewStoreResponseModel

`func NewStoreResponseModel(address1 string, city string, createdTime time.Time, lastModifiedTime time.Time, merchantToken string, mid string, name string, state string, ) *StoreResponseModel`

NewStoreResponseModel instantiates a new StoreResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoreResponseModelWithDefaults

`func NewStoreResponseModelWithDefaults() *StoreResponseModel`

NewStoreResponseModelWithDefaults instantiates a new StoreResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *StoreResponseModel) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *StoreResponseModel) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *StoreResponseModel) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *StoreResponseModel) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAddress1

`func (o *StoreResponseModel) GetAddress1() string`

GetAddress1 returns the Address1 field if non-nil, zero value otherwise.

### GetAddress1Ok

`func (o *StoreResponseModel) GetAddress1Ok() (*string, bool)`

GetAddress1Ok returns a tuple with the Address1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress1

`func (o *StoreResponseModel) SetAddress1(v string)`

SetAddress1 sets Address1 field to given value.


### GetAddress2

`func (o *StoreResponseModel) GetAddress2() string`

GetAddress2 returns the Address2 field if non-nil, zero value otherwise.

### GetAddress2Ok

`func (o *StoreResponseModel) GetAddress2Ok() (*string, bool)`

GetAddress2Ok returns a tuple with the Address2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress2

`func (o *StoreResponseModel) SetAddress2(v string)`

SetAddress2 sets Address2 field to given value.

### HasAddress2

`func (o *StoreResponseModel) HasAddress2() bool`

HasAddress2 returns a boolean if a field has been set.

### GetCity

`func (o *StoreResponseModel) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *StoreResponseModel) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *StoreResponseModel) SetCity(v string)`

SetCity sets City field to given value.


### GetContact

`func (o *StoreResponseModel) GetContact() string`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *StoreResponseModel) GetContactOk() (*string, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *StoreResponseModel) SetContact(v string)`

SetContact sets Contact field to given value.

### HasContact

`func (o *StoreResponseModel) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetContactEmail

`func (o *StoreResponseModel) GetContactEmail() string`

GetContactEmail returns the ContactEmail field if non-nil, zero value otherwise.

### GetContactEmailOk

`func (o *StoreResponseModel) GetContactEmailOk() (*string, bool)`

GetContactEmailOk returns a tuple with the ContactEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactEmail

`func (o *StoreResponseModel) SetContactEmail(v string)`

SetContactEmail sets ContactEmail field to given value.

### HasContactEmail

`func (o *StoreResponseModel) HasContactEmail() bool`

HasContactEmail returns a boolean if a field has been set.

### GetCountry

`func (o *StoreResponseModel) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *StoreResponseModel) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *StoreResponseModel) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *StoreResponseModel) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCreatedTime

`func (o *StoreResponseModel) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *StoreResponseModel) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *StoreResponseModel) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetKeyedAuthCvvEnforced

`func (o *StoreResponseModel) GetKeyedAuthCvvEnforced() bool`

GetKeyedAuthCvvEnforced returns the KeyedAuthCvvEnforced field if non-nil, zero value otherwise.

### GetKeyedAuthCvvEnforcedOk

`func (o *StoreResponseModel) GetKeyedAuthCvvEnforcedOk() (*bool, bool)`

GetKeyedAuthCvvEnforcedOk returns a tuple with the KeyedAuthCvvEnforced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyedAuthCvvEnforced

`func (o *StoreResponseModel) SetKeyedAuthCvvEnforced(v bool)`

SetKeyedAuthCvvEnforced sets KeyedAuthCvvEnforced field to given value.

### HasKeyedAuthCvvEnforced

`func (o *StoreResponseModel) HasKeyedAuthCvvEnforced() bool`

HasKeyedAuthCvvEnforced returns a boolean if a field has been set.

### GetLastModifiedTime

`func (o *StoreResponseModel) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *StoreResponseModel) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *StoreResponseModel) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.


### GetLatitude

`func (o *StoreResponseModel) GetLatitude() float32`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *StoreResponseModel) GetLatitudeOk() (*float32, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *StoreResponseModel) SetLatitude(v float32)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *StoreResponseModel) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### GetLongitude

`func (o *StoreResponseModel) GetLongitude() float32`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *StoreResponseModel) GetLongitudeOk() (*float32, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *StoreResponseModel) SetLongitude(v float32)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *StoreResponseModel) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### GetMerchantToken

`func (o *StoreResponseModel) GetMerchantToken() string`

GetMerchantToken returns the MerchantToken field if non-nil, zero value otherwise.

### GetMerchantTokenOk

`func (o *StoreResponseModel) GetMerchantTokenOk() (*string, bool)`

GetMerchantTokenOk returns a tuple with the MerchantToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantToken

`func (o *StoreResponseModel) SetMerchantToken(v string)`

SetMerchantToken sets MerchantToken field to given value.


### GetMid

`func (o *StoreResponseModel) GetMid() string`

GetMid returns the Mid field if non-nil, zero value otherwise.

### GetMidOk

`func (o *StoreResponseModel) GetMidOk() (*string, bool)`

GetMidOk returns a tuple with the Mid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMid

`func (o *StoreResponseModel) SetMid(v string)`

SetMid sets Mid field to given value.


### GetName

`func (o *StoreResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StoreResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StoreResponseModel) SetName(v string)`

SetName sets Name field to given value.


### GetNetworkMid

`func (o *StoreResponseModel) GetNetworkMid() string`

GetNetworkMid returns the NetworkMid field if non-nil, zero value otherwise.

### GetNetworkMidOk

`func (o *StoreResponseModel) GetNetworkMidOk() (*string, bool)`

GetNetworkMidOk returns a tuple with the NetworkMid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkMid

`func (o *StoreResponseModel) SetNetworkMid(v string)`

SetNetworkMid sets NetworkMid field to given value.

### HasNetworkMid

`func (o *StoreResponseModel) HasNetworkMid() bool`

HasNetworkMid returns a boolean if a field has been set.

### GetPartialApprovalCapable

`func (o *StoreResponseModel) GetPartialApprovalCapable() bool`

GetPartialApprovalCapable returns the PartialApprovalCapable field if non-nil, zero value otherwise.

### GetPartialApprovalCapableOk

`func (o *StoreResponseModel) GetPartialApprovalCapableOk() (*bool, bool)`

GetPartialApprovalCapableOk returns a tuple with the PartialApprovalCapable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartialApprovalCapable

`func (o *StoreResponseModel) SetPartialApprovalCapable(v bool)`

SetPartialApprovalCapable sets PartialApprovalCapable field to given value.

### HasPartialApprovalCapable

`func (o *StoreResponseModel) HasPartialApprovalCapable() bool`

HasPartialApprovalCapable returns a boolean if a field has been set.

### GetPartialAuthFlag

`func (o *StoreResponseModel) GetPartialAuthFlag() bool`

GetPartialAuthFlag returns the PartialAuthFlag field if non-nil, zero value otherwise.

### GetPartialAuthFlagOk

`func (o *StoreResponseModel) GetPartialAuthFlagOk() (*bool, bool)`

GetPartialAuthFlagOk returns a tuple with the PartialAuthFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartialAuthFlag

`func (o *StoreResponseModel) SetPartialAuthFlag(v bool)`

SetPartialAuthFlag sets PartialAuthFlag field to given value.

### HasPartialAuthFlag

`func (o *StoreResponseModel) HasPartialAuthFlag() bool`

HasPartialAuthFlag returns a boolean if a field has been set.

### GetPhone

`func (o *StoreResponseModel) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *StoreResponseModel) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *StoreResponseModel) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *StoreResponseModel) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetPostalCode

`func (o *StoreResponseModel) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *StoreResponseModel) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *StoreResponseModel) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *StoreResponseModel) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetProvince

`func (o *StoreResponseModel) GetProvince() string`

GetProvince returns the Province field if non-nil, zero value otherwise.

### GetProvinceOk

`func (o *StoreResponseModel) GetProvinceOk() (*string, bool)`

GetProvinceOk returns a tuple with the Province field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvince

`func (o *StoreResponseModel) SetProvince(v string)`

SetProvince sets Province field to given value.

### HasProvince

`func (o *StoreResponseModel) HasProvince() bool`

HasProvince returns a boolean if a field has been set.

### GetState

`func (o *StoreResponseModel) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *StoreResponseModel) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *StoreResponseModel) SetState(v string)`

SetState sets State field to given value.


### GetToken

`func (o *StoreResponseModel) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *StoreResponseModel) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *StoreResponseModel) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *StoreResponseModel) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetZip

`func (o *StoreResponseModel) GetZip() string`

GetZip returns the Zip field if non-nil, zero value otherwise.

### GetZipOk

`func (o *StoreResponseModel) GetZipOk() (*string, bool)`

GetZipOk returns a tuple with the Zip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZip

`func (o *StoreResponseModel) SetZip(v string)`

SetZip sets Zip field to given value.

### HasZip

`func (o *StoreResponseModel) HasZip() bool`

HasZip returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


