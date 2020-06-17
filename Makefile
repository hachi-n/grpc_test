clean-multi:
	rm -fr ./bin/grpc_client
	rm -fr ./bin/grpc_server
clean-server:
	rm -fr ./bin/grpc_server
clean-client:
	rm -fr ./bin/grpc_client
go-build-multi:
	go build -o ./bin/grpc_client ./cmd/grpc_client
	go build -o ./bin/grpc_server ./server/grpc_server
go-build-server:
	go build -o ./bin/grpc_server ./server/grpc_server
go-build-client:
	go build -o ./bin/grpc_client ./cmd/grpc_client
build-multi:
	$(MAKE) clean-multi
	$(MAKE) go-build-multi
build-server:
	$(MAKE) clean-server
	$(MAKE) go-build-server
build-client:
	$(MAKE) clean-client
	$(MAKE) go-build-client
run:
	$(MAKE) build
	$(MAKE) go-run
go-run:
	sudo ./bin/grpc_server
pb-build:
	rm -fr ./pb/auto_gen/*
	protoc -I./pb/src --go_out=plugins=grpc:./pb/auto_gen ./pb/src/*.proto
pb-doc:
	rm -fr ./proto_doc/*
	protoc --doc_out=html,index.html:./proto_doc/ pb/src/*.proto

build:
	$(MAKE) build-multi
