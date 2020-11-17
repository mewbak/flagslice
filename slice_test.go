package flagslice

import (
	"reflect"
	"testing"
	"time"
)

func TestSlice(t *testing.T) {
	tests := []struct {
		name   string
		args   []string
		slice  interface{}
		expect interface{}
	}{
		{
			name:   "string",
			args:   []string{"foo", "bar", "baz"},
			expect: []string{"foo", "bar", "baz"},
		},
		{
			name:   "int",
			args:   []string{"1", "42", "999"},
			expect: []int{1, 42, 999},
		},
		{
			name:   "duration",
			args:   []string{"2s", "10m", "0"},
			expect: []time.Duration{2 * time.Second, 10 * time.Minute, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			slice := reflect.MakeSlice(reflect.TypeOf(tt.expect), 0, 0)
			ptr := reflect.New(slice.Type())
			ptr.Elem().Set(slice)
			v := Slice(ptr.Interface())
			for _, a := range tt.args {
				if err := v.Set(a); err != nil {
					t.Errorf("arg %q: %v", a, err)
				}
			}
			if reflect.DeepEqual(tt.slice, tt.expect) {
				t.Errorf("expected %v, got %v", tt.expect, tt.slice)
			}
		})
	}

}
