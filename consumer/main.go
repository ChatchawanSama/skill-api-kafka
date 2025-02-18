package main

// SIGUSR1 toggle the pause/resume consumption
import (
	"consumer/database"
	"consumer/skill"
	"context"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"

	"github.com/IBM/sarama"
)

// Sarama configuration options
var (
	brokers = os.Getenv("BROKERS")
	version = sarama.DefaultVersion.String()
	group   = os.Getenv("GROUP")
	topics  = os.Getenv("TOPICS")
	verbose = false
	oldest  = false
)

func init() {
	if len(brokers) == 0 {
		panic("no Kafka bootstrap brokers defined, please set the -brokers flag")
	}
	if len(topics) == 0 {
		panic("no topics given to be consumed, please set the -topics flag")
	}
	if len(group) == 0 {
		panic("no Kafka consumer group defined, please set the -group flag")
	}
}
func main() {
	//Set Config
	config := sarama.NewConfig()
	if oldest {
		config.Consumer.Offsets.Initial = sarama.OffsetOldest
	}
	config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategyRoundRobin()}
	consumer := Consumer{
		ready: make(chan struct{}),
	}
	client, err := sarama.NewConsumerGroup(strings.Split(brokers, ","), group, config)
	if err != nil {
		log.Panicf("new client: %v", err)
	}
	defer func() {
		if err = client.Close(); err != nil {
			log.Panicf("closing client: %v", err)
		}
	}()

	//Gracefully
	wg := &sync.WaitGroup{}
	wg.Add(1)
	ctx, gracefully := context.WithCancel(context.Background())
	go func() {
		defer wg.Done()
		for {
			if err := client.Consume(ctx, strings.Split(topics, ","), &consumer); err != nil {
				if errors.Is(err, sarama.ErrClosedConsumerGroup) {
					return
				}
				log.Panicf("consume: %v", err)
			}
			// check if context was cancelled, signaling that the consumer should stop
			if err := ctx.Err(); err != nil {
				if errors.Is(err, context.Canceled) {
					slog.Info("the consumer context has cancelled for gracefully shutting down")
					return
				}
				slog.Error(ctx.Err().Error())
				return
			}
			slog.Info("rebalancing...")
			consumer.ready = make(chan struct{})
		}
	}()
	<-consumer.ready
	slog.Info("consumer up and running...")
	sigCtx, unregistered := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer unregistered()
keepRunning:
	for {
		select {
		case <-ctx.Done():
			slog.Info("terminating: consumer context cancel")
			break keepRunning
		case <-sigCtx.Done():
			slog.Info("terminating: via signal")
			unregistered()
			break keepRunning
		}
	}
	gracefully()
	wg.Wait() // waiting for gracefully consumer stopping
	if err = client.Close(); err != nil {
		log.Panicf("closing client: %v", err)
	}
}

type Consumer struct {
	ready chan struct{}
}

func (consumer Consumer) Setup(_ sarama.ConsumerGroupSession) error {
	close(consumer.ready)
	return nil
}
func (Consumer) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (Consumer) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	// NOTE:
	// Do not move the code below to a goroutine.
	// The ConsumeClaim itself is called within a goroutine, see:
	// https://github.com/IBM/sarama/blob/main/consumer_group.go#L27-L29
	//Set Database
	// Load environment variables from .env file
	// if err := godotenv.Load(); err != nil {
	// 	log.Fatalf("Error loading .env file: %v", err)
	// }

	database.ConnectDB()
	db := database.DB
	defer database.DB.Close()
	database.CreateTable()

	repo := skill.NewSkillRepo(db)

	skillHandler := skill.NewSkillHandler(repo)

consume:
	for {
		select {
		case msg, ok := <-claim.Messages():
			if !ok {
				slog.Info("message channel was closed")
				break consume
			}
			skill.ConsumeMessage(msg, *skillHandler)
			fmt.Printf("Message topic:%q partition:%d offset:%d message:%s\n", msg.Topic, msg.Partition, msg.Offset, msg.Value)
			sess.MarkMessage(msg, "")
		// Should return when session.Context() is done.
		// If not, will raise ErrRebalanceInProgress or read tcp <ip>:<port>: i/o timeout when kafka rebalance. see:
		// https://github.com/IBM/sarama/issues/1192
		case <-sess.Context().Done():
			break consume
		}
	}
	return sess.Context().Err()
}
