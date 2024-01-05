
run: certificates gen_proto build_server build_client
	./server

certificates: 
	go run ./cmd/cert/main.go

migrate: build_migrate
	./migrator  --migrations-path=./migrations

build_migrate:
	go build -o . ./cmd/migrator

gen_proto:
	protoc --go_out=./proto/gen/keeper --go_opt=paths=source_relative --go-grpc_out=./proto/gen/keeper --go-grpc_opt=paths=source_relative proto/goph_keeper.proto 

build_server: 
	GOARCH=amd64 GOOS=windows go build -o . ./cmd/server

build_client: 
	GOARCH=amd64 GOOS=windows go build -o . ./cmd/client

run_client: build_client
	./client


gen_mocks:
	mockgen -destination=mocks/mock_keeperClient -package=mocks github.com/DEHbNO4b/practicum_project2/proto/gen/keeper/proto GophKeeperClient