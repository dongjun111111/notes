Js的面向对象
------------------
### 起因
其实就是因为Js的基础太薄弱了，看了N遍的原型还有Js对象的问题，还是没有搞懂，终于翻到这篇博文了，[Js中的self和this小结](http://www.cnblogs.com/fullhouse/archive/2012/03/20/2407647.html)

直接上段代码：

	var Class = {
		create: function() {
			return function() {
				this.initialize.apply(this, arguments);
			}
		}
	}

Class类的使用如下：

	var A = Class.create();
	A.prototype = {
		initialize.function(v) {
			this.value = v;
		}
		showValue: function() {
			alert(this.value);
		}
	}

	var a = new A('hello-world');
	a.showValue(); //结果弹出对话框显示"hello-world"

对于上面的代码，这就有了疑问：

- initialize是什么东西呀？
- apply方法是干啥子的么？
- arguments变量又是搞墨子的么？
- 为什么new A之后会立即执行initialize方法？

### 寻找答案 - Js的面向对象

#### initialize
只不过是个变量，代表一个方法，用途是类的构造函数。

其具体功能靠js的面向对象支持，那么js的面向对象是什么样子的那？和java 的有什么相同与不同？

看代码永远是最直观的：

	var ClassName = function(v) {
		this.value = v;
		this.getValue = function() {
			return this.value;
		}
		this.setValue = function() {
			this.value = v;
		}
	}

那么**Js中的函数和类到底有什么不同**呢？

其实是**一样的**，是的，我被震惊了，但它们确实一样~~

拿上面代码来讲，有两种情况：

- ClassName就是一个函数，出现在new之后就作为一个**构造函数**来`构造对象`。如：

	var obj1 = new ClassName("a"); // 得到一个对象

其中obj1就是执行ClassName构造函数后得到的对象，而在ClassName函数中的this指的就是new之后构造出来的对象，所以objectName1会后一个属性和两个方法。可以通过这样来调用他们：

	obj1.setValue(''hello''); 
	alert(obj1.getValue());//对话框hello 
	alert(obj1.value) ;//对话框hello

- 如果直接调用，如：

	var obj2 = ClassName("b"); //同样得到一个对象

但是，这里的obj2，显然是**函数ClassName()执行后的返回值**，这里ClassName只作为了一个普通的函数（虽然首字母大写了）。但是在之前写的ClassName中并没有返回值，所以obj2会是undifinded的。那么"b"赋给谁了呢？在这并没有产生一个对象，而只是单纯的执行这个方法，所以这个"b"赋值给了调用这个方法的对象window，证据如下，在上面obj2的基础上：

	alert(window.value); //得到对话框，显示"b"


