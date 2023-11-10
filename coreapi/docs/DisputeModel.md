# DisputeModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaseManagementIdentifier** | Pointer to **string** | The unique identifier of the dispute case. | [optional] 
**Reason** | Pointer to **string** | The reason for the dispute. | [optional] 

## Methods

### NewDisputeModel

`func NewDisputeModel() *DisputeModel`

NewDisputeModel instantiates a new DisputeModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDisputeModelWithDefaults

`func NewDisputeModelWithDefaults() *DisputeModel`

NewDisputeModelWithDefaults instantiates a new DisputeModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaseManagementIdentifier

`func (o *DisputeModel) GetCaseManagementIdentifier() string`

GetCaseManagementIdentifier returns the CaseManagementIdentifier field if non-nil, zero value otherwise.

### GetCaseManagementIdentifierOk

`func (o *DisputeModel) GetCaseManagementIdentifierOk() (*string, bool)`

GetCaseManagementIdentifierOk returns a tuple with the CaseManagementIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseManagementIdentifier

`func (o *DisputeModel) SetCaseManagementIdentifier(v string)`

SetCaseManagementIdentifier sets CaseManagementIdentifier field to given value.

### HasCaseManagementIdentifier

`func (o *DisputeModel) HasCaseManagementIdentifier() bool`

HasCaseManagementIdentifier returns a boolean if a field has been set.

### GetReason

`func (o *DisputeModel) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *DisputeModel) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *DisputeModel) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *DisputeModel) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


