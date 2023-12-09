
gen_proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./proto/goph_keeper.proto 

build_server: 
	GOARCH=amd64 GOOS=windows go build -o ./cmd/server/server ./cmd/server

build_client: 
	GOARCH=amd64 GOOS=windows go build -o ./cmd/client/client ./cmd/client

run: gen_proto build_server build_client
	./cmd/client/client
	./cmd/server/server

