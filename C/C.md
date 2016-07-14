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
###C实现栈的线性表的链表数据结构
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
###基础知识
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
###C string.h
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
 #define

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