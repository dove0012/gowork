package json

import (
	"common/utils/log"
	j "encoding/json"
)

func Unmarshal(data []byte, v interface{}) {
	err := j.Unmarshal(data, v)
	log.Error(err, "json.Unmarshal error")
}
