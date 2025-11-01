# go-api

Minimal Go HTTP server for demos.

Run locally (PowerShell):

```powershell
cd go-api
# Run directly (requires Go installed and on PATH)
go run .

# Build binary
go build -o server .
./server
```

Endpoints:
- GET / -> {"message":"Hello, world"}
- GET /health -> {"status":"ok"}
- GET /items/{id}?q=... -> {"item_id": <int>, "q": <string|null>}

Tests:

```powershell
cd go-api
go test ./... -v
```

If `go` isn't installed, download from https://golang.org/dl and ensure `go` is on your PATH.


