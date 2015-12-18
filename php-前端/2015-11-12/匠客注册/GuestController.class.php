<?php 
namespace Admindj\Controller;
/**
     * 云匠客后台控制器
*/
class GuestController extends AdminController {


	public function kparate() {
		$get = I('get.');
		if (isset($get['id'])) {
            $map['id'] = $get['id'];
        }

        if (isset($get['did'])) {
            $map['des_id'] = $get['did'];
        }
		$list = paging('Keseparate', $map, $field = true, $c = '50', $order = 'ctime desc',$pamap);
		$allnum = sizeof($list['list']);
		//var_dump($list['list']);
		$this->assign('arr_nums', $allnum);
		$this->assign('list', $list["list"]);
		$this->assign('_page', $list['show']);
		$this->meta_title = '云匠客开启表';
        $this->display();
	}

	public function keseller() {
		$get = I('get.');
		if (isset($get['id'])) {
            $map['id'] = $get['id'];
        }

        if (isset($get['kid'])) {
            $map['kid'] = $get['kid'];
        }
		$list = paging('Keseller', $map, $field = true, $c = '50', $order = 'id desc',$pamap);
		$allnum = sizeof($list['list']);
		$this->assign('arr_nums', $allnum);
		$this->assign('list', $list["list"]);
		$this->assign('_page', $list['show']);
		$this->meta_title = '云匠客雇主绑定表';
        $this->display();
	}

	public function jklist() {
		$get = I('get.');
		if (isset($get['kid'])) {
            $map['id'] = $get['kid'];
        }
        if (isset($get['tel'])) {
            $map['lxtel'] = $get['tel'];
        }
		$list = paging('Jiangke', $map, $field = true, $c = '50', $order = 'id desc',$pamap);
		$allnum = sizeof($list['list']);
		$this->assign('arr_nums', $allnum);
		$this->assign('list', $list["list"]);
		$this->assign('_page', $list['show']);
		$this->meta_title = '云匠客用户表';
        $this->display();
	}

	public function klist(){
		$sql1 = 'SELECT a.*,count(b.id) as p_num
		FROM ks_jiangke a inner join ks_project b on a.id = b.kid
		WHERE ( b.is_tuog = 1 and b.kid > 0 ) 
		GROUP BY b.kid ORDER BY p_num desc ';
		$c_info = M('jiangke')->query($sql1);
		$count = sizeof($c_info);


		$style = array('css_prev'=>'prev','css_num'=>'num','css_current'=>'current','css_next'=>'next','css_end'=>'num','prev'=>'上一页','next'=>'下一页','css_first'=>'prev','css_end'=>'next','first'=>'上一页','last'=>'下一页','convert_url'=>true); 
      	$arr_pages = pages($count,20,$map,$style);

		$sql = 'SELECT a.*,count(b.id) as p_num
		FROM ks_jiangke a inner join ks_project b on a.id = b.kid
		WHERE ( b.is_tuog = 1 and b.kid > 0 ) 
		GROUP BY b.kid ORDER BY p_num desc 
		LIMIT '.$arr_pages['first'].','.$arr_pages['lim'];
		$info = M('jiangke')->query($sql);
		foreach ($info as $key => $value) {
			if($value['type'] == 1){$maps = ' and keseparate > 0';}else{$maps = '';}
			$a_pid = M('Project')->where('kid='.$value['id'].' and is_tuog = 1'.$maps)->getField('id',true);
			$info[$key]['p_num'] = sizeof($a_pid);
			$arr_pid = implode(",", $a_pid);
        	$p_a = M('Order')->where('pid IN ('.$arr_pid.')')->getField('allprice',true);
    		$info[$key]['allprice'] = array_sum($p_a); 
    		$info[$key]['yprice'] =  M('jiangke_record')->where('kid='.$value['id'].' and pid in('.$arr_pid.') and rtype = 3')->sum('userb');
    		$p_o = M('jiangke_record')->where('kid='.$value['id'].' and pid in('.$arr_pid.') and rtype = 3')->getField('ordersn',true);
    		$arr_or = implode(",", $p_o);
    		$info[$key]['sprice'] = M('Cash_shift')->where('ordersn in('.$arr_or.')')->sum('avalue');
    		$info[$key]['b_num'] = M('keseller')->where('kid='.$value['id'])->count();
    		$info[$key]['yprice'] = $info[$key]['yprice']?$info[$key]['yprice']:0;
    		$info[$key]['sprice'] = $info[$key]['sprice']?$info[$key]['sprice']:0;
		}
		//var_dump(M('jiangke')->_sql());
		$allnum = sizeof($info);
        $this->assign('arr_nums', $allnum);
        $this->assign('list',$info);
        $this->assign('_page',$arr_pages);
        $this->meta_title = '云匠客统计表';
        $this->display();

	}



	public function plist(){
		$id = I('get.yid');
		if (isset($get['sid'])) {
            $map['sid'] = $get['sid'];
        }
		$map['kid'] = $id;
		$map['is_tuog'] = 1;
		$type = M('Jiangke')->where('id='.$id)->getField('type');
		if($type == 1){
			$map['keseparate'] =array('gt',0);
		}
		//$field = 'id,sid,kid,ctype,kestatus,wjtime,wjstatus,payprice,status,begintime,endtime';
		$list = paging('Project', $map, $field = true, $c = '30', $order = 'ctime desc',$pamap);
		foreach ($list['list'] as $key => $value) {
			$y_info = M('jiangke_record')->where('pid='.$value['id'].' and rtype = 3')->find();
			$list['list'][$key]['yprice'] = $y_info['userb'];
			$list['list'][$key]['sprice'] = M('Cash_shift')->where('ordersn='.$y_info['ordersn'])->getField('avalue');
			$list['list'][$key]['allprice'] = M('order')->where('pid = '.$value['id'])->getField('allprice');
			$list['list'][$key]['yprice'] = $list['list'][$key]['yprice']?$list['list'][$key]['yprice']:0;
			$list['list'][$key]['sprice'] = $list['list'][$key]['sprice']?$list['list'][$key]['sprice']:0;
		}
		//var_dump($list['list']);
		$allnum = sizeof($list['list']);
		$this->assign('arr_nums', $allnum);
		$this->assign('list', $list["list"]);
		$this->assign('_page', $list['show']);
		$this->meta_title = '云匠客订单表';
        $this->display();
	}

	public function glist(){
		//$ids = M('Jiangke')->where('type = 2')->getField('id',true);
		$map['kid'] = array('gt',0);
		if (isset($_GET['sid'])) {
            $pamap['sid'] = $map['sid'] = $$_GET['sid'];
        }
        if (isset($_GET['type'])) {
            $ids = M('Jiangke')->where('type = '.$_GET['type'])->getField('id',true);
            $a_id = implode(",", $ids);
            $pamap['type'] = $_GET['type'];
            $map['kid'] = array('IN',$a_id);
            $this->assign('type', $_GET['type']);
        }
		$list = paging('Project', $map, $field = true, $c = '30', $order = 'ctime desc',$pamap);
		foreach ($list['list'] as $key => $value) {
			$type = M('Jiangke')->where('id='.$value['kid'])->getField('type');
			$list['list'][$key]['k_type'] = $type;
			if($type == 1){
				$y_info = M('jiangke_record')->where('pid='.$value['id'].' and rtype = 3')->find();
				$list['list'][$key]['yprice'] = $y_info['userb'];
				$list['list'][$key]['sprice'] = M('Cash_shift')->where('ordersn='.$y_info['ordersn'])->getField('avalue');
				$list['list'][$key]['allprice'] = M('order')->where('pid = '.$value['id'])->getField('allprice');
				$list['list'][$key]['yprice'] = $list['list'][$key]['yprice']?$list['list'][$key]['yprice']:0;
				$list['list'][$key]['sprice'] = $list['list'][$key]['sprice']?$list['list'][$key]['sprice']:0;
			}elseif($type == 2){
				$o_p = M('order')->where('pid='.$value['id'])->find();
				$o_c = M('Cash_shift')->where('ordersn='.$o_p['ordersn'])->find();
				if($o_c){
					$o_c['avalue'] = $o_c['avalue']?$o_c['avalue']:0;
					$list['list'][$key]['sprice'] = $o_c['avalue'];
					$list['list'][$key]['yprice'] = $o_c['avalue'] * 0.2;
				}else{
					$ol_p = M('Order_all')->where('ctype = "xuqian" and xqstatus = 1')->getField('ordersn');
					$a_o = implode(",", $ol_p);
					$time = $o_p['endtime'] + 5*86400; 
					$avalue =M('Cash_shift')->where('ordersn IN ('.$a_o.') and ctime <='.$time)->getField('avalue');
					$list['list'][$key]['sprice'] = $avalue?$avalue:0;
					$list['list'][$key]['yprice'] = $avalue * 0.2;
				}
				$list['list'][$key]['allprice'] = $o_p['allprice'];
			}
			
		}
		//var_dump($list['list']);
		$allnum = sizeof($list['list']);
		$this->assign('arr_nums', $allnum);
		$this->assign('list', $list["list"]);
		$this->assign('_page', $list['show']);
		$this->meta_title = '云匠客项目表';
        $this->display();
	}


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