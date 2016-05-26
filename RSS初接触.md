###通过实例总结RSS基本语法知识
RSS 文档使用一种简单的自我描述的语法。下面是一个简单的RSS文档：
<pre>
<?xml version="1.0" encoding="ISO-8859-1"?>
<rss version="2.0">
<channel>
  <title>W3School Home Page</title>
  <link>http://www.w3school.com.cn</link>
  <description>Free web building tutorials</description>
  <item>
    <title>RSS Tutorial</title>
    <link>http://www.w3school.com.cn/rss</link>
    <description>New RSS tutorial on W3School</description>
  </item>
  <item>
    <title>XML Tutorial</title>
    <link>http://www.w3school.com.cn/xml</link>
    <description>New XML tutorial on W3School</description>
  </item>
</channel>
</rss>
</pre>
文档中的第一行：XML声明 - 定义了文档中使用的 XML 版本和字符编码。此例子遵守 1.0 规范，并使用 ISO-8859-1 (Latin-1/West European) 字符集。

下一行是标识此文档是一个 RSS 文档的 RSS 声明（此例是 RSS version 2.0）。

下一行含有 <channel>元素。此元素用于描述RSS feed。

<channel>元素有三个必需的子元素：

-  title- 定义频道的标题。（比如 w3school 首页）
-  link - 定义到达频道的超链接。（比如 www.w3school.com.cn）
-  description - 描述此频道（比如免费的网站建设教程）

每个 <channel> 元素可拥有一个或多个 <item> 元素。

每个<item>元素可定义RSS feed 中的一篇文章或 "story"。

item 元素拥有三个必需的子元素：

- title - 定义项目的标题。（比如 RSS 教程）
- link - 定义到达项目的超链接。（比如 http://www.w3school.com.cn/rss）
- description - 描述此项目（比如 w3school 的 RSS 教程）

最后，后面的两行关闭 <channel> 和 <rss> 元素。

在RSS中书写注释的语法与HTML的语法类似：
<pre>
<!-- This is an RSS comment -->
</pre>
几点注意事项,因为RSS也是XML，所以：

1. 所有的元素必许拥有关闭标签
2. 元素对大小写敏感
3. 元素必需被正确地嵌套
4. 属性值必须带引号