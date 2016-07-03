#C
###用c实现小闹钟
<pre>
#include <stdio.h>
#include <time.h>
int jason();
void sleep(); 
int main(){
	printf("小闹钟，Now it is working...\n");
	while(1){
	time_t nowtime;
	struct tm *timeinfo;
	time(&nowtime);
	timeinfo = localtime(&nowtime);
	int year,month,day,hour,minute,second;
	year = timeinfo->tm_year +1900;
	month = timeinfo->tm_mon +1;
	day = timeinfo->tm_mday;
	hour = timeinfo->tm_hour;
	minute = timeinfo->tm_min;
	second = timeinfo->tm_sec;
		if(hour == 8 && minute == 8 && second == 0){
			printf("%d年%d月%d日 %d时%d分%d秒\n",year,month,day,hour,minute,second);
			printf("\a\a\a\a\a\a\a\a\a\a");		 			
		}else{			
		}
		sleep(1000);
			
	}
	return 0;
}

int jason(){
	int i;
	for (i =0;i<10;i++){
		printf("\a");
	}
} 
//这个函数流弊 
void sleep(long wait){
	long goal=clock()+wait;
	while(goal>clock());
}
/*
	tm_year 从1900年计算，所以要加1900
    tm_mon，从0计算，所以要加1
    
    struct tm -- 时间结构，time.h 定义如下： 
	int tm_sec; 
	int tm_min; 
	int tm_hour; 
	int tm_mday; 
	int tm_mon; 
	int tm_year; 
	int tm_wday; 
	int tm_yday; 
	int tm_isdst; 
*/
/*自己实现Sleep类型功能的函数 
#include<stdio.h>
#include<time.h>
main()
{
    void sleep(long wait);       
    sleep(1000);
    printf("hello!");
    return 0;
}
void sleep(long wait)
{
	long goal=clock()+wait;
	while(goal>clock());
}
*/
</pre>