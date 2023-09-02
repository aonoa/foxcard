
## 一个使用redis做存储的开源的卡密服务
#### 起了个头，凑合着也能用
#### 运行起来后，接口地址在 http://127.0.0.1:8000/q/swagger-ui#/
```api
POST
​/card​/login   (卡密认证)

POST
​/card​/heartbeat (每分钟发一次心跳)

POST
​/card​/logout   (退出，不主动退出的的话3分钟后也会自动退)
```

#### PS:(有人用的话我就继续更，没人用的话我就懒得继续写了，毕竟我也不需要这个，不清楚需求，有问题或者建议提issus)


# Kratos Project Template

## Install Kratos
```
go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
```
## Create a service
```
# Create a template project
kratos new server

cd server
# Add a proto template
kratos proto add api/server/server.proto
# Generate the proto code
kratos proto client api/server/server.proto
# Generate the source code of service by proto file
kratos proto server api/server/server.proto -t internal/service

go generate ./...
go build -o ./bin/ ./...
./bin/server -conf ./configs
```
## Generate other auxiliary files by Makefile
```
# Download and update dependencies
make init
# Generate API files (include: pb.go, http, grpc, validate, swagger) by proto file
make api
# Generate all files
make all
```
## Automated Initialization (wire)
```
# install wire
go get github.com/google/wire/cmd/wire

# generate wire
cd cmd/server
wire
```

## Docker
```bash
# build
docker build -t <your-docker-image-name> .

# run
docker run --rm -p 8000:8000 -p 9000:9000 -v </path/to/your/configs>:/data/conf <your-docker-image-name>
```


## 授权许可
本项目采用 MIT 开源授权许可证，完整的授权说明已放置在 [LICENSE](LICENSE) 文件中。