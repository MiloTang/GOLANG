//包声明,程序运行的入口是包 main
package main

//引入包
import (
	"errors"
	"fmt"
	/*也可写成. "fmt"省略前缀的包名，调用的时候fmt.Println("hello world")可以写成Println("hello world")
	P "fmt" 调用时P.Println("hello world")
	_ "fmt" _操作其实是引入该包，而不直接使用包里面的函数，而是调用了该包里面的init函数。
	*/
	"math"
	"reflect"
	"runtime"
	"time"
	/*上面这个fmt是Go语言的标准库，其实是去 GOROOT 环境变量指定目录下去加载该
	模块，当然Go的import还支持如下两种方式来加载自己写的模块：*/
	//相对路径import "./model"当前文件同一目录的model目录
	//绝对路径import "model" 加载gopath/src/model模块
)

//函数外定义的变量称为全局变量

//var identifier type
var var1 string

/*
var 变量名1,变量名2... = 值1, 值2...
不需要显示声明类型，自动推断
*/
var var2, var3 = "study", 2

//因式分解关键字的写法一般用于声明全局变量
var (

	//布尔型的值只可以是常量 true 或者 false。
	var5 bool

	//整型 int 和浮点型 float，Go 语言支持整型和浮点型数字，并且原生支持复数，其中位的运算采用补码。
	var6  uint8
	var7  uint16
	var8  uint32
	var9  uint64
	var10 int8
	var11 int16
	var12 int32
	var13 int64
	var14 float32
	var15 float64
	var16 complex64
	var17 complex128
	var18 byte
	var19 rune
	var20 uint
	var21 int
	var22 uintptr

	//字符串类型
	var23 string

	/*派生类型*/
	//数组
	var24 [10][6]string
	var25 = [2]string{"milo", "tang"}
	//语言会根据元素的个数来设置数组的大小
	var26 = [...]float32{1002.0, 22.0, 33.5, 87.0, 150.0}

	//指针
	var33 *int
	var34 *string

	//结构化类型(struct)
	amen     Persons
	aboby    = Persons{"Joy", 11, "男", 1.30}
	vertu    = Mobile{"Vertu", 100000.00, "电信"}
	huaweip8 = HuaWei{Mobile{"化为P8", 3000.00, "移动"}, "一个年轻人的专享"}

	//channel 可以是 带缓冲的。为 make 提供第二个参数作为缓冲长度来初始化一个缓冲
	ch1 = make(chan int)
	ch2 = make(chan int, 2)
	//函数类型
	func1 FuncT

	/*切片类型,切片是对数组的抽象var identifier []type
	make([]T, length, capacity)
	*/
	slice1 []int
	slice2 = make([]int, 2, 3)
	slice3 = make([]string, 2)
	slice4 = []int{1, 2, 3}

	//接口类型（interface）空interface(interface{})不包含任何的method，正因为如此，所有的类型都实现了空interface
	phone   Phone
	emptyin interface{}

	/*Map 类型M是一种无序的键值对的集合。Map 最重要的一点是通过 key 来快速检索数据，
	key 类似于索引，指向数据的值,var map_variable map[key_data_type]value_data_type
	map_variable = make(map[key_data_type]value_data_type)*/
	map1 = map[string]string{"苏州": "南京", "山东": "济南", "四川": "成都", "西藏": "拉萨", "新疆": "乌鲁木齐", "陕西": "西安", "山西": "太原"}
	map2 = make(map[string]string)
)

//结构化类型(struct)
type Persons struct {
	name   string
	age    int
	sex    string
	height float32
}
type HuaWei struct {
	Mobile
	characteristic string
}
type IPhone struct {
	Mobile
	characteristic string
}
type Mobile struct {
	version string
	price   float32
	nettype string
}
type Human struct {
	name   string
	age    int
	weight int
}
type Student struct {
	Human      // 匿名字段，那么默认Student就包含了Human的所有字段
	speciality string
}
type Employee struct {
	Human   //匿名字段
	company string
}
type AnimalData struct {
	typename string
	age      int
	sex      string
}
type DogData struct {
	AnimalData
	dogtye  string
	dogname string
}
type CowData struct {
	AnimalData
	cowcolor string
}

//函数类型
type FuncT func(int, int)

//接口类型（interface）
type Phone interface {
	call()
	net()
}
type Animal interface {
	Move()
	Sound()
	Eat()
}

//常量const identifier [type] = value
const str1 string = "str1"
const str2 = "str2"

//枚举
const (
	Monday    = "星期一"
	Tuesday   = "星期二"
	Wednesday = "星期三"
	Thursday  = "星期四"
	Friday    = "星期五"
	Saturday  = "星期六"
	Sunday    = "星期日"
)

/*
iota，特殊常量，可以认为是一个可以被编译器修改的常量。第一个 iota 等于 0，
每当 iota 在新的一行被使用时，它的值都会自动加 1
*/
const (
	var27 = 1 << iota
	var28 = 2 << iota
	var29
	var30
	var31
	var32
)

//函数定义中的变量称为形式参数
func Fslice(sl []int) {
	if sl == nil {
		fmt.Printf("len=%d cap=%d slice=%v\n", len(sl), cap(sl), sl)
		fmt.Printf("切片是空的\n")
	} else {
		fmt.Printf("len=%d cap=%d slice=%v\n", len(sl), cap(sl), sl)
		sl = append(sl, 100)
		fmt.Printf("len=%d cap=%d slice=%v\n", len(sl), cap(sl), sl)
		sl = append(sl, 101, 102)
		fmt.Printf("len=%d cap=%d slice=%v\n", len(sl), cap(sl), sl)
		sl = make([]int, len(sl), (cap(sl))*2)
		fmt.Printf("len=%d cap=%d slice=%v\n", len(sl), cap(sl), sl)
		s := []int{200, 201, 202}
		copy(sl, s)
		fmt.Printf("len=%d cap=%d slice=%v\n", len(sl), cap(sl), sl)
		fmt.Printf("len=%d cap=%d slice=%v\n", len(sl), cap(sl), sl[0:3])
	}

}

func Fstruct(p Persons) {
	fmt.Printf("Person name : %s\t", p.name)
	fmt.Printf("Person age : %d\t", p.age)
	fmt.Printf("Person sex : %s\t", p.sex)
	fmt.Printf("Person height : %.2f\n", p.height)
}

func Fpointer() {
	/*
		var34指针的地址
		*var34指针指向的值
		&var34指针的存储地址
	*/
	var govar = 100
	var point *int
	println("指针相关", var23, var33, &var1, var34, *var34, &var34, govar, point)
	point = &govar
	*point = 200
	fmt.Println(*point, point, govar)
}

func FRange() {
	/* range 关键字用于for循环中迭代数组(array)、切片(slice)、链表(channel)或集合(map)的元素。
	   在数组和切片中它返回元素的索引值，在集合中返回 key-value 对的 key 值。*/
	nums := []int{10, 20, 30, 40, 50, 60}
	//_抛弃不需要的值
	for index, _ := range nums {
		println("index==>", index)
	}
	for _, value := range nums {
		println("value==>", value)
	}
	for index, value := range nums {
		println(index, "==>", value)
	}
}

func FMap() {
	map2["湖北"] = "武汉"
	map2["湖南"] = "长沙"
	map2["云南"] = "昆明"
	map2["贵州"] = "贵阳"
	map2["广东"] = "广州"

	for province, city := range map2 {
		println(province, "==>", city)
	}
	for province, city := range map1 {
		println(province, "==>", city)
	}
	delete(map1, "西藏")
	for province, city := range map1 {
		println(province, "==>", city)
	}
	c, ok := map2["贵州"]
	fmt.Println("The :", c, "City?", ok)

}

func functype(int, int) {
	fmt.Println("this is a func type")
}

func (f FuncT) FuncTest() {
	fmt.Println("this is a func type test")
}

func Ffunc() {
	func1 = FuncT(functype)
	func1(1, 2)
	func1.FuncTest()
}

func Frecursion(x int) (result int) {
	if x == 0 {
		result = 1
	} else {
		result = x * Frecursion(x-1)
	}
	return
}

func (huawei HuaWei) call() {
	fmt.Println("hello, huawei!")
	fmt.Println("版本:", huawei.version)
	fmt.Printf("价格:%0.2f\n", huawei.price)
	fmt.Println("其他特点:", huawei.characteristic)
}

func (huawei HuaWei) net() {
	fmt.Println("网络类型", huawei.nettype)
}

func (iPhone IPhone) call() {
	fmt.Println("hello, Iphone")
	fmt.Println("版本:", iPhone.version)
	fmt.Printf("价格:%0.2f\n", iPhone.price)
	fmt.Println("其他特点:", iPhone.characteristic)
}
func (iPhone IPhone) net() {
	fmt.Println("网络类型:", iPhone.nettype)
}
func (mb Mobile) call() {
	fmt.Println("版本:", mb.version)
	fmt.Printf("价格:%0.2f\n", mb.price)
}
func (mb Mobile) net() {
	fmt.Println("网络类型:", mb.nettype)
}
func Finterface() {
	phone = new(HuaWei)
	phone.call()
	phone = HuaWei{Mobile{"版本: meta10", 99999, "电信"}, "产于中国"}
	phone.call()
	phone = new(IPhone)
	phone.call()
	phone = IPhone{Mobile{"版本: 苹果9", 99999, "电信"}, "没啥可说"}
	phone = new(Mobile)
	phone.call()
	phone.net()

	vertu.call()
	vertu.net()
	phone = vertu
	phone.call()
	phone.net()
	huaweip8.call()
	huaweip8.net()
	phone = huaweip8
	phone.call()
	phone.net()

}

/*
int 0
int8 0
int32 0
int64 0
uint 0x0
rune 0 //rune的实际类型是 int32
byte 0x0 // byte的实际类型是 uint8
float32 0 //长度为 4 byte
float64 0 //长度为 8 byte
bool false
string ""
指针为nil
*/
func Finit() {
	var vari int
	var varf float64
	var varb bool
	var vars string
	var varby byte
	var varui uint
	var varpo *int
	fmt.Println("初始化:", vari, varf, varb, vars, varby, varui, varpo)
}

//返回多个值
func Freturn(str1, str2 string) (string, string, int) {
	return str2, str1, 2017
}

func Fconvert() {
	f := 22.22
	u := uint(f)
	fmt.Println(f, u)
}

/*在循环里面有两个关键操作 break 和 continue , break 操作是跳出当前循
环， continue 是跳过本次循环。当嵌套过深的时候， break 可以配合标签使
用，即跳转至标签所指定的位置*/
func Ffor() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	sum = 1
	for sum < 2000 {
		sum += sum
		if sum < 1500 {
			fmt.Println("continue")
			continue
		}
		if sum > 1800 {
			fmt.Println("break")
			break
		}

	}
	fmt.Println(sum)
}

func Fif(x int) {
	//if 语句也不要求用 ( ) 将条件括起来
	if x < 100 {
		fmt.Println(x * 10)
	} else {
		fmt.Println(x)
	}
	//if 语句可以在条件之前执行一个简单语句,在 if 的便捷语句定义的变量同样可以在任何对应的 else 块中使用。
	if v := x * 10; v < 10000 {
		fmt.Println("值太小了")
	} else {
		fmt.Println(v)
	}

}

func Fswicth() {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s", os)
	}
	fmt.Println(runtime.GOROOT())
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!", t)
	case t.Hour() < 17:
		fmt.Println("Good afternoon.", t)
	default:
		fmt.Println("Good evening.", t)
	}
}
func Fdefer() {
	defer fmt.Println("done")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	defer fmt.Print(" world\n")
	fmt.Print("hello ")
}
func FstructPoint() {
	p := &aboby
	fmt.Println(aboby)
	p.age = 20
	fmt.Println(*p, aboby, &*p, &aboby, &p)

}

//函数可以是一个闭包。闭包是一个函数值，它引用了函数体之外的变量
func FprivateFunc() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
func (m Mobile) Fstructclass(version string, price float32, nettype string) Mobile {
	m.version = version
	m.price = price
	m.nettype = nettype
	return m
}
func (m *Mobile) FpointParams(discount float32) {
	m.price = m.price * discount
}

//goroutine 是由 Go 运行时环境管理的轻量级线程
func Fgoroutine(s string) {
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s, "goroutine", i)
	}
}
func Fchanel(num int) {
	ch1 <- num
	ch2 <- num - 1
}
func Fstrings() {
	instr := `hello
				world`
	instr2 := instr + "2017"
	instr3 := "aAhello world"
	c := []byte(instr3)
	c[7] = '>'
	instr4 := "够浪阿"
	c2 := []byte(instr4)
	fmt.Println(instr, "\n", instr2, "\n", instr3, "\n", c, "\n", string(c), c2)
}
func Ferror() {
	err := errors.New("gou lang\n")
	if err != nil {
		fmt.Print(err)
	}
}

//用 goto 跳转到必须在当前函数内定义的标签
func Fgoto() {
	i := 0
Here:
	println(i)
	i++
	fmt.Println(i)
	if i < 10 {
		goto Here
	}

}

//变参
func Fvarparams(params ...string) {
	for _, str := range params {
		fmt.Println(str)
	}
}

/*传值与传指针,传指针比较轻量级 (8bytes),只是传内存地址，我们可以用指针传递体积大的结
构体。如果用参数值传递的话, 在每次copy上面就会花费相对较多的系统开销
（内存和时间）。所以当你要传递大的结构体的时候，用指针是一个明智的选
择。channel ， slice ， map 这三种类型的实现机制类似指针，所
以可以直接传递，而不用取地址后传递指针*/
func Fparamandadd(x int, y *int) {
	x = x + 100
	*y = *y + 100
}
func init() {
	stra := `Go里面有两个保留的函数： init 函数（能够应用于所有的 package ）
和 main 函数（只能应用于 package main ）。这两个函数在定义时不能有任何
的参数和返回值。虽然一个 package 里面可以写任意多个 init 函数，但这无论
是对于可读性还是以后的可维护性来说，我们都强烈建议用户在一个 package 中
每个文件只写一个 init 函数。Go程序会自动调用 init() 和 main() ，所以你不需要在任何地方调用这两个函
数。每个 package 中的 init 函数都是可选的，但 package main 就必须包含一个 main 函数。
程序的初始化和执行都起始于 main 包。如果 main 包还导入了其它的包，那么
就会在编译时将它们依次导入。有时一个包会被多个包同时导入，那么它只会被导
入一次（例如很多包可能都会用到 fmt 包，但它只会被导入一次，因为没有必要
导入多次）。当一个包被导入时，如果该包还导入了其它的包，那么会先将其它包
导入进来，然后再对这些包中的包级常量和变量进行初始化，接着执行 init 函数
（如果有的话），依次类推。等所有被导入的包都加载完毕了，就会开始
对 main 包中的包级常量和变量进行初始化，然后执行 main 包中的 init 函数
（如果存在的话），最后执行 main 函数。`
	fmt.Println(stra)
	var23 = "我是全局变量，在init函数被初始化"

}

/*method,method不是只能作用在struct上面，他可以定义在任何你自定
义的类型、内置类型、struct等各种类型上面。自定义类型申明如下：
type typeName typeLiteral
type ages int
type money float32
type months map[string]int
来实现。*/

type Rectangle struct {
	width, height float64
}
type Circle struct {
	radius float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}
func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

//method
func (h *Human) Finherit() {
	fmt.Printf("Hi, I am %s , %d years old and i have weight %d kg \n", h.name, h.age, h.weight)
}

//重载
func (e *Employee) Finherit() {
	fmt.Printf("Hi, I am %s , %d years old and i have weight %d kg \n", e.name, e.age, e.weight)
	fmt.Println("是", e.company, "的职员")
}

//动物接口方法
func (animal AnimalData) Move() {
	fmt.Println(animal.typename, "===Move")
}
func (animal AnimalData) Sound() {
	fmt.Println(animal.typename, "===Sound")
}
func (animal AnimalData) Eat() {
	fmt.Println(animal.typename, "===Eat")
}

//重载AnimalData的Move方法
func (dog DogData) Move() {
	fmt.Println(dog.typename, "===Move", "品种是:", dog.dogtye)
}

func Finterfaceexample() {
	var anm Animal
	pitter := AnimalData{"猫", 1, "母"}
	noob := DogData{AnimalData{"狗", 2, "公"}, "藏敖", "小菜"}
	niu := CowData{AnimalData{"牛", 5, "公"}, "黄色"}
	anm = pitter
	anm.Move()
	anm.Sound()
	anm.Eat()
	anm = noob
	anm.Move()
	anm.Sound()
	anm.Eat()
	anm = niu
	anm.Move()
	anm.Sound()
	anm.Eat()

}

//反射Go语言实现了反射，所谓反射就是能检查程序在运行时的状态
func Freflect() {
	i := 100
	t := reflect.TypeOf(i)
	v := reflect.ValueOf(&i)
	v.Elem().SetInt(200)
	t2 := reflect.TypeOf(aboby)
	v2 := reflect.ValueOf(aboby)
	fmt.Println(t, reflect.ValueOf(i))
	fmt.Println(t2, v2)
}

//果存在多个channel的时候，可过 select 可以监听channel上的数据流动。
func fibonacci(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
func Fselect() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
func Fpanic() {
	panic("error")
}

//入口函数
func main() {

	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			fmt.Println(err) // 这里的err其实就是panic传入的内容，55
		}
	}()
	//函数内定义的变量称为局部变量

	//变量名 := 值，出现在:=左侧的变量不应该是已经被声明过的，否则会导致编译错误，这种不带声明格式的只能在函数体中出现
	var1 = "golang"
	var4 := "begin"
	fmt.Println("begin")
	var34 = &var1
	println(var1, var2, var3, var4, var25[0], var26[0], Monday)
	println(var27, var28, var29, var30, var31, var32)
	Fpointer()
	amen.name = "milo"
	amen.age = 20
	amen.sex = "男"
	amen.height = 1.92
	agirl := Persons{"meta", 11, "女", 1.70}
	Fstruct(amen)
	Fstruct(agirl)
	Fstruct(aboby)
	Fslice(slice1)
	slice1 = []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	Fslice(slice1)
	slice1 = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	Fslice(slice1)
	FRange()
	FMap()
	Ffunc()
	fmt.Println(Frecursion(10))
	Finterface()
	fmt.Println(Freturn("wolrd", "hello"))
	Finit()
	Fconvert()
	Ffor()
	Fif(200)
	Fswicth()
	Fdefer()
	FstructPoint()
	gosum := FprivateFunc()
	for i := 0; i < 10; i++ {
		fmt.Println(gosum(i))
	}
	m := new(Mobile)
	fmt.Println(m.Fstructclass("诺基亚999", 9999, "移动"))
	m2 := &Mobile{"小米", 1000, "电信"}
	fmt.Println(*m2)
	m2.FpointParams(0.8)
	fmt.Println(*m2)
	go Fgoroutine("is ")
	Fgoroutine("not is ")
	go Fchanel(100)
	go Fchanel(200)
	go Fchanel(300)
	go Fchanel(400)
	fmt.Println(<-ch1, <-ch1, <-ch1, <-ch1)
	fmt.Println(<-ch2, <-ch2, <-ch2, <-ch2)
	Fstrings()
	Ferror()
	//Fgoto()
	Fvarparams("test string1", "test string2", "test string3", "test string4")
	parama := 100
	pointa := 100
	Fparamandadd(parama, &pointa)
	fmt.Println(parama, pointa)
	r1 := Rectangle{12, 2}
	r2 := Rectangle{9, 4}
	c1 := Circle{10}
	c2 := Circle{25}
	fmt.Println("Area of r1 is: ", r1.area())
	fmt.Println("Area of r2 is: ", r2.area())
	fmt.Println("Area of c1 is: ", c1.area())
	fmt.Println("Area of c2 is: ", c2.area())
	milo := Human{"Milo", 25, 70}
	mark := Student{Human{"Mark", 25, 70}, "computer"}
	sam := Employee{Human{"Sam", 45, 90}, "Golang Inc"}
	milo.Finherit()
	//方法继承
	mark.Finherit()
	//方法重载
	sam.Finherit()
	Finterfaceexample()
	Freflect()
	Fselect()
	//当我们不希望给函数起名字的时候，可以使用匿名函数
	func() {
		sum := 0
		for i := 1; i <= 100; i++ {
			sum += i
		}
		fmt.Println(sum)
	}()
	Fpanic()
	fmt.Println("end")

}

/*
GOPATH(bin,src,pkg)src go代码install的地方，pkg go代码install之后存放的地方，可以允许多个GOPATH存在
GOBIN 可执行的main pakage go代码install存放的地方
GOROOT go安装目录
go语言有一个获取远程包的工具就是 go get
go build这个命令主要用于编译代码。在包的编译过程中，若有必要，会同时编译与之相关联的包。
当你执行 go build 之后，它不会产生任何文件。如果你需要在 $GOPATH/pkg 下生
成相应的文件，那就得执行 go install 。
go clean这个命令是用来移除当前源码包和关联源码包里面编译生成的文件。
go fmt格式化代码
go test执行这个命令，会自动读取源码目录下面名为 *_test.go 的文件，生成并运行测试用的可执行文件。
go tool
go generate这个命令是从Go1.4开始才设计的，用于在编译前自动化生成某类代码。
go version 查看go当前的版本
go env 查看当前go的环境变量
go list 列出当前全部安装的package
go run 编译并运行Go程序

Go之所以会那么简洁，是因为它有一些默认的行为：
大写字母开头的变量是可导出的，也就是其它包可以读取的，是公用变量；小写字母开头的就是不可导出的，是私有变量。
大写字母开头的函数也是一样，相当于 class 中的带 public 关键词的公有
函数；小写字母开头的就是有 private 关键词的私有函数。
make 用于内建类型（ map 、 slice 和 channel ）的内存分配。 new 用于各种类型的内存分配。
*/
