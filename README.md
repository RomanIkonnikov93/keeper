# keeper

The keeper allows you to safely store private information on a remote server.

## Such as: 
* `credentials (login, password)`
* `payment card data`
* `any binary data (maximum size for one record 2Mb)`

# Usage:

### Repository:
The `PostgeSQL` database is used as a repository.

To get started, you need to create a new empty PostgeSQL database, the migration mechanism occurs automatically when the server starts.

The table structure is in the directory : `keeper/server/migrations`

### Server:
This application uses `gRPC` to transfer data between client and server.

Server configuration file `keeper/server/internal/config/config.go`
```go
type Config struct {
DSN          string `env:"KEEPER_DSN"` // "postgres://login:password@address:port/databasename"
GRPCAddress  string `env:"GRPC_PORT" envDefault:":3200"` // server address
JWTSecretKey string `env:"JWT_SECRET_KEY"` // needed for JSON Web Token
SecretKey    string `env:"KEEPER_SECRET_KEY"` // 32 byte - needed to AES encrypt private data
}
```

You need to compile the server binary while in the directory: `keeper/server`
```
Execute: go build ./cmd/keeper/main.go
```

### Client:
The application client implements a terminal user interface for interacting with the server.

Client configuration file `keeper/client/internal/config/config.go`
```go
type Config struct {
ServerAddress     string `env:"KEEPER_SERVER_ADDRESS" envDefault:":3200"` // server connection address
DownloadFilesPath string `env:"KEEPER_DOWNLOAD_FILES_PATH" envDefault:""` // directory for downloading binary data, if this variable is empty, the download will take place in the root directory of the application
}
```

You need to compile the client binary while in the directory: `keeper/client`
```
Execute: go build ./cmd/keeper/main.go
```


