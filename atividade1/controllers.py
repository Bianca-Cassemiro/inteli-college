# api/controllers.py
from api.models import User, Task

# Lista simulada de usuários e tarefas
users = []
tasks = []

def authenticate_user(username, password):
    for user in users:
        if user.username == username and user.password == password:
            return user
    return None

def create_user(username, password):
    # Verifica se o usuário já existe
    for user in users:
        if user.username == username:
            return None
    # Cria um novo usuário
    new_user = User(username, password)
    users.append(new_user)
    return new_user

def get_all_tasks():
    return tasks

def get_task_by_id(id):
    for task in tasks:
        if task.id == id:
            return task
    return None

def create_task(title, description, user_id):
    new_task = Task(title, description, user_id)
    tasks.append(new_task)
    return new_task

def update_task(id, title, description):
    for task in tasks:
        if task.id == id:
            task.title = title
            task.description = description
            return task
    return None

def delete_task(id):
    for task in tasks:
        if task.id == id:
            tasks.remove(task)
            return True
    return False
