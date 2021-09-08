package convert

import (
	"testing"
)

func TestJson2Yaml(t *testing.T) {
	type args struct {
		j string
	}
	tests := []struct {
		name  string
		args  args
		wantY string
	}{
		// TODO: Add test cases.
		{
			name: "test-1",
			args: args{
				j: `{
    "test": {
        "hello": {
            "world": {
                "name": "demo-be",
                "namespace": "default"
            }
        }
    }
}`,
			},
			wantY: `test:
    hello:
        world:
            name: demo-be
            namespace: default
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotY := Json2Yaml(tt.args.j); gotY != tt.wantY {
				t.Errorf("Json2Yaml() = %v, want %v", gotY, tt.wantY)
			}
		})
	}
}
