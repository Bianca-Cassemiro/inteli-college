# Sistema de Logs com Gateway e Serviço de Eventos

## Descrição

Este projeto implementa um sistema de logs com um gateway NGINX, um serviço de backend, um serviço de eventos, e envio de logs para o Elastic Search usando Filebeat.

## Componentes

1. **NGINX**: Utilizado como gateway para rotear as solicitações entre os serviços.
2. **Serviço de Eventos**: Implementado em Python usando Flask com operações CRUD.
3. **Filebeat**: Utilizado para enviar logs dos serviços para o Elastic Search.
4. **Elasticsearch**: Utilizado para armazenar e buscar logs.
5. **Kibana**: Utilizado para visualizar os logs armazenados no Elasticsearch.

## Estrutura do Projeto
```
├── backend/
├── filebeat/
├── elasticsearch/
├── kibana/
├── events_service/
│ ├── Dockerfile
│ ├── app.py
│ └── requirements.txt
├── logs/
├── nginx/
│ └── nginx.conf
└── docker-compose.yml
```

## Configuração

### NGINX

1. O NGINX é configurado para rotear as solicitações para os serviços backend e events_service.
2. Arquivo de configuração: `nginx/nginx.conf`.

### Backend

1. O serviço backend está configurado para escutar na porta 8000.
2. Os logs do backend são armazenados no diretório `logs`.

### Serviço de Eventos

1. Implementado em Python usando Flask, este serviço está configurado para escutar na porta 5000.
2. Os logs do serviço de eventos são armazenados no diretório `logs`.

### Filebeat

1. O Filebeat é configurado para monitorar os arquivos de log dos serviços e enviá-los para o Elasticsearch.
2. Arquivo de configuração: `filebeat/filebeat.yml`.

### Elasticsearch e Kibana

1. Elasticsearch está configurado para escutar nas portas 9200 e 9300.
2. Kibana está configurado para escutar na porta 5601 e depende do Elasticsearch.

## Executando o Sistema

1. **Execute os Conteiners**:
```
   docker-compose up --build
```
