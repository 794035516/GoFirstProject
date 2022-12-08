package main

//一个独立可执行的go语言程序必须要有package main的声明

import (
	"errors"
	"fmt"
)

/*
type 用于设定结构体的名称 student

struct 用于定义一个新的数据类型
*/
type student struct {
	age  int
	name string
}

type File struct {
	fd int
}

/*
把参数移动到函数名的开头，这个函数就成了File类型独有的方法，而不是File对象方法
*/
func (f *File) CloseFile() error {
	if f == nil {
		return errors.New("file null")
	}
	return nil
}

func (f *File) ReadFile(offset int64, data []byte) error {
	if f == nil {
		return errors.New("file null")
	}
	fmt.Println(offset, " -- ", data)
	return nil
}

// 接口类型转换为具体类型（也就是“反射”, reflect），具体里诶慈宁宫向接口类型赋值，都涉及到了类型转换
// 下面是空接口的数据结构：
//
//		type interfaceName interface {}
//	没有任何方法的接口就是空接口，实际上每个类型都实现了空接口，所以空接口类型可以接受任何类型的数据
type Eface interface {
	speak()
}

type Cat struct {
	class string
}

func (animal *Cat) Speak(q interface{}) {
	fmt.Println("cat miao miao miao...", animal.class)
	temp, ok := q.(string)
	if ok {
		fmt.Println("类型转换成功", temp)
	} else {
		fmt.Println("类型转换失败", temp)
	}
}

func Print(arr ...interface{}) {
	fmt.Println(arr)
}

func showmpType(q interface{}) {
	fmt.Println("type:%T, value:%v", q, q)
}

// test

// 为了创建一个二进制可执行文件，我们需要让我们的程序成为 main 包的一部分，而且必须要有一个作为执行入口点的 main 函数
func main() {
	fmt.Println("hello world!!!")
	//声明变量 var identify typename
	var name string
	var num int = 10
	//上述语句等同于 name := 10
	//多变量声明
	var num1, num2 = 1, 2
	str1, str2 := 11, "str2"
	fmt.Println(num1, num2, str1, str2)
	//全局变量声明 var ( vname1 v_type1
	//	vname2 v_type2)

	var (
		a int
		b bool
	)
	fmt.Println(a, b)
	name = "lijx"
	var d = true
	fmt.Println(d, name)
	//指针的使用
	var ip *int
	ip = &num
	fmt.Println(ip, *ip, &ip)
	ip = nil
	fmt.Println(ip)
	//数组的使用
	var fruit = [5]string{"peach"}
	var integers [10]int
	for i := 0; i < len(integers); i++ {
		integers[i] = i
	}
	fmt.Println(fruit) //go中数组名仅仅表示整个数组，是一个完整的值，一个数组变量即表示整个数组
	fmt.Println(integers)

	//数组指针
	var arr1 = []int{11, 12, 13}
	var arrB = &arr1
	for i, v := range *arrB {
		fmt.Println(i, v)
	}
	for _, v := range arr1 {
		fmt.Println(v)
	}

	//结构体的使用
	var lijx student = student{age: 21, name: "lijx"}
	var lijx2 student
	lijx2.age = 33
	lijx2.name = "lijx2"
	fmt.Println(lijx, len(lijx.name))
	fmt.Println(lijx2, len(lijx2.name))
	lijx.age = 18
	fmt.Println(lijx)

	//数组切片的使用
	arr := []int{1, 2, 3, 4, 5}
	arr = append(arr[:0], arr[1:]...)
	fmt.Println(arr)
	arr = []int{1, 2, 3, 4, 5}
	arr = append(arr[:0], arr[3:]...)
	fmt.Println(arr)
	fmt.Println("------------------")
	arr = []int{1, 2, 3, 4, 5}
	arr = arr[:copy(arr, arr[1:])]
	fmt.Println(arr)
	arr = []int{1, 2, 3, 4, 5}
	arr = arr[:copy(arr, arr[3:])]
	fmt.Println(arr)
	arr = []int{1, 2, 3, 4, 5}
	slice := arr[0:3:len(arr)] //第三个参数过大会被识别成木马程序
	fmt.Println(slice)

	//函数传参
	Print(arr)
	//Print(arr...)

	//defer语句延迟执行的其实是一个匿名函数，因为这个匿名函数捕获了外部函数的局部变量，这种函数一般叫外包。
	//闭包对捕获的外部变量并不是传值方式访问，而是以引用的方式访问
	//但是这种方式往往会带来一些问题，修复方法为在每一轮迭代中都为defer函数提供一个独有的变量
	for i := 0; i < 3; i++ {
		fmt.Println(i)
		//通过函数传入参数
		//defer 语句会马上对调用参数传值
		//不在捕获，而是直接传值
		//defer func(i int) { fmt.Println(i) }(i)
	}
	fmt.Println("------")
	for i := 0; i < 3; i++ {
		//j := i
		//定义一个循环体内局部变量
		//defer func() { fmt.Println(i, j) }()
	}
	showmpType(nil)
	var cat Cat = Cat{class: " I am a cat..."}
	cat.Speak(Cat{})
}
