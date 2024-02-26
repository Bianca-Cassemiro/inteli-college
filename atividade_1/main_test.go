package main_test

import (
	"testing"
	"time"
	"fmt"
	"os"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/eclipse/paho.mqtt.golang"
)
type SensorData struct { // Define the SensorData structure
    Value     float64 `json:"value"`
    Unit      string  `json:"unit"`
    Timestamp string  `json:"timestamp"`
    Location  string  `json:"location"`
}
var receivedMessages chan mqtt.Message 

func TestRecebimento(t *testing.T) {
    // Cria um cliente MQTT de teste.
	sensorDataList := []SensorData{
        {Value: 12.5, Unit: "°C", Timestamp: "2024-02-26T15:18:00Z", Location: "Sala"},
        {Value: 200, Unit: "RPM", Timestamp: "2024-02-26T15:18:01Z", Location: "Motor"},
    }
    client := createTestMQTTServer("localhost:1883", "Bia", nil)

    // Publica mensagens simuladas.
    for _, sensorData := range sensorDataList {
        message := fmt.Sprintf("Valor: %.2f %s\nTimestamp: %s\nLocalização: %s", sensorData.Value, sensorData.Unit, sensorData.Timestamp, sensorData.Location)
        client.Publish("Bia", 0, false, message)
    }

    // Verifica se as mensagens foram recebidas.
    for _, expectedData := range sensorDataList {
        received := <-receivedMessages
        assert.Contains(t, received, fmt.Sprintf("Valor: %.2f %s", expectedData.Value, expectedData.Unit))
        assert.Contains(t, received, fmt.Sprintf("Timestamp: %s", expectedData.Timestamp))
        assert.Contains(t, received, fmt.Sprintf("Localização: %s", expectedData.Location))
    }

    // Desconexão do cliente MQTT.
    client.Disconnect(0)
}

func TestValidacaoDosDados(t *testing.T) {
    // Lê os dados de sensores do arquivo JSON.
	client := createTestMQTTServer("localhost:1883", "Bia", nil)

    file, err := os.Open("sensor_data.json")
    if err != nil {
        t.Fatal("erro ao abrir o arquivo sensor_data.json:", err)
    }
    defer file.Close()

    var expectedDataList []SensorData
    if err := json.NewDecoder(file).Decode(&expectedDataList); err != nil {
        t.Fatal("erro ao decodificar o arquivo JSON:", err)
    }

    // Publica os dados e recebe as mensagens.
    for _, expectedData := range expectedDataList {
        message := fmt.Sprintf("Valor: %.2f %s\nTimestamp: %s\nLocalização: %s", expectedData.Value, expectedData.Unit, expectedData.Timestamp, expectedData.Location)
        client.Publish("Bia", 0, false, message)
        received := <-receivedMessages

        // Compara os dados.
        assert.Equal(t, expectedData.Value, received.Value)
        assert.Equal(t, expectedData.Unit, received.Unit)
        assert.Equal(t, expectedData.Timestamp, received.Timestamp)
        assert.Equal(t, expectedData.Location, received.Location)
    }
}

func TestConfirmacaoDaTaxaDeDisparo(t *testing.T) {
    // Define a quantidade de mensagens e o intervalo de tempo.
    const messageCount = 10
    const interval = 2 * time.Second
	client := createTestMQTTServer("localhost:1883", "Bia", nil)
    // Publica as mensagens.
    start := time.Now()
    for i := 0; i < messageCount; i++ {
        message := fmt.Sprintf("Valor: %.2f %s\nTimestamp: %s\nLocalização: %s", sensorData.Value, sensorData.Unit, sensorData.Timestamp, sensorData.Location)
        client.Publish("Bia", 0, false, message)
    }

    // Calcula a taxa de disparo real
    elapsed := time.Since(start)
    actualRate := float64(messageCount) / elapsed.Seconds()

    // Define a margem de erro aceitável.
    const allowedError = 0.1

    // Verifica se a taxa de disparo está dentro da margem de erro.
    expectedRate := float64(messageCount) / interval.Seconds()
    assert.InDelta(t, expectedRate, actualRate, allowedError, "taxa de disparo fora da margem de erro")
}

func createTestMQTTServer(brokerURL, topic string, messageHandler mqtt.MessageHandler) mqtt.Client {
    // Cria as opções do cliente MQTT.
    opts := mqtt.NewClientOptions().AddBroker(brokerURL)
    opts.SetClientID("test-client")
    opts.SetCleanSession(true)

    // Cria um novo cliente MQTT.
    client := mqtt.NewClient(opts)

    // Conecta ao broker.
    if token := client.Connect(); token.Wait() && token.Error() != nil {
        panic(token.Error())
    }

    // Se um manipulador de mensagens for fornecido, inscreve-se no tópico com o manipulador.
    if messageHandler != nil {
        if token := client.Subscribe(topic, 0, messageHandler); token.Wait() && token.Error() != nil {
            panic(token.Error())
        }
    }

    // Retorna o cliente MQTT.
    return client
}
