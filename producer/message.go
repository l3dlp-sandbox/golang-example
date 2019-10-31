package producer

import (
	"fmt"
	"github.com/Shopify/sarama"


	"gitlab.mytaxi.lk/pickme/go-util/log"
)

func Message(genEventBytes []byte, topic string) {

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(genEventBytes),
	}

	//log.Println("Producing kafka message: ", msg.Value)

	partition, offset, err := Producer.SendMessage(msg)
	if err != nil {
		log.Error("ERROR sending message to kafka: ", err)
		log.Error(log.WithPrefix("driver-ledger.producer.CreateKafkaProducer", fmt.Sprintf("Error sending to kafka- %v", err)))
		return
	}
	log.Info(log.WithPrefix("driver-ledger.producer.CreateKafkaProducer", fmt.Sprintf("Message is stored in topic %v,partition %v, offset%v", topic, partition, offset)))

}