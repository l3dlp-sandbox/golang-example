package boostrap

import (
	"encoding/json"

	"gitlab.mytaxi.lk/pickme/go/schema_registry"
	"test/test/domain"
)

var (
	Registry *schema_registry.SchemaRegistry
)

func Initializer() {

	Registry = schema_registry.NewSchemaRegistry("http://35.184.181.97:8089/")

	Registry.RegisterLatest(
		"payment_services.error_events.auto_settle",
		func(data []byte) (interface{}, error) {
			e := domain.RetryEvent{}
			err := json.Unmarshal(data, &e)
			return e, err
		})

	Registry.RegisterLatest(
		"com.pickme.events.finance.DriverLedgerTransaction",
		func(data []byte) (interface{}, error) {
			e := domain.DriverTripTransaction{}
			err := json.Unmarshal(data, &e)
			return e, err
		})
}
