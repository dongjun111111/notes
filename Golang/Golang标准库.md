##go标准库
bufio 实现缓冲的I/O<br>
bytes 提供了对字节切片操作的函数 <br>
crypto 收集了常见的加密常数 <br>
errors 实现了操作错误的函数 <br>
Expvar 为公共变量提供了一个标准的接口，如服务器中的运算计数器 <br>
flag 实现了命令行标记解析 <br>
fmt 实现了格式化输入输出 <br>
hash 提供了哈希函数接口 <br>
html 实现了一个HTML5兼容的分词器和解析器 <br>
image 实现了一个基本的二维图像库 <br>
io 提供了对I/O原语的基本接口 <br>
log 它是一个简单的记录包，提供最基本的日志功能 <br>
math 提供了一些基本的常量和数学函数 <br>
mine 实现了部分的MIME规范 <br>
net 提供了一个对UNIX网络套接字的可移植接口，包括TCP/IP、UDP域名解析和 <br>
UNIX域套接字 <br>
os 为操作系统功能实现了一个平台无关的接口 <br>
path 实现了对斜线分割的文件名路径的操作 <br>
reflect 实现了运行时反射，允许一个程序以任意类型操作对象 <br>
regexp 实现了一个简单的正则表达式库 <br>
runtime 包含与Go运行时系统交互的操作，如控制goroutine的函数 <br>
sort 提供对集合排序的基础函数集 <br>
strconv 实现了在基本数据类型和字符串之间的转换 <br>
strings 实现了操作字符串的简单函数 <br>
sync 提供了基本的同步机制，如互斥锁 <br>
syscall 包含一个低级的操作系统原语的接口 <br>
testing 提供对自动测试Go包的支持 <br>
time 提供测量和显示时间的功能 <br>
unicode Unicode编码相关的基础函数 <br>
archive tar 实现对tar压缩文档的访问 <br>
zip 提供对ZIP压缩文档的读和写支持 <br>
bzip2 实现了bzip2解压缩 <br>
flate 实现了RFC 1951中所定义的DEFLATE压缩数据格式 <br>
gzip 实现了RFC 1951中所定义的gzip格式压缩文件的读和写 <br>
lzw 实现了Lempel-Ziv-Welch编码格式的压缩的数据格式，参见T. A. Welch, A <br>
Technique for High-Performance Data Compression, Computer, 17(6) (June 1984), pp 
8-19 
compress 
zlib 实现了RFC 1950中所定义的zlib格式压缩数据的读和写 <br>
heap 提供了实现heap.Interface接口的任何类型的堆操作 <br>
list 实现了一个双链表 <br>
container <br>
ring 实现了对循环链表的操作 <br>
aes 实现了AES加密（以前的Rijndael），详见美国联邦信息处理标准（197号文） <br>
cipher 实现了标准的密码块模式，该模式可包装进低级的块加密实现中 <br>
des 实现了数据加密标准（Data Encryption Standard，DES）和三重数据加密算法（Triple 
Data Encryption Algorithm，TDEA），详见美国联邦信息处理标准（46-3号文） <br>
dsa 实现了FIPS 186-3所定义的数据签名算法（Digital Signature Algorithm） <br>
ecdsa 实现了FIPS 186-3所定义的椭圆曲线数据签名算法（Elliptic Curve Digital Signature Algorithm） <br>
elliptic 实现了素数域上几个标准的椭圆曲线 <br>
hmac 实现了键控哈希消息身份验证码（Keyed-Hash Message Authentication Code， 
HMAC），详见美国联邦信息处理标准（198号文） <br>
md5 实现了RFC 1321中所定义的MD5哈希算法 <br>
rand 实现了一个加密安全的伪随机数生成器 <br>
rc4 实现了RC4加密，其定义见Bruce Schneier的应用密码学（Applied Cryptography） <br>
rsa 实现了PKCS#1中所定义的RSA加密 <br>
sha1 实现了RFC 3174中所定义的SHA1哈希算法 <br>
sha256 实现了FIPS 180-2中所定义的SHA224和SHA256哈希算法 <br>
sha512 实现了FIPS 180-2中所定义的SHA384和SHA512哈希算法 <br>
subtle 实现了一些有用的加密函数，但需要仔细考虑以便正确应用它们 <br>
tls 部分实现了RFC 4346所定义的TLS 1.1协议 <br>
x509 可解析X.509编码的键值和证书 <br>
crypto <br>
x509/pkix 包含用于对X.509证书、CRL和OCSP的ASN.1解析和序列化的共享的、低级的结构 <br>
database sql 围绕SQL提供了一个通用的接口 <br>
sql/driver 定义了数据库驱动所需实现的接口，同sql包的使用方式 <br>
dwarf 提供了对从可执行文件加载的DWARF调试信息的访问，这个包对于实现Go语言
的调试器非常有价值 
elf 实现了对ELF对象文件的访问。ELF是一种常见的二进制可执行文件和共享库的  <br>
文件格式。Linux采用了ELF格式  <br>
gosym 访问Go语言二进制程序中的调试信息。对于可视化调试很有价值  <br>
macho 实现了对 http://developer.apple.com/mac/library/documentation/DeveloperTools/Conceptual/ 
MachORuntime/Reference/reference.html 所定义的Mach-O对象文件的访问  <br>
debug 
pe 实现了对PE（Microsoft Windows Portable Executable）文件的访问  <br>
ascii85 实现了ascii85数据编码，用于btoa工具和Adobe’s PostScript以及PDF文档格式  <br>
asn1 实现了解析DER编码的ASN.1数据结构，其定义见ITU-T Rec X.690  <br>
base32 实现了RFC 4648中所定义的base32编码  <br>
base64 实现了RFC 4648中所定义的base64编码  <br>
binary 实现了在无符号整数值和字节串之间的转化，以及对固定尺寸值的读和写  <br>
csv 可读和写由逗号分割的数值（csv）文件  <br>
gob 管理gob流——在编码器（发送者）和解码器（接收者）之间进行二进制值交换  <br>
hex 实现了十六进制的编码和解码  <br>
json 实现了定义于RFC 4627中的JSON对象的编码和解码  <br>
pem 实现了PEM（Privacy Enhanced Mail）数据编码  <br>
encoding  <br>
xml 实现了一个简单的可理解XML名字空间的XML 1.0解析器  <br>
ast 声明了用于展示Go包中的语法树类型  <br>
build 提供了构建Go包的工具  <br>
doc 从一个Go AST（抽象语法树）中提取源代码文档  <br>
parser 实现了一个Go源文件解析器  <br>
printer 实现了对AST（抽象语法树）的打印  <br>
scanner 实现了一个Go源代码文本的扫描器  <br>
go 
token 定义了代表Go编程语言中词法标记以及基本操作标记（printing、predicates）的常 
量  <br>
adler32 实现了Adler-32校验和  <br>
crc32 实现了32位的循环冗余校验或CRC-32校验和  <br>
crc64 实现了64位的循环冗余校验或CRC-64校验和  <br>
hash 
fnv 实现了Glenn Fowler、Landon Curt Noll和Phong Vo所创建的FNV-1和FNV-1a未加 
密哈希函数  <br>
html template 它自动构建HTML输出，并可防止代码注入  <br>
color 实现了一个基本的颜色库  <br>
draw 提供一些做图函数  <br>
gif 实现了一个GIF图像解码器  <br>
jpeg 实现了一个JPEG图像解码器和编码器  <br>
image 
png 实现了一个PNG图像解码器和编码器  <br>
index suffixarray 通过构建内存索引实现的高速字符串匹配查找算法  <br>
io ioutil 实现了一些实用的I/O函数  <br>
log syslog 提供了对系统日志服务的简单接口  <br>
big 实现了多精度的算术运算（大数）  <br>
cmplx 为复数提供了基本的常量和数学函数  <br>
Math rand 实现了伪随机数生成器  <br>
mime multipart 实现了在RFC 2046中定义的MIME多个部分的解析  <br>
http 提供了HTTP客户端和服务器的实现  <br>
mail 实现了对邮件消息的解析  <br>
rpc 提供了对一个来自网络或其他I/O连接的对象可导出的方法的访问  <br>
smtp 实现了定义于RFC 5321中的简单邮件传输协议（Simple Mail Transfer Protocol) 
textproto 实现了在HTTP、NNTP和SMTP中基于文本的通用的请求/响应协议  <br>
url 解析URL并实现查询转义  <br>
http/cgi 实现了定义于RFC 3875中的CGI（通用网关接口）  <br>
http/fcgi 实现了FastCGI协议  <br>
http/httptest 提供了一些HTTP测试应用  <br>
http/httputil 提供了一些HTTP应用函数，这些是对net/http包中的东西的补充，只不过相对 
不太常用  <br>
http/pprof 通过其HTTP服务器运行时提供性能测试数据，该数据的格式正是pprof可视化工 
具需要的  <br>
net 
rpc/jsonrpc 为rpc包实现了一个JSON-RPC ClientCodec和ServerCodec  <br>
os exec 可运行外部命令  <br>
user 通过名称和id进行用户账户检查  <br>
path filepath 实现了以与目标操作系统定义文件路径相兼容的方式处理文件名路径  <br>
regexp syntax 将正则表达式解析为语法树  <br>
runtime debug 包含当程序在运行时调试其自身的功能  <br>
pprof 以pprof可视化工具需要的格式写运行时性能测试数据  <br>
sync atomic 提供了低级的用于实现同步算法的原子级的内存机制  <br>
iotest 提供一系列测试目的的类型，实现了Reader和Writer标准接口  <br>
quick 实现了用于黑箱测试的实用函数  <br>
testing  <br>
script 帮助测试使用通道的 <br>