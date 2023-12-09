
build_server: 
	GOARCH=amd64 GOOS=windows go build -o ./cmd/server/server ./cmd/server


build_client: 
	GOARCH=amd64 GOOS=windows go build -o ./cmd/client/client ./cmd/client

run: build_server build_client
	./cmd/client/client
	./cmd/server/server

