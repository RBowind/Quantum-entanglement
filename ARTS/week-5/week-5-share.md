## Supervisor 帮你解决进程崩溃烦恼

supervisor是用Python开发的一个进程管理工具。可以很方便的监听、启动、停止、重启一个或多个进程。用supervisor管理的进程，当一个进程意外被杀死，supervisor监听到进程死后，会自动将它重启，很方便的做到进程自动恢复的功能，不再需要自己写shell脚本来控制。

那么进程管理又有什么场景呢？比如以守护进程方式启动的程序，比如一个后台任务、一组 Web 服务的进程、消息推送的后台服务、日志数据的处理分析服务等等。

[官网链接](http://www.supervisord.org/introduction.html)

### 安装
`sudo pip install supervisor`

### 配置文件
`/etc/supervisord.conf`，如不存在自己新建。
```bash
[unix_http_server]
;file=/tmp/supervisor.sock   ; (the path to the socket file)
;修改为 /home/supervisor 目录，避免被系统删除
file=/home/supervisor/supervisor.sock   ; (the path to the socket file)
;chmod=0700                 ; socket file mode (default 0700)
;chown=nobody:nogroup       ; socket file uid:gid owner
;username=user              ; (default is no username (open server))
;password=123               ; (default is no password (open server))

;[inet_http_server]         ; inet (TCP) server disabled by default
;port=127.0.0.1:9001        ; (ip_address:port specifier, *:port for ;all iface)
;username=user              ; (default is no username (open server))
;password=123               ; (default is no password (open server))
...

[supervisord]
;logfile=/tmp/supervisord.log ; (main log file;default $CWD/supervisord.log)
;修改为 /var/log 目录，避免被系统删除
logfile=/var/log/supervisor/supervisord.log ; (main log file;default $CWD/supervisord.log)
; 日志文件多大时进行分割
logfile_maxbytes=50MB        ; (max main logfile bytes b4 rotation;default 50MB)
; 最多保留多少份日志文件
logfile_backups=10           ; (num of main logfile rotation backups;default 10)
loglevel=info                ; (log level;default info; others: debug,warn,trace)
;pidfile=/tmp/supervisord.pid ; (supervisord pidfile;default supervisord.pid)
;修改为 /home/supervisor 目录，避免被系统删除
pidfile=/home/supervisor/supervisord.pid ; (supervisord pidfile;default supervisord.pid)
...
;设置启动supervisord的用户，一般情况下不要轻易用root用户来启动，除非你真的确定要这么做
;user=chrism                 ; (default is current user, required if root)
...

[supervisorctl]
; 必须和'unix_http_server'里面的设定匹配
;serverurl=unix:///tmp/supervisor.sock ; use a unix:// URL  for a unix socket
;修改为 /home/supervisor 目录，避免被系统删除
serverurl=unix:///home/supervisor/supervisor.sock ; use a unix:// URL  for a unix socket
;serverurl=http://127.0.0.1:9001 ; use an http:// url to specify an inet socket
;username=chris              ; should be same as http_username if set
;password=123                ; should be same as http_password if set
...
```

