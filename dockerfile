# 源镜像
FROM golang:latest

# 作者
MAINTAINER Razil "hezhixiong"

# 设置工作目录
WORKDIR $GOPATH/src/docker/docker_hello

# 将服务器的go工程代码加入到docker容器中
COPY . $GOPATH/src/docker/docker_hello

# 构建可执行文件
RUN go build .

# 暴露端口
EXPOSE 4000

# 最终运行docker的命令
ENTRYPOINT  ["./docker_hello"]