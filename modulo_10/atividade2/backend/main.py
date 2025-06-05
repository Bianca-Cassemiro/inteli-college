from fastapi import FastAPI
from router.task_routs import router as tasks_router
from router.user_routs import router as user_router

import uvicorn

app = FastAPI()

app.include_router(user_router, tags=["users"])
app.include_router(tasks_router,tags=["tasks"])

@app.get("/")
def read_root():
    return {"Health": "ok"}

if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=5000)