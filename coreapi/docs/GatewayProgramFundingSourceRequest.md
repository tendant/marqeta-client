# GatewayProgramFundingSourceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Indicates whether the program gateway funding source is active. | [optional] 
**BasicAuthPassword** | **string** | Password for authenticating your environment. | 
**BasicAuthUsername** | **string** | Username for authenticating your environment. | 
**CustomHeader** | Pointer to **map[string]string** | Additional custom information included in the HTTP header. For example, this might contain security information, along with Basic Authentication, when making a JIT Funding request. Custom headers also appear in the associated webhook&#39;s notifications. | [optional] 
**Name** | **string** | Name of the program gateway funding source. | 
**TimeoutMillis** | Pointer to **int64** | Total timeout in milliseconds for gateway processing. | [optional] 
**Token** | Pointer to **string** | Unique identifier of the program gateway funding source. If you do not include a token, the system will generate one automatically. As this token is necessary for use in other calls, we recommend that you define a simple and easy to remember string rather than letting the system generate a token for you. This value cannot be updated. | [optional] 
**Url** | **string** | URL of the gateway endpoint hosted in your environment, to which &#x60;POST&#x60; requests are submitted by the Marqeta platform. | 
**UseMtls** | Pointer to **bool** | Specifies whether or not to use mutual transport layer security (mTLS) authentication for the funding request. | [optional] [default to false]

## Methods

### NewGatewayProgramFundingSourceRequest

`func NewGatewayProgramFundingSourceRequest(basicAuthPassword string, basicAuthUsername string, name string, url string, ) *GatewayProgramFundingSourceRequest`

NewGatewayProgramFundingSourceRequest instantiates a new GatewayProgramFundingSourceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayProgramFundingSourceRequestWithDefaults

`func NewGatewayProgramFundingSourceRequestWithDefaults() *GatewayProgramFundingSourceRequest`

NewGatewayProgramFundingSourceRequestWithDefaults instantiates a new GatewayProgramFundingSourceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *GatewayProgramFundingSourceRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *GatewayProgramFundingSourceRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *GatewayProgramFundingSourceRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *GatewayProgramFundingSourceRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetBasicAuthPassword

`func (o *GatewayProgramFundingSourceRequest) GetBasicAuthPassword() string`

GetBasicAuthPassword returns the BasicAuthPassword field if non-nil, zero value otherwise.

### GetBasicAuthPasswordOk

`func (o *GatewayProgramFundingSourceRequest) GetBasicAuthPasswordOk() (*string, bool)`

GetBasicAuthPasswordOk returns a tuple with the BasicAuthPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuthPassword

`func (o *GatewayProgramFundingSourceRequest) SetBasicAuthPassword(v string)`

SetBasicAuthPassword sets BasicAuthPassword field to given value.


### GetBasicAuthUsername

`func (o *GatewayProgramFundingSourceRequest) GetBasicAuthUsername() string`

GetBasicAuthUsername returns the BasicAuthUsername field if non-nil, zero value otherwise.

### GetBasicAuthUsernameOk

`func (o *GatewayProgramFundingSourceRequest) GetBasicAuthUsernameOk() (*string, bool)`

GetBasicAuthUsernameOk returns a tuple with the BasicAuthUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuthUsername

`func (o *GatewayProgramFundingSourceRequest) SetBasicAuthUsername(v string)`

SetBasicAuthUsername sets BasicAuthUsername field to given value.


### GetCustomHeader

`func (o *GatewayProgramFundingSourceRequest) GetCustomHeader() map[string]string`

GetCustomHeader returns the CustomHeader field if non-nil, zero value otherwise.

### GetCustomHeaderOk

`func (o *GatewayProgramFundingSourceRequest) GetCustomHeaderOk() (*map[string]string, bool)`

GetCustomHeaderOk returns a tuple with the CustomHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomHeader

`func (o *GatewayProgramFundingSourceRequest) SetCustomHeader(v map[string]string)`

SetCustomHeader sets CustomHeader field to given value.

### HasCustomHeader

`func (o *GatewayProgramFundingSourceRequest) HasCustomHeader() bool`

HasCustomHeader returns a boolean if a field has been set.

### GetName

`func (o *GatewayProgramFundingSourceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayProgramFundingSourceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayProgramFundingSourceRequest) SetName(v string)`

SetName sets Name field to given value.


### GetTimeoutMillis

`func (o *GatewayProgramFundingSourceRequest) GetTimeoutMillis() int64`

GetTimeoutMillis returns the TimeoutMillis field if non-nil, zero value otherwise.

### GetTimeoutMillisOk

`func (o *GatewayProgramFundingSourceRequest) GetTimeoutMillisOk() (*int64, bool)`

GetTimeoutMillisOk returns a tuple with the TimeoutMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutMillis

`func (o *GatewayProgramFundingSourceRequest) SetTimeoutMillis(v int64)`

SetTimeoutMillis sets TimeoutMillis field to given value.

### HasTimeoutMillis

`func (o *GatewayProgramFundingSourceRequest) HasTimeoutMillis() bool`

HasTimeoutMillis returns a boolean if a field has been set.

### GetToken

`func (o *GatewayProgramFundingSourceRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GatewayProgramFundingSourceRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GatewayProgramFundingSourceRequest) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GatewayProgramFundingSourceRequest) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUrl

`func (o *GatewayProgramFundingSourceRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GatewayProgramFundingSourceRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GatewayProgramFundingSourceRequest) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetUseMtls

`func (o *GatewayProgramFundingSourceRequest) GetUseMtls() bool`

GetUseMtls returns the UseMtls field if non-nil, zero value otherwise.

### GetUseMtlsOk

`func (o *GatewayProgramFundingSourceRequest) GetUseMtlsOk() (*bool, bool)`

GetUseMtlsOk returns a tuple with the UseMtls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseMtls

`func (o *GatewayProgramFundingSourceRequest) SetUseMtls(v bool)`

SetUseMtls sets UseMtls field to given value.

### HasUseMtls

`func (o *GatewayProgramFundingSourceRequest) HasUseMtls() bool`

HasUseMtls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


