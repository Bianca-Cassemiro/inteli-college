from fastapi import FastAPI, HTTPException
from pydantic import BaseModel
import logging

app = FastAPI()

blog_posts = []

class BlogPost(BaseModel):
    id: int
    title: str
    content: str

logging.basicConfig(filename='app.log', level=logging.WARNING,
                    format='%(asctime)s %(levelname)s %(message)s')

@app.post('/blog', status_code=201)
async def create_blog_post(post: BlogPost):
    try:
        blog_posts.append(post)
        return {"status": "success"}
    except Exception as e:
        logging.error(f"Error creating blog post: {str(e)}")
        raise HTTPException(status_code=500, detail=str(e))

@app.get('/blog')
async def get_blog_posts():
    return {"posts": [post.dict() for post in blog_posts]}

@app.get('/blog/{id}')
async def get_blog_post(id: int):
    for post in blog_posts:
        if post.id == id:
            return {"post": post.dict()}
    logging.warning(f"Post with id {id} not found")
    raise HTTPException(status_code=404, detail="Post not found")

@app.delete('/blog/{id}')
async def delete_blog_post(id: int):
    for post in blog_posts:
        if post.id == id:
            blog_posts.remove(post)
            return {"status": "success"}
    logging.warning(f"Post with id {id} not found for deletion")
    raise HTTPException(status_code=404, detail="Post not found")

@app.put('/blog/{id}')
async def update_blog_post(id: int, updated_post: BlogPost):
    try:
        for post in blog_posts:
            if post.id == id:
                post.title = updated_post.title
                post.content = updated_post.content
                return {"status": "success"}
        logging.warning(f"Post with id {id} not found for update")
        raise HTTPException(status_code=404, detail="Post not found")
    except Exception as e:
        logging.error(f"Error updating blog post: {str(e)}")
        raise HTTPException(status_code=500, detail=str(e))

