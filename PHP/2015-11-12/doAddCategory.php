<?php
error_reporting(0);
header("content-type:text/html; charset=utf-8");
$conn=mysql_connect("localhost","root","");
mysql_select_db("root");
mysql_query("set names utf8");
$categoryname=$_POST['categoryname'];
$highlevel=$_POST['id'];

if($highlevel==null)
{
$highlevel=$categoryname;
}

$categorymessage=$_POST['categorymessage'];
$categorysort=$_POST['categorysort'];
$categorystatus=$_POST['categorystatus'];
if($highlevel==$categoryname) //没有选择顶级分类，则是顶级分类
{

	$add=mysql_query("insert into addcategory(categoryname,highlevel,categorymessage,categorysort,categorystatus) values ('$categoryname','顶级分类','$categorymessage','$categorysort','$categorystatus')");
	}
else //有选择顶级分类
{
  $add=mysql_query("insert into addcategory(categoryname,highlevel,categorymessage,categorysort,categorystatus) values ('$categoryname','$highlevel','$categorymessage','$categorysort','$categorystatus')");
}

if($add)
{
 echo "<script>alert('添加成功');location.href='addCategory.php';</script>";
}
else
{
 echo "<script>alert('添加失败');location.href='addCategory.php';</script>";

}
?>