# ARTS-Week4

## Algorithm
**4 Sums**
### problem:
Given an array nums of n integers and an integer target, are there elements a, b, c, and d in nums such that a + b + c + d = target? Find all unique quadruplets in the array which gives the sum of target.

> Note:The solution set must not contain duplicate quadruplets.

### Solution:
Base on the solution of sum3.
 [code](https://github.com/RBowind/RBlog/blob/master/ARTS/week-4/sum4.go)

## Review
I believe that most of us are 996 because of their ineffectiveness. We can't focus on the main task . We are interrupted all the time. Why? We should find the solutions from ourselves.
[The Average Employee Works 3 Hours Out Of Every 8](https://medium.com/swlh/the-average-employee-works-3-hours-out-of-every-8-135f2f042268)


## Tip
这里想分享一些docker命令

查看所有镜像
`docker images`

查看运行容器
`docker ps`

查看所有容器
`docker ps -a`

进入容器其中字符串为容器ID:
`docker exec -it d27bd3008ad9 /bin/bash`

停用全部运行中的容器:
`docker stop $(docker ps -q)`

删除全部容器：
`docker rm $(docker ps -aq)`

一条命令实现停用并删除容器：
`docker stop $(docker ps -q) & docker rm $(docker ps -aq)`

删除none的镜像
`docker rmi $(docker images | grep "^<none>" | awk "{print $3}")`

其实就是些shell 变量的小技巧


## Share
[《Docker 搭建 Golang+Nginx+Mysql 开发环境》](https://github.com/RBowind/RBlog/blob/master/ARTS/week-4/week-4-share.md)


