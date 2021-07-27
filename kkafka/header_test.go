package kkafka

import (
	"reflect"
	"testing"

	"github.com/segmentio/kafka-go"
)

func TestHeadersBatchAdd(t *testing.T) {
	type args struct {
		headers        []kafka.Header
		needAddHeaders []kafka.Header
	}
	tests := []struct {
		name string
		args args
		want []kafka.Header
	}{
		// TODO: Add test cases.
		{
			name: "test-1",
			args: args{
				headers: []kafka.Header{{
					Key:   "A",
					Value: nil,
				}, {
					Key:   "B",
					Value: nil,
				}, {
					Key:   "C",
					Value: nil,
				}},
				needAddHeaders: []kafka.Header{{
					Key:   "A",
					Value: nil,
				}, {
					Key:   "B",
					Value: nil,
				}, {
					Key:   "D",
					Value: nil,
				}},
			},
			want: []kafka.Header{{
				Key:   "A",
				Value: nil,
			}, {
				Key:   "B",
				Value: nil,
			}, {
				Key:   "D",
				Value: nil,
			}, {
				Key:   "C",
				Value: nil,
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HeadersBatchAdd(tt.args.headers, tt.args.needAddHeaders); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HeadersBatchAdd() = %v, want %v", got, tt.want)
			}
		})
	}
}
