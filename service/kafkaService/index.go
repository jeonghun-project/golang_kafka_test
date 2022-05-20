package kafkaService

import (
	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"kafka/config"
	"kafka/model/key"
)

const (
	ServerAPIGroupId = "kafka_APIGroup"
)

var topics = map[string][]string{
	config.ServerTypeAPI:      []string{KeyTopic},
	config.ServerTypeData:     []string{KeyTopic},
	config.ServerTypeAnalysis: []string{KeyTopic},
}

func Connect() {
	subscribeFunc := map[string]func(kafka.Message){
		KeyTopic: keyTopicCallback,
	}

	switch config.ServerType {
	case config.ServerTypeAPI:
		for i := range topics[config.ServerTypeAPI] {
			topic := topics[config.ServerTypeAPI][i]
			callback := subscribeFunc[topic]
			go subscribeTopic(topic, ServerAPIGroupId, callback)
		}
	}
}

func keyTopicCallback(msg kafka.Message) {
	switch string(msg.Key) {
	case MsgKeyInvalidated:
		v, _ := primitive.ObjectIDFromHex(string(msg.Value))
		key.DeleteByKeyId(v)
	case MsgKeyUpdated:
		v, _ := primitive.ObjectIDFromHex(string(msg.Value))
		key.UpdateByKeyId(v)
	}
}
