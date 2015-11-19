<?php
namespace Admin\Controller;
use \Think\Page;
header('content-type:text/html;charset=utf-8');
class CategoryController extends AdminController {

    public function index(){
		$Model=D('addcategory'); 
	/*	$result=$Model->field('*')->limit(5)->page($_GET['p'])->select();
        $count=$Model->count();
		$Everypage=5;
		$pagecount=ceil($count/$Everypage);
		foreach($result as $num=>$highlevellist){      //显示数据替换	
		if($highlevellist['highlevel']==0){
			$result[$num]['highlevel']='顶级分类';
			}
		if($highlevellist['categorystatus']==0){
		   $result[$num]['categorystatus']='正常';
		    }
			else{
		$result[$num]['categorystatus']='禁用';
			}
		}
        $this->assign('count',$count);
		$this->assign('pagecount',$pagecount);
		$this->assign('data',$result);
	*/	

		$this->testst="1";
		$count=$Model->count();
		$this->assign('count',$count);
		$list = $Model->page($_GET['p'].',5')->select();
	/*	foreach($list as $num=>$highlevellist){      //显示数据替换	
			if($highlevellist['highlevel']==0){
				$list[$num]['highlevel']='顶级分类';
				}
			if($highlevellist['categorystatus']==0){
			   $list[$num]['categorystatus']='正常';
			    }
				else{
			$list[$num]['categorystatus']='禁用';
				}
			}
			*/
		$this->assign('data',$list);// 赋值数据集
		import(“ORG.Util.Page”);// 导入分页类
		$count      = $Model->count();// 查询满足要求的总记录数
		$pagecount=ceil($count/5);
		$Page       = new Page($count,5);// 实例化分页类 传入总记录数和每页显示的记录数
		$show       = $Page->show();// 分页显示输出
		//var_dump($show);  这里已经把所有分页链接生成
		$this->assign('pagecount',$pagecount);
		$this->assign('_page',$show);// 赋值分页输出
		$this->display(); // 输出模板
		/*$Home=A('Home/Index');
		$kuai=$Home->testA();
		$this->assign('testa',$kuai);
		$Second=R('Model/showName');  //A与R方法写法上不一样
		$this->assign('second',$Second);  */
        //$this->display();
      
         
    



	}

    /* 编辑分类 */
    public function edit(){
		$Model=D('addcategory'); 
        if(!IS_POST){
	 //显示start	
	    $id=$_GET['id'];
		$result=$Model->where("id=$id")->select();
	    if($result[0]['categorystatus']=="0"){
		 $result[0]['categorystatus']="正常";
		  }
		  else{
		 $result[0]['categorystatus']="禁用";
		  }
		$this->assign('categorylist',$result);
		$all=$Model->field('categoryname')->select();
        $this->assign('alll',$all);
		$this->display();
      //显示end
	   }
		else{  //修该
	     $uid=$_POST['uid'];
		 $data=$Model->create(); //create可以创建那边form表单内POST的字段
	     if(!$data['categorystatus']=="禁用"){
		 $data['categorystatus']=0;
		   }
		 else{
		 $data['categorystatus']=1;
		   }
		  $res=$Model->where("id=$uid")->save($data);   //TMD把这里的'id=$uid'改成"id=$uid"就成功了，也是醉了 ！！！！！！！！！！！
			if($res){
			$this->success('修改成功！',U('Index'));
			}
			else{
			$this->error(empty($error) ? '未知错误！' : $error);
			}	
		  }
	}	
	 public function delete(){
		$id=$_GET['id'];
		if(IS_GET){             //get与post要分清楚
		$Mo=M('addcategory');
		$update=$Mo->where("id=$id")->delete();
		if($update){
		$this->success('删除成功！',U('Index'));
		}
		else{
		$this->error(empty($error) ? '未知错误！' : $error);
		}
	   }
		
	}

	public function add(){
		if(!IS_POST){
		$Model=M('addcategory');
        $resul=$Model->field('categoryname')->select();
		$this->assign('categorylist',$resul);
		$this->display();
		}
		else{
	  $Model=M('addcategory');
	  $arr=$Model->create();
	  $resu=$Model->add($arr);
	if($resu){
	$this->success('添加成功',U('Index')); 
    	 }
		 else{
	$this->error('添加失败',U('Index')); 
		 }
		}
	
	}

	public function dochange(){
       $Model=D('addcategory');
       $id=$_GET['id'];
       $data['categorystatus']=intval($_GET['categorystatus']);
       $ress=$Model->where("id=$id")->save($data);
        if($ress)
         {
         $this->success('修改成功！',U('index'));
         }
         else
         {
         $this->error('修改失败',U('index'));
         }

	}
}
