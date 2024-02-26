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

## Vídeo Demonstração
[Screencast from 26-02-2024 11:00:46.webm](https://github.com/Bianca-Cassemiro/modulo-9/assets/99203402/d185830d-4fb7-497f-9d55-a431907167cd)

