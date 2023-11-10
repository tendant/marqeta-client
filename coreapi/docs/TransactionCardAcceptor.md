# TransactionCardAcceptor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | Card acceptor&#39;s address. May be returned if the request uses Transaction Model v1 of the Marqeta Core API. Not returned for Transaction Model v2 requests. | [optional] 
**City** | Pointer to **string** | Card acceptor&#39;s city. | [optional] 
**CountryCode** | Pointer to **string** | Card acceptor&#39;s country code. May be returned if the request uses Transaction Model v2 of the Marqeta Core API. Not returned for Transaction Model v1 requests. | [optional] 
**CountryOfOrigin** | Pointer to **string** | The merchant&#39;s country of origin.  A merchant&#39;s country of origin can be different from the country in which the merchant is located. For example, embassies are physically located in countries that are not their country of origin: a Mexican embassy might be physically located in Singapore, but the country of origin is Mexico.  This field is included when the cardholder conducts a transaction with a government-controlled merchant using a Marqeta-issued card. | [optional] 
**CustomerServicePhone** | Pointer to **string** |  | [optional] 
**IndependentSalesOrganizationId** | Pointer to **string** |  | [optional] 
**Mcc** | Pointer to **string** | Merchant category code (MCC). | [optional] 
**MccGroups** | Pointer to **[]string** | An array of &#x60;mcc_groups&#x60;. | [optional] 
**MerchantTaxId** | Pointer to **string** |  | [optional] 
**Mid** | Pointer to **string** | The merchant identifier. | [optional] 
**Name** | Pointer to **string** | Card acceptor&#39;s name. | [optional] 
**NetworkAssignedId** | Pointer to **string** |  | [optional] 
**NetworkMid** | Pointer to **string** | The merchant identifier on the card network. | [optional] 
**PaymentFacilitatorId** | Pointer to **string** |  | [optional] 
**PaymentFacilitatorName** | Pointer to **string** | The name of the payment facilitator, including the sub-merchant identifier, of an original credit transaction. This field is returned when a payment facilitator participates in the transaction. | [optional] 
**Phone** | Pointer to **string** |  | [optional] 
**Poi** | Pointer to [**TerminalModel**](TerminalModel.md) |  | [optional] 
**PostalCode** | Pointer to **string** | Card acceptor&#39;s postal code. | [optional] 
**SpecialMerchantId** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** | Two-character state, province, or territorial abbreviation.  For a complete list of valid state and province abbreviations, see &lt;&lt;/core-api/kyc-verification#_valid_state_provincial_and_territorial_abbreviations, Valid state, provincial, and territorial abbreviations&gt;&gt;.  *Note*: Non-US merchants may return more than 2 char for this field. | [optional] 
**SubMerchantId** | Pointer to **string** |  | [optional] 
**TransferServiceProviderName** | Pointer to **string** | The name of the transfer service provider of a money transfer original credit transaction. This field is included when a transfer service provider participates in the transaction. | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewTransactionCardAcceptor

`func NewTransactionCardAcceptor() *TransactionCardAcceptor`

NewTransactionCardAcceptor instantiates a new TransactionCardAcceptor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionCardAcceptorWithDefaults

`func NewTransactionCardAcceptorWithDefaults() *TransactionCardAcceptor`

NewTransactionCardAcceptorWithDefaults instantiates a new TransactionCardAcceptor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *TransactionCardAcceptor) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *TransactionCardAcceptor) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *TransactionCardAcceptor) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *TransactionCardAcceptor) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetCity

`func (o *TransactionCardAcceptor) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *TransactionCardAcceptor) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *TransactionCardAcceptor) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *TransactionCardAcceptor) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCountryCode

`func (o *TransactionCardAcceptor) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *TransactionCardAcceptor) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *TransactionCardAcceptor) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *TransactionCardAcceptor) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetCountryOfOrigin

`func (o *TransactionCardAcceptor) GetCountryOfOrigin() string`

GetCountryOfOrigin returns the CountryOfOrigin field if non-nil, zero value otherwise.

### GetCountryOfOriginOk

`func (o *TransactionCardAcceptor) GetCountryOfOriginOk() (*string, bool)`

GetCountryOfOriginOk returns a tuple with the CountryOfOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryOfOrigin

`func (o *TransactionCardAcceptor) SetCountryOfOrigin(v string)`

SetCountryOfOrigin sets CountryOfOrigin field to given value.

### HasCountryOfOrigin

`func (o *TransactionCardAcceptor) HasCountryOfOrigin() bool`

HasCountryOfOrigin returns a boolean if a field has been set.

### GetCustomerServicePhone

`func (o *TransactionCardAcceptor) GetCustomerServicePhone() string`

GetCustomerServicePhone returns the CustomerServicePhone field if non-nil, zero value otherwise.

### GetCustomerServicePhoneOk

`func (o *TransactionCardAcceptor) GetCustomerServicePhoneOk() (*string, bool)`

GetCustomerServicePhoneOk returns a tuple with the CustomerServicePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerServicePhone

`func (o *TransactionCardAcceptor) SetCustomerServicePhone(v string)`

SetCustomerServicePhone sets CustomerServicePhone field to given value.

### HasCustomerServicePhone

`func (o *TransactionCardAcceptor) HasCustomerServicePhone() bool`

HasCustomerServicePhone returns a boolean if a field has been set.

### GetIndependentSalesOrganizationId

`func (o *TransactionCardAcceptor) GetIndependentSalesOrganizationId() string`

GetIndependentSalesOrganizationId returns the IndependentSalesOrganizationId field if non-nil, zero value otherwise.

### GetIndependentSalesOrganizationIdOk

`func (o *TransactionCardAcceptor) GetIndependentSalesOrganizationIdOk() (*string, bool)`

GetIndependentSalesOrganizationIdOk returns a tuple with the IndependentSalesOrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndependentSalesOrganizationId

`func (o *TransactionCardAcceptor) SetIndependentSalesOrganizationId(v string)`

SetIndependentSalesOrganizationId sets IndependentSalesOrganizationId field to given value.

### HasIndependentSalesOrganizationId

`func (o *TransactionCardAcceptor) HasIndependentSalesOrganizationId() bool`

HasIndependentSalesOrganizationId returns a boolean if a field has been set.

### GetMcc

`func (o *TransactionCardAcceptor) GetMcc() string`

GetMcc returns the Mcc field if non-nil, zero value otherwise.

### GetMccOk

`func (o *TransactionCardAcceptor) GetMccOk() (*string, bool)`

GetMccOk returns a tuple with the Mcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcc

`func (o *TransactionCardAcceptor) SetMcc(v string)`

SetMcc sets Mcc field to given value.

### HasMcc

`func (o *TransactionCardAcceptor) HasMcc() bool`

HasMcc returns a boolean if a field has been set.

### GetMccGroups

`func (o *TransactionCardAcceptor) GetMccGroups() []string`

GetMccGroups returns the MccGroups field if non-nil, zero value otherwise.

### GetMccGroupsOk

`func (o *TransactionCardAcceptor) GetMccGroupsOk() (*[]string, bool)`

GetMccGroupsOk returns a tuple with the MccGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMccGroups

`func (o *TransactionCardAcceptor) SetMccGroups(v []string)`

SetMccGroups sets MccGroups field to given value.

### HasMccGroups

`func (o *TransactionCardAcceptor) HasMccGroups() bool`

HasMccGroups returns a boolean if a field has been set.

### GetMerchantTaxId

`func (o *TransactionCardAcceptor) GetMerchantTaxId() string`

GetMerchantTaxId returns the MerchantTaxId field if non-nil, zero value otherwise.

### GetMerchantTaxIdOk

`func (o *TransactionCardAcceptor) GetMerchantTaxIdOk() (*string, bool)`

GetMerchantTaxIdOk returns a tuple with the MerchantTaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantTaxId

`func (o *TransactionCardAcceptor) SetMerchantTaxId(v string)`

SetMerchantTaxId sets MerchantTaxId field to given value.

### HasMerchantTaxId

`func (o *TransactionCardAcceptor) HasMerchantTaxId() bool`

HasMerchantTaxId returns a boolean if a field has been set.

### GetMid

`func (o *TransactionCardAcceptor) GetMid() string`

GetMid returns the Mid field if non-nil, zero value otherwise.

### GetMidOk

`func (o *TransactionCardAcceptor) GetMidOk() (*string, bool)`

GetMidOk returns a tuple with the Mid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMid

`func (o *TransactionCardAcceptor) SetMid(v string)`

SetMid sets Mid field to given value.

### HasMid

`func (o *TransactionCardAcceptor) HasMid() bool`

HasMid returns a boolean if a field has been set.

### GetName

`func (o *TransactionCardAcceptor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TransactionCardAcceptor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TransactionCardAcceptor) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TransactionCardAcceptor) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworkAssignedId

`func (o *TransactionCardAcceptor) GetNetworkAssignedId() string`

GetNetworkAssignedId returns the NetworkAssignedId field if non-nil, zero value otherwise.

### GetNetworkAssignedIdOk

`func (o *TransactionCardAcceptor) GetNetworkAssignedIdOk() (*string, bool)`

GetNetworkAssignedIdOk returns a tuple with the NetworkAssignedId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkAssignedId

`func (o *TransactionCardAcceptor) SetNetworkAssignedId(v string)`

SetNetworkAssignedId sets NetworkAssignedId field to given value.

### HasNetworkAssignedId

`func (o *TransactionCardAcceptor) HasNetworkAssignedId() bool`

HasNetworkAssignedId returns a boolean if a field has been set.

### GetNetworkMid

`func (o *TransactionCardAcceptor) GetNetworkMid() string`

GetNetworkMid returns the NetworkMid field if non-nil, zero value otherwise.

### GetNetworkMidOk

`func (o *TransactionCardAcceptor) GetNetworkMidOk() (*string, bool)`

GetNetworkMidOk returns a tuple with the NetworkMid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkMid

`func (o *TransactionCardAcceptor) SetNetworkMid(v string)`

SetNetworkMid sets NetworkMid field to given value.

### HasNetworkMid

`func (o *TransactionCardAcceptor) HasNetworkMid() bool`

HasNetworkMid returns a boolean if a field has been set.

### GetPaymentFacilitatorId

`func (o *TransactionCardAcceptor) GetPaymentFacilitatorId() string`

GetPaymentFacilitatorId returns the PaymentFacilitatorId field if non-nil, zero value otherwise.

### GetPaymentFacilitatorIdOk

`func (o *TransactionCardAcceptor) GetPaymentFacilitatorIdOk() (*string, bool)`

GetPaymentFacilitatorIdOk returns a tuple with the PaymentFacilitatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentFacilitatorId

`func (o *TransactionCardAcceptor) SetPaymentFacilitatorId(v string)`

SetPaymentFacilitatorId sets PaymentFacilitatorId field to given value.

### HasPaymentFacilitatorId

`func (o *TransactionCardAcceptor) HasPaymentFacilitatorId() bool`

HasPaymentFacilitatorId returns a boolean if a field has been set.

### GetPaymentFacilitatorName

`func (o *TransactionCardAcceptor) GetPaymentFacilitatorName() string`

GetPaymentFacilitatorName returns the PaymentFacilitatorName field if non-nil, zero value otherwise.

### GetPaymentFacilitatorNameOk

`func (o *TransactionCardAcceptor) GetPaymentFacilitatorNameOk() (*string, bool)`

GetPaymentFacilitatorNameOk returns a tuple with the PaymentFacilitatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentFacilitatorName

`func (o *TransactionCardAcceptor) SetPaymentFacilitatorName(v string)`

SetPaymentFacilitatorName sets PaymentFacilitatorName field to given value.

### HasPaymentFacilitatorName

`func (o *TransactionCardAcceptor) HasPaymentFacilitatorName() bool`

HasPaymentFacilitatorName returns a boolean if a field has been set.

### GetPhone

`func (o *TransactionCardAcceptor) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *TransactionCardAcceptor) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *TransactionCardAcceptor) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *TransactionCardAcceptor) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetPoi

`func (o *TransactionCardAcceptor) GetPoi() TerminalModel`

GetPoi returns the Poi field if non-nil, zero value otherwise.

### GetPoiOk

`func (o *TransactionCardAcceptor) GetPoiOk() (*TerminalModel, bool)`

GetPoiOk returns a tuple with the Poi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoi

`func (o *TransactionCardAcceptor) SetPoi(v TerminalModel)`

SetPoi sets Poi field to given value.

### HasPoi

`func (o *TransactionCardAcceptor) HasPoi() bool`

HasPoi returns a boolean if a field has been set.

### GetPostalCode

`func (o *TransactionCardAcceptor) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *TransactionCardAcceptor) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *TransactionCardAcceptor) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *TransactionCardAcceptor) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetSpecialMerchantId

`func (o *TransactionCardAcceptor) GetSpecialMerchantId() string`

GetSpecialMerchantId returns the SpecialMerchantId field if non-nil, zero value otherwise.

### GetSpecialMerchantIdOk

`func (o *TransactionCardAcceptor) GetSpecialMerchantIdOk() (*string, bool)`

GetSpecialMerchantIdOk returns a tuple with the SpecialMerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialMerchantId

`func (o *TransactionCardAcceptor) SetSpecialMerchantId(v string)`

SetSpecialMerchantId sets SpecialMerchantId field to given value.

### HasSpecialMerchantId

`func (o *TransactionCardAcceptor) HasSpecialMerchantId() bool`

HasSpecialMerchantId returns a boolean if a field has been set.

### GetState

`func (o *TransactionCardAcceptor) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *TransactionCardAcceptor) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *TransactionCardAcceptor) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *TransactionCardAcceptor) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSubMerchantId

`func (o *TransactionCardAcceptor) GetSubMerchantId() string`

GetSubMerchantId returns the SubMerchantId field if non-nil, zero value otherwise.

### GetSubMerchantIdOk

`func (o *TransactionCardAcceptor) GetSubMerchantIdOk() (*string, bool)`

GetSubMerchantIdOk returns a tuple with the SubMerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubMerchantId

`func (o *TransactionCardAcceptor) SetSubMerchantId(v string)`

SetSubMerchantId sets SubMerchantId field to given value.

### HasSubMerchantId

`func (o *TransactionCardAcceptor) HasSubMerchantId() bool`

HasSubMerchantId returns a boolean if a field has been set.

### GetTransferServiceProviderName

`func (o *TransactionCardAcceptor) GetTransferServiceProviderName() string`

GetTransferServiceProviderName returns the TransferServiceProviderName field if non-nil, zero value otherwise.

### GetTransferServiceProviderNameOk

`func (o *TransactionCardAcceptor) GetTransferServiceProviderNameOk() (*string, bool)`

GetTransferServiceProviderNameOk returns a tuple with the TransferServiceProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferServiceProviderName

`func (o *TransactionCardAcceptor) SetTransferServiceProviderName(v string)`

SetTransferServiceProviderName sets TransferServiceProviderName field to given value.

### HasTransferServiceProviderName

`func (o *TransactionCardAcceptor) HasTransferServiceProviderName() bool`

HasTransferServiceProviderName returns a boolean if a field has been set.

### GetUrl

`func (o *TransactionCardAcceptor) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TransactionCardAcceptor) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TransactionCardAcceptor) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *TransactionCardAcceptor) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


