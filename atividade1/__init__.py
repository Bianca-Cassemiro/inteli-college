# api/__init__.py
from flask import Flask

# Criar a instância do aplicativo Flask
app = Flask(__name__)

# Importar os módulos de rotas após a criação do aplicativo para evitar ciclos de importação
from api import routes
