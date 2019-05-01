## Supervisor 介绍与使用

supervisor是用Python开发的一个进程管理工具。可以很方便的监听、启动、停止、重启一个或多个进程。用supervisor管理的进程，当一个进程意外被杀死，supervisor监听到进程死后，会自动将它重启，很方便的做到进程自动恢复的功能，不再需要自己写shell脚本来控制。

那么进程管理又有什么场景呢？比如以守护进程方式启动的程序，比如一个后台任务、一组 Web 服务的进程、消息推送的后台服务、日志数据的处理分析服务等等。

它是通过fork/exec的方式把这些被管理的进程当作supervisor的子进程来启动，这样只要在supervisor的配置文件中，把要管理的进程的可执行文件的路径写进去即可。也实现当子进程挂掉的时候，父进程可以准确获取子进程挂掉的信息的，可以选择是否自己启动和报警。supervisor还提供了一个功能，可以为supervisord或者每个子进程，设置一个非root的user，这个user就可以管理它对应的进程。

> Supervisor只能管理非daemon的进程，也就是说Supervisor不能管理守护进程。



[官网链接](http://www.supervisord.org/introduction.html)

### 安装
使用 `sudo pip install supervisor`
> 保证版本最新

### 配置文件
`/etc/supervisord.conf`，如不存在自己新建。
配置文件见[附件](supervisord.conf)

### 一些特性
- Web界面操作(不安全)
  
    其中有段配置如下，取消注释后可以使用:
    ```
    ;[inet_http_server]         ; inet (TCP) server disabled by default
    ;port=127.0.0.1:9001        ; ip_address:port specifier, *:port for all iface
    ;username=user              ; default is no username (open server)
    ;password=123               ; default is no password (open server)
    ```

- 自定义配置
    ```
    [include]
    files = /etc/supervisorconf/*.conf
    ;files = relative/directory/*.ini
    ```

`/etc/supervisorconf`文件夹若不存在则自建，笔者使用场景是用来管理Golang的Web服务进程
```conf
[program:gowork]
GOPATH=/root/gopath

user=root
command=$GOPATH/bin/gowork
autostart=true
autorestart=true
startsecs=10
stderr_logfile=$HOME/gowork_stderr.log   //错误日志记录
stdout_logfile=$HOME/gowork_stdout.log
```
### 常用命令

**启动**
```bash
#启动
sudo /usr/bin/supervisord -c /etc/supervisord.conf

#如果报错：Another program is already listening on a port that one of our HTTP servers is configured to use.  Shut this program down first before starting supervisord.

sudo unlink /tmp/supervisor.sock
sudo unlink /var/run/supervisor.sock
# This .sock file is defined in /etc/supervisord.conf's [unix_http_server]'s file config value (default is /tmp/supervisor.sock or /var/run/supervisor.sock).

```

**查看进程状态**

`sudo supervisorctl status processName`

结果：`processName   RUNNING   pid 1954, uptime 14 days, 23:18:14`

## 管理进程
```bash
 supervisorctl status
 supervisorctl stop processname
 supervisorctl start processname
 supervisorctl restart processname
 supervisorctl reread       #重新读取配置文件
 supervisorctl update
 ```

