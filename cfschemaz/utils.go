package cfschemaz

import (
	"encoding/json"

	"github.com/ibrt/golang-utils/errorz"
	"github.com/ibrt/golang-utils/jsonz"
)

var (
	_ json.Unmarshaler = (*Enum)(nil)
)

// Enum describes part of the CloudFormation JSON schema for a resource.
type Enum struct {
	ss []string
	nn []json.Number
}

// IsSet returns true if the enum was set in the schema.
func (ue *Enum) IsSet() bool {
	return ue != nil
}

// CanBeString returns true if the enum is set and can be parsed as []string.
func (ue *Enum) CanBeString() bool {
	return ue != nil && len(ue.ss) > 0
}

// MaybeString returns the enum as []string if possible.
func (ue *Enum) MaybeString() []string {
	if ue != nil && len(ue.ss) > 0 {
		return ue.ss
	}

	return nil
}

// CanBeFloat64 returns true if the enum is set and can be parsed as []float64.
func (ue *Enum) CanBeFloat64() bool {
	if ue == nil || len(ue.nn) == 0 {
		return false
	}

	for _, n := range ue.nn {
		if _, err := n.Float64(); err != nil {
			return false
		}
	}

	return true
}

// MaybeFloat64 returns the enum as []float64 if possible.
func (ue *Enum) MaybeFloat64() []float64 {
	if ue == nil || len(ue.nn) == 0 {
		return nil
	}

	ff := make([]float64, 0, len(ue.nn))

	for _, n := range ue.nn {
		f, err := n.Float64()
		if err != nil {
			return nil
		}

		ff = append(ff, f)
	}

	return ff
}

// CanBeInt64 returns true if the enum is set and can be parsed as []int64.
func (ue *Enum) CanBeInt64() bool {
	if ue == nil || len(ue.nn) == 0 {
		return false
	}

	for _, n := range ue.nn {
		if _, err := n.Int64(); err != nil {
			return false
		}
	}

	return true
}

// MaybeInt64 returns the enum as []int64 if possible.
func (ue *Enum) MaybeInt64() []int64 {
	if ue == nil || len(ue.nn) == 0 {
		return nil
	}

	ii := make([]int64, 0, len(ue.nn))

	for _, n := range ue.nn {
		i, err := n.Int64()
		if err != nil {
			return nil
		}

		ii = append(ii, i)
	}

	return ii
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (ue *Enum) UnmarshalJSON(bytes []byte) error {
	if string(bytes) == "null" {
		return nil
	}

	if ss, err := jsonz.Unmarshal[[]string](bytes); err == nil {
		*ue = Enum{ss: ss}
		return nil
	}

	if nn, err := jsonz.Unmarshal[[]json.Number](bytes); err == nil {
		*ue = Enum{nn: nn}
		return nil
	}

	return errorz.Errorf("invalid enum value: '%s'", bytes)
}

func canJSONNumberBeFloat64(n json.Number) bool {
	_, err := n.Float64()
	return err == nil
}

func maybeJSONNumberPtrFloat64(n *json.Number) *float64 {
	if n != nil {
		if f, err := n.Float64(); err == nil {
			return &f
		}
	}
	return nil
}

func canJSONNumberBeInt64(n json.Number) bool {
	_, err := n.Int64()
	return err == nil
}

func maybeJSONNumberPtrInt64(n *json.Number) *int64 {
	if n != nil {
		if i, err := n.Int64(); err == nil {
			return &i
		}
	}
	return nil
}
