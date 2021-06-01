# micro

## install micro
```
go get -v -u github.com/micro/micro
go get -v -u github.com/micro/go-micro
go get -v -u github.com/micro/go-micro/v2
```

## install protoc
```
https://github.com/protocolbuffers/protobuf/releases
```

## install protoc-gen-go
```
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
```

## install protoc-gen-micro
```
go get -u github.com/micro/protoc-gen-micro
go get -u github.com/micro/protoc-gen-micro/v2
```

## create service
```
micro new micro
```

## generate proto golang
```
cd micro
protoc --proto_path=proto --go_out=pb --micro_out=pb proto/*.proto
```

## start server
```
go run server.go
```

## check server
```
micro list services
```

## start restful api
```
micro api --namespace=go.micro.srv
```

## start etcd
```
#stop firewall
systemctl stop firewalld
systemctl disable firewalld
systemctl status  firewalld
#start etcd
./etcd --listen-client-urls http://0.0.0.0:2379 --advertise-client-urls http://0.0.0.0:2379 --listen-peer-urls http://0.0.0.0:2380
```

## start consul
```
#stop firewall
systemctl stop firewalld
systemctl disable firewalld
systemctl status  firewalld
#start consul(-dev & -server)
./consul agent -dev -client 0.0.0.0 -ui
```

```
#cluster
consul agent -server -bootstrap-expect=3 -data-dir=/tmp/consul -node=10.200.110.91 -bind=10.200.110.91 -client=0.0.0.0 -datacenter=origin -ui
consul agent -server -bootstrap-expect=3 -data-dir=/tmp/consul -node=10.200.110.92 -bind=10.200.110.92 -client=0.0.0.0 -datacenter=origin -ui
consul agent -server -bootstrap-expect=3 -data-dir=/tmp/consul -node=10.200.110.93 -bind=10.200.110.93 -client=0.0.0.0 -datacenter=origin -ui
```
- server run as server
- bootstrap-expect cluster number
- data-dir data dir
- node node id
- bind listen ip, default 0.0.0.0
- client client ip, remote use 0.0.0.0
- ui view ui
- config-dir config dir
- datacenter center name
