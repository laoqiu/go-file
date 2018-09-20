# Go File [![License](https://img.shields.io/:license-apache-blue.svg)](https://opensource.org/licenses/Apache-2.0) [![GoDoc](https://godoc.org/github.com/micro/go-file?status.svg)](https://godoc.org/github.com/micro/go-file)

forked from asim/go-file

主要是在原项目上增加上传、查看目录等功能

## 使用


### Server

```go
import "github.com/laoqiu/go-file"

service := micro.NewService(
	micro.Name("go.micro.srv.file"),
)

proto.RegisterFileHandler(service.Server(), file.NewHandler("/tmp"))

service.Init()
service.Run()
```

### Client

```go
import "github.com/laoqiu/go-file"

// use new service or default client
service := micro.NewService()
service.Init()

client := file.NewClient("go.micro.srv.file", service.Client())
client.Download("remote.file", "local.file")
```

## Hand Wavy Bench

Local hand wavy benchmarks for rough estimates on transfer speed

size	|	time taken
----	|	----------
1mb	|	15.590542ms
8mb	|	75.184788ms
64mb	|	516.236417ms
128mb	|	1.141906576s
1024mb	|	9.794891634s

Using connection pooling and caching selector

size	|	time taken
----	|	----------
1mb	|	13.521179ms
8mb	|	53.160487ms
64mb	|	415.388025ms
128mb	|	889.409332ms
512mb	|	4.177052391s
1024mb	|	8.347038098s
