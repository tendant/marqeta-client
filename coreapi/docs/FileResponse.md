# FileResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileType** | [**FileType**](FileType.md) |  | 
**Links** | [**FileLinks**](FileLinks.md) |  | 
**TrackingToken** | Pointer to **string** | Unique identifier used to acknowledge that the file has been disclosed to the applicant.  If the file was already disclosed, a null value is returned.  *NOTE*: The tracking token is only valid for 14 calendar days. | [optional] 

## Methods

### NewFileResponse

`func NewFileResponse(fileType FileType, links FileLinks, ) *FileResponse`

NewFileResponse instantiates a new FileResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileResponseWithDefaults

`func NewFileResponseWithDefaults() *FileResponse`

NewFileResponseWithDefaults instantiates a new FileResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileType

`func (o *FileResponse) GetFileType() FileType`

GetFileType returns the FileType field if non-nil, zero value otherwise.

### GetFileTypeOk

`func (o *FileResponse) GetFileTypeOk() (*FileType, bool)`

GetFileTypeOk returns a tuple with the FileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileType

`func (o *FileResponse) SetFileType(v FileType)`

SetFileType sets FileType field to given value.


### GetLinks

`func (o *FileResponse) GetLinks() FileLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *FileResponse) GetLinksOk() (*FileLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *FileResponse) SetLinks(v FileLinks)`

SetLinks sets Links field to given value.


### GetTrackingToken

`func (o *FileResponse) GetTrackingToken() string`

GetTrackingToken returns the TrackingToken field if non-nil, zero value otherwise.

### GetTrackingTokenOk

`func (o *FileResponse) GetTrackingTokenOk() (*string, bool)`

GetTrackingTokenOk returns a tuple with the TrackingToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingToken

`func (o *FileResponse) SetTrackingToken(v string)`

SetTrackingToken sets TrackingToken field to given value.

### HasTrackingToken

`func (o *FileResponse) HasTrackingToken() bool`

HasTrackingToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


