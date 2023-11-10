# TransactionMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Airline** | Pointer to [**Airline**](Airline.md) |  | [optional] 
**AuthorizationLifeCycle** | Pointer to **int32** | Number of days the pre-authorization is in effect. | [optional] 
**CrossBorderTransaction** | Pointer to **bool** | Whether the transaction is cross-border, i.e., when the merchant and the cardholder are located in two different countries. | [optional] [default to false]
**IsDeferredAuthorization** | Pointer to **bool** |  | [optional] 
**IsLodgingAutoRental** | Pointer to **bool** | Whether the transaction is a lodging or vehicle rental. | [optional] [default to false]
**LodgingAutoRentalStartDate** | Pointer to **time.Time** | Date and time when the lodging check-in or vehicle rental began. | [optional] 
**MotoIndicator** | Pointer to **string** | Indicates the type of mail or telephone order transaction. | [optional] 
**OneLegOut** | Pointer to **bool** |  | [optional] 
**PaymentChannel** | Pointer to **string** | Channel from which the transaction was originated. | [optional] 
**SpecialPurchaseId** | Pointer to **string** |  | [optional] 
**TransactionCategory** | Pointer to **string** | Industry for which the transaction was originated. | [optional] 
**Transit** | Pointer to [**Transit**](Transit.md) |  | [optional] 

## Methods

### NewTransactionMetadata

`func NewTransactionMetadata() *TransactionMetadata`

NewTransactionMetadata instantiates a new TransactionMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionMetadataWithDefaults

`func NewTransactionMetadataWithDefaults() *TransactionMetadata`

NewTransactionMetadataWithDefaults instantiates a new TransactionMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAirline

`func (o *TransactionMetadata) GetAirline() Airline`

GetAirline returns the Airline field if non-nil, zero value otherwise.

### GetAirlineOk

`func (o *TransactionMetadata) GetAirlineOk() (*Airline, bool)`

GetAirlineOk returns a tuple with the Airline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirline

`func (o *TransactionMetadata) SetAirline(v Airline)`

SetAirline sets Airline field to given value.

### HasAirline

`func (o *TransactionMetadata) HasAirline() bool`

HasAirline returns a boolean if a field has been set.

### GetAuthorizationLifeCycle

`func (o *TransactionMetadata) GetAuthorizationLifeCycle() int32`

GetAuthorizationLifeCycle returns the AuthorizationLifeCycle field if non-nil, zero value otherwise.

### GetAuthorizationLifeCycleOk

`func (o *TransactionMetadata) GetAuthorizationLifeCycleOk() (*int32, bool)`

GetAuthorizationLifeCycleOk returns a tuple with the AuthorizationLifeCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationLifeCycle

`func (o *TransactionMetadata) SetAuthorizationLifeCycle(v int32)`

SetAuthorizationLifeCycle sets AuthorizationLifeCycle field to given value.

### HasAuthorizationLifeCycle

`func (o *TransactionMetadata) HasAuthorizationLifeCycle() bool`

HasAuthorizationLifeCycle returns a boolean if a field has been set.

### GetCrossBorderTransaction

`func (o *TransactionMetadata) GetCrossBorderTransaction() bool`

GetCrossBorderTransaction returns the CrossBorderTransaction field if non-nil, zero value otherwise.

### GetCrossBorderTransactionOk

`func (o *TransactionMetadata) GetCrossBorderTransactionOk() (*bool, bool)`

GetCrossBorderTransactionOk returns a tuple with the CrossBorderTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossBorderTransaction

`func (o *TransactionMetadata) SetCrossBorderTransaction(v bool)`

SetCrossBorderTransaction sets CrossBorderTransaction field to given value.

### HasCrossBorderTransaction

`func (o *TransactionMetadata) HasCrossBorderTransaction() bool`

HasCrossBorderTransaction returns a boolean if a field has been set.

### GetIsDeferredAuthorization

`func (o *TransactionMetadata) GetIsDeferredAuthorization() bool`

GetIsDeferredAuthorization returns the IsDeferredAuthorization field if non-nil, zero value otherwise.

### GetIsDeferredAuthorizationOk

`func (o *TransactionMetadata) GetIsDeferredAuthorizationOk() (*bool, bool)`

GetIsDeferredAuthorizationOk returns a tuple with the IsDeferredAuthorization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeferredAuthorization

`func (o *TransactionMetadata) SetIsDeferredAuthorization(v bool)`

SetIsDeferredAuthorization sets IsDeferredAuthorization field to given value.

### HasIsDeferredAuthorization

`func (o *TransactionMetadata) HasIsDeferredAuthorization() bool`

HasIsDeferredAuthorization returns a boolean if a field has been set.

### GetIsLodgingAutoRental

`func (o *TransactionMetadata) GetIsLodgingAutoRental() bool`

GetIsLodgingAutoRental returns the IsLodgingAutoRental field if non-nil, zero value otherwise.

### GetIsLodgingAutoRentalOk

`func (o *TransactionMetadata) GetIsLodgingAutoRentalOk() (*bool, bool)`

GetIsLodgingAutoRentalOk returns a tuple with the IsLodgingAutoRental field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLodgingAutoRental

`func (o *TransactionMetadata) SetIsLodgingAutoRental(v bool)`

SetIsLodgingAutoRental sets IsLodgingAutoRental field to given value.

### HasIsLodgingAutoRental

`func (o *TransactionMetadata) HasIsLodgingAutoRental() bool`

HasIsLodgingAutoRental returns a boolean if a field has been set.

### GetLodgingAutoRentalStartDate

`func (o *TransactionMetadata) GetLodgingAutoRentalStartDate() time.Time`

GetLodgingAutoRentalStartDate returns the LodgingAutoRentalStartDate field if non-nil, zero value otherwise.

### GetLodgingAutoRentalStartDateOk

`func (o *TransactionMetadata) GetLodgingAutoRentalStartDateOk() (*time.Time, bool)`

GetLodgingAutoRentalStartDateOk returns a tuple with the LodgingAutoRentalStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLodgingAutoRentalStartDate

`func (o *TransactionMetadata) SetLodgingAutoRentalStartDate(v time.Time)`

SetLodgingAutoRentalStartDate sets LodgingAutoRentalStartDate field to given value.

### HasLodgingAutoRentalStartDate

`func (o *TransactionMetadata) HasLodgingAutoRentalStartDate() bool`

HasLodgingAutoRentalStartDate returns a boolean if a field has been set.

### GetMotoIndicator

`func (o *TransactionMetadata) GetMotoIndicator() string`

GetMotoIndicator returns the MotoIndicator field if non-nil, zero value otherwise.

### GetMotoIndicatorOk

`func (o *TransactionMetadata) GetMotoIndicatorOk() (*string, bool)`

GetMotoIndicatorOk returns a tuple with the MotoIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotoIndicator

`func (o *TransactionMetadata) SetMotoIndicator(v string)`

SetMotoIndicator sets MotoIndicator field to given value.

### HasMotoIndicator

`func (o *TransactionMetadata) HasMotoIndicator() bool`

HasMotoIndicator returns a boolean if a field has been set.

### GetOneLegOut

`func (o *TransactionMetadata) GetOneLegOut() bool`

GetOneLegOut returns the OneLegOut field if non-nil, zero value otherwise.

### GetOneLegOutOk

`func (o *TransactionMetadata) GetOneLegOutOk() (*bool, bool)`

GetOneLegOutOk returns a tuple with the OneLegOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOneLegOut

`func (o *TransactionMetadata) SetOneLegOut(v bool)`

SetOneLegOut sets OneLegOut field to given value.

### HasOneLegOut

`func (o *TransactionMetadata) HasOneLegOut() bool`

HasOneLegOut returns a boolean if a field has been set.

### GetPaymentChannel

`func (o *TransactionMetadata) GetPaymentChannel() string`

GetPaymentChannel returns the PaymentChannel field if non-nil, zero value otherwise.

### GetPaymentChannelOk

`func (o *TransactionMetadata) GetPaymentChannelOk() (*string, bool)`

GetPaymentChannelOk returns a tuple with the PaymentChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentChannel

`func (o *TransactionMetadata) SetPaymentChannel(v string)`

SetPaymentChannel sets PaymentChannel field to given value.

### HasPaymentChannel

`func (o *TransactionMetadata) HasPaymentChannel() bool`

HasPaymentChannel returns a boolean if a field has been set.

### GetSpecialPurchaseId

`func (o *TransactionMetadata) GetSpecialPurchaseId() string`

GetSpecialPurchaseId returns the SpecialPurchaseId field if non-nil, zero value otherwise.

### GetSpecialPurchaseIdOk

`func (o *TransactionMetadata) GetSpecialPurchaseIdOk() (*string, bool)`

GetSpecialPurchaseIdOk returns a tuple with the SpecialPurchaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialPurchaseId

`func (o *TransactionMetadata) SetSpecialPurchaseId(v string)`

SetSpecialPurchaseId sets SpecialPurchaseId field to given value.

### HasSpecialPurchaseId

`func (o *TransactionMetadata) HasSpecialPurchaseId() bool`

HasSpecialPurchaseId returns a boolean if a field has been set.

### GetTransactionCategory

`func (o *TransactionMetadata) GetTransactionCategory() string`

GetTransactionCategory returns the TransactionCategory field if non-nil, zero value otherwise.

### GetTransactionCategoryOk

`func (o *TransactionMetadata) GetTransactionCategoryOk() (*string, bool)`

GetTransactionCategoryOk returns a tuple with the TransactionCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionCategory

`func (o *TransactionMetadata) SetTransactionCategory(v string)`

SetTransactionCategory sets TransactionCategory field to given value.

### HasTransactionCategory

`func (o *TransactionMetadata) HasTransactionCategory() bool`

HasTransactionCategory returns a boolean if a field has been set.

### GetTransit

`func (o *TransactionMetadata) GetTransit() Transit`

GetTransit returns the Transit field if non-nil, zero value otherwise.

### GetTransitOk

`func (o *TransactionMetadata) GetTransitOk() (*Transit, bool)`

GetTransitOk returns a tuple with the Transit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransit

`func (o *TransactionMetadata) SetTransit(v Transit)`

SetTransit sets Transit field to given value.

### HasTransit

`func (o *TransactionMetadata) HasTransit() bool`

HasTransit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


