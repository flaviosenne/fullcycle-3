package main

import (
	"encoding/json"
	"fmt"
	"sync"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/flaviosenne/fullcycle-3/invest-golang/internal/infra/kafka"
	"github.com/flaviosenne/fullcycle-3/invest-golang/internal/market/dto"
	"github.com/flaviosenne/fullcycle-3/invest-golang/internal/market/entity"
	"github.com/flaviosenne/fullcycle-3/invest-golang/internal/market/transformer"
)

func main() {
	ordersIn := make(chan *entity.Order)
	ordersOut := make(chan *entity.Order)
	wg := &sync.WaitGroup{}
	defer wg.Wait()

	kafkaMsgChan := make(chan *ckafka.Message)
	configMap := ckafka.ConfigMap{
		"bootstrap.servers": "localhost:9094",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	}

	producer := kafka.NewProducer(&configMap)
	kafka := kafka.NewConsumer(&configMap, []string{"input"})

	go kafka.Consume(kafkaMsgChan) // thread 2

	book := entity.NewBook(ordersIn, ordersOut, wg)
	go book.Trade() // thread 3

	go func() {
		for msg := range kafkaMsgChan {
			wg.Add(1)
			fmt.Println(string(msg.Value))
			tradeInput := dto.TradeInput{}
			err := json.Unmarshal(msg.Value, &tradeInput)
			if err != nil {
				panic(err)
			}
			order := transformer.TransformeInput(tradeInput)
			ordersIn <- order
		}
	}()

	for res := range ordersOut {
		output := transformer.TransformeOutput(res)
		outputJson, err := json.MarshalIndent(output, "", "  ")
		if err != nil {
			fmt.Println(err)
		}
		producer.Publish(outputJson, []byte("orders"), "output")
	}
}
