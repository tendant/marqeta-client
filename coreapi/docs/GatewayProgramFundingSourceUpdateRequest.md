# GatewayProgramFundingSourceUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Indicates whether the program gateway funding source is active. | [optional] 
**BasicAuthPassword** | **string** | Password for authenticating your environment. | 
**BasicAuthUsername** | **string** | Username for authenticating your environment. | 
**CustomHeader** | Pointer to **map[string]string** | Additional custom information included in the HTTP header. For example, this might contain security information, along with Basic Authentication, when making a JIT Funding request. Custom headers also appear in the associated webhook&#39;s notifications. | [optional] 
**Name** | Pointer to **string** | Name of the program gateway funding source. | [optional] 
**TimeoutMillis** | Pointer to **int64** | Total timeout in milliseconds for gateway processing. | [optional] 
**Url** | **string** | URL of the gateway endpoint hosted in your environment, to which &#x60;POST&#x60; requests are submitted by the Marqeta platform. | 
**UseMtls** | Pointer to **bool** | Specifies whether or not to use mutual transport layer security (mTLS) authentication for the funding request. | [optional] [default to false]

## Methods

### NewGatewayProgramFundingSourceUpdateRequest

`func NewGatewayProgramFundingSourceUpdateRequest(basicAuthPassword string, basicAuthUsername string, url string, ) *GatewayProgramFundingSourceUpdateRequest`

NewGatewayProgramFundingSourceUpdateRequest instantiates a new GatewayProgramFundingSourceUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayProgramFundingSourceUpdateRequestWithDefaults

`func NewGatewayProgramFundingSourceUpdateRequestWithDefaults() *GatewayProgramFundingSourceUpdateRequest`

NewGatewayProgramFundingSourceUpdateRequestWithDefaults instantiates a new GatewayProgramFundingSourceUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *GatewayProgramFundingSourceUpdateRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *GatewayProgramFundingSourceUpdateRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *GatewayProgramFundingSourceUpdateRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *GatewayProgramFundingSourceUpdateRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetBasicAuthPassword

`func (o *GatewayProgramFundingSourceUpdateRequest) GetBasicAuthPassword() string`

GetBasicAuthPassword returns the BasicAuthPassword field if non-nil, zero value otherwise.

### GetBasicAuthPasswordOk

`func (o *GatewayProgramFundingSourceUpdateRequest) GetBasicAuthPasswordOk() (*string, bool)`

GetBasicAuthPasswordOk returns a tuple with the BasicAuthPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuthPassword

`func (o *GatewayProgramFundingSourceUpdateRequest) SetBasicAuthPassword(v string)`

SetBasicAuthPassword sets BasicAuthPassword field to given value.


### GetBasicAuthUsername

`func (o *GatewayProgramFundingSourceUpdateRequest) GetBasicAuthUsername() string`

GetBasicAuthUsername returns the BasicAuthUsername field if non-nil, zero value otherwise.

### GetBasicAuthUsernameOk

`func (o *GatewayProgramFundingSourceUpdateRequest) GetBasicAuthUsernameOk() (*string, bool)`

GetBasicAuthUsernameOk returns a tuple with the BasicAuthUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuthUsername

`func (o *GatewayProgramFundingSourceUpdateRequest) SetBasicAuthUsername(v string)`

SetBasicAuthUsername sets BasicAuthUsername field to given value.


### GetCustomHeader

`func (o *GatewayProgramFundingSourceUpdateRequest) GetCustomHeader() map[string]string`

GetCustomHeader returns the CustomHeader field if non-nil, zero value otherwise.

### GetCustomHeaderOk

`func (o *GatewayProgramFundingSourceUpdateRequest) GetCustomHeaderOk() (*map[string]string, bool)`

GetCustomHeaderOk returns a tuple with the CustomHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomHeader

`func (o *GatewayProgramFundingSourceUpdateRequest) SetCustomHeader(v map[string]string)`

SetCustomHeader sets CustomHeader field to given value.

### HasCustomHeader

`func (o *GatewayProgramFundingSourceUpdateRequest) HasCustomHeader() bool`

HasCustomHeader returns a boolean if a field has been set.

### GetName

`func (o *GatewayProgramFundingSourceUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayProgramFundingSourceUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayProgramFundingSourceUpdateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GatewayProgramFundingSourceUpdateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTimeoutMillis

`func (o *GatewayProgramFundingSourceUpdateRequest) GetTimeoutMillis() int64`

GetTimeoutMillis returns the TimeoutMillis field if non-nil, zero value otherwise.

### GetTimeoutMillisOk

`func (o *GatewayProgramFundingSourceUpdateRequest) GetTimeoutMillisOk() (*int64, bool)`

GetTimeoutMillisOk returns a tuple with the TimeoutMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutMillis

`func (o *GatewayProgramFundingSourceUpdateRequest) SetTimeoutMillis(v int64)`

SetTimeoutMillis sets TimeoutMillis field to given value.

### HasTimeoutMillis

`func (o *GatewayProgramFundingSourceUpdateRequest) HasTimeoutMillis() bool`

HasTimeoutMillis returns a boolean if a field has been set.

### GetUrl

`func (o *GatewayProgramFundingSourceUpdateRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GatewayProgramFundingSourceUpdateRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GatewayProgramFundingSourceUpdateRequest) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetUseMtls

`func (o *GatewayProgramFundingSourceUpdateRequest) GetUseMtls() bool`

GetUseMtls returns the UseMtls field if non-nil, zero value otherwise.

### GetUseMtlsOk

`func (o *GatewayProgramFundingSourceUpdateRequest) GetUseMtlsOk() (*bool, bool)`

GetUseMtlsOk returns a tuple with the UseMtls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseMtls

`func (o *GatewayProgramFundingSourceUpdateRequest) SetUseMtls(v bool)`

SetUseMtls sets UseMtls field to given value.

### HasUseMtls

`func (o *GatewayProgramFundingSourceUpdateRequest) HasUseMtls() bool`

HasUseMtls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


