from fastapi import APIRouter, Depends 
import sqlite3
from models.user_model import User
from passlib.context import CryptContext
from auth.auth import verify_password, get_password_hash, generate_token, oauth2_scheme, validate_token
from typing_extensions import Annotated

pwd_context = CryptContext(schemes=["bcrypt"], deprecated="auto")
router = APIRouter()

@router.post('/add_user')
async def add_user(new_user: User):
    connection = sqlite3.connect("banco.db")
    cursor = connection.cursor()
    username = new_user.username
    password = get_password_hash(new_user.password)
    cursor.execute("SELECT * FROM users WHERE username = ?", (username,))
    existing_user = cursor.fetchone()
    if existing_user:
        return {"status": "failed user already exists"}
    cursor.execute("INSERT INTO users (username, password) VALUES (?, ?)", (username, password))
    connection.commit()
    return {"status": f"success creating user {username}"}

@router.post('/login')
async def login(user_info: User):
    connection = sqlite3.connect("banco.db")
    cursor = connection.cursor()
    username = user_info.username
    password = user_info.password
    
    cursor.execute("SELECT * FROM users WHERE username = ?", (username,))
    user = cursor.fetchone()
    if not user:
        return {"status": "failed user not found"}
    if not verify_password(password, user[2]):
        return {"status": "failed password incorrect"}
    data = {"username": user[1]}
    token = generate_token(data)
    return {'token': token}

@router.get('/get_user')
async def get_user(token: str = Depends(oauth2_scheme)):
    connection = sqlite3.connect("banco.db")
    cursor = connection.cursor()
    username = validate_token(token)
    return {"user": username}
