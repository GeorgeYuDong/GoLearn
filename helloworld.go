package main
import (
	"fmt"
	"unsafe"
) // 我们需要使用fmt包中的Println()函数

var vname4,vname2 int

func main() {
	fmt.Println("hello")
	fmt.Println("123")

	fmt.Println("abc"); fmt.Println("456")
	//标识符不能以数字开头，可以是字母或下划线
	//变量的声明必须使用空格隔开
	var age int = 10;
	fmt.Println(age)
	//类型可以不写，go编译器可以自动推导
	var number int = 5
	//can't user ... as type int in assignment
	var quote = "life is simple,i use python"
	fmt.Println(number, quote)
	//var 声明变量，如果不赋初值，必须制定类型，如果赋初值，则可自动推导类型，不是必须的
	var a int
	fmt.Println(a)
	//:=声明变量，变量名必须是新的，不能是已经声明过的,此时var关键字必须拿掉
	b := 20
	b = 10 //可以，因为已经声明过，可以对他赋值
	fmt.Println(b)
	//:=声明方式只能用在函数体内，不能用于全局变量

	//函数前面必须加包名，否则undefined func
	//可以多个变量同时声明
	vname1, vname2,vname3 := 1,2,3
	fmt.Println(vname1,vname2,vname3)

	var (
		integer int
		bo bool //bool值默认是false
	)
	fmt.Println(integer, bo)
	//可以并行赋值，可用于多个函数返回值
	integer, bo = 5, true
	fmt.Println(integer, bo)
	//全局变量允许声明但不使用，函数内变量则例外，会报错
	//如果函数内变量声明后未被使用，也会报c declared and not used
//	c := 20
	abc := 1;
	abc = 2
	fmt.Println(abc)

	//定义常量
	const LENGTH = 10
	const WIDTH = 5
	area := LENGTH * WIDTH
	fmt.Println("面积为:", area)

	//常量定义枚举,常量表达式必须是内置函数，否则编译不过
	const (
		Unknow = "abcdef"
		Female = len(Unknow)
		//sizeof 字符串，16位
		//字符串类型在 go 里是个结构, 包含指向底层数组的指针和长度,
		//这两部分每部分都是 8 个字节，所以字符串类型大小为 16 个字节
		Male = unsafe.Sizeof(Unknow)
	)

	fmt.Println(Female,Male)
	//特殊常量,第一次出现置0，下面再出现加1
	const (
		aa = iota  //itoa值0
		ab         //1
		ac		   	//2
		ad = "hello" //3
		ae 			//4
		am = iota //此处为5
		an 			//6
	)
	fmt.Println(aa,ab,ac,ad,ae,am,an)
	//不写则按照原模式顺序拷贝
	const (
		i = 1 << iota    //1*2^0
		j = 3 << iota    //3*2^1
		k               //3*2^2
		l               //3*2^3
	)
	fmt.Println(i,j,k,l)

 	a = 10
 	b = 100

 	a += 10
	fmt.Println("a+10=",a)
 	b /= a
	fmt.Println("b/a=",b)
	b *= a
	fmt.Println("b*a", b)
	b %= a
	fmt.Println("b%a",b)
 	b |= a
	fmt.Println("b|a",b)
 	b &= a
	fmt.Println("b&a",b)
 	b ^= a
	fmt.Println("b^a",b)









}