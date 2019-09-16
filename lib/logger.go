package lib

import (
"fmt"
"log"
"os"
"time"
)

func LogHander(logs string,err error) {
	file:="./" + time.Now().Format("20160102") + ".log"

	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if nil != err {
		panic(err)
		//return false
	}

	//创建一个Logger
	//参数1：日志写入目的地
	//参数2：每条日志的前缀
	//参数3：日志属性
	loger := log.New(logFile, "operation", log.Ldate|log.Ltime|log.Lshortfile)


	//Flags返回Logger的输出选项
	//fmt.Println(loger.Flags())
	//fmt.Println("************")
	//SetFlags设置输出选项
	loger.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)


	//返回输出前缀
	//fmt.Println(loger.Prefix())
	//loger.Println("dsfsdfsdf")
	loger.Println(logs)
	loger.Println(err)
	loger.Println("------------------fey-------------------------")

	//设置输出前缀
	//loger.SetPrefix("test_")


	//输出一条日志
	//loger.Output(2, "打印一条日志信息")

	//格式化输出日志
	//loger.Printf("第%d行 内容:%s", 11, "我是错误k")


	// //等价于print();panic();
	// loger.Panic("我是错误5")


	// //等价于print();os.Exit(1);
	// loger.Fatal("我是错误1")

	//log的导出函数
	//导出函数基于std,std是标准错误输出
	//var std = New(os.Stderr, "", LstdFlags)


	//获取输出项
	//fmt.Println(log.Flags())
	//获取前缀
	//fmt.Printf(log.Prefix())
	//输出内容
	//loger.Output(2, "输出内容")
	//格式化输出
	//loger.Printf("第%d行 内容:%s", 22, "我是错误")
	//loger.Fatal("我是错误3")
	//loger.Panic("我是错误4")

	//return true
}

func InfoHander(logs string) {
	file:="./" + time.Now().Format("20160102") + ".log"

	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if nil != err {
		panic(err)
		//return false
	}

	//创建一个Logger
	//参数1：日志写入目的地
	//参数2：每条日志的前缀
	//参数3：日志属性
	loger := log.New(logFile, "operation", log.Ldate|log.Ltime|log.Lshortfile)


	//Flags返回Logger的输出选项
	//SetFlags设置输出选项
	loger.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	//返回输出前缀
	loger.Println(logs)
	loger.Println("------------------info-------------------------")

}

// 定义一个 DivideError 结构
type DivideError struct {
	dividee int
	divider int
}

// 实现 `error` 接口
func (de *DivideError) Error() string {
	strFormat := `
    Cannot proceed, the divider is zero.
    dividee: %d
    divider: 0
`
	return fmt.Sprintf(strFormat, de.dividee)
}

// 定义 `int` 类型除法运算的函数
func Divide(varDividee int, varDivider int) (result int, errorMsg string) {
	if varDivider == 0 {
		dData := DivideError{
			dividee: varDividee,
			divider: varDivider,
		}
		errorMsg = dData.Error()
		return
	} else {
		return varDividee / varDivider, ""
	}

}

