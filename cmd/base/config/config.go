package config

import (
	"kafka/cmd/kafka/config"
	"kafka/pkg/common/db"
)

type Config struct {
	DBConfig             db.Config                   `split_words:"true" required:"true"`
	KafkaConfig          config.KafkaConfig          `split_words:"true" required:"true"`
	SchemaRegistryConfig config.SchemaRegistryConfig `split_words:"true" required:"true"`
}
