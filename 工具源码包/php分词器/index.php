<?php
$sub  = @$_POST['submit'];
$text = @$_POST['text'];
if(!empty($sub)&&!empty($text)){
require_once "lib_splitword_full.php";
echo $text;
echo  '<br /><hr />';
$sp = new SplitWord();
echo $sp->SplitRMM($text);
$sp->Clear();
}
?>
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
<title>智能分词功能-张存超php个人博客</title>
</head>

<body>
<br />
<form action="" method="POST">
请输入要分词的字符：<input type="text" name="text" value="" /><br />
<input type="submit" name="submit" value="确定" />
</form>
<br /><br /><br /><br /><br /><br /><br /><br /><br /><br /><br /><br /><br /><br /><br />
<div align="center">技术支持：<a href="http://www.521php.com">php个人博客</a><span style="display:none;"><script src="http://s96.cnzz.com/stat.php?id=4200165&web_id=4200165" language="JavaScript"></script></span></div>
</body>
</html>