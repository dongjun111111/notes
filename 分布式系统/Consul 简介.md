##Consul 简介
在 GitHub 上，Consul 被称为“一种用于服务发现、监控和配置的工具”。Consul 是由 Vagrant 的开发公司 HashiCorp 开发的开源项目之一。 它可以提供一个具有高可用性的分布式系统，用以注册服务、存储共享配置并保持多个数据中心的准确视图。此外，它作为一个简单的 Go 程序，因此不需要部署。
###声明式服务
下面让我们一起来了解下注册、外部服务及 Docker 在我们的解决方案中所扮演的角色。为了便于说明，我们首先想象一个在 MongoDB 中存储数据并通过 Mailgun 发送电子邮件的现代应用。后者是一种外部服务，我们会自行运行前者
###注册
为了找到这些重要属性，首先需要注册服务。接下来，将会在集群的每个节点上运行一个 Consul 代理，负责连接 Consul 服务器，以便找到节点的服务并执行健康检查
<pre>
// 注册服务
# download and install the latest version
wget https://dl.bintray.com/mitchellh/consul/0.5.2_linux_amd64.zip -O 
/tmp/consul.zip
cd /usr/local/bin && unzip /tmp/consul.zip

# create state and configuration directories
mkdir -p {/srv/consul,/etc/consul.d}

# check that everything worked
consul --help
</pre>
MongoDB 的下载量超过 1,000 万次，是文档数据库的理想之选。我们使用该服务并将以下文件保存在 /etc/consul.d/mongo.json.
<pre>
{
    "service":{
        "name":"mongo",
        "tags":[
            "database",
            "nosql"
        ],
         "port":27017,
         "check":{
             "name":"status",
             "script":"mongo --eval 'printjson(rs.status())'",
             "interval":"30s"
         }
     }
}
</pre>
上述语法提供了简明且可读性强的声明式方法，可供您定义服务属性及健康检查。您可以在版本控制系统中提取这些文件，并立即识别应用的组件。上述文件声明了端口 27017 上一个名为“mongo”的服务。 检查部分为 Consul 代理提供了一个脚本，可用于测试节点是否处于健康状态。在向服务器请求服务时，您需要确保服务器能返回可靠的终端设备。

启动实际的 Mongo 服务器及本地 Consul 代理
<pre>
# launch mongodb server on default port 27017
mongod

# launch local agent
consul agent \
    -join $CONSUL_HOST \  # explicitly provide how to reach the server
    -data-dir /data/consul \  # internal state storage
    -config-dir /etc/consul.d  # configuration directory where services and checks 
                               # are expected to be defined
</pre>
是否有作用？让我们来查询 Consul HTTP API
<pre>
# fetch infrastructure overview
curl $CONSUL_IP:8500/v1/catalog/nodes
[{"Node":"consul-server-1","Address":"172.17.0.1"},{"Node":"mongo-1","Address"
:"172.17.0.2"}]

# consul correctly registered mongo service
curl $CONSUL_IP:8500/v1/catalog/service/mongo
[{
    "Node":"mongo-1",
    "Address":"172.17.0.2",
    "ServiceID":"mongo",
    "ServiceName":"mongo",
    "ServiceTags":["database", "no-sql"],
    "ServiceAddress":"",
    "ServicePort":27017
}]

# it also exposes health state
curl $CONSUL_IP:8500/v1/health/service/mongo
[{
    "Node":{
        "Node":"mongo-1",
    },
    "Service":{
        "ID":"mongo",
        "Service":"mongo",
        "Tags":["database","no-sql"],
        "Address":"",
    },
    "Checks":[{
        "Node":"mongo-1",
        "CheckID":"service:mongo",
        "Name":"Service 'mongo' check",
        "Status":"passing",
        "Notes":"",
        "Output":"MongoDB shell version:3.0.3\nconnecting to: test\n{ \"ok\" :0, 
    \"errmsg\" :\"not running with --replSet\", \"code\" :76 }\n",
        "ServiceID":"mongo",
        "ServiceName":"mongo"
    },{
        "Node":"mongo-1",
        "CheckID":"serfHealth",
        "Status":"passing",
        "Notes":"",
        "Output":"Agent alive and reachable",
        "ServiceID":"",
        "ServiceName":""
    }]
}]
</pre>
在给定 Consul 代理或服务器地址的情况下，能够处理 HTTP 请求的集群中的任何代码均可使用该信息。 下面我将会处理过程做简要说明，但首先让我们来了解一下如何注册超出控制访问的服务，以及如何借助 Docker 实现自动化。
###外部服务
为了避免做无用功，将第三方服务集成到应用中是比较明智的方法。但在这种情况下，您不能再适当的节点上启动 Consul 代理。Consul 会再次将您覆盖在内
<pre>
# manually register mailgun service through the HTTP API
curl -X PUT -d \
    '{"Datacenter":"dc1", "Node":"mailgun", "Address":"http://www.mailgun.com",
 "Service":{"Service":"email", "Port":80}, "Check":{"Name":"mailgun api", 
 "http":"www.status.mailgun.com", "interval":"360s", "timeout":"1s"}}' \
    http://$CONSUL_IP:8500/v1/catalog/register

# looks like we're all good !
curl $CONSUL_IP:8500/v1/catalog/services
{"consul":[],"email":[],"mongo":["database","nosql"]}
</pre>
由于 Mailgun 是一个 Web 服务，因此使用 HTTP 字段来检查 API 的可用性。若要深入了解 Consul 的强大功能，请参阅综合性说明文档。
###Docker 集成
到目前为止，Go 二进制库、单个 JSON 文件以及一些 HTTP 请求均支持服务发现工作流。您当然无需束缚于某种特定技术，但正如前面所说，这种灵活的设置特别适合于微服务。
在这种情况下，借助 Docker，可以将服务打包至可复写的自注册容器中。在现有的 mongo.json 中，仅使用 清单 8 中的 Dockerfile 和 Procfile。
<pre>
//将服务打包到可复写的自注册容器中
# Dockerfile
# start from official mongo image
FROM mongo:3.0

RUN apt-get update && apt-get install -y unzip

# install consul agent
ADD https://dl.bintray.com/mitchellh/consul/0.5.2_linux_amd64.zip /tmp/consul.zip
RUN cd /bin && \
    unzip /tmp/consul.zip&& \
    chmod +x /bin/consul && \
    mkdir -p {/data/consul,/etc/consul.d} && \
    rm /tmp/consul.zip

# copy service and check definition, as we wrote them earlier
ADD mongo.json /etc/consul.d/mongo.json

# Install goreman - foreman clone written in Go language
ADD https://github.com/mattn/goreman/releases/download/v0.0.6
/goreman_linux_amd64.tar.gz /tmp/goreman.tar.gz
RUN tar -xvzf /tmp/goreman.tar.gz -C /usr/local/bin --strip-components 1 && \
    rm -r  /tmp/goreman*

# copy startup script
ADD Procfile /root/Procfile

# launch both mongo server and consul agent
ENTRYPOINT ["goreman"]
CMD ["-f", "/root/Procfile", "start"]
</pre>
Dockerfile 用于定义在启动容器时运行的单个命令。 不过，我们需要同时运行 MongoDB 和 Consul. 我们可以通过 Goreman 实现这一点。它能够读取名为 Procfile 的配置文件，用以定义多个管理流程（生命周期、环境、日志等）。在容器领域，这种方法是一个悖论，而且其他解决方案也存在，但现在我们可以通过更简单的方式做到这一点。
<pre>
# Procfile
database: mongod
consul: consul agent -join $CONSUL_HOST -data-dir /data/consul -config-dir
/etc/consul.d
</pre>
<pre>
//构建容器的外壳命令
ls
Dockerfile  mongo.json  Procfile

docker build -t article/mongo .
# ...

docker run --detach --name docker-mongo \
    --hostname docker-mongo-2 \  # if not explicitly configured, consul agent 
set its name to the node hostname
    --env CONSUL_HOST=$CONSUL_IP article/mongo

curl $CONSUL_IP:8500/v1/catalog/nodes
[
    {
        "Node":"consul-server-1",
        "Address":"172.17.0.1"
    }, {
        "Node":"docker-mongo-2",
        "Address":"172.17.0.3"
    }, {
        "Node":"mailgun",
        "Address":"http://www.mailgun.com"
    }, {
        "Node":"mongo-1",
        "Address":"172.17.0.2"
    }
]
</pre>
太棒了！将 Docker 结合到服务发现流程中，效果非常好！<br>
我们可以按照 清单 6 中所述查询 $CONSUL_IP:8500/v1/catalog/service/mongo，找到服务端口，从而获得更多详情。Consul 可以提供容器 IP，以此作为服务地址。即便 Docker 将其映射到主机上一个随机值上，只要是容器提供端口，该方法都适用。不过，在多主机拓扑中，您需要明确地将容器的端口映射到主机的相同端口上。为了避免这一限制，我们可以考虑采用 Weave。

总的来说，在提供多个数据中心的服务信息时，大致步骤如下：
至少启动 

1.  个 Consul 服务器，并存储其地址。
2. 在每个节点上：

- 下载 Consul 二进制库。
- 写入服务并检查其在 Consul 配置目录中的定义。
- 启动应用。
- 使用另一代理或服务器的地址启动 Consul 代理。