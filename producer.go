package main

import (
	"kafka/service/kafkaService"
)

const (
	topic         = "test"
	brokerAddress = "localhost:9092"
)

//var mechanism sasl.Mechanism
//var dialer *kafka.Dialer

func init() {
	//mechanism, err := scram.Mechanism(scram.SHA512, "username", "password")
	//if err != nil {
	//	panic(err)
	//}
	//
	//dialer = &kafka.Dialer{
	//	Timeout:   10 * time.Second,
	//	DualStack: true,
	//	//SASLMechanism: mechanism,
	//	ClientID: "1",
	//}
}

func main() {
	//ctx := context.Background()

	//createTopic(ctx)
	produce()
	//consume(ctx)
}

//func createTopic(context context.Context) {
//	conn, err := kafkaService.DialLeader(context, "tcp", "localhost:9093", topic, 0)
//	if err != nil {
//		panic(err.Error())
//	}
//	defer conn.Close()
//}

func produce() {
	kafkaService.PublishToKeyInvalidated("binance", "something KeyId")
	kafkaService.PublishToKeyInvalidated("mexc", "something KeyId")
	//defer ctx.Done()
	//i := 0
	//writer := &kafkaService.Writer{
	//	Addr:        kafkaService.TCP(brokerAddress),
	//	Compression: kafkaService.Gzip,
	//	Transport: &kafkaService.Transport{
	//		Dial:     dialer.DialFunc,
	//		ClientID: "1",
	//		//TLS: &tls.Config{
	//		//	MinVersion: tls.VersionTLS12,
	//		//},
	//		//SASL: mechanism,
	//	},
	//	Balancer: &kafkaService.LeastBytes{},
	//}
	//
	//for {
	//	err := writer.WriteMessages(ctx, kafkaService.Message{
	//		Key:   []byte(strconv.Itoa(i)),
	//		Value: []byte("this is message" + strconv.Itoa(i)),
	//	})
	//
	//	if err != nil {
	//		panic("could not write message " + err.Error())
	//	}
	//
	//	fmt.Println("writes:", i)
	//	i++
	//
	//	time.Sleep(time.Millisecond * 10)
	//}

}
