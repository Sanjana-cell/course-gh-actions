from typing import Optional
from fastapi import FastAPI

app = FastAPI(title="api-demo")


@app.get("/")
def read_root():
    return {"message": "Hello, world"}


@app.get("/health")
def health():
    return {"status": "ok"}


@app.get("/items/{item_id}")
def read_item(item_id: int, q: Optional[str] = None):
    return {"item_id": item_id, "q": q}
