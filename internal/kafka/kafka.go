package kafkaBroker

import (
	"context"
	"net"
	"strconv"

	"github.com/segmentio/kafka-go"
)

type KafkaBroker struct {
	writer *kafka.Writer
	reader *kafka.Reader
}

func NewKafkaBroker(adress, topic string) (*KafkaBroker, error) {
	err := createTopic(topic, adress)
	if err != nil {
		return nil, err
	}
	writer := &kafka.Writer{
		Addr: kafka.TCP(adress),
		Topic: topic,
	}
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{adress},
		Topic: topic,
	})
	
	return &KafkaBroker{writer: writer, reader: reader}, nil
}

func (kb *KafkaBroker) Receive() ([]byte, error) {
	msg, err := kb.reader.ReadMessage(context.Background())
	if err != nil {
		return nil, err
	}

	return msg.Value, nil
}

func (kb *KafkaBroker) Send(msg []byte) error {
	data := kafka.Message{
		Value: msg,
	}
	err := kb.writer.WriteMessages(context.Background(), data)
	if err != nil {
		return err
	}
	return nil
}

func createTopic(topic, broker string) error {
	conn, err := kafka.Dial("tcp", broker)
	if err != nil {
		return err
	}
	defer conn.Close()

	controller, err := conn.Controller()
	if err != nil {
		return err
	}

	controllerConn, err := kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	if err != nil {
		return err
	}
	defer controllerConn.Close()

	topicConfig := kafka.TopicConfig{
		Topic:             topic,
		NumPartitions:     1,
		ReplicationFactor: 1,
	}

	err = controllerConn.CreateTopics(topicConfig)
	if err != nil {
		return err
	}

	return nil
}