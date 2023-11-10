# TerminalModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardPresence** | Pointer to **string** | Indicates whether the card was present during the transaction. | [optional] 
**CardholderPresence** | Pointer to **string** | Indicates whether the cardholder was present during the transaction. | [optional] 
**PartialApprovalCapable** | Pointer to **string** | Indicates whether the card acceptor or terminal supports partial-approval transactions. | [optional] 
**PinPresent** | Pointer to **string** | Indicates whether the cardholder entered a PIN during the transaction. | [optional] 
**SpecialConditionIndicator** | Pointer to **string** | Indicates a higher-risk operation, such as a quasi-cash or cryptocurrency transaction.  These transactions typically involve non-financial institutions. | [optional] 
**Tid** | Pointer to **string** | Card acceptor or terminal identification number. | [optional] 

## Methods

### NewTerminalModel

`func NewTerminalModel() *TerminalModel`

NewTerminalModel instantiates a new TerminalModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTerminalModelWithDefaults

`func NewTerminalModelWithDefaults() *TerminalModel`

NewTerminalModelWithDefaults instantiates a new TerminalModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardPresence

`func (o *TerminalModel) GetCardPresence() string`

GetCardPresence returns the CardPresence field if non-nil, zero value otherwise.

### GetCardPresenceOk

`func (o *TerminalModel) GetCardPresenceOk() (*string, bool)`

GetCardPresenceOk returns a tuple with the CardPresence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardPresence

`func (o *TerminalModel) SetCardPresence(v string)`

SetCardPresence sets CardPresence field to given value.

### HasCardPresence

`func (o *TerminalModel) HasCardPresence() bool`

HasCardPresence returns a boolean if a field has been set.

### GetCardholderPresence

`func (o *TerminalModel) GetCardholderPresence() string`

GetCardholderPresence returns the CardholderPresence field if non-nil, zero value otherwise.

### GetCardholderPresenceOk

`func (o *TerminalModel) GetCardholderPresenceOk() (*string, bool)`

GetCardholderPresenceOk returns a tuple with the CardholderPresence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardholderPresence

`func (o *TerminalModel) SetCardholderPresence(v string)`

SetCardholderPresence sets CardholderPresence field to given value.

### HasCardholderPresence

`func (o *TerminalModel) HasCardholderPresence() bool`

HasCardholderPresence returns a boolean if a field has been set.

### GetPartialApprovalCapable

`func (o *TerminalModel) GetPartialApprovalCapable() string`

GetPartialApprovalCapable returns the PartialApprovalCapable field if non-nil, zero value otherwise.

### GetPartialApprovalCapableOk

`func (o *TerminalModel) GetPartialApprovalCapableOk() (*string, bool)`

GetPartialApprovalCapableOk returns a tuple with the PartialApprovalCapable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartialApprovalCapable

`func (o *TerminalModel) SetPartialApprovalCapable(v string)`

SetPartialApprovalCapable sets PartialApprovalCapable field to given value.

### HasPartialApprovalCapable

`func (o *TerminalModel) HasPartialApprovalCapable() bool`

HasPartialApprovalCapable returns a boolean if a field has been set.

### GetPinPresent

`func (o *TerminalModel) GetPinPresent() string`

GetPinPresent returns the PinPresent field if non-nil, zero value otherwise.

### GetPinPresentOk

`func (o *TerminalModel) GetPinPresentOk() (*string, bool)`

GetPinPresentOk returns a tuple with the PinPresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinPresent

`func (o *TerminalModel) SetPinPresent(v string)`

SetPinPresent sets PinPresent field to given value.

### HasPinPresent

`func (o *TerminalModel) HasPinPresent() bool`

HasPinPresent returns a boolean if a field has been set.

### GetSpecialConditionIndicator

`func (o *TerminalModel) GetSpecialConditionIndicator() string`

GetSpecialConditionIndicator returns the SpecialConditionIndicator field if non-nil, zero value otherwise.

### GetSpecialConditionIndicatorOk

`func (o *TerminalModel) GetSpecialConditionIndicatorOk() (*string, bool)`

GetSpecialConditionIndicatorOk returns a tuple with the SpecialConditionIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialConditionIndicator

`func (o *TerminalModel) SetSpecialConditionIndicator(v string)`

SetSpecialConditionIndicator sets SpecialConditionIndicator field to given value.

### HasSpecialConditionIndicator

`func (o *TerminalModel) HasSpecialConditionIndicator() bool`

HasSpecialConditionIndicator returns a boolean if a field has been set.

### GetTid

`func (o *TerminalModel) GetTid() string`

GetTid returns the Tid field if non-nil, zero value otherwise.

### GetTidOk

`func (o *TerminalModel) GetTidOk() (*string, bool)`

GetTidOk returns a tuple with the Tid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTid

`func (o *TerminalModel) SetTid(v string)`

SetTid sets Tid field to given value.

### HasTid

`func (o *TerminalModel) HasTid() bool`

HasTid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


