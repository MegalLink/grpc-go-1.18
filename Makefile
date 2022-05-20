generate_greet:
	protoc -Igreet/proto --go_out=. --go_opt=module=github.com/MegalLink/grpc-go-1.18 --go-grpc_out=. --go-grpc_opt=module=github.com/MegalLink/grpc-go-1.18 greet/proto/greet.proto
	go build -o bin/greet/server ./greet/server
	go build -o bin/greet/client ./greet/client

run_server:
	./bin/greet/server

run_client:
	./bin/greet/client