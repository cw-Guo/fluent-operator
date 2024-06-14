package plugins

import (
	"encoding/json"
	"github.com/fluent/fluent-operator/v2/apis/fluentbit/v1alpha2/plugins/params"
)

// +kubebuilder:object:generate:=true

type CommonParams struct {

	// Alias for the plugin
	Alias string `json:"alias,omitempty"`
	// RetryLimit describes how many times fluent-bit should retry to send data to a specific output. If set to false fluent-bit will try indefinetly. If set to any integer N>0 it will try at most N+1 times. Leading zeros are not allowed (values such as 007, 0150, 01 do not work). If this property is not defined fluent-bit will use the default value: 1.
	// +kubebuilder:validation:Pattern="^(((f|F)alse)|(no_limits)|(no_retries)|([1-9]+[0-9]*))$"
	RetryLimit string `json:"retryLimit,omitempty"`
}

func (c *CommonParams) AddCommonParams(kvs *params.KVs) error {
	if c.Alias != "" {
		kvs.Insert("Alias", c.Alias)
	}
	if c.RetryLimit != "" {
		kvs.Insert("Retry_Limit", c.RetryLimit)
	}
	return nil
}

// Config represents untyped YAML configuration.
// +kubebuilder:validation:Type=object
type Config struct {
	// Data holds the configuration keys and values.
	Data map[string]interface{} `json:"-"`
}

// MarshalJSON implements the Marshaler interface.
func (c *Config) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Data)
}

// UnmarshalJSON implements the Unmarshaler interface.
func (c *Config) UnmarshalJSON(data []byte) error {
	var out map[string]interface{}
	err := json.Unmarshal(data, &out)
	if err != nil {
		return err
	}
	c.Data = out
	return nil
}

// MarshalYAML implements the yaml.Marshaler interface.
func (c *Config) MarshalYAML() (interface{}, error) {
	return c.Data, nil
}

// DeepCopyInto is an ~autogenerated~ deepcopy function, copying the receiver, writing into out. in must be non-nil.
// This exists here to work around https://github.com/kubernetes/code-generator/issues/50
func (c *Config) DeepCopyInto(out *Config) {
	bytes, err := json.Marshal(c.Data)
	if err != nil {
		// we assume that it marshals cleanly because otherwise the resource would not have been
		// created in the API server
		panic(err)
	}
	var clone map[string]interface{}
	err = json.Unmarshal(bytes, &clone)
	if err != nil {
		// we assume again optimistically because we just marshalled that the round trip works as well
		panic(err)
	}
	out.Data = clone
}
