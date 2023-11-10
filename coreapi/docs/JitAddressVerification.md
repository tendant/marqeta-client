# JitAddressVerification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gateway** | Pointer to [**AddressVerificationSource**](AddressVerificationSource.md) |  | [optional] 
**Issuer** | Pointer to [**AddressVerificationSource**](AddressVerificationSource.md) |  | [optional] 
**Request** | Pointer to [**AvsInformation**](AvsInformation.md) |  | [optional] 

## Methods

### NewJitAddressVerification

`func NewJitAddressVerification() *JitAddressVerification`

NewJitAddressVerification instantiates a new JitAddressVerification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJitAddressVerificationWithDefaults

`func NewJitAddressVerificationWithDefaults() *JitAddressVerification`

NewJitAddressVerificationWithDefaults instantiates a new JitAddressVerification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGateway

`func (o *JitAddressVerification) GetGateway() AddressVerificationSource`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *JitAddressVerification) GetGatewayOk() (*AddressVerificationSource, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *JitAddressVerification) SetGateway(v AddressVerificationSource)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *JitAddressVerification) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetIssuer

`func (o *JitAddressVerification) GetIssuer() AddressVerificationSource`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *JitAddressVerification) GetIssuerOk() (*AddressVerificationSource, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *JitAddressVerification) SetIssuer(v AddressVerificationSource)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *JitAddressVerification) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetRequest

`func (o *JitAddressVerification) GetRequest() AvsInformation`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *JitAddressVerification) GetRequestOk() (*AvsInformation, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *JitAddressVerification) SetRequest(v AvsInformation)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *JitAddressVerification) HasRequest() bool`

HasRequest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


