# API (FastAPI)

Minimal FastAPI backend for demos.

Quick start (PowerShell):

```powershell
cd api
python -m venv .venv
# Activate the venv (PowerShell)
.\.venv\Scripts\Activate.ps1
pip install -r requirements.txt
# Run the app
uvicorn app.main:app --reload --port 8000
```

Tests:

```powershell
# from api\ directory
.\.venv\Scripts\Activate.ps1
pytest -q
```

Endpoints:
- GET / -> {"message": "Hello, world"}
- GET /health -> {"status": "ok"}
- GET /items/{item_id}?q=... -> {"item_id": <int>, "q": <str|null>}
