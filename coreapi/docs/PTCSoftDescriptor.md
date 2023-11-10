# PTCSoftDescriptor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | [**PTCAddress**](PTCAddress.md) |  | 
**Email** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Phone** | Pointer to [**PTCPhone**](PTCPhone.md) |  | [optional] 

## Methods

### NewPTCSoftDescriptor

`func NewPTCSoftDescriptor(address PTCAddress, name string, ) *PTCSoftDescriptor`

NewPTCSoftDescriptor instantiates a new PTCSoftDescriptor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPTCSoftDescriptorWithDefaults

`func NewPTCSoftDescriptorWithDefaults() *PTCSoftDescriptor`

NewPTCSoftDescriptorWithDefaults instantiates a new PTCSoftDescriptor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *PTCSoftDescriptor) GetAddress() PTCAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *PTCSoftDescriptor) GetAddressOk() (*PTCAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *PTCSoftDescriptor) SetAddress(v PTCAddress)`

SetAddress sets Address field to given value.


### GetEmail

`func (o *PTCSoftDescriptor) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PTCSoftDescriptor) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PTCSoftDescriptor) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *PTCSoftDescriptor) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetName

`func (o *PTCSoftDescriptor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PTCSoftDescriptor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PTCSoftDescriptor) SetName(v string)`

SetName sets Name field to given value.


### GetPhone

`func (o *PTCSoftDescriptor) GetPhone() PTCPhone`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *PTCSoftDescriptor) GetPhoneOk() (*PTCPhone, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *PTCSoftDescriptor) SetPhone(v PTCPhone)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *PTCSoftDescriptor) HasPhone() bool`

HasPhone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


