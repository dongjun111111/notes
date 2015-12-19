panic 抛出异常 <br>
recover 捕获异常 捕获后会返回外层函数继续执行 <br>
defer 函数退出之前执行

###cap()函数
cap()函数返回的是数组切片分配的空间大小。
<pre>
package main
import "fmt"
func main(){
	myslice := make([]int,5,10)
	fmt.Println("len(myslice):",len(myslice))
	fmt.Println("cap(myslice):",cap(myslice))
}

Output==>
len(myslice):5
cap(myslice):10
</pre>
###unsafe
由于Go语言不能和C语言一样直接进行指针运算，所以需要引入unsafe包，通过它进行运算.unsafe.Pointer其实就是类似C的void *，在golang中是用于各种指针相互转换的桥梁。
在Go语言中，指针的本质是什么呢？是unsafe.Pointer和uintptr。