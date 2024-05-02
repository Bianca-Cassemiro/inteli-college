from fastapi import APIRouter, Depends
import sqlite3
from models.task_model import Task
from router.user_routs import getUser
from typing import List

router = APIRouter()

@router.post('/add_task')
async def add_task(new_task: Task, user_token: dict = Depends(getUser)):
    connection = sqlite3.connect("banco.db")
    cursor = connection.cursor()
    title = new_task.title
    content = new_task.content
    if not content:
        return {"status": "failed content is required"}
    username = user_token['username']
    cursor.execute("SELECT * FROM user_tasks WHERE title = ? AND username = ?", (title, username))
    existing_task = cursor.fetchone()
    if existing_task:
        return {"status": "failed task already exists"}
    cursor.execute("INSERT INTO user_tasks (title, content, username) VALUES (?, ?, ?)", (title, content, username,))
    connection.commit()
    return {"status": "success creating task with title: " + title}

@router.get('/get_tasks', response_model=List[Task])
async def get_tasks(user_token: dict = Depends(getUser)):
    connection = sqlite3.connect("banco.db")
    cursor = connection.cursor()
    username = user_token['username']
    cursor.execute("SELECT title, content FROM user_tasks WHERE username = ?", (username,))
    tasks = [{"title": task[0], "content": task[1]} for task in cursor.fetchall()]
    return tasks

@router.delete('/delete_task')
async def delete_task(task: Task, user_token: dict = Depends(getUser)):
    connection = sqlite3.connect("banco.db")
    cursor = connection.cursor()
    username = user_token['username']
    title = task.title
    cursor.execute("SELECT * FROM user_tasks WHERE title = ? AND username = ?", (title, username,))
    existing_task = cursor.fetchone()
    if not existing_task:
        return {"status": "failed task not found"}
    cursor.execute("DELETE FROM user_tasks WHERE title = ? AND username =?", (title,username))
    connection.commit()
    return {"status": "success deleting task with title: " + title}

@router.put('/update_task')
async def update_task(task: Task, user_token: dict = Depends(getUser)):
    connection = sqlite3.connect("banco.db")
    cursor = connection.cursor()
    username = user_token['username']
    title = task.title
    content = task.content
    cursor.execute("SELECT * FROM user_tasks WHERE title = ? AND username = ?", (title, username))
    existing_task = cursor.fetchone()
    if not existing_task:
        return {"status": "failed task not found"}
    cursor.execute("UPDATE user_tasks SET content = ? WHERE title = ? and username = ?", (content, title,username))
    connection.commit()
    return {"status": "success updating task with title: " + title}
