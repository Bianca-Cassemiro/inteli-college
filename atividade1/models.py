# api/models.py
class User:
    def __init__(self, id, username, password):
        self.id = id
        self.username = username
        self.password = password

class Task:
    def __init__(self, id, title, description, user_id):
        self.id = id
        self.title = title
        self.description = description
        self.user_id = user_id
