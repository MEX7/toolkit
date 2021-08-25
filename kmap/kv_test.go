package kmap

import (
	"reflect"
	"testing"
)

func TestKvUnFormat(t *testing.T) {
	type args struct {
		in string
	}
	var tests = []struct {
		name string
		args args
		want []Kv
	}{
		// TODO: Add test cases.
		{
			name: "test-1",
			args: args{
				in: "TCP:9001,TCP:9003",
			},
			want: []Kv{
				{
					Key:   "TCP",
					Value: "9001",
				},
				{
					Key:   "TCP",
					Value: "9003",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := KvUnFormat(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("KvUnFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}
