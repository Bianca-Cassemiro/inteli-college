import paho.mqtt.client as mqtt

# Callback quando uma mensagem é recebida do servidor.
def on_message(client, userdata, message):
    print(f"\nRecebido:\n{message.payload.decode()} \nNo tópico {message.topic}")

# Callback para quando o cliente recebe uma resposta CONNACK do servidor.
def on_connect(cclient, userdata, flags, reason_code, properties):
    print("Conectado com código de resultado "+str(reason_code))
    # Inscreva no tópico aqui, ou se perder a conexão e se reconectar, então as
    # subscrições serão renovadas.
    client.subscribe("Bia")

# Configuração do cliente
client = mqtt.Client(mqtt.CallbackAPIVersion.VERSION2)
client.on_connect = on_connect
client.on_message = on_message

# Conecte ao broker
client.connect("localhost", 1883, 60)

# Loop para manter o cliente executando e escutando por mensagens
client.loop_forever()