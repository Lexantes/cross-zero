# cross-zero

Команда для билда находясь в корне

```bash
go build -buildmode=pie -x -trimpath -ldflags="-s -w" -o cross-zero.exe ./cmd/main.go
```