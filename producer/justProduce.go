package producer

import (
	"github.com/Shopify/sarama"
	"fmt"
)

func JustProduce()  {
	data:=[]byte("tesing")
	h :=sarama.RecordHeader{

	}
	header:=[]sarama.RecordHeader{h}
	msg:=&sarama.ProducerMessage{
		Topic:"driver_ledger_transaction",
		Value:sarama.ByteEncoder(data),
		Headers:header,
	}
	_,offset,_:=ProducerNew.SendMessage(msg)

	fmt.Println(offset)
}