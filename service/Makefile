build:
	protoc --proto_path=. --micro_out=. --go_out=. proto/demo/demo.proto
	GOOS=linux GOARCH=amd64 go build
	docker build -t laracom-service .

run:
	docker run -d -p 9091:9091 -e MICRO_SERVER_ADDRESS=:9091 -e MICRO_REGISTRY=mdns laracom-service