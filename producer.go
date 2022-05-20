package main

import (
	"kafka/service/kafkaService"
)

func main() {
	produce()
}

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
