package mq

import "github.com/nats-io/nats.go"

var (
	Nc  *nats.Conn
	err error
)

func Init() {
	Nc, err = nats.Connect("nats://127.0.0.1:4444")
	if err != nil {
		panic(err)
	}
}
