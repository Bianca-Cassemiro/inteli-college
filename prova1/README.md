# Atividade prática - prova 1

Durante essa atividade foi desenvolvida uma API que permite gerenciar pedidos com funcionalidades de cadastro, listagem, obtenção por ID, edição e exclusão.

## Como executar

1. Clone o repositório.
2. Navegue até a pasta do projeto.
3. Construa a imagem Docker:
    ```
    docker build -t prova .
    ```
4. Execute o container:
    ```
    docker run -p 5000:5000 prova
    ```

## Testes de funcionamento
### Post /novo 
### Get /novo 
### Get /novo/id 
### Put /novo/id
### Delete /novo/id 

## Rotas 

- `POST/novo`: Cadastrar um novo pedido. 
- `GET /pedidos`: Retorna todos os pedidos cadastrados.
- `GET /pedidos/<id>`: Retorna o pedido do ID fornecido. 
- `PUT /pedidos/<id>`: Atualiza o pedido do ID fornecido.
- `DELETE /pedidos/<id>`: Exclui o pedido do ID fornecido.



