<?php 
namespace Admindj\Controller;
class GuestController extends AdminController {
	public function register(){		
		if(IS_POST){
			/*function magic($str){
				if(!get_magic_quotes_gpc()){
					$str=addslashes($str);
				}
				return $str;
			}
			*/
			function randkcode(){
				$str="qwertyuioplkjhgfdsazxcvbnm";
				for($i=0;$i<6;$i++){
					$t=rand(0,strlen($str)-1);
					$strr.=$str{$t};
				}
				return $strr;
			}
			function randsalt(){
				$str="1234567890";
				for($i=0;$i<6;$i++){
					$t=rand(0,strlen($str)-1);
					$strr.=$str{$t};
				}
				return $strr;
			}
			$map['username']=I("post.username");
			$map['lxtel']=I("post.lxtel");
			$map['email']=I("post.email");
			$map['kcode'] =randkcode();
			$map['salt'] =randsalt();
			$map['password']=md5(I("post.password").$map['salt']);
			$map['paypass']=md5(I("post.password").$map['salt']);
			$map['reg_time']=time();
			$map['reg_ip']=ip2long(get_client_ip());
			$map['type'] =2;
			$map['status'] =1;
			$repassword=I("post.repassword");			
			if(I("post.password") != $repassword){
				$this->error("确认密码出现错误！",U("Admindj/Guest/register"));
			}
			if(!check_verify($_POST['verify']))
			        {
			            $this->error("验证码错误！",U("Admindj/Guest/register"));
			}
			$rules = array(
			    array('verify','require','验证码必须！'), //默认情况下用正则进行验证
			    array('username','',"
			    	<html><body><p>帐号名称已经存在!2秒后自动返回...</p><script>
					setTimeout('window.history.back()',2000);
                </script></body></html>",0,'unique',1), // 在新增的时候验证name字段是否唯一
			    array('repassword','password','确认密码不正确',0,'confirm'), // 验证确认密码是否和密码一致
			    array('password','checkPwd','密码格式不正确',0,'function'), // 自定义函数验证密码格式
			);
			$jiangke = M("jiangke"); // 实例化
			if (!$jiangke->validate($rules)->create()){
				exit($jiangke->getError());
			}else{
				$res=$jiangke->add($map);
				if($res){
					$this->success("注册成功！",U("Admindj/Guest/jklist"));
				}
				else{
					$this->error("注册失败!",U("Admindj/Guest/register"));
				}
			}
			


		}
		else{		
			$this->display();
		}

	}

}