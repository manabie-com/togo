package config

import (
	"testing"
)

type sampleApplicationConfig struct {
	SampleConfig sampleConfig `json:"sample_config"`
}

type sampleConfig struct {
	ConfigTypeInt            int                `json:"config_type_int"`
	ConfigTypeString         string             `json:"config_type_string"`
	ConfigTypeBool           bool               `json:"config_type_bool"`
	ConfigTypeIntUnmapped    int                `json:"config_type_int_unmaped"`
	ConfigTypeStringUnmapped string             `json:"config_type_string_unmaped"`
	ConfigTypeBoolUnmapped   bool               `json:"config_type_bool_unmaped"`
	NestedConfig             sampleNestedConfig `json:"nested_config"`
}

type sampleNestedConfig struct {
	ConfigNestedTypeInt            int    `json:"config_nested_type_int"`
	ConfigNestedTypeString         string `json:"config_nested_type_string"`
	ConfigNestedTypeBool           bool   `json:"config_nested_type_bool"`
	ConfigNestedTypeIntUnmapped    int    `json:"config_nested_type_int_unmaped"`
	ConfigNestedTypeStringUnmapped string `json:"config_nested_type_string_unmaped"`
	ConfigNestedTypeBoolUnmapped   bool   `json:"config_nested_type_bool_unmaped"`
}

func TestLoadConfig(t *testing.T) {
	var got sampleApplicationConfig
	Load("./sample.json", &got)
	want := &sampleApplicationConfig{
		SampleConfig: sampleConfig{
			ConfigTypeInt:    3,
			ConfigTypeString: "some string",
			ConfigTypeBool:   true,
			NestedConfig: sampleNestedConfig{
				ConfigNestedTypeInt:    33,
				ConfigNestedTypeString: "some string some string",
				ConfigNestedTypeBool:   true,
			},
		},
	}
	if got.SampleConfig.ConfigTypeInt != want.SampleConfig.ConfigTypeInt {
		t.Errorf("ConfigTypeInt is not mapped")
	}
	if got.SampleConfig.ConfigTypeString != want.SampleConfig.ConfigTypeString {
		t.Errorf("ConfigTypeString is not mapped")
	}
	if got.SampleConfig.ConfigTypeBool != want.SampleConfig.ConfigTypeBool {
		t.Errorf("ConfigTypeBool is not mapped")
	}
	if got.SampleConfig.ConfigTypeIntUnmapped != 0 {
		t.Errorf("ConfigTypeIntUnmapped is unmapped but does not have default value (0)")
	}
	if got.SampleConfig.ConfigTypeStringUnmapped != "" {
		t.Errorf(`ConfigTypeStringUnmapped is unmapped but does not have default value ("")`)
	}
	if got.SampleConfig.ConfigTypeBoolUnmapped != false {
		t.Errorf("ConfigTypeBoolUnmapped is unmapped but does not have default value (false)")
	}

	if got.SampleConfig.NestedConfig.ConfigNestedTypeInt != want.SampleConfig.NestedConfig.ConfigNestedTypeInt {
		t.Errorf("NestedConfig.ConfigNestedTypeInt is not mapped")
	}
	if got.SampleConfig.NestedConfig.ConfigNestedTypeString != want.SampleConfig.NestedConfig.ConfigNestedTypeString {
		t.Errorf("NestedConfig.ConfigNestedTypeString is not mapped")
	}
	if got.SampleConfig.NestedConfig.ConfigNestedTypeBool != want.SampleConfig.NestedConfig.ConfigNestedTypeBool {
		t.Errorf("NestedConfig.ConfigNestedTypeBool is not mapped")
	}
	if got.SampleConfig.NestedConfig.ConfigNestedTypeIntUnmapped != 0 {
		t.Errorf("NestedConfig.ConfigNestedTypeIntUnmapped is unmapped but does not have default value (0)")
	}
	if got.SampleConfig.NestedConfig.ConfigNestedTypeStringUnmapped != "" {
		t.Errorf(`NestedConfig.ConfigNestedTypeStringUnmapped is unmapped but does not have default value ("")`)
	}
	if got.SampleConfig.NestedConfig.ConfigNestedTypeBoolUnmapped != false {
		t.Errorf("NestedConfig.ConfigNestedTypeBoolUnmapped is unmapped but does not have default value (false)")
	}
}
