# Atividade prática - prova 1

Durante essa atividade foi desenvolvida uma API que permite gerenciar pedidos com funcionalidades de cadastro, listagem, obtenção por ID, edição e exclusão.

## Identificação do estudante
- Nome: Bianca Cassemiro Lima
- Turma: Eng. Comp. - módulo 10

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
![image](https://github.com/Bianca-Cassemiro/modulo-10/assets/99203402/faeb7b92-074c-4345-bd2d-b195563cf620)

### Get /pedidos
![image](https://github.com/Bianca-Cassemiro/modulo-10/assets/99203402/10d60023-62d7-40bb-91db-c7543c6ab83f)

### Get /pedidos/id 
![image](https://github.com/Bianca-Cassemiro/modulo-10/assets/99203402/0aea14ac-a6e4-4b7d-9214-96210df6f5c7)

### Put /pedidos/id
![image](https://github.com/Bianca-Cassemiro/modulo-10/assets/99203402/32066901-c007-4be6-83a1-9d890abdea22)

### Delete /pedidos/id 
![image](https://github.com/Bianca-Cassemiro/modulo-10/assets/99203402/700a1e6e-3d3e-4d81-bc10-892f550afbba)


## Rotas 

- `POST/novo`: Cadastrar um novo pedido. 
- `GET /pedidos`: Retorna todos os pedidos cadastrados.
- `GET /pedidos/<id>`: Retorna o pedido do ID fornecido. 
- `PUT /pedidos/<id>`: Atualiza o pedido do ID fornecido.
- `DELETE /pedidos/<id>`: Exclui o pedido do ID fornecido.



