# Prova 2
Essa atividade tem como objetivo  

# Como executar 
1) Para baixar todas as dependências.
```
go mod tidy
```
2) Vá até a kafka e execute
```
docker compose up
```
3) Vá até a pasta go e execute
```
go run main.go
```
4) Para testar, vá até a paste testes e execute
```
go test
```

# Testes e resultados
1) Para demonstrar a persistência dos dados criei um arquivo txt que adiciona as mensagens recebidas.

2) Para mostrar a integridade das informações, foi criado um teste específico na pasta testes, o qual cria um producer e um consumer e verifica se a mensagem recebida é a mesma enviada.

3) A mensagem no terminal é apresentada por parte do consumer do modo estruturado.
   
# Video demonstração
