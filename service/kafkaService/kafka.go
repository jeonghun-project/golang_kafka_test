package kafkaService

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"sync"
	"time"
)

const (
	brokerAddress = "localhost:9092"
	clientId      = "kafkaService-client"
)

var dialer *kafka.Dialer

type Client struct {
	sync.Mutex
	ctx   context.Context
	close context.CancelFunc
}

func init() {
	dialer = &kafka.Dialer{
		Timeout:   10 * time.Second,
		DualStack: true,
		//SASLMechanism: mechanism,
	}
}

func getClient() *Client {
	return &Client{}
}

func (c *Client) setContext() {
	c.ctx, c.close = context.WithCancel(context.Background())
}

func setWriter(topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:        kafka.TCP(brokerAddress),
		Compression: kafka.Gzip,
		Topic:       topic,
		Balancer:    &kafka.RoundRobin{},
		Transport: &kafka.Transport{
			Dial: dialer.DialFunc,
		},
	}
}

func (c *Client) publish(topic, key string, value []byte) error {
	c.setContext()
	c.Lock()
	defer func() {
		c.Unlock()
		c.close()

	}()
	//publish
	writer := setWriter(topic)
	fmt.Println(topic)
	err := writer.WriteMessages(c.ctx, kafka.Message{
		Key:       []byte(key),
		Value:     value,
		Partition: 3,
	})
	if err != nil {
		// some kafkaService error
		return err
	}

	if err := writer.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
	return nil
}

func (c *Client) subscribe(topic, groupId string, cb func(msg kafka.Message)) {
	c.setContext()
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        []string{brokerAddress},
		Topic:          topic,
		Dialer:         dialer,
		MinBytes:       10e3, // 10kb
		MaxBytes:       10e6, // 10mb
		MaxWait:        250 * time.Millisecond,
		CommitInterval: 250 * time.Millisecond,
	})

	//setOffset
	//r.SetOffset(c.ctx, )
	defer func() {
		if err := r.Close(); err != nil {
			log.Fatal("failed to close reader:", err)
		}
	}()

	for {
		m, err := r.ReadMessage(c.ctx)
		if err != nil {
			fmt.Println("Read messge Error: ", err)
			return
		}
		fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
		cb(m)
	}
}
