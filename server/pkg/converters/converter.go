package converters

import "encoding/json"

func MapToString(mapData map[string]string) string {
	str, _ := json.Marshal(mapData)

	return string(str)

}

func StringToMap(stringData string) map[string]string {
	var out map[string]string

	_ = json.Unmarshal([]byte(stringData), &out)

	return out
}

func URLValuesToString(mapData map[string][]string) string {
	str, _ := json.Marshal(mapData)

	return string(str)

}

func StringToURLValues(stringData string) map[string][]string {
	var out map[string][]string

	_ = json.Unmarshal([]byte(stringData), &out)

	return out
}
