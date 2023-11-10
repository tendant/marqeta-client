# ErrorDetailsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationToken** | **string** | Unique identifier of the application that resulted in an error. | 
**Message** | **string** | Message describing the error. | 
**Token** | **string** | Unique identifier of the error details. | 
**Type** | **string** | Type of error. | 

## Methods

### NewErrorDetailsResponse

`func NewErrorDetailsResponse(applicationToken string, message string, token string, type_ string, ) *ErrorDetailsResponse`

NewErrorDetailsResponse instantiates a new ErrorDetailsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorDetailsResponseWithDefaults

`func NewErrorDetailsResponseWithDefaults() *ErrorDetailsResponse`

NewErrorDetailsResponseWithDefaults instantiates a new ErrorDetailsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationToken

`func (o *ErrorDetailsResponse) GetApplicationToken() string`

GetApplicationToken returns the ApplicationToken field if non-nil, zero value otherwise.

### GetApplicationTokenOk

`func (o *ErrorDetailsResponse) GetApplicationTokenOk() (*string, bool)`

GetApplicationTokenOk returns a tuple with the ApplicationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationToken

`func (o *ErrorDetailsResponse) SetApplicationToken(v string)`

SetApplicationToken sets ApplicationToken field to given value.


### GetMessage

`func (o *ErrorDetailsResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ErrorDetailsResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ErrorDetailsResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetToken

`func (o *ErrorDetailsResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ErrorDetailsResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ErrorDetailsResponse) SetToken(v string)`

SetToken sets Token field to given value.


### GetType

`func (o *ErrorDetailsResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ErrorDetailsResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ErrorDetailsResponse) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


