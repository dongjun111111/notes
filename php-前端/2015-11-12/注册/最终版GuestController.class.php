<?php 
namespace Admindj\Controller;

	public function autoregister($data){  

		if($data['username'] && $data['lxtel'] && $data['password']){
			$username = $data['username'];
			$lxtel = $data['lxtel'];
			$password = $data['password'];
			$type = $data['type'];
		}else{ 
			return false;
		}  
		if(empty($lxtel)){
			return false;
		}
		if(empty($password)){
			return false;
		} 
        if(!check_empty_char($lxtel)){
              return false;
        }
        if(!check_empty_char($password)){
               return false;
        } 
        if(doreg($lxtel,'cell')){
            //$this->error = '手机格式不正确1'; return false;
        }
   		$einfo = M('jiangke')->where('lxtel="'.$lxtel.'"')->find();        
		if($einfo){return false;}				
		$model = M(); 
        $model->startTrans();		        
		$data1 = array();
		$data1['username'] = $lxtel;  
		$data1['salt'] =  getRandPwd(1,6);
		$data1['password'] = D('Ke/jiangke')->setPass($password,$data1['salt']);
		$data1['last_login_ip'] = $data1['reg_ip'] =  getIP(); 
		$data1['last_login_time'] = $data1['reg_time'] =  time();   
		$data1['lxtel'] = $lxtel;
		$data1['type'] = $type;  
		$data1['kcode'] =  strtolower(getRandPwd(2,6));
        $id =  M('jiangke')->add($data1);
        if($id){    
            $arr = array(); 
            $bkey = 0;
            $ykey = 0;
            $akey = 0;
            $mdate = time();
            $BKEY = '7d5f48d731695f9383ac04969550dfd6';
            $YKEY = '51c3cd83a438aadcc019938ec77c5f25';
            $AKEY = 'edd217b749a14789895e055982d85201';
            $sql = 'INSERT INTO ks_kaccount (uid,userb,usery,usera,mdate,bkey,ykey,akey,status) VALUES ("'.$id.'","0","0","0",'.$mdate.',ENCODE('.$bkey.',"'.$BKEY.'"),ENCODE('.$ykey.',"'.$YKEY.'"),ENCODE('.$akey.',"'.$AKEY.'"),1)';
            $rs = M('kaccount')->execute($sql);
            if($rs == false){
				$model->rollback();	
				return false;
			}
		} else { 
			$model->rollback();	
			return false;
		}				
		$model->commit(); 
		return true; 
    }
	public  function gJiangke($num){   
		$j = 0;
		while($j < $num){
	        $randun='12'.getRandPwd(1,9);
	        for($i=1;$i<=10;$i++){
	            $kid=M('jiangke')->where('username = "'.$randun.'" or lxtel = "'.$randun.'"')->getField('id');
	            if(empty($kid)){
	                $username =$randun;
	                break;
	            }
	        }
	        if(empty($username)) continue;
	        $data['password'] = 'a123456';
	        $data['username'] = $username;
	        $data['type'] = 2;
	        $data['lxtel'] = $username;
	     	$id = $this -> autoregister($data);
	        if(0 < $id){
	            echo intval($j+1).'-->'.$data['username']." was  successful<br>";
	        }
	        else{
	        	echo "<span style='color:red'>".intval($j+1).'-->'.$data['username']." was failed</span><br>";
	        	continue; 
	        } 
	    	$j++;
	    }
		return true;
	}
	public function register(){	
		$num = I('post.num',1);
		$preg="/^[0-9]$/";
		preg_match_all($preg,$num,$arr);
		if(!$arr){
			$this->error("输入格式有误",U("Admindj/Guest/register"));
		}
		if($num<=0 || $num>30){
			$this->error("输入数量有误!",U("Admindj/Guest/register"));
		}  
		if(IS_POST){
			$res = $this -> gJiangke($num);
			if($res){
				$this->success("添加成功",U("Admindj/Guest/jklist"),6);
			}
			else{
				$this->error("添加失败!",U("Admindj/Guest/register"));
			} 
		}
		else{			
			$this->display();
		}

	}

	 
}