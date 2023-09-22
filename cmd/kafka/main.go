package main

import (
	"fmt"
	"kafka/cmd/base/config"
	"kafka/pkg/common/db"

	con "kafka/pkg/consumer"
	pro "kafka/pkg/producer"
	"sync"

	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg := config.Config{}
	envconfig.MustProcess("KAFKA_CONFIG", &cfg)
	logrus.Infof("Kafka Config: %+v", cfg)

	_db, err := db.GetDatabase(&cfg.DBConfig)
	if err != nil {
		panic(err)
	}

	kafkaCfg := cfg.KafkaConfig.GetBaseConfigMap()
	fmt.Printf("kafkaCfg : %+v\n", kafkaCfg)

	topic := "topic_send_data_to_database"

	// create consumer
	consumer, err := con.CreateConsumer(kafkaCfg)
	if err != nil {
		panic(err)
	}

	cfgProducer := pro.ConfigProducer{
		KafkaCfgMap:       kafkaCfg,
		KafkaCfg:          &cfg.KafkaConfig,
		SchemaRegistryCfg: &cfg.SchemaRegistryConfig,
	}

	// create producer and serializer
	srProducer, err := pro.CreateProducerAndSerializer(cfgProducer)
	if err != nil {
		panic(err)
	}

	logrus.Info("============= Processing Run ================")
	var wg sync.WaitGroup
	wg.Add(1)
	// process consumer
	go func() {
		logrus.Info("============= Processing Consumer ================")
		if err := con.ProcessConsumer(consumer, topic, _db); err != nil {
			panic(err)
		}
		logrus.Info("============= Done Run Consumer ================")
		wg.Done()
	}()

	wg.Add(1)
	// process producer
	go func() {
		logrus.Info("============= Processing Producer ================")
		pro.ProcessProducer(srProducer, topic)
		logrus.Info("============= Done Run Producer ================")
		wg.Done()
	}()

	wg.Wait()
}
