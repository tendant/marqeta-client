# BundleTransitionReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**BundleResourceStatus**](BundleResourceStatus.md) |  | 

## Methods

### NewBundleTransitionReq

`func NewBundleTransitionReq(status BundleResourceStatus, ) *BundleTransitionReq`

NewBundleTransitionReq instantiates a new BundleTransitionReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBundleTransitionReqWithDefaults

`func NewBundleTransitionReqWithDefaults() *BundleTransitionReq`

NewBundleTransitionReqWithDefaults instantiates a new BundleTransitionReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *BundleTransitionReq) GetStatus() BundleResourceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BundleTransitionReq) GetStatusOk() (*BundleResourceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BundleTransitionReq) SetStatus(v BundleResourceStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


