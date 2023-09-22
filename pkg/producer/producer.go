package producer

import (
	"fmt"
	"kafka/cmd/kafka/config"
	types_data "kafka/pkg/proto"
	"math/rand"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"

	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/confluentinc/confluent-kafka-go/v2/schemaregistry"
	"github.com/confluentinc/confluent-kafka-go/v2/schemaregistry/serde"
	"github.com/confluentinc/confluent-kafka-go/v2/schemaregistry/serde/protobuf"
)

type SRProducer struct {
	producer   *kafka.Producer
	serializer serde.Serializer
}

type ConfigProducer struct {
	KafkaCfgMap       *kafka.ConfigMap
	KafkaCfg          *config.KafkaConfig
	SchemaRegistryCfg *config.SchemaRegistryConfig
}

func CreateProducerAndSerializer(cfg ConfigProducer) (*SRProducer, error) {
	cfg.KafkaCfgMap.SetKey("enable.idempotence", true)
	cfg.KafkaCfgMap.SetKey("acks", "all")
	cfg.KafkaCfgMap.SetKey("retries", 3)
	cfg.KafkaCfgMap.SetKey("retry.backoff.ms", 100)
	cfg.KafkaCfgMap.SetKey("max.in.flight.requests.per.connection", 1)
	p, err := kafka.NewProducer(cfg.KafkaCfgMap)
	if err != nil {
		return nil, err
	}

	// client schema registry
	client, err := schemaregistry.NewClient(schemaregistry.NewConfigWithAuthentication(cfg.SchemaRegistryCfg.Url, cfg.SchemaRegistryCfg.Username, cfg.SchemaRegistryCfg.Password))
	if err != nil {
		return nil, err
	}

	// serializer
	ser, err := protobuf.NewSerializer(client, serde.ValueSerde, protobuf.NewSerializerConfig())
	if err != nil {
		return nil, err
	}

	return &SRProducer{
		producer:   p,
		serializer: ser,
	}, nil
}

func ProcessProducer(sr *SRProducer, topic string) {
	go func() {
		for e := range sr.producer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Failed to deliver message: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Produced event to topic %s: key = %-10s value = %s\n", *ev.TopicPartition.Topic, string(ev.Key), string(ev.Value))
				}
			}
		}
	}()

	users := [...]string{"key_1", "key_2", "key_3", "key_4", "key_5"}
	items := [...]string{"book", "alarm clock", "t-shirts", "gift card", "batteries"}

	for n := 0; n < 10; n++ {
		_time := time.Now()
		key := fmt.Sprintf("%v_%v", users[rand.Intn(len(users))], _time.Unix())
		value := items[rand.Intn(len(items))]
		tp := &types_data.TestData{
			Id:    fmt.Sprintf("%v_%v:%v", key, value, (_time.Unix() + int64(n+1))),
			Topic: topic,
			Key:   key,
			Value: value,
			Time:  timestamppb.New(_time),
		}
		b, err := Serializer(sr, tp)
		if err != nil {
			logrus.Warnf("error protobuf serialize data: %+v", err)
			panic(err)
		}
		payload := b
		sr.producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Key:            []byte(key),
			Value:          payload,
			Timestamp:      _time,
		}, nil)
	}

	sr.producer.Flush(3000 * 1000)
	defer sr.producer.Close()
	defer sr.serializer.Close()
}

func Serializer(sr *SRProducer, tp *types_data.TestData) ([]byte, error) {
	return sr.serializer.Serialize(tp.Topic, tp)
}
