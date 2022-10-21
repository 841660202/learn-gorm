package main

import (
	"flag"
	"fmt"
	"math/rand" //math包中的rand函数可以用来生成随机数
	"sort"
	"time"
)

func main() {
	// guess_random()

	// sum2array()

	// sumarray()

	// sum2ItemInArray()

	// sortArray()

	// autoMima()

	// range a-z

	// rangeaz()
	// deletemap()
	// printint1()
	// printint2()
	changeString()

}

func changeString() {
	s1 := "hello博客"
	// 强制类型转换
	byteS1 := []byte(s1) //  ASCII
	// 数组
	fmt.Println(byteS1)
	// 数组更改
	byteS1[0] = 'H'
	// 转化字符串
	fmt.Println(string(byteS1))

	s2 := "hello博客"
	runeS2 := []rune(s2) // UTF-8字符。
	fmt.Println(runeS2)
	runeS2[0] = '狗'
	fmt.Println(string(runeS2))
}

func printint2() {
	const a int8 = -1
	// var b int8 = -128 / a // 越界

	println(a)
	// println(b)
}

func printint1() {
	var a int8 = -1
	var b int8 = -128 / a

	println(a)
	println(b)
}

func deletemap() {

	var m map[string]int
	delete(m, "oh noes!")
	fmt.Println(m)
}

func rangeaz() {
	start := 'a'
	for i := 0; i < 26; i++ {
		fmt.Printf("%c ", start)
		start += 1
	}
}

func autoMima() {
	// 实现一个密码生成工具，支持以下功能：

	// a) 用户可以通过-l指定生成密码的长度

	// b) 用户可以通过-t指定生成密码的字符集，比如-t num生成全数字的密码

	// -t char 生成包含全英文字符的密码， -t mix包含生成数字和英文的密码，

	// -t advance 生成包含数字、英文以及特殊字符的密码

	// 提示：可以用标准包 “flag”解析命令行参数

	// 思路：

	// 1.先解析命令行参数；

	// 2.定义变量接收用户选择；

	// 3.生成随机密码for循环
	var (
		numCharset     = "0123456789"
		strCharset     = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
		mixCharset     = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
		advanceCharset = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ`~!@#$%^&*()-_+=|?"
	)
	var length int                                           //定义一个密码长度的变量
	var charset string                                       //定义一个密码字符集的变量
	flag.IntVar(&length, "l", 16, "-l the length of passwd") //因为int是值类型，所以终端输入传参必须要用&取地址，
	//IntVar参数依次是：传入值-终端参数-默认值-用法
	flag.StringVar(&charset, "t", "harden", "-t the charset of password")
	//StringVar参数依次是：传入值-终端参数-默认值-用法
	flag.Parse() //真正解析命令行参数，IntVar、StringVar只是设置命令行需要的一些参数

	//判断用户输入的到底是什么？switch和if else都可以
	var userCharset string //定义一个变量来接收用户选择
	switch charset {
	case "num":
		userCharset = numCharset
	case "char":
		userCharset = strCharset
	case "mix":
		userCharset = mixCharset
	case "advance":
		userCharset = advanceCharset
	default:
		userCharset = mixCharset
	} //字符集功能编写完毕

	//下面写生成随机密码最终部分
	var password []byte
	rand.Seed(time.Now().UnixNano()) //设置随机总值，这样就避免了随机密码重复
	for i := 0; i < length; i++ {    //想要几位密码就循环几次
		index := rand.Intn(len(userCharset)) //生成随机下标
		char := userCharset[index]           //现在是字符，最终要存到一个字符串中
		password = append(password, char)    //字符串中的字符是不能修改的，想要修改，需要将其放入切片，所以将一个个字符放入到切片中，切片是可以修改的

	}
	strPassword := string(password) //将字符切片强制转换成字符串
	fmt.Printf("%s\n", strPassword)
}

func sortArray() {
	var arr [10]int //使用一个随机的初始化数组
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(10000)
	}
	fmt.Printf("排序前 arr:%v\n", arr) //排序前
	sort.Ints(arr[:])               //Ints是升序

	fmt.Printf("排序后 arr:%v\n", arr) //排序后
}

// 题目：找出数组中和为给定值的两个元素的下标，比如数组:[1,3,5,8,7]，找出两个元素之和等于8的下标分别是(0, 4)和(1,2)。
func sum2ItemInArray() {
	var arr [10]int
	for i := 0; i < len(arr); i++ {
		arr[i] = i //数组初始化从0-9（因为是试题所以这里为了方便索引和元素值都是0-9）
	}

	var sum int = 12                //求两个元素之和为12的所有数组下标
	for i := 0; i < len(arr); i++ { //思路就是i=0 和1-9的下标加一遍，和为12即满足条件，其余下标也同理，这里写2个for循环即可
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == sum { //如果元素和为12则打印
				fmt.Printf("i=%d j=%d \n", i, j)
			}
		}
	}
}

func sumarray() {
	// panic("unimplemented")
	var arr [10]int
	rand.Seed(time.Now().UnixNano()) //通过rand函数的Seed方法，来设定随机数总值
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(10000) //数组取值为随机数，随机数范围为[1-10000)
	}

	var sum int
	for i := 0; i < len(arr); i++ {
		sum = sum + arr[i]
	}
	fmt.Println("sum", sum)
}

func guess_random() {
	var number int
	/*
	   for i := 0;i < 10; i++ {
	       number = rand.Intn(100)
	       fmt.Printf("number:%dn",number)
	   }
	   因为伪随机的存在，rand生成的随机数都有一个固定的序列，比如第一次生成1，第二次生成20等等，随机数的序列是固定，我们可以通过给其一个随机的总值，如果随机的总值不一样，那么伪随机数的序列也不一样。我们不给，他就是按照一个默认的随机总值，所以伪随机数序列永远是一样的。
	*/
	rand.Seed(time.Now().UnixNano()) //通过rand函数的Seed方法，来设置总值，这里我们以当前时间来设置总值，并且用的纳秒，十分精确了
	number = rand.Intn(100)          //随机数的范围是0-100，但不包括100
	fmt.Printf("请猜一个数字，数字的范围是：[0-100), n = ")
	for { //因为并不知道用户什么时候输入正确，所以是没有限制条件的，我们只能做一个死循环，配合switch语句满足条件了即可退出；
		var input int
		fmt.Scanf("%dn", &input) //Scanf表示让用户输入,Scanf从终端读取一个整数，并传值给input变量，&表示获取到该变量内存地址
		var flag bool = false    //通过设置flag变量，解决的是用户输入正确后可以退出
		switch {
		case number > input:
			fmt.Printf("您输入的数字太小, n = ")
		case number == input:
			fmt.Printf("恭喜您，答对了！%v", number)
			flag = true
		case number < input:
			fmt.Printf("您输入的数字太大,n = ")
		}
		if flag { //表示如果flag为真，则break退出这个for循环
			break
		}
	}
}

// 题目：求数组所有元素之和
func sum2array() {
	a := [3]int{1, 5, 10}
	sum := 0
	for i := 0; i < len(a); i++ {
		sum += a[i]
	}
	fmt.Println(sum)
}
