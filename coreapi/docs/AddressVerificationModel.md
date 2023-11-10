# AddressVerificationModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OnFile** | Pointer to [**AvsInformation**](AvsInformation.md) |  | [optional] 
**Request** | Pointer to [**AvsInformation**](AvsInformation.md) |  | [optional] 
**Response** | Pointer to [**Response**](Response.md) |  | [optional] 

## Methods

### NewAddressVerificationModel

`func NewAddressVerificationModel() *AddressVerificationModel`

NewAddressVerificationModel instantiates a new AddressVerificationModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressVerificationModelWithDefaults

`func NewAddressVerificationModelWithDefaults() *AddressVerificationModel`

NewAddressVerificationModelWithDefaults instantiates a new AddressVerificationModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOnFile

`func (o *AddressVerificationModel) GetOnFile() AvsInformation`

GetOnFile returns the OnFile field if non-nil, zero value otherwise.

### GetOnFileOk

`func (o *AddressVerificationModel) GetOnFileOk() (*AvsInformation, bool)`

GetOnFileOk returns a tuple with the OnFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnFile

`func (o *AddressVerificationModel) SetOnFile(v AvsInformation)`

SetOnFile sets OnFile field to given value.

### HasOnFile

`func (o *AddressVerificationModel) HasOnFile() bool`

HasOnFile returns a boolean if a field has been set.

### GetRequest

`func (o *AddressVerificationModel) GetRequest() AvsInformation`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *AddressVerificationModel) GetRequestOk() (*AvsInformation, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *AddressVerificationModel) SetRequest(v AvsInformation)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *AddressVerificationModel) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetResponse

`func (o *AddressVerificationModel) GetResponse() Response`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *AddressVerificationModel) GetResponseOk() (*Response, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *AddressVerificationModel) SetResponse(v Response)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *AddressVerificationModel) HasResponse() bool`

HasResponse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


