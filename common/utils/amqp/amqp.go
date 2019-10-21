package amqp

import (
	"common/utils/log"
	"errors"
	rabbitmq "github.com/streadway/amqp"
)

type Amqp struct {
	Url  string
	conn *rabbitmq.Connection
	ch   *rabbitmq.Channel
	Qd   QueueDeclare
	C    Consume
}

type QueueDeclare struct {
	Name      string
	Durable   bool
	AutoAck   bool
	Exclusive bool
	NoWait    bool
	Arguments map[string]interface{}
}

type Consume struct {
	Consumer  string
	AutoAck   bool
	Exclusive bool
	NoLocal   bool
	NoWait    bool
	Arguments map[string]interface{}
}

func NewAmqp() *Amqp {
	return &Amqp{}
}

func (amqp *Amqp) Close() {
	if amqp.conn != nil {
		amqp.conn.Close()
	}
}

func (amqb *Amqp) initMq() {
	var err error
	if amqb.Url == "" {
		log.Error2Exit(errors.New("amqp url is empty"), "error")
	}
	amqb.conn, err = rabbitmq.Dial(amqb.Url)
	log.Error2Exit(err, "rabbitmq.Dial error")
	amqb.ch, err = amqb.conn.Channel()
	log.Error2Exit(err, "conn.Channel error")
}

func (amqp *Amqp) Receive() <-chan rabbitmq.Delivery {
	amqp.initMq()

	q, err := amqp.ch.QueueDeclare(
		amqp.Qd.Name,      // name
		amqp.Qd.Durable,   // durable
		amqp.Qd.AutoAck,   // delete when usused
		amqp.Qd.Exclusive, // exclusive
		amqp.Qd.NoWait,    // no-wait
		amqp.Qd.Arguments, // arguments
	)
	log.Error2Exit(err, "ch.QueueDeclare error")

	msgs, err := amqp.ch.Consume(
		q.Name,           // queue
		amqp.C.Consumer,  // consumer
		amqp.C.AutoAck,   // auto-ack
		amqp.C.Exclusive, // exclusive
		amqp.C.NoLocal,   // no-local
		amqp.C.NoWait,    // no-wait
		amqp.C.Arguments, // args
	)
	log.Error2Exit(err, "ch.Consume error")

	return msgs
}
