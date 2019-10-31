package consumer

import (
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/bsm/sarama-cluster"
	"gitlab.mytaxi.lk/pickme/go-util/log"
	"gitlab.mytaxi.lk/pickme/go/schema_registry"
	"os"
	"os/signal"
)

var (
	RegistryLedger *schema_registry.SchemaRegistry
)

type Body struct {
	Amount                  float64 `json:"amount"`
	CreatedBy               int     `json:"created_by"`
	CreatedTime             int64   `json:"created_time"`
	Description             string  `json:"description"`
	DriverID                int     `json:"driver_id"`
	TransactionCategory     int     `json:"transaction_category"`
	TransactionCategoryName string  `json:"transaction_category_name"`
	TransactionType         string  `json:"transaction_type"`
	TripID                  int     `json:"trip_id"`
}

type DriverLedgerTransactionRequest struct {
	Body      []Body `json:"body"`
	CreatedAt int64  `json:"created_at"`
	Expiry    int64  `json:"expiry"`
	ID        string `json:"id"`
	TraceInfo struct {
		ParentID int  `json:"parent_id"`
		Sampled  bool `json:"sampled"`
		SpanID   int  `json:"span_id"`
		TraceID  struct {
			High int `json:"high"`
			Low  int `json:"low"`
		} `json:"trace_id"`
	} `json:"trace_info"`
	Type    string `json:"type"`
	Version int    `json:"version"`
}

type KconsumerLedger struct {
}

func (k KconsumerLedger) InitializerLedger() {

	RegistryLedger = schema_registry.NewSchemaRegistry("http://35.184.181.97:8089/")

	RegistryLedger.Register(
		"com.pickme.events.driver.DriverLedgerTransactionRequest",
		2,
		func(data []byte) (interface{}, error) {
			e := DriverLedgerTransactionRequest{}
			err := json.Unmarshal(data, &e)
			return e, err
		},
	)
}

func (k KconsumerLedger) ConsumerLedger() {

	config := cluster.NewConfig()
	config.Consumer.Return.Errors = true
	config.Group.Return.Notifications = true
	config.Group.Mode = cluster.ConsumerModeMultiplex
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	config.Version = sarama.V1_0_0_0

	brokers := []string{"kafka-2.c.charming-opus-833.internal:9092", "kafka-0.c.charming-opus-833.internal:9092", "kafka-1.c.charming-opus-833.internal:9092"}

	//brokers := []string{"localhost:9092"}

	topics := []string{"driver_ledger_transaction_request"}

	consumer, err := cluster.NewConsumer(brokers, "test_go_java1", topics, config)
	if err != nil {
		log.Error(log.WithPrefix("driver_disconnect.consumer.InitConsumer", "Init new consumer error"), err)
		return
	}
	defer consumer.Close()

	// trap SIGINT to trigger a shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	// consume errors
	go func() {
		for err := range consumer.Errors() {
			log.Error(log.WithPrefix("driver_disconnect.consumer.InitConsumer", "consumer error"), err)

		}
	}()

	// consume notifications
	go func() {
		for ntf := range consumer.Notifications() {
			log.Info(fmt.Sprintf("Rebalanced: %+v", ntf))
		}
	}()

	// consume messages, watch signals
	for {
		select {
		case msg, ok := <-consumer.Messages():
			if ok {
				log.Debug(log.WithPrefix("driver-ledger.common-consumer", "latest offset: "), msg.Offset)
				data, err := RegistryLedger.WithSchema("com.pickme.events.driver.DriverLedgerTransactionRequest").Decode(msg.Value)

				if err != nil {
					log.Error(log.WithPrefix("", err))
				}
				dData, ok := data.(DriverLedgerTransactionRequest)

				if !ok {
					log.Error(log.WithPrefix("", err))

				}
				fmt.Println(msg)
				fmt.Println(dData)
				consumer.MarkOffset(msg, "")
				consumer.CommitOffsets()
				log.Info(log.WithPrefix("driver-ledger.common-consumer", "commited message with offset-"), msg.Offset)
				//deserializer.DriverTripDeserializer(msg.Value, 0)
			}
		case msg := <-signals:
			log.Info(log.WithPrefix("driver_trip_transaction.deserializer.InitConsumer", "Os message"), msg)
			return
		}
	}
}
