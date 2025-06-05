import json
import jwt

JWT_ALGORITHM = "HS256"
JWT_SECRET = "mysecret"

def token_response(token: str):
    return json.dumps({"token de acesso": token})

def signJWT(userId : int) -> str:
    payload = { 
        "user" : userId,}
    token = jwt.encode(payload, JWT_SECRET, algorithm = JWT_ALGORITHM)
    return token_response(token)

def decodeJWT(token: str):
    decode_token = jwt.decode(token, JWT_SECRET, algorithms=JWT_ALGORITHM)
    return decode_token
   