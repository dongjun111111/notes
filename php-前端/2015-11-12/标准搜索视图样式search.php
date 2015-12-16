<!--搜索教科书般的样式-->
<div class="search-form fr cf">
            <div class="btn-group-click adv-sch-pannel fl">
                <div class="cf">
                	<div class="row">
                		<label>项目时间：</label>
                		<input type="text" id="time-start" name="begintime" class="text input-2x" value="{:I('begintime')}" placeholder="起始时间" /> -                		
                        <div class="input-append date" id="datetimepicker"  style="display:inline-block">
                            <input type="text" id="time-end" name="endtime" class="text input-2x" value="{:I('endtime')}" placeholder="结束时间" />
                            <span class="add-on"><i class="icon-th"></i></span>
                        </div>
                	</div>
                </div>
            </div>
            <div class="sleft">
				<div class="drop-down">
					<span id="sch-sort-txt" class="sort-txt" data="{$status}"><if condition="$status">{$status|get_pro_status}<else/>所有</if></span>
                                        <i class="arrow arrow-down"></i>
                    
                    <ul id="sub-sch-menu" class="nav-list hidden">
						<li><a href="javascript:;" value="">所有</a></li>
                         <foreach name="prostatus" item="_list">
							<li><a href="javascript:;" value="{$key}">{$_list}</a></li>
						</foreach>
					</ul>   
				</div>
				<input type="text" name="id" class="search-input" value="{:I('id')}" placeholder="请输入ID">
                <input type="text" name="did" class="search-input" value="{:I('did')}" placeholder="设计师ID" style="width:100px;">
                <input type="text" name="dname" class="search-input" value="{:I('dname')}" placeholder="设计师花名" style="width:100px;">
                <input type="text" name="bid" class="search-input" value="{:I('bid')}" placeholder="雇主ID" style="width:100px;">
                <input type="text" name="bname" class="search-input" value="{:I('bname')}" placeholder="雇主花名" style="width:100px;">
				<a class="sch-btn" href="javascript:;" id="search" url="{:U('Admindj/project/plist')}"><i class="btn-search"></i></a>
			</div> 
		</div>