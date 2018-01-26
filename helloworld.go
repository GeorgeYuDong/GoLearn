package main
import (
	"fmt"
	"unsafe"
	"errors"
	"math"
) // 我们需要使用fmt包中的Println()函数

type Phone interface {
	call()
}

type NokiaPhone struct{

}

type IPhone struct {

}

func (iPhone IPhone) call(){
	fmt.Println("iphone call")
}

func (nokiaPhone NokiaPhone) call(){
	fmt.Println("nokia call")
}


type Circle struct{
	radius float64
}
//递归函数
func Factorial(n uint64)(result uint64) {
	if (n > 0) {
		result = n * Factorial(n - 1)
		return result
	}
	return 1
}


//切片
func ModifyArray(array [5]int){
	array[0] = 100
}

//方法
func (c Circle) getArea() float64{
	return 3.14*c.radius*c.radius
}


//第二个func是匿名函数，返回的是匿名函数
func getSequence() func() int{
	ba := 0
	return func() int { //此为闭包
		ba+=1
		return ba
	}
}

var vname4,vname2 int = 10,20

//:=声明不能用于全局变量,全局变量的初始化形式进行统一,也就是必须出现var关键字
//vname5 := "string"
//err的类型为error,需要导入包errors
func add(a int, b int)(ret int, err error){

	if a < 0 || b < 0 {
		err = errors.New("should be non-negative numbers")
		return
	}

	return a + b,nil
}

func swap(x string,y string )(string, string){
	return y,x
}

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
	vname1, vname5,vname3 := 1,2,3
	fmt.Println(vname1,vname5,vname3)

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

	var jj int = 10
	var mm uint = 2 //移位的数目必须是无符号数
	jj <<= mm
	fmt.Println("jj<<mm",jj)
	jj >>= mm
	fmt.Println("jj>>mm",jj)
	//括号提升运算符优先级
	ptr := &mm
	fmt.Println(*ptr)

	//select语句，类似case,不同之处，在于选择一个可运行的case，没有则阻塞
	//for循环语句
	a = 1
	for a < 100{
		a++
	}
	fmt.Println("a循环后",a)

	var nn int = 8
	var ii int
	//返回两个参数，调用的时候，需要输出两个，不使用的用_代替
	//函数名首字母小写，只能在内部访问，大写，则可被外部文件访问
	ii, _ = add(jj,nn)
	fmt.Println("a+b",ii)
	//两个数交换,支持并行赋值
	jj,nn = nn,jj
	fmt.Println("jj,nn",jj,nn)

	s_x := "abc"
	s_y := "kxyz"
	var s_z string
	var s_zz string
	s_z,s_zz = swap(s_x, s_y)
	fmt.Println("swap(s_x,s_y)",s_z,s_zz)
	//函数传递，默认情况使用值传递，不影响实参
	//全局变量和局部变量名称可以相同，但优先使用局部变量
	var vname2 int = 30
	fmt.Println("vname2:",vname2)

	//指针类型的初始化默认值为nil
	var pptr *int
	fmt.Println("pptr:",pptr)
	//数组声明,关键字var,首地址名称，个数，元素类型
//	var bal [5] int
	//数组的初始化有以下方式
	 bal :=[...]int{1,2,3,4,5}
//	var bal = [...]int{1,2,3,4,5}

	//类型写在最后面，前面强调首地址，个数，类型
	//	bal := [...]int
	fmt.Println("数组元素:",bal[0],bal[1],bal[2])

 	//go语言里有个select，涉及到channel，到后面学习
	//switch里面case var1，var1可以是任何类型的变量,case后面不需要加break
	//switch后面可以同时测试多个执行项, case var1,var2,var3
	//switch还涉及到了接口，后面再测

	//函数本身可以作为变量使用,使用的时候，变量里面传参
	getSquareRoot := func(x float64) float64{
		return math.Sqrt(x)
	}
	fmt.Println("函数作为值：",getSquareRoot(9))

	//闭包是一个内联函数,回去继续学习
	nextNumber := getSequence()

	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())

	nextNumber1 := getSequence()
	fmt.Println(nextNumber1())
	fmt.Println(nextNumber1())

    //函数方法,func 接受者，方法名，返回类型
	var c1 Circle
	c1.radius = 10.00
	fmt.Println("Area of Circle(c1) =",c1.getArea())

	//数组:var,首地址名，个数，类型
	//bal := []float,首地址，个数，类型,这种思维方式要慢慢习惯
	//不传地址，占用内存。闭包，占用内存.
	//三行四列，先写列，再写行,注意这种差别
	//访问的时候，先访问行，再访问列,行在外层，列在内层,里面的是最后访问的

	var iptr *int
	var flp *float64
	//Printf格式化输出,%x,十六进制
	fmt.Printf("ip的地址 =%x\n",iptr)
	iptr = &a
	flp = &c1.radius
	fmt.Println("a的地址 =",iptr)
	fmt.Println("radius的地址 =",flp)
	//if (iptr == nil),指针与空值比较
	iiptr := &a //也可以这样声明，与整数类型声明一样，自动推导
	fmt.Printf("iiptr的地址 =%x\n",iiptr)

	var abcd [3] int //指针默认值都是nil,指针数组
	fmt.Println("abcd的值 =",abcd[0],abcd[1])
	abcd = [...]int{4,5,6} //赋值的时候，必须要有省略号
	fmt.Println("abcd1234的值 =",abcd[0],abcd[1])

	//初始化的时候省略号可以拿掉,数组个数可以不输,把报错情况考虑进去
	bbc := [...]int{1,2,3} //数组强调首地址名，也是一个变量，定义的时候要注意
	fmt.Println("bbc的值 =",bbc[0],bbc[1])
	var ccd = []int{3,5,6} //越界报index out of range
	fmt.Println("ccd的值 =",ccd[0],ccd[1])


 	//指向指针的指针，同C语言，两个**
 	//指针可作为参数传递，同C语言

 	//结构体定义,type设定了结构体名称，struct是结构体关键字
 	//type类似于C语言typedef
	type Books struct {
		title string
		book_id int
	}

	var Book1 Books
	var Book2 Books

	Book1.title = "Go languange"
	Book1.book_id = 1
	Book2.title = "C languange"
	Book2.book_id = 2

	fmt.Println("Book1.title:",Book1.title)
	fmt.Println("Book1.book_id:",Book1.book_id)
	fmt.Println("Book2.title:",Book2.title)
	fmt.Println("Book2.book_id:",Book2.book_id)


	//结构体指针可以返回结构体,结构体名称就是结构体类型
	//访问结构体元素方法：.操作符
	var pBook *Books //结构体指针 ,type类似于typedef
	pBook = &Book1
	fmt.Println("pBook.title:",pBook.title)

	fmt.Printf("ii的值:")
	for ii = 0; ii < 10 ; ii++{
		fmt.Printf("%d ",ii)
	}
	fmt.Printf("\n")

	//bal :=[...]int{1,2,3,4,5}
	//	var bal [5] int
	//Slice切片，对数组的抽象，动态数组，不限定数组数量
	slice1 := [] int {1,2,3}
	fmt.Println("slice:", slice1[0])

	slice1 =[] int {5,6,7}
	//[]标识为切片
//	s := make([]int, 3, 3)
//	fmt.Println("s:", s[0],s[1])
	//初始化切片是数组bal的引用
	s := bal[:]
	fmt.Println("s:", s[0],s[1])
	//直接打印s,bal，打印出来的整个数组
	fmt.Println("s", s)
	fmt.Println("bal", bal)
	//3为初始长度，5为最大长度(容量)
	var numbers = make([]int, 3, 5)
	numbers = []int{1,2,4,7,8,9,10,11}
//	fmt.Println("numbers:",numbers[6])
//	fmt.Println("numbers len:",len(numbers))
//	fmt.Println("numbers cap:",cap(numbers))
	//含索引1，不包含3
	fmt.Println("numbers:",numbers[1:3])
	//0到索引2，不包含3
	fmt.Println("numbers:",numbers[:3])
	//开始值，是从索引1到最后
	fmt.Println("numbers:",numbers[1:])

	var num []int
	if (num == nil){
		fmt.Println("num", num)
	}

	numbers = append(numbers, 10,20,30)
	fmt.Println("numbers:",numbers[1:])
/*
	fmt.Println("before modify,numbers:",numbers)
	ModifyArray(numbers)
	fmt.Println("modified,numbers:",numbers)
*/
	//两者一样，数组是值传递,函数内改变数组元素值不会生效，切片可以
	fmt.Println("before modify,bal:",bal)
	ModifyArray(bal)
	fmt.Println("modified,bal:",bal)
	//切片，指向数组的指针,传的是指针，所以可以改变数组内元素值
	numbers1 := make([]int,len(numbers),(cap(numbers)*2))
	copy(numbers1,numbers)
	fmt.Println("cap(numbers1):",cap(numbers1))
	fmt.Println("cap(numbers):",cap(numbers))

	var ss = make([]int, 5, 5)
	ss = []int{1,2,3,4,5}
	fmt.Println("ss:",ss[1:])

	//range的用法
	nums := []int{1,2,3,4}
	sum := 0
	for _, num := range nums{
		sum += num
	} //_为索引，num为索引对应的值
	fmt.Println("sum:",sum)

	for i, num := range nums{
		if (num == 3){
			fmt.Println("index:",i)
		}
	}

	kvs := map[string]string{"a":"apple","b":"banana"}
	for k, v := range kvs {
	//	fmt.Printf("%s->%s\n",k,v)
			fmt.Println(k,v)
	}
	//i为索引，c为索引对应的值
	for xx,c := range "ab" {
		fmt.Println(xx,c)
	}
	//map的两种初始化方式，第一种:
	//map集合,第一个string是key_type,第二个string是value_type
	var countryCapitalMap map[string]string
	fmt.Println(countryCapitalMap)
	countryCapitalMap = make(map[string]string)

	//第二种
	kvvs := map[string]string{"a":"apple","b":"banana"}
	fmt.Println(kvvs)

	countryCapitalMap["china"] = "beijing"
	countryCapitalMap["korea"] = "hancheng"
	//countryCapitalMap是一个键值对
	fmt.Println(countryCapitalMap)

	//for后面可以只有索引
	for country := range countryCapitalMap {
		fmt.Println("capital of",country,"is",countryCapitalMap[country])
	}

	for country,capital := range countryCapitalMap {
		fmt.Println(country,",",capital)
	}
	//第2个参数代表是否存在
	capital, ok := countryCapitalMap["china"]
	fmt.Println(ok,capital)

	//delete函数删除集合元素
	delete(countryCapitalMap,"korea")
	fmt.Println(countryCapitalMap)

	//递归完成,确定到哪边递归完成,避免无限循环
	factorial := 3
	fmt.Println(Factorial(uint64(factorial)))

	mean := float32(17) /float32(5)
	fmt.Printf("%f\n",mean)

	//非倾入式接口
	var phone Phone
	phone = new(IPhone)
	phone.call()

	phone = new(NokiaPhone)
	phone.call()



}