# PolicyDocumentTemplateReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TemplateToken** | **string** | Unique identifier of a template, which is a document that serves as an initial disclosure but does not contain finalized values. | 

## Methods

### NewPolicyDocumentTemplateReq

`func NewPolicyDocumentTemplateReq(templateToken string, ) *PolicyDocumentTemplateReq`

NewPolicyDocumentTemplateReq instantiates a new PolicyDocumentTemplateReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyDocumentTemplateReqWithDefaults

`func NewPolicyDocumentTemplateReqWithDefaults() *PolicyDocumentTemplateReq`

NewPolicyDocumentTemplateReqWithDefaults instantiates a new PolicyDocumentTemplateReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplateToken

`func (o *PolicyDocumentTemplateReq) GetTemplateToken() string`

GetTemplateToken returns the TemplateToken field if non-nil, zero value otherwise.

### GetTemplateTokenOk

`func (o *PolicyDocumentTemplateReq) GetTemplateTokenOk() (*string, bool)`

GetTemplateTokenOk returns a tuple with the TemplateToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateToken

`func (o *PolicyDocumentTemplateReq) SetTemplateToken(v string)`

SetTemplateToken sets TemplateToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


