# GatewayProgramFundingSourceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | **string** | Bank account number. | 
**Active** | Pointer to **bool** | Indicates whether the program gateway funding source is active. This field is returned if it exists in the resource. | [optional] 
**BasicAuthPassword** | **string** | Password for authenticating your environment. | 
**BasicAuthUsername** | **string** | Username for authenticating your environment. | 
**CreatedTime** | **time.Time** | Date and time when the resource was created, in UTC. | 
**CustomHeader** | **map[string]string** | Additional custom information included in the HTTP header. | 
**LastModifiedTime** | **time.Time** | Date and time when the resource was last modified, in UTC. | 
**Name** | **string** | Name of the program gateway funding source. | 
**TimeoutMillis** | **int64** | Total timeout in milliseconds for gateway processing. | 
**Token** | **string** | Unique identifier of the program gateway funding source. | 
**Url** | **string** | URL of the gateway endpoint hosted in your environment, to which &#x60;POST&#x60; requests are submitted by the Marqeta platform. | 
**UseMtls** | **bool** | Specifies whether or not to use mutual transport layer security (mTLS) authentication for the funding request. | [default to false]
**Version** | **string** | Program gateway funding source object version. | 

## Methods

### NewGatewayProgramFundingSourceResponse

`func NewGatewayProgramFundingSourceResponse(account string, basicAuthPassword string, basicAuthUsername string, createdTime time.Time, customHeader map[string]string, lastModifiedTime time.Time, name string, timeoutMillis int64, token string, url string, useMtls bool, version string, ) *GatewayProgramFundingSourceResponse`

NewGatewayProgramFundingSourceResponse instantiates a new GatewayProgramFundingSourceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayProgramFundingSourceResponseWithDefaults

`func NewGatewayProgramFundingSourceResponseWithDefaults() *GatewayProgramFundingSourceResponse`

NewGatewayProgramFundingSourceResponseWithDefaults instantiates a new GatewayProgramFundingSourceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *GatewayProgramFundingSourceResponse) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *GatewayProgramFundingSourceResponse) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *GatewayProgramFundingSourceResponse) SetAccount(v string)`

SetAccount sets Account field to given value.


### GetActive

`func (o *GatewayProgramFundingSourceResponse) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *GatewayProgramFundingSourceResponse) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *GatewayProgramFundingSourceResponse) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *GatewayProgramFundingSourceResponse) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetBasicAuthPassword

`func (o *GatewayProgramFundingSourceResponse) GetBasicAuthPassword() string`

GetBasicAuthPassword returns the BasicAuthPassword field if non-nil, zero value otherwise.

### GetBasicAuthPasswordOk

`func (o *GatewayProgramFundingSourceResponse) GetBasicAuthPasswordOk() (*string, bool)`

GetBasicAuthPasswordOk returns a tuple with the BasicAuthPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuthPassword

`func (o *GatewayProgramFundingSourceResponse) SetBasicAuthPassword(v string)`

SetBasicAuthPassword sets BasicAuthPassword field to given value.


### GetBasicAuthUsername

`func (o *GatewayProgramFundingSourceResponse) GetBasicAuthUsername() string`

GetBasicAuthUsername returns the BasicAuthUsername field if non-nil, zero value otherwise.

### GetBasicAuthUsernameOk

`func (o *GatewayProgramFundingSourceResponse) GetBasicAuthUsernameOk() (*string, bool)`

GetBasicAuthUsernameOk returns a tuple with the BasicAuthUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuthUsername

`func (o *GatewayProgramFundingSourceResponse) SetBasicAuthUsername(v string)`

SetBasicAuthUsername sets BasicAuthUsername field to given value.


### GetCreatedTime

`func (o *GatewayProgramFundingSourceResponse) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *GatewayProgramFundingSourceResponse) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *GatewayProgramFundingSourceResponse) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetCustomHeader

`func (o *GatewayProgramFundingSourceResponse) GetCustomHeader() map[string]string`

GetCustomHeader returns the CustomHeader field if non-nil, zero value otherwise.

### GetCustomHeaderOk

`func (o *GatewayProgramFundingSourceResponse) GetCustomHeaderOk() (*map[string]string, bool)`

GetCustomHeaderOk returns a tuple with the CustomHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomHeader

`func (o *GatewayProgramFundingSourceResponse) SetCustomHeader(v map[string]string)`

SetCustomHeader sets CustomHeader field to given value.


### GetLastModifiedTime

`func (o *GatewayProgramFundingSourceResponse) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *GatewayProgramFundingSourceResponse) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *GatewayProgramFundingSourceResponse) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.


### GetName

`func (o *GatewayProgramFundingSourceResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayProgramFundingSourceResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayProgramFundingSourceResponse) SetName(v string)`

SetName sets Name field to given value.


### GetTimeoutMillis

`func (o *GatewayProgramFundingSourceResponse) GetTimeoutMillis() int64`

GetTimeoutMillis returns the TimeoutMillis field if non-nil, zero value otherwise.

### GetTimeoutMillisOk

`func (o *GatewayProgramFundingSourceResponse) GetTimeoutMillisOk() (*int64, bool)`

GetTimeoutMillisOk returns a tuple with the TimeoutMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutMillis

`func (o *GatewayProgramFundingSourceResponse) SetTimeoutMillis(v int64)`

SetTimeoutMillis sets TimeoutMillis field to given value.


### GetToken

`func (o *GatewayProgramFundingSourceResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GatewayProgramFundingSourceResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GatewayProgramFundingSourceResponse) SetToken(v string)`

SetToken sets Token field to given value.


### GetUrl

`func (o *GatewayProgramFundingSourceResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GatewayProgramFundingSourceResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GatewayProgramFundingSourceResponse) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetUseMtls

`func (o *GatewayProgramFundingSourceResponse) GetUseMtls() bool`

GetUseMtls returns the UseMtls field if non-nil, zero value otherwise.

### GetUseMtlsOk

`func (o *GatewayProgramFundingSourceResponse) GetUseMtlsOk() (*bool, bool)`

GetUseMtlsOk returns a tuple with the UseMtls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseMtls

`func (o *GatewayProgramFundingSourceResponse) SetUseMtls(v bool)`

SetUseMtls sets UseMtls field to given value.


### GetVersion

`func (o *GatewayProgramFundingSourceResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GatewayProgramFundingSourceResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GatewayProgramFundingSourceResponse) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


