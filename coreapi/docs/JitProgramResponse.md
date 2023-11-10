# JitProgramResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JitFunding** | [**JitFundingApi**](JitFundingApi.md) |  | 
**NetworkMetadata** | Pointer to [**NetworkMetadata**](NetworkMetadata.md) |  | [optional] 

## Methods

### NewJitProgramResponse

`func NewJitProgramResponse(jitFunding JitFundingApi, ) *JitProgramResponse`

NewJitProgramResponse instantiates a new JitProgramResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJitProgramResponseWithDefaults

`func NewJitProgramResponseWithDefaults() *JitProgramResponse`

NewJitProgramResponseWithDefaults instantiates a new JitProgramResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJitFunding

`func (o *JitProgramResponse) GetJitFunding() JitFundingApi`

GetJitFunding returns the JitFunding field if non-nil, zero value otherwise.

### GetJitFundingOk

`func (o *JitProgramResponse) GetJitFundingOk() (*JitFundingApi, bool)`

GetJitFundingOk returns a tuple with the JitFunding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJitFunding

`func (o *JitProgramResponse) SetJitFunding(v JitFundingApi)`

SetJitFunding sets JitFunding field to given value.


### GetNetworkMetadata

`func (o *JitProgramResponse) GetNetworkMetadata() NetworkMetadata`

GetNetworkMetadata returns the NetworkMetadata field if non-nil, zero value otherwise.

### GetNetworkMetadataOk

`func (o *JitProgramResponse) GetNetworkMetadataOk() (*NetworkMetadata, bool)`

GetNetworkMetadataOk returns a tuple with the NetworkMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkMetadata

`func (o *JitProgramResponse) SetNetworkMetadata(v NetworkMetadata)`

SetNetworkMetadata sets NetworkMetadata field to given value.

### HasNetworkMetadata

`func (o *JitProgramResponse) HasNetworkMetadata() bool`

HasNetworkMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


