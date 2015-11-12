<?php
error_reporting(0);
header("content-type:text/html; charset=utf-8");
$conn=mysql_connect("localhost","root","");
mysql_select_db("root");
mysql_query("set names utf8");
$co=mysql_query("select * from addcategory");

?>
<html>
<body>
<a href="addCategory.php">添加</a>
<table>
<thead style="color:red;">
<tr><td>分类名</td><td>上级分类</td><td>注释</td><td>排序</td><td>状态</td><td>修改</td></tr>
</thead>
<tbody>
<?php
while($rows=mysql_fetch_array($co))
{//状态开始
$status=mysql_query("select * from category_status");
	while($rowstatus=mysql_fetch_array($status))
	{
	if($rows['categorystatus']==$rowstatus['status_id'])
	  {
	$rows['categorystatus']=$rowstatus['status_message'];
	  }
	}
//状态结束

//上级分类开始
$yes=mysql_query("select * from addcategory where id=".$rows['highlevel']." ");
while($rowsyes=mysql_fetch_array($yes))
{	
	$rows['highlevel']=$rowsyes['categoryname']; 
}
	if($rows['highlevel']=='0')
	{
	$rows['highlevel']="顶级分类";
	}
	
//上级分类结束

	echo "
<tr><td>".$rows['categoryname']."</td><td>".$rows['highlevel']."</td><td>".$rows['categorymessage']."</td><td>".$rows['categorysort']."</td><td>".$rows['categorystatus']."</td><td><a href=''>修改</a></td></tr> ";
	}
?>
</tbody>
</table>


</body>
</html>