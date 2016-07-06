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
###extern 关键字释解
1 基本解释：extern可以置于变量或者函数前，以标示变量或者函数的定义在别的文件中，提示编译器遇到此变量和函数时在其他模块中寻找其定义。此外extern也可用来进行链接指定。

也就是说extern有两个作用:

- 第一个,当它与"C"一起连用时，如: extern "C" void fun(int a, int b);则告诉编译器在编译fun这个函数名时按着C的规则去翻译相应的函数名而不是C++的，C++的规则在翻译这个函数名时会把fun这个名字变得面目全非，可能是fun@aBc_int_int#%$也可能是别的，这要看编译器的"脾气"了(不同的编译器采用的方法不一样)，为什么这么做呢，因为C++支持函数的重载啊，在这里不去过多的论述这个问题，如果你有兴趣可以去网上搜索，相信你可以得到满意的解释!
- 第二，当extern不与"C"在一起修饰变量或函数时，如在头文件中: extern int g_Int; 它的作用就是声明函数或全局变量的作用范围的关键字，其声明的函数和变量可以在本模块活其他模块中使用，记住它是一个声明不是定义!也就是说B模块(编译单元)要是引用模块(编译单元)A中定义的全局变量或函数时，它只要包含A模块的头文件即可,在编译阶段，模块B虽然找不到该函数或变量，但它不会报错，它会在连接时从模块A生成的目标代码中找到此函数。

示例：

main.c 
<pre>
#include <stdio.h>
 
int count ;
extern void write_extern();
 
main()
{
   count = 5;
   write_extern();
}
</pre>
support.c
<pre>
#include <stdio.h>
 
extern int count;
 
void write_extern(void)
{
   printf("count is %d\n", count);
}
</pre>
然后切换到终端：
<pre>
gcc main.c support.c   //成功则生成a.out
./a.out                //运行 count is 5
</pre>