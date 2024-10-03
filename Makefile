.PHONY: gen del

gen:
	protoc -I=proto \
      --go_out=proto/gen --go_opt=paths=source_relative \
      --go-grpc_out=proto/gen --go-grpc_opt=paths=source_relative \
      proto/*.proto

del:
	for /r %%x in (*.pb.go) do del "%%x"

