package faceit_cc

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	kafka "github.com/segmentio/kafka-go"
)

type Notificater interface {
	CreateTopic(topic string) error
	StartWriter()
	Publish(input string) error
	Initialize() error
}

type KafkaInstance struct {
	Ctx     context.Context
	Brokers []string
	Topic   string
	Writer  *kafka.Writer
}

func NewKafkaInstance() KafkaInstance {
	return KafkaInstance{}
}

// CreateTopic creates a topic in kafka environment
func (k *KafkaInstance) CreateTopic(topic string) error {
	_, err := kafka.DialLeader(k.Ctx, "tcp", k.Brokers[0], topic, 0)
	if err != nil {
		return err
	}
	return nil
}

// StartWriter intializes the writer with the broker addresses, and the topic
func (k *KafkaInstance) StartWriter() {
	k.Writer = kafka.NewWriter(kafka.WriterConfig{
		Brokers: k.Brokers,
		Topic:   k.Topic,
	})
}

// Publish publis a new message in the configured topic
func (k *KafkaInstance) Publish(input string) error {
	if err := k.Writer.WriteMessages(k.Ctx, kafka.Message{
		Key:   []byte(strconv.Itoa(1)),
		Value: []byte(input),
	}); err != nil {
		return err
	}
	return nil
}

// Initialize checks for kafka status during a set of time
// when the kafka instance is ready, returns the kafka instance and a nil error
// If the initialization time is reached, returns a nil kafka instance entity and an error
func (k *KafkaInstance) Initialize() error {
	log.Println("Initializing Kafka connection ...")

	KAFKA_BROKER, err := GetSetting("KAFKA_BROKER")
	if err != nil {
		// return KafkaInstance{}, err
		return nil
	}
	KAFKA_BROKER_PORT, err := GetSetting("KAFKA_BROKER_PORT")
	if err != nil {
		// return KafkaInstance{}, err
		return nil
	}
	brokerstr := fmt.Sprintf("%s:%s", KAFKA_BROKER, KAFKA_BROKER_PORT)

	c := make(chan Result)
	waitfor := 60

	for i := 0; i < waitfor; i++ {
		time.Sleep(1 * time.Second)

		go func(broker string) {
			kafka := KafkaInstance{
				Brokers: []string{broker},
				Ctx:     context.Background(),
				Topic:   "user_events",
			}

			if err := kafka.CreateTopic(kafka.Topic); err != nil {
				log.Println("error creating topic, kafka is not ready")
				result := Result{
					Success: false,
					Error:   err,
					Kafka:   nil,
				}
				c <- result
				return
			}
			kafka.StartWriter()

			result := Result{
				Success: true,
				Error:   nil,
				Kafka:   &kafka,
			}
			c <- result
		}(brokerstr)

		select {
		case res := <-c:
			if res.Success {
				fmt.Println("***** Kafka is ready *****")
				k.Writer = res.Kafka.Writer
				k.Brokers = res.Kafka.Brokers
				k.Ctx = res.Kafka.Ctx
				k.Topic = res.Kafka.Topic
				return nil
			}
		case <-time.After(time.Duration(waitfor) * time.Second):
			fmt.Println("timeout %n", waitfor)
		}
	}
	return fmt.Errorf("error establishing kafka connection")
}

type DataEvent struct {
	User   User
	Action string
}

// CreateString returns an string from a DataEvent struct
func (de *DataEvent) CreateString() string {
	return fmt.Sprintf("User %s: Id -> %s", de.Action, de.User.Id)
}
