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