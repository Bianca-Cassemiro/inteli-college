package testes

import (
    "fmt"
    "testing"
    "github.com/confluentinc/confluent-kafka-go/kafka"
)

func TestIntegrity(t *testing.T) {
    // Configuração do produtor
    p, err := kafka.NewProducer(&kafka.ConfigMap{
        "bootstrap.servers": "localhost:29092,localhost:39092",
        "client.id":         "go-producer",
    })
    if err != nil {
		fmt.Printf("Erro ao criar o produtor: %v\n", err)
		return
	}
    // Configuração do consumidor
    c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:29092,localhost:39092",
		"group.id":          "go-consumer-group",
		"auto.offset.reset": "earliest",
	})
    if err != nil {
        t.Errorf("Failed to create consumer: %s\n", err)
    }

    topic := "myTopic"
    message := "Test message"

    // Envia a mensagem
    p.Produce(&kafka.Message{
        TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
        Value:          []byte(message),
    }, nil)

    // Consome a mensagem
    c.SubscribeTopics([]string{topic}, nil)
    msg, err := c.ReadMessage(-1)
    if err != nil {
        t.Errorf("Failed to read message: %s\n", err)
    }

    // Verifica se a mensagem recebida é a mesma que a enviada
    if string(msg.Value) != message {
        t.Errorf("Expected message %q, got %q", message, string(msg.Value))
    }
}