package tools

import "encoding/json"

func JsonToMap(json_byte []byte) map[int]int {
	result := map[int]int{}
	json.Unmarshal(json_byte, &result)
	return result
}
