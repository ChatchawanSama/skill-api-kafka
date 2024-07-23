package skill

import (
	"fmt"
	"log"
	"os"
	"strings"

	_ "net/http/pprof"

	"github.com/IBM/sarama"
)

// Sarama configuration options
var (
	brokerURLS string
	topic      string
)

func init() {
	brokerURLS = os.Getenv("BROKER")
	topic = os.Getenv("TOPIC")

	fmt.Printf("BROKER: %s\n", brokerURLS)
	fmt.Printf("TOPIC: %s\n", topic)

	if len(topic) == 0 {
		panic("no topic given to be consumed, please set the -topic flag")
	}
}

func produceMessage(message, action string) {
	fmt.Println("--------------------------------------------------------->")
	config := configProducer()

	fmt.Println("Broker urls", brokerURLS)
	producer, err := sarama.NewSyncProducer(strings.Split(brokerURLS, ","), config)
	if err != nil {
		log.Fatalln(err)
	}

	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	fmt.Println("<---------------------------------------------------------")
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(action),
		Value: sarama.StringEncoder(message)}

	partition, offset, err := producer.SendMessage(msg)
	fmt.Println("Bat Man: ---------------------------------------------------------")
	if err != nil {
		log.Printf("FAILED to send message: %s\n", err)
	} else {
		log.Printf("> message sent to partition %d at offset %d\n", partition, offset)
	}
}

func configProducer() *sarama.Config {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	// config.Producer.Partitioner = sarama.NewRoundRobinPartitioner
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll

	return config
}
