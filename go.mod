module go-micro-etcd

go 1.15

require (
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5
	github.com/asim/go-micro/plugins/registry/etcd/v3 v3.0.0-20210130181356-ca3dfc4580fa
	github.com/asim/go-micro/v3 v3.5.0
	github.com/golang/protobuf v1.4.3
	github.com/micro/go-micro v1.18.0 // indirect
	google.golang.org/protobuf v1.25.0
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
