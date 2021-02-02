module go-micro-etcd

go 1.15

require (
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5
	github.com/asim/go-micro/plugins/registry/etcd/v3 v3.0.0-20210130181356-ca3dfc4580fa
	github.com/asim/go-micro/plugins/wrapper/ratelimiter/uber/v3 v3.0.0-20210130181356-ca3dfc4580fa
	github.com/asim/go-micro/v3 v3.5.0
	github.com/coreos/bbolt v1.3.3 // indirect
	github.com/coreos/go-systemd v0.0.0-20190719114852-fd7a80b32e1f // indirect
	github.com/golang/protobuf v1.4.3
	github.com/gorilla/websocket v1.4.1 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.1.0 // indirect
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0 // indirect
	github.com/jonboulle/clockwork v0.1.0 // indirect
	github.com/json-iterator/go v1.1.8 // indirect
	github.com/soheilhy/cmux v0.1.4 // indirect
	github.com/tmc/grpc-websocket-proxy v0.0.0-20190109142713-0ad062ec5ee5 // indirect
	github.com/xiang90/probing v0.0.0-20190116061207-43a291ad63a2 // indirect
	go.etcd.io/bbolt v1.3.3 // indirect
	google.golang.org/protobuf v1.25.0
	sigs.k8s.io/yaml v1.1.0 // indirect
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
