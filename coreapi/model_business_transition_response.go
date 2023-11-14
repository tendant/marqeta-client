/*
Core API

Marqeta's Core API endpoints, conveniently annotated to enable code generation (including SDKs), test cases, and documentation. Currently in beta.

API version: 3.0.11
Contact: support@marqeta.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package coreapi

import (
	"encoding/json"
	"time"
)

// checks if the BusinessTransitionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BusinessTransitionResponse{}

// BusinessTransitionResponse struct for BusinessTransitionResponse
type BusinessTransitionResponse struct {
	// Unique identifier of the business whose status you want to transition.
	BusinessToken *string `json:"business_token,omitempty"`
	// Mechanism by which the transaction was initiated.
	Channel string `json:"channel"`
	// Date and time when the resource was created, in UTC.
	CreatedTime *time.Time `json:"created_time,omitempty"`
	// Date and time when the resource was last updated, in UTC.
	LastModifiedTime *time.Time `json:"last_modified_time,omitempty"`
	// Additional information about the status change.
	Reason *string `json:"reason,omitempty"`
	// Identifies the standardized reason for the transition:  *00:* Object activated for the first time.  *01:* Requested by you.  *02:* Inactivity over time.  *03:* This address cannot accept mail or the addressee is unknown.  *04:* Negative account balance.  *05:* Account under review.  *06:* Suspicious activity was identified.  *07:* Activity outside the program parameters was identified.  *08:* Confirmed fraud was identified.  *09:* Matched with an Office of Foreign Assets Control list.  *10:* Card was reported lost.  *11:* Card information was cloned.  *12:* Account or card information was compromised.  *13:* Temporary status change while on hold/leave.  *14:* Initiated by Marqeta.  *15:* Initiated by issuer.  *16:* Card expired.  *17:* Failed KYC.  *18:* Changed to `ACTIVE` because information was properly validated.  *19:* Changed to `ACTIVE` because account activity was properly validated.  *20:* Change occurred prior to the normalization of reason codes.  *21:* Initiated by a third party, often a digital wallet provider.  *22:* PIN retry limit reached.  *23:* Card was reported stolen.  *24:* Address issue.  *25:* Name issue.  *26:* SSN issue.  *27:* DOB issue.  *28:* Email issue.  *29:* Phone issue.  *30:* Account/fulfillment mismatch.  *31:* Other reason.
	ReasonCode string `json:"reason_code"`
	// Specifies the status of the business.
	Status string `json:"status"`
	// Unique identifier of the business transition.
	Token string `json:"token"`
}

// NewBusinessTransitionResponse instantiates a new BusinessTransitionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBusinessTransitionResponse(channel string, reasonCode string, status string, token string) *BusinessTransitionResponse {
	this := BusinessTransitionResponse{}
	this.Channel = channel
	this.ReasonCode = reasonCode
	this.Status = status
	this.Token = token
	return &this
}

// NewBusinessTransitionResponseWithDefaults instantiates a new BusinessTransitionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBusinessTransitionResponseWithDefaults() *BusinessTransitionResponse {
	this := BusinessTransitionResponse{}
	return &this
}

// GetBusinessToken returns the BusinessToken field value if set, zero value otherwise.
func (o *BusinessTransitionResponse) GetBusinessToken() string {
	if o == nil || IsNil(o.BusinessToken) {
		var ret string
		return ret
	}
	return *o.BusinessToken
}

// GetBusinessTokenOk returns a tuple with the BusinessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BusinessTransitionResponse) GetBusinessTokenOk() (*string, bool) {
	if o == nil || IsNil(o.BusinessToken) {
		return nil, false
	}
	return o.BusinessToken, true
}

// HasBusinessToken returns a boolean if a field has been set.
func (o *BusinessTransitionResponse) HasBusinessToken() bool {
	if o != nil && !IsNil(o.BusinessToken) {
		return true
	}

	return false
}

// SetBusinessToken gets a reference to the given string and assigns it to the BusinessToken field.
func (o *BusinessTransitionResponse) SetBusinessToken(v string) {
	o.BusinessToken = &v
}

// GetChannel returns the Channel field value
func (o *BusinessTransitionResponse) GetChannel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Channel
}

// GetChannelOk returns a tuple with the Channel field value
// and a boolean to check if the value has been set.
func (o *BusinessTransitionResponse) GetChannelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Channel, true
}

// SetChannel sets field value
func (o *BusinessTransitionResponse) SetChannel(v string) {
	o.Channel = v
}

// GetCreatedTime returns the CreatedTime field value if set, zero value otherwise.
func (o *BusinessTransitionResponse) GetCreatedTime() time.Time {
	if o == nil || IsNil(o.CreatedTime) {
		var ret time.Time
		return ret
	}
	return *o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BusinessTransitionResponse) GetCreatedTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedTime) {
		return nil, false
	}
	return o.CreatedTime, true
}

// HasCreatedTime returns a boolean if a field has been set.
func (o *BusinessTransitionResponse) HasCreatedTime() bool {
	if o != nil && !IsNil(o.CreatedTime) {
		return true
	}

	return false
}

// SetCreatedTime gets a reference to the given time.Time and assigns it to the CreatedTime field.
func (o *BusinessTransitionResponse) SetCreatedTime(v time.Time) {
	o.CreatedTime = &v
}

// GetLastModifiedTime returns the LastModifiedTime field value if set, zero value otherwise.
func (o *BusinessTransitionResponse) GetLastModifiedTime() time.Time {
	if o == nil || IsNil(o.LastModifiedTime) {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedTime
}

// GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BusinessTransitionResponse) GetLastModifiedTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastModifiedTime) {
		return nil, false
	}
	return o.LastModifiedTime, true
}

// HasLastModifiedTime returns a boolean if a field has been set.
func (o *BusinessTransitionResponse) HasLastModifiedTime() bool {
	if o != nil && !IsNil(o.LastModifiedTime) {
		return true
	}

	return false
}

// SetLastModifiedTime gets a reference to the given time.Time and assigns it to the LastModifiedTime field.
func (o *BusinessTransitionResponse) SetLastModifiedTime(v time.Time) {
	o.LastModifiedTime = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *BusinessTransitionResponse) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BusinessTransitionResponse) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *BusinessTransitionResponse) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *BusinessTransitionResponse) SetReason(v string) {
	o.Reason = &v
}

// GetReasonCode returns the ReasonCode field value
func (o *BusinessTransitionResponse) GetReasonCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReasonCode
}

// GetReasonCodeOk returns a tuple with the ReasonCode field value
// and a boolean to check if the value has been set.
func (o *BusinessTransitionResponse) GetReasonCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReasonCode, true
}

// SetReasonCode sets field value
func (o *BusinessTransitionResponse) SetReasonCode(v string) {
	o.ReasonCode = v
}

// GetStatus returns the Status field value
func (o *BusinessTransitionResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *BusinessTransitionResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *BusinessTransitionResponse) SetStatus(v string) {
	o.Status = v
}

// GetToken returns the Token field value
func (o *BusinessTransitionResponse) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *BusinessTransitionResponse) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *BusinessTransitionResponse) SetToken(v string) {
	o.Token = v
}

func (o BusinessTransitionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BusinessTransitionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BusinessToken) {
		toSerialize["business_token"] = o.BusinessToken
	}
	toSerialize["channel"] = o.Channel
	if !IsNil(o.CreatedTime) {
		toSerialize["created_time"] = o.CreatedTime
	}
	if !IsNil(o.LastModifiedTime) {
		toSerialize["last_modified_time"] = o.LastModifiedTime
	}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	toSerialize["reason_code"] = o.ReasonCode
	toSerialize["status"] = o.Status
	toSerialize["token"] = o.Token
	return toSerialize, nil
}

type NullableBusinessTransitionResponse struct {
	value *BusinessTransitionResponse
	isSet bool
}

func (v NullableBusinessTransitionResponse) Get() *BusinessTransitionResponse {
	return v.value
}

func (v *NullableBusinessTransitionResponse) Set(val *BusinessTransitionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBusinessTransitionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBusinessTransitionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBusinessTransitionResponse(val *BusinessTransitionResponse) *NullableBusinessTransitionResponse {
	return &NullableBusinessTransitionResponse{value: val, isSet: true}
}

func (v NullableBusinessTransitionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBusinessTransitionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


