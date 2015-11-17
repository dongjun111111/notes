<?php
namespace Admin\Controller;
class CategoryController extends AdminController {

    public function index(){
 
		$Model=D('addcategory');   //对应的是表名
		$result=$Model->field('*')->select();
		$this->assign('data',$result);
        $this->display('');
		
	}

    /* 编辑分类 */
    public function edit()
   {  $Model=D('addcategory'); 
        if(!IS_POST) //展示
	   {
	 //显示start	
	    $id=$_GET['id'];
		$result=$Model->where("id=$id")->select();
		$this->assign('data',$result);
		$this->display();
      //显示end
	   }
		else //新增
	    {
	  $uid=$_POST['uid'];
      $data['categoryname']=$_POST["categoryname"];
	
	  $res=$Model->where("id=$uid")->save($data); //TMD把这里的'id=$uid'改成"id=$uid"就成功了，也是醉了 ！！！！！！！！！！！
		if($res)
	    {
		$this->success('修改成功！',U('Index'));
		}
		else
	   {
		$this->error(empty($error) ? '未知错误！' : $error);
		}
	   }
	}	


	 public function delete(){
         //只在本页面的不用this->display();
		$id=$_GET['id'];
		if(IS_GET)  //get与post要分清楚
	   {
		$Mo=M('addcategory');
		$update=$Mo->where("id=$id")->delete();
		if($update)
	    {
		$this->success('删除成功！',U('Index'));
		}
		else
	   {
		$this->error(empty($error) ? '未知错误！' : $error);
		}
	   }
		
	}

	public function add(){
		if(!IS_POST)
		{  
		$this->display();
		}
		else
		{
		$Model=M('addcategory');
		 $arr['categoryname']=$_POST['categoryname'];
	  $arr['highlevel']=$_POST['highlevel'];
	  $arr['categorymessage']=$_POST['categorymessage'];
	  $arr['categorysort']=$_POST['categorysort'];
	  $arr['categorystatus']=$_POST['categorystatus']; 
	  $resu=$Model->add($arr);
	if($resu)
		{
	$this->success('添加成功',U('Index')); 
    	 }
		 else
		{
	$this->error('添加失败',U('Index')); 
		 }
		}
    
	
	
	}
}
