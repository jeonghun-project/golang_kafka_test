package kafkaService

import (
	"fmt"
	"github.com/segmentio/kafka-go"
)

//key collection
const (
	KeyInvalidated = "key/invalidated"
)

func subscribeKeyInvalidated(groupId string, cb func([]byte)) {
	//subscribe to kafkaService client
	client := getClient()
	//onresult
	//parse result
	client.subscribe(KeyInvalidated, groupId, func(msg kafka.Message) {
		fmt.Println("Some thing do it : ", KeyInvalidated, string(msg.Value))
		cb(msg.Value)
	})
	//callback("keyID")
}
func PublishToKeyInvalidated(exchange string, keyID string) {
	msg := []byte(keyID)
	client := getClient()
	err := client.publish(exchange, KeyInvalidated, msg)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
	//client.Publish(msg)
	//client.publish(dfsdfsdf)
}
