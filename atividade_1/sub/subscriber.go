package main

import (
	"fmt"
	"os"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	godotenv "github.com/joho/godotenv"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var messagePubHandler MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("Recebido: %s do tópico: %s\n", msg.Payload(), msg.Topic())
}

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Printf("Error loading .env file: %s", err)
	}

	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:29092,localhost:39092",
		"group.id":          "go-consumer-group",
		"auto.offset.reset": "earliest",
	})

	var broker = os.Getenv("BROKER_ADDR")
	var port = 8883
	opts := MQTT.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tls://%s:%d", broker, port))
	opts.SetClientID("Subscriber")
	opts.SetUsername(os.Getenv("HIVE_USER"))
	opts.SetPassword(os.Getenv("HIVE_PSWD"))
	opts.SetDefaultPublishHandler(messagePubHandler)

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	if token := client.Subscribe("KafkaBia", 1, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		return
	}

	topic := "KafkaBia"
	consumer.SubscribeTopics([]string{topic}, nil)

	// Consumir mensagens
	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Received message: %s\n", string(msg.Value))
		} else {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
			break
		}
	}

	fmt.Println("Subscriber está rodando. Pressione CTRL+C para sair.")
	select {} 
}