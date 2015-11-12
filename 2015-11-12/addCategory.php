<?php
error_reporting(0);
header("content-type:text/html; charset=utf-8");
$conn=mysql_connect("localhost","root","");
mysql_select_db("root");
mysql_query("set names utf8");

$con=mysql_query("select * from addcategory");
?>
<html>
<body>
<a href="Category.php">返回</a>

<form action="doAddCategory.php" method="post">
类名:<input type="text" name="categoryname" /><br>
上级分类:<select name="id">
<option value="" selected>-请选择-</option>
<?php
while($rows=mysql_fetch_array($con))
{
echo "<option value=".$rows['id']." >".$rows['categoryname']."</option> ";
}
?>
</select>
<br>
注释:<textarea cols="40" rows="4" name="categorymessage"></textarea><br>
排序:<input type="text" name="categorysort" value="999" /><br>
状态:
<label><input type="radio" name="categorystatus" value="0" checked /></label>正常
<label>
<input type="radio" name="categorystatus" value="1" />禁用</label><br>
<input type="submit" value="提交"/>
</form>
</body>
</html>