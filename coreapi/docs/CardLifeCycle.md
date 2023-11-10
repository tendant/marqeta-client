# CardLifeCycle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivateUponIssue** | Pointer to **bool** | A value of &#x60;true&#x60; indicates that cards of this card product type are active once they are issued. | [optional] [default to false]
**CardServiceCode** | Pointer to **int32** | Sequence of three digits that defines various services, differentiates card usage in international or domestic interchange, designates personal identification number (PIN) and authorization requirements, and identifies card restrictions. The following values are commonly used:  *First digit*  * *1* — International interchange OK * *2* — International interchange, use IC (chip) where feasible * *5* — National interchange only except under bilateral agreement * *6* — National interchange only except under bilateral agreement, use IC (chip) where feasible * *7* — No interchange except under bilateral agreement (closed loop) * *9* — Test  *Second digit*  * *0* — Normal * *2* — Contact issuer via online means * *4* — Contact issuer via online means except under bilateral agreement  *Third digit*  * *0* — No restrictions, PIN required * *1* — No restrictions * *2* — Goods and services only (no cash) * *3* — ATM only, PIN required * *4* — Cash only * *5* — Goods and services only (no cash), PIN required * *6* — No restrictions, use PIN where feasible * *7* — Goods and services only (no cash), use PIN where feasible | [optional] [default to 101]
**ExpirationOffset** | Pointer to [**ExpirationOffsetWithMinimum**](ExpirationOffsetWithMinimum.md) |  | [optional] 
**UpdateExpirationUponActivation** | Pointer to **bool** | Normally, the &#x60;expiration_offset&#x60; is measured from the date of issue. Set this field to &#x60;true&#x60; to measure &#x60;expiration_offset&#x60; from the date of activation instead. | [optional] [default to false]

## Methods

### NewCardLifeCycle

`func NewCardLifeCycle() *CardLifeCycle`

NewCardLifeCycle instantiates a new CardLifeCycle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardLifeCycleWithDefaults

`func NewCardLifeCycleWithDefaults() *CardLifeCycle`

NewCardLifeCycleWithDefaults instantiates a new CardLifeCycle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivateUponIssue

`func (o *CardLifeCycle) GetActivateUponIssue() bool`

GetActivateUponIssue returns the ActivateUponIssue field if non-nil, zero value otherwise.

### GetActivateUponIssueOk

`func (o *CardLifeCycle) GetActivateUponIssueOk() (*bool, bool)`

GetActivateUponIssueOk returns a tuple with the ActivateUponIssue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivateUponIssue

`func (o *CardLifeCycle) SetActivateUponIssue(v bool)`

SetActivateUponIssue sets ActivateUponIssue field to given value.

### HasActivateUponIssue

`func (o *CardLifeCycle) HasActivateUponIssue() bool`

HasActivateUponIssue returns a boolean if a field has been set.

### GetCardServiceCode

`func (o *CardLifeCycle) GetCardServiceCode() int32`

GetCardServiceCode returns the CardServiceCode field if non-nil, zero value otherwise.

### GetCardServiceCodeOk

`func (o *CardLifeCycle) GetCardServiceCodeOk() (*int32, bool)`

GetCardServiceCodeOk returns a tuple with the CardServiceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardServiceCode

`func (o *CardLifeCycle) SetCardServiceCode(v int32)`

SetCardServiceCode sets CardServiceCode field to given value.

### HasCardServiceCode

`func (o *CardLifeCycle) HasCardServiceCode() bool`

HasCardServiceCode returns a boolean if a field has been set.

### GetExpirationOffset

`func (o *CardLifeCycle) GetExpirationOffset() ExpirationOffsetWithMinimum`

GetExpirationOffset returns the ExpirationOffset field if non-nil, zero value otherwise.

### GetExpirationOffsetOk

`func (o *CardLifeCycle) GetExpirationOffsetOk() (*ExpirationOffsetWithMinimum, bool)`

GetExpirationOffsetOk returns a tuple with the ExpirationOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationOffset

`func (o *CardLifeCycle) SetExpirationOffset(v ExpirationOffsetWithMinimum)`

SetExpirationOffset sets ExpirationOffset field to given value.

### HasExpirationOffset

`func (o *CardLifeCycle) HasExpirationOffset() bool`

HasExpirationOffset returns a boolean if a field has been set.

### GetUpdateExpirationUponActivation

`func (o *CardLifeCycle) GetUpdateExpirationUponActivation() bool`

GetUpdateExpirationUponActivation returns the UpdateExpirationUponActivation field if non-nil, zero value otherwise.

### GetUpdateExpirationUponActivationOk

`func (o *CardLifeCycle) GetUpdateExpirationUponActivationOk() (*bool, bool)`

GetUpdateExpirationUponActivationOk returns a tuple with the UpdateExpirationUponActivation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateExpirationUponActivation

`func (o *CardLifeCycle) SetUpdateExpirationUponActivation(v bool)`

SetUpdateExpirationUponActivation sets UpdateExpirationUponActivation field to given value.

### HasUpdateExpirationUponActivation

`func (o *CardLifeCycle) HasUpdateExpirationUponActivation() bool`

HasUpdateExpirationUponActivation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


