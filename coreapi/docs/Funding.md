# Funding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **float32** | Amount of funds loaded into the GPA. | [optional] 
**GatewayLog** | Pointer to [**GatewayLogModel**](GatewayLogModel.md) |  | [optional] 
**Source** | [**FundingSourceModel**](FundingSourceModel.md) |  | 
**SourceAddress** | Pointer to [**CardholderAddressResponse**](CardholderAddressResponse.md) |  | [optional] 

## Methods

### NewFunding

`func NewFunding(source FundingSourceModel, ) *Funding`

NewFunding instantiates a new Funding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFundingWithDefaults

`func NewFundingWithDefaults() *Funding`

NewFundingWithDefaults instantiates a new Funding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *Funding) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Funding) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Funding) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *Funding) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetGatewayLog

`func (o *Funding) GetGatewayLog() GatewayLogModel`

GetGatewayLog returns the GatewayLog field if non-nil, zero value otherwise.

### GetGatewayLogOk

`func (o *Funding) GetGatewayLogOk() (*GatewayLogModel, bool)`

GetGatewayLogOk returns a tuple with the GatewayLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayLog

`func (o *Funding) SetGatewayLog(v GatewayLogModel)`

SetGatewayLog sets GatewayLog field to given value.

### HasGatewayLog

`func (o *Funding) HasGatewayLog() bool`

HasGatewayLog returns a boolean if a field has been set.

### GetSource

`func (o *Funding) GetSource() FundingSourceModel`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Funding) GetSourceOk() (*FundingSourceModel, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Funding) SetSource(v FundingSourceModel)`

SetSource sets Source field to given value.


### GetSourceAddress

`func (o *Funding) GetSourceAddress() CardholderAddressResponse`

GetSourceAddress returns the SourceAddress field if non-nil, zero value otherwise.

### GetSourceAddressOk

`func (o *Funding) GetSourceAddressOk() (*CardholderAddressResponse, bool)`

GetSourceAddressOk returns a tuple with the SourceAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAddress

`func (o *Funding) SetSourceAddress(v CardholderAddressResponse)`

SetSourceAddress sets SourceAddress field to given value.

### HasSourceAddress

`func (o *Funding) HasSourceAddress() bool`

HasSourceAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


