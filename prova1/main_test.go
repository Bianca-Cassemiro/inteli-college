package main_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/eclipse/paho.mqtt.golang"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

type SensorData struct {
	Id          string  `json:"id"`
	Tipo        string  `json:"tipo"`
	Temperatura float64 `json:"temperatura"`
	Timestamp   string  `json:"timestamp"`
}

var sensorDataList = []SensorData{
	{ Id: "1", Tipo: "freezer", Temperatura:-13,Timestamp:"01/03/2024 14:30"},
	{ Id: "2", Tipo: "geladeira", Temperatura:10,Timestamp:"01/03/2024 14:30"},
}

var receivedMessages = make(chan mqtt.Message)

func TestRecebimento(t *testing.T) {
	client := createTestMQTTServer("Bia", func(client mqtt.Client, msg mqtt.Message) {
		receivedMessages <- msg
	})

	defer client.Disconnect(0)

	for _, sensorData := range sensorDataList {
		message := fmt.Sprintf( "Id: %s, Tipo: %s, Temperatura:%f,Timestamp:%s", sensorData.Id, sensorData.Tipo, sensorData.Temperatura,sensorData.Timestamp)
		client.Publish("Bia", 0, false, message)
	}

	for _, expectedData := range sensorDataList {
		received := <-receivedMessages

		assert.Contains(t, string(received.Payload()), fmt.Sprintf("Id: %s", expectedData.Id))
		assert.Contains(t, string(received.Payload()), fmt.Sprintf("Tipo: %s", expectedData.Tipo))
		assert.Contains(t, string(received.Payload()), fmt.Sprintf("Temperatura: %f", expectedData.Temperatura))
		assert.Contains(t, string(received.Payload()), fmt.Sprintf("Timestamp: %s", expectedData.Timestamp))
	}
}

func TestValidacaoDosDados(t *testing.T) {
	client := createTestMQTTServer("Bia", func(client mqtt.Client, msg mqtt.Message) {
		receivedMessages <- msg
	})

	defer client.Disconnect(0)

	for _, sensorData := range sensorDataList {
		message := fmt.Sprintf("Id: %s \nTipo: %s\nTemperatura: %f   \nTimestamp: %s", sensorData.Id, sensorData.Tipo, sensorData.Temperatura,sensorData.Timestamp)
		client.Publish("Bia", 0, false, message)
	}

	for _, expectedData := range sensorDataList {
		received := <-receivedMessages

		var receivedData SensorData
		err := json.Unmarshal(received.Payload(), &receivedData)
		if err != nil {
			t.Errorf("Erro ao decodificar a mensagem recebida: %s", err)
		}

		assert.Equal(t, expectedData, receivedData)
	}
}

func TestConfirmacaoDaTaxaDeDisparo(t *testing.T) {
	client := createTestMQTTServer("Bia", nil)

	defer client.Disconnect(0)

	// Publica mensagens simuladas.
	for _, sensorData := range sensorDataList {
		message := fmt.Sprintf("Id: %s \nTipo: %s\nTemperatura: %f   \nTimestamp: %s", sensorData.Id, sensorData.Tipo, sensorData.Temperatura,sensorData.Timestamp)
		client.Publish("Bia", 0, false, message)
		time.Sleep(1 * time.Second)
	}
}

func TestIntegracaoHiveMQ(t *testing.T) {
	godotenv.Load(".env")
	var broker = os.Getenv("BROKER_ADDR")
	var port = 8883
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tls://%s:%d", broker, port))
	opts.SetClientID("Publisher")
	opts.SetUsername(os.Getenv("HIVE_USER"))
	opts.SetPassword(os.Getenv("HIVE_PSWD"))

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		t.Fatalf("Falha ao conectar ao HiveMQ: %v", token.Error())
	}
	defer client.Disconnect(0)

	// Publica mensagens simuladas.
	for _, sensorData := range sensorDataList {
		message := fmt.Sprintf("Id: %s \nTipo: %s\nTemperatura: %f   \nTimestamp: %s", sensorData.Id, sensorData.Tipo, sensorData.Temperatura,sensorData.Timestamp)
		token := client.Publish("Bia", 0, false, message)
		token.Wait()
		if token.Error() != nil {
			t.Fatalf("Falha ao publicar mensagem: %v", token.Error())
		}
	}
}

func createTestMQTTServer(topic string, messageHandler mqtt.MessageHandler) mqtt.Client {
	godotenv.Load(".env")
	var broker = os.Getenv("BROKER_ADDR")
	var port = 8883
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tls://%s:%d", broker, port))
	opts.SetClientID("test-client")
	opts.SetCleanSession(true)
	opts.SetUsername(os.Getenv("HIVE_USER"))
	opts.SetPassword(os.Getenv("HIVE_PSWD"))

	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	if messageHandler != nil {
		if token := client.Subscribe(topic, 0, messageHandler); token.Wait() && token.Error() != nil {
			panic(token.Error())
		}
	}

	return client
}