# Atividade prática - prova 1


## Identificação do estudante
- Nome: Bianca Cassemiro Lima
- Turma: Eng. Comp. - módulo 10

## Como configurar o ambiente virtual 

1. Crie um ambiente virtual:
    ```sh
    python -m venv venv
    ```

2. Ative o ambiente virtual:
    ```sh
    source venv/bin/activate 
    ```

3. Instale as dependências:
    ```sh
    pip install
    ```
## Como executar a aplicação

Para rodar a aplicação localmente:
- Instale as dependências:
```sh
    pip install -r requirements.txt
```

```sh
    uvicorn main:app --reload
```
Para rodar a aplicação com docker:
```
docker-compose up --build
```

## Rotas definidas

- POST /blog para criar um novo post.
- GET /blog para obter todos os posts.
- GET /blog/{id} para obter um post específico.
- PUT /blog/{id} para atualizar um post.
- DELETE /blog/{id} para deletar um post.

## Evidências de teste
- Arquivo json do collection contendo as rotas no postman
- Prints de telas com o resultado
https://docs.google.com/document/d/18DCx2NRJ9t0DWVN_6hiTiMogrvXK8RUeqG5bBzCHoOg/edit
