protoc -I=proto --go_out=plugins=grpc:pb proto/data.proto
python -m grpc_tools.protoc -I=proto --python_out=pb --grpc_python_out=pb proto/data.proto