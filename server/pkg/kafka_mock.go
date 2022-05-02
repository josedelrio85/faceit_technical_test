package faceit_cc

import (
	"context"

	kafka "github.com/segmentio/kafka-go"
)

type KafkaMock struct {
	Ctx     context.Context
	Brokers []string
	Topic   string
	Writer  *kafka.Writer
}

func NewKafkaMock() KafkaMock {
	return KafkaMock{}
}

func (km *KafkaMock) CreateTopic(topic string) error {
	return nil
}
func (km *KafkaMock) StartWriter() {

}
func (km *KafkaMock) Publish(input string) error {
	return nil
}

func (km *KafkaMock) Initialize() error {
	return nil
}
