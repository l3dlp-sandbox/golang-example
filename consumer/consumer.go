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
	Registry *schema_registry.SchemaRegistry
)

type DriverTripTransaction struct {
	Body struct {
		Amount                  float64 `json:"amount"`
		CreatedBy               int     `json:"created_by"`
		CreatedTime             int64   `json:"created_time"`
		Description             string  `json:"description"`
		DriverID                int     `json:"driver_id"`
		TransactionCategory     int     `json:"transaction_category"`
		TransactionCategoryName string  `json:"transaction_category_name"`
		TransactionID           int     `json:"transaction_id"`
		TransactionType         string  `json:"transaction_type"`
		TripID                  int     `json:"trip_id"`
	} `json:"body"`
	CreatedAt int    `json:"created_at"`
	Expiry    int    `json:"expiry"`
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

type Kconsumer struct {
}

func (k Kconsumer) Initializer() {

	Registry = schema_registry.NewSchemaRegistry("http://35.184.181.97:8089/")

	Registry.Register(
		"com.pickme.events.finance.DriverLedgerTransaction",
		3,
		func(data []byte) (interface{}, error) {
			e := DriverTripTransaction{}
			err := json.Unmarshal(data, &e)
			return e, err
		},
	)
}

func (k Kconsumer) Consumer() {

	config := cluster.NewConfig()
	config.Consumer.Return.Errors = true
	config.Group.Return.Notifications = true
	config.Group.Mode = cluster.ConsumerModeMultiplex
	config.Consumer.Offsets.Initial = sarama.OffsetNewest
	config.Version = sarama.V1_0_0_0

	//brokers := []string{"kafka-2.c.charming-opus-833.internal:9092","kafka-0.c.charming-opus-833.internal:9092","kafka-1.c.charming-opus-833.internal:9092"}

	brokers := []string{"localhost:9092"}

	topics := []string{"driver_ledger_transaction"}

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
				data, err := Registry.WithSchema("com.pickme.events.finance.DriverLedgerTransaction").Decode(msg.Value)

				if err != nil {
					log.Error(log.WithPrefix("", err))
				}
				dData, ok := data.(DriverTripTransaction)

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
