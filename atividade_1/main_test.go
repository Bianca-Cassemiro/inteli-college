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
	Value     float64 `json:"value"`
	Unit      string  `json:"unit"`
	Timestamp string  `json:"timestamp"`
	Location  string  `json:"location"`
}

var sensorDataList = []SensorData{
	{Value: 12.5, Unit: "°C", Timestamp: "2024-02-26T15:18:00Z", Location: "Sala"},
	{Value: 200, Unit: "RPM", Timestamp: "2024-02-26T15:18:01Z", Location: "Quarto"},
}

var receivedMessages = make(chan mqtt.Message)

func TestRecebimento(t *testing.T) {
	client := createTestMQTTServer("Bia", func(client mqtt.Client, msg mqtt.Message) {
		receivedMessages <- msg
	})

	defer client.Disconnect(0)

	// Publica mensagens simuladas.
	for _, sensorData := range sensorDataList {
		message := fmt.Sprintf("Valor: %.2f %s\nTimestamp: %s\nLocalização: %s", sensorData.Value, sensorData.Unit, sensorData.Timestamp, sensorData.Location)
		client.Publish("Bia", 0, false, message)
	}

	// Verifica se as mensagens foram recebidas e se contêm os dados esperados.
	for _, expectedData := range sensorDataList {
		received := <-receivedMessages

		assert.Contains(t, string(received.Payload()), fmt.Sprintf("Valor: %.2f %s", expectedData.Value, expectedData.Unit))
		assert.Contains(t, string(received.Payload()), fmt.Sprintf("Timestamp: %s", expectedData.Timestamp))
		assert.Contains(t, string(received.Payload()), fmt.Sprintf("Localização: %s", expectedData.Location))
	}
}

func TestValidacaoDosDados(t *testing.T) {
	client := createTestMQTTServer("Bia", func(client mqtt.Client, msg mqtt.Message) {
		receivedMessages <- msg
	})

	defer client.Disconnect(0)

	// Publica mensagens simuladas.
	for _, sensorData := range sensorDataList {
		message := fmt.Sprintf(`{"value": %.2f, "unit": "%s", "timestamp": "%s", "location": "%s"}`, sensorData.Value, sensorData.Unit, sensorData.Timestamp, sensorData.Location)
		client.Publish("Bia", 0, false, message)
	}

	// Verifica se as mensagens foram recebidas e se contêm os dados esperados.
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
		message := fmt.Sprintf("Valor: %.2f %s\nTimestamp: %s\nLocalização: %s", sensorData.Value, sensorData.Unit, sensorData.Timestamp, sensorData.Location)
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
		message := fmt.Sprintf("Valor: %.2f %s\nTimestamp: %s\nLocalização: %s", sensorData.Value, sensorData.Unit, sensorData.Timestamp, sensorData.Location)
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
