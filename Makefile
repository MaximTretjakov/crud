generate:
	protoc -I . --grpc-gateway_out=logtostderr=true:. --swagger_out=allow_merge=true,merge_file_name=api:. --go_out=. --go-grpc_out=. crud.proto