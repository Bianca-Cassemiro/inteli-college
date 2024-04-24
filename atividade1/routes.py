# api/routes.py
from flask import request, jsonify
from api import app
from api.controllers import *

@app.route('/')
def index():
    return 'API is running'

# Rotas para autenticação de usuários
@app.route('/login', methods=['POST'])
def login():
    data = request.json
    if 'username' not in data or 'password' not in data:
        return jsonify({'error': 'Missing username or password'}), 400
    
    user = authenticate_user(data['username'], data['password'])
    if user:
        return jsonify({'message': 'Login successful', 'user_id': user.id}), 200
    else:
        return jsonify({'error': 'Invalid username or password'}), 401

@app.route('/register', methods=['POST'])
def register():
    data = request.json
    if 'username' not in data or 'password' not in data:
        return jsonify({'error': 'Missing username or password'}), 400
    
    user = create_user(data['username'], data['password'])
    if user:
        return jsonify({'message': 'User created successfully', 'user_id': user.id}), 201
    else:
        return jsonify({'error': 'Username already exists'}), 409

# Rotas para operações CRUD de tarefas
@app.route('/tasks', methods=['GET'])
def get_tasks():
    return jsonify([task.__dict__ for task in get_all_tasks()]), 200

@app.route('/tasks/<int:id>', methods=['GET'])
def get_task(id):
    task = get_task_by_id(id)
    if task:
        return jsonify(task.__dict__), 200
    else:
        return jsonify({'error': 'Task not found'}), 404

@app.route('/tasks', methods=['POST'])
def create_task():
    data = request.json
    if 'title' not in data or 'description' not in data or 'user_id' not in data:
        return jsonify({'error': 'Missing required fields'}), 400
    
    task = create_task(data['title'], data['description'], data['user_id'])
    return jsonify(task.__dict__), 201

@app.route('/tasks/<int:id>', methods=['PUT'])
def update_task(id):
    data = request.json
    if 'title' not in data or 'description' not in data:
        return jsonify({'error': 'Missing required fields'}), 400
    
    task = update_task(id, data['title'], data['description'])
    if task:
        return jsonify(task.__dict__), 200
    else:
        return jsonify({'error': 'Task not found'}), 404

@app.route('/tasks/<int:id>', methods=['DELETE'])
def delete_task(id):
    success = delete_task(id)
    if success:
        return jsonify({'message': 'Task deleted successfully'}), 200
    else:
        return jsonify({'error': 'Task not found'}), 404
