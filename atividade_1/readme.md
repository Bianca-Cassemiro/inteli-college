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
Após isso criei um script publisher.py e configurei para realizar a leitura dos dados estabelecidos no Json e publica-los em um tópico chamado "Bia"
Criei um subscriber.py que se inscrevia no tópico Bia e realizava a leitura do que estava sendo publicado.

## Como executar 
1) Clone o repositório
2) Execute os comandos abaixo
```
python3 publisher.py
```
```
python3 subscriber.py
```

## Vídeo Demonstração
[Screencast from 18-02-2024 22:00:36.webm](https://github.com/Bianca-Cassemiro/modulo-9/assets/99203402/09b345ba-9b5c-42f4-af92-cf2bbc0b01ae)
