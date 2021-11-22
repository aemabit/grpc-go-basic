### GO 

##### MAC OS

```
export GO_PATH=~/go
export PATH=$PATH:/$GO_PATH/bin
source ~/.bash_profile
```

##### Commands

- Initialize Go: 
    base -> `go mod init <MODULE_NAME>`
    sample -> `go mod init dummyFrontier/frontier-grpc`
- Go get package:
    base -> `go get <PATH>`
    sample -> `go get google.golang.org/grpc`
- Run Server: 
    base -> `go run <PATH>`
    sample -> `go run usermgmt_server/usermgmt_server.go`

### GRPC

##### Commands
- Run Protoc Script:
    `sh execProto.sh`
    `blockchain/blockchain.proto`

- Compile proto: 
    base -> `protoc --go_out=<PATH> --go_opt=paths=source_relative --go-grpc_out=<PATH> --go-grpc_opt=paths=source_relative <PATH-PROTO-FILE>`
    sample -> `protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative userpb/userpb.proto`

- Compile proto with `option go_package="/userpb;userpb"`:
    base -> `protoc --go_out=<PATH> --go-grpc_out=<PATH> <PATH-PROTO-FILE>`
    sample -> `protoc --go_out=./pb --go-grpc_out=./pb proto/user.proto`
