/*
Core API

Marqeta's Core API endpoints, conveniently annotated to enable code generation (including SDKs), test cases, and documentation. Currently in beta.

API version: 3.0.11
Contact: support@marqeta.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package marqeta_coreapi_client

import (
	"encoding/json"
)

// checks if the Carrier type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Carrier{}

// Carrier Specifies attributes of the card carrier.
type Carrier struct {
	// Specifies an image to display on the card carrier.
	LogoFile *string `json:"logo_file,omitempty"`
	// Specifies a thumbnail-sized rendering of the image specified in the `logo_file` field.
	LogoThumbnailFile *string `json:"logo_thumbnail_file,omitempty"`
	// Specifies a text file containing a custom message to print on the card carrier.
	MessageFile *string `json:"message_file,omitempty"`
	// Specifies a custom message that appears on the card carrier.
	MessageLine *string `json:"message_line,omitempty"`
	// Specifies the card carrier template to use.
	TemplateId *string `json:"template_id,omitempty"`
}

// NewCarrier instantiates a new Carrier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCarrier() *Carrier {
	this := Carrier{}
	return &this
}

// NewCarrierWithDefaults instantiates a new Carrier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCarrierWithDefaults() *Carrier {
	this := Carrier{}
	return &this
}

// GetLogoFile returns the LogoFile field value if set, zero value otherwise.
func (o *Carrier) GetLogoFile() string {
	if o == nil || IsNil(o.LogoFile) {
		var ret string
		return ret
	}
	return *o.LogoFile
}

// GetLogoFileOk returns a tuple with the LogoFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Carrier) GetLogoFileOk() (*string, bool) {
	if o == nil || IsNil(o.LogoFile) {
		return nil, false
	}
	return o.LogoFile, true
}

// HasLogoFile returns a boolean if a field has been set.
func (o *Carrier) HasLogoFile() bool {
	if o != nil && !IsNil(o.LogoFile) {
		return true
	}

	return false
}

// SetLogoFile gets a reference to the given string and assigns it to the LogoFile field.
func (o *Carrier) SetLogoFile(v string) {
	o.LogoFile = &v
}

// GetLogoThumbnailFile returns the LogoThumbnailFile field value if set, zero value otherwise.
func (o *Carrier) GetLogoThumbnailFile() string {
	if o == nil || IsNil(o.LogoThumbnailFile) {
		var ret string
		return ret
	}
	return *o.LogoThumbnailFile
}

// GetLogoThumbnailFileOk returns a tuple with the LogoThumbnailFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Carrier) GetLogoThumbnailFileOk() (*string, bool) {
	if o == nil || IsNil(o.LogoThumbnailFile) {
		return nil, false
	}
	return o.LogoThumbnailFile, true
}

// HasLogoThumbnailFile returns a boolean if a field has been set.
func (o *Carrier) HasLogoThumbnailFile() bool {
	if o != nil && !IsNil(o.LogoThumbnailFile) {
		return true
	}

	return false
}

// SetLogoThumbnailFile gets a reference to the given string and assigns it to the LogoThumbnailFile field.
func (o *Carrier) SetLogoThumbnailFile(v string) {
	o.LogoThumbnailFile = &v
}

// GetMessageFile returns the MessageFile field value if set, zero value otherwise.
func (o *Carrier) GetMessageFile() string {
	if o == nil || IsNil(o.MessageFile) {
		var ret string
		return ret
	}
	return *o.MessageFile
}

// GetMessageFileOk returns a tuple with the MessageFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Carrier) GetMessageFileOk() (*string, bool) {
	if o == nil || IsNil(o.MessageFile) {
		return nil, false
	}
	return o.MessageFile, true
}

// HasMessageFile returns a boolean if a field has been set.
func (o *Carrier) HasMessageFile() bool {
	if o != nil && !IsNil(o.MessageFile) {
		return true
	}

	return false
}

// SetMessageFile gets a reference to the given string and assigns it to the MessageFile field.
func (o *Carrier) SetMessageFile(v string) {
	o.MessageFile = &v
}

// GetMessageLine returns the MessageLine field value if set, zero value otherwise.
func (o *Carrier) GetMessageLine() string {
	if o == nil || IsNil(o.MessageLine) {
		var ret string
		return ret
	}
	return *o.MessageLine
}

// GetMessageLineOk returns a tuple with the MessageLine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Carrier) GetMessageLineOk() (*string, bool) {
	if o == nil || IsNil(o.MessageLine) {
		return nil, false
	}
	return o.MessageLine, true
}

// HasMessageLine returns a boolean if a field has been set.
func (o *Carrier) HasMessageLine() bool {
	if o != nil && !IsNil(o.MessageLine) {
		return true
	}

	return false
}

// SetMessageLine gets a reference to the given string and assigns it to the MessageLine field.
func (o *Carrier) SetMessageLine(v string) {
	o.MessageLine = &v
}

// GetTemplateId returns the TemplateId field value if set, zero value otherwise.
func (o *Carrier) GetTemplateId() string {
	if o == nil || IsNil(o.TemplateId) {
		var ret string
		return ret
	}
	return *o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Carrier) GetTemplateIdOk() (*string, bool) {
	if o == nil || IsNil(o.TemplateId) {
		return nil, false
	}
	return o.TemplateId, true
}

// HasTemplateId returns a boolean if a field has been set.
func (o *Carrier) HasTemplateId() bool {
	if o != nil && !IsNil(o.TemplateId) {
		return true
	}

	return false
}

// SetTemplateId gets a reference to the given string and assigns it to the TemplateId field.
func (o *Carrier) SetTemplateId(v string) {
	o.TemplateId = &v
}

func (o Carrier) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Carrier) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LogoFile) {
		toSerialize["logo_file"] = o.LogoFile
	}
	if !IsNil(o.LogoThumbnailFile) {
		toSerialize["logo_thumbnail_file"] = o.LogoThumbnailFile
	}
	if !IsNil(o.MessageFile) {
		toSerialize["message_file"] = o.MessageFile
	}
	if !IsNil(o.MessageLine) {
		toSerialize["message_line"] = o.MessageLine
	}
	if !IsNil(o.TemplateId) {
		toSerialize["template_id"] = o.TemplateId
	}
	return toSerialize, nil
}

type NullableCarrier struct {
	value *Carrier
	isSet bool
}

func (v NullableCarrier) Get() *Carrier {
	return v.value
}

func (v *NullableCarrier) Set(val *Carrier) {
	v.value = val
	v.isSet = true
}

func (v NullableCarrier) IsSet() bool {
	return v.isSet
}

func (v *NullableCarrier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCarrier(val *Carrier) *NullableCarrier {
	return &NullableCarrier{value: val, isSet: true}
}

func (v NullableCarrier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCarrier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


