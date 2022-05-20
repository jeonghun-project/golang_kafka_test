package kafkaService

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"kafka/config"
	"kafka/model/key"
)

const (
	ServerAPIGroupId = "kafka_APIGroup"
)

var topics = map[string][]string{
	config.ServerTypeAPI:      []string{KeyInvalidated},
	config.ServerTypeData:     []string{KeyInvalidated},
	config.ServerTypeAnalysis: []string{KeyInvalidated},
}

func Connect() {
	subscribeFunc := map[string]func(string, func([]byte)){
		KeyInvalidated: subscribeKeyInvalidated,
	}

	switch config.ServerType {
	case config.ServerTypeAPI:
		for i := range topics[config.ServerTypeAPI] {
			topic := topics[config.ServerTypeAPI][i]
			subscribe := subscribeFunc[topic]
			go subscribe(ServerAPIGroupId, keyInvalidatedCallback)
		}
	}
}

func keyInvalidatedCallback(keyId []byte) {
	v, _ := primitive.ObjectIDFromHex(string(keyId))
	key.DeleteByKeyId(v)
}
