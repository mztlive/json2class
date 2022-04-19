package format

import (
	"bytes"
	"encoding/json"
	"log"
)

//return 格式化后的json字符串
func FromatJSON(jsonString string) string {

	out := &bytes.Buffer{}

	err := json.Indent(out, []byte(jsonString), "", "  ")

	if err != nil {
		log.Fatalf("json format failed. reason: %s\n", err)
	}

	return out.String()
}
