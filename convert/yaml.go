package convert

import (
	"encoding/json"

	"gopkg.in/yaml.v3"
)

func Json2Yaml(j string) string {
	var out interface{}
	_ = json.Unmarshal([]byte(j), &out)
	y, _ := yaml.Marshal(out)
	return string(y)
}
