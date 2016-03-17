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