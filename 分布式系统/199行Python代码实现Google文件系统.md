<h1>199行Python代码实现Google文件系统</h1>
GFS,Google文件系统,位于整个Google基础设施的支柱。 然而,对很多人来说,这是一个谜,特别是对于那些有幸比低级C更熟悉高级python代码的操作系统资源。 但是不要害怕,我们要突破GFS的面纱和描述一个实现199行python。 自然,您可能想要读到GFS在最初的理论和设计谷歌研究GFS纸。 但我们的目标是给的核心概念,本文用真正的python代码工作。 完成运行的python 2.6源代码,如果你愿意,可以在这篇文章的结束。

GFS的简要总结如下。 GFS包含三个组件:一个客户,一个主人,和一个或多个chunkservers。 客户端是唯一的用户可见,programmer-accessible,系统的一部分。 它的功能类似于一个标准POSIX库文件。 主是单个服务器保存所有文件系统的元数据。 通过元数据,我们指的是每个文件的信息,其组成成分称为块,这些块的位置各种chunkservers。 chunkservers实际数据的存储位置,和绝大多数的网络流量客户端和chunkservers之间发生,为了避免主作为一个瓶颈。 下面我们将给更详细的描述通过GFS客户机,主人,和chunkserver在python中实现类,并关闭一个测试脚本和它的输出。

客户端类是唯一GFS图书馆用户可见的部分。 它搭建起了所有请求客户端文件系统访问和主人之间chunkservers数据存储和检索。 重要的是要注意,GFS似乎非常熟悉程序员正常的文件系统,不需要分布式知识。 所有的这些都是在客户端实现抽象出来。 当然也有一些例外,比如局部块知识分配处理文件的最有效的利用,比如在地图上减少算法,但我们避免了这种复杂性在这个实现。 最关键的是要注意正常的读,写,追加,存在,和删除调用可用的常见形式,以及如何这些都是由客户端实现类;我们也简化打开,关闭和下创建,将他们之前的方法。 每个方法的要点是一样的:主的提出元数据包括chunkservers块id和块的位置,然后更新任何必要的元数据与主,最后实际交易数据流只有chunkservers。
<pre>
class GFSClient:
    def __init__(self, master):
        self.master = master

    def write(self, filename, data): # filename is full namespace path
        if self.exists(filename): # if already exists, overwrite
            self.delete(filename)
        num_chunks = self.num_chunks(len(data))
        chunkuuids = self.master.alloc(filename, num_chunks)
        self.write_chunks(chunkuuids, data)

    def write_chunks(self, chunkuuids, data):
        chunks = [ data[x:x+self.master.chunksize] \
            for x in range(0, len(data), self.master.chunksize) ]
        chunkservers = self.master.get_chunkservers()
        for i in range(0, len(chunkuuids)): # write to each chunkserver
            chunkuuid = chunkuuids[i]
            chunkloc = self.master.get_chunkloc(chunkuuid)
            chunkservers[chunkloc].write(chunkuuid, chunks[i])

    def num_chunks(self, size):
        return (size // self.master.chunksize) \
            + (1 if size % self.master.chunksize > 0 else 0)

    def write_append(self, filename, data):
        if not self.exists(filename):
            raise Exception("append error, file does not exist: " \
                 + filename)
        num_append_chunks = self.num_chunks(len(data))
        append_chunkuuids = self.master.alloc_append(filename, \
            num_append_chunks)
        self.write_chunks(append_chunkuuids, data)

    def exists(self, filename):
        return self.master.exists(filename)

    def read(self, filename): # get metadata, then read chunks direct
        if not self.exists(filename):
            raise Exception("read error, file does not exist: " \
                + filename)
        chunks = []
        chunkuuids = self.master.get_chunkuuids(filename)
        chunkservers = self.master.get_chunkservers()
        for chunkuuid in chunkuuids:
            chunkloc = self.master.get_chunkloc(chunkuuid)
            chunk = chunkservers[chunkloc].read(chunkuuid)
            chunks.append(chunk)
        data = reduce(lambda x, y: x + y, chunks) # reassemble in order
        return data

    def delete(self, filename):
        self.master.delete(filename)
</pre>
主类模拟GFS主服务器。 这是所有存储的元数据,整个系统的核心节点。 客户机请求发起的主人,然后检索元数据之后,他们个人chunkservers直接对话。 这避免了主是一个瓶颈,因为元数据通常是短而低延迟。 元数据被实现为一个系列的词典,尽管在实际系统你的字典文件系统支持。 chunkservers变得可用的和不可用的通知通过心跳chunkserver身份验证和本地化信息高效存储都是简化这里分配chunkservers主本身。 但是我们仍然保留直接客户读/写chunkservers,绕过主,分布式系统是如何工作的。
<pre>
class GFSMaster:
    def __init__(self):
        self.num_chunkservers = 5
        self.max_chunkservers = 10
        self.max_chunksperfile = 100
        self.chunksize = 10
        self.chunkrobin = 0
        self.filetable = {} # file to chunk mapping
        self.chunktable = {} # chunkuuid to chunkloc mapping
        self.chunkservers = {} # loc id to chunkserver mapping
        self.init_chunkservers()

    def init_chunkservers(self):
        for i in range(0, self.num_chunkservers):
            chunkserver = GFSChunkserver(i)
            self.chunkservers[i] = chunkserver

    def get_chunkservers(self):
        return self.chunkservers

    def alloc(self, filename, num_chunks): # return ordered chunkuuid list
        chunkuuids = self.alloc_chunks(num_chunks)
        self.filetable[filename] = chunkuuids
        return chunkuuids

    def alloc_chunks(self, num_chunks):
        chunkuuids = []
        for i in range(0, num_chunks):
            chunkuuid = uuid.uuid1()
            chunkloc = self.chunkrobin
            self.chunktable[chunkuuid] = chunkloc
            chunkuuids.append(chunkuuid)
            self.chunkrobin = (self.chunkrobin + 1) % self.num_chunkservers
        return chunkuuids

    def alloc_append(self, filename, num_append_chunks): # append chunks
        chunkuuids = self.filetable[filename]
        append_chunkuuids = self.alloc_chunks(num_append_chunks)
        chunkuuids.extend(append_chunkuuids)
        return append_chunkuuids

    def get_chunkloc(self, chunkuuid):
        return self.chunktable[chunkuuid]

    def get_chunkuuids(self, filename):
        return self.filetable[filename]

    def exists(self, filename):
        return True if filename in self.filetable else False

    def delete(self, filename): # rename for later garbage collection
        chunkuuids = self.filetable[filename]
        del self.filetable[filename]
        timestamp = repr(time.time())
        deleted_filename = "/hidden/deleted/" + timestamp + filename
        self.filetable[deleted_filename] = chunkuuids
        print "deleted file: " + filename + " renamed to " + \
             deleted_filename + " ready for gc"

    def dump_metadata(self):
        print "Filetable:",
        for filename, chunkuuids in self.filetable.items():
            print filename, "with", len(chunkuuids),"chunks"
        print "Chunkservers: ", len(self.chunkservers)
        print "Chunkserver Data:"
        for chunkuuid, chunkloc in sorted(self.chunktable.iteritems(), key=operator.itemgetter(1)):
            chunk = self.chunkservers[chunkloc].read(chunkuuid)
            print chunkloc, chunkuuid, chunk
</pre>
chunkserver类是最小的在这个项目。 这是一个实际的不同的盒子在大规模数据中心运行,连接到一个网络通过主和客户取得联系。 在GFS chunkservers相对“哑巴”,他们知道只有块,也就是说,文件数据分解成碎片。 他们看不到整个文件的全貌,在整个文件系统,相关的元数据,等。我们实现这个类作为一个简单的本地存储,运行测试代码后,您可以检查通过查看目录路径“/ tmp / gfs /块”。 您想要在一个真正的系统持久性存储块信息的备份。
<pre>
class GFSChunkserver:
    def __init__(self, chunkloc):
        self.chunkloc = chunkloc
        self.chunktable = {}
        self.local_filesystem_root = "/tmp/gfs/chunks/" + repr(chunkloc)
        if not os.access(self.local_filesystem_root, os.W_OK):
            os.makedirs(self.local_filesystem_root)

    def write(self, chunkuuid, chunk):
        local_filename = self.chunk_filename(chunkuuid)
        with open(local_filename, "w") as f:
            f.write(chunk)
        self.chunktable[chunkuuid] = local_filename

    def read(self, chunkuuid):
        data = None
        local_filename = self.chunk_filename(chunkuuid)
        with open(local_filename, "r") as f:
            data = f.read()
        return data

    def chunk_filename(self, chunkuuid):
        local_filename = self.local_filesystem_root + "/" \
            + str(chunkuuid) + '.gfs'
        return local_filename
</pre>
我们使用main()作为所有客户的测试方法,包括异常。 我们首先创建一个主和客户对象,然后写一个文件。 这写是由客户机相同的方式执行真正的GFS:第一块的元数据从主,然后直接向每个chunkserver写块。 附加功能类似。 删除处理的GFS时尚,重命名文件隐藏名称空间和树叶后垃圾收集。 一个转储显示元数据内容。 请注意,这是一个单线程的测试,这个演示程序不支持并发性,尽管这可以添加元数据以适当的锁。
<pre>
def main():
    # test script for filesystem

    # setup
    master = GFSMaster()
    client = GFSClient(master)

    # test write, exist, read
    print "\nWriting..."
    client.write("/usr/python/readme.txt", """
        This file tells you all about python that you ever wanted to know.
        Not every README is as informative as this one, but we aim to please.
        Never yet has there been so much information in so little space.
        """)
    print "File exists? ", client.exists("/usr/python/readme.txt")
    print client.read("/usr/python/readme.txt")

    # test append, read after append
    print "\nAppending..."
    client.write_append("/usr/python/readme.txt", \
        "I'm a little sentence that just snuck in at the end.\n")
    print client.read("/usr/python/readme.txt")

    # test delete
    print "\nDeleting..."
    client.delete("/usr/python/readme.txt")
    print "File exists? ", client.exists("/usr/python/readme.txt")

    # test exceptions
    print "\nTesting Exceptions..."
    try:
        client.read("/usr/python/readme.txt")
    except Exception as e:
        print "This exception should be thrown:", e
    try:
        client.write_append("/usr/python/readme.txt", "foo")
    except Exception as e:
        print "This exception should be thrown:", e

    # show structure of the filesystem
    print "\nMetadata Dump..."
    print master.dump_metadata()

</pre>
并把它放在一起,这是测试脚本的输出从python解释器运行。 特别注意到主元数据转储到最后,在那里你可以看到数据块分布在chunkservers打乱顺序,只能重新由客户机指定的顺序由主元数据。

测试：
<pre>
$ python gfs.py

Writing...
File exists?  True

        This file tells you all about python that you ever wanted to know.
        Not every README is as informative as this one, but we aim to please.
        Never yet has there been so much information in so little space.


Appending...

        This file tells you all about python that you ever wanted to know.
        Not every README is as informative as this one, but we aim to please.
        Never yet has there been so much information in so little space.
        I'm a little sentence that just snuck in at the end.


Deleting...
deleted file: /usr/python/readme.txt renamed to /hidden/deleted/1289928955.7363091/usr/python/readme.txt ready for gc
File exists?  False

Testing Exceptions...
This exception should be thrown: read error, file does not exist: /usr/python/readme.txt
This exception should be thrown: append error, file does not exist: /usr/python/readme.txt

Metadata Dump...
Filetable: /hidden/deleted/1289928955.7363091/usr/python/readme.txt with 30 chunks
Chunkservers:  5
Chunkserver Data:
0 f76734ce-f1a7-11df-b529-001d09d5b664 mation in
0 f7671750-f1a7-11df-b529-001d09d5b664  you ever
0 f7670bd4-f1a7-11df-b529-001d09d5b664
        T
0 f767b656-f1a7-11df-b529-001d09d5b664 le sentenc
0 f7672182-f1a7-11df-b529-001d09d5b664  is as inf
0 f7672b0a-f1a7-11df-b529-001d09d5b664 se.

1 f767b85e-f1a7-11df-b529-001d09d5b664 e that jus
1 f76736b8-f1a7-11df-b529-001d09d5b664 so little
1 f767193a-f1a7-11df-b529-001d09d5b664 wanted to
1 f7670f3a-f1a7-11df-b529-001d09d5b664 his file t
1 f767236c-f1a7-11df-b529-001d09d5b664 ormative a
1 f7672cf4-f1a7-11df-b529-001d09d5b664   Never ye
2 f7671b2e-f1a7-11df-b529-001d09d5b664 know.

2 f7671142-f1a7-11df-b529-001d09d5b664 ells you a
2 f767ba48-f1a7-11df-b529-001d09d5b664 t snuck in
2 f7672556-f1a7-11df-b529-001d09d5b664 s this one
2 f76738e8-f1a7-11df-b529-001d09d5b664 space.

2 f7672ee8-f1a7-11df-b529-001d09d5b664 t has ther
3 f767bcb4-f1a7-11df-b529-001d09d5b664  at the en
3 f7671d40-f1a7-11df-b529-001d09d5b664     Not ev
3 f76730d2-f1a7-11df-b529-001d09d5b664 e been so
3 f7673adc-f1a7-11df-b529-001d09d5b664
3 f767135e-f1a7-11df-b529-001d09d5b664 ll about p
3 f7672740-f1a7-11df-b529-001d09d5b664 , but we a
4 f7672920-f1a7-11df-b529-001d09d5b664 im to plea
4 f767b3f4-f1a7-11df-b529-001d09d5b664 I'm a litt
4 f767bea8-f1a7-11df-b529-001d09d5b664 d.

4 f7671552-f1a7-11df-b529-001d09d5b664 ython that
4 f76732e4-f1a7-11df-b529-001d09d5b664 much infor
4 f7671f8e-f1a7-11df-b529-001d09d5b664 ery README
</pre>
当然我们现在缺乏一些必要的复杂性GFS对一个完整的功能系统:元数据锁定,租赁,复制,掌握故障转移,本地化的数据,垃圾收集chunkserver心跳、删除文件。 但是我们这里演示GFS的要点,并将帮助你更好的理解基础。 也可以为自己的探索到一个起点更详细的分布式文件系统在python代码。
####原文链接
http://clouddbs.blogspot.jp/2010/11/gfs-google-file-system-in-199-lines-of.html