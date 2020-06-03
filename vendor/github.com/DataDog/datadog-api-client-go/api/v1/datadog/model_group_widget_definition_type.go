/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
	"fmt"
)

// GroupWidgetDefinitionType Type of the group widget.
type GroupWidgetDefinitionType string

// List of GroupWidgetDefinitionType
const (
	GROUPWIDGETDEFINITIONTYPE_GROUP GroupWidgetDefinitionType = "group"
)

func (v *GroupWidgetDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GroupWidgetDefinitionType(value)
	for _, existing := range []GroupWidgetDefinitionType{"group"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GroupWidgetDefinitionType", *v)
}

// Ptr returns reference to GroupWidgetDefinitionType value
func (v GroupWidgetDefinitionType) Ptr() *GroupWidgetDefinitionType {
	return &v
}

type NullableGroupWidgetDefinitionType struct {
	value *GroupWidgetDefinitionType
	isSet bool
}

func (v NullableGroupWidgetDefinitionType) Get() *GroupWidgetDefinitionType {
	return v.value
}

func (v *NullableGroupWidgetDefinitionType) Set(val *GroupWidgetDefinitionType) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupWidgetDefinitionType) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupWidgetDefinitionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupWidgetDefinitionType(val *GroupWidgetDefinitionType) *NullableGroupWidgetDefinitionType {
	return &NullableGroupWidgetDefinitionType{value: val, isSet: true}
}

func (v NullableGroupWidgetDefinitionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupWidgetDefinitionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
