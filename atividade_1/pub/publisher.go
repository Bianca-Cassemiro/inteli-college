package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
	godotenv "github.com/joho/godotenv"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/confluentinc/confluent-kafka-go/kafka"

)

type SensorData struct {
	Value     float64 `json:"value"`
	Unit      string  `json:"unit"`
	Timestamp string  `json:"timestamp"`
	Location  string  `json:"location"`
}

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Printf("Error loading .env file: %s", err)
	}

	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:29092,localhost:39092",
		"client.id":         "go-producer",
	})

	// Abre arquivo JSON
	file, err := os.Open("../sensor_data.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var sensorDataList []SensorData
	if err := json.NewDecoder(file).Decode(&sensorDataList); err != nil {
		panic(err)
	}

	// Configuração do cliente MQTT
	var broker = os.Getenv("BROKER_ADDR")
	var port = 8883
	opts := MQTT.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tls://%s:%d", broker, port))
	opts.SetClientID("Publisher")
	opts.SetUsername(os.Getenv("HIVE_USER"))
	opts.SetPassword(os.Getenv("HIVE_PSWD"))
	

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// Loop para publicar mensagens continuamente

		for {
			for _, sensorData := range sensorDataList {
				message := fmt.Sprintf("Valor: %.2f %s\nTimestamp: %s\nLocalização: %s", sensorData.Value, sensorData.Unit, sensorData.Timestamp, sensorData.Location)
				token := client.Publish("KafkaBia", 0, false, message)
				topic := "KafkaBia"
				producer.Produce(&kafka.Message{
					TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
					Value:          []byte(message),
				}, nil)
				token.Wait()
				fmt.Printf("Leitura do sensor:\n\n%s\n\n", message)
				time.Sleep(2 * time.Second)
			}
		}
			
	}
	
