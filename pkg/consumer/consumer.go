package consumer

import (
	"fmt"
	"kafka/pkg/common/db"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/sirupsen/logrus"
)

func CreateConsumer(cfg *kafka.ConfigMap) (*kafka.Consumer, error) {
	cfg.SetKey("session.timeout.ms", "6000")
	cfg.SetKey("group.id", "send-data")
	cfg.SetKey("auto.offset.reset", "earliest")
	cfg.SetKey("enable.auto.offset.store", false)
	c, err := kafka.NewConsumer(cfg)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func ProcessConsumer(c *kafka.Consumer, topic string, _db *db.DB) error {
	// pgdb := dbf.UseDatabase(_db)
	if err := c.SubscribeTopics([]string{topic}, nil); err != nil {
		logrus.Warnf("error subscribe topic: %v", err)
		return err
	}

	// Set up a channel for handling Ctrl-C, etc
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	// Process messages
	run := true
	for run {
		select {
		case sig := <-sigchan:
			fmt.Printf("Caught signal %v: terminating\n", sig)
			run = false
		default:
			ev, err := c.ReadMessage(100 * time.Millisecond)
			if err != nil {
				// Errors are informational and automatically handled by the consumer
				continue
			}
			fmt.Println("=====================================")
			fmt.Printf("ev : %+v\n", ev)
			fmt.Printf("Consumed event from topic %s: key = %-10s value = %s\n",
				*ev.TopicPartition.Topic, string(ev.Key), string(ev.Value))
			fmt.Println("=====================================")

			if ev.Value != nil {
				logrus.Infof("=> value not nil : %+v", string(ev.Value))
				// insert data to database
				// if err := pgdb.InsertValueToDatabase(context.Background(), types.Dtests{
				// 	Topic: *ev.TopicPartition.Topic,
				// 	Key:   string(ev.Key),
				// 	Value: string(ev.Value),
				// 	Time:  ev.Timestamp,
				// }); err != nil {
				// 	logrus.Warnf("error insert data to database: %v", err)
				// 	panic(err)
				// }
			}
		}
	}

	c.Close()

	return nil
}
