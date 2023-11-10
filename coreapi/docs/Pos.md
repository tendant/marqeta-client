# Pos

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardDataInputCapability** | Pointer to **string** | How the terminal accepts card data. | [optional] 
**CardHolderPresence** | Pointer to **bool** | Whether the cardholder was present during the transaction. | [optional] [default to false]
**CardPresence** | Pointer to **bool** | Whether the card was present during the transaction. | [optional] [default to false]
**CardholderAuthenticationMethod** | Pointer to **string** | Method used to authenticate the cardholder. | [optional] 
**CountryCode** | Pointer to **string** | Country code of the card acceptor or terminal. | [optional] 
**IsInstallment** | Pointer to **bool** | Whether the transaction is an installment payment. | [optional] [default to false]
**IsRecurring** | Pointer to **bool** | Whether the transaction is recurring. | [optional] [default to false]
**PanEntryMode** | Pointer to **string** | Method used for capturing the card primary account number (PAN) during the transaction. | [optional] 
**PartialApprovalCapable** | Pointer to **bool** | Indicates whether the card acceptor or terminal supports partial-approval transactions. | [optional] [default to false]
**PinEntryMode** | Pointer to **string** | Indicates whether the card acceptor or terminal can capture card personal identification numbers (PINs).  *NOTE:* This field does not indicate whether a PIN was entered. | [optional] 
**PinPresent** | Pointer to **bool** | Indicates whether the cardholder entered a PIN during the transaction. | [optional] [default to false]
**PurchaseAmountOnly** | Pointer to **bool** | Indicates whether the card acceptor or terminal supports purchase-only approvals. | [optional] [default to false]
**SpecialConditionIndicator** | Pointer to **string** | Indicates a higher-risk operation, such as a quasi-cash or cryptocurrency transaction.  These transactions typically involve non-financial institutions. | [optional] 
**TerminalAttendance** | Pointer to **string** | Whether the card acceptor/terminal was attended. | [optional] 
**TerminalId** | Pointer to **string** | Card acceptor or terminal identification number. | [optional] 
**TerminalLocation** | Pointer to **string** | Location of the card acceptor/terminal. | [optional] 
**TerminalType** | Pointer to **string** | Type of card acceptor/terminal. | [optional] 
**Zip** | Pointer to **string** | United States ZIP code of the card acceptor or terminal. | [optional] 

## Methods

### NewPos

`func NewPos() *Pos`

NewPos instantiates a new Pos object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPosWithDefaults

`func NewPosWithDefaults() *Pos`

NewPosWithDefaults instantiates a new Pos object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardDataInputCapability

`func (o *Pos) GetCardDataInputCapability() string`

GetCardDataInputCapability returns the CardDataInputCapability field if non-nil, zero value otherwise.

### GetCardDataInputCapabilityOk

`func (o *Pos) GetCardDataInputCapabilityOk() (*string, bool)`

GetCardDataInputCapabilityOk returns a tuple with the CardDataInputCapability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardDataInputCapability

`func (o *Pos) SetCardDataInputCapability(v string)`

SetCardDataInputCapability sets CardDataInputCapability field to given value.

### HasCardDataInputCapability

`func (o *Pos) HasCardDataInputCapability() bool`

HasCardDataInputCapability returns a boolean if a field has been set.

### GetCardHolderPresence

`func (o *Pos) GetCardHolderPresence() bool`

GetCardHolderPresence returns the CardHolderPresence field if non-nil, zero value otherwise.

### GetCardHolderPresenceOk

`func (o *Pos) GetCardHolderPresenceOk() (*bool, bool)`

GetCardHolderPresenceOk returns a tuple with the CardHolderPresence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardHolderPresence

`func (o *Pos) SetCardHolderPresence(v bool)`

SetCardHolderPresence sets CardHolderPresence field to given value.

### HasCardHolderPresence

`func (o *Pos) HasCardHolderPresence() bool`

HasCardHolderPresence returns a boolean if a field has been set.

### GetCardPresence

`func (o *Pos) GetCardPresence() bool`

GetCardPresence returns the CardPresence field if non-nil, zero value otherwise.

### GetCardPresenceOk

`func (o *Pos) GetCardPresenceOk() (*bool, bool)`

GetCardPresenceOk returns a tuple with the CardPresence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardPresence

`func (o *Pos) SetCardPresence(v bool)`

SetCardPresence sets CardPresence field to given value.

### HasCardPresence

`func (o *Pos) HasCardPresence() bool`

HasCardPresence returns a boolean if a field has been set.

### GetCardholderAuthenticationMethod

`func (o *Pos) GetCardholderAuthenticationMethod() string`

GetCardholderAuthenticationMethod returns the CardholderAuthenticationMethod field if non-nil, zero value otherwise.

### GetCardholderAuthenticationMethodOk

`func (o *Pos) GetCardholderAuthenticationMethodOk() (*string, bool)`

GetCardholderAuthenticationMethodOk returns a tuple with the CardholderAuthenticationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardholderAuthenticationMethod

`func (o *Pos) SetCardholderAuthenticationMethod(v string)`

SetCardholderAuthenticationMethod sets CardholderAuthenticationMethod field to given value.

### HasCardholderAuthenticationMethod

`func (o *Pos) HasCardholderAuthenticationMethod() bool`

HasCardholderAuthenticationMethod returns a boolean if a field has been set.

### GetCountryCode

`func (o *Pos) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *Pos) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *Pos) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *Pos) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetIsInstallment

`func (o *Pos) GetIsInstallment() bool`

GetIsInstallment returns the IsInstallment field if non-nil, zero value otherwise.

### GetIsInstallmentOk

`func (o *Pos) GetIsInstallmentOk() (*bool, bool)`

GetIsInstallmentOk returns a tuple with the IsInstallment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInstallment

`func (o *Pos) SetIsInstallment(v bool)`

SetIsInstallment sets IsInstallment field to given value.

### HasIsInstallment

`func (o *Pos) HasIsInstallment() bool`

HasIsInstallment returns a boolean if a field has been set.

### GetIsRecurring

`func (o *Pos) GetIsRecurring() bool`

GetIsRecurring returns the IsRecurring field if non-nil, zero value otherwise.

### GetIsRecurringOk

`func (o *Pos) GetIsRecurringOk() (*bool, bool)`

GetIsRecurringOk returns a tuple with the IsRecurring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRecurring

`func (o *Pos) SetIsRecurring(v bool)`

SetIsRecurring sets IsRecurring field to given value.

### HasIsRecurring

`func (o *Pos) HasIsRecurring() bool`

HasIsRecurring returns a boolean if a field has been set.

### GetPanEntryMode

`func (o *Pos) GetPanEntryMode() string`

GetPanEntryMode returns the PanEntryMode field if non-nil, zero value otherwise.

### GetPanEntryModeOk

`func (o *Pos) GetPanEntryModeOk() (*string, bool)`

GetPanEntryModeOk returns a tuple with the PanEntryMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPanEntryMode

`func (o *Pos) SetPanEntryMode(v string)`

SetPanEntryMode sets PanEntryMode field to given value.

### HasPanEntryMode

`func (o *Pos) HasPanEntryMode() bool`

HasPanEntryMode returns a boolean if a field has been set.

### GetPartialApprovalCapable

`func (o *Pos) GetPartialApprovalCapable() bool`

GetPartialApprovalCapable returns the PartialApprovalCapable field if non-nil, zero value otherwise.

### GetPartialApprovalCapableOk

`func (o *Pos) GetPartialApprovalCapableOk() (*bool, bool)`

GetPartialApprovalCapableOk returns a tuple with the PartialApprovalCapable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartialApprovalCapable

`func (o *Pos) SetPartialApprovalCapable(v bool)`

SetPartialApprovalCapable sets PartialApprovalCapable field to given value.

### HasPartialApprovalCapable

`func (o *Pos) HasPartialApprovalCapable() bool`

HasPartialApprovalCapable returns a boolean if a field has been set.

### GetPinEntryMode

`func (o *Pos) GetPinEntryMode() string`

GetPinEntryMode returns the PinEntryMode field if non-nil, zero value otherwise.

### GetPinEntryModeOk

`func (o *Pos) GetPinEntryModeOk() (*string, bool)`

GetPinEntryModeOk returns a tuple with the PinEntryMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinEntryMode

`func (o *Pos) SetPinEntryMode(v string)`

SetPinEntryMode sets PinEntryMode field to given value.

### HasPinEntryMode

`func (o *Pos) HasPinEntryMode() bool`

HasPinEntryMode returns a boolean if a field has been set.

### GetPinPresent

`func (o *Pos) GetPinPresent() bool`

GetPinPresent returns the PinPresent field if non-nil, zero value otherwise.

### GetPinPresentOk

`func (o *Pos) GetPinPresentOk() (*bool, bool)`

GetPinPresentOk returns a tuple with the PinPresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinPresent

`func (o *Pos) SetPinPresent(v bool)`

SetPinPresent sets PinPresent field to given value.

### HasPinPresent

`func (o *Pos) HasPinPresent() bool`

HasPinPresent returns a boolean if a field has been set.

### GetPurchaseAmountOnly

`func (o *Pos) GetPurchaseAmountOnly() bool`

GetPurchaseAmountOnly returns the PurchaseAmountOnly field if non-nil, zero value otherwise.

### GetPurchaseAmountOnlyOk

`func (o *Pos) GetPurchaseAmountOnlyOk() (*bool, bool)`

GetPurchaseAmountOnlyOk returns a tuple with the PurchaseAmountOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseAmountOnly

`func (o *Pos) SetPurchaseAmountOnly(v bool)`

SetPurchaseAmountOnly sets PurchaseAmountOnly field to given value.

### HasPurchaseAmountOnly

`func (o *Pos) HasPurchaseAmountOnly() bool`

HasPurchaseAmountOnly returns a boolean if a field has been set.

### GetSpecialConditionIndicator

`func (o *Pos) GetSpecialConditionIndicator() string`

GetSpecialConditionIndicator returns the SpecialConditionIndicator field if non-nil, zero value otherwise.

### GetSpecialConditionIndicatorOk

`func (o *Pos) GetSpecialConditionIndicatorOk() (*string, bool)`

GetSpecialConditionIndicatorOk returns a tuple with the SpecialConditionIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialConditionIndicator

`func (o *Pos) SetSpecialConditionIndicator(v string)`

SetSpecialConditionIndicator sets SpecialConditionIndicator field to given value.

### HasSpecialConditionIndicator

`func (o *Pos) HasSpecialConditionIndicator() bool`

HasSpecialConditionIndicator returns a boolean if a field has been set.

### GetTerminalAttendance

`func (o *Pos) GetTerminalAttendance() string`

GetTerminalAttendance returns the TerminalAttendance field if non-nil, zero value otherwise.

### GetTerminalAttendanceOk

`func (o *Pos) GetTerminalAttendanceOk() (*string, bool)`

GetTerminalAttendanceOk returns a tuple with the TerminalAttendance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalAttendance

`func (o *Pos) SetTerminalAttendance(v string)`

SetTerminalAttendance sets TerminalAttendance field to given value.

### HasTerminalAttendance

`func (o *Pos) HasTerminalAttendance() bool`

HasTerminalAttendance returns a boolean if a field has been set.

### GetTerminalId

`func (o *Pos) GetTerminalId() string`

GetTerminalId returns the TerminalId field if non-nil, zero value otherwise.

### GetTerminalIdOk

`func (o *Pos) GetTerminalIdOk() (*string, bool)`

GetTerminalIdOk returns a tuple with the TerminalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalId

`func (o *Pos) SetTerminalId(v string)`

SetTerminalId sets TerminalId field to given value.

### HasTerminalId

`func (o *Pos) HasTerminalId() bool`

HasTerminalId returns a boolean if a field has been set.

### GetTerminalLocation

`func (o *Pos) GetTerminalLocation() string`

GetTerminalLocation returns the TerminalLocation field if non-nil, zero value otherwise.

### GetTerminalLocationOk

`func (o *Pos) GetTerminalLocationOk() (*string, bool)`

GetTerminalLocationOk returns a tuple with the TerminalLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalLocation

`func (o *Pos) SetTerminalLocation(v string)`

SetTerminalLocation sets TerminalLocation field to given value.

### HasTerminalLocation

`func (o *Pos) HasTerminalLocation() bool`

HasTerminalLocation returns a boolean if a field has been set.

### GetTerminalType

`func (o *Pos) GetTerminalType() string`

GetTerminalType returns the TerminalType field if non-nil, zero value otherwise.

### GetTerminalTypeOk

`func (o *Pos) GetTerminalTypeOk() (*string, bool)`

GetTerminalTypeOk returns a tuple with the TerminalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalType

`func (o *Pos) SetTerminalType(v string)`

SetTerminalType sets TerminalType field to given value.

### HasTerminalType

`func (o *Pos) HasTerminalType() bool`

HasTerminalType returns a boolean if a field has been set.

### GetZip

`func (o *Pos) GetZip() string`

GetZip returns the Zip field if non-nil, zero value otherwise.

### GetZipOk

`func (o *Pos) GetZipOk() (*string, bool)`

GetZipOk returns a tuple with the Zip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZip

`func (o *Pos) SetZip(v string)`

SetZip sets Zip field to given value.

### HasZip

`func (o *Pos) HasZip() bool`

HasZip returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


