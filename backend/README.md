## How to run

### Fill up environment variables using `.env.example` file

### Generate entgo database schemas

```bash
go run -mod=mod entgo.io/ent/cmd/ent generate ./ent/schema
```

### Reset then migrate database

```bash
go run . migrate reset
```

### Run the server

```bash
go run . server
```

## Enable hot reloading

1. Install air `go install github.com/cosmtrek/air@latest`
2. Start the server with air `air server`
