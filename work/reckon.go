package work

import (
	"errors"
	"fmt"
	"gowork/common/model"
	"gowork/common/utils/amqp"
	"gowork/common/utils/config"
	"gowork/common/utils/convert"
	"gowork/common/utils/json"
	"gowork/common/utils/log"
	"gowork/common/utils/mgodb"
	"gowork/common/utils/time"
	rec "gowork/work/reckon"
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
