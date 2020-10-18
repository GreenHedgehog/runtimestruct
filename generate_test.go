package runtimestruct_test

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"

	"github.com/GreenHedgehog/runtimestruct"
)

func TestNewFromJSON(t *testing.T) {
	var testdata = `
	{
		"StringField": "asdadqdw",
		"NumberField": 12312312,
		"BoolenField": true,
		"NullField": null,
		"ArrayField": [1, "2", false],
		"ObjectField": {
			"StringField": "asdadqdw",
			"NumberField": 12312312,
			"BoolenField": true,
			"NullField": null,
			"ArrayField": [1, "2", false]
		},
		"unexportedField": null
	}`
	var want = struct {
		StringField string
		NumberField float64
		BoolenField bool
		NullField   *struct{}
		ArrayField  []interface{}
		ObjectField struct {
			StringField string
			NumberField float64
			BoolenField bool
			NullField   *struct{}
			ArrayField  []interface{}
		}
		unexportedField *struct{}
	}{
		StringField: "asdadqdw",
		NumberField: 12312312,
		BoolenField: true,
		NullField:   nil,
		ArrayField:  []interface{}{1, "2", false},
		ObjectField: struct {
			StringField string
			NumberField float64
			BoolenField bool
			NullField   *struct{}
			ArrayField  []interface{}
		}{
			StringField: "asdadqdw",
			NumberField: 12312312,
			BoolenField: true,
			NullField:   nil,
			ArrayField:  []interface{}{1, "2", false},
		},
		unexportedField: nil,
	}

	got, err := runtimestruct.NewFromJSON(strings.NewReader(testdata))
	if err != nil {
		t.Errorf("NewFromJSON() error = %v", err)
		return
	}

	equal, err := equal(got, want)
	if err != nil {
		t.Errorf("eqequalFields() error = %v", err)
		return
	}
	if !equal {
		t.Errorf("NewFromJSON() = %v, want %v", got, want)
	}
}

func equal(a, b interface{}) (bool, error) {
	var m1, m2 interface{}

	data, err := json.Marshal(a)
	if err != nil {
		return false, err
	}
	if err = json.Unmarshal(data, &m1); err != nil {
		return false, err
	}

	data, err = json.Marshal(b)
	if err != nil {
		return false, err
	}
	if err = json.Unmarshal(data, &m2); err != nil {
		return false, err
	}

	return reflect.DeepEqual(m1, m2), nil
}
