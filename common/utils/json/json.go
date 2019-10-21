package json

import (
	j "encoding/json"
	"gowork/common/utils/log"
)

func Unmarshal(data []byte, v interface{}) {
	err := j.Unmarshal(data, v)
	log.Error(err, "json.Unmarshal error")
}
