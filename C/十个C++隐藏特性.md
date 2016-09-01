title:  十个C++隐藏特性
description:
time: 2015/04/11 14:08
category: Languages
++++++++

## 1. []运算符的真相

因为`a[8]`是`*(a+8)`的同义表达，`*(a+8)`又等于`*(8+a)`，所以`a[8]`可以写成……`8[a]`。

```
int a[3] = {0};
2[a] = 1;
2[a][a] = 2;
// a = {0,2,1}
```

## 2. 三元运算符的返回值

三元运算:前后的表达式并不要求类型相同，只需要是同一类型种类(category)，最后的返回值是二者中最为通用的类型。

```
void foo (int) {}
void foo (double) {}
struct X {
  X (double d = 0.0) {}
};
void foo (X) {}

int main(void) {
  int i = 1;
  foo(i ? 0 : 0.0); // 调用foo(double)
  X x;
  foo(i ? 0.0 : x);  // 调用foo(X)
}
```

## 3. 函数里面可以直接写URL

```
void foo() {
    http://disksing.com/
    ...
}
```

可以编译过不报错的。好吧，这个其实比较囧。

## 4. 不用memset初始化结构体

```
struct A {
   int x;
   int y;
};
A a = {0};   // 或者A a = {};
```

## 5. 不用memset初始化成员数组

```
class A
{
public:
     A():m_a(),m_c(){}
     
     int m_a[10];
     char m_c[10];
};
```

## 6. 规避前置声明

```
struct global
{
    void main()
    {
        a = 1;
        b();
    }
    int a;
    void b() {}
} singleton;
```

## 7. 三元运算符

大部分书上讲三元运算符都是这么用的：
```
x = (y < 0) ? -1 : 1;
```
其实左值也可以是?:表达式：
```
(a < b ? a : b) = 10;
```
等价于：
```
if (a < b)
    a = 10;
else
    b = 10;
```
可以左值右值都是三元运算符：
```
(a < b ? a : b) = (y < 0) ? -1 : 1;
```
还可以用来调用函数的：
```
void even(int x) {printf("%d is even.", x);}
void odd(int x) {printf("%d is odd.", x);}
(x%2==0 ? even : odd)(x);
```

## 8. namespace可以用当变量用

```
namespace fs = boost::filesystem;
fs::path myPath( strPath, fs::native );
```

在超大型的项目中，或者需要在不同的命名空间切换时比较有用。

## 9. for循环中可以定义struct或class

```
for(struct {int x; int y;} point = {0,0}; ...; ...) {
    ...
}
```

## 10. 函数默认参数的值可以修改

这是因为C++支持用静态变量当参数的默认值(不一定要是const的)

```
static int d = 1;

int foo(int x = d)
{
     return x * 2;
}

int main() {
     int x = foo(); // x = 2
     d = 2;
     int y = foo(); // y = 4
}
```

