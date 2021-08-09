package stock

import (
	"testing"
)

func TestLogX(t *testing.T) {
	type args struct {
		bef float64
		end float64
	}
	tests := []struct {
		name  string
		args  args
		wantN int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				bef: 2,
				end: 8,
			},
			wantN: 3,
		},
		{
			name: "test2",
			args: args{
				bef: 1.025,
				end: 1.28,
			},
			wantN: 10,
		}, {
			name: "test3",
			args: args{
				bef: 1.01,
				end: 1.72,
			},
			wantN: 49,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotN := LogX(tt.args.bef, tt.args.end); gotN != tt.wantN {
				t.Errorf("LogX() = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}
