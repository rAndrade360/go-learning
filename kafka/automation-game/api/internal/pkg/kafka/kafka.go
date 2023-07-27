package kafka

import (
	"time"

	kfk "github.com/segmentio/kafka-go"
)

type kafka struct {
	conn *kfk.Conn
}

type KafkaPkg interface {
	Produce(data []byte, timeout ...time.Time) error
}

func NewKafkaPkg(conn *kfk.Conn) KafkaPkg {
	return &kafka{
		conn: conn,
	}
}

func (k *kafka) Produce(data []byte, timeout ...time.Time) error {
	if len(timeout) > 0 {
		err := k.conn.SetWriteDeadline(timeout[0])
		if err != nil {
			return err
		}
	}

	_, err := k.conn.WriteMessages(kfk.Message{
		Value: data,
	})
	if err != nil {
		return err
	}

	return nil
}
