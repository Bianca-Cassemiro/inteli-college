package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
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

	// Configurações do produtor Kafka
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:29092,localhost:39092",
		"client.id":         "go-producer",
	})
	if err != nil {
		fmt.Printf("Erro ao criar o produtor: %v\n", err)
		return
	}
	defer producer.Close()

	// Conexão com o MongoDB Atlas
	mongoURI := "mongodb+srv://@cluster-bia.htq4jn2.mongodb.net/?retryWrites=true&w=majority&appName=Cluster-Bia"
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

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

	// Loop para publicar mensagens continuamente
	for {
		for _, sensorData := range sensorDataList {
	

						message := fmt.Sprintf("Valor: %.2f %s\nTimestamp: %s\nLocalização: %s", sensorData.Value, sensorData.Unit, sensorData.Timestamp, sensorData.Location)
						topic := "KafkaBia"
						producer.Produce(&kafka.Message{
							TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
							Value:          []byte(message),
						}, nil)
						fmt.Printf("Leitura do sensor:\n\n%s\n\n", message)

						// Salvar a mensagem no MongoDB
						collection := client.Database("sensor_data").Collection("data") 
						_, err := collection.InsertOne(ctx, bson.M{"message": message})
						if err != nil {
							log.Println("Erro ao salvar a mensagem no MongoDB:", err)
							continue
						}
			log.Println("Mensagem salva no MongoDB:", message)

			time.Sleep(2 * time.Second)
		}
	}
}
