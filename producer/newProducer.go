package producer

import (
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"gitlab.mytaxi.lk/pickme/go-util/log"
	"gitlab.mytaxi.lk/pickme/go/schema_registry"
	"test/test/domain"
	"time"
)

func NewProducer() {
	Initializer()
	ProducerNew = CreateNewProducer()
	//SendAutoSettlementNew()
	//SendRestaurantBankAccountCreated()

	SendTripCompleted()
	//SendTripCancelled()
	//SendTripCreated()

	//sendExternalPayment()
	//SendBankStatusChanged(143)
	//for i:=0;i<11;i++{
	//	//sendExternalPayment()
	//	SendBankStatusChanged(i)
	//}
	//SendBankStatusChanged(11)
	//sendExternalPayment()

	//SetMessageDriverLedgerTransaction()
	//SetMessageLocationChange()
	//JustProduce()
	//p.SendMessage(msg)
}

var (
	Registry *schema_registry.SchemaRegistry
)

var (
	ProducerNew sarama.SyncProducer
)

func Initializer() {

	Registry = schema_registry.NewSchemaRegistry("http://capp-schemaregistry.dev-mytaxi.com:8081")

	//Registry.Register("payment_services.payment_events.auto_settle",4)
	//Registry.Register("generic_event",2)
	//Registry.Register("com.pickme.events.AutoSettlement.AutoSettlementNew",3)

	//Registry.RegisterLatest(
	//	"com.pickme.events.driver.DriverLedgerTransactionRequest",
	//	func(data []byte) (interface{}, error) {
	//		e := domain.DriverTripTransactionRequest{}
	//		err := json.Unmarshal(data, &e)
	//		return e, err
	//	})
	//
	//Registry.RegisterLatest(
	//	"com.pickme.events.finance.DriverLedgerTransaction",
	//	func(data []byte) (interface{}, error) {
	//		e := domain.DriverTripTransactionRequest{}
	//		err := json.Unmarshal(data, &e)
	//		return e, err
	//	})
	//
	//Registry.RegisterLatest(
	//	`com.pickme.events.driver.DriverLocationChanged`,
	//	func(data []byte) (interface{}, error) {
	//		e := domain.DriverLocationChanged{}
	//		err := json.Unmarshal(data, &e)
	//		return e, err
	//	})
	//
	//Registry.RegisterLatest(
	//	`com.pickme.events.payment.BankStatusChanged`,
	//	func(data []byte) (interface{}, error) {
	//		e := domain.BankStatusChanged{}
	//		err := json.Unmarshal(data, &e)
	//		return e, err
	//	})
	//
	//Registry.RegisterLatest(
	//	`com.pickme.events.finance.externalPaymentRequestStatus`,
	//	func(data []byte) (interface{}, error) {
	//		e:=domain.ExternalPaymentRequestStatus{}
	//		err:=json.Unmarshal(data,&e)
	//		return e,err
	//	})
	//
	Registry.RegisterLatest(
		`com.pickme.events.trip.TripCompleted`,
		func(data []byte) (interface{}, error) {
			e := domain.TripCompleted{}
			err := json.Unmarshal(data, &e)
			return e, err
		})
	//Registry.RegisterLatest(
	//	`com.pickme.events.trip.TripCreated`,
	//	func(data []byte) (interface{}, error) {
	//		e:=domain.TripCreated{}
	//		err:=json.Unmarshal(data,&e)
	//		return e,err
	//	})
	//Registry.RegisterLatest(
	//	`com.pickme.events.trip.TripCancelled`,
	//	func(data []byte) (interface{},  error) {
	//		e:=domain.TripCancelled{}
	//		err:=json.Unmarshal(data,&e)
	//		return e,err
	//	})
	//Registry.RegisterLatest(
	//	`com.pickme.events.finance.RestaurantBankAccountCreated`,
	//	func(data []byte) (interface{}, error) {
	//		e:=domain.RestaurantBankAccountCreated{}
	//		err:=json.Unmarshal(data,&e)
	//		return e,err
	//	})
	//Registry.RegisterLatest(`payment_services.payment_events.auto_settle`,
	//	func(data []byte) (interface{},error) {
	//	e:=domain.AutoSettlementEvent{}
	//	err:=json.Unmarshal(data,&e)
	//	return e,err
	//})
}

func CreateNewProducer() sarama.SyncProducer {

	brokers := []string{"capp-kafka-kfk-001.dev-mytaxi.com:9092"}

	//brokers :=[]string{"localhost:9092"}
	saramaConfig := sarama.NewConfig()
	saramaConfig.Producer.RequiredAcks = sarama.WaitForAll
	saramaConfig.Producer.Return.Successes = true
	saramaConfig.Version = sarama.V0_10_0_0

	producer, err := sarama.NewSyncProducer(brokers, saramaConfig)

	if err != nil {
		log.Error(log.WithPrefix("com.pickme.events.driver.DriverLedgerTransaction", err))
		return nil
	}

	return producer

}

//func SendRestaurantBankAccountCreated()  {
//	b:=domain.BodyRestaurant{
//		MerchantId:123,
//		BankName:"NTB",
//		BankBranch:"Nugegoda",
//		AccountName:"fix",
//		AccountNumber:"456621",
//		UpdatedBy:5,
//		CompanyId:3,
//	}
//	R:=domain.RestaurantBankAccountCreated{
//	}
//	R.Body=b
//	R.CreatedAt=45661223
//	R.Expiry=456123
//	R.ID="14ddd"
//	R.Version=1
//	R.Type="RestaurantBankAccountCreated"
//	R.TraceInfo.Sampled=true
//	R.TraceInfo.ParentID=15
//	R.TraceInfo.SpanID=45
//	R.TraceInfo.TraceID.High=12
//	R.TraceInfo.TraceID.Low=5
//	data,err:=Registry.WithSchema("com.pickme.events.finance.RestaurantBankAccountCreated").Encode(R)
//	if err!=nil{
//		log.Error("serialize error",err)
//	}
//
//	msg:=&sarama.ProducerMessage{
//		Key:sarama.StringEncoder("123456"),
//		Topic:"restaurants",
//		Value:sarama.ByteEncoder(data),
//	}
//
//	_,offset,err:=ProducerNew.SendMessage(msg)
//	if err!=nil{
//		fmt.Println(err)
//	}
//	fmt.Println(offset)
//}
//func SetMessageLocationChange()  {
//	L:= domain.Location{
//		Lat: 123456,
//		Lng: 123654,
//		Accuracy: 321,
//		Speed: 23.2,
//		Timestamp:1539855332000,
//		Bearing: 45,
//	}
//	TID:=domain.TraceId{
//		High:123456789321,
//		Low:141785233698,
//	}
//	TIF:= domain.TraceInfo{
//		TraceId: TID,
//		ParentId:7654321,
//		Sampled:true,
//		SpanID:1234567,
//	}
//
//	d:=domain.DriverLocationChanged{}
//	d.Id="123driverLocation"
//	d.Type="driver_location_changed"
//	d.CreatedAt=1539855332000
//	d.Expiry = 100000
//	d.Body.Location=L
//	d.Body.DriverId=123456
//	d.Version = 3
//	d.TraceInfo= TIF
//
//
//	data,err:=Registry.WithSchema("com.pickme.events.driver.DriverLocationChanged").Encode(d)
//	if err!=nil{
//		log.Error("serialize error",err)
//	}
//
//	//h :=sarama.RecordHeader{
//	//
//	//}
////	header:=[]sarama.RecordHeader{h}
//	msg:=&sarama.ProducerMessage{
//		Key:sarama.StringEncoder("123456"),
//		Topic:"driver_location",
//		Value:sarama.ByteEncoder(data),
//		//Headers:header,
//	}
//
//	_,offset,err:=ProducerNew.SendMessage(msg)
//	if err!=nil{
//		fmt.Println(err)
//	}
//	fmt.Println(offset)
//}
//
//func SetMessageDriverLedgerTransaction(){
//
//	d:= domain.DriverTripTransaction{
//
//	}
//
//	d.Body.Amount=560.2377820114711120
//	d.Body.CreatedTime=1538987048
//	d.Body.Description="hi there"
//	d.Body.DriverID=145301
//	d.Body.TransactionCategory=5
//	d.Body.TransactionCategoryName="van"
//	//d.Body.TransactionID=1233559190
//	d.Body.TransactionType="CREDIT"
//	d.Body.TripID=555544
//	d.Body.CreatedBy=1
//
//	d.CreatedAt=time.Now().Unix()
//
//	d.CreatedAt=1538987048
//	d.Version=1
//	d.Type="driver_ledger_transaction"
//	d.TraceInfo.Sampled=true
//	d.TraceInfo.ParentID=1235456
//	d.TraceInfo.SpanID=123456
//
//	d.TraceInfo.TraceID.Low=123456
//	d.TraceInfo.TraceID.High=1234566
//
//	d.ID="123asd"
//	d.Expiry=1232456
//
//
//	data,err:=Registry.WithSchema("com.pickme.events.finance.DriverLedgerTransaction").Encode(d)
//	if err!=nil{
//		log.Error("serialize error",err)
//	}
//
//	//h :=sarama.RecordHeader{
//	//
//	//}
//	//header:=[]sarama.RecordHeader{h}
//	msg:=&sarama.ProducerMessage{
//		Topic:"driver_ledger_transaction",
//		Value:sarama.ByteEncoder(data),
//		//Headers:header,
//	}
//
//	_,offset,_:=ProducerNew.SendMessage(msg)
//	fmt.Println(offset)
//
//}
//func SendBankStatusChanged(i int){
//	body:=domain.BankStatusChangedBody{
//		Id:int64(i),
//		Status:4,
//		Bank: 2,
//		PaymentType:1,
//		PaymentTypeReferenceId:45,
//		TransactionRefId:"sds",
//		CreatedDatetime:456123,
//		UpdatedDatetime:45615852,
//	}
//
//	b:=domain.BankStatusChanged{
//		ID:"AA52",
//		Body:body,
//		CreatedAt:time.Now().UnixNano()/1000000,
//		Type:"bank_settlements",
//		Expiry:123654,
//		Version:3,
//	}
//
//
//
//	data,err:=Registry.WithSchema("com.pickme.events.payment.BankStatusChanged").Encode(b)
//	if err!=nil{
//		log.Error("serialize error",err)
//	}
//
//	msg:=&sarama.ProducerMessage{
//		Key:sarama.StringEncoder("123456"),
//		//Topic:"bank_settlements",
//		Topic:"bank_settlements",
//		Value:sarama.ByteEncoder(data),
//		//Headers:header,
//	}
//
//	_,offset,err:=ProducerNew.SendMessage(msg)
//	if err!=nil{
//		fmt.Println(err)
//	}
//	fmt.Println(offset)
//}
//func sendExternalPayment(){
//	b:=domain.ExternalPaymentRequestStatus{
//
//	}
//	body:=b.Body
//	body.ID=10
//	body.CreatedDateTime=time.Now().UnixNano()/1000000
//	body.DriverID=1
//	body.OrderID="whyddamitha"
//	body.PaymentMethod=3
//	body.PaymentStatus=2
//	body.PaymentTry=0
//
//
//
//	traceInfo:=b.TraceInfo
//	traceid:=traceInfo.TraceID
//
//	traceid.High=123456
//	traceid.Low=123456
//
//	traceInfo.ParentID=123456
//	traceInfo.SpanID=123456
//	traceInfo.Sampled=true
//	traceInfo.TraceID=traceid
//
//	b.Body=body
//	b.CreatedAt=time.Now().UnixNano()/1000000
//	b.ID="A1"
//	b.Version=1
//	b.Type="external_payment_request_status"
//	b.Expiry=123456
//	b.TraceInfo=traceInfo
//
//	daa,_:=json.Marshal(&b)
//	t:=domain.ExternalPaymentRequestStatus{}
//	err:=json.Unmarshal(daa,&t)
//	data,err:=Registry.WithSchema("com.pickme.events.finance.externalPaymentRequestStatus").Encode(b)
//	if err!=nil{
//		log.Error("serialize error",err)
//	}
//	//dataLater,err:=Registry.WithSchema("com.pickme.events.finance.externalPaymentRequestStatus").Decode(data)
//	//result,_:=json.MarshalIndent(t,"","")
//	//fmt.Println(string(result))
//	//h :=sarama.RecordHeader{
//	//
//	//}
//	//header:=[]sarama.RecordHeader{h}
//	msg:=&sarama.ProducerMessage{
//		Topic:"external_payment_request_status",
//		Value:sarama.ByteEncoder(data),
//		//Headers:header,
//	}
//
//	_,offset,_:=ProducerNew.SendMessage(msg)
//	fmt.Println(offset)
//}
func SendTripCompleted() {
	t := domain.TripCompleted{}
	t.ID = "100"
	t.Type = "trip_completed"
	t.CreatedAt = 1551329541000
	t.Expiry = 123456789
	t.Version = 2

	t.Body.PassengerID = 1000
	t.Body.DriverID = 21841
	t.Body.CurrencyCode = "LKR"

	t.Body.Trip.ID = 3
	t.Body.Trip.Distance = 100
	t.Body.Trip.BookedBy = 1
	t.Body.Trip.TripCost = 166.24

	t.Body.Trip.Flags.Itc = true

	pay := []domain.Payment{}
	p := domain.Payment{}
	p.Method = 1
	p.Amount = 123.12
	pay = append(pay, p)
	t.Body.Trip.Payment = pay

	t.Body.Trip.Discount = 0.0
	t.Body.Trip.PromoCode = "pickme"

	t.Body.Trip.ActualPickup.Address = "Pickme headquarters"
	t.Body.Trip.ActualPickup.Lng = 234567.4567
	t.Body.Trip.ActualPickup.Lat = 234567.4567

	t.Body.Trip.ActualDrop.Address = "Pickme headquarters"
	t.Body.Trip.ActualDrop.Lat = 234567.4567
	t.Body.Trip.ActualDrop.Lng = 234567.4567

	t.TraceInfo.TraceID.Low = 141785233698
	t.TraceInfo.TraceID.High = 141785233698

	t.TraceInfo.SpanID = 123456
	t.TraceInfo.ParentID = 654321
	t.TraceInfo.Sampled = true

	data, err := Registry.WithSchema("com.pickme.events.trip.TripCompleted").Encode(t)
	if err != nil {
		log.Error("serialize error", err)
	}

	msg := &sarama.ProducerMessage{
		Key: sarama.StringEncoder("80"),
		//Topic:"bank_settlements",
		Topic: "trip",
		Value: sarama.ByteEncoder(data),
		//Timestamp:time.Now(),
		//Headers:header,
	}

	_, offset, err := ProducerNew.SendMessage(msg)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(offset)
}
func SendTripCancelled() {
	t := domain.TripCancelled{}
	t.ID = "100"
	t.Type = "trip_cancelled"
	t.CreatedAt = 152312146123
	t.Expiry = 123456789
	t.Version = 1
	t.Body.TripID = 3
	t.Body.CancelledBy = 1
	t.Body.CancelledFrom = 10
	t.Body.CancelType = 1
	t.Body.Note = "trip has been cancelled"
	t.Body.ReasonID = 1

	t.TraceInfo.TraceID.High = 1223554
	t.TraceInfo.TraceID.Low = 12345625

	t.TraceInfo.SpanID = 11222
	t.TraceInfo.ParentID = 112233
	t.TraceInfo.Sampled = true

	data, err := Registry.WithSchema("com.pickme.events.trip.TripCancelled").Encode(t)
	if err != nil {
		log.Error("serialize error", err)
	}

	msg := &sarama.ProducerMessage{
		Key:   sarama.StringEncoder("75"),
		Topic: "testtrip",
		Value: sarama.ByteEncoder(data),
	}
	_, offset, err := ProducerNew.SendMessage(msg)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(offset)
}
func SendTripCreated() {
	t := domain.TripCreated{}
	t.ID = "105"
	t.Type = "trip_created"
	t.CreatedAt = 1551329541000
	t.Expiry = 123456789
	t.Version = 2

	t.Body.Module = 1
	t.Body.ServiceGroupCode = "RIDES"
	t.Body.BookedBy = 1
	t.Body.TripID = 1234567
	t.Body.VehicleType = 3
	t.Body.PreBooking = true

	t.Body.Passenger.ID = 1

	t.Body.Driver.ID = 0

	t.Body.Corporate.ID = 1
	t.Body.Corporate.DepID = 1

	t.Body.Pickup.Time = 1551

	loc := []domain.Location2{}
	l := domain.Location2{}
	l.Address = "Pickme HeadQuarters"
	l.Lat = 6.89
	l.Lng = 79.85
	loc = append(loc, l)
	t.Body.Pickup.Location = loc

	loc2 := []domain.Location2{}
	loc2 = append(loc2, l)
	t.Body.Drop.Location = loc2

	t.Body.Promotion.Code = "test"

	r := []int{}
	r = append(r, 1)
	t.Body.Region.Ids = r

	t.Body.Payment.PrimaryMethod = 1
	t.Body.Payment.SecondaryMethod = 1

	t.Body.Comments.Remark = ""
	t.Body.Comments.DriverNotes = ""

	t.Body.Filters.Driver.LanguageID = 0
	t.Body.Filters.Vehicle.CompanyID = 0
	t.Body.Filters.Vehicle.BrandID = 0
	t.Body.Filters.Vehicle.ColorID = 0

	t.Body.Surge.RegionID = 111
	t.Body.Surge.Value = 100

	t.Body.FareDetails.FareType = "GEO"
	t.Body.FareDetails.MinKm = 25
	t.Body.FareDetails.MinFare = 1000
	t.Body.FareDetails.AdditionalKmFare = 10
	t.Body.FareDetails.WaitingTimeFare = 25
	t.Body.FareDetails.FreeWaitingTime = 25
	t.Body.FareDetails.NightFare = 10
	t.Body.FareDetails.RideHours = 10
	t.Body.FareDetails.ExtraRideFare = 10
	t.Body.FareDetails.DriverBata = 50
	t.Body.FareDetails.TripType = 1

	t.TraceInfo.TraceID.Low = 141785233698
	t.TraceInfo.TraceID.High = 141785233698

	t.TraceInfo.SpanID = 123456
	t.TraceInfo.ParentID = 654321
	t.TraceInfo.Sampled = true

	data, err := Registry.WithSchema("com.pickme.events.trip.TripCreated").Encode(t)
	if err != nil {
		log.Error("serialize error", err)
	}

	msg := &sarama.ProducerMessage{
		Key: sarama.StringEncoder("80"),
		//Topic:"bank_settlements",
		Topic:     "testtrip",
		Value:     sarama.ByteEncoder(data),
		Timestamp: time.Now(),
		//Headers:header,
	}

	_, offset, err := ProducerNew.SendMessage(msg)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(offset)
}

//func encodePrefix(id int) []byte {
//	byt := make([]byte, 5)
//	binary.BigEndian.PutUint32(byt[1:], uint32(id))
//	return byt
//}
//func SerializeWithSchema(event interface{}, schemaName string) ([]byte, error) {
//	// serialize with schema
//	byt, err := Registry.WithSchema(schemaName).Encode(&event)
//	if err != nil {
//		//log.Fatal(err)
//		log.Error(err)
//		return nil, err
//	}
//	schema,err:=Registry.GetBySubject(schemaName,1)
//	log.Debug("schema id",schema.Id)
//	prefix:=encodePrefix(schema.Id)
//	return append(prefix,byt...), nil
//}
//func SendAutoSettlement(){
//	r:=domain.AutoSettlementEvent{}
//
//	p:=domain.PayRecord{}
//	p.TransactionId=111111111
//	p.AccountHolder="cc"
//	p.BankName=5
//
//	r.PaymentRecords =append(r.PaymentRecords,p)
//	data,err:=SerializeWithSchema(r,"payment_services.payment_events.auto_settle")
//	if err!=nil{
//		log.Error("serialize error",err)
//		return
//	}
//	//message := domain.Event{
//	//	Id:        "",
//	//	Type:      "auto_settlement",
//	//	CreatedAt: time.Now().Unix(),
//	//	Expiry:    int64(time.Duration(time.Hour * 24).Seconds()),
//	//	Version:   4,
//	//	Body:      data,
//	//}
//	//genEventByte,err:=SerializeWithSchema(message,"generic_event")
//	msg:=&sarama.ProducerMessage{
//		Key:sarama.StringEncoder("123456"),
//		//Topic:"bank_settlements",
//		Topic:"AutoSettlement_New",
//		Value:sarama.ByteEncoder(data),
//		//Headers:header,
//	}
//
//	_,offset,err:=ProducerNew.SendMessage(msg)
//	if err!=nil{
//		fmt.Println(err)
//	}
//	fmt.Println(offset)
//}
//func SendAutoSettlementNew(){
//	r:=domain.AutoSettlementNew{}
//
//	p:=domain.PayRecord{}
//	p.TransactionId=111111111
//	p.AccountHolder="cc"
//	p.BankName=5
//
//	r.Body.PaymentRecords =append(r.Body.PaymentRecords,p)
//	data,err:=SerializeWithSchema(r,"com.pickme.events.AutoSettlement.AutoSettlementNew")
//	if err!=nil{
//		log.Error("serialize error",err)
//		return
//	}
//	//message := domain.Event{
//	//	Id:        "",
//	//	Type:      "auto_settlement",
//	//	CreatedAt: time.Now().Unix(),
//	//	Expiry:    int64(time.Duration(time.Hour * 24).Seconds()),
//	//	Version:   4,
//	//	Body:      data,
//	//}
//	//genEventByte,err:=SerializeWithSchema(message,"generic_event")
//	msg:=&sarama.ProducerMessage{
//		Key:sarama.StringEncoder("123456"),
//		//Topic:"bank_settlements",
//		Topic:"AutoSettlement_New",
//		Value:sarama.ByteEncoder(data),
//		//Headers:header,
//	}
//
//	_,offset,err:=ProducerNew.SendMessage(msg)
//	if err!=nil{
//		fmt.Println(err)
//	}
//	fmt.Println(offset)
//}
