package kafka

import (
	"test/configs"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/rs/zerolog"
)

type (
	kafkaDatabase struct {
		Consumer *kafka.Consumer
		Producer *kafka.Producer
		Logs     zerolog.Logger
	}
	KafkaDatabase interface {
		GetConsumer() *kafka.Consumer
		GetProducer() *kafka.Producer
	}
)

func InitializeKafkaDatabase(consumer *kafka.Consumer, producer *kafka.Producer, log zerolog.Logger) KafkaDatabase {
	return &kafkaDatabase{
		Consumer: consumer,
		Producer: producer,
		Logs:     log,
	}
}

func ConnectKafkaReader(log zerolog.Logger) *kafka.Consumer {
	conf := configs.Cfg.Database.Kafka
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": conf.Address,
		"group.id":          conf.GroupID,
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		log.Error().Err(err).Msg("Error creating kafka consumer")
		panic(err)
	}
	log.Info().Msg("Kafka reader connected successfully")
	return c
}

func ConnectKafkaWriter(log zerolog.Logger) *kafka.Producer {
	conf := configs.Cfg.Database.Kafka
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": conf.Address,
	})
	if err != nil {
		log.Error().Err(err).Msg("Error creating kafka producer")
		panic(err)
	}
	log.Info().Msg("Kafka writer connected successfully")
	return p
}

func (k *kafkaDatabase) GetConsumer() *kafka.Consumer {
	return k.Consumer
}

func (k *kafkaDatabase) GetProducer() *kafka.Producer {
	return k.Producer
}
