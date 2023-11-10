# Link

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | **string** | URL of the requested &#x60;/balances&#x60; resource. | [readonly] 
**Method** | **string** | The HTTP method of the link. | 
**Rel** | **string** | Specifies the relationship between the current resource and the linked resource. | 

## Methods

### NewLink

`func NewLink(href string, method string, rel string, ) *Link`

NewLink instantiates a new Link object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkWithDefaults

`func NewLinkWithDefaults() *Link`

NewLinkWithDefaults instantiates a new Link object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *Link) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *Link) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *Link) SetHref(v string)`

SetHref sets Href field to given value.


### GetMethod

`func (o *Link) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *Link) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *Link) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetRel

`func (o *Link) GetRel() string`

GetRel returns the Rel field if non-nil, zero value otherwise.

### GetRelOk

`func (o *Link) GetRelOk() (*string, bool)`

GetRelOk returns a tuple with the Rel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRel

`func (o *Link) SetRel(v string)`

SetRel sets Rel field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


