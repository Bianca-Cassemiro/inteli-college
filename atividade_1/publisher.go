package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

type SensorData struct {
	Value     float64 `json:"value"`
	Unit      string  `json:"unit"`
	Timestamp string  `json:"timestamp"`
	Location  string  `json:"location"`
}

func main() {
	// Abre arquivo JSON
	file, err := os.Open("sensor_data.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var sensorDataList []SensorData
	if err := json.NewDecoder(file).Decode(&sensorDataList); err != nil {
		panic(err)
	}

	// Configuração do cliente MQTT
	opts := MQTT.NewClientOptions().AddBroker("tcp://localhost:1883")
	opts.SetClientID("go_publisher")

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// Loop para publicar mensagens continuamente

		for {
			for _, sensorData := range sensorDataList {
				message := fmt.Sprintf("Valor: %.2f %s\nTimestamp: %s\nLocalização: %s", sensorData.Value, sensorData.Unit, sensorData.Timestamp, sensorData.Location)
				token := client.Publish("Bia", 0, false, message)
				token.Wait()
				fmt.Printf("Leitura do sensor:\n\n%s\n\n", message)
				time.Sleep(2 * time.Second)
			}
		}

		
			
	}
	
