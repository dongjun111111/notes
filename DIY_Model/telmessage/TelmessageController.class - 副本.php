<?php 
namespace Admindj\Controller;

class TelmessageController extends AdminController {
    public function index(){    
		 $my=D('Telmessage');
		 $array=$my->getData();
         if(isset($_GET['tel'])){
            $map['tel']=$_GET['tel'];
         }
         if(isset($_GET['ctime'])){
            $map['ctime'] =  $pamap['ctime']= array('LIKE' , '%'.$_GET['ctime'].'%');
         }

        $list = paging('tel_message', $map, true, '30', 'id desc', $pamap);
        $this->assign('data',$list[0]);

		 $this->assign('_page',$array['show']);
         foreach($array['data'] as $num=>$value){
                $array['data'][$num]['ctime']=date('Y-m-d H:i:s',$array['data'][$num]['ctime']);
         }
		 $this->assign('data',$array['data']);
         $this->assign('count',$array['count']);
		 $this->display();
    }
	
	  public function edit(){
    	$my=D('Telmessage');
      if(IS_POST){
             $Model=D('tel_message');
             $data=$Model->create();
             $id=intval($data['id']);
             array_splice($data,0,1);
             $data['ctime']=time();
             $res=$Model->where("id=$id")->save($data);   
             if(!$res){
                $this->error(empty($error)?"Î´Öª´íÎó":$error);
             }
             else{
                $this->success('±à¼­³É¹¦',U('Admindj/Telmessage/index'));
             }
   		}
   	   else{
        $array=$my->edit();
        $this->assign('data',$array['data']);
        $this->display();
       }
    }

    public function delete(){
    	$my=D('tel_message');
    	$id=intval($_GET['id']);
    	$resu=$my->where("id=$id")->limit(1)->delete();
    	if(!$resu){
            $this->error('Error',U('Admindj/Telmessage/index'));
    	}
    	else{
    		$this->success('Success',U('Admindj/Telmessage/index'));
    	}
    }
	public function add(){
	    if(IS_POST){
            $model=D('Telmessage');
            $data=$model->addData();
            $my=D('tel_message');
            $resul=$my->add($data);
            if(!$resul){
                $this->error('Error',U('Admindj/Telmessage/index'));
            }
            else{
                $this->success('Success',U('Admindj/Telmessage/index'));
            }
            
        }
        else{
            $this->display();
			
        }
	}
    //电话号码搜索，时间搜索
	public function search(){
        if(empty($search)){
        $this->redirect("Admindj/Telmessage/index");
        }
        else{
        $Model=D('tel_message');
        $map = array();   
        $map  = array('status' => 1);

            if(isset($_GET['group'])){
            $map['group']   =   I('group',0);
            }
            if(isset($_GET['name'])){
            $map['name']    =   array('like', '%'.(string)I('name').'%');
        }

        $data = $this->lists('Telmessage', $map,'id');  
		$this->assign('data',$data);
		$this->display('index');	
	  }
	}
	function myPaging($mod,$map,$field=true,$limit='10',$order='id desc',$pamap,$stime,$etime) {
        $model=M($mod);
        $count = $model->where($map)->count();
        $Page = new \Think\Page($count,$limit);
                $Page->setConfig('next', '下一页');
                $Page->setConfig('prev', '上一页');
        //分页跳转的时候保证查询条件
        if(empty($pamap)) {$pamap =  $map ;}
        foreach($pamap as $key=>$val){
            $Page->parameter[$key] = urlencode($val); //urlencode 将字符串以URL编码 不支持数组
        }
        $show = $Page->show();// 分页显示输出
        empty($map['ctime'])?0:array('between','$stime,$etime');
        // 进行分页数据查询 注意limit方法的参数要使用Page类的属性
        $list = $model->where($map)->order($order)->limit($Page->firstRow.','.$Page->listRows)->select();
        $arr=array();
        $arr['list']=$list;
        $arr['show']=$show;
                $arr['tolnum'] = $count;
        return $arr;
         
     }

}
