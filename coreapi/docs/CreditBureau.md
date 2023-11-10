# CreditBureau

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | [**CreditBureauAddress**](CreditBureauAddress.md) |  | 
**Name** | **string** | Name of the credit bureau used to obtain the user&#39;s credit data. | 
**PhoneNumber** | **string** | Phone number of the credit bureau. | 
**Website** | **string** | Website of the credit bureau. | 

## Methods

### NewCreditBureau

`func NewCreditBureau(address CreditBureauAddress, name string, phoneNumber string, website string, ) *CreditBureau`

NewCreditBureau instantiates a new CreditBureau object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditBureauWithDefaults

`func NewCreditBureauWithDefaults() *CreditBureau`

NewCreditBureauWithDefaults instantiates a new CreditBureau object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *CreditBureau) GetAddress() CreditBureauAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *CreditBureau) GetAddressOk() (*CreditBureauAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *CreditBureau) SetAddress(v CreditBureauAddress)`

SetAddress sets Address field to given value.


### GetName

`func (o *CreditBureau) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreditBureau) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreditBureau) SetName(v string)`

SetName sets Name field to given value.


### GetPhoneNumber

`func (o *CreditBureau) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *CreditBureau) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *CreditBureau) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.


### GetWebsite

`func (o *CreditBureau) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *CreditBureau) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *CreditBureau) SetWebsite(v string)`

SetWebsite sets Website field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


