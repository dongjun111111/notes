#mysql
##mysql中的isnull()、ifnull()、is null与 isnotnull
###isnull()
找到username是null的行
<pre>
select * from 表 where isnull(username)
</pre>
###is null
找到username是null的行
<pre>
SELECT * FROM 表 WHERE username is null
</pre>
###is not null
找到username不是null的行
<pre>
SELECT * FROM 表 WHERE username is not null
</pre>
###ifnull()
列出所有行，并且将username是null的字段替换成jason
<pre>
select ifnull(username,'jason'),* from 表
</pre>
##mysql数据库设置为只读
1. 设置用户权限；
2. set GLOBAL read_only = true.


##Linux mysql使用常用命令
1.linux下启动mysql的命令：

mysqladmin start
/ect/init.d/mysql start (前面为mysql的安装路径)

2.linux下重启mysql的命令：

mysqladmin restart
/ect/init.d/mysql restart (前面为mysql的安装路径)

3.linux下关闭mysql的命令：

mysqladmin shutdown
/ect/init.d/mysql shutdown (前面为mysql的安装路径)

4.连接本机上的mysql：

进入目录mysql\bin，再键入命令mysql -uroot -p， 回车后提示输入密码。
退出mysql命令：exit（回车）

5.修改mysql密码：

mysqladmin -u用户名 -p旧密码 password 新密码
或进入mysql命令行SET PASSWORD FOR root=PASSWORD("root");

6.增加新用户。（注意：mysql环境中的命令后面都带一个分号作为命令结束符）

grant select on 数据库.* to 用户名@登录主机 identified by "密码"
如增加一个用户test密码为123，让他可以在任何主机上登录， 并对所有数据库有查询、插入、修改、删除的权限。首先用以root用户连入mysql，然后键入以下命令：
grant select,insert,update,delete on *.* to " Identified by "123";

二、有关mysql数据库方面的操作
必须首先登录到mysql中，有关操作都是在mysql的提示符下进行，而且每个命令以分号结束

1、显示数据库列表。

show databases;

2、显示库中的数据表：

use mysql； ／／打开库
show tables;

3、显示数据表的结构：

describe 表名;

4、建库：

create database 库名;

GBK: create database test2 DEFAULT CHARACTER SET gbk COLLATE gbk_chinese_ci;
UTF8: CREATE DATABASE `test2` DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;

5、建表：

use 库名；
create table 表名(字段设定列表)；

6、删库和删表:

drop database 库名;
drop table 表名；

7、将表中记录清空：

delete from 表名;

truncate table  表名;

8、显示表中的记录：

select * from 表名;

9、编码的修改

如果要改变整个mysql的编码格式： 
启动mysql的时候，mysqld_safe命令行加入 
--default-character-set=gbk 

如果要改变某个库的编码格式：在mysql提示符后输入命令 
alter database db_name default character set gbk;

10.重命名表

alter table t1 rename t2;

11.查看sql语句的效率

 explain < table_name >

例如：explain select * from t3 where id=3952602;

12.用文本方式将数据装入数据库表中(例如D:/mysql.txt)

mysql> LOAD DATA LOCAL INFILE "D:/mysql.txt" INTO TABLE MYTABLE;

三、数据的导入导出

1、文本数据转到数据库中

文本数据应符合的格式：字段数据之间用tab键隔开，null值用来代替。例：
1 name duty 2006-11-23
数据传入命令 load data local infile "文件名" into table 表名;

2、导出数据库和表

mysqldump --opt news > news.sql（将数据库news中的所有表备份到news.sql文件，news.sql是一个文本文件，文件名任取。）
mysqldump --opt news author article > author.article.sql（将数据库news中的author表和article表备份到author.article.sql文件， author.article.sql是一个文本文件，文件名任取。）
mysqldump --databases db1 db2 > news.sql（将数据库dbl和db2备份到news.sql文件，news.sql是一个文本文件，文件名任取。）
mysqldump -h host -u user -p pass --databases dbname > file.dump
就是把host上的以名字user，口令pass的数据库dbname导入到文件file.dump中
mysqldump --all-databases > all-databases.sql（将所有数据库备份到all-databases.sql文件，all-databases.sql是一个文本文件，文件名任取。）

3、导入数据
mysql < all-databases.sql（导入数据库）
mysql>source news.sql;（在mysql命令下执行，可导入表）