//包名，入口包名必须为main
package main

/**导入包，可以是自带的包，也可以是客户化的包,在 Golang 包中，
一个名字如果首字母大写则表示此名字被导出(包括变量，函数等)
**/
import (
	"bemashow/common" //自定义包
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

/**
声明变量
基本数据类型：
1.bool
2.string
3.int int8 int16 int32 int64
4.uint uint8 uint16 uint32 uint64
5.uintptr
6.byte（等价于 uint8）
7.rune（等价于 int32，用于表示一个 unicode code point）
8.float32 float64
9.complex64 complex128
需要注意的就是，变量名在类型之前，这和很多语言都不一样。
按照 Golang 的语法，在函数外的任何结构（construct）都通过一个关键字开始，
例如变量使用 var 关键字开始，函数使用 func 关键字开始。
在函数内，变量赋值可以使用 := 操作符
**/
var (
	sum int
	i   int
	f   float64
)

/**
常量声明
、**/
const (
	Big   = 1 << 100
	Small = Big >> 99
	Pi    = 3.1415926
)

/**
其他数据类型
**/
type Employee struct {
	Name   string
	Age    int8
	Height float32
	Years  int8
}
type Student struct {
	Name   string
	Age    int8
	Height float32
}
type Myint int64

var (
	e   *Employee
	es  [10]Employee
	ess []Employee
	esm map[string]Employee
)

/**
接口也是一种类型（就像结构体一样）。一个接口类型包含了一组方法，一个接口类型能够持有那些实现了这些方法的值。
**/
type Human interface {
	ShowInfo()
	Say()
	Work()
}

/**
结构体中可以存在只有类型而没有名字的域，它们被叫做匿名域。
一个结构体的匿名域中的域或者方法可以被此结构体实例直接访问
**/
type Car struct {
	string
}

/**
使用 channel 和 goroutine 通讯能够避免显式使用锁机制，
通过 channel 发送和接收值时默认是阻塞的。
channel 可以带有一个缓冲区（buffer）来缓存被传递的值，向 channel
中发送时只有缓冲区满的情况下会阻塞，接收 channel 中的值时只有在缓冲区空的情况下阻塞
通过 make 函数创建 channel
**/
var ch = make(chan string, 3)

/**
初始化函数，可对数据进行初始化，在main方法之前运行
**/
func init() {
	fmt.Println("init")
	i = 1
	sum = 1
}
func main() {
	fmt.Println("begin main")
	Statement()
	s := Func()
	fmt.Println(s("test"))
	e = &Employee{"Baidu", 18, 1.80, 5} //指针初始化
	e.ShowInfo()
	e.Say()
	emp := new(Employee)
	emp.Name = "Ben"
	emp.Age = 30
	emp.Height = 1.22
	emp.Years = 10
	InterFaceFunc(emp)
	st := new(Student)
	st.Name = "A Ben"
	st.Age = 29
	st.Height = 1.32
	InterFaceFunc(st)
	myint := Myint(2222)
	myint.Int()
	fmt.Println(Car{"Ferrari"})
	GoRoutineChanel()
	fmt.Println(common.Sum())
	fmt.Println("end main")
}

func GoRoutineChanel() {
	go Goroutine("string a")
	go Goroutine("string b")
	go Goroutine("string c")
	fmt.Println("take out from", <-ch)
	fmt.Println("take out from", <-ch)
	fmt.Println("take out from", <-ch)
	c1 := make(chan int, 10)
	go fibonacciA(cap(c1), c1)
	// 这里 for 和 range 组合使用
	// 不断的接收 c 中的值一直到它被关闭
	for i = range c1 {
		fmt.Println(i)
	}
	c2 := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c2)
		}
		quit <- 0
	}()
	fibonacciB(c2, quit)
	Select()
}

/**
函数定义，函数可以返回多个值,
func name(x, y string) (string, string) {
    return y, x
}
定义方式：func [类型] 名字([传值类型]) [类型]{}第一个类型非基础类型
**/
func Statement() {
	/**
	类型转换
	**/
	f = float64(i)
	/**
	switch 语句,Golang 中无需使用 break 语句来跳出 switch,另外，switch 可以没有条件
	**/
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("System of OS X")
	case "linux":
		fmt.Println("System of Linux")
	case "windows":
		fmt.Println("System of Windows")
	default:
		fmt.Printf("%s.", os)
	}
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
	/**
	if 语句可以在执行条件判断前带一个语句（这常被叫做 if 带上一个短语句），
	此语句中变量的生命周期在 if 语句结束后结束。
	**/
	rand.Seed(time.Now().UnixNano())
	if n := rand.Intn(1000); n <= 500 {
		fmt.Println(n, " 点,买小的赢")
	} else {
		fmt.Println(n, " 点,买大的赢")
	}
	/**
	使用 for 语句时不需要 () 并且 {} 是必须的,golang没有 while 语句，
	但是可以用for来实现while 语句
	**/
	for sum < 100 {
		sum += sum
	}
	for ; i < 100; i++ {
		sum += sum
	}
	/**
	一个 defer 语句能够将一个函数调用加入一个列表中
	（这个函数调用被叫做 deferred 函数调用），在当前函数调用结束时调用列表中的函数。
	**/
	for i = 0; i < 2; i++ {
		defer fmt.Println(i)
	}
	/**
	range 被用在 for 中来迭代一个array,slice 或者一个 map：
	**/
	for _, em := range es {
		fmt.Println(em)
	}
	fmt.Println("==================怎么能少得了华丽的分割线呢=====================")
	fmt.Println(ess, len(ess), cap(ess))
	/**
	初始化切片,这里 make 函数会创建一个数组（其元素初始化为 0）并返回一个 slice 指向此数组。
	make 可以带第三个参数，用于指定容量,切片可以不初始化而直接赋值，但是map就必须先要初始化才能
	赋值,可以使用append函数来给某个切片（slice）追加元素，可以用delete(mapname,key)来删除
	map
	**/
	ess = make([]Employee, 0, 10)
	ess = []Employee{{"Milo", 18, 1.80, 5}}
	ess = append(ess, Employee{"Ben", 18, 1.80, 5})
	for _, em := range ess {
		fmt.Println(em)
	}
	fmt.Println(ess, len(ess), cap(ess))
	fmt.Println("==================怎么能少得了华丽的分割线呢=====================")
	esm = make(map[string]Employee)
	esm["006"] = Employee{"Milo", 18, 1.80, 5}
	esm["007"] = Employee{"Ben", 18, 1.80, 5}
	for k, v := range esm {
		fmt.Println(k, v)
	}

}

/**
闭包Golang 中函数也是一个值（就像 int 值一样），且函数可以是一个闭包。闭包是一个引用了外部变量的函数。
**/
func Func() func(f string) string {
	fmt.Println(f)
	return func(f string) string {
		fmt.Println(f)
		return f
	}
}

/** 这里的方法接收者（method receiver）e 的类型为 *Employee
在这里方法的接收者使用指针类型而非值类型主要出于以下几点考虑（类似 C/C++ 等语言）：

1.避免方法每次调用时，对接收者的不必要的拷贝
2.在方法内可以修改接收者的值

我们可以为任意类型定义方法，但以下情况除外：

1.如果类型定义在其他包中，不能为其定义方法
2.如果类型是基础类型，不能为其定义方法
**/
func (e *Employee) ShowInfo() {
	fmt.Println("名字=>", e.Name)
	fmt.Println("年龄=>", e.Age)
	fmt.Println("身高=>", e.Height)
	fmt.Println("工作年限=>", e.Years)
}
func (e *Employee) Say() {
	fmt.Println(e.Name, " say hello")
}
func (e *Employee) Work() {
	fmt.Println(e.Name, " work hard")
}
func (s *Student) ShowInfo() {
	fmt.Println("名字=>", s.Name)
	fmt.Println("年龄=>", s.Age)
	fmt.Println("身高=>", s.Height)

}
func (s *Student) Say() {
	fmt.Println(s.Name, " say hello")
}
func (s *Student) Work() {
	fmt.Println(s.Name, " study hard")
}
func InterFaceFunc(h Human) {
	h.Say()
	h.ShowInfo()
	h.Work()
}
func (i Myint) Int() {
	fmt.Println(i)
}

/**
Golang 运行时（runtime）管理了一种轻量级线程，被叫做 goroutine。
我们使用 channel 和 goroutine 通讯。channel 中是一种带有类型的通道，
被用于接收和发送特定类型的值。操作符 <- 被叫做
channel 操作符（这个操作符中箭头表明了值的流向）
**/
func Goroutine(s string) {
	s = s + " in ch"
	fmt.Println(s)
	ch <- s
}
func fibonacciA(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	// 必须要关闭 c通常来说，我们不需要主动的关闭 channel。但有时候接收者必须被告知已经没有值可以接收了，
	//这时候主动关闭是必要的，例如终止 for range 循环
	close(c)
}
func fibonacciB(c, quit chan int) {
	x, y := 0, 1
	//使用 select 语句可以让一个 goroutine 等待多个通讯操作。
	//select 会阻塞直到某个 case 能够运行，如果同时存在多个可执行的，那么将随机选择一个
	//select 中的 default 会在没有任何 case 可执行时执行
	for {
		select {
		case c <- x:
			x, y = y, x+y
		// 控制此线程退出
		case <-quit:
			fmt.Println("quit")
			return
		}
	}

}
func Select() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(200 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("100")
		case <-boom:
			fmt.Println("200")
			return
		default:
			fmt.Println("50")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

/**
GOPATH(bin,src,pkg)src go代码install的地方，pkg go代码install之后存放的地方，可以允许多个GOPATH存在
GOBIN 可执行的main pakage go代码install存放的地方
GOROOT go安装目录
go语言有一个获取远程包的工具就是 go get
go build这个命令主要用于编译代码。在包的编译过程中，若有必要，会同时编译与之相关联的包。
go build --gcflags=-m
当你执行 go build 之后，它不会产生任何文件。如果你需要在 $GOPATH/pkg 下生
成相应的文件，那就得执行 go install 。
go clean这个命令是用来移除当前源码包和关联源码包里面编译生成的文件。
go fmt格式化代码
go test执行这个命令，会自动读取源码目录下面名为 *_test.go 的文件，生成并运行测试用的可执行文件。
go tool
go tool compile -S *.go >> *.s生成汇编代码

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

**/
/**

**/
