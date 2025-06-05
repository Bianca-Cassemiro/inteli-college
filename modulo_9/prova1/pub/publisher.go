package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	godotenv "github.com/joho/godotenv"
)

type SensorData struct {
	Id          string  `json:"id"`
	Tipo        string  `json:"tipo"`
	Temperatura float64 `json:"temperatura"`
	Timestamp   string  `json:"timestamp"`
}

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Printf("Error loading .env file: %s", err)
	}

	// Abre arquivo JSON
	file, err := os.Open("../data.json")
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
			
			var caution = ""

			if (sensorData.Temperatura > -15) && (sensorData.Tipo == "freezer") { 
				caution = "|| ALERTA: Temperatura ALTA no freezer"
			}

			if (sensorData.Temperatura < -26) && (sensorData.Tipo == "freezer") { 
				caution = "|| ALERTA: Temperatura BAIXA no freezer"
			}
			
			if (sensorData.Temperatura > 10) && (sensorData.Tipo == "geladeira") { 
				caution = "|| ALERTA: Temperatura ALTA na geladeira"
			}
			if (sensorData.Temperatura < 2) && (sensorData.Tipo == "geladeira") { 
				caution = "|| ALERTA: Temperatura BAIXA na geladeira"
			}

			message := fmt.Sprintf("Id: %s \nTipo: %s\nTemperatura: %f   %s\nTimestamp: %s", sensorData.Id, sensorData.Tipo, sensorData.Temperatura,caution, sensorData.Timestamp)
			token := client.Publish("Bia", 0, false, message)
			token.Wait()
			fmt.Printf("Leitura do sensor:\n\n%s\n\n", message)
			time.Sleep(2 * time.Second)
		}
	}
	
}
