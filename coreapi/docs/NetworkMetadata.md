# NetworkMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountIdentification1** | Pointer to **string** |  | [optional] 
**IncomingResponseCode** | Pointer to **string** | Visa Risk Management esponse code &#x60;59&#x60;, indicating suspected fraud. | [optional] 
**ProductId** | Pointer to **string** | Product identification value assigned by the card network to each card product. Can be used to track card-level activity by individual account number for premium card products. | [optional] 
**ProgramId** | Pointer to **string** | Program identification number used with &#x60;product_id&#x60; that identifies the programs associated with a card within a program registered by the issuer with the card network. | [optional] 
**SpendQualifier** | Pointer to **string** | Indicates whether or not the base spend-assessment threshold defined by the card network has been met. | [optional] 
**SurchargeFreeAtmNetwork** | Pointer to **string** | Name of the surcharge-free ATM network used to complete the transaction. | [optional] 

## Methods

### NewNetworkMetadata

`func NewNetworkMetadata() *NetworkMetadata`

NewNetworkMetadata instantiates a new NetworkMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkMetadataWithDefaults

`func NewNetworkMetadataWithDefaults() *NetworkMetadata`

NewNetworkMetadataWithDefaults instantiates a new NetworkMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountIdentification1

`func (o *NetworkMetadata) GetAccountIdentification1() string`

GetAccountIdentification1 returns the AccountIdentification1 field if non-nil, zero value otherwise.

### GetAccountIdentification1Ok

`func (o *NetworkMetadata) GetAccountIdentification1Ok() (*string, bool)`

GetAccountIdentification1Ok returns a tuple with the AccountIdentification1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountIdentification1

`func (o *NetworkMetadata) SetAccountIdentification1(v string)`

SetAccountIdentification1 sets AccountIdentification1 field to given value.

### HasAccountIdentification1

`func (o *NetworkMetadata) HasAccountIdentification1() bool`

HasAccountIdentification1 returns a boolean if a field has been set.

### GetIncomingResponseCode

`func (o *NetworkMetadata) GetIncomingResponseCode() string`

GetIncomingResponseCode returns the IncomingResponseCode field if non-nil, zero value otherwise.

### GetIncomingResponseCodeOk

`func (o *NetworkMetadata) GetIncomingResponseCodeOk() (*string, bool)`

GetIncomingResponseCodeOk returns a tuple with the IncomingResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomingResponseCode

`func (o *NetworkMetadata) SetIncomingResponseCode(v string)`

SetIncomingResponseCode sets IncomingResponseCode field to given value.

### HasIncomingResponseCode

`func (o *NetworkMetadata) HasIncomingResponseCode() bool`

HasIncomingResponseCode returns a boolean if a field has been set.

### GetProductId

`func (o *NetworkMetadata) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *NetworkMetadata) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *NetworkMetadata) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *NetworkMetadata) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetProgramId

`func (o *NetworkMetadata) GetProgramId() string`

GetProgramId returns the ProgramId field if non-nil, zero value otherwise.

### GetProgramIdOk

`func (o *NetworkMetadata) GetProgramIdOk() (*string, bool)`

GetProgramIdOk returns a tuple with the ProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramId

`func (o *NetworkMetadata) SetProgramId(v string)`

SetProgramId sets ProgramId field to given value.

### HasProgramId

`func (o *NetworkMetadata) HasProgramId() bool`

HasProgramId returns a boolean if a field has been set.

### GetSpendQualifier

`func (o *NetworkMetadata) GetSpendQualifier() string`

GetSpendQualifier returns the SpendQualifier field if non-nil, zero value otherwise.

### GetSpendQualifierOk

`func (o *NetworkMetadata) GetSpendQualifierOk() (*string, bool)`

GetSpendQualifierOk returns a tuple with the SpendQualifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendQualifier

`func (o *NetworkMetadata) SetSpendQualifier(v string)`

SetSpendQualifier sets SpendQualifier field to given value.

### HasSpendQualifier

`func (o *NetworkMetadata) HasSpendQualifier() bool`

HasSpendQualifier returns a boolean if a field has been set.

### GetSurchargeFreeAtmNetwork

`func (o *NetworkMetadata) GetSurchargeFreeAtmNetwork() string`

GetSurchargeFreeAtmNetwork returns the SurchargeFreeAtmNetwork field if non-nil, zero value otherwise.

### GetSurchargeFreeAtmNetworkOk

`func (o *NetworkMetadata) GetSurchargeFreeAtmNetworkOk() (*string, bool)`

GetSurchargeFreeAtmNetworkOk returns a tuple with the SurchargeFreeAtmNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurchargeFreeAtmNetwork

`func (o *NetworkMetadata) SetSurchargeFreeAtmNetwork(v string)`

SetSurchargeFreeAtmNetwork sets SurchargeFreeAtmNetwork field to given value.

### HasSurchargeFreeAtmNetwork

`func (o *NetworkMetadata) HasSurchargeFreeAtmNetwork() bool`

HasSurchargeFreeAtmNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


