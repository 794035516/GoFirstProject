package main

//一个独立可执行的go语言程序必须要有package main的声明

import "fmt"

/*
	type 用于设定结构体的名称 student

struct 用于定义一个新的数据类型
*/
type student struct {
	age  int
	name string
}

// g
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
	var ip *int
	ip = &num
	fmt.Println(ip, *ip, &ip)
	ip = nil
	fmt.Println(ip)
	var fruit = [5]string{"peach"}
	var integers [10]int
	for i := 0; i < len(integers); i++ {
		integers[i] = i
	}
	fmt.Println(fruit)
	fmt.Println(integers)
	var lijx student = student{age: 21, name: "lijx"}
	fmt.Println(lijx)
	lijx.age = 18
	fmt.Println(lijx)

}
