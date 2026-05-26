package generator

import (
	"reflect"
	"testing"
)

func TestFormatAllValues(t *testing.T) {
	cases := []struct {
		name   string
		input  []string
		want   string
	}{
		{"empty slice", []string{}, "[]"},
		{"nil slice", nil, "[]"},
		{"single value", []string{"RED"}, "[RED]"},
		{"two values", []string{"RED", "BLUE"}, "[RED, BLUE]"},
		{"multiple values", []string{"RED", "GREEN", "BLUE"}, "[RED, GREEN, BLUE]"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := formatAllValues(tc.input)
			if got != tc.want {
				t.Errorf("formatAllValues(%v) = %q, want %q", tc.input, got, tc.want)
			}
		})
	}
}

func TestProcessIndexedValues_IntType(t *testing.T) {
	values := []string{"PENDING", "RUNNING", "DONE"}

	gotValues, gotKeys := processIndexedValues(values, true)

	wantValues := []Value{
		{Key: "PENDING", Value: "0", IsInt: true},
		{Key: "RUNNING", Value: "1", IsInt: true},
		{Key: "DONE", Value: "2", IsInt: true},
	}
	wantKeys := []string{"PENDING", "RUNNING", "DONE"}

	if !reflect.DeepEqual(gotValues, wantValues) {
		t.Errorf("processIndexedValues values mismatch:\n got:  %#v\n want: %#v", gotValues, wantValues)
	}
	if !reflect.DeepEqual(gotKeys, wantKeys) {
		t.Errorf("processIndexedValues keys mismatch:\n got:  %#v\n want: %#v", gotKeys, wantKeys)
	}
}

func TestProcessIndexedValues_StringType(t *testing.T) {
	values := []string{"RED", "GREEN", "BLUE"}

	gotValues, gotKeys := processIndexedValues(values, false)

	wantValues := []Value{
		{Key: "RED", Value: `"RED"`, IsInt: false},
		{Key: "GREEN", Value: `"GREEN"`, IsInt: false},
		{Key: "BLUE", Value: `"BLUE"`, IsInt: false},
	}
	wantKeys := []string{"RED", "GREEN", "BLUE"}

	if !reflect.DeepEqual(gotValues, wantValues) {
		t.Errorf("processIndexedValues values mismatch:\n got:  %#v\n want: %#v", gotValues, wantValues)
	}
	if !reflect.DeepEqual(gotKeys, wantKeys) {
		t.Errorf("processIndexedValues keys mismatch:\n got:  %#v\n want: %#v", gotKeys, wantKeys)
	}
}

func TestProcessIndexedValues_Empty(t *testing.T) {
	gotValues, gotKeys := processIndexedValues(nil, true)

	if len(gotValues) != 0 {
		t.Errorf("expected empty values slice, got %#v", gotValues)
	}
	if len(gotKeys) != 0 {
		t.Errorf("expected empty keys slice, got %#v", gotKeys)
	}
}
