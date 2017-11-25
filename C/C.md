# C
### 用c实现小闹钟
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
	year = timeinfo->tm_year + 1900;
	month = timeinfo->tm_mon + 1;
	day = timeinfo->tm_mday;
	hour = timeinfo->tm_hour;
	minute = timeinfo->tm_min;
	second = timeinfo->tm_sec;
		if(hour == 8 && minute == 8 && second == 0){
			printf("%d年%d月%d日 %d时%d分%d秒\n",year,month,day,hour,minute,second);
			printf("\a\a\a\a\a\a\a\a\a\a");		 			
		}else{			
		}
		sleep(1000) ;
			
	}
	return 0;
}

int jason(){
	int i;
	for (i =0;i<10;i++){
		printf("\a");
	}
} 
//这个函数很流弊 
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
### extern 关键字释解
1 基本解释：extern可以置于变量或者函数前，以标示变量或者函数的定义在别的文件中，提示编译器遇到此变量和函数时在其他模块中寻找其定义,此外extern也可用来进行链接指定。

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
//如果出现deny提示，则chmod ug+x 文件名  //为文件所有者和组添加执行权限
</pre>
由此引申：
<pre>
chmod u+x 文件名或者目录 //为此文件所有者添加执行权限
chmod g-x 文件名或者目录 //为组用户减去执行权限
chmod 777 文件名或者目录 //为所有用户添加可读可写可执行权限  （最高权限）
chmod 755 文件名或者目录 //为所有者添加读写与执行权限，组用户与其他用户添加读与执行权限
</pre>
###把一个字符串的大写字母放到字符串的后面，各个字符的相对位置不变，不能申请额外的空间
<pre>
#include <stdio.h>  
#include <string.h>  
//题目以及要求：把一个字符串的大写字母放到字符串的后面，  
//各个字符的相对位置不变，不能申请额外的空间。   
//判断是不是大写字母   
int isUpperAlpha(char c){  
if(c >= 'A' && c <= 'Z'){  
return 1;  
}  
return 0;   
}  
//交换两个字母   
void swap(char *a, char *b){  
char temp = *a;  
*a = *b;  
*b = temp;  
}   
char * mySort(char *arr, int len){  
if(arr == NULL || len <= 0){  
return NULL;  
}  
int i = 0, j = 0, k = 0;  
for(i = 0; i < len; i++){  
for(j = len - 1 - i; j >= 0; j--){  
if(isUpperAlpha(arr[j])){  
for(k = j; k < len - i - 1; k++){  
swap(&arr[k], &arr[k + 1]);  
}  
break;  
}  
//遍历完了字符数组，但是没发现大写字母，所以没必要再遍历下去  
if(j == 0 && !isUpperAlpha(arr[j])){  
//结束;  
                           return arr;  
}  
}  
}  
//over:   
return arr;  
}  
int main(){  
char arr[] = "aaaaaaaAbcAdeBbDc";  
printf("%s\n", mySort(arr, strlen(arr)));  
return 0;  
} 
output==>
aaaaaaabcdeAABD
</pre>
### C实现栈的线性表的链表数据结构
<pre>
//#include <stdio.h>
//#include <malloc.h>
//
//struct student
//{
// char score;
// struct student *next;
//};
//
//int main()
//{
// struct student *head,*p,*q;
// head =(struct student*)malloc(sizeof(struct student));
// if(!head)
// { 
//    printf("wrong！");
//    return 0;
//  }
// head->score=getchar();
// head->next=NULL;
// p=head;
//
// while(1)
// {
//   char c=getchar();
//   if(c=='\n')
//   break;
//   q = (struct student*)malloc(sizeof(struct student));
//   if(!q) return 0;
//   q->score=c;
//   p->next=q;
//   p=q;
//   p->next=NULL;
// }
//
// while(head!=NULL)
// { 
//   printf("%c",head->score);
//   head=head->next;
// }
// printf("\n");
// return 0;
//}   
//

#include <stdio.h>  
#include <stdlib.h>  
#include <time.h>   
#define ERROR 0  
#define OK 1  
#define MAXSIZE 20  
#define  random(x) (rand()%x)  
  
typedef int SElemType, Status;                                
  
typedef struct {    //定义一个结构体栈并且用字符串SqStack来代表该结构体类型   
    SElemType data[MAXSIZE] ;  
    int top;  
} SqStack;  
  
//在栈中插入元素 e   
Status push(SqStack *s, SElemType e) {  
    if(s->top == MAXSIZE - 1) {  
        printf("栈已满");  
        return ERROR;   
    }  
    s->top++;  
    s->data[s->top] = e;  
    return OK;  
}   
  
//若栈不空，则删除S的栈顶元素，用e返回其值，并返回OK，否则返回ERROR   
Status pop(SqStack *s, SElemType *e) {  
    if(s->top == -1) {  
        printf("栈为空");  
        return ERROR;  
    }  
    *e = s->data[s->top];  
    s->top--;  
}  
  
//栈元素展示   
void display(SqStack s) {  
    int i =0;  
    for(i; i <= s.top; i++) {  
        printf("%d ", s.data[i]);  
    }  
    printf("\n");  
}  
  
int main() {  
    srand((int)time(NULL)); //用当前的时间作为随机数种子，这样就能保证每次运行时都能取到不同的随机数序列  
  
    SqStack s;  
    s.top = -1;  
    int i = 0;  
      
    for(i; i < random(MAXSIZE); i++) {   //创建一个原始栈并为其赋值   
        s.data[i] = random(MAXSIZE);  
        s.top++;      
    }  
    printf("原始栈为  ：");  
    display(s);  
      
    printf("插入后栈为：");  
    push(&s, random(MAXSIZE));  
    display(s);  
      
    printf("弹出后栈为：");   
    int e;  
    pop(&s,&e);   
    display(s);  
    printf("弹出元素为：%d\n", e);   
}  
output==>
原始栈为  ：5  
插入后栈为：5 0  
弹出后栈为：5  
弹出元素为：0  
请按任意键继续. . .  
</pre>
### 基础知识
堆和栈分为三种：

一、堆栈空间分配区别：
1. 栈（操作系统）：由操作系统自动分配释放 ，存放函数的参数值，局部变量的值等。其操作方式类似于数据结构中的栈；
2. 堆（操作系统）： 一般由程序员分配释放， 若程序员不释放，程序结束时可能由OS回收，分配方式倒是类似于链表。

二、堆栈缓存方式区别：

1. 栈使用的是一级缓存， 他们通常都是被调用时处于存储空间中，调用完毕立即释放；
2. 堆是存放在二级缓存中，生命周期由虚拟机的垃圾回收算法来决定（并不是一旦成为孤儿对象就能被回收）。所以调用这些对象的速度要相对来得低一些。

三、堆栈数据结构区别：

- 堆（数据结构）：堆可以被看成是一棵树，如：堆排序；
- 栈（数据结构）：一种先进后出的数据结构。 


数据结构的栈：限定插入和删除数据元素的操作只能在线性表的一端进行，先进先出了。然而这里又牵扯到线性表的知识。线性表的存储结构分两种：分别是顺序存储结构和链式存储结构。

第一种顺序存储结构的主要特点是：
- （1）结点中只有自身的信息域，没有关联信息域。因此，顺序存储结构的存储密度大、存储空间利用率高。
- （2）通过计算地址直接访问任何数据元素，即可以随机访问。
- （3）插入和删除操作会引起大量元素的移动。
第二种链式存储结构的主要特点是：
- （1）结点除自身的信息域外，还有表示关联信息的指针域。因此，链式存储结构的存储密度小、存储空间利用率低。
- （2）在逻辑上相邻的结点在物理上不必相邻，因此，不可以随机存取，只能顺序存取。 
- （3）插入和删除操作方便灵活，不必移动结点只需修改结点中的指针域即可。 

队列(Queue)也是一种运算受限的线性表，它的运算限制与栈不同，是两头都有限制，插入只能在表的一端进行(只进不出)，而删除只能在表的另一端进行(只出不进)，允许删除的一端称为队尾(rear)，允许插入的一端称为队头 (Front) 。
###C实现堆排序算法
<pre>
/** 
 * author:gubojun 
 * time:2012-12-23 
 * name:堆排序 
 */    
 /* 
  解析：本程序对数列  312,126,272,226,28,165,123,8,12  进行排序 
  首先进入heapsort函数对所有元素建堆 
 
  初始状态------------------------------------------------------------------------ 
        0        1       2       3       4       5       6       7       8       9 
     NULL      312     126     272     226      28     165     123       8      12 
  4)------------------------------------------------------------------------------ 
        0        1       2       3       4       5       6       7       8       9 
      226      312     126     272       8      28     165     123     226      12 
  3)------------------------------------------------------------------------------ 
        0        1       2       3       4       5       6       7       8       9 
      272      312     126     123       8      28     165     272     226      12 
  2)------------------------------------------------------------------------------ 
        0        1       2       3       4       5       6       7       8       9 
      126      312       8     123     126      28     165     272     226      12 
   4    0        1       2       3       4       5       6       7       8       9 
      126      312       8     123      12      28     165     272     226     126 
  1)------------------------------------------------------------------------------ 
        0        1       2       3       4       5       6       7       8       9 
      312        8     312     123      12      28     165     272     226     126 
   2    0        1       2       3       4       5       6       7       8       9 
      312        8      12     123     312      28     165     272     226     126 
   4    0        1       2       3       4       5       6       7       8       9 
      312        8      12     123     126      28     165     272     226     312 
--------------------------------完成 
  */  
#include<stdio.h>  
typedef struct{  
    int r[101];  
    int length;  
}table;  
/************************************************** 
 *  函数功能：筛选算法 
 *  函数参数：结构类型table的指针变量tab 
 *            整型变量k为调整位置 
 *            整型变量m为堆的大小 
 *  函数返回值：空 
 *  文件名：sift.c 函数名：sift() 
 **************************************************/  
void sift(table *tab,int k,int m){  
    int i,j,finished;  
    i=k;j=2*i;  
    tab->r[0]=tab->r[k];  
    finished=0;  
    while((j<=m)&&!finished){  
        if(j<m && tab->r[j+1] < tab->r[j]) j++;//如果看成二叉树，右子树的元素值小于左子树的元素值  
        if(tab->r[0] <= tab->r[j]) finished=1;//如果r[0]是最小的就完成  
        else{  
            tab->r[i]=tab->r[j];  
            i=j;  
            j=j*2;  
        }  
    }  
    tab->r[i]=tab->r[0];  
}  
/************************************************** 
 *  函数功能：堆排序算法 
 *  函数参数：结构类型table的指针变量tab 
 *  函数返回值：空 
 *  文件名：heapsort.c，函数名：heapsort() 
 **************************************************/  
void heapsort(table *tab){  
    int i;  
    for(i=tab->length/2;i>=1;i--)  
        sift(tab,i,tab->length);/*对所有元素建堆,最小的数字在tab.r[1]中*/  
    for(i=tab->length;i>=2;i--){/* i表示当前堆的大小，即等待排序的元素的个数*/  
        tab->r[0]=tab->r[i];  
        tab->r[i]=tab->r[1];  
        tab->r[1]=tab->r[0];/*上述3条语句为将堆中最小元素和最后一个元素交换*/  
        sift(tab,1,i-1);/*对剩下的元素建堆,最小的数字在tab.r[1]中*/  
    }  
}  
int main(){  
    table tab;  
    int num[10]={272,126,312,226,28,165,123,8,12};  
    int i;  
    //初始化  
    for(i=1;i<=9;i++){  
        tab.r[i]=num[i-1];  
    }  
    tab.length=9;  
    //打印初始表  
    for(i=1;i<=9;i++){  
        printf("%d  ",tab.r[i]);  
    }  
    printf("\n");  
    //进行堆排序  
    heapsort(&tab);  
    //打印排序后的表  
    for(i=1;i<=9;i++){  
        printf("%d  ",tab.r[i]);  
    }  
    printf("\n");  
}  
</pre>
### C string.h
字符串常用函数
<pre>
#include <stdio.h>
#include <string.h>

int main(){
	char str1[12] = "HELLO";
	char str2[12] = "jason";
	char str3[23];
	int len;
	
	/* 复制 str1 到 str3 */
	strcpy(str3,str1);
	printf("strcpy(str3,str1):%s\n",str3);
	
	/* 连接 str1 和 str2 */
	strcat(str1,str2);
	printf("strcat(str1,str2)str1:%s\n",str1);
	printf("strcat(str1,str2)str2:%s\n",str2);
	
	len = strlen(str1);
	printf("strlen(str1):%d\n",len);
	return 0;
}
output==>
strcpy(str3,str1):HEELO
strcat(str1,str2)str1:HEELLOjason
strcat(str1,str2)str2:jason
strlen(str1):10
</pre>
结构体
<pre>
#include <stdio.h>
#include <string.h>

struct stuinfo{
	char name[80];
	int score;
}student1;          //注意这里

int main(){
	scanf("%s",&student1.name);
	scanf("%d",&student1.score);
	printf("姓名：%s 成绩：%d\n",student1.name,student1.score); 
	return 0;
}
</pre>
位域

所谓"位域"是把一个字节中的二进位划分为几个不同的区域，并说明每个区域的位数。每个域有一个域名，允许在程序中按域名进行操作。这样就可以把几个不同的对象用一个字节的二进制位域来表示。
典型的实例：

- 用 1 位二进位存放一个开关量时，只有 0 和 1 两种状态。
- 读取外部文件格式——可以读取非标准的文件格式。例如：9 位的整数。
<pre>
#include <stdio.h>
main(){
    struct bs{
        unsigned a:1;
        unsigned b:3;
        unsigned c:4;
    } bit,*pbit;
    bit.a=1;	/* 给位域赋值（应注意赋值不能超过该位域的允许范围） */
    bit.b=7;	/* 给位域赋值（应注意赋值不能超过该位域的允许范围） */
    bit.c=15;	/* 给位域赋值（应注意赋值不能超过该位域的允许范围） */
    printf("%d,%d,%d\n",bit.a,bit.b,bit.c);	/* 以整型量格式输出三个域的内容 */
    pbit=&bit;	/* 把位域变量 bit 的地址送给指针变量 pbit */
    pbit->a=0;	/* 用指针方式给位域 a 重新赋值，赋为 0 */
    pbit->b&=3;	/* 使用了复合的位运算符 "&="，相当于：pbit->b=pbit->b&3，位域 b 中原有值为 7，与 3 作按位与运算的结果为 3（111&011=011，十进制值为 3） */
    pbit->c|=1;	/* 使用了复合位运算符"|="，相当于：pbit->c=pbit->c|1，其结果为 15 */
    printf("%d,%d,%d\n",pbit->a,pbit->b,pbit->c);	/* 用指针方式输出了这三个域的值 */
}
</pre>
共用体

共用体是一种特殊的数据类型，允许您在相同的内存位置存储不同的数据类型。您可以定义一个带有多成员的共用体，但是任何时候只能有一个成员带有值。共用体提供了一种使用相同的内存位置的有效方式。
<pre>
#include <stdio.h>
#include <string.h>
 
union Data
{
   int i;
   float f;
   char  str[20];
};
 
int main( )
{
   union Data data;        

   data.i = 10;
   data.f = 220.5;
   strcpy( data.str, "C Programming");

   printf( "data.i : %d\n", data.i);
   printf( "data.f : %f\n", data.f);
   printf( "data.str : %s\n", data.str);

   return 0;
}
</pre>
typedef

C 语言提供了typedef关键字，您可以使用它来为类型取一个新的名字。您也可以使用 typedef 来为用户自定义的数据类型取一个新的名字。例如，您可以对结构体使用 typedef 来定义一个新的数据类型，然后使用这个新的数据类型来直接定义结构变量，如
<pre>
#include <stdio.h>
#include <string.h>
 
typedef struct Books
{
   char  title[50];
   char  author[50];
   char  subject[100];
   int   book_id;
} Book;
 
int main( )
{
   Book book;
 
   strcpy( book.title, "C Programming");
   strcpy( book.author, "Nuha Ali"); 
   strcpy( book.subject, "C Programming Tutorial");
   book.book_id = 6495407;
 
   printf( "Book title : %s\n", book.title);
   printf( "Book author : %s\n", book.author);
   printf( "Book subject : %s\n", book.subject);
   printf( "Book book_id : %d\n", book.book_id);

   return 0;
}
</pre>
 # define

 #define 是 C 指令，用于为各种数据类型定义别名，与 typedef 类似，但是它们有以下几点不同：

- typedef 仅限于为类型定义符号名称，#define 不仅可以为类型定义别名，也能为数值定义别名，比如您可以定义 1 为 ONE;
- typedef 是由编译器执行解释的，#define 语句是由预编译器进行处理的。
<pre>
#include <stdio.h>
 
#define TRUE  1
#define FALSE 0
 
int main( )
{
   printf( "Value of TRUE : %d\n", TRUE);
   printf( "Value of FALSE : %d\n", FALSE);
   return 0;
}
</pre>
###解决linux vi/vim 编辑无权限
你用一个普通用户编辑了半天，等保存时候猛然发现，你没有权限。。。这是多么蛋疼的情景啊。不过当你试过这个命令就笑了:

w !sudo tee %

查看cpu使用情况：

sar -u 5 720 > cpu.out &

chkconfig

说明：chkconfig命令主要用来更新（启动或停止）和查询系统服务的运行级信息。谨记chkconfig不是立即自动禁止或激活一个服务，它只是简单的改变了符号连接。

语法：chkconfig [--add][--del][--list][系统服务] 或 chkconfig [--level <等级代号>][系统服务][on/off/reset]

linux os 将操作环境分为以下7个等级:

- 0:开机(请不要切换到此等级)
- 1:单人使用者模式的文字界面
- 2:多人使用者模式的文字界面,不具有网络档案系统(NFS)功能
- 3:多人使用者模式的文字界面,具有网络档案系统(NFS)功能
- 4:某些发行版的linux使用此等级进入x windows system
- 5:某些发行版的linux使用此等级进入x windows system
- 6:重新启动

例如：

chkconfig [--level levels] name <on|off|reset>：设置某一服务在指定的运行级是被启动，停止还是重置。例如，要在3，4，5运行级停止nfs服务，则命令如下：
<pre>
chkconfig --level 345 nfs off
</pre>

####/etc/profile与/etc/profile.d/
1. /etc/profile是永久性的环境变量,是全局变量，/etc/profile.d/设置所有用户生效；
2. /etc/profile.d/比/etc/profile好维护，不想要什么变量直接删除/etc/profile.d/下对应的shell脚本即可，不用像/etc/profile需要改动此文件。
###各种变量类型的最大值
以下均为C语言。以int类型为例，int的取值范围在 -32768~32767,无符号类型unsigned int取值范围在0~65535。int类型在不同编译器类型下又有着不同的取值长度，不过基本的计算式为：2^(n-1) n是位数。

因此，在16位编译器中，int占16位（2字节），int最大值为2^(16-1) = 32767；对于32位或64位编译器，int占32位（4字节），int最大值为 2^(32-1) = 2147483647。

### C 文件操作函数
基本函数模型是：
<pre>
FILE *fopen(const char * filename,const char * mode);
</pre>
其中访问模式mode的分下面几种：
<pre>
r   打开已经存在的文件，只读，不能写；不能创建；
r+  打开已经存在的文件，可读可写；不能创建；
w   打开已经存在的文件，从头部写入，不能读；无则创建；
w+  打开已经存在的文件，可读可清空之后写（以前的内容会销毁）；无则创建；
a   打开已经存在的文件，追加写，不能读;无则创建；
a+  打开已经存在的文件，追加写，可读；无则创建。
</pre>
如果处理的是二进制文件，则需使用下面的访问模式来取代上面的访问模式：
<pre>
//b 是 binary
"rb", "wb", "ab", "rb+", "r+b", "wb+", "w+b", "ab+", "a+b"
</pre>
写示例：
<pre>
#include <stdio.h>

main()
{
   FILE *fp;

   fp = fopen("/tmp/test.txt", "w+");
   fprintf(fp, "This is testing for fprintf...\n");    //注意这里语法
   fputs("This is testing for fputs...\n", fp);        //注意这里语法
   fclose(fp);
}
</pre>
读示例：
<pre>
#include <stdio.h>

main(){
	FILE *file;
	char buff[2555];
	
	file =fopen("jason.txt","r");
	fscanf(file, "%s", buff); //读到第一个\n或空格结束       
   	printf("1 content: %s\n", buff);
	fgets(buff,2550,(FILE*)file);                       //注意这里语法
	printf("2 content:%s\n",buff);//最多读取2550个字符 
	fclose(file);
}
</pre>
递归求阶乘
<pre>
#include <stdio.h> 

//阶乘 
double factorial(int i){
	if(i<= 1){
		return 1;
	}
	return i*factorial(i-1);
}

int main(){
	int i = 15;
	printf("%d 的阶乘为%f\n",i,factorial(i));
}
</pre>
不定参数
<pre>
#include <stdio.h>
#include <stdarg.h> 

double average(int num,...){
	va_list valist;
	double sum = 0.0;
	int i;
	va_start(valist,num);
	
	for(i=0;i<num;i++){
		sum += va_arg(valist,int);
	}
	va_end(valist); 
	return sum/num;
}

int main(){
	printf("average of 3,4,5 = %f\n",average(3,3,4,5));
}
</pre>
## 内存操作
主要用到的是 stdlib.h 下的calloc .free . malloc .realloc函数。

- 1.void *calloc(int num, int size);
该函数分配一个带有 function allocates an array of num 个元素的数组，每个元素的大小为 size 字节。
- 2.void free(void *address); 
该函数释放 address 所指向的h内存块。
- 3.void *malloc(int num); 
该函数分配一个 num 字节的数组，并把它们进行初始化。
- 4.void *realloc(void *address, int newsize); 
该函数重新分配内存，把内存扩展到 newsize。
#### malloc动态分配内存
<pre>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int main(){
	char name[100];
	char *description; //动态分配
	
	strcpy(name,"Jason"); 
	
	description = malloc(100 * sizeof(char));//动态分配
	if(description == NULL){
		fprintf(stderr,"Error-unable to allocate requried memory\n");
	}else{
		strcpy(description,"Jason is my nickname,welcome to contactting with me");
	}
	printf("Name = %s\n",name);
	printf("Description:%s\n",description);
}
</pre>
#### 重新调整内存的大小和释放内存
当程序退出时，操作系统会自动释放所有分配给程序的内存，但是，建议您在不需要内存时，都应该调用函数 free() 来释放内存。或者，您可以通过调用函数 realloc() 来增加或减少已分配的内存块的大小。让我们使用 realloc() 和 free() 函数。
<pre>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int main(){
	char name[100];
	char *description;
	
	strcpy(name,"Jason");
	
	description = malloc(30*sizeof(char));
	if(description == NULL){
		fprintf(stderr,"Error-unable to allocate requried memory\n") ;
	}else{
		strcpy(description,"Jason I am");
	}
	
	description =realloc(description,100*sizeof(char));
	if(description == NULL){
		fprintf(stderr,"Error-unable to allocate required memory\n");
	}else{
		strcat(description,"She is in class 10th\n");
	}
	printf("Name = %s\n",name);
	printf("Description:%s\n",description);
	
	free(description);
}
</pre>
C访问结构体成员两种方法

- 第一种比较简单就是使用结构体运算符，点号就行了，也不必分配内存给它；
- 第二种也就是下面演示的比较麻烦或者说有难度，它使用的是结构体指针运算符，需要手动分配内存，同时注意不要导致内存泄漏。
<pre>
#include <stdio.h>  
#include <malloc.h>  
#include <string.h>  
typedef struct changeable{  
        int iCnt;  
        char pc[0];  
}schangeable;  
  
main(){  
        printf("size of struct changeable : %d\n",sizeof(schangeable));  
  
        schangeable *pchangeable = (schangeable *)malloc(sizeof(schangeable) + 10*sizeof(char));  
        printf("size of pchangeable : %d\n",sizeof(pchangeable));  
  
        schangeable *pchangeable2 = (schangeable *)malloc(sizeof(schangeable) + 20*sizeof(char));  
        pchangeable2->iCnt = 20;  
        printf("pchangeable2->iCnt : %d\n",pchangeable2->iCnt);  
        strncpy(pchangeable2->pc,"hello world",11);  
        printf("%s\n",pchangeable2->pc);  
        printf("size of pchangeable2 : %d\n",sizeof(pchangeable2));  
}  
</pre>
自己写一个
<pre>
#include <stdio.h>
#include <malloc.h> 

typedef struct{
	int a;
}Djason;

main(){
	//手动给结构体分配内存,这个是用指针的形式，涉及到底层
	Djason *jason = (Djason *)malloc(sizeof(Djason)+20*sizeof(char));
	jason->a = 34;
	printf("jason->a`s value is %d",jason->a);
}
output==>
jason->a`s value is 34
</pre>
