package producer

import (
	"github.com/Shopify/sarama"

	"gitlab.mytaxi.lk/pickme/go-util/log"
)

var Producer sarama.SyncProducer

func CreateKafkaProducer() {


	//brokers :=[]string{"kafka-2.c.charming-opus-833.internal:9092","kafka-0.c.charming-opus-833.internal:9092","kafka-1.c.charming-opus-833.internal:9092"}
	brokers :=[]string{"localhost:9092"}
	saramaConfig := sarama.NewConfig()
	saramaConfig.Producer.RequiredAcks = sarama.WaitForAll
	saramaConfig.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(brokers, saramaConfig)

	if err != nil {
		log.Error(log.WithPrefix("driver-ledger.producer.CreateKafkaProducer", err))
		return
	}

	Producer = producer

}