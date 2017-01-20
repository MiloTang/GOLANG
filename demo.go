//包声明
package main

//引入包
import "fmt"

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
	huaweip8 = Mobile{"化为P8", 3000.00, "移动"}

	//Channel 类型

	//函数类型
	func1 FuncT

	/*切片类型,切片是对数组的抽象var identifier []type
	make([]T, length, capacity)
	*/
	slice1 []int
	slice2 = make([]int, 2, 3)
	slice3 = make([]string, 2)
	slice4 = []int{1, 2, 3}

	//接口类型（interface）
	phone Phone

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
	version string
	price   float32
	nettype string
}
type IPhone struct {
	version string
	price   float32
	nettype string
}
type Mobile struct {
	version string
	price   float32
	nettype string
}

//函数类型
type FuncT func(int, int)

//接口类型（interface）
type Phone interface {
	call()
	net()
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

//入口函数
func main() {

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
	fmt.Println("end")

}

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
	println(var23, var33, &var1, var34, *var34, &var34)
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
	huawei.version = "meta9"
	huawei.price = 10000.01
	fmt.Println("版本:", huawei.version)
	fmt.Printf("价格:%0.2f\n", huawei.price)
}

func (huawei HuaWei) net() {
	fmt.Println("网络类型", huawei.nettype)
}

func (iPhone IPhone) call() {
	fmt.Println("hello, Iphone")
	iPhone.version = "iphone7"
	iPhone.price = 7000.01
	fmt.Println("版本:", iPhone.version)
	fmt.Printf("价格:%0.2f\n", iPhone.price)
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
	phone = new(IPhone)
	phone.call()
	phone = new(Mobile)
	phone.call()
	phone.net()

	vertu.call()
	vertu.net()
	huaweip8.call()
	huaweip8.net()

}

/*
不同类型的局部和全局变量默认值为：int=0 float32=0 pointer=nil,string=空格 

*/
