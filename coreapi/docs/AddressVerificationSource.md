# AddressVerificationSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OnFile** | Pointer to [**AvsInformation**](AvsInformation.md) |  | [optional] 
**Response** | Pointer to [**Response**](Response.md) |  | [optional] 

## Methods

### NewAddressVerificationSource

`func NewAddressVerificationSource() *AddressVerificationSource`

NewAddressVerificationSource instantiates a new AddressVerificationSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressVerificationSourceWithDefaults

`func NewAddressVerificationSourceWithDefaults() *AddressVerificationSource`

NewAddressVerificationSourceWithDefaults instantiates a new AddressVerificationSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOnFile

`func (o *AddressVerificationSource) GetOnFile() AvsInformation`

GetOnFile returns the OnFile field if non-nil, zero value otherwise.

### GetOnFileOk

`func (o *AddressVerificationSource) GetOnFileOk() (*AvsInformation, bool)`

GetOnFileOk returns a tuple with the OnFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnFile

`func (o *AddressVerificationSource) SetOnFile(v AvsInformation)`

SetOnFile sets OnFile field to given value.

### HasOnFile

`func (o *AddressVerificationSource) HasOnFile() bool`

HasOnFile returns a boolean if a field has been set.

### GetResponse

`func (o *AddressVerificationSource) GetResponse() Response`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *AddressVerificationSource) GetResponseOk() (*Response, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *AddressVerificationSource) SetResponse(v Response)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *AddressVerificationSource) HasResponse() bool`

HasResponse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


