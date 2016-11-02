CENTOS7 建立私有仓库 docker

一、准备

地址规划：
Docker私有仓库地址：10.0.2.8
Docker客户端地址：10.0.2.6

1、关闭本地防火墙并设置开机不自启动
# systemctl stop firewalld.service
# systemctl disable firewalld.service

2、关闭本地selinux防火墙
  	
# sed -i 's,SELINUX=enforcing,SELINUX=disabled,g' /etc/sysconfig/selinux 
# setenforce 0   


二、安装

1、安装docker
# yum install docker
# yum upgrade device-mapper-libs
# service docker start
# chkconfig docker on

2、本地私有仓库registry
[root@localhost ~]# docker pull registry
Trying to pull repository docker.io/registry ...
24dd746e9b9f: Download complete 
706766fe1019: Download complete 
a62a42e77c9c: Download complete 
2c014f14d3d9: Download complete 
b7cf8f0d9e82: Download complete 
d0b170ebeeab: Download complete 
171efc310edf: Download complete 
522ed614b07a: Download complete 
605ed9113b86: Download complete 
22b93b23ebb9: Download complete 
2ac557a88fda: Download complete 
1f3b4c532640: Download complete 
27ebaac643a7: Download complete 
ce630195cb45: Download complete 
Status: Downloaded newer image for docker.io/registry:latest
[root@localhost ~]# docker images 
REPOSITORY                TAG                 IMAGE ID            CREATED             SIZE
docker.io/registry        latest              bca04f698ba8        6 months ago        422.8 MB

3、基于私有仓库镜像运行容器
[root@localhost ~]# docker run -d -p 5000:5000 -v /opt/data/registry:/tmp/registry docker.io/registry

[root@localhost ~]# docker ps -a
CONTAINER ID        IMAGE                COMMAND             CREATED             STATUS              PORTS                    NAMES
07c0c34925fd        docker.io/registry   "docker-registry"   About an hour ago   Up About an hour    0.0.0.0:5000->5000/tcp   elated_curie

4、访问私有仓库
[root@localhost ~]# curl 127.0.0.1:5000/v1/search
{"num_results": 0, "query": "", "results": []} //私有仓库为空，没有提交新镜像到仓库中

5、导入一个本地mysql镜像为镜像打个标签

[root@localhost ~]# docker load <yangqd_mysql5.5.tar
[root@localhost ~]# docker tag cdba6dd07aa2 127.0.0.1:5000/mysql5.5
[root@localhost ~]# docker images 
REPOSITORY                TAG                 IMAGE ID            CREATED             SIZE
127.0.0.1:5000/mysql5.5   latest              cdba6dd07aa2        3 days ago          304 MB
yangqd/mysql5.5           latest              cdba6dd07aa2        3 days ago          304 MB
docker.io/registry        latest              bca04f698ba8        6 months ago        422.8 MB

6、修改Docker配置文件制定私有仓库url
[root@localhost ~]# vim /etc/sysconfig/docker
添加此行
OPTIONS='--selinux-enabled --insecure-registry 10.0.2.8:5000'
[root@localhost ~]# service docker restart
Redirecting to /bin/systemctl restart  docker.service

7、提交镜像到本地私有仓库中
[root@localhost ~]# docker push 127.0.0.1:5000/mysql5.5
The push refers to a repository [127.0.0.1:5000/mysql5.5] (len: 1)
Sending image list
Pushing repository 127.0.0.1:5000/mysql5.5 (1 tags)
511136ea3c5a: Image successfully pushed 
00a0c78eeb6d: Image successfully pushed 
834629358fe2: Image successfully pushed 
571e8a51403c: Image successfully pushed 
87d5d42e693c: Image successfully pushed 
92b5ef05fe68: Image successfully pushed 
92d3910dc33c: Image successfully pushed 
cf2e9fa11368: Image successfully pushed 
2aeb2b6d9705: Image successfully pushed 
Pushing tag for rev [2aeb2b6d9705] on {http://127.0.0.1:5000/v1/repositories/mysql5.5/tags/latest}

8、查看私有仓库是否存在对应的镜像
[root@localhost ~]# curl 127.0.0.1:5000/v1/search
{"num_results": 1, "query": "", "results": [{"description": "", "name": "library/mysql5.5"}]}


三、从私有仓库中下载已有的镜像

1、登陆另外一台Docker客户端
安装docker过程同上，过程略……

2、修改Docker配置文件
[root@localhost ~]# vim /etc/sysconfig/docker
修改此行
OPTIONS='--selinux-enabled --insecure-registry 10.0.2.8:5000'        //添加私有仓库地址
[root@localhost ~]# service docker restart
Redirecting to /bin/systemctl restart  docker.service

3、从私有仓库中下载已有的镜像
[root@localhost ~]# docker pull 10.0.2.8:5000/mysql5.5
Using default tag: latest
Trying to pull repository 10.0.2.8:5000/mysql5.5 ... 
Pulling repository 10.0.2.8:5000/mysql5.5
cdba6dd07aa2: Pull complete 
42755cf4ee95: Pull complete 
6e1e82e2ceff: Pull complete 
ea04363d1a22: Pull complete 
24f9bc7215bd: Pull complete 
2dd0c49555bb: Pull complete 
639f172b6283: Pull complete 
9b62e8258306: Pull complete 
b369be27ec6c: Pull complete 
fec369fff1f3: Pull complete 
7aee551f831e: Pull complete 
ce6b31a0e623: Pull complete 
653b38759fc5: Pull complete 
Status: Downloaded newer image for 10.0.2.8:5000/mysql5.5:latest
10.0.2.8:5000/mysql5.5: this image was pulled from a legacy registry.  Important: This registry version will not be supported in future versions of docker.
[root@localhost ~]# docker images 
REPOSITORY               TAG                 IMAGE ID            CREATED             SIZE
10.0.2.8:5000/mysql5.5   latest              a60a2388c39f        3 days ago          304 MB





===============================bash脚本===============================================
#!/bin/bash
#
systemctl stop firewalld.service
systemctl disable firewalld.service 	
sed -i 's,SELINUX=enforcing,SELINUX=disabled,g' /etc/sysconfig/selinux 
setenforce 0	
yum install net-tools -y
yum install docker -y
yum upgrade device-mapper-libs
service docker start
chkconfig docker on




vim /etc/sysconfig/docker

OPTIONS='--selinux-enabled --insecure-registry 10.0.2.8:5000'

service docker restart
docker pull 10.0.2.8:5000/mysql5.5