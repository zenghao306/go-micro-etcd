# micro-service
go语言go-micro微服务工程

## Getting Started

- [教程](https://micro.mu/docs/helloworld.html)
- [用户服务](https://github.com/micro-in-cn/tutorials/tree/master/microservice-in-micro/part1)


## 项目目录说明
本项目基于github.com/asim/go-micro/v3构造，苦于官方给出例子较少，故分享出来。

目前的项目结构看起来像是一个独立的项目，实际真实的项目结构不是这样的，本项目为了简单使用而故意为之。
实际情况是`go-micro-etce`下的client/server目录都是一个独立的微服务，需要独立部署在不同的服务器上，每个服务都需要自己
的`go.mod`,`proto`等文件和目录。

## 依赖

grpc v1.25.1，
v1.27.1会有类型这样的错误`undefined: resolver.BuildOption`,`undefined: resolver.ResolveNowOption`

## etcd
```etcd
先启动etcd程序
注意etcd端口，在项目client/main.go和server/main.go中填上etcd的完整路径（默认写的是本机的）
生成*.pb.go代码
```shell
$ protoc --proto_path=$GOPATH/src:. --micro_out=. --go_out=. *.proto
```
### 测试client
```shell
go run server/main.go
```
```shell
go run client/main.go
```

### 测试api服务
```shell
连上后client端会自动向服务端发送一次g rpc请求，并输出服务端返回结果
```
返回返回结果: 200 ,return uid: 1
item_id:"12345"  item_scene_id:1  item_type_id:1  item_hot_rank:0.5  item_name:"test"  item_info:"test"  item_price:99  item_tags:"test"  item_modif_time:1610096132  item_create_timestamp:1610096132  item_city:"深圳"  item_create_userid:"12345"  reserve_property_1:"ReserveProperty_1"  reserve_property_2:"RerveProperty_2"  reserve_property_3:"ReserveProperty_3"  reserve_property_4:"ReserveProperty_4"
