package work

import (
	"common/utils/log"
	"common/utils/amqp"
	"common/utils/config"
	"common/utils/json"
	"common/utils/time"
	"common/utils/convert"
	"common/model"
	"common/utils/mgodb"
	rec "work/reckon"
	"errors"
	"fmt"
)

type reckon struct {
	Base
}

func newReckon() Work {
	return &reckon{
		Base{
			name:       "reckon",
			usage:      "Bigame reckon server",
			reboot:     true,
			rebootTime: 1,
		},
	}
}

func (r *reckon) Run() {
	mq := amqp.NewAmqp()
	mq.Url = config.C.GetStringByDefault("mq_url")
	mq.Qd.Name = r.GetName()
	msgs := mq.Receive()
	defer mq.Close()

	log.Info("[*] Waiting for messages. To exit press CTRL+C")
	for d := range msgs {
		go func() {
			log.Info(fmt.Sprintf("Received a message: %s", d.Body))
			startTime := time.NowUnixMilli()
			msgs := &model.Msgs{}
			json.Unmarshal(d.Body, &msgs)
			r := rec.NewReckonHandler()
			defer func() {
				log.TimeConsuming(startTime, "[handicap "+convert.ToStr(msgs.Han_id)+"] is over")
				d.Ack(false)
				mgodb.Close()
			}()
			if msgs.Han_id > 0 {
				r.Run(msgs)
			} else {
				log.Error(errors.New("[handicap "+convert.ToStr(msgs.Han_id)+"] is not gt zero error"), "")
			}
		}()
	}
}
