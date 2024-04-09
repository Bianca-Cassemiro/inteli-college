package main

import (
	"encoding/json"
	"fmt"
	"os"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type SensorData struct {
	Id        string  `json:"idSensor"`
	Timestamp string  `json:"timestamp"`
	Poluente  string  `json:"tipoPoluente"`
	Nivel     float32 `json:"nivel"`
}

func main() {

	// Cria um arquivo chamado "mensagens.txt"
    fileTxt, err := os.Create("mensagens.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer fileTxt.Close()



	// Configurações do produtor
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
			"bootstrap.servers": "localhost:29092,localhost:39092",
			"client.id":         "go-producer",
		})
	if err != nil {
		fmt.Printf("Erro ao criar o produtor: %v\n", err)
		return
	}
	defer producer.Close()

	// Abre arquivo JSON
	file, err := os.Open("../sensor_data.json")
	if err != nil {
		fmt.Printf("Erro ao abrir o arquivo JSON: %v\n", err)
		return
	}
	defer file.Close()

	var sensorDataList []SensorData
	if err := json.NewDecoder(file).Decode(&sensorDataList); err != nil {
		fmt.Printf("Erro ao decodificar o JSON: %v\n", err)
		return
	}

	// Enviar mensagem
	topic := "qualidadeAr"
	for _, sensorData := range sensorDataList {
		message := fmt.Sprintf("Id: %s\nTimestamp: %s\nPoluente: %s\nNivel: %f", sensorData.Id, sensorData.Timestamp, sensorData.Timestamp, sensorData.Nivel)
		producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(message),
		}, nil)

		_, err = fileTxt.WriteString("\n"+ message + "\n")
   		 if err != nil {
        fmt.Println(err)
        return
   		 }

	}

	err = file.Sync()
    if err != nil {
        fmt.Println(err)
        return
    }

	// Aguardar a entrega de todas as mensagens
	producer.Flush(15 * 1000)

	// Configurações do consumidor
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:29092,localhost:39092",
		"group.id":          "go-consumer-group",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		fmt.Printf("Erro ao criar o consumidor: %v\n", err)
		return
	}
	defer consumer.Close()

	// Assinar tópico
	consumer.SubscribeTopics([]string{topic}, nil)

	// Consumir mensagens
	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			fmt.Printf("\nReceived message:\n%s\n", string(msg.Value))
		} else {
			fmt.Printf("Erro do consumidor: %v (%v)\n", err, msg)
			break
		}
	}
}
