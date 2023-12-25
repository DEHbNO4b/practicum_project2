
gen:
	protoc --go_out=./proto/gen --go_opt=paths=source_relative --go-grpc_out=./proto/gen --go-grpc_opt=paths=source_relative proto/goph_keeper.proto 

build_server: 
	GOARCH=amd64 GOOS=windows go build -o . ./cmd/server

build_client: 
	GOARCH=amd64 GOOS=windows go build -o . ./cmd/client

run: gen build_server build_client
	./server

build_migrate:
	go build -o . ./cmd/migrator

migrate: build_migrate
	./migrator  --migrations-path=./migrations
