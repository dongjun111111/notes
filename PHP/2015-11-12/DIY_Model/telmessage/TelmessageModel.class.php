<?php 
namespace Admindj\Model;
use Think\Model;
use \Think\Page;
class TelmessageModel extends Model {

     public function getData(){
		 import('ORG.Util.Page');
		 $Model=D('tel_message');
		 if(empty($_GET['p'])){
        	$_GET['p']=1;
         }
		 $data=$Model->field(true)->page($_GET['p'].',30')->select();
		 $count=$Model->count();
         $Page=new Page($count,30);
         $show=$Page->show();
		 $array=array();
         $array['data']=$data;
         $array['show']=$show;
         $array['count']=$count;
         return $array;   
	}
     
	 public function edit(){
        $Model=D('tel_message');
 		$id=$_GET['id'];
 		if(!empty($id)){
        $data=$Model->field(true)->where("id=$id")->select();
        }
 		$array=array();
 		$array['data']=$data;
 		return $array;
	}
	public function addData(){ 
	  $Model=D('tel_message');
	  $data['tel']=$_POST['tel'];
	  $data['contents']=$_POST['contents'];
	  $data['request']=$_POST['request'];
	  $data['status']=$_POST['status'];
	  $data['ctime']=time();
	  return $data;
	}

	 public function lists($map, $field = true,$order = '`id` DESC', $limit = '0'){	
        return $this->field($field)->where($map)->order($order)->limit($limit)->select(); 
    }

}