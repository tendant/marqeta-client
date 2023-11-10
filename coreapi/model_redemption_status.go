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
	"fmt"
)

// RedemptionStatus Status of the redemption.  If <</core-api/reward-redemptions#postRedemptionTransition, transitioning the redemption's status>>:  * `new_state` is the state to which you want to transition the redemption; must be `COMPLETED` or `RETURNED`.  * `initial_status` is the initial status of the redemption prior to transition.
type RedemptionStatus string

// List of RedemptionStatus
const (
	REDEMPTIONSTATUS_INITIATED RedemptionStatus = "INITIATED"
	REDEMPTIONSTATUS_COMPLETED RedemptionStatus = "COMPLETED"
	REDEMPTIONSTATUS_RETURNED RedemptionStatus = "RETURNED"
	REDEMPTIONSTATUS_PROCESSING RedemptionStatus = "PROCESSING"
	REDEMPTIONSTATUS_SUBMITTED RedemptionStatus = "SUBMITTED"
)

// All allowed values of RedemptionStatus enum
var AllowedRedemptionStatusEnumValues = []RedemptionStatus{
	"INITIATED",
	"COMPLETED",
	"RETURNED",
	"PROCESSING",
	"SUBMITTED",
}

func (v *RedemptionStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RedemptionStatus(value)
	for _, existing := range AllowedRedemptionStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RedemptionStatus", value)
}

// NewRedemptionStatusFromValue returns a pointer to a valid RedemptionStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRedemptionStatusFromValue(v string) (*RedemptionStatus, error) {
	ev := RedemptionStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RedemptionStatus: valid values are %v", v, AllowedRedemptionStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RedemptionStatus) IsValid() bool {
	for _, existing := range AllowedRedemptionStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RedemptionStatus value
func (v RedemptionStatus) Ptr() *RedemptionStatus {
	return &v
}

type NullableRedemptionStatus struct {
	value *RedemptionStatus
	isSet bool
}

func (v NullableRedemptionStatus) Get() *RedemptionStatus {
	return v.value
}

func (v *NullableRedemptionStatus) Set(val *RedemptionStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableRedemptionStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableRedemptionStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRedemptionStatus(val *RedemptionStatus) *NullableRedemptionStatus {
	return &NullableRedemptionStatus{value: val, isSet: true}
}

func (v NullableRedemptionStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRedemptionStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
