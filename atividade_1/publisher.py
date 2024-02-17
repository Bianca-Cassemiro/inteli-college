import paho.mqtt.client as mqtt
import time
import json


# Abre arquivo JSON
with open('sensor_data.json') as f:
    data = json.load(f)

sensor_data_list = data["sensor_data"]

# Configuração do cliente MQTT
client = mqtt.Client(mqtt.CallbackAPIVersion.VERSION2)

# Associa a função de callback ao cliente MQTT
# Conecta ao broker MQTT
client.connect("localhost", 1883, 60)

# Inicia o loop de comunicação MQTT
client.loop_start()

# Loop para publicar mensagens continuamente
try:
    while True:
        for sensor_data in sensor_data_list:
            message = " Valor: {} {}\n Timestamp: {}\n Localização: {}".format(sensor_data["value"], sensor_data["unit"], sensor_data["timestamp"], sensor_data["location"])
            client.publish("Bia", message)
            print(f"Leitura do sensor:\n\n{message}\n")
            time.sleep(2)
except KeyboardInterrupt:
    print("Publicação encerrada")

# Finaliza o loop de comunicação MQTT
client.loop_stop()

# Desconecta do broker MQTT
client.disconnect()
