package dynamic_json

import (
	"reflect"
	"testing"
)

const (
	raw1 =`
{
    "a": "1",
    "b": {
        "c":"2",
        "d": "3",
        "b": {
            "c":"2",
            "d": "3"
        }
    },
    "e":"4"
}
`)
	var want1 = map[string]string{
		"a": "1",
		"b.c": "2",
		"b.d": "3",
		"b.b.c": "2",
		"b.b.d": "3",
		"e": "4",
	}


func TestParse(t *testing.T) {
	type args struct {
		raw string
	}
	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		{"", args{raw: raw1}, want1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Parse(tt.args.raw); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}