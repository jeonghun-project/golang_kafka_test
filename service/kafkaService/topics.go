package kafkaService

import (
	"fmt"
	"github.com/segmentio/kafka-go"
)

//key collection
const (
	MsgKeyInvalidated = "invalidated"
	MsgKeyUpdated     = "updated"
)

const (
	KeyTopic = "key"
)

func subscribeTopic(topic, groupId string, cb func(kafka.Message)) {
	//subscribe to kafkaService client
	client := getClient()
	//onresult
	//parse result
	client.subscribe(topic, groupId, func(msg kafka.Message) {
		cb(msg)
		fmt.Println("Some thing do it : ", KeyTopic, string(msg.Value))
	})
	//callback("keyID")
}
func PublishToKeyInvalidated(KeyTopic string, keyID string) {
	msg := []byte(keyID)
	client := getClient()
	err := client.publish(KeyTopic, MsgKeyInvalidated, msg)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
	//client.Publish(msg)
	//client.publish(dfsdfsdf)
}
