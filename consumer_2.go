package main

import (
	"context"
	"fmt"
	"kafka/service/kafkaService"
	"os"
	"os/signal"
	"syscall"
)

func init() {

}

func main() {
	signals := make(chan os.Signal, 1)

	signal.Notify(signals, syscall.SIGINT, syscall.SIGKILL)

	ctx, cancel := context.WithCancel(context.Background())

	// go routine for getting signals asynchronously
	go func() {
		sig := <-signals
		fmt.Println("Got signal: ", sig)
		cancel()
	}()

	kafkaService.Connect()
	defer ctx.Done()
	// TODO: kafkaService-go
	//topic := "test"
	//brokerAddress := "localhost:9092"
	//logger := log.New(os.Stdout, "kafkaService Consumer: ", 0)
	//errlogger := log.New(os.Stdout, "kafkaService Consumer Error: ", 0)
	//ctx := context.Background()
	//defer ctx.Done()
	//
	//startTime := time.Now()
	//
	//dialer := &kafkaService.Dialer{
	//	Timeout:   10 * time.Second,
	//	DualStack: true,
	//	//SASLMechanism: mechanism,
	//	ClientID: "1",
	//}
	//
	//r := kafkaService.NewReader(kafkaService.ReaderConfig{
	//	Brokers:     []string{brokerAddress},
	//	GroupID:     "stream-test",
	//	Topic:       topic,
	//	Dialer:      dialer,
	//	Logger:      logger,
	//	ErrorLogger: errlogger,
	//	//CommitInterval: time.Millisecond, // flushes commits to Kafka every second,
	//	MinBytes: 10e3,
	//	MaxBytes: 10e6,
	//	MaxWait:  250 * time.Millisecond,
	//})
	//
	//r.SetOffsetAt(ctx, startTime)
	//c := make(chan struct{}, 0)
	//
	//defer close(c)
	//for {
	//	//m, err := r.ReadMessage(ctx)
	//	m, err := r.FetchMessage(ctx)
	//	if err != nil {
	//		fmt.Println(err)
	//		break
	//	}
	//	//if m.Time.After(endTime) {
	//	//	break
	//	//}
	//	// TODO: process message
	//	fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
	//}
	//
	//if err := r.Close(); err != nil {
	//	log.Fatal("failed to close reader:", err)
	//}
	//
	//for {
	//	select {
	//	case <-c:
	//		time.Sleep(5 * time.Second)
	//		log.Println("close context")
	//		return
	//	}
	//}

}
