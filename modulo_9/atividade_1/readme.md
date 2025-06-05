# Simulador Iot
A primeira atividade do módulo tem como objetivo a criação de um simulador de dispositivos IoT capaz de enviar informações em um tópico com o formato de dados consistente.

## Entrega
Criei um arquivo Json que continha as informações dos sensores bem como no exemplo abaixo:
```
       {
            "unit": "W/m2",
            "value": 1024,
            "timestamp": "2024-02-13 00:00:00",
            "location": "51.5074, -0.1278"  
        }
```
Após isso criei um script publisher.go e configurei para realizar a leitura dos dados estabelecidos no Json e publica-los em um tópico chamado "Bia"
Criei um subscriber.go que se inscrevia no tópico Bia e realizava a leitura do que estava sendo publicado.
# Testes

## Teste de Recebimento:
Garante que o broker recebe as mensagens publicadas pelo publisher.
Implementação:
- Publica mensagens simuladas no tópico "Bia".
- Verifica se o cliente MQTT de teste recebe as mensagens.

Resultados Esperados:
- Todas as mensagens publicadas devem ser recebidas pelo cliente MQTT.
- A ordem das mensagens recebidas deve ser a mesma da ordem de publicação.

## Teste de Validação dos Dados:
Garante que os dados enviados pelo publisher não são alterados durante a comunicação.
Implementação:
- Compara os dados publicados com os dados originais no arquivo JSON.

Resultados Esperados:
- Os dados publicados devem ser idênticos aos dados do arquivo JSON.
- Não deve haver erros de validação ou inconsistências nos dados.

## Teste de Confirmação da Taxa de Disparo:
Garante que o publisher atende à taxa de disparo especificada.
Implementação:
- Publica um número conhecido de mensagens em um intervalo de tempo específico.
- Calcula a taxa de disparo real e a compara com a taxa esperada.
Resultados Esperados:
- A taxa de disparo real deve estar dentro da margem de erro especificada.
- O publisher deve atender às suas especificações de desempenho.

## Teste de Integração 
Verifica se a integração com o HiveMQ está sendo realizada com sucesso.
  
## Como executar 
1) Clone o repositório
2) Execute os comandos abaixo
```
go run publisher.go
```
```
go run subscriber.go
```
Para testar
```
go mod tidy
```
```
go test
```

## Video Demonstração
[Screencast from 04-03-2024 08:35:23.webm](https://github.com/Bianca-Cassemiro/modulo-9/assets/99203402/82b084f5-4668-4fe6-8456-fa200d385a6a)

## Video demonstração Kafka
[Gravação de tela de 09-04-2024 10:24:10.webm](https://github.com/Bianca-Cassemiro/modulo-9/assets/99203402/28064480-4a94-444a-8787-e90f8ce3d6c9)
