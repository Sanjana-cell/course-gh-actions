from fastapi.testclient import TestClient
from app.main import app

client = TestClient(app)


def test_root():
    r = client.get("/")
    assert r.status_code == 200
    assert r.json() == {"message": "Hello, world"}


def test_health():
    r = client.get("/health")
    assert r.status_code == 200
    assert r.json().get("status") == "ok"


def test_item():
    r = client.get("/items/42?q=test")
    assert r.status_code == 200
    assert r.json()["item_id"] == 42
    assert r.json()["q"] == "test"
