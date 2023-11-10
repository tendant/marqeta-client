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

// checks if the UserTransitionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserTransitionRequest{}

// UserTransitionRequest struct for UserTransitionRequest
type UserTransitionRequest struct {
	// Mechanism by which the transaction was initiated.
	Channel string `json:"channel"`
	// Unique hashed value that identifies subsequent submissions of the user transition request.
	IdempotentHash *string `json:"idempotentHash,omitempty"`
	// Additional information about the status change.
	Reason *string `json:"reason,omitempty"`
	// Identifies the standardized reason for the transition:  *00:* Object activated for the first time.  *01:* Requested by you.  *02:* Inactivity over time.  *03:* This address cannot accept mail or the addressee is unknown.  *04:* Negative account balance.  *05:* Account under review.  *06:* Suspicious activity was identified.  *07:* Activity outside the program parameters was identified.  *08:* Confirmed fraud was identified.  *09:* Matched with an Office of Foreign Assets Control list.  *10:* Card was reported lost.  *11:* Card information was cloned.  *12:* Account or card information was compromised.  *13:* Temporary status change while on hold/leave.  *14:* Initiated by Marqeta.  *15:* Initiated by issuer.  *16:* Card expired.  *17:* Failed KYC.  *18:* Changed to `ACTIVE` because information was properly validated.  *19:* Changed to `ACTIVE` because account activity was properly validated.  *20:* Change occurred prior to the normalization of reason codes.  *21:* Initiated by a third party, often a digital wallet provider.  *22:* PIN retry limit reached.  *23:* Card was reported stolen.  *24:* Address issue.  *25:* Name issue.  *26:* SSN issue.  *27:* DOB issue.  *28:* Email issue.  *29:* Phone issue.  *30:* Account/fulfillment mismatch.  *31:* Other reason.
	ReasonCode string `json:"reason_code"`
	// Specifies the new status of the user.
	Status string `json:"status"`
	// Unique identifier of the user transition.  If you do not include a token, the system generates one automatically. This token is referenced in other API calls, so we recommend that you define a simple string that is easy to remember. This value cannot be updated.
	Token *string `json:"token,omitempty"`
	// Unique identifier of the user whose status you want to transition.
	UserToken string `json:"user_token"`
}

// NewUserTransitionRequest instantiates a new UserTransitionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserTransitionRequest(channel string, reasonCode string, status string, userToken string) *UserTransitionRequest {
	this := UserTransitionRequest{}
	this.Channel = channel
	this.ReasonCode = reasonCode
	this.Status = status
	this.UserToken = userToken
	return &this
}

// NewUserTransitionRequestWithDefaults instantiates a new UserTransitionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserTransitionRequestWithDefaults() *UserTransitionRequest {
	this := UserTransitionRequest{}
	return &this
}

// GetChannel returns the Channel field value
func (o *UserTransitionRequest) GetChannel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Channel
}

// GetChannelOk returns a tuple with the Channel field value
// and a boolean to check if the value has been set.
func (o *UserTransitionRequest) GetChannelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Channel, true
}

// SetChannel sets field value
func (o *UserTransitionRequest) SetChannel(v string) {
	o.Channel = v
}

// GetIdempotentHash returns the IdempotentHash field value if set, zero value otherwise.
func (o *UserTransitionRequest) GetIdempotentHash() string {
	if o == nil || IsNil(o.IdempotentHash) {
		var ret string
		return ret
	}
	return *o.IdempotentHash
}

// GetIdempotentHashOk returns a tuple with the IdempotentHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserTransitionRequest) GetIdempotentHashOk() (*string, bool) {
	if o == nil || IsNil(o.IdempotentHash) {
		return nil, false
	}
	return o.IdempotentHash, true
}

// HasIdempotentHash returns a boolean if a field has been set.
func (o *UserTransitionRequest) HasIdempotentHash() bool {
	if o != nil && !IsNil(o.IdempotentHash) {
		return true
	}

	return false
}

// SetIdempotentHash gets a reference to the given string and assigns it to the IdempotentHash field.
func (o *UserTransitionRequest) SetIdempotentHash(v string) {
	o.IdempotentHash = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *UserTransitionRequest) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserTransitionRequest) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *UserTransitionRequest) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *UserTransitionRequest) SetReason(v string) {
	o.Reason = &v
}

// GetReasonCode returns the ReasonCode field value
func (o *UserTransitionRequest) GetReasonCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReasonCode
}

// GetReasonCodeOk returns a tuple with the ReasonCode field value
// and a boolean to check if the value has been set.
func (o *UserTransitionRequest) GetReasonCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReasonCode, true
}

// SetReasonCode sets field value
func (o *UserTransitionRequest) SetReasonCode(v string) {
	o.ReasonCode = v
}

// GetStatus returns the Status field value
func (o *UserTransitionRequest) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *UserTransitionRequest) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *UserTransitionRequest) SetStatus(v string) {
	o.Status = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *UserTransitionRequest) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserTransitionRequest) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *UserTransitionRequest) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *UserTransitionRequest) SetToken(v string) {
	o.Token = &v
}

// GetUserToken returns the UserToken field value
func (o *UserTransitionRequest) GetUserToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserToken
}

// GetUserTokenOk returns a tuple with the UserToken field value
// and a boolean to check if the value has been set.
func (o *UserTransitionRequest) GetUserTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserToken, true
}

// SetUserToken sets field value
func (o *UserTransitionRequest) SetUserToken(v string) {
	o.UserToken = v
}

func (o UserTransitionRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserTransitionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["channel"] = o.Channel
	if !IsNil(o.IdempotentHash) {
		toSerialize["idempotentHash"] = o.IdempotentHash
	}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	toSerialize["reason_code"] = o.ReasonCode
	toSerialize["status"] = o.Status
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	toSerialize["user_token"] = o.UserToken
	return toSerialize, nil
}

type NullableUserTransitionRequest struct {
	value *UserTransitionRequest
	isSet bool
}

func (v NullableUserTransitionRequest) Get() *UserTransitionRequest {
	return v.value
}

func (v *NullableUserTransitionRequest) Set(val *UserTransitionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUserTransitionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUserTransitionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserTransitionRequest(val *UserTransitionRequest) *NullableUserTransitionRequest {
	return &NullableUserTransitionRequest{value: val, isSet: true}
}

func (v NullableUserTransitionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserTransitionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


