/*!
*description: delay for AllyesDeliver
*email:lei_ding@allyes.com
*date:2013.12.1
*version:2.1
*/
!function(a){var b,c,d;"undefined"!=typeof AllyesDeliver&&(d=function(b){a.document.write(b)},b=a.allyes_inter||0,c=AllyesDeliver.cdnSrc("delay.min.js"),++b<5?(a.allyes_inter=b,d('<script type="text/javascript" data-belong="allyes" src="'+c+'"></script>')):(d("</div>"),b=null,c=null,a.allyes_inter=null,AllyesDeliver.checking()))}(window);