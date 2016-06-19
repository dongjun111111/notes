var starDate, allDate;
function HS_DateAdd(interval,number,date){
 number = parseInt(number);
 if (typeof(date)=="string"){var date = new Date(date.split("-")[0],date.split("-")[1],date.split("-")[2])}
 if (typeof(date)=="object"){var date = date}
 switch(interval){
  case "y":return new Date(date.getFullYear()+number,date.getMonth(),date.getDate()); break;
  case "m":return new Date(date.getFullYear(),date.getMonth()+number,checkDate(date.getFullYear(),date.getMonth()+number)); break;
  case "d":return new Date(date.getFullYear(),date.getMonth(),date.getDate()+number); break;
  case "w":return new Date(date.getFullYear(),date.getMonth(),7*number+date.getDate()); break;
 }
}
function checkDate(year,month){
 return new Date(year, month, '01').getDate();
}

function WeekDay(date){
 var theDate;
 if (typeof(date)=="string"){theDate = new Date(date.split("-")[0],date.split("-")[1],date.split("-")[2]);}
 if (typeof(date)=="object"){theDate = date}
 return theDate.getDay();
}

function HS_calender(){
 var lis = "";
 var style = "";
 style +="<style type='text/css'>";
 style +=".calender {z-index:99;position:absolute; width:170px; height:auto; font-size:12px; margin-top:20px; margin-left:-150px; background:#fff; border:1px solid #CCC; padding:1px;color:#000}";
 style +=".calender ol {width:170px;list-style-type:none; margin:0; padding:0;border:0;}";
 style +=".calender .day { background-color:#EDF5FF; height:20px;}";
 style +=".calender .day span,.calender .date span{float:left; display:inline-block; width:14%; height:20px; line-height:20px; text-align:center; margin:0; padding:0;border:0;color:#000;}";
 style +=".calender span a { text-decoration:none; font-family:Tahoma; font-size:11px; color:#333}";
 style +=".calender span a.hasArticle {font-weight:bold; color:#f60 !important}";
 style +=".lastMonthDate, .nextMonthDate {color:#bbb;font-size:11px}";
 style +=".selectThisYear a, .selectThisMonth a{text-decoration:none; margin:0 2px; color:#000; font-weight:bold}";
 style +=".calender .LastMonth, .calender .NextMonth{ text-decoration:none; color:#000; font-size:18px; font-weight:bold; line-height:16px;}";
 style +=".calender .LastMonth { float:left;}";
 style +=".calender .NextMonth { float:right;}";
 style +=".calenderBody {clear:both}";
 style +=".calenderTitle {text-align:center;height:20px; line-height:20px; }";
 style +=".today,.calender span a:hover { background-color:#ffffaa;border:1px solid #f60; padding:2px}";
 style +=".today a { color:#f30; }";
 style +=".calenderBottom {border-bottom:1px solid #CCC; height:20px; line-height:20px;text-align:left}";
 style +=".calenderBottom a {text-decoration:none; margin:2px !important; color:#000}";
 style +=".calenderBottom a.closeCalender{float:right}";
 style +=".closeCalenderBox {float:right; border:1px solid #000; background:#fff; font-size:9px; width:11px; height:11px; line-height:11px; text-align:center;overflow:hidden; font-weight:normal !important}";
 style +="</style>";

 var now;
 if (typeof(arguments[0])=="string"){
  selectDate = arguments[0].split("-");
  var year = selectDate[0];
  var month = parseInt(selectDate[1])-1+"";
  var date = selectDate[2];
  now = new Date(year,month,date);
 }else if (typeof(arguments[0])=="object"){
  now = arguments[0];
 }

 var lastMonthEndDate = HS_DateAdd("d","-1",now.getFullYear()+"-"+now.getMonth()+"-01").getDate();
 var lastMonthDate = WeekDay(now.getFullYear()+"-"+now.getMonth()+"-01");
 var thisMonthLastDate = HS_DateAdd("d","-1",now.getFullYear()+"-"+(parseInt(now.getMonth())+1).toString()+"-01");
 var thisMonthEndDate = thisMonthLastDate.getDate();
 var thisMonthEndDay = thisMonthLastDate.getDay();
 var todayObj = new Date();
 today = todayObj.getFullYear()+"-"+todayObj.getMonth()+"-"+todayObj.getDate();

 for (i=0; i<lastMonthDate; i++){  // Last Month's Date
  lis = "<span class='lastMonthDate'>&nbsp; </span>" + lis;
  lastMonthEndDate--;
 }

 for (i=1; i<=thisMonthEndDate; i++){ // Current Month's Date
  if(allDate) {
   if((new Date(now.getFullYear(), now.getMonth(), i)).getMonth() < (new Date()).getMonth() || (new Date(now.getFullYear(), now.getMonth(), i)).getFullYear() < (new Date()).getFullYear()) {
    lis += "<span><a disabled='disabled' title='"+todayString+"'>"+i+"</a></span>";
   } else if((new Date(now.getFullYear(), now.getMonth(), i)).getMonth() == (new Date()).getMonth()) {
    if(today == now.getFullYear()+"-"+now.getMonth()+"-"+i ){
     var todayString = now.getFullYear()+"-"+(parseInt(now.getMonth())+1).toString()+"-"+i;
     if(now.getDate() > (new Date()).getDate()) {
      lis += "<span><a disabled='disabled' class='today' title='"+todayString+"'>"+i+"</a></span>";
     } else {
      lis += "<span><a href=javascript:void(0) class='today' onclick='_selectThisDay(this)' title='"+todayString+"'>"+i+"</a></span>";
     }
    }else if(now.getDate() > i){
     lis += "<span><a disabled='disabled' title='"+now.getFullYear()+"-"+(parseInt(now.getMonth())+1)+"-"+i+"'>"+i+"</a></span>";
    } else {
     lis += "<span><a href=javascript:void(0) onclick='_selectThisDay(this)' title='"+now.getFullYear()+"-"+(parseInt(now.getMonth())+1)+"-"+i+"'>"+i+"</a></span>";
    }
   } else {
    lis += "<span><a href=javascript:void(0) onclick='_selectThisDay(this)' title='"+now.getFullYear()+"-"+(parseInt(now.getMonth())+1)+"-"+i+"'>"+i+"</a></span>";
   }
  } else {
   if(today == now.getFullYear()+"-"+now.getMonth()+"-"+i ){
    var todayString = now.getFullYear()+"-"+(parseInt(now.getMonth())+1).toString()+"-"+i;
    lis += "<span><a href=javascript:void(0) class='today' onclick='_selectThisDay(this)' title='"+todayString+"'>"+i+"</a></span>";
   } else {
    lis += "<span><a href=javascript:void(0) onclick='_selectThisDay(this)' title='"+now.getFullYear()+"-"+(parseInt(now.getMonth())+1)+"-"+i+"'>"+i+"</a></span>";
   }
  }
 }

 var j=1;
 for (i=thisMonthEndDay; i<6; i++){  // Next Month's Date
  lis += "<span class='nextMonthDate'>&nbsp; </span>";
  j++;
 }
 lis += style;

 var CalenderTitle = "<a href='javascript:void(0)' class='NextMonth' onclick=HS_calender(HS_DateAdd('m',1,'"+now.getFullYear()+"-"+now.getMonth()+"-"+now.getDate()+"'),this) title='下月'>&raquo;</a>";
 CalenderTitle += "<a href='javascript:void(0)' class='LastMonth' onclick=HS_calender(HS_DateAdd('m',-1,'"+now.getFullYear()+"-"+now.getMonth()+"-"+now.getDate()+"'),this) title='上月'>&laquo;</a>";
 CalenderTitle += "<span class='selectThisYear'><a href='javascript:void(0)' onclick='CalenderselectYear(this)' title='点这里选择年份' >"+now.getFullYear()+"</a></span>年<span class='selectThisMonth'><a href='javascript:void(0)' onclick='CalenderselectMonth(this)' title='点这里选月'>"+(parseInt(now.getMonth())+1).toString()+"</a></span>月";

 if (arguments.length>1){
  arguments[1].parentNode.parentNode.getElementsByTagName("ol")[1].innerHTML = lis;
  arguments[1].parentNode.innerHTML = CalenderTitle;
 }else{
  var CalenderBox = style+"<div class='calender'><div class='calenderBottom'><a href='javascript:void(0)' class='closeCalender' onclick='closeCalender(this)'>×</a><span></span></div><div class='calenderTitle'>"+CalenderTitle+"</div><div class='calenderBody'><ol class='day'><span>日</span><span>一</span><span>二</span><span>三</span><span>四</span><span>五</span><span>六</span></ol><ol class='date' id='thisMonthDate'>"+lis+"</ol></div></div>";
  return CalenderBox;
 }
}
function _selectThisDay(d){
 var boxObj = d.parentNode.parentNode.parentNode.parentNode.parentNode;
 boxObj.targetObj.value = d.title;
 boxObj.parentNode.removeChild(boxObj);
}
function closeCalender(d){
 var boxObj = d.parentNode.parentNode.parentNode;
 boxObj.parentNode.removeChild(boxObj);
}

function CalenderselectYear(obj){
 var opt = "";
 var thisYear = obj.innerHTML;
 for (i=1970; i<=2020; i++){
  if (i==thisYear){
   opt += "<option value="+i+" selected>"+i+"</option>";
  }else{
   opt += "<option value="+i+">"+i+"</option>";
  }
 }
 opt = "<select onblur='selectThisYear(this)' onchange='selectThisYear(this)' style='font-size:11px'>"+opt+"</select>";
 obj.parentNode.innerHTML = opt;
}

function selectThisYear(obj){
 HS_calender(obj.value+"-"+obj.parentNode.parentNode.getElementsByTagName("span")[1].getElementsByTagName("a")[0].innerHTML+"-1",obj.parentNode);
}

function CalenderselectMonth(obj){
 var opt = "";
 var thisMonth = obj.innerHTML;
 for (i=1; i<=12; i++){
  if (i==thisMonth){
   opt += "<option value="+i+" selected>"+i+"</option>";
  }else{
   opt += "<option value="+i+">"+i+"</option>";
  }
 }
 opt = "<select onblur='selectThisMonth(this)' onchange='selectThisMonth(this)' style='font-size:11px'>"+opt+"</select>";
 obj.parentNode.innerHTML = opt;
}
function selectThisMonth(obj){
 HS_calender(obj.parentNode.parentNode.getElementsByTagName("span")[0].getElementsByTagName("a")[0].innerHTML+"-"+obj.value+"-"+starDate,obj.parentNode);
}

function HS_setDate(inputObj,ID,show){
 var calenderObj = document.createElement("span");
 if(ID) {
  selectDate = (document.getElementById(ID).value)?(document.getElementById(ID).value):(new Date());
  starDate = selectDate.split("-")[2];
 } else {
  selectDate = new Date();
  starDate = new Date().getDate();
 }
 allDate = show?false:true;
 calenderObj.innerHTML  = HS_calender(selectDate);
 calenderObj.style.position = "absolute";
 calenderObj.targetObj  = inputObj;
 inputObj.parentNode.insertBefore(calenderObj,inputObj.nextSibling);
}