## Docker 搭建 Golang+Nginx+Mysql 开发环境

本文主要描述笔者在MacOS环境下使用Docker搭建Golang+Nginx+MySQL的经历。

### 初衷
一开始使用Docker的目的很简单：可移植性以及环境隔离性。

### 思路
一共三个容器：
- nginx
- golang
- MySQL

1. Golang开启web服务器，监听8080端口
2. Nginx 监听80端口，反向代理到8080端口
3. MySQL 映射宿主机3306端口，通过数据持久化映射到本地。

### 步骤
### 1. 先配置 Golang 的运行环境
先在$GOPATH下写一个简单的 web demo
```golang
package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}

```
然后在与此同一目录下新建 Dockerfile 如下：

```Dockerfile
FROM golang:1.12 as build
MAINTAINER  RBowind "rbowind@163.com"

# 设置环境变量，开启 go mod，设置代理
ENV GOPROXY https://goproxy.io
ENV GO111MODULE on

WORKDIR /go/cache

# go module 包管理工具，需要go1.11版本后才支持，强烈推荐。
COPY go.mod .
COPY go.sum .
RUN go mod download

WORKDIR /go/release

ADD . .

# -w 禁止生成debug信息,注意使用该选项后，无法使用 gdb 进行调试
# -s 禁用符号表
# 可以使用 go tool link --help 查看 ldflags 各参数含义 ，下面可以减小编译后可执行程序的大小
# GOOS=linux 是将交叉编译的目标设置为Linux，这样你在Mac或者Win下也不会出现问题。 -installsuffix cgo 是为了在静态编译中导入net
# -installsuffix :在软件包安装的目录中增加后缀标识，以保持输出与默认版本分开
# -o : 指定编译后的可执行文件名称


RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -installsuffix nocgo -o app .

FROM scratch as prod

COPY --from=build /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
COPY --from=build /go/release/app /

 
CMD ["/app"]
```
之后运行:
- `docker build -t app .`
- `docker run --name app -p 8080:8080 -d app`

此时访问 http://127.0.0.1:8080 会有你想要的结果。

### 2.配置Nginx
nginx 主要是配置文件的编辑，因为`/etc/nginx/nginx.conf`中有一句`include /etc/nginx/conf.d/*.conf;`,那么主要修改位于/etc/nginx/conf.d中的conf文件。可自己先生成一个原始的nginx容易，然后把配置文件同步到本地，再修改。修改后default.conf:如下：
```nginx
server {
    listen       80;
    server_name  test.com;

    location / {
         
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
        proxy_set_header Host $http_host;
        proxy_set_header X-NginX-Proxy true;
        proxy_redirect     off;    
        # warning:这里须是宿主机的IP地址+端口，我这里是手动配置了固定的IP地址。若使用127.0.0.1或localhost，会因为容器与宿主机之间的网络隔离，导致502错误。
        proxy_pass  "http://192.168.1.100:8080"; 
    }

    error_page   500 502 503 504  /50x.html;
    location = /50x.html {
        root   /usr/share/nginx/html;
    }
}
```
- ` docker pull nginx`
- `docker run -p 80:80 -p 9000:9000 --name=nginx  -v $GOPATH/src/nginx/conf.d:/etc/nginx/conf.d -d nginx`


这个时候你可以试着访问下 http://127.0.0.1 ，也会有你想要的结果。


### 3. MySQL 配置
- `mkdir $HOME/Docker/mysqlData`
- `docker pull mysql`
- `docker run --name mysql -d --restart always -p 3306:3306 -e MYSQL_ROOT_PASSWORD=password -v $HOME/Docker/mysqlData:/var/lib/mysql mysql`

**warning**： MySQL8.0 用户密码都OK，同时防火墙端口也OK，但远程管理软件还是连接不上，如果报错 `Authentication plugin 'caching_sha2_password' cannot be loaded` balabala , 可执以下操作修改加密规则，具体原因自行谷歌：
- `docker exec -it mysql /bin/bash`
- `mysql -uroot -p`
- `ALTER USER 'root'@'%' IDENTIFIED WITH mysql_native_password BY 'password'`
- `FLUSH PRIVILEGES;`
- `exit`

这时候 Navicat 连上，建个库看看效果。然后把该容器删了,可以看到 $HOME/Docker/MysqlData 文件夹下数据依旧存在，再次建立新的容器，数据依旧在。


