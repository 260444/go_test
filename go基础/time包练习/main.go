package main

import (
	"fmt"
	"time"
)

func main() {
	//定义一个开始时间并转换为时间戳毫秒级
	start := time.Now()
	startmicro := start.UnixMicro()

	now := time.Now()
	fmt.Printf("now: %v\n", now)
	year := now.Year() //年
	fmt.Printf("year: %v\n", year)
	month := now.Month() //月
	fmt.Printf("month: %v\n", month)
	day := now.Day() //日
	fmt.Printf("day: %v\n", day)
	hour := now.Hour() //小时
	fmt.Printf("hour: %v\n", hour)
	minute := now.Minute() //分钟
	fmt.Printf("minute: %v\n", minute)
	second := now.Second() //秒
	fmt.Printf("second: %v\n", second)
	fmt.Println("===============================")

	//unix time
	seconds := now.Unix() // 获得秒级时间戳
	fmt.Printf("Seconds: %v\n", seconds)
	milli := now.UnixMilli() // 获得毫秒级时间戳
	fmt.Printf("milli: %v\n", milli)
	micro := now.UnixMicro() // 获得微秒级时间戳
	fmt.Printf("micro: %v\n", micro)
	nano := now.UnixNano() // 获得纳秒级时间戳
	fmt.Printf("nano: %v\n", nano)
	fmt.Println("===============================")

	//将时间戳转换为时间(第二个参数为不足一苗的纳秒数)
	t := time.Unix(1662187811, 0)
	fmt.Printf("t: %v\n", t)
	mill2 := time.UnixMilli(1662187992650) //毫秒级转化为时间对象
	fmt.Printf("mill2: %v\n", mill2)
	micro2 := time.UnixMicro(1662187992650750) //微妙级转化为时间对象
	fmt.Printf("micro2: %v\n", micro2)
	fmt.Println("===============================")

	//时间间隔
	// time.Sleep(2 * time.Second) //睡眠两秒
	fmt.Println("sleep over!!!")
	// n := 2 //type int
	// time.Sleep(time.Duration(n) * time.Second)
	fmt.Println("sleep2 over!!!")
	fmt.Println("===============================")

	//时间操作
	afterhour := now.Add(time.Hour)
	fmt.Printf("afterhour: %v\n", afterhour)
	//求两个时间的差值
	s := afterhour.Sub(now)
	fmt.Printf("s: %v\n", s)
	// 第一个时间和第二个时间是否相等返回一个bool
	b := time.Now().Equal(afterhour)
	b2 := afterhour.Equal(afterhour)
	fmt.Printf("b: %v\n", b)
	fmt.Printf("b2: %v\n", b2)
	// 第一个时间在第二个时间之前返回一个bool
	b = time.Now().Before(afterhour)
	fmt.Printf("b: %v\n", b)

	// 第一个时间在第二个时间之后返回一个bool
	b = time.Now().After(afterhour)
	fmt.Printf("b: %v\n", b)
	fmt.Println("===============================")

	//定时器
	// c := time.Tick(2 * time.Second) //每隔两秒秒计时一次
	// c2 := time.Tick(2 * time.Millisecond) //每隔两毫秒计时一次
	// for v := range c2 {
	// 	fmt.Println(v)
	// }
	fmt.Println("===============================")

	// 时间格式化
	// 制定布局的文本表示形式为：2006-01-02 15:04:05.0000
	/////////////////////////// 年-月-日 时:分:秒
	// 如果想格式化为12小时制15改为03末尾加 PM

	fmt.Println(now.Format("2006.01.02 15:04"))      //24小时制
	fmt.Println(now.Format("2006.01.02 03:04 PM"))   //12小时制
	fmt.Println(now.Format("2006.01.02 15:04.000 ")) //小数点后有三个零表示结果保留三位小数
	fmt.Println(now.Format("2006.01.02 15:04.999 ")) //小数点后有三个9表示结果省略末尾出现的零

	//解析字符串时间
	// 在没有指定时区的情况下默认返回UTC时间
	t2, err := time.Parse("2006-01-02 15:04:05", "2099-01-01 15:15:20")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t2)

	// 获得时区
	loc, err2 := time.LoadLocation("Asia/Shanghai")
	if err2 != nil {
		fmt.Println(err2)
		return
	}

	timeObj, err3 := time.ParseInLocation("2006-01-02 15:04:05", "2099-01-01 15:15:20", loc)
	if err3 != nil {
		fmt.Println(err3)
		return
	}
	fmt.Println(timeObj)
	fmt.Println(timeObj.Sub(now))

	// 定义一个结束时间并转换为微秒级时间戳
	end := time.Now()
	endmicro := end.UnixMicro()

	// 计算程序运行的了多少微妙
	timeall := endmicro - startmicro
	fmt.Printf("执行了%v微秒\n", timeall)

	//计算开始到程序结束多少毫秒
	a := time.Since(start)
	fmt.Printf("a: %v\n", a)
	// 计算程序运行的了多少微妙
	i := int(a) / 1000
	fmt.Printf("i: %v\n", i)

	fmt.Println(time.Now().Year())
}
