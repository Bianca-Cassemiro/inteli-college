package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/eclipse/paho.mqtt.golang"
	"github.com/joho/godotenv"
)

var (
	brokerAddress = os.Getenv("BROKER_ADDR")
	port          = 8883
	topic         = "my/test/topic"
	username      = os.Getenv("HIVE_USER")
	password      = os.Getenv("HIVE_PSWD")
)

func onConnectHandler(client mqtt.Client) {
	fmt.Println("CONNACK received")
	if token := client.Subscribe(topic, 1, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
	}
}

func onMessageHandler(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("%s (QoS: %d) - %s\n", msg.Topic(), msg.Qos(), msg.Payload())
	insertData(msg.Payload())
}

func insertData(payload []byte) {
	db, err := sql.Open("sqlite3", "dados.db")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}
	defer db.Close()

	sqlStmt := "CREATE TABLE IF NOT EXISTS dados (id INTEGER PRIMARY KEY, valor TEXT)"
	_, err = db.Exec(sqlStmt)
	if err != nil {
		fmt.Println("Error creating table:", err)
		return
	}

	sqlStmt = "INSERT INTO dados (valor) VALUES (?)"
	_, err = db.Exec(sqlStmt, string(payload))
	if err != nil {
		fmt.Println("Error inserting data:", err)
		return
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
	}

	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tls://%s:%d", brokerAddress, port))
	opts.SetUsername(username)
	opts.SetPassword(password)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println("Error connecting to broker:", token.Error())
		return
	}

	client.Subscribe(topic, 1, onMessageHandler)

	select {}
}
