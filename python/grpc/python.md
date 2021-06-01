## install gRPC
```
pip install grpcio
pip install grpcio-tools
```

## install protobuf
```
pip install protobuf
```

## proto file data.proto
```
syntax = "proto3";
package pb;

message actionRequest {
  string text = 1;
  enum Corpus {
    UNIVERSAL = 0;
    WEB = 1;
    IMAGES = 2;
    LOCAL = 3;
    NEWS = 4;
    PRODUCTS = 5;
    VIDEO = 6;
  }
  Corpus corpus = 4;
}
message actionResponse{
  string text = 1;
  int32 age = 2;
  message Result {
    string url = 1;
    string title = 2;
    repeated string snippets = 3;
  }
  repeated Result result = 3;
}

service FormatData {
  rpc DoFormat(actionRequest) returns (actionResponse){}
}
```

## protoc
```
python -m grpc_tools.protoc -I=proto --python_out=pb --grpc_python_out=pb proto/data.proto
```
