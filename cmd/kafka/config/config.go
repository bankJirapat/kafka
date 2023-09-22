package config

import "github.com/confluentinc/confluent-kafka-go/kafka"

type KafkaConfig struct {
	BootstrapServer  string `split_words:"true" required:"true"`
	RestEndpoint     string `split_words:"true" required:"true"`
	SecurityProtocol string `split_words:"true" default:"PLAINTEXT"`
	SaslMechanisms   string `split_words:"true" default:"PLAIN"`
	Username         string `required:"true"`
	Password         string `required:"true"`
}

type SchemaRegistryConfig struct {
	Url      string `required:"true"`
	Username string `required:"true"`
	Password string `required:"true"`
}

func (k *KafkaConfig) GetBaseConfigMap() *kafka.ConfigMap {
	cfg := kafka.ConfigMap{}
	cfg.SetKey("bootstrap.servers", k.BootstrapServer)
	cfg.SetKey("security.protocol", k.SecurityProtocol)
	cfg.SetKey("sasl.mechanisms", k.SaslMechanisms)
	cfg.SetKey("sasl.username", k.Username)
	cfg.SetKey("sasl.password", k.Password)
	return &cfg
}
