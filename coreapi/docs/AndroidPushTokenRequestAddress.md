# AndroidPushTokenRequestAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address1** | Pointer to **string** | Street address of the cardholder. | [optional] 
**Address2** | Pointer to **string** | Additional address information for the cardholder, such as a suite or apartment number.  &#x60;Suite 600&#x60;, for example. | [optional] 
**City** | Pointer to **string** | City of the cardholder. | [optional] 
**Country** | Pointer to **string** | Two-character link:https://www.iso.org/iso-3166-country-codes.html[ISO alpha-2 country code, window&#x3D;\&quot;_blank\&quot;]. &#x60;US&#x60;, for example. | [optional] 
**Name** | Pointer to **string** | Name of the cardholder. | [optional] 
**Phone** | Pointer to **string** | Telephone number of the cardholder. | [optional] 
**PostalCode** | Pointer to **string** | Postal code of the cardholder, such as a United States ZIP code. &#x60;94612&#x60;, for example. | [optional] 
**State** | Pointer to **string** | Two-character &lt;&lt;/core-api/kyc-verification#_valid_state_provincial_and_territorial_abbreviations, state or province code&gt;&gt;. &#x60;CA&#x60;, for example. | [optional] 
**Zip** | Pointer to **string** |  | [optional] 

## Methods

### NewAndroidPushTokenRequestAddress

`func NewAndroidPushTokenRequestAddress() *AndroidPushTokenRequestAddress`

NewAndroidPushTokenRequestAddress instantiates a new AndroidPushTokenRequestAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAndroidPushTokenRequestAddressWithDefaults

`func NewAndroidPushTokenRequestAddressWithDefaults() *AndroidPushTokenRequestAddress`

NewAndroidPushTokenRequestAddressWithDefaults instantiates a new AndroidPushTokenRequestAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress1

`func (o *AndroidPushTokenRequestAddress) GetAddress1() string`

GetAddress1 returns the Address1 field if non-nil, zero value otherwise.

### GetAddress1Ok

`func (o *AndroidPushTokenRequestAddress) GetAddress1Ok() (*string, bool)`

GetAddress1Ok returns a tuple with the Address1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress1

`func (o *AndroidPushTokenRequestAddress) SetAddress1(v string)`

SetAddress1 sets Address1 field to given value.

### HasAddress1

`func (o *AndroidPushTokenRequestAddress) HasAddress1() bool`

HasAddress1 returns a boolean if a field has been set.

### GetAddress2

`func (o *AndroidPushTokenRequestAddress) GetAddress2() string`

GetAddress2 returns the Address2 field if non-nil, zero value otherwise.

### GetAddress2Ok

`func (o *AndroidPushTokenRequestAddress) GetAddress2Ok() (*string, bool)`

GetAddress2Ok returns a tuple with the Address2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress2

`func (o *AndroidPushTokenRequestAddress) SetAddress2(v string)`

SetAddress2 sets Address2 field to given value.

### HasAddress2

`func (o *AndroidPushTokenRequestAddress) HasAddress2() bool`

HasAddress2 returns a boolean if a field has been set.

### GetCity

`func (o *AndroidPushTokenRequestAddress) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *AndroidPushTokenRequestAddress) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *AndroidPushTokenRequestAddress) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *AndroidPushTokenRequestAddress) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCountry

`func (o *AndroidPushTokenRequestAddress) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *AndroidPushTokenRequestAddress) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *AndroidPushTokenRequestAddress) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *AndroidPushTokenRequestAddress) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetName

`func (o *AndroidPushTokenRequestAddress) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AndroidPushTokenRequestAddress) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AndroidPushTokenRequestAddress) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AndroidPushTokenRequestAddress) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPhone

`func (o *AndroidPushTokenRequestAddress) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *AndroidPushTokenRequestAddress) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *AndroidPushTokenRequestAddress) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *AndroidPushTokenRequestAddress) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetPostalCode

`func (o *AndroidPushTokenRequestAddress) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *AndroidPushTokenRequestAddress) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *AndroidPushTokenRequestAddress) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *AndroidPushTokenRequestAddress) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetState

`func (o *AndroidPushTokenRequestAddress) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AndroidPushTokenRequestAddress) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AndroidPushTokenRequestAddress) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *AndroidPushTokenRequestAddress) HasState() bool`

HasState returns a boolean if a field has been set.

### GetZip

`func (o *AndroidPushTokenRequestAddress) GetZip() string`

GetZip returns the Zip field if non-nil, zero value otherwise.

### GetZipOk

`func (o *AndroidPushTokenRequestAddress) GetZipOk() (*string, bool)`

GetZipOk returns a tuple with the Zip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZip

`func (o *AndroidPushTokenRequestAddress) SetZip(v string)`

SetZip sets Zip field to given value.

### HasZip

`func (o *AndroidPushTokenRequestAddress) HasZip() bool`

HasZip returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


