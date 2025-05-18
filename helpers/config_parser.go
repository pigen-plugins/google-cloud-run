package helpers

import (
	"encoding/json"

	"gopkg.in/yaml.v2"
)


func YamlConfigParser(in map[string] interface{}, output interface{}) error {
	out, err := yaml.Marshal(in)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(out, output)
	if err != nil {
		return err
	}
	return nil
}

// I'm using json marshal and unmarshal to convert struct to map
// Using yaml marshal and unmarshal would be cause issues with the map[string]interface{} type
// Yaml marshal and unmarshal would convert the struct to map[interface{}]interface{} which is not what we want
func StructToMap(v any) (map[string]any, error) {
	var result map[string]any
	data, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}